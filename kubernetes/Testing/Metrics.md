## Constants

```go

const (
    ProxyTimeout = 2 * time.Minute
)

```

***

## ApiServerMetrics

* type ApiServerMetrics Metrics
* func NewApiServerMetrics

* func NewApiServerMetrics() ApiServerMetrics
* func (*ApiServerMetrics) Equal

* func (m *ApiServerMetrics) Equal(o ApiServerMetrics) bool





## ClusterAutoscalerMetrics

* func NewClusterAutoscalerMetrics() ClusterAutoscalerMetrics
* func (m *ClusterAutoscalerMetrics) Equal(o ClusterAutoscalerMetrics) bool

## type ControllerManagerMetrics
* func NewControllerManagerMetrics() ControllerManagerMetrics
* func (m *ControllerManagerMetrics) Equal(o ControllerManagerMetrics) bool
## type KubeletMetrics
func GrabKubeletMetricsWithoutProxy(nodeName string) (KubeletMetrics, error)
func NewKubeletMetrics() KubeletMetrics
func (m *KubeletMetrics) Equal(o KubeletMetrics) bool
type Metrics
func NewMetrics() Metrics
func (m *Metrics) Equal(o Metrics) bool
type MetricsCollection
type MetricsGrabber
func NewMetricsGrabber(c clientset.Interface, ec clientset.Interface, kubelets bool, scheduler bool, controllers bool, apiServer bool, clusterAutoscaler bool) (*MetricsGrabber, error)
func (g *MetricsGrabber) Grab() (MetricsCollection, error)
func (g *MetricsGrabber) GrabFromApiServer() (ApiServerMetrics, error)
func (g *MetricsGrabber) GrabFromClusterAutoscaler() (ClusterAutoscalerMetrics, error)
func (g *MetricsGrabber) GrabFromControllerManager() (ControllerManagerMetrics, error)
func (g *MetricsGrabber) GrabFromKubelet(nodeName string) (KubeletMetrics, error)
func (g *MetricsGrabber) GrabFromScheduler() (SchedulerMetrics, error)
func (g *MetricsGrabber) HasRegisteredMaster() bool
type SchedulerMetrics
func NewSchedulerMetrics() SchedulerMetrics
func (m *SchedulerMetrics) Equal(o SchedulerMetrics) bool
