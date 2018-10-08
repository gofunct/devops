# KSonnet Study Guide

KSonnet is used to define resources for Kubernetes.

## Create a deployment

<details><summary>show</summary>
<p>

Note how this wraps a container object in a deployment object, specifying the number of replicas.  The `core.v1.list.new` call is used to convert to a Kubernetes resource list.

```
// Import KSonnet library
local k = import "ksonnet.beta.2/k.libsonnet";

// Specify the import objects that we need
local container = k.extensions.v1beta1.deployment.mixin.spec.template.spec.containersType;
local depl = k.extensions.v1beta1.deployment;

// Define containers
local containers = [
      container.new("gaffer", "gcr.io/trust-networks/gaffer:0.7.4b")
];

// Define deployment with 3 replicas
local deployment = 
    depl.new("gaffer", 3, containers, {app: "gaffer"});

local resources = [ deployment ];

// Return list of resources.
k.core.v1.list.new(resources)
```

</p>
</details>

## Add environment variables

<details><summary>show</summary>
<p>

Notice how the `envs` array is initialised with environment variables.

```

// Import KSonnet library
local k = import "ksonnet.beta.2/k.libsonnet";

// Specify the import objects that we need
local container = k.extensions.v1beta1.deployment.mixin.spec.template.spec.containersType;
local depl = k.extensions.v1beta1.deployment;
local env = container.envType;

// Environment variables
local envs = [

    // List of Zookeepers.
    env.new("ZOOKEEPERS", "zk1,zk2,zk3")

];

// Define containers
local containers = [
      container.new("gaffer", "gcr.io/trust-networks/gaffer:0.7.4b") +
      container.env(envs)
];

// Define deployment with 3 replicas
local deployment = 
    depl.new("gaffer", 3, containers, {app: "gaffer"});

local resources = [ deployment ];

// Return list of resources.
k.core.v1.list.new(resources)

```

</p>
</details>

## Resource specifications

<details><summary>show</summary>
<p>

Add resource limits - notice the extra `limits` and `requests` references in the container definition.

```

// Import KSonnet library
local k = import "ksonnet.beta.2/k.libsonnet";

// Specify the import objects that we need
local container = k.extensions.v1beta1.deployment.mixin.spec.template.spec.containersType;
local depl = k.extensions.v1beta1.deployment;
local env = container.envType;

// Environment variables
local envs = [

    // List of Zookeepers.
    env.new("ZOOKEEPERS", "zk1,zk2,zk3")

];

// Define containers
local containers = [
    container.new("gaffer", "gcr.io/trust-networks/gaffer:0.7.4b") +
        container.env(envs) +
        container.mixin.resources.limits({
            memory: "1G", cpu: "1.5"
        }) +
        container.mixin.resources.requests({
            memory: "1G", cpu: "1.0"
        })
];

// Define deployment with 3 replicas
local deployment = 
    depl.new("gaffer", 3, containers, {app: "gaffer"});

local resources = [ deployment ];

// Return list of resources.
k.core.v1.list.new(resources)

```

</p>
</details>

## Container ports

<details><summary>show</summary>
<p>

Notice the `containerPort` definition which is refered to by a `container.ports` reference, in order to specify ports on the container.
```

// Import KSonnet library
local k = import "ksonnet.beta.2/k.libsonnet";

// Specify the import objects that we need
local container = k.extensions.v1beta1.deployment.mixin.spec.template.spec.containersType;
local containerPort = container.portsType;
local depl = k.extensions.v1beta1.deployment;
local env = container.envType;

// Environment variables
local envs = [

    // List of Zookeepers.
    env.new("ZOOKEEPERS", "zk1,zk2,zk3")

];

// Ports used by deployments
local ports = [
    containerPort.newNamed("rest", 8080)
];

// Define containers
local containers = [
    container.new("gaffer", "gcr.io/trust-networks/gaffer:0.7.4b") +
        container.ports(ports) +
        container.env(envs) +
        container.mixin.resources.limits({
            memory: "1G", cpu: "1.5"
        }) +
        container.mixin.resources.requests({
            memory: "1G", cpu: "1.0"
        })
];

// Define deployment with 3 replicas
local deployment = 
    depl.new("gaffer", 3, containers, {app: "gaffer"});

local resources = [ deployment ];

// Return list of resources.
k.core.v1.list.new(resources)

```

</p>
</details>

## Volumes

<details><summary>show</summary>
<p>

Notice the `volumeMounts` declaration referenced in `container.volumeMounts`, and the `volumes` definition which is used in the deployment definition.
```

// Import KSonnet library
local k = import "ksonnet.beta.2/k.libsonnet";

// Specify the import objects that we need
local container = k.extensions.v1beta1.deployment.mixin.spec.template.spec.containersType;
local containerPort = container.portsType;
local mount = container.volumeMountsType;
local depl = k.extensions.v1beta1.deployment;
local env = container.envType;
local volume = depl.mixin.spec.template.spec.volumesType;
local gceDisk = volume.mixin.gcePersistentDisk;

// Environment variables
local envs = [

    // List of Zookeepers.
    env.new("ZOOKEEPERS", "zk1,zk2,zk3")

];

// Ports used by deployments
local ports = [
    containerPort.newNamed("rest", 8080)
];

// Volume mount points
local volumeMounts = [
    mount.new("data", "/data")
];

// Volumes - this invokes a GCE permanent disk.
local volumes = [
    volume.name("data") + gceDisk.fsType("ext4") +
          gceDisk.pdName("data-disk")
];

// Define containers
local containers = [
    container.new("gaffer", "gcr.io/trust-networks/gaffer:0.7.4b") +
        container.ports(ports) +
        container.env(envs) +
  container.volumeMounts(volumeMounts) +
        container.mixin.resources.limits({
            memory: "1G", cpu: "1.5"
        }) +
        container.mixin.resources.requests({
            memory: "1G", cpu: "1.0"
        })
];

// Define deployment with 3 replicas
local deployment = 
    depl.new("gaffer", 3, containers, {app: "gaffer"}) +
    depl.mixin.spec.template.spec.volumes(volumes);

local resources = [ deployment ];

// Return list of resources.
k.core.v1.list.new(resources)

```

</p>
</details>

## Add a service

<details><summary>show</summary>
<p>

The `servicePorts` declaration maps the external port 8080 to the container's 8080 port.  Then, `service` defines the service which is added to the `resource` array.

```

// Import KSonnet library
local k = import "ksonnet.beta.2/k.libsonnet";

// Specify the import objects that we need
local container = k.extensions.v1beta1.deployment.mixin.spec.template.spec.containersType;
local containerPort = container.portsType;
local mount = container.volumeMountsType;
local depl = k.extensions.v1beta1.deployment;
local env = container.envType;
local volume = depl.mixin.spec.template.spec.volumesType;
local gceDisk = volume.mixin.gcePersistentDisk;
local svc = k.core.v1.service;
local svcPort = svc.mixin.spec.portsType;
local svcLabels = svc.mixin.metadata.labels;

// Environment variables
local envs = [

    // List of Zookeepers.
    env.new("ZOOKEEPERS", "zk1,zk2,zk3")

];

// Ports used by deployments
local ports = [
    containerPort.newNamed("rest", 8080)
];

// Volume mount points
local volumeMounts = [
    mount.new("data", "/data")
];

// Volumes - this invokes a GCE permanent disk.
local volumes = [
    volume.name("data") + gceDisk.fsType("ext4") +
        gceDisk.pdName("data-disk")
];

// Define containers
local containers = [
   container.new("gaffer", "gcr.io/trust-networks/gaffer:0.7.4b") +
        container.ports(ports) +
        container.env(envs) +
  container.volumeMounts(volumeMounts) +
        container.mixin.resources.limits({
            memory: "1G", cpu: "1.5"
        }) +
        container.mixin.resources.requests({
            memory: "1G", cpu: "1.0"
        })
];

// Define deployment with 3 replicas
local deployment = 
    depl.new("gaffer", 3, containers, {app: "gaffer"}) +
    depl.mixin.spec.template.spec.volumes(volumes);

// Ports declared on the service.
local servicePorts = [
    svcPort.newNamed("rest", 8080, 8080) + svcPort.protocol("TCP")
];

local service =
    // One service load-balanced across the replicas
    svc.new("gaffer", {app: "gaffer"}, servicePorts) +
    svcLabels({app: "gaffer", component: "gaffer"});

local resources = [ deployment, service ];

// Return list of resources.
k.core.v1.list.new(resources)
```

</p>
</details>
