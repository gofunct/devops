package daemon

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"encoding/json"

	"sync"

	"go.uber.org/zap"
	"google.golang.org/grpc/health/grpc_health_v1"
)

type livenessResponse struct {
	Version string `json:"version"`
}

type livenessHandler struct {
	livenessResponse

	logger *zap.Logger
}

func (lh *livenessHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(rw).Encode(lh.livenessResponse); err != nil {
		http.Error(rw, fmt.Sprintf("liveness check response encode failure: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	lh.logger.Debug("liveness check done")
	rw.WriteHeader(http.StatusOK)
}

type readinessResponse struct {
	sync.Mutex `json:"-"`

	livenessResponse

	Probes struct {
		Postgres probeStatus            `json:"postgres"`
		Cluster  map[string]probeStatus `json:"cluster"`
	} `json:"probes"`
}

type readinessHandler struct {
	livenessResponse

	logger   *zap.Logger
	postgres *sql.DB
	cluster  *Cluster
}

func (rh *readinessHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	var (
		wg     sync.WaitGroup
		status readinessResponse
	)

	status.livenessResponse = rh.livenessResponse
	status.Probes.Cluster = make(map[string]probeStatus, rh.cluster.Len())

	ctx, cancel := context.WithTimeout(r.Context(), 20*time.Second)
	defer cancel()

	wg.Add(1)
	go func() {
		defer wg.Done()

		if rh.postgres != nil {
			if err := rh.postgres.PingContext(ctx); err != nil {
				status.Probes.Postgres = probeStatus(err.Error())
			}
		}
	}()

	for _, n := range rh.cluster.ExternalNodes() {
		wg.Add(1)

		go func(n *Node) {
			defer wg.Done()

			res, err := n.Health.Check(ctx, &grpc_health_v1.HealthCheckRequest{})

			status.Lock()
			defer status.Unlock()

			if err != nil {
				status.Probes.Cluster[n.Addr] = probeStatus(err.Error())
			} else {
				status.Probes.Cluster[n.Addr] = probeStatus(res.Status.String())
			}
		}(n)
	}

	wg.Wait()

	if err := json.NewEncoder(rw).Encode(status); err != nil {
		http.Error(rw, fmt.Sprintf("readiness check response encode failure: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	rh.logger.Debug("readiness check done", zap.Any("probes", status.Probes), zap.Int("nb_of_external_nodes", len(rh.cluster.ExternalNodes())))
}

type probeStatus string

func (ps probeStatus) IsOK() bool {
	return ps == ""
}

func (ps probeStatus) MarshalJSON() ([]byte, error) {
	if ps.IsOK() {
		return json.Marshal(grpc_health_v1.HealthCheckResponse_SERVING.String())
	}

	return json.Marshal(string(ps))
}