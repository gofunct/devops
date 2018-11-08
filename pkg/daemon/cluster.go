package daemon

import (
	"context"
	"encoding/json"
	"sort"
	"strconv"
	"strings"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/metadata"
)

// Node ...
type Node struct {
	ID     int
	Addr   string                            `json:"addr"`
	Client mnemosynerpc.SessionManagerClient `json:"-"`
	Health grpc_health_v1.HealthClient       `json:"-"`
}

// Cluster ...
type Cluster struct {
	listen  string
	buckets int
	nodes   []*Node
	logger  *zap.Logger
}

// Opts ...
type Opts struct {
	Listen string
	Seeds  []string
	Logger *zap.Logger
}

// New ...
func New(opts Opts) (csr *Cluster, err error) {
	var (
		nodes  []string
		exists bool
	)
	for _, seed := range opts.Seeds {
		if seed == opts.Listen {
			exists = true
		}
	}
	if !exists {
		nodes = append(nodes, opts.Listen)
	}

	nodes = append(nodes, opts.Seeds...)
	sort.Strings(nodes)

	csr = &Cluster{
		nodes:  make([]*Node, 0),
		listen: opts.Listen,
		logger: opts.Logger,
	}

	for i, addr := range nodes {
		if addr == "" {
			continue
		}
		csr.buckets++
		csr.nodes = append(csr.nodes, &Node{
			ID:   i,
			Addr: addr,
		})
	}
	return csr, nil
}

// Connect ...
func (c *Cluster) Connect(ctx context.Context, opts ...grpc.DialOption) error {
	for i, n := range c.nodes {
		if n.Addr == c.listen {
			continue
		}

		if c.logger != nil {
			c.logger.Debug("cluster node attempt to connect", zap.String("address", n.Addr), zap.Int("index", i))
		}

		conn, err := grpc.DialContext(ctx, n.Addr, opts...)
		if err != nil {
			return err
		}

		if c.logger != nil {
			c.logger.Debug("cluster node connection success", zap.String("address", n.Addr), zap.Int("index", i))
		}

		n.Client = mnemosynerpc.NewSessionManagerClient(conn)
		n.Health = grpc_health_v1.NewHealthClient(conn)
	}

	return nil
}

// Get if possible returns node for a given bucket id.
func (c *Cluster) Get(k int32) (*Node, bool) {
	if len(c.nodes) == 0 {
		return nil, false
	}
	if len(c.nodes)-1 < int(k) {
		return nil, false
	}
	return c.nodes[k], true
}

// Nodes returns all available nodes.
func (c *Cluster) Nodes() []*Node {
	return c.nodes
}

// ExternalNodes returns all available nodes except host.
func (c *Cluster) ExternalNodes() (res []*Node) {
	for _, n := range c.nodes {
		if n.Addr != c.listen {
			res = append(res, n)
		}
	}
	return
}

// Len returns number of nodes.
func (c *Cluster) Len() int {
	return c.buckets
}

// Listen returns address of current node.
func (c *Cluster) Listen() string {
	return c.listen
}

// GetOther returns node for given access token.
// Returns false if cluster is nil, has only one element or if node that was found has same listen address as current one.
func (c *Cluster) GetOther(accessToken string) (*Node, bool) {
	if c == nil {
		return nil, false
	}
	if c.Len() == 1 {
		return nil, false
	}

	if node, ok := c.Get(jump.HashString(accessToken, c.Len())); ok {
		if node.Addr != c.listen {
			if node.Client != nil {
				return node, true
			}
		}
	}
	return nil, false
}

// GoString implements fmt GoStringer interface.
func (c *Cluster) GoString() string {
	buf, _ := json.Marshal(map[string]interface{}{
		"listen":  c.listen,
		"nodes":   c.nodes,
		"buckets": strconv.FormatInt(int64(c.buckets), 10),
	})
	return string(buf)
}

func IsInternalRequest(ctx context.Context) bool {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if ua, ok := md["user-agent"]; ok {
			if len(ua) > 0 {
				return strings.Contains(ua[0], Subsystem)
			}
		}
	}
	return false
}