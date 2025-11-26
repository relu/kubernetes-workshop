# Kubernetes Introduction Workshop

![Kubernetes](https://img.shields.io/badge/Kubernetes-v1.34-326CE5?logo=kubernetes&logoColor=white)

![Pile of blue lego pieces](https://images.unsplash.com/photo-1646729470883-9602c07b1810?w=1200&h=300&fit=crop&auto=format&q=80)
<sup>Photo by [Charles Snow](https://unsplash.com/@charles_snow)</sup>

A hands-on, progressive workshop that takes you from Kubernetes basics to advanced deployment concepts. Learn by doing with practical examples and real applications.

## Table of Contents

- [What You'll Learn](#-what-youll-learn)
- [Prerequisites](#-prerequisites)
  - [Required Tools](#required-tools)
  - [Clone This Repository](#clone-this-repository)
- [Getting Started](#-getting-started)
  - [Set Up Your Local Cluster](#set-up-your-local-cluster)
  - [Verify Your Setup](#verify-your-setup)
- [Workshop Modules](#-workshop-modules)
  - [1. Namespaces - Your Workspace](#1--namespaces---your-workspace)
  - [2. Pods - Your First Application](#2--pods---your-first-application)
  - [3. Services - Networking Your Apps](#3--services---networking-your-apps)
  - [4. ReplicaSets - Scaling for Resilience](#4--replicasets---scaling-for-resilience)
  - [5. Deployments - Managing Updates](#5--deployments---managing-updates)
  - [6. Ingress - External Access](#6--ingress---external-access)
  - [7. Logging](#7-logging)
  - [8. Resource Management](#8-resource-management)
  - [9. Configuration Management](#9-️-configuration-management)
  - [10. Headlamp - Kubernetes UI](#10-headlamp---kubernetes-ui)
  - [11. Helm - Package Management](#11--helm---package-management)
- [Troubleshooting](#-troubleshooting)
- [Cleanup](#-cleanup)
- [Additional Resources](#-additional-resources)

## 🎯 What You'll Learn

By the end of this workshop, you will:

- Understand core Kubernetes concepts and architecture
- Deploy and manage containerized applications
- Scale applications for high availability
- Perform rolling updates and rollbacks
- Configure networking and external access
- Manage application configuration and secrets
- Monitor resource usage
- Use Helm for package management
- Apply Kubernetes best practices

**Estimated time**: 3-4 hours

## ✅ Prerequisites

**Note for beginners**: This workshop uses the command line (Terminal on Mac/Linux, PowerShell or Command Prompt on Windows). If you're new to the command line, don't worry - we'll guide you through each step!

### What is Kubernetes?

Kubernetes (often abbreviated as "K8s") is a system for automating the deployment, scaling, and management of containerized applications. Think of it as an orchestrator that manages where your applications run and ensures they stay healthy.

**Key terms you'll encounter**:
- **Cluster**: A group of computers (called nodes) that run your applications
- **Container**: A packaged application with everything it needs to run
- **Pod**: The smallest unit in Kubernetes - usually runs one container
- **Manifest**: A YAML file that describes what you want Kubernetes to create

**📚 Learn more about Kubernetes**:
- [Kubernetes Overview](https://kubernetes.io/docs/concepts/overview/) - What is Kubernetes and why use it
- [Kubernetes Components](https://kubernetes.io/docs/concepts/overview/components/) - Architecture and control plane
- [Understanding Kubernetes Objects](https://kubernetes.io/docs/concepts/overview/working-with-objects/) - Core concepts
- [What is a Container?](https://kubernetes.io/docs/concepts/containers/) - Container basics

### Required Tools

Install the following tools on your machine before starting the workshop. You'll need all of them, and they must be installed in this order.

#### Container Runtime (Docker/Colima)

Before installing Kind, you need a container runtime. Kind runs Kubernetes clusters using containers, so you need Docker or an alternative.

**First, check if you already have one installed:**

```bash
docker ps
```

**If you see**:
- An empty table with column headers (CONTAINER ID, IMAGE, etc.) → **You're all set! Skip to the Kind installation section.**
- "command not found" or "Cannot connect to Docker daemon" → **Continue below to install a container runtime.**

---

**Only install a container runtime if you don't have one already:**

**On Mac (Recommended: Colima)**:

Colima is a lightweight, open-source container runtime for Mac.

```bash
# Install Colima using Homebrew
brew install colima

# Start Colima
colima start
```

**What to expect**: Colima will download and start a Linux VM (this takes a minute or two the first time). You'll see output like:
```
INFO[0000] starting colima
INFO[0000] runtime: docker
INFO[0001] starting ...
INFO[0030] done
```

Verify it's working:
```bash
docker ps
```

You should see an empty list of containers (no errors).

**Alternative for Mac: Docker Desktop**

If you prefer, you can use [Docker Desktop for Mac](https://www.docker.com/products/docker-desktop/). Download and install it, then start Docker Desktop from your Applications folder.

**On Windows**:

**Option 1: Docker Desktop (Recommended)**

1. Download [Docker Desktop for Windows](https://www.docker.com/products/docker-desktop/)
2. Run the installer
3. During installation, ensure "Use WSL 2 instead of Hyper-V" is selected
4. After installation, start Docker Desktop from the Start menu
5. Wait for Docker to start (you'll see a green icon in the system tray)

Verify it's working in PowerShell or Command Prompt:
```powershell
docker ps
```

**Option 2: Podman Desktop**

An open-source alternative to Docker Desktop:

1. Download [Podman Desktop](https://podman-desktop.io/downloads)
2. Run the installer
3. Start Podman Desktop
4. Click "Initialize" to set up Podman

**Note for Windows**: You'll need WSL 2 (Windows Subsystem for Linux 2). Windows 10 version 2004 or higher, or Windows 11 is required. If you don't have WSL 2, follow [Microsoft's WSL installation guide](https://learn.microsoft.com/en-us/windows/wsl/install).

**On Linux**:

Install Docker Engine using your distribution's package manager:

```bash
# Ubuntu/Debian
sudo apt-get update
sudo apt-get install docker.io
sudo systemctl start docker
sudo systemctl enable docker

# Add your user to the docker group to run without sudo
sudo usermod -aG docker $USER

# Log out and back in for group changes to take effect
```

Or use [Podman](https://podman.io/getting-started/installation):

```bash
# Ubuntu/Debian
sudo apt-get update
sudo apt-get install podman
```

Verify it's working:
```bash
docker ps
```

#### Kind

Kind (Kubernetes in Docker) allows you to run Kubernetes clusters locally on your laptop for learning and development.

**What it does**: Creates a local Kubernetes cluster using Docker containers instead of virtual machines.

**On Mac (using Homebrew)**:
```bash
brew install kind
```

**On Windows (using Chocolatey)**:
```powershell
choco install kind
```

Or download the binary from [Kind releases](https://github.com/kubernetes-sigs/kind/releases) and add it to your PATH.

**On Linux**:
```bash
# Download the latest version
curl -Lo ./kind https://kind.sigs.k8s.io/dl/latest/kind-linux-amd64
chmod +x ./kind
sudo mv ./kind /usr/local/bin/kind
```

Or follow the complete installation guide [here](https://kind.sigs.k8s.io/docs/user/quick-start/#installation).

Verify Kind is installed:
```bash
kind version
```

#### Kubectl

Kubectl (pronounced "kube-control" or "kube-C-T-L") is the command-line tool for interacting with Kubernetes clusters.

**What it does**: Lets you run commands to create, view, update, and delete Kubernetes resources.

Follow the installation guide [here](https://kubernetes.io/docs/tasks/tools/install-kubectl/).

#### Helm

Helm is a package manager for Kubernetes that simplifies deployment management.

**What it does**: Packages multiple Kubernetes resources together and lets you install them as a unit.

Install using the guide [here](https://helm.sh/docs/intro/install/).

### Clone This Repository

All commands in this workshop should be run from the repository's base directory (the folder you clone).

```bash
git clone https://github.com/relu/kubernetes-workshop.git
cd kubernetes-workshop
```

**Note**: If you received a different repository URL, use that instead. The `cd` command (change directory) navigates into the workshop folder you just cloned. All subsequent commands in this workshop should be run from this directory.

## 🚀 Getting Started

**Before you begin**: Make sure your container runtime (Colima or Docker Desktop) is running! Check by running `docker ps` in your terminal. If you get an error, go back to the Container Runtime installation section.

### Understanding kubectl Commands

Before we begin, let's understand the basic kubectl command structure you'll be using:

```
kubectl <action> <resource-type> <resource-name> [options]
```

**Common actions**:
- `apply`: Create or update resources from a file
- `get`: List resources
- `describe`: Show detailed information about a resource
- `delete`: Remove resources
- `logs`: View logs from a container
- `exec`: Run commands inside a container

**Common flags** you'll see throughout:
- `-f <file>`: Read from a file (the "f" stands for "file")
- `-n <namespace>`: Specify which namespace to use
- `-l <label>`: Select resources by label (like a filter)
- `-o yaml`: Output in YAML format
- `--watch` or `-w`: Watch for changes in real-time

**Line continuation**: When you see a backslash `\` at the end of a line, it means the command continues on the next line. You can type it all on one line without the `\` if you prefer.

**📚 Learn more about kubectl**:
- [kubectl Overview](https://kubernetes.io/docs/reference/kubectl/) - Official reference
- [kubectl Cheat Sheet](https://kubernetes.io/docs/reference/kubectl/cheatsheet/) - Common commands
- [kubectl Command Reference](https://kubernetes.io/docs/reference/generated/kubectl/kubectl-commands) - Complete command list

### Set Up Your Local Cluster

Create a local Kubernetes cluster using Kind with the configuration that enables ingress support:

```bash
kind create cluster --name workshop --config manifests/00-kind-config.yaml
```

**What this command does**:
- `kind create cluster`: Creates a new Kubernetes cluster
- `--name workshop`: Names the cluster "workshop"
- `--config manifests/00-kind-config.yaml`: Uses our custom configuration file with special port settings

The configuration includes port mappings (30080, 30443) that allow external traffic to reach your cluster, which we'll need later for accessing applications from your browser.

**What to expect**: The cluster creation takes 1-2 minutes. You'll see output like:
```
Creating cluster "workshop" ...
 ✓ Ensuring node image (kindest/node:v1.27.3) 🖼
 ✓ Preparing nodes 📦
 ✓ Writing configuration 📜
 ✓ Starting control-plane 🕹️
 ✓ Installing CNI 🔌
 ✓ Installing StorageClass 💾
Set kubectl context to "kind-workshop"
```

### Verify Your Setup

Confirm all tools are properly installed and working:

```bash
# Check kubectl is configured for your cluster
kubectl cluster-info

# Verify cluster nodes are ready
kubectl get nodes

# Check Kind cluster is running
kind get clusters

# Verify Helm is installed
helm version
```

**Expected output**: You should see your cluster responding, one control-plane node in "Ready" status, "workshop" in the Kind clusters list, and Helm version information.

### Install Metrics Server

Before starting the workshop modules, let's install the Metrics Server. This component collects resource usage data (CPU and memory) from your cluster, which you'll use throughout the workshop.

```bash
# Add the metrics-server repository
helm repo add metrics-server https://kubernetes-sigs.github.io/metrics-server/
helm repo update

# Install metrics-server in the kube-system namespace
helm install metrics-server metrics-server/metrics-server \
  -n kube-system \
  --set args[0]=--kubelet-insecure-tls \
  --set args[1]=--metric-resolution=10s
```

**Note**: The `--kubelet-insecure-tls` flag is needed for local development with Kind. The `--metric-resolution=10s` flag makes metrics-server scrape more frequently (every 10 seconds instead of the default 60 seconds), giving you faster feedback during the workshop. In production, the default 60s is usually sufficient.

Wait for the metrics server to be ready (1-2 minutes):

```bash
kubectl wait --for=condition=ready pod -l app.kubernetes.io/name=metrics-server -n kube-system --timeout=300s
```

Verify it's working:

```bash
kubectl top nodes
```

You should see CPU and memory usage for your cluster node. If you see an error, wait another minute and try again - metrics take time to collect.

## 📚 Workshop Modules

### 1. 📁 Namespaces - Your Workspace

**Time**: ~10 minutes

A namespace is a virtual partition within your Kubernetes cluster. Think of it as a folder that groups related resources together.

**Why use namespaces?**
- Organize resources (e.g., separate dev, staging, and production)
- Isolate different projects or teams
- Easy cleanup - delete a namespace and all its resources are gone

**What's "namespaced"?**
- Most Kubernetes resources live in a namespace (pods, services, deployments, etc.)
- Some resources are cluster-wide and don't belong to a namespace (like nodes or the namespace itself)

**📚 Learn more**: [Kubernetes Namespaces](https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/)

We'll create a dedicated namespace called `workshop` for this tutorial. When we're done, deleting this namespace will clean up everything we created.

#### Create the namespace

```bash
kubectl create namespace workshop
```

#### List all namespaces

```bash
kubectl get ns
```

**Expected output**: You should see your `workshop` namespace with status "Active", along with system namespaces like `default`, `kube-system`, and `kube-public`.

**Note**: `ns` is a short name (alias) for `namespaces`. Kubernetes provides these shortcuts to save typing. Other examples: `po` for pods, `svc` for services, `deploy` for deployments.

#### Inspect the namespace

```bash
kubectl describe ns workshop
```

This shows a detailed view of the namespace including labels, resource quotas (if any), and status.

**Note**: You can describe any Kubernetes resource using the `kubectl describe <resource-type> <resource-name>` command. This is invaluable for troubleshooting as it shows detailed information and recent events.

#### View the namespace manifest

Even though we created the namespace using an imperative command, it's backed by a declarative manifest. Let's see it:

```bash
kubectl get -o yaml ns workshop
```

This outputs the namespace definition in YAML format, showing all fields including metadata and status.

**About YAML**: YAML is a human-readable format for representing data. It uses indentation (spaces, not tabs) to show structure. Kubernetes uses YAML files to describe what you want to create. Think of it like a recipe that tells Kubernetes how to make your resources.

**Note**: You can use `kubectl get -o yaml` on all Kubernetes objects to see their full specifications.

#### Edit a resource in place

You can modify resources directly using your editor:

```bash
kubectl edit ns workshop
```

This opens your default editor with the namespace manifest. Try adding a label under `metadata.labels`, save, and exit. The changes are applied immediately.

**Warning**: While `kubectl edit` is convenient for learning and debugging, in production you should always edit the YAML files and use `kubectl apply` instead. This way your changes are tracked in version control (like git), and you have a history of what changed.

#### Set your context to the workshop namespace

To avoid typing `-n workshop` with every command, let's make our namespace the default for the current context:

```bash
kubectl config set-context --current --namespace workshop
```

**Note**: This setting persists - even if you close your terminal and open a new one, `workshop` will still be your default namespace. Before this command, kubectl used the `default` namespace automatically.

**Verify**: Run `kubectl config get-contexts` to see the current context configuration.

#### Challenge Exercise

Try creating a second namespace called `workshop-dev` and list all namespaces. What differences do you notice?

---

### 2. 🎁 Pods - Your First Application

**Time**: ~15 minutes

Pods are the smallest deployable units in Kubernetes. A pod represents a single instance of a running process and can contain one or more containers that share networking and storage. In most cases, you'll run one container per pod.

**Understanding pods**:
- Each pod gets its own IP address
- Containers in a pod share the same network namespace (can communicate via localhost)
- A special "pause" container (infrastructure container) holds the network namespace for the pod
- Pods are designed to be ephemeral and immutable - you replace them rather than modify them

**📚 Learn more about pods**:
- [Pod Concepts](https://kubernetes.io/docs/concepts/workloads/pods/) - Official overview
- [Pod Lifecycle](https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/) - Pod phases and container states
- [Init Containers](https://kubernetes.io/docs/concepts/workloads/pods/init-containers/) - Specialized containers that run before app containers
- [Sidecar Containers](https://kubernetes.io/docs/concepts/workloads/pods/sidecar-containers/) - Multi-container pod patterns

For this workshop, we'll use four example applications written in different languages. Each application is packaged as a **container image** (think of it as a ready-to-run package containing the application and everything it needs).

- `ghcr.io/relu/example-app-python` - Python web application
- `ghcr.io/relu/example-app-ruby` - Ruby web application
- `ghcr.io/relu/example-app-go` - Go web application
- `ghcr.io/relu/example-app-rust` - Rust web application

The format `ghcr.io/relu/example-app-ruby` is a container image reference where:
- `ghcr.io` is the registry (where the image is stored, like Docker Hub)
- `relu/example-app-ruby` is the image name

**Note**: These applications are in the `apps/` directory of this repository. Feel free to explore the code and build your own images.

**📚 Learn more about container images**:
- [Container Images](https://kubernetes.io/docs/concepts/containers/images/) - How Kubernetes uses images
- [Image Pull Policy](https://kubernetes.io/docs/concepts/containers/images/#image-pull-policy) - When images are pulled
- [Container Registries](https://kubernetes.io/docs/concepts/containers/images/#image-names) - Where images are stored

**What the applications display**:
When you access these applications, they show:
- **Version**: The runtime version (Ruby/Python/Go/Rust version)
- **Request Path**: The current URL path being accessed
- **Hostname**: The pod's hostname (which is the pod name in Kubernetes)
- **Request Count**: A counter that increments with each request to that specific pod

You can customize the subtitle text by setting the `SUBTITLE` environment variable (defaults to "Kubernetes Workshop Example Application").

#### Create your first pod

```bash
kubectl apply -f manifests/01-pod.yaml
```

**What happened**: Kubernetes received the pod specification, scheduled it to a node, pulled the container image, and started the container.

#### Check if the pod is running

```bash
kubectl get po
```

**Expected output**: You should see a pod named `web-app` with status "Running" and 1/1 containers ready.

**Note**: `po` is short for `pods`. The pod name comes from the `metadata.name` field in the manifest and must be unique within the namespace.

#### Get more details

```bash
kubectl get po -o wide
```

The `-o wide` flag shows additional information like the node IP, pod IP, and which node the pod is running on.

#### Describe the pod

```bash
kubectl describe po web-app
```

This shows comprehensive information including events, container status, resource limits, and volume mounts.

#### Interact with your pod

**Port forwarding**: Connect to the pod from your local machine:

```bash
kubectl port-forward pod/web-app 3000
```

**What this does**:
- Creates a tunnel from your local computer to the pod
- `localhost:3000` on your computer now forwards traffic to port 3000 in the pod
- `localhost` means "this computer" - it's a special address (127.0.0.1) that always refers to your own machine

Open http://localhost:3000 in your browser to see the application running.

**Expected output**: You should see a simple web page showing which application is running (Ruby, Python, Go, or Rust).

Press `Ctrl+C` to stop the port-forward when you're done.

**Note**: Port forwarding works for Pod, ReplicaSet, Deployment, and Service resources. It's a quick way to test applications without exposing them publicly.

#### Execute commands in the pod

You can get a shell inside a running pod (if the container image includes one):

```bash
kubectl exec -ti web-app -- ash
```

**What this does**:
- `kubectl exec`: Execute a command in a container
- `-t` allocates a pseudo-TTY (terminal) - gives you an interactive terminal
- `-i` keeps STDIN open for interactive use - allows you to type commands
- `web-app`: The name of the pod
- `--`: Separates kubectl options from the command you want to run in the pod
- `ash`: The Alpine Linux shell (our images are based on Alpine, a lightweight Linux distribution)

Try running commands like `ps` (show running processes), `env` (show environment variables), or `ls` (list files) inside the pod.

To exit the shell, type `exit` and press Enter, or press `Ctrl+D`.

#### View the pod manifest

```bash
kubectl get po web-app -o yaml
```

Notice how Kubernetes added many fields to what we defined in `manifests/01-pod.yaml`, including status information.

#### Challenge Exercise

1. Check the pod's logs using `kubectl logs web-app`
2. Try editing the pod with `kubectl edit po web-app` and change something (like adding a label). What happens?
3. What happens if you delete the pod? Does it come back?

---

### 3. 🌐 Services - Networking Your Apps

**Time**: ~15 minutes

Pods are ephemeral - they can be deleted and recreated with new IP addresses. Services provide stable networking for pods using label selectors to automatically discover and route traffic to matching pods.

**📚 Learn more about services**:
- [Service Concepts](https://kubernetes.io/docs/concepts/services-networking/service/) - Official service documentation
- [Connecting Applications with Services](https://kubernetes.io/docs/tutorials/services/connect-applications-service/) - Step-by-step tutorial
- [DNS for Services and Pods](https://kubernetes.io/docs/concepts/services-networking/dns-pod-service/) - How service discovery works

#### Create a service

```bash
kubectl apply -f manifests/02-service.yaml
```

#### List services

```bash
kubectl get svc
```

**Expected output**: You should see a service named `web-app` with a ClusterIP and a NodePort.

**Service types explained**:
- **ClusterIP** (default): Only accessible within the cluster
- **NodePort**: Exposes the service on each node's IP at a static port, accessible from outside the cluster
- **LoadBalancer**: Provisions a cloud provider's load balancer (not applicable for local Kind clusters)

Our service is type `NodePort`, which creates a high-numbered port (30000-32767) on the host.

#### Understand how services find pods

Services use **label selectors** to discover pods. Labels are key-value pairs attached to Kubernetes resources (like tags) that help organize and select them.

**Understanding labels**:
- Labels are metadata you attach to objects (like `app: web-app` or `version: v1`)
- Label selectors let you find resources with specific labels
- Services use label selectors to find which pods they should send traffic to

Let's see this in action:

1. Look at the service manifest: `kubectl get svc web-app -o yaml`
2. Find the `selector` section - it matches pods with label `app: web-app`
3. Look at the pod labels: `kubectl get po --show-labels`

Any pod with matching labels will automatically become an endpoint for this service. This is powerful because pods can come and go, but the service always finds the right ones based on labels.

#### View service endpoints

```bash
kubectl get endpointslices -l kubernetes.io/service-name=web-app
```

or simply describe the service to see endpoints:

```bash
kubectl describe svc web-app
```

You'll see the pod's IP address listed as an endpoint. When you have multiple pods, all their IPs will be listed here, and the service will load balance traffic between them.

#### Connect to your service

```bash
kubectl port-forward svc/web-app 3000
```

Now http://localhost:3000 forwards to the service, which proxies to the pod. This looks the same as before, but now we're going through the service abstraction.

**Why is this useful?** The service provides a stable network identity. Pods can come and go, but the service endpoint remains constant.

#### Challenge Exercise

1. Add another label to your pod: `kubectl label po web-app version=v1`
2. What happens if you change the service selector to look for a non-existent label?
3. Check the endpoints - are there any?

---

### 4. 📈 ReplicaSets - Scaling for Resilience

**Time**: ~15 minutes

Running a single pod has limitations - if it crashes or gets deleted, your application goes down. ReplicaSets solve this by maintaining a specified number of identical pod replicas running at all times.

**📚 Learn more about ReplicaSets**:
- [ReplicaSet Concepts](https://kubernetes.io/docs/concepts/workloads/controllers/replicaset/) - Official documentation
- [Labels and Selectors](https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/) - How Kubernetes uses labels to organize resources

#### Create a ReplicaSet

```bash
kubectl apply -f manifests/03-replicaset.yaml
```

This ReplicaSet is configured to maintain 3 pod replicas.

#### Check the ReplicaSet

```bash
kubectl get rs
```

**Expected output**: You should see a ReplicaSet named `web-app` with 3 desired, 3 current, and 3 ready replicas.

#### List the pods

```bash
kubectl get pod
```

**What happened?** You might notice we only have 3 pods total, not 4. The ReplicaSet "adopted" our original `web-app` pod (because it matches the label selector) and created 2 additional pods to reach the desired count of 3.

**ReplicaSet pod naming**: Pods managed by a ReplicaSet have generated names like `web-app-xxxxx` where `xxxxx` is a random suffix automatically generated by Kubernetes. This ensures each pod has a unique name.

#### Test self-healing

Let's delete our original static pod and see what happens:

```bash
kubectl delete -f manifests/01-pod.yaml
```

**Note**: You can delete resources using either the manifest file or the resource name directly: `kubectl delete pod web-app`

Now check the pods again:

```bash
kubectl get po
```

We still have 3 pods! The ReplicaSet detected that a pod was deleted and immediately created a new one to maintain the desired count.

#### Check service endpoints

```bash
kubectl describe svc web-app
```

You should now see 3 endpoint IP addresses corresponding to the 3 pods. The service automatically load balances traffic across all matching pods.

#### Observe multiple pods

You can see logs from all pods to verify they're all running:

```bash
kubectl logs -f -l app=web-app --prefix
```

The `-f` flag streams logs in real-time, `-l` filters by label, and `--prefix` shows which pod generated each log line. Press `Ctrl+C` to stop following the logs.

**Tip**: Consider opening a separate terminal window for log streaming so you can continue working in your main terminal.

**Note about load balancing**: Port-forwarding (`kubectl port-forward`) connects to a single pod, even when targeting a service, so it won't demonstrate load balancing. True load balancing happens when traffic comes through the service from within the cluster, or later when we set up the Ingress controller which will distribute requests across all healthy pods.

#### Scale manually

You can change the replica count:

```bash
kubectl scale rs web-app --replicas 5
```

Check the pods: `kubectl get po`

Scale back down:

```bash
kubectl scale rs web-app --replicas 3
```

**What happens when scaling down?** Kubernetes terminates excess pods gracefully, giving them time to finish processing requests.

#### Challenge Exercise

1. Delete one pod manually: `kubectl delete pod <pod-name>`. How long does it take for a replacement to appear?
2. Scale to 10 replicas. Check resource usage with `kubectl top pod`
3. What happens if you create a standalone pod with the same labels as the ReplicaSet's selector?

---

### 5. 🔄 Deployments - Managing Updates

**Time**: ~20 minutes

ReplicaSets are great for maintaining pod replicas, but they don't handle updates well. If you want to roll out a new version, you'd need to manually create a new ReplicaSet, scale it up while scaling down the old one, and then clean up. Deployments automate this process.

A Deployment is a higher-level controller that manages ReplicaSets. It provides:
- Declarative updates with rolling deployments
- Rollback capabilities
- Revision history
- Pausing and resuming rollouts

**📚 Learn more about Deployments**:
- [Deployment Concepts](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/) - Official documentation
- [Performing a Rolling Update](https://kubernetes.io/docs/tutorials/kubernetes-basics/update/update-intro/) - Interactive tutorial
- [Deployment Strategies](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/#strategy) - Rolling update vs recreate

#### Create a deployment

```bash
kubectl apply -f manifests/04-deployment.yaml
```

**Note**: The Deployment manifest looks almost identical to the ReplicaSet manifest, except for the `kind: Deployment` field. This is because Deployments create and manage ReplicaSets under the hood.

The manifest also includes `revisionHistoryLimit: 10`, which tells Kubernetes to keep the last 10 ReplicaSets for rollback purposes. The default is also 10, but it's good practice to set it explicitly.

#### Check the deployment

```bash
kubectl get deployments -o wide
```

**Expected output**: You should see a deployment named `web-app` with 3/3 ready replicas and the current image being used.

#### Look at what was created

```bash
kubectl get rs
```

You'll now see two ReplicaSets:
1. The old one we created directly (scaled to 0)
2. A new one created by the Deployment (with 3 replicas)

```bash
kubectl get pod
```

The pods remain at 3 because the Deployment's selector matches the existing pods, preventing the creation of duplicates.

#### Understand the hierarchy

```
Deployment
    └─ ReplicaSet
        └─ Pod
        └─ Pod
        └─ Pod
```

The Deployment manages ReplicaSets, which in turn manage Pods.

#### Perform a rolling update

Let's update our application to use the Python image instead of Ruby:

```bash
kubectl set image deployment/web-app app=ghcr.io/relu/example-app-python
```

**What's happening**: Kubernetes is performing a rolling update:
1. Creates a new ReplicaSet with the new image
2. Gradually scales up the new ReplicaSet
3. Simultaneously scales down the old ReplicaSet
4. Ensures the desired number of pods are always available

Watch the rollout in real-time:

```bash
kubectl rollout status deployment web-app
```

**Tip**: Open a second terminal window/tab to watch the pods change in real-time while the rollout happens:

```bash
kubectl get po -w
```

The `-w` flag watches for changes. You'll see old pods terminating and new pods starting. Press `Ctrl+C` to stop watching.

#### Examine the ReplicaSets

```bash
kubectl get rs
```

You should now see two ReplicaSets:
- The old one (Ruby image) scaled to 0 replicas
- The new one (Python image) with 3 replicas

The old ReplicaSet is kept for rollback purposes.

#### View rollout history

```bash
kubectl rollout history deployment web-app
```

This shows all revisions of your deployment. Each update creates a new revision.

**Understanding revision history**: Kubernetes keeps the old ReplicaSets around (scaled to 0 replicas) to enable rollbacks. The `revisionHistoryLimit` field in the deployment spec controls how many old ReplicaSets to keep. We've set it to 10 in our manifests, meaning you can rollback through the last 10 changes. When you exceed this limit, the oldest revisions are automatically deleted.

You can view detailed information about a specific revision:

```bash
kubectl rollout history deployment web-app --revision=1
```

This shows the pod template used in that revision, including the image and other configuration.

#### Rollback to the previous version

Made a mistake? Easy to rollback:

```bash
kubectl rollout undo deployment web-app
```

This reverts to the previous revision. Check the pods and you'll see the Ruby image is back.

#### Rollback to a specific revision

First, check the available revisions:

```bash
kubectl rollout history deployment web-app
```

You'll see a list of revisions with their change causes. Each update or rollback creates a new revision.

To rollback to a specific revision, use one of the revision numbers shown in the history output. For example, if your history shows revisions 5, 6, and 7, you can rollback to revision 5:

```bash
kubectl rollout undo deployment web-app --to-revision=5
```

You can view details about what changed in a specific revision before rolling back:

```bash
kubectl rollout history deployment web-app --revision=5
```

**Note**: Kubernetes keeps up to `revisionHistoryLimit` (configured: 10) old ReplicaSets for rollback. Our deployment manifests explicitly set this value to ensure you can rollback through your deployment history.

**Important**: If you see fewer than expected revisions in the history, it's because `revisionHistoryLimit` only prevents *future* deletions of old ReplicaSets. Any ReplicaSets that were already cleaned up before setting this limit cannot be recovered. This is normal if you're working through the workshop - future updates will be retained up to the limit.

#### Scale the deployment

Scaling works the same as with ReplicaSets:

```bash
kubectl scale deployment web-app --replicas 5
```

Check the pods:

```bash
kubectl get po
```

The Deployment updated its ReplicaSet, which created the additional pods.

Scale back:

```bash
kubectl scale deployment web-app --replicas 3
```

#### Update to the Go image

Let's try one more update:

```bash
kubectl set image deployment/web-app app=ghcr.io/relu/example-app-go
```

**In practice**: Instead of using `kubectl set image`, you would typically update the manifest file and run `kubectl apply -f manifests/04-deployment.yaml`. This is the declarative approach and is preferred for production use.

#### Pause and resume rollouts

You can pause a rollout to make multiple changes:

```bash
kubectl rollout pause deployment web-app
```

Make changes (they won't be applied yet), then resume:

```bash
kubectl rollout resume deployment web-app
```

All changes are applied together in a single rollout.

#### Challenge Exercise

1. Perform a rolling update with an invalid image name. What happens? How do you recover?
2. Check the deployment events: `kubectl describe deployment web-app`
3. Try pausing a rollout mid-way, making additional changes, then resuming. How does this differ from multiple sequential updates?

---

### 6. 🚪 Ingress - External Access

**Time**: ~20 minutes

So far, we've used port-forwarding to access our applications locally. In production, you need a proper way to route external traffic to your services. Ingress provides HTTP and HTTPS routing from outside the cluster to services within the cluster.

An Ingress requires an **Ingress Controller** - a component that watches for Ingress resources and configures a reverse proxy accordingly. We'll use **Traefik**, a modern, cloud-native ingress controller.

**📚 Learn more about Ingress**:
- [Ingress Concepts](https://kubernetes.io/docs/concepts/services-networking/ingress/) - Official documentation
- [Ingress Controllers](https://kubernetes.io/docs/concepts/services-networking/ingress-controllers/) - Available ingress controller options
- [Traefik Documentation](https://doc.traefik.io/traefik/providers/kubernetes-ingress/) - Traefik as a Kubernetes ingress controller

#### Install Traefik Ingress Controller

Traefik can be installed using Helm (we'll learn more about Helm later):

```bash
# Add the Traefik Helm repository
helm repo add traefik https://traefik.github.io/charts

# Update the repository
helm repo update

# Install Traefik in the kube-system namespace
helm install traefik traefik/traefik \
  --namespace kube-system \
  --set ports.web.hostPort=30080 \
  --set ports.websecure.hostPort=30443
```

**What this does**: Installs Traefik with hostPort bindings on ports 30080 and 30443 (which don't require root permissions). These ports are forwarded to your host machine via the Kind cluster's `extraPortMappings` configuration.

#### Verify Traefik is running

```bash
kubectl get pods -n kube-system -l app.kubernetes.io/name=traefik
```

**Expected output**: You should see a Traefik pod in "Running" status.

Wait for it to be ready (this may take 1-2 minutes):

```bash
kubectl wait --for=condition=ready pod -l app.kubernetes.io/name=traefik -n kube-system --timeout=300s
```

#### Create the Ingress resource

```bash
kubectl apply -f manifests/05-ingress.yaml
```

#### Check the Ingress

```bash
kubectl get ingress
```

**Expected output**: You should see an ingress named `web-app` with a host of `*` (matches any hostname).

For more details:

```bash
kubectl describe ingress web-app
```

Look for the "Rules" section showing the routing configuration and the "Backend" section showing which service receives the traffic.

#### Access your application through Ingress

Since we're running locally with Kind, we can access the application at:

```
http://localhost:30080
```

Open this URL in your browser. You should see your application!

**What's happening**:
1. Your browser sends a request to localhost:30080
2. Kind forwards port 30080 to the Traefik ingress controller
3. Traefik matches the request against Ingress rules
4. The request is routed to the `web-app` service
5. The service load balances to one of the pods
6. The pod responds and the response travels back

#### Test with different paths

The Ingress is configured with `pathType: Prefix` and `path: /`, so all paths are routed to the service. Try:

```
http://localhost:30080/test
http://localhost:30080/api
http://localhost:30080/anything
```

All should work (though your application might return 404 for undefined routes).

#### Advanced: Path-based routing

You can route different paths to different services. Let's deploy our four example applications (Ruby, Python, Go, Rust) as separate deployments, each with their own service, and route to them based on URL paths.

**Note**: The manifest creates four new deployments (web-app-ruby, web-app-python, web-app-go, web-app-rust) with corresponding services. Your existing `web-app` deployment and service will remain unchanged and will handle the root path (`/`).

First, deploy the four applications with their own services:

```bash
kubectl apply -f manifests/05-multiapp-deployments.yaml
```

This creates:
- 4 Deployments (web-app-ruby, web-app-python, web-app-go, web-app-rust)
- 4 Services (web-app-ruby, web-app-python, web-app-go, web-app-rust)

Wait for all pods to be ready:

```bash
kubectl get pods -l 'app in (web-app-ruby,web-app-python,web-app-go,web-app-rust)'
```

Check all services:

```bash
kubectl get svc
```

You should see five services: `web-app` (original), `web-app-ruby`, `web-app-python`, `web-app-go`, and `web-app-rust`.

Now update the Ingress to route based on paths:

```bash
kubectl apply -f manifests/05-ingress-multipath.yaml
```

Check the updated Ingress:

```bash
kubectl describe ingress web-app
```

You should see multiple path rules configured.

Test the different paths:

```bash
curl http://localhost:30080/ruby
curl http://localhost:30080/python
curl http://localhost:30080/go
curl http://localhost:30080/rust
curl http://localhost:30080/
```

Each path routes to a different backend service! The `/` path acts as a catch-all for requests that don't match the more specific paths.

#### Advanced: Host-based routing (Optional)

You can also route based on hostnames instead of paths. This is useful when you want each application to have its own domain name.

**Prerequisites**: Make sure you've deployed the multi-app setup from the path-based routing section above.

**Step 1: Configure /etc/hosts**

You need to add local DNS entries to map hostnames to localhost.

**On Linux/Mac**, edit `/etc/hosts` (requires sudo):

```bash
sudo nano /etc/hosts
```

Add these lines at the end of the file:

```
127.0.0.1 ruby.local
127.0.0.1 python.local
127.0.0.1 go.local
127.0.0.1 rust.local
127.0.0.1 app.local
```

Save and exit (in nano: `Ctrl+O`, `Enter`, then `Ctrl+X`).

**On Windows**, edit `C:\Windows\System32\drivers\etc\hosts` as Administrator and add:

```
127.0.0.1 ruby.local
127.0.0.1 python.local
127.0.0.1 go.local
127.0.0.1 rust.local
127.0.0.1 app.local
```

**Step 2: Apply host-based Ingress**

```bash
kubectl apply -f manifests/05-ingress-hostbased.yaml
```

**Step 3: Verify the configuration**

```bash
kubectl describe ingress web-app
```

You should see five host rules instead of path rules.

**Step 4: Test the host-based routing**

```bash
curl http://ruby.local:30080/
curl http://python.local:30080/
curl http://go.local:30080/
curl http://rust.local:30080/
curl http://app.local:30080/
```

Each hostname routes to a different service! You can also open these URLs in your browser.

**How it works**:
1. Your browser/curl resolves `ruby.local` to `127.0.0.1` (localhost) via `/etc/hosts`
2. The request reaches Traefik on port 30080
3. Traefik inspects the `Host` header and routes to the appropriate service
4. Different hostnames → different backend services

**To revert back**: Apply the original ingress or the path-based ingress:

```bash
# Back to single-app ingress
kubectl apply -f manifests/05-ingress.yaml

# Or back to path-based routing
kubectl apply -f manifests/05-ingress-multipath.yaml
```

#### Challenge Exercise

1. Try the path-based routing example above with all three applications
2. Modify `manifests/05-ingress-multipath.yaml` to add a `/test` path that routes to the Ruby service
3. (Optional) Try the host-based routing setup above - it's a great way to see how virtual hosting works
4. Use `kubectl logs` to see which pods handle requests to different paths or hosts
5. Combine both! Create an Ingress with both host-based AND path-based rules (e.g., `ruby.local/api`, `python.local/data`)

---

### 7. Logging

**Time**: ~10 minutes

Kubernetes provides built-in access to container logs through kubectl. This is essential for debugging and monitoring applications.

**📚 Learn more about logging**:
- [Logging Architecture](https://kubernetes.io/docs/concepts/cluster-administration/logging/) - How Kubernetes handles logs
- [Application Introspection and Debugging](https://kubernetes.io/docs/tasks/debug/debug-application/) - Debugging techniques
- [Logging Best Practices](https://kubernetes.io/docs/concepts/cluster-administration/logging/#logging-at-the-node-level) - Node-level logging patterns

#### View logs from all pods

```bash
kubectl logs -l app=web-app
```

The `-l` flag uses a label selector to get logs from all pods matching `app=web-app`.

#### Follow logs in real-time

```bash
kubectl logs -f -l app=web-app
```

The `-f` flag streams logs as they're generated, similar to `tail -f`.

Press `Ctrl+C` to stop following.

**Tip**: When following logs, it's helpful to keep them running in a separate terminal window so you can see real-time output while working in another terminal.

#### View logs from a specific pod

```bash
kubectl logs <pod-name>
```

#### View logs with timestamps

```bash
kubectl logs <pod-name> --timestamps
```

#### View recent logs

Show the last 20 lines:

```bash
kubectl logs <pod-name> --tail=20
```

Show logs from the last hour:

```bash
kubectl logs <pod-name> --since=1h
```

#### Logs from previous container instance

If a container crashed and restarted, you can view the previous container's logs:

```bash
kubectl logs <pod-name> --previous
```

This is crucial for debugging crash loops.

#### Logs from multi-container pods

If a pod has multiple containers, specify which one:

```bash
kubectl logs <pod-name> -c <container-name>
```

List containers in a pod:

```bash
kubectl get pod <pod-name> -o jsonpath='{.spec.containers[*].name}'
```

#### Streaming logs from multiple pods

The `--prefix` flag shows which pod each log line comes from:

```bash
kubectl logs -f -l app=web-app --prefix
```

This is useful when you have multiple replicas and want to see which pod is handling requests.

#### Generate some logs

Generate traffic to your application to see logs. **Open two terminal windows** for this exercise:

**Terminal 1** - Follow logs:
```bash
kubectl logs -f -l app=web-app --prefix
```

**Terminal 2** - Make requests:
```bash
curl http://localhost:30080/
curl http://localhost:30080/test
curl http://localhost:30080/api
```

**Note**: `curl` is a command-line tool for making HTTP requests. `http://localhost:30080/` means "make a request to port 30080 on this computer". You can also open these URLs in your web browser instead of using curl.

You should see HTTP request logs appearing in Terminal 1 from different pods.

#### Logging best practices

1. **Log to stdout/stderr** - Kubernetes captures this automatically
2. **Use structured logging** (JSON) for easier parsing
3. **Include correlation IDs** for tracing requests across services
4. **Don't log sensitive information** (passwords, tokens, PII)
5. **Use appropriate log levels** (debug, info, warn, error)
6. **For production, use a log aggregation system** (ELK stack, Loki, Splunk)

#### Limitations of kubectl logs

- Logs are lost when pods are deleted
- Logs are limited to the most recent entries (typically 10MB per container)
- No log aggregation across pods or time ranges
- No advanced querying or analysis

For production systems, deploy a centralized logging solution like:
- **ELK Stack** (Elasticsearch, Logstash, Kibana)
- **Grafana Loki** (lightweight, Prometheus-like)
- **Fluentd** + backend of choice
- Cloud provider solutions (CloudWatch Logs, Stackdriver, etc.)

#### Challenge Exercise

1. Generate an error in your application (e.g., request a non-existent route) and find it in the logs
2. Simulate a pod crash and view the previous container's logs
3. Use `kubectl logs` with `--since-time` to view logs from a specific timestamp

---

### 8. Resource Management

**Time**: ~20 minutes

Kubernetes needs to know how much CPU and memory each pod needs to make smart scheduling decisions and prevent resource exhaustion. This section covers resource requests and limits.

**📚 Learn more about resource management**:
- [Resource Management for Pods and Containers](https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/) - Official documentation
- [Quality of Service Classes](https://kubernetes.io/docs/concepts/workloads/pods/pod-qos/) - QoS classes explained
- [Resource Quotas](https://kubernetes.io/docs/concepts/policy/resource-quotas/) - Limiting resource usage per namespace
- [Understanding Kubernetes CPU Throttling](https://www.datadoghq.com/blog/kubernetes-cpu-requests-limits/) - Deep dive into CPU limits debate

#### View resource usage

```bash
kubectl top pod
```

**Expected output**: CPU and memory usage for each pod. If you see an error, wait a minute and try again - metrics take time to collect.

You can also check node-level resource usage:

```bash
kubectl top nodes
```

#### Understand resource requests and limits

There are two types of resource constraints:

**Requests**: The amount of resources guaranteed to a container. The scheduler uses requests to decide which node can run the pod.

**Limits**: The maximum amount of resources a container can use. Exceeding limits results in throttling (CPU) or termination (memory).

```yaml
resources:
  requests:
    cpu: 0.1        # 100 millicores (10% of one CPU core)
    memory: 50Mi    # 50 Mebibytes
  limits:
    cpu: 0.25       # 250 millicores (25% of one CPU core)
    memory: 100Mi   # 100 Mebibytes
```

**CPU units**:
- `1` CPU = `1000m` (millicores)
- `0.1` CPU = `100m` = 10% of one CPU core
- `0.5` CPU = `500m` = half of one CPU core

**Memory units**:
- `Mi` = Mebibyte = 1024 × 1024 bytes
- `Gi` = Gibibyte = 1024 × 1024 × 1024 bytes
- `M` = Megabyte = 1000 × 1000 bytes (less common)

**Why set both?**
- **Requests**: Ensures your pod gets scheduled on a node with enough resources
- **Limits**: Prevents a pod from consuming too much and starving other pods

**Note on CPU limits in production**: Many production environments choose **not to set CPU limits**, only CPU requests. This is because CPU limits can cause unnecessary throttling even when CPU is available on the node, leading to degraded application performance. Memory limits are still important since exceeding memory causes OOM kills. Read more about this debate: [Understanding Kubernetes CPU throttling](https://www.datadoghq.com/blog/kubernetes-cpu-requests-limits/)

#### Apply resource constraints

We have a deployment manifest with resource settings already prepared:

```bash
kubectl apply -f manifests/06-deployment-with-resources.yaml
```

This triggers a rolling update. The new pods will have resource requests and limits applied.

#### Verify the resource settings

This shows both limits and requests sections:

```bash
kubectl get pod -l app=web-app -o yaml | grep -A 6 "resources:"
```

#### Observe resource behavior

Monitor resource usage continuously using the `watch` command:

```bash
watch kubectl top pod
```

This refreshes the resource usage every 2 seconds. Press `Ctrl+C` to stop.

**Tip**: This is another good command to run in a separate terminal window so you can observe resource changes while performing actions in your main terminal.

**Note**: On Mac, you may need to install `watch` with Homebrew: `brew install watch`. Alternatively, use this one-liner:

```bash
while true; do clear; kubectl top pod; sleep 2; done
```

#### What happens when limits are exceeded?

**CPU**: The container is throttled. It won't be terminated, but it can't use more CPU than the limit.

**Memory**: If a container exceeds its memory limit, Kubernetes kills it with an OOMKilled (Out Of Memory) status. The pod's restart policy determines what happens next.

#### What happens when you exceed limits? (Optional demonstration)

If you want to see OOMKilled in action, you can force a pod to exceed its memory limit gradually so you can observe the behavior.

**Open two terminal windows for this exercise:**

**Terminal 1** - Watch memory usage in real-time:
```bash
watch kubectl top pod
```

**Terminal 2** - Gradually allocate memory (the `/dev/shm` directory is a tmpfs filesystem that lives in RAM):
```bash
kubectl exec -ti <pod-name> -- sh -c 'i=0; while [ $i -lt 200 ]; do dd if=/dev/zero of=/dev/shm/fill bs=1M count=1 seek=$i 2>/dev/null; i=$((i+5)); sleep 1; done'
```

This gradually writes 5MB at a time with a 1-second pause between writes. Watch Terminal 1 - you'll see memory usage climb until the pod hits the 100Mi limit and gets killed.

After the pod is killed, check its status:

```bash
kubectl get pods
kubectl describe pod <pod-name>
```

Look for `OOMKilled` (Out Of Memory Killed) in the pod's status and events. The pod will automatically restart due to the default restart policy.

**Note**: This is just for educational purposes to understand Kubernetes behavior. In production, properly tune your limits based on actual application requirements.

#### Resource management best practices

1. **Always set requests and limits** for production workloads
2. **Requests should be realistic** based on actual usage
3. **Limits should provide headroom** but prevent runaway processes
4. **Monitor actual usage** with metrics to tune settings
5. **Consider Quality of Service (QoS) classes**:
   - **Guaranteed**: requests = limits (highest priority)
   - **Burstable**: requests < limits (medium priority)
   - **BestEffort**: no requests or limits (lowest priority)

Check a pod's QoS class:

```bash
kubectl get pod <pod-name> -o jsonpath='{.status.qosClass}'
```

#### Challenge Exercise

1. Set unrealistic resource requests (e.g., `cpu: 100`). What happens when you try to schedule the pod?
2. Create a deployment with no limits. Is this good practice? Why or why not?
3. Monitor `kubectl top pod` over time. Do your resource settings match actual usage?

---

### 9. ⚙️ Configuration Management

**Time**: ~25 minutes

Applications need configuration - database URLs, API keys, feature flags, etc. Hard-coding these values in your container images is inflexible and insecure. Kubernetes provides ConfigMaps for general configuration and Secrets for sensitive data.

**📚 Learn more about configuration management**:
- [ConfigMaps](https://kubernetes.io/docs/concepts/configuration/configmap/) - Official ConfigMap documentation
- [Secrets](https://kubernetes.io/docs/concepts/configuration/secret/) - Official Secrets documentation
- [Good Practices for Kubernetes Secrets](https://kubernetes.io/docs/concepts/security/secrets-good-practices/) - Security best practices
- [Encrypting Secret Data at Rest](https://kubernetes.io/docs/tasks/administer-cluster/encrypt-data/) - How to enable encryption

#### Apply the configuration manifest

```bash
kubectl apply -f manifests/07-configuration.yaml
```

**Note**: Kubernetes manifests can contain multiple object definitions separated by `---` delimiters. This manifest creates three ConfigMaps, one Secret, and updates our Deployment.

#### ConfigMaps

ConfigMaps store non-confidential key-value pairs. Let's explore the three ConfigMaps we created:

```bash
kubectl get cm
```

**Expected output**: You should see three ConfigMaps: `web-app-reference`, `web-app-environment`, and `web-app-file`, plus the `kube-root-ca.crt` (created automatically by Kubernetes).

#### View ConfigMap contents

```bash
kubectl describe cm web-app-reference
kubectl describe cm web-app-environment
kubectl describe cm web-app-file
```

Notice how each stores data differently:
- `web-app-reference`: Single key-value pair
- `web-app-environment`: Multiple key-value pairs
- `web-app-file`: File-like data with newlines

#### Three ways to use ConfigMaps

Our deployment demonstrates three methods of consuming ConfigMap data:

**1. Individual environment variables**

```yaml
env:
  - name: OTHER_NAME
    valueFrom:
      configMapKeyRef:
        name: web-app-reference
        key: somekey
```

**2. All keys as environment variables**

```yaml
envFrom:
  - configMapRef:
      name: web-app-environment
```

**3. Mounted as files in a volume**

```yaml
volumes:
  - name: config-volume
    configMap:
      name: web-app-file
volumeMounts:
  - name: config-volume
    mountPath: /tmp/config
```

#### Verify ConfigMap usage

Describe one of the pods to see how ConfigMaps are attached:

```bash
kubectl describe pod -l app=web-app
```

Look for the "Environment", "Mounts", and "Volumes" sections.

**Note**: The `-l app=web-app` flag is a label selector (the "l" stands for "label"). It filters to show only pods with the `app=web-app` label, which matches all our deployment's pods. This is much faster than looking through all pods manually!

#### Exec into a pod to see the configuration

```bash
kubectl exec -ti deployment/web-app -- sh
```

Inside the pod, check environment variables:

```sh
env | grep -E 'OTHER_NAME|NAME|ENV_VAR'
```

Check the mounted configuration file:

```sh
cat /tmp/config/file.txt
```

Type `exit` and press Enter (or press `Ctrl+D`) to leave the pod shell.

#### Secrets

Secrets are similar to ConfigMaps but designed for sensitive data like passwords, tokens, and keys.

**Security warning**: Secrets are only base64-encoded by default, not encrypted. Anyone with access to the manifest or API can decode them. For production:
- Use RBAC to restrict secret access
- Enable encryption at rest in etcd
- Consider external secret management (e.g., HashiCorp Vault, AWS Secrets Manager)

#### View secrets

```bash
kubectl get secrets
```

You'll see the secret we created plus any default secrets in the namespace.

#### Create secrets in manifests using stringData

The easiest way to create secrets in manifests is using `stringData`, which accepts plain text:

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: my-secret
type: Opaque
stringData:
  username: admin
  password: secret123
```

Kubernetes automatically converts `stringData` to base64-encoded `data` when you apply the manifest.

#### Create secrets from command line

For quick testing or imperative creation:

```bash
kubectl create secret generic db-secret \
  --from-literal=username=admin \
  --from-literal=password=secret123
```

Or from files:

```bash
kubectl create secret generic db-secret-from-file \
  --from-file=username=username.txt \
  --from-file=password=password.txt
```

#### Secret types

The `type: Opaque` means the secret contains arbitrary user-defined data. Other types include:

- `kubernetes.io/service-account-token`: Service account tokens
- `kubernetes.io/dockerconfigjson`: Docker registry credentials
- `kubernetes.io/tls`: TLS certificates and keys
- `kubernetes.io/basic-auth`: Basic authentication credentials

#### Use secrets in pods

Secrets can be used the same three ways as ConfigMaps:

1. Individual environment variables
2. All keys as environment variables
3. Mounted as files

Example with environment variables:

```yaml
env:
  - name: DB_PASSWORD
    valueFrom:
      secretKeyRef:
        name: my-secret
        key: password
```

#### Update configuration

When you update a ConfigMap or Secret:
- Environment variables do NOT update automatically (requires pod restart)
- Mounted volumes DO update automatically (but with a delay of ~1 minute)

Try updating a ConfigMap:

```bash
kubectl edit cm web-app-file
```

Change the file content, save, and exit. Then check the mounted file in a pod:

```bash
kubectl exec -ti deployment/web-app -- cat /tmp/config/file.txt
```

It may take up to a minute to see the change.

#### Configuration best practices

1. **Use ConfigMaps for non-sensitive config**
2. **Use Secrets for sensitive data** (passwords, tokens, keys)
3. **Don't commit secrets to version control** (use templating or external secret management)
4. **Prefer volume mounts over environment variables** for large configurations
5. **Use separate ConfigMaps/Secrets per environment** (dev, staging, production)
6. **Version your ConfigMaps** (e.g., `app-config-v1`, `app-config-v2`) for safer updates

#### Challenge Exercise

1. Create a ConfigMap with multiple files and mount them all
2. Create a Secret for database credentials and use it in your deployment
3. Update a ConfigMap and observe how long it takes for the change to appear in a mounted volume
4. What happens if you reference a non-existent ConfigMap in a pod? (Hint: try it!)

---

### 10. Headlamp - Kubernetes UI

**Time**: ~15 minutes

Headlamp is a modern, lightweight Kubernetes UI that runs on your local machine. Unlike the traditional Kubernetes Dashboard, Headlamp is designed for developers and provides an excellent user experience for local development and cluster management.

**Why Headlamp?**
- Runs locally on your machine (not in-cluster)
- Modern, fast, and responsive interface
- Uses your kubeconfig automatically - no token management needed
- Great for development workflows
- Open source and actively maintained

**📚 Learn more about Headlamp**:
- [Headlamp Documentation](https://headlamp.dev/docs/) - Official documentation
- [Headlamp GitHub](https://github.com/headlamp-k8s/headlamp) - Source code and issues

#### Install Headlamp

**On Mac (using Homebrew)**:
```bash
brew install headlamp
```

**On Linux (using package managers)**:

```bash
# Debian/Ubuntu
wget https://github.com/headlamp-k8s/headlamp/releases/latest/download/Headlamp_amd64.deb
sudo dpkg -i Headlamp_amd64.deb

# Fedora/RHEL
wget https://github.com/headlamp-k8s/headlamp/releases/latest/download/Headlamp.x86_64.rpm
sudo rpm -i Headlamp.x86_64.rpm
```

**On Windows**:

Download the installer from [Headlamp Releases](https://github.com/headlamp-k8s/headlamp/releases) and run it.

**Important Note**: The desktop applications for Windows and Mac are currently unsigned because Headlamp recently moved under Kubernetes SIG UI. Your operating system may warn about or block the application:
- **On Windows**: Click "More info" then "Run anyway" when prompted
- **On Mac**: If blocked, run `xattr -dr com.apple.quarantine /Applications/Headlamp.app` in Terminal before launching

Once the signing process is complete, updated signed versions will be released. Linux packages are unaffected.

**Alternative: Run in-cluster (optional)**:

If you prefer, you can also deploy Headlamp in your cluster using Helm, but for this workshop, the desktop app is simpler.

#### Launch Headlamp

After installation, launch Headlamp:

**On Mac/Linux**:
```bash
headlamp
```

Or launch it from your applications menu.

**On first launch**: Headlamp will automatically detect your kubeconfig and show available clusters. Select the `kind-workshop` cluster.

**What you'll see**: Headlamp opens in your default browser at `http://localhost:4466` (or similar).

#### Explore Headlamp

Headlamp provides an intuitive interface with several key sections:

**1. Cluster Overview**
- Dashboard showing cluster health
- Node status and resource usage
- Recent events and warnings

**2. Workloads**
- View and manage Pods, Deployments, ReplicaSets, StatefulSets, DaemonSets, Jobs, CronJobs
- Real-time status updates
- Quick actions (scale, restart, delete)

**3. Network**
- Services and their endpoints
- Ingresses with routing rules
- Network Policies

**4. Storage**
- PersistentVolumes and PersistentVolumeClaims
- StorageClasses
- Volume usage and status

**5. Configuration**
- ConfigMaps with content preview
- Secrets (values are hidden for security)
- Easy editing and creation

#### Navigate to your workshop resources

1. **Select namespace**: Use the namespace dropdown at the top and select `workshop`
2. **View Deployments**: Click "Workloads" → "Deployments" to see your `web-app` deployment
3. **Inspect a deployment**: Click on `web-app` to see:
   - Pod status and replicas
   - Container images
   - Resource requests/limits
   - Revision history
   - Events
4. **View pods**: Click on a pod name to see detailed information including logs and resource usage
5. **Check services**: Navigate to "Network" → "Services" to see your services and endpoints

#### Useful Headlamp features

**Live logs**:
- Click on any pod
- Go to the "Logs" tab
- View live streaming logs with filtering and search
- Toggle between containers in multi-container pods

**Resource metrics**:
- View CPU and memory usage graphs (requires Metrics Server)
- See resource utilization trends
- Identify resource-hungry pods

**Shell access**:
- Click on a pod
- Click the "Terminal" icon or tab
- Get an interactive shell inside the container
- No need for kubectl exec commands!

**YAML editing**:
- Click the "Edit" icon on any resource
- Modify the YAML directly in the browser
- Validation and syntax highlighting included
- See changes reflected immediately

**Scaling**:
- Go to a Deployment
- Click the "Scale" button
- Adjust replica count with a slider or input
- Watch pods spin up or down in real-time

#### Explore your workshop resources

Try these tasks in Headlamp:

1. **Scale your deployment**: Go to the `web-app` deployment and scale it to 5 replicas
2. **View logs**: Click on a pod and check the logs tab
3. **Inspect ConfigMaps**: Go to "Config & Storage" → "Config Maps" and view your `web-app-*` ConfigMaps
4. **Monitor resources**: Check the resource usage graphs for your pods
5. **View Ingress**: Navigate to "Network" → "Ingresses" to see your ingress rules

#### Advantages over kubectl for development

While kubectl is essential for scripting and automation, Headlamp excels at:

- **Visual exploration**: See relationships between resources at a glance
- **Quick debugging**: Jump from deployment → pods → logs in seconds
- **Resource discovery**: Browse available resources without memorizing commands
- **Live updates**: Watch changes happen in real-time without refreshing
- **Context switching**: Easy switching between clusters and namespaces

#### Headlamp best practices

1. **Use it alongside kubectl**: Headlamp is great for exploration, kubectl for automation
2. **Keep it updated**: Headlamp is actively developed with frequent improvements
3. **Multiple clusters**: Headlamp can manage multiple clusters - great for dev/staging/prod
4. **RBAC applies**: Headlamp respects your kubeconfig permissions
5. **Local only for sensitive clusters**: For production clusters, use appropriate security measures

#### Challenge Exercise

1. Use Headlamp to trigger a rolling update by changing the image of your deployment
2. Find the pod using the most memory and check its logs
3. View the events for the `workshop` namespace to see what's been happening
4. Open a terminal in a running pod and explore the filesystem
5. Try editing a ConfigMap and watch a pod restart to pick up changes

---

### 11. 📦 Helm - Package Management

**Time**: ~30 minutes

Helm is the package manager for Kubernetes. It allows you to define, install, and upgrade complex Kubernetes applications using reusable templates called Charts.

**Why Helm?**
- **Templating**: Reuse manifests with different values
- **Versioning**: Track and rollback application releases
- **Dependencies**: Charts can depend on other charts
- **Sharing**: Publish and use community charts

**📚 Learn more about Helm**:
- [Helm Documentation](https://helm.sh/docs/) - Official Helm documentation
- [Helm Chart Best Practices](https://helm.sh/docs/chart_best_practices/) - Creating better charts
- [Helm Chart Template Guide](https://helm.sh/docs/chart_template_guide/) - Templating reference
- [Artifact Hub](https://artifacthub.io/) - Discover and share Helm charts

#### Helm concepts

- **Chart**: A package containing Kubernetes manifest templates
- **Release**: An instance of a chart deployed to a cluster
- **Repository**: A collection of charts available for download
- **Values**: Configuration parameters for customizing charts

#### Explore the example chart

Our workshop includes a Helm chart in the `helm/example-app/` directory:

```bash
ls -la helm/example-app/
```

Chart structure:

```
example-app/
├── Chart.yaml          # Chart metadata (name, version, description)
├── values.yaml         # Default configuration values
├── templates/          # Kubernetes manifest templates
│   ├── deployment.yaml
│   ├── service.yaml
│   ├── ingress.yaml
│   └── _helpers.tpl   # Template helpers and functions
└── charts/            # Dependent charts (e.g., redis)
```

#### View the chart metadata

```bash
cat helm/example-app/Chart.yaml
```

This shows the chart name, version, description, and dependencies (like Redis).

#### View default values

```bash
cat helm/example-app/values.yaml
```

These are the default configuration values. They can be overridden during installation.

#### Install a Helm release

Install our chart as the Ruby application:

```bash
helm install example-app-ruby ./helm/example-app \
  --set image.repository=ghcr.io/relu/example-app-ruby \
  --wait
```

**What this does**:
- Creates a release named `example-app-ruby`
- Uses the chart in `./helm/example-app`
- Overrides the image repository to use the Ruby app
- `--wait` waits for all pods to be ready before returning

**Note**: Helm applies all templates in the chart, creating deployment, service, and potentially other resources.

#### List Helm releases

```bash
helm ls
```

**Expected output**: You should see one release named `example-app-ruby` with status "deployed".

**Note**: Helm releases are namespaced. Use `helm ls -A` to see releases in all namespaces.

#### View release details

```bash
helm status example-app-ruby
```

This shows the release status, namespace, and deployed resources.

#### Check the deployed resources

```bash
kubectl get all -l app.kubernetes.io/instance=example-app-ruby
```

Helm automatically adds labels to track resources belonging to a release.

#### Upgrade a release

Let's enable the ingress for our Ruby app:

```bash
helm upgrade example-app-ruby ./helm/example-app \
  --reuse-values \
  -f ./helm/values-ruby.yaml \
  --set ingress.enabled=true
```

**What this does**:
- Upgrades the existing `example-app-ruby` release
- `--reuse-values` keeps previously set values (like image.repository)
- `-f ./helm/values-ruby.yaml` loads additional values from a file
- `--set ingress.enabled=true` enables the ingress

Check the ingress was created:

```bash
kubectl get ingress -l app.kubernetes.io/instance=example-app-ruby
```

#### View release history

```bash
helm history example-app-ruby
```

You should see two revisions: the initial install and the upgrade.

#### Rollback a release

If an upgrade goes wrong, easily rollback:

```bash
helm rollback example-app-ruby 1
```

This reverts to revision 1 (the initial installation). The ingress will be removed.

Upgrade again to bring back the ingress:

```bash
helm upgrade example-app-ruby ./helm/example-app \
  --reuse-values \
  -f ./helm/values-ruby.yaml \
  --set ingress.enabled=true
```

#### Install releases for Python, Go, and Rust apps

Use the same chart to deploy the other applications:

```bash
helm install example-app-python ./helm/example-app \
  --set image.repository=ghcr.io/relu/example-app-python \
  -f ./helm/values-python.yaml \
  -f ./helm/overrides.yaml \
  --wait

helm install example-app-go ./helm/example-app \
  --set image.repository=ghcr.io/relu/example-app-go \
  -f ./helm/values-go.yaml \
  -f ./helm/overrides.yaml \
  --wait

helm install example-app-rust ./helm/example-app \
  --set image.repository=ghcr.io/relu/example-app-rust \
  -f ./helm/values-rust.yaml \
  -f ./helm/overrides.yaml \
  --wait
```

**Note**: We're using multiple `-f` flags to layer values files. Values are merged in order, with later files overriding earlier ones.

List all releases:

```bash
helm ls
```

You should see three releases, all using the same chart with different configurations.

#### Values precedence

When using Helm, values are applied in this order (later values override earlier ones):

1. Default values in `values.yaml`
2. Values from `-f` files (in order specified)
3. Values from `--set` flags
4. Values from `--set-file` flags
5. Values from `--set-string` flags

#### Understand Helm templating

Helm templates use Go templating syntax. Let's look at an example:

```bash
cat helm/example-app/templates/deployment.yaml
```

You'll see placeholders like:
- `{{ .Values.image.repository }}` - References values from values.yaml
- `{{ include "example-app.fullname" . }}` - Calls a helper function
- `{{ .Release.Name }}` - Built-in variable with the release name

#### Render templates without installing

Preview what Helm will deploy:

```bash
helm template example-app-test ./helm/example-app \
  --set image.repository=ghcr.io/relu/example-app-ruby
```

This outputs the rendered Kubernetes manifests without installing anything. Useful for debugging templates.

#### Chart dependencies

Our chart has a dependency on Redis (see `Chart.yaml`). To manage dependencies:

```bash
# Update dependencies (downloads dependent charts)
helm dependency update ./helm/example-app

# List dependencies
helm dependency list ./helm/example-app
```

#### Helm repositories

Helm can install charts from remote repositories. We used this earlier for Traefik and Metrics Server.

List configured repositories:

```bash
helm repo list
```

**Major Helm repository change**: The stable and incubator chart repositories were deprecated. Use [Artifact Hub](https://artifacthub.io) to discover charts.

Add a repository:

```bash
helm repo add bitnami https://charts.bitnami.com/bitnami
helm repo update
```

Search for charts:

```bash
helm search repo redis
```

Install a chart from a repository:

```bash
helm install my-redis bitnami/redis --namespace workshop
```

#### Uninstall a release

Remove a Helm release and all its resources:

```bash
helm uninstall example-app-ruby
```

Verify it's gone:

```bash
helm ls
kubectl get all -l app.kubernetes.io/instance=example-app-ruby
```

#### Helm best practices

1. **Use version pinning** in Chart.yaml dependencies
2. **Validate charts** with `helm lint`
3. **Test charts** with `helm test`
4. **Document values** in values.yaml with comments
5. **Use semantic versioning** for chart versions
6. **Keep secrets out of values files** (use external secret management)
7. **Review rendered templates** before deploying (`helm template`)

#### Challenge Exercise

1. Create a custom values file that changes the replica count and resource limits
2. Use `helm template` to preview the changes before applying
3. Install a chart from a public repository (e.g., PostgreSQL, MongoDB)
4. Try adding a new template file to the example-app chart

---

## 🔧 Troubleshooting

Common issues and their solutions:

### Container runtime not running

**Symptoms**: `kind create cluster` fails with errors like "Cannot connect to the Docker daemon" or "docker: command not found"

**Solution**:

**On Mac (Colima)**:
```bash
# Check if Colima is running
colima status

# If not running, start it
colima start
```

**On Windows/Mac (Docker Desktop)**:
- Make sure Docker Desktop is running (check system tray for the whale icon)
- If Docker Desktop shows an error, try restarting it
- On Windows, ensure WSL 2 is running: `wsl --status`

**On Linux**:
```bash
# Check if Docker is running
sudo systemctl status docker

# Start Docker if needed
sudo systemctl start docker
```

**Note**: If you restart your computer, you'll need to start Colima or Docker Desktop again.

### Pods not starting

**Symptoms**: Pod status shows `Pending`, `ImagePullBackOff`, `CrashLoopBackOff`

**Debug steps**:
```bash
# Check pod status
kubectl get po

# Describe the pod for events
kubectl describe po <pod-name>

# Check logs
kubectl logs <pod-name>

# Check previous container logs if it crashed
kubectl logs <pod-name> --previous
```

**Common causes**:
- **Pending**: Insufficient cluster resources, check `kubectl describe po`
- **ImagePullBackOff**: Image doesn't exist or is private, verify image name
- **CrashLoopBackOff**: Application is crashing on startup, check logs

### Service not reachable

**Symptoms**: Can't connect to a service

**Debug steps**:
```bash
# Check if service exists
kubectl get svc

# Check if service has endpoints
kubectl get endpointslices -l kubernetes.io/service-name=<service-name>

# Describe the service
kubectl describe svc <service-name>

# Verify pod labels match service selector
kubectl get po --show-labels
```

**Common causes**:
- No endpoints: Label selector doesn't match any pods
- Wrong port: Check service port vs target port vs container port

### Ingress not working

**Symptoms**: Can't access application through ingress

**Debug steps**:
```bash
# Check ingress exists
kubectl get ingress

# Describe ingress
kubectl describe ingress <ingress-name>

# Check ingress controller is running
kubectl get po -n kube-system -l app.kubernetes.io/name=traefik

# Check ingress controller logs
kubectl logs -n kube-system -l app.kubernetes.io/name=traefik
```

**Common causes**:
- Ingress controller not installed or not ready
- Service name or port mismatch in ingress spec
- Path or host rules not matching your request

### Metrics Server not working

**Symptoms**: `kubectl top` returns errors

**Debug steps**:
```bash
# Check metrics-server is running
kubectl get po -n kube-system -l app.kubernetes.io/name=metrics-server

# Check logs
kubectl logs -n kube-system -l app.kubernetes.io/name=metrics-server
```

**Common cause**: Missing `--kubelet-insecure-tls` flag for Kind clusters

**Fix**: Patch the deployment:
```bash
kubectl patch deployment metrics-server -n kube-system --type='json' \
  -p='[{"op": "add", "path": "/spec/template/spec/containers/0/args/-", "value": "--kubelet-insecure-tls"}]'
```

### Dashboard access issues

**Symptoms**: Can't access dashboard or authentication fails

**Common causes**:
- Token expired (valid for 1 hour)
- kubectl proxy not running
- Wrong URL

**Fix**: Generate a new token:
```bash
kubectl -n kubernetes-dashboard create token admin-user
```

### Helm installation fails

**Symptoms**: `helm install` returns errors

**Debug steps**:
```bash
# Validate chart
helm lint ./helm/example-app

# Render templates to see what will be deployed
helm template test ./helm/example-app

# Check helm release status
helm status <release-name>

# View helm release history
helm history <release-name>
```

### General debugging techniques

**Get cluster events**:
```bash
kubectl get events --sort-by='.lastTimestamp'
```

**Check resource quotas**:
```bash
kubectl describe quota
```

**Interactive debugging**:
```bash
# Run a debug container in the cluster
kubectl run debug --image=busybox -it --rm --restart=Never -- sh
```

**Check DNS**:
```bash
kubectl run -it --rm debug --image=busybox --restart=Never -- nslookup kubernetes.default
```

### Still stuck?

1. Check the [Kubernetes documentation](https://kubernetes.io/docs/)
2. Search for error messages in GitHub issues
3. Ask in the Kubernetes Slack community
4. Review pod events: `kubectl describe pod <pod-name>`

---

## 🧹 Cleanup

When you're finished with the workshop, clean up all resources:

### Option 1: Delete the namespace (quick)

This removes all resources in the workshop namespace:

```bash
kubectl delete namespace workshop
```

### Option 2: Delete the entire cluster (complete cleanup)

```bash
kind delete cluster --name workshop
```

This removes the entire local Kubernetes cluster and all its data.

**Note**: If you delete the cluster, you'll need to recreate it from the beginning if you want to repeat the workshop.

### Verify cleanup

```bash
# Check clusters
kind get clusters

# Check contexts (the workshop context should be gone)
kubectl config get-contexts
```

### Stop container runtime (Optional)

If you want to free up system resources after the workshop:

**On Mac (Colima)**:
```bash
colima stop
```

To start it again later: `colima start`

**On Windows/Mac (Docker Desktop)**:
- Right-click the Docker Desktop icon in the system tray
- Select "Quit Docker Desktop"

**On Linux (Docker)**:
```bash
sudo systemctl stop docker
```

**Note**: You'll need to start your container runtime again before running Kind in the future.

Au revoir!

---

## 📖 Additional Resources

**Note**: Portions of this workshop content were enhanced and reviewed with assistance from Claude (Anthropic).

### Official Documentation

- [Kubernetes Documentation](https://kubernetes.io/docs/home/)
- [kubectl Cheat Sheet](https://kubernetes.io/docs/reference/kubectl/cheatsheet/)
- [Helm Documentation](https://helm.sh/docs/)
- [Kind Documentation](https://kind.sigs.k8s.io/)

### Learning Resources

- [Kubernetes Basics Tutorial](https://kubernetes.io/docs/tutorials/kubernetes-basics/)
- [Kubernetes the Hard Way](https://github.com/kelseyhightower/kubernetes-the-hard-way)
- [Helm Charts Best Practices](https://helm.sh/docs/chart_best_practices/)

### Tools

- [k9s](https://k9scli.io/) - Terminal UI for Kubernetes
- [Headlamp](https://headlamp.dev/) - Modern Kubernetes UI
- [kubectx/kubens](https://github.com/ahmetb/kubectx) - Fast context/namespace switching
- [stern](https://github.com/stern/stern) - Multi-pod log tailing

### Community

- [Kubernetes Slack](https://slack.k8s.io/)
- [Helm Slack](https://slack.helm.sh/)
- [Artifact Hub](https://artifacthub.io/) - Discover Helm charts

### Practice

- [Play with Kubernetes](https://labs.play-with-k8s.com/) - Free browser-based playground
- [Killercoda Kubernetes Scenarios](https://killercoda.com/kubernetes) - Interactive tutorials

---

**Congratulations!** You've completed the Kubernetes Introduction Workshop. You now have hands-on experience with core Kubernetes concepts and are ready to explore more advanced topics like StatefulSets, DaemonSets, Jobs, Network Policies, and production deployment strategies.

Keep practicing, and happy clustering!
