package function

import (
	"encoding/json"
	"errors"
	"net"
	"os"
	"time"

	"github.com/ymotongpoo/datemaki"
)

type Request struct {
	Format string
	Server string
	Query  string
	Start  string
	End    string
	Step   string
}

func NewRequest(data []byte) (*Request, error) {
	r := &Request{}

	if err := json.Unmarshal(data, r); err != nil {
		// not json, use input as raw PromQL string
		r.Query = string(data)
	}

	promURL := os.Getenv("PROMETHEUS_URL")
	if promURL == "" {
		_, err := net.LookupIP("prometheus")
		if err != nil {
			promURL = "http://prometheus.openfaas:9090"
		} else {
			promURL = "http://prometheus:9090"
		}
	}
	if r.Server == "" && len(promURL) > 0 {
		r.Server = promURL
	}

	if r.Server == "" {
		return nil, errors.New("no Prometheus server specified")
	}

	if r.Query == "" {
		return nil, errors.New("no query specified")
	}

	if r.Start == "" {
		r.Start = "1 hour ago"
	}

	if r.End == "" {
		r.End = "now"
	}

	if r.Step == "" {
		r.Step = "1m"
	}

	if r.Format != "table" && r.Format != "json" {
		r.Format = "table"
	}
	return r, nil
}

func (r *Request) GetQueryRange() (start time.Time, end time.Time, step time.Duration, err error) {
	if start, err = datemaki.Parse(r.Start); err != nil {
		return
	}

	if end, err = datemaki.Parse(r.End); err != nil {
		return
	}

	if step, err = time.ParseDuration(r.Step); err != nil {
		return
	}

	return
}
