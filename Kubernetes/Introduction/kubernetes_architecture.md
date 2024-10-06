# Kubernetes Architecture

## Nodes
A node is a machine, physical or virtual which kubernetes is installed on. A node is a worker machine and this is where containers run. We need more than one node to ensure availability if one of the nodes fails. Multiple nodes grouped together is known as a cluster. Having multiple nodes also allows for sharing load. 

## Master
The master node monitors and manages the worker nodes. The master is responsible for the orchestration of containers on the worker nodes. 

## Components
When you install Kubernetes on a system, you are also installing the following;
- API server: Acts as the frontend for Kubernetes. Users, management devices and CLIs all talk to API server to communicate with the kubernetes cluster. (master)
- Etcd: A distributed, reliable key value store used by Kubernetes to store all data used to manage the cluster. Responsible for implementing logs within the cluster to ensure there are no conflicts within the masters. (master)
- Kubelet: Agent that runs on each node in the cluster. Responsible for making sure containers are running on the nodes as expected. (worker)
- Scheduler: Responsible for distributing work or containers across multiple nodes. (master)
- Container runtime: Underlying software used to run containers. (worker)
- Controller: Responsible for monitoring to see if nodes, containers or endpoints go down and respond to these issues. The controllers make decisions to bring up new containers in such cases. (master)

## Kubectl commands
- `kubectl run hello-minikube`: Runs the specified application on the cluster
- `kubectl cluster-info`: Get information of the cluster
- `kubectl get nodes`: Displays nodes in the cluster