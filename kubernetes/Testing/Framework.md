# Package Framework Documentation

**Largely Based on the Kubernetes Testing Package**

```go
package framework
import "k8s.io/kubernetes/test/e2e/framework"
```


**Package framework contains provider-independent helper code for building and running E2E tests with Ginkgo. The actual Ginkgo test suites gets assembled by combining this framework, the optional provider support 
code and specific tests via a separate .go file like Kubernetes' test/e2e.go.**


## Constants

<details><summary>show</summary>
<p>



```go
const (
       // How long to wait for a job to finish.
        JobTimeout = 15 * time.Minute
   
       // Job selector name
        JobSelectorKey = "job"
        EndpointHttpPort      = 8080
        EndpointUdpPort       = 8081
        TestContainerHttpPort = 8080
        ClusterHttpPort       = 80
        ClusterUdpPort        = 90
    // Number of checks to hit a given set of endpoints when enable session affinity.
        SessionAffinityChecks = 10
        // KubeProxyLagTimeout is the maximum time a kube-proxy daemon on a node is allowed
            // to not notice a Service update, such as type=NodePort.
            // TODO: This timeout should be O(10s), observed values are O(1m), 5m is very
            // liberal. Fix tracked in #20567.
            KubeProxyLagTimeout = 5 * time.Minute
        
            // KubeProxyEndpointLagTimeout is the maximum time a kube-proxy daemon on a node is allowed
            // to not notice an Endpoint update.
            KubeProxyEndpointLagTimeout = 30 * time.Second
        
            // LoadBalancerLagTimeoutDefault is the maximum time a load balancer is allowed to
            // not respond after creation.
            LoadBalancerLagTimeoutDefault = 2 * time.Minute
        
            // LoadBalancerLagTimeoutAWS is the delay between ELB creation and serving traffic
            // on AWS. A few minutes is typical, so use 10m.
            LoadBalancerLagTimeoutAWS = 10 * time.Minute
        
            // How long to wait for a load balancer to be created/modified.
            //TODO: once support ticket 21807001 is resolved, reduce this timeout back to something reasonable
            LoadBalancerCreateTimeoutDefault = 20 * time.Minute
            LoadBalancerCreateTimeoutLarge   = 2 * time.Hour
        
            // Time required by the loadbalancer to cleanup, proportional to numApps/Ing.
            // Bring the cleanup timeout back down to 5m once b/33588344 is resolved.
            LoadBalancerCleanupTimeout = 15 * time.Minute
        
            // On average it takes ~6 minutes for a single backend to come online in GCE.
            LoadBalancerPollTimeout  = 15 * time.Minute
            LoadBalancerPollInterval = 30 * time.Second
        
            LargeClusterMinNodesNumber = 100
        
            // Don't test with more than 3 nodes.
            // Many tests create an endpoint per node, in large clusters, this is
            // resource and time intensive.
            MaxNodesForEndpointsTests = 3
        
            // ServiceTestTimeout is used for most polling/waiting activities
            ServiceTestTimeout = 60 * time.Second
        
            // GCPMaxInstancesInInstanceGroup is the maximum number of instances supported in
            // one instance group on GCP.
            GCPMaxInstancesInInstanceGroup = 2000
        
            // AffinityConfirmCount is the number of needed continuous requests to confirm that
            // affinity is enabled.
            AffinityConfirmCount = 15
            // Poll interval for StatefulSet tests
            StatefulSetPoll = 10 * time.Second
            // Timeout interval for StatefulSet operations
            StatefulSetTimeout = 10 * time.Minute
            // Timeout for stateful pods to change state
            StatefulPodTimeout = 5 * time.Minute
            // How long to wait for the pod to be listable
            PodListTimeout = time.Minute
            // Initial pod start can be delayed O(minutes) by slow docker pulls
            // TODO: Make this 30 seconds once #4566 is resolved.
            PodStartTimeout = 5 * time.Minute
            
            // Same as `PodStartTimeout` to wait for the pod to be started, but shorter.
            // Use it case by case when we are sure pod start will not be delayed
            // minutes by slow docker pulls or something else.
            PodStartShortTimeout = 2 * time.Minute
            
            // How long to wait for a pod to be deleted
                PodDeleteTimeout = 5 * time.Minute
            
                // PodEventTimeout is how much we wait for a pod event to occur.
                PodEventTimeout = 2 * time.Minute
            
                // If there are any orphaned namespaces to clean up, this test is running
                // on a long lived cluster. A long wait here is preferably to spurious test
                // failures caused by leaked resources from a previous test run.
                NamespaceCleanupTimeout = 15 * time.Minute
            
                // How long to wait for a service endpoint to be resolvable.
                ServiceStartTimeout = 3 * time.Minute
            
                // How often to Poll pods, nodes and claims.
                Poll = 2 * time.Second
            
                // service accounts are provisioned after namespace creation
                // a service account is required to support pod creation in a namespace as part of admission control
                ServiceAccountProvisionTimeout = 2 * time.Minute
            
                // How long to try single API calls (like 'get' or 'list'). Used to prevent
                // transient failures from failing tests.
                // TODO: client should not apply this timeout to Watch calls. Increased from 30s until that is fixed.
                SingleCallTimeout = 5 * time.Minute
            
                // How long nodes have to be "ready" when a test begins. They should already
                // be "ready" before the test starts, so this is small.
                NodeReadyInitialTimeout = 20 * time.Second
            
                // How long pods have to be "ready" when a test begins.
                PodReadyBeforeTimeout = 5 * time.Minute
            
                ServiceRespondingTimeout = 2 * time.Minute
                EndpointRegisterTimeout  = time.Minute
            
                // How long claims have to become dynamically provisioned
                ClaimProvisionTimeout = 5 * time.Minute
            
                // Same as `ClaimProvisionTimeout` to wait for claim to be dynamically provisioned, but shorter.
                // Use it case by case when we are sure this timeout is enough.
                ClaimProvisionShortTimeout = 1 * time.Minute
            
                // How long claims have to become bound
                ClaimBindingTimeout = 3 * time.Minute
            
                // How long claims have to become deleted
                ClaimDeletingTimeout = 3 * time.Minute
            
                // How long PVs have to beome reclaimed
                PVReclaimingTimeout = 3 * time.Minute
            
                // How long PVs have to become bound
                PVBindingTimeout = 3 * time.Minute
            
                // How long PVs have to become deleted
                PVDeletingTimeout = 3 * time.Minute
            
                // How long a node is allowed to become "Ready" after it is restarted before
                // the test is considered failed.
                RestartNodeReadyAgainTimeout = 5 * time.Minute
            
                // How long a pod is allowed to become "running" and "ready" after a node
                // restart before test is considered failed.
                RestartPodReadyAgainTimeout = 5 * time.Minute
                 Kb  int64 = 1000
                    Mb  int64 = 1000 * Kb
                    Gb  int64 = 1000 * Mb
                    Tb  int64 = 1000 * Gb
                    KiB int64 = 1024
                    MiB int64 = 1024 * KiB
                    GiB int64 = 1024 * MiB
                    TiB int64 = 1024 * GiB
                
                    // Waiting period for volume server (Ceph, ...) to initialize itself.
                    VolumeServerPodStartupTimeout = 3 * time.Minute
                
                    // Waiting period for pod to be cleaned up and unmount its volumes so we
                    // don't tear down containers with NFS/Ceph/Gluster server too early.
                    PodCleanupTimeout = 20 * time.Second
             
                    // CurrentKubeletPerfMetricsVersion is the current kubelet performance 
               // metrics version. This is used by mutiple perf related data structures. 
              // We should bump up the version each time we make an incompatible change 
              // to the metrics
              CurrentKubeletPerfMetricsVersion = "v2"
              // Default value for how long the CPU profile is gathered for. DefaultCPUProfileSeconds = 30 )
              // Default value for how long the CPU profile is gathered for.
              DefaultCPUProfileSeconds = 30    
              // TODO(mikedanese): reset this to 5 minutes once #47135 is resolved.
              // ref https://github.com/kubernetes/kubernetes/issues/47135
              DefaultNamespaceDeletionTimeout = 10 * time.Minute
              DefaultPodDeletionTimeout = 3 * time.Minute
              NoCPUConstraint = math.MaxFloat64
              // NodeStartupThreshold is a rough estimate of the time allocated for a pod to start on a node.
              NodeStartupThreshold = 4 * time.Second
            )
)
```

</p>
</details>

***

## Variables

<details><summary>show</summary>
<p>

var (
    BusyBoxImage = imageutils.GetE2EImage(imageutils.BusyBox)

    // Serve hostname image name
    ServeHostnameImage = imageutils.GetE2EImage(imageutils.ServeHostname)

    ImageWhiteList sets.String
    // ImageWhiteList is the images used in the current test suite. It should be initialized in test suite and the images 
    // in the white list should be pre-pulled in the test suite. Currently, this is only used by node e2e test.

    InterestingApiServerMetrics = []string{
        "apiserver_request_count",
        "apiserver_request_latencies_summary",
        "etcd_helper_cache_entry_count",
        "etcd_helper_cache_hit_count",
        "etcd_helper_cache_miss_count",
        "etcd_request_cache_add_latencies_summary",
        "etcd_request_cache_get_latencies_summary",
        "etcd_request_latencies_summary",
    }
    InterestingClusterAutoscalerMetrics = []string{
      "function_duration_seconds",
      "errors_total",
      "evicted_pods_total",
    }
    InterestingControllerManagerMetrics = []string{
      "garbage_collector_attempt_to_delete_queue_latency",
      "garbage_collector_attempt_to_delete_work_duration",
      "garbage_collector_attempt_to_orphan_queue_latency",
      "garbage_collector_attempt_to_orphan_work_duration",
      "garbage_collector_dirty_processing_latency_microseconds",
      "garbage_collector_event_processing_latency_microseconds",
      "garbage_collector_graph_changes_queue_latency",
      "garbage_collector_graph_changes_work_duration",
      "garbage_collector_orphan_processing_latency_microseconds",

      "namespace_queue_latency",
      "namespace_queue_latency_sum",
      "namespace_queue_latency_count",
      "namespace_retries",
      "namespace_work_duration",
      "namespace_work_duration_sum",
      "namespace_work_duration_count",
    }
    InterestingKubeletMetrics = []string{
      "kubelet_container_manager_latency_microseconds",
      "kubelet_docker_errors",
      "kubelet_docker_operations_latency_microseconds",
      "kubelet_generate_pod_status_latency_microseconds",
      "kubelet_pod_start_latency_microseconds",
      "kubelet_pod_worker_latency_microseconds",
      "kubelet_pod_worker_start_latency_microseconds",
      "kubelet_sync_pods_latency_microseconds",
    }
    NetexecImageName = imageutils.GetE2EImage(imageutils.Netexec)
    ProvidersWithSSH = []string{"gce", "gke", "aws", "local"}
    // ProvidersWithSSH are those providers where each node is accessible with SSH

    RunId = uuid.NewUUID()
    // unique identifier of the e2e run


    // Common selinux labels
    SELinuxLabel = &v1.SELinuxOptions{
        Level: "s0:c0,c1"}

    SchedulingLatencyMetricName = model.LabelValue(schedulermetric.SchedulerSubsystem + "_" +       schedulermetric.SchedulingLatencyName)
    ServiceNodePortRange = utilnet.PortRange{Base: 30000, Size: 2768}
    // This should match whatever the default/configured range is
)

</p>
</details>

***

