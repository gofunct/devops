# Certified Kubernetes Administrator Exam Prep
by: Coleman Word

**Based on Kubernetes the Hard Way**

- [Certified Kubernetes Administrator Exam Prep](#certified-kubernetes-administrator-exam-prep)
- [Overview](#overview)
- [Questions](#questions)
- [Provisioning a CA and Generating TLS Certificates](#provisioning-a-ca-and-generating-tls-certificates)
  - [Certificate files are generated from..?](#certificate-files-are-generated-from)
  - [Certificate authority is generated from..?](#certificate-authority-is-generated-from)
  - [The kube-proxy, kube-controller-manager, kube-scheduler, and kubelet client certificates will be used to..?](#the-kube-proxy-kube-controller-manager-kube-scheduler-and-kubelet-client-certificates-will-be-used-to)
  - [What keys are copied to workers?](#what-keys-are-copied-to-workers)
  - [What keys are copied to the controllers?](#what-keys-are-copied-to-the-controllers)
- [Generating Kubernetes Configuration Files for Authentication](#generating-kubernetes-configuration-files-for-authentication)
  - [What configs are copied to the workers?](#what-configs-are-copied-to-the-workers)
  - [What kubeconfigs are copied to the controllers?](#what-kubeconfigs-are-copied-to-the-controllers)
  - [How do you generate kubeconfig files from certificates using kubectl?](#how-do-you-generate-kubeconfig-files-from-certificates-using-kubectl)
- [Generating the Data Encryption Config and Key](#generating-the-data-encryption-config-and-key)
  - [How do you generate an encryption key?](#how-do-you-generate-an-encryption-key)
  - [How do you generate an encryption configuration for controllers?](#how-do-you-generate-an-encryption-configuration-for-controllers)
- [Bootstrapping the etcd Cluster](#bootstrapping-the-etcd-cluster)
  - [What is etcd used for?](#what-is-etcd-used-for)
  - [How do you configure the etcd server?](#how-do-you-configure-the-etcd-server)
- [Bootstrapping the Kubernetes Control Plane](#bootstrapping-the-kubernetes-control-plane)
  - [How do you create the Kubernetes config directory?](#how-do-you-create-the-kubernetes-config-directory)
  - [What should you move to /usr/local/bin when bootstrapping the control plane?](#what-should-you-move-to-usrlocalbin-when-bootstrapping-the-control-plane)
  - [What should be moved to /var/lib/kubernetes when bootstrapping the control plane?](#what-should-be-moved-to-varlibkubernetes-when-bootstrapping-the-control-plane)
  - [What should you move to /etc/systemd/system/ when bootstraping the controle plane?](#what-should-you-move-to-etcsystemdsystem-when-bootstraping-the-controle-plane)
  - [How do you start the controller services?](#how-do-you-start-the-controller-services)
  - [What do you need for controller health-checks?](#what-do-you-need-for-controller-health-checks)
- [Bootstrapping the Kubernetes Worker Nodes](#bootstrapping-the-kubernetes-worker-nodes)
  - [Why do you need RBAC for kubelet on worker nodes?](#why-do-you-need-rbac-for-kubelet-on-worker-nodes)
  - [What do you need to create to let the Kubernetes API server on the controllers communicate with the kubelet api on each worker?](#what-do-you-need-to-create-to-let-the-kubernetes-api-server-on-the-controllers-communicate-with-the-kubelet-api-on-each-worker)
  - [A ClusterRole must be bound between what to components to activate communication between the Kubernetes API and the kubelet api?](#a-clusterrole-must-be-bound-between-what-to-components-to-activate-communication-between-the-kubernetes-api-and-the-kubelet-api)
  - [What should the external load balancer attach to?](#what-should-the-external-load-balancer-attach-to)
  - [After setting up the external load balancer, how can you check the Kubernetes version info?](#after-setting-up-the-external-load-balancer-how-can-you-check-the-kubernetes-version-info)
  - [What needs to be installed on each worker node to bootstrap?](#what-needs-to-be-installed-on-each-worker-node-to-bootstrap)
  - [What do you need to do to configure CNI plugins?](#what-do-you-need-to-do-to-configure-cni-plugins)
  - [How do you need to configure the CNI containerd?](#how-do-you-need-to-configure-the-cni-containerd)
  - [How do you configure the Kubelet when bootstraping the worker nodes?](#how-do-you-configure-the-kubelet-when-bootstraping-the-worker-nodes)
  - [What components are being added to what directories when bootsrapping worker nodes?](#what-components-are-being-added-to-what-directories-when-bootsrapping-worker-nodes)
  - [How do you start worker services after bootsraping them?](#how-do-you-start-worker-services-after-bootsraping-them)
  - [How do you verify that the worker bootstrap was successful?](#how-do-you-verify-that-the-worker-bootstrap-was-successful)
- [Configuring Kubectl for Remote Access](#configuring-kubectl-for-remote-access)
  - [What does each config file point to?](#what-does-each-config-file-point-to)
  - [How do you Generate a kubeconfig file suitable for authenticating as the admin user?](#how-do-you-generate-a-kubeconfig-file-suitable-for-authenticating-as-the-admin-user)
  - [How can you check cluster health after setting up the config for remote admin auth?](#how-can-you-check-cluster-health-after-setting-up-the-config-for-remote-admin-auth)
- [Provisioning Pod Network Routes](#provisioning-pod-network-routes)
  - [How can you enable pods to communicate with eachother?](#how-can-you-enable-pods-to-communicate-with-eachother)
  - [Print the internal IP address and Pod CIDR range for each worker instance:](#print-the-internal-ip-address-and-pod-cidr-range-for-each-worker-instance)
  - [Create network routes for each worker instance:](#create-network-routes-for-each-worker-instance)
- [Deploying the DNS Cluster Add-on](#deploying-the-dns-cluster-add-on)
  - [How do you implement DNS based service discovery?](#how-do-you-implement-dns-based-service-discovery)
  - [How can you execute a DNS lookup for a kubernetes service inside a pod?](#how-can-you-execute-a-dns-lookup-for-a-kubernetes-service-inside-a-pod)
- [Smoke Test](#smoke-test)
  - [How can you encrypt secret data at rest?](#how-can-you-encrypt-secret-data-at-rest)
  - [How can you enable port forwarding to an nginx deployment?](#how-can-you-enable-port-forwarding-to-an-nginx-deployment)
  - [How do you retrieve logs from containers?](#how-do-you-retrieve-logs-from-containers)
  - [How can you print the nginx version by executing on the container?](#how-can-you-print-the-nginx-version-by-executing-on-the-container)
  - [How can you expose a pod using a service?](#how-can-you-expose-a-pod-using-a-service)
  - [How do you verify the ability to run untrusted workloads using gVisor?](#how-do-you-verify-the-ability-to-run-untrusted-workloads-using-gvisor)

***
# Overview


***

# Questions

# Provisioning a CA and Generating TLS Certificates

***

## Certificate files are generated from..?

<details><summary>show</summary>
<p>


* ca.pem 
* ca-key.pem
* csr.json file
* generated with cfssl and cfssl json


</p>
</details>

***

## Certificate authority is generated from..?

<details><summary>show</summary>
<p>

ca-config.json file that specifies:

* server auth, client auth

ca-csr.json specifies:
* location 
* name
* rsa 2048



```
{

cat > ca-config.json <<EOF
{
  "signing": {
    "default": {
      "expiry": "8760h"
    },
    "profiles": {
      "kubernetes": {
        "usages": ["signing", "key encipherment", "server auth", "client auth"],
        "expiry": "8760h"
      }
    }
  }
}
EOF

cat > ca-csr.json <<EOF
{
  "CN": "Kubernetes",
  "key": {
    "algo": "rsa",
    "size": 2048
  },
  "names": [
    {
      "C": "US",
      "L": "Portland",
      "O": "Kubernetes",
      "OU": "CA",
      "ST": "Oregon"
    }
  ]
}
EOF

cfssl gencert -initca ca-csr.json | cfssljson -bare ca

}
```

</p>
</details>

***

## The kube-proxy, kube-controller-manager, kube-scheduler, and kubelet client certificates will be used to..?

<details><summary>show</summary>
<p>

client authentication configuration files

</p>
</details>

***

## What keys are copied to workers?

<details><summary>show</summary>
<p>

```
ca.pem ${instance}-key.pem ${instance}.pem

```

</p>
</details>

***

## What keys are copied to the controllers?

<details><summary>show</summary>
<p>

ca.pem 
ca-key.pem 
kubernetes-key.pem 
kubernetes.pem
service-account-key.pem 
service-account.pem

</p>
</details>

***

# Generating Kubernetes Configuration Files for Authentication


***

## What configs are copied to the workers?

<details><summary>show</summary>
<p>

* ${instance}.kubeconfig 
* kube-proxy.kubeconfig

</p>
</details>

***

## What kubeconfigs are copied to the controllers?

<details><summary>show</summary>
<p>

```
* admin.kubeconfig 
* kube-controller-manager.kubeconfig 
* kube-scheduler.kubeconfig 

```

</p>
</details>

***

## How do you generate kubeconfig files from certificates using kubectl?

<details><summary>show</summary>
<p>

* kubectl set-cluster with cert-auth, embedded certs, server addr, and the designated kubeconfig
* kubectl set-credentials with designated system, client cert, client key, embedded certs, and set the kubeconfig
* kubectl set-context to default, system user, and kubeconfig
* kubectl config use-context default with desired kubeconfig

Example for kube-controller-manager:

```
{
  kubectl config set-cluster kubernetes-the-hard-way \
    --certificate-authority=ca.pem \
    --embed-certs=true \
    --server=https://127.0.0.1:6443 \
    --kubeconfig=kube-controller-manager.kubeconfig

  kubectl config set-credentials system:kube-controller-manager \
    --client-certificate=kube-controller-manager.pem \
    --client-key=kube-controller-manager-key.pem \
    --embed-certs=true \
    --kubeconfig=kube-controller-manager.kubeconfig

  kubectl config set-context default \
    --cluster=kubernetes-the-hard-way \
    --user=system:kube-controller-manager \
    --kubeconfig=kube-controller-manager.kubeconfig

  kubectl config use-context default --kubeconfig=kube-controller-manager.kubeconfig
}
```

</p>
</details>

***

# Generating the Data Encryption Config and Key


***

## How do you generate an encryption key?

<details><summary>show</summary>
<p>

```
ENCRYPTION_KEY=$(head -c 32 /dev/urandom | base64)

```

</p>
</details>

***

## How do you generate an encryption configuration for controllers?

<details><summary>show</summary>
<p>
* Need the encryption key

```
cat > encryption-config.yaml <<EOF
kind: EncryptionConfig
apiVersion: v1
resources:
  - resources:
      - secrets
    providers:
      - aescbc:
          keys:
            - name: key1
              secret: ${ENCRYPTION_KEY}
      - identity: {}
EOF
```

</p>
</details>

***

# Bootstrapping the etcd Cluster


***

## What is etcd used for?

<details><summary>show</summary>
<p>

* storing cluster state
* located on controllers

</p>
</details>

***

## How do you configure the etcd server?

<details><summary>show</summary>
<p>
**must be configured on each controller**

1. Create the directories /etc/etcd and /var/lib/etcd and add them to your path
2. Copy ca.pem, kubernetes-key.pem, kubernetes.pem to /etc/etcd/ directory

```

{
  sudo mkdir -p /etc/etcd /var/lib/etcd
  sudo cp ca.pem kubernetes-key.pem kubernetes.pem /etc/etcd/
}

```

3. Set the internal ip of each controller 

```

INTERNAL_IP=$(curl -s -H "Metadata-Flavor: Google" \
  http://metadata.google.internal/computeMetadata/v1/instance/network-interfaces/0/ip)
```
4. Set the etcd name to match the hostname of the current compute instance
```
ETCD_NAME=$(hostname -s)

```

5. Create the etcd.service systemd unit file

```

cat <<EOF | sudo tee /etc/systemd/system/etcd.service
[Unit]
Description=etcd
Documentation=https://github.com/coreos

[Service]
ExecStart=/usr/local/bin/etcd \\
  --name ${ETCD_NAME} \\
  --cert-file=/etc/etcd/kubernetes.pem \\
  --key-file=/etc/etcd/kubernetes-key.pem \\
  --peer-cert-file=/etc/etcd/kubernetes.pem \\
  --peer-key-file=/etc/etcd/kubernetes-key.pem \\
  --trusted-ca-file=/etc/etcd/ca.pem \\
  --peer-trusted-ca-file=/etc/etcd/ca.pem \\
  --peer-client-cert-auth \\
  --client-cert-auth \\
  --initial-advertise-peer-urls https://${INTERNAL_IP}:2380 \\
  --listen-peer-urls https://${INTERNAL_IP}:2380 \\
  --listen-client-urls https://${INTERNAL_IP}:2379,https://127.0.0.1:2379 \\
  --advertise-client-urls https://${INTERNAL_IP}:2379 \\
  --initial-cluster-token etcd-cluster-0 \\
  --initial-cluster controller-0=https://10.240.0.10:2380,controller-1=https://10.240.0.11:2380,controller-2=https://10.240.0.12:2380 \\
  --initial-cluster-state new \\
  --data-dir=/var/lib/etcd
Restart=on-failure
RestartSec=5

[Install]
WantedBy=multi-user.target
EOF

```

6. Start the etcd server

```

{
  sudo systemctl daemon-reload
  sudo systemctl enable etcd
  sudo systemctl start etcd
}

```

7. Verify 

```

sudo ETCDCTL_API=3 etcdctl member list \
  --endpoints=https://127.0.0.1:2379 \
  --cacert=/etc/etcd/ca.pem \
  --cert=/etc/etcd/kubernetes.pem \
  --key=/etc/etcd/kubernetes-key.pem
```



</p>
</details>

***

# Bootstrapping the Kubernetes Control Plane

***

## How do you create the Kubernetes config directory?

<details><summary>show</summary>
<p>

```
sudo mkdir -p /etc/kubernetes/config


```

</p>
</details>

***

## What should you move to /usr/local/bin when bootstrapping the control plane?

<details><summary>show</summary>
<p>
**make executable**

* kube-apiserver 
* kube-controller-manager 
* kube-scheduler 
* kubectl
* etcd 


</p>
</details>

***

## What should be moved to /var/lib/kubernetes when bootstrapping the control plane?

<details><summary>show</summary>
<p>

```
* ca.pem 
* ca-key.pem 
* kubernetes-key.pem 
* kubernetes.pem 
* service-account-key.pem 
* service-account.pem 
* encryption-config.yaml
* kube-controller-manager.kubeconfig
* kube-scheduler.kubeconfig
* 
* 
```

</p>
</details>

***

## What should you move to /etc/systemd/system/ when bootstraping the controle plane?

<details><summary>show</summary>
<p>

* kube-apiserver.service
* kube-controller-manager.service
* kube-scheduler.service
* 
* 

</p>
</details>

***

## How do you start the controller services?

<details><summary>show</summary>
<p>

```
{
  sudo systemctl daemon-reload
  sudo systemctl enable kube-apiserver kube-controller-manager kube-scheduler
  sudo systemctl start kube-apiserver kube-controller-manager kube-scheduler
}

```

</p>
</details>

***

## What do you need for controller health-checks?

<details><summary>show</summary>
<p>
* A basic web server

```
sudo apt-get install -y nginx
```

* Configure NGINX

```
cat > kubernetes.default.svc.cluster.local <<EOF
server {
  listen      80;
  server_name kubernetes.default.svc.cluster.local;

  location /healthz {
     proxy_pass                    https://127.0.0.1:6443/healthz;
     proxy_ssl_trusted_certificate /var/lib/kubernetes/ca.pem;
  }
}
EOF
```

* mv config to /etc/nginx/sites-available/kubernetes.default.svc.cluster.local

```
{
  sudo mv kubernetes.default.svc.cluster.local \
    /etc/nginx/sites-available/kubernetes.default.svc.cluster.local

  sudo ln -s /etc/nginx/sites-available/kubernetes.default.svc.cluster.local /etc/nginx/sites-enabled/
}
```
* Start the NGINX service

```
sudo systemctl restart nginx
sudo systemctl enable nginx

```
* Verify

```
kubectl get componentstatuses --kubeconfig admin.kubeconfig

```
* Test access

```
curl -H "Host: kubernetes.default.svc.cluster.local" -i http://127.0.0.1/healthz

```

</p>
</details>

***

# Bootstrapping the Kubernetes Worker Nodes

***

## Why do you need RBAC for kubelet on worker nodes?

<details><summary>show</summary>
<p>

**Access to the Kubelet API from the Kubernetes API is required for retrieving:**

* metrics
* logs
* executing commands in pods

</p>
</details>

***

## What do you need to create to let the Kubernetes API server on the controllers communicate with the kubelet api on each worker?

<details><summary>show</summary>
<p>

* system:kube-apiserver-to-kubelet ClusterRole

```
cat <<EOF | kubectl apply --kubeconfig admin.kubeconfig -f -
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRole
metadata:
  annotations:
    rbac.authorization.kubernetes.io/autoupdate: "true"
  labels:
    kubernetes.io/bootstrapping: rbac-defaults
  name: system:kube-apiserver-to-kubelet
rules:
  - apiGroups:
      - ""
    resources:
      - nodes/proxy
      - nodes/stats
      - nodes/log
      - nodes/spec
      - nodes/metrics
    verbs:
      - "*"
EOF

```

The Kubernetes API Server authenticates to the Kubelet as the kubernetes user using the client certificate as defined by the --kubelet-client-certificate flag.

</p>
</details>

***

## A ClusterRole must be bound between what to components to activate communication between the Kubernetes API and the kubelet api?

<details><summary>show</summary>
<p>

* system:kube-apiserver-to-kubelet & kubernetes user

```
cat <<EOF | kubectl apply --kubeconfig admin.kubeconfig -f -
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRoleBinding
metadata:
  name: system:kube-apiserver
  namespace: ""
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: system:kube-apiserver-to-kubelet
subjects:
  - apiGroup: rbac.authorization.k8s.io
    kind: User
    name: kubernetes
EOF
```

</p>
</details>

***

## What should the external load balancer attach to?

<details><summary>show</summary>
<p>

* The kubernetes-the-hard-way static IP address 

</p>
</details>

***

## After setting up the external load balancer, how can you check the Kubernetes version info?

<details><summary>show</summary>
<p>

```
curl --cacert ca.pem https://${KUBERNETES_PUBLIC_ADDRESS}:6443/version


```

</p>
</details>

***

## What needs to be installed on each worker node to bootstrap?

<details><summary>show</summary>
<p>

**Each component needs a config file in /etc/{component}/{config.toml}**
**Each component needs a service file in /etc/systemd/system/{component}.service**

* runc
* gVisor
* container networking plugins
* containerd
* kubelet
* kube-proxy
* socat
* conntrack
* ipset



**Download socat,contrack,ipset**
```
{
  sudo apt-get update
  sudo apt-get -y install socat conntrack ipset
}
```

**Download Binaries**

```
wget -q --show-progress --https-only --timestamping \
  https://github.com/kubernetes-incubator/cri-tools/releases/download/v1.0.0-beta.0/crictl-v1.0.0-beta.0-linux-amd64.tar.gz \
  https://storage.googleapis.com/kubernetes-the-hard-way/runsc \
  https://github.com/opencontainers/runc/releases/download/v1.0.0-rc5/runc.amd64 \
  https://github.com/containernetworking/plugins/releases/download/v0.6.0/cni-plugins-amd64-v0.6.0.tgz \
  https://github.com/containerd/containerd/releases/download/v1.1.0/containerd-1.1.0.linux-amd64.tar.gz \
  https://storage.googleapis.com/kubernetes-release/release/v1.10.2/bin/linux/amd64/kubectl \
  https://storage.googleapis.com/kubernetes-release/release/v1.10.2/bin/linux/amd64/kube-proxy \
  https://storage.googleapis.com/kubernetes-release/release/v1.10.2/bin/linux/amd64/kubelet
```

**Make the installation directiories**

```
sudo mkdir -p \
  /etc/cni/net.d \
  /opt/cni/bin \
  /var/lib/kubelet \
  /var/lib/kube-proxy \
  /var/lib/kubernetes \
  /var/run/kubernetes
```

**Install the binaries**

```
{
  chmod +x kubectl kube-proxy kubelet runc.amd64 runsc
  sudo mv runc.amd64 runc
  sudo mv kubectl kube-proxy kubelet runc runsc /usr/local/bin/
  sudo tar -xvf crictl-v1.0.0-beta.0-linux-amd64.tar.gz -C /usr/local/bin/
  sudo tar -xvf cni-plugins-amd64-v0.6.0.tgz -C /opt/cni/bin/
  sudo tar -xvf containerd-1.1.0.linux-amd64.tar.gz -C /
}
```


</p>
</details>

***

## What do you need to do to configure CNI plugins?

<details><summary>show</summary>
<p>

**The pod''s CIDR range**

```
POD_CIDR=$(curl -s -H "Metadata-Flavor: Google" \
  http://metadata.google.internal/computeMetadata/v1/instance/attributes/pod-cidr)

```

**a bridge network config**

```
cat <<EOF | sudo tee /etc/cni/net.d/10-bridge.conf
{
    "cniVersion": "0.3.1",
    "name": "bridge",
    "type": "bridge",
    "bridge": "cnio0",
    "isGateway": true,
    "ipMasq": true,
    "ipam": {
        "type": "host-local",
        "ranges": [
          [{"subnet": "${POD_CIDR}"}]
        ],
        "routes": [{"dst": "0.0.0.0/0"}]
    }
}
EOF
```

**a loopback network config**

```
cat <<EOF | sudo tee /etc/cni/net.d/99-loopback.conf
{
    "cniVersion": "0.3.1",
    "type": "loopback"
}
EOF
```

</p>
</details>

***

## How do you need to configure the CNI containerd?

<details><summary>show</summary>
<p>

**A config file**

```
sudo mkdir -p /etc/containerd/
cat << EOF | sudo tee /etc/containerd/config.toml
[plugins]
  [plugins.cri.containerd]
    snapshotter = "overlayfs"
    [plugins.cri.containerd.default_runtime]
      runtime_type = "io.containerd.runtime.v1.linux"
      runtime_engine = "/usr/local/bin/runc"
      runtime_root = ""
    [plugins.cri.containerd.untrusted_workload_runtime]
      runtime_type = "io.containerd.runtime.v1.linux"
      runtime_engine = "/usr/local/bin/runsc"
      runtime_root = "/run/containerd/runsc"
EOF
```
* Untrusted workloads will be run using the gVisor (runsc) runtime.

**a containerd.service systemd unit file**


```
cat <<EOF | sudo tee /etc/systemd/system/containerd.service
[Unit]
Description=containerd container runtime
Documentation=https://containerd.io
After=network.target

[Service]
ExecStartPre=/sbin/modprobe overlay
ExecStart=/bin/containerd
Restart=always
RestartSec=5
Delegate=yes
KillMode=process
OOMScoreAdjust=-999
LimitNOFILE=1048576
LimitNPROC=infinity
LimitCORE=infinity

[Install]
WantedBy=multi-user.target
EOF

```

</p>
</details>

***

## How do you configure the Kubelet when bootstraping the worker nodes? 

<details><summary>show</summary>
<p>

Configure the Kubelet

```
{
  sudo mv ${HOSTNAME}-key.pem ${HOSTNAME}.pem /var/lib/kubelet/
  sudo mv ${HOSTNAME}.kubeconfig /var/lib/kubelet/kubeconfig
  sudo mv ca.pem /var/lib/kubernetes/
}
```

**a kubelet-config.yaml configuration file**

```
cat <<EOF | sudo tee /var/lib/kubelet/kubelet-config.yaml
kind: KubeletConfiguration
apiVersion: kubelet.config.k8s.io/v1beta1
authentication:
  anonymous:
    enabled: false
  webhook:
    enabled: true
  x509:
    clientCAFile: "/var/lib/kubernetes/ca.pem"
authorization:
  mode: Webhook
clusterDomain: "cluster.local"
clusterDNS:
  - "10.32.0.10"
podCIDR: "${POD_CIDR}"
runtimeRequestTimeout: "15m"
tlsCertFile: "/var/lib/kubelet/${HOSTNAME}.pem"
tlsPrivateKeyFile: "/var/lib/kubelet/${HOSTNAME}-key.pem"
EOF
```

**a kubelet.service systemd unit file**


```
cat <<EOF | sudo tee /etc/systemd/system/kubelet.service
[Unit]
Description=Kubernetes Kubelet
Documentation=https://github.com/kubernetes/kubernetes
After=containerd.service
Requires=containerd.service

[Service]
ExecStart=/usr/local/bin/kubelet \\
  --config=/var/lib/kubelet/kubelet-config.yaml \\
  --container-runtime=remote \\
  --container-runtime-endpoint=unix:///var/run/containerd/containerd.sock \\
  --image-pull-progress-deadline=2m \\
  --kubeconfig=/var/lib/kubelet/kubeconfig \\
  --network-plugin=cni \\
  --register-node=true \\
  --v=2
Restart=on-failure
RestartSec=5

[Install]
WantedBy=multi-user.target
EOF


```

</p>
</details>

***

## What components are being added to what directories when bootsrapping worker nodes?

<details><summary>show</summary>
<p>

**Installation Directories**

1. /user/local/bin
* kubectl 
* kube-proxy 
* kubelet 
* runc 
* runsc
* crictl

2. /etc/cni/net.d 
* b10-bridge.conf
* 99-loopback.conf

3. /opt/cni/bin 
* cni-plugins

4. /var/lib/kubelet 
* ${HOSTNAME}-key.pem 
* ${HOSTNAME}.pem
* kubelet-config.yaml

5. /var/lib/kubelet/kubeconfig/
* ${HOSTNAME}.kubeconfig

6. /var/lib/kube-proxy 

7. /var/lib/kube-proxy/kubeconfig
* kube-proxy.kubeconfig

6. /var/lib/kubernetes
* ca.pem

7. /var/run/kubernetes

8. /
* containerd

9. /etc/containerd/
* config.toml

10. /etc/systemd/service/
* containerd.service
* kubelet.service
* kube-proxy.service


</p>
</details>

***

## How do you start worker services after bootsraping them?

<details><summary>show</summary>
<p>

```
{
  sudo systemctl daemon-reload
  sudo systemctl enable containerd kubelet kube-proxy
  sudo systemctl start containerd kubelet kube-proxy
}

```

</p>
</details>

***

## How do you verify that the worker bootstrap was successful?

<details><summary>show</summary>
<p>

```
gcloud compute ssh controller-0 \
  --command "kubectl get nodes --kubeconfig admin.kubeconfig"

```

</p>
</details>

***

# Configuring Kubectl for Remote Access

***

## What does each config file point to?

<details><summary>show</summary>
<p>

* the external load balancer fronting the Kubernetes API Servers

</p>
</details>

***

## How do you Generate a kubeconfig file suitable for authenticating as the admin user?



<details><summary>show</summary>
<p>

**Requirements:**
* kubernetes public ip address
* set-cluster
* set-credentials
* set-context
* use-context
```
{
  KUBERNETES_PUBLIC_ADDRESS=$(gcloud compute addresses describe kubernetes-the-hard-way \
    --region $(gcloud config get-value compute/region) \
    --format 'value(address)')

  kubectl config set-cluster kubernetes-the-hard-way \
    --certificate-authority=ca.pem \
    --embed-certs=true \
    --server=https://${KUBERNETES_PUBLIC_ADDRESS}:6443

  kubectl config set-credentials admin \
    --client-certificate=admin.pem \
    --client-key=admin-key.pem

  kubectl config set-context kubernetes-the-hard-way \
    --cluster=kubernetes-the-hard-way \
    --user=admin

  kubectl config use-context kubernetes-the-hard-way
}

```

</p>
</details>


***

## How can you check cluster health after setting up the config for remote admin auth?

<details><summary>show</summary>
<p>

```
kubectl get componentstatuses
kubectl get nodes


```

</p>
</details>

***

# Provisioning Pod Network Routes


***

## How can you enable pods to communicate with eachother?

<details><summary>show</summary>
<p>

* Pods scheduled to a node receive an IP address from the node's Pod CIDR range. 

**Create a route for each worker node that maps the node's Pod CIDR range to the node's internal IP address**

</p>
</details>


***

## Print the internal IP address and Pod CIDR range for each worker instance:



<details><summary>show</summary>
<p>

```
for instance in worker-0 worker-1 worker-2; do
  gcloud compute instances describe ${instance} \
    --format 'value[separator=" "](networkInterfaces[0].networkIP,metadata.items[0].value)'
done
```

</p>
</details>

***

## Create network routes for each worker instance:



<details><summary>show</summary>
<p>

```
for i in 0 1 2; do
  gcloud compute routes create kubernetes-route-10-200-${i}-0-24 \
    --network kubernetes-the-hard-way \
    --next-hop-address 10.240.0.2${i} \
    --destination-range 10.200.${i}.0/24
done

```

</p>
</details>

***

# Deploying the DNS Cluster Add-on


***

## How do you implement DNS based service discovery?

<details><summary>show</summary>
<p>
  
Deploy the DNS add-on which provides DNS based service discovery to applications running inside the Kubernetes cluster.

**Deploy the kube-dns cluster add-on**

```

kubectl create -f https://storage.googleapis.com/kubernetes-the-hard-way/kube-dns.yaml

```

**List the pods created by the kube-dns deployment**


```

kubectl get pods -l k8s-app=kube-dns -n kube-system

```


</p>
</details>


***

## How can you execute a DNS lookup for a kubernetes service inside a pod?



<details><summary>show</summary>
<p>

```
kubectl exec -ti $POD_NAME -- nslookup kubernetes


```

</p>
</details>

***

# Smoke Test

***

## How can you encrypt secret data at rest?

<details><summary>show</summary>
<p>
  
  **Create a generic secret**

```
kubectl create secret generic kubernetes-the-hard-way \
  --from-literal="mykey=mydata"

```

**Print a hexdump of the kubernetes-the-hard-way secret stored in etcd**

```
gcloud compute ssh controller-0 \
  --command "sudo ETCDCTL_API=3 etcdctl get \
  --endpoints=https://127.0.0.1:2379 \
  --cacert=/etc/etcd/ca.pem \
  --cert=/etc/etcd/kubernetes.pem \
  --key=/etc/etcd/kubernetes-key.pem\
  /registry/secrets/default/kubernetes-the-hard-way | hexdump -C"


```

The etcd key should be prefixed with k8s:enc:aescbc:v1:key1, which indicates the aescbc provider was used to encrypt the data with the key1 encryption key.

</p>
</details>


***

## How can you enable port forwarding to an nginx deployment?

<details><summary>show</summary>
<p>

**Create a deployment for the nginx web server**

```
kubectl run nginx --image=nginx
```

**List the pod created by the nginx deployment**

```
kubectl get pods -l run=nginx
```

**Retrieve the full name of the nginx pod**

```
POD_NAME=$(kubectl get pods -l run=nginx -o jsonpath="{.items[0].metadata.name}")

```

**Forward port 8080 on your local machine to port 80 of the nginx pod**

```
kubectl port-forward $POD_NAME 8080:80

```

**In a new terminal make an HTTP request using the forwarding address**

```
curl --head http://127.0.0.1:8080

```

</p>
</details>

## How do you retrieve logs from containers?

<details><summary>show</summary>
<p>
    
Print pod logs:

```
kubectl logs $POD_NAME

```

</p>
</details>

***

## How can you print the nginx version by executing on the container?

<details><summary>show</show>
<p>
  
```
kubectl exec -ti $POD_NAME -- nginx -v
  
```

</p>
</details

***

## How can you expose a pod using a service?

<details><summary>show</show>
<p>

**Expose the nginx deployment using a NodePort service**

```
kubectl expose deployment nginx --port 80 --type NodePort
```

**Retrieve the node port assigned to the nginx service**

```
NODE_PORT=$(kubectl get svc nginx \
  --output=jsonpath='{range .spec.ports[0]}{.nodePort}')
```

**Create a firewall rule that allows remote access to the nginx node port**

```
gcloud compute firewall-rules create kubernetes-the-hard-way-allow-nginx-service \
  --allow=tcp:${NODE_PORT} \
  --network kubernetes-the-hard-way
```

**Retrieve the external IP address of a worker instance**

```
EXTERNAL_IP=$(gcloud compute instances describe worker-0 \
  --format 'value(networkInterfaces[0].accessConfigs[0].natIP)')
```

**Make an HTTP request using the external IP address and the nginx node port**

```
curl -I http://${EXTERNAL_IP}:${NODE_PORT}
```


</p>
</details

***

## How do you verify the ability to run untrusted workloads using gVisor?

<details><summary>show</show>
<p>

**Create the untrusted pod**

```
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Pod
metadata:
  name: untrusted
  annotations:
    io.kubernetes.cri.untrusted-workload: "true"
spec:
  containers:
    - name: webserver
      image: gcr.io/hightowerlabs/helloworld:2.0.0
EOF
```

**Verify the untrusted pod is running**

```
kubectl get pods -o wide
```

**Get the node name where the untrusted pod is running**

```
INSTANCE_NAME=$(kubectl get pod untrusted --output=jsonpath='{.spec.nodeName}')
```

**SSH into the worker node**

```
gcloud compute ssh ${INSTANCE_NAME}
```

**List the containers running under gVisor**

```
sudo runsc --root  /run/containerd/runsc/k8s.io list
```

**Get the ID of the untrusted pod**

```
POD_ID=$(sudo crictl -r unix:///var/run/containerd/containerd.sock \
  pods --name untrusted -q)
```

**Get the ID of the webserver container running in the untrusted pod**

```
CONTAINER_ID=$(sudo crictl -r unix:///var/run/containerd/containerd.sock \
  ps -p ${POD_ID} -q)
```
**Use the gVisor runsc command to display the processes running inside the webserver container**

```
sudo runsc --root /run/containerd/runsc/k8s.io ps ${CONTAINER_ID}
```


</p>
</details

***


