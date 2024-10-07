# Pods

## Assumptions
We would like to assume the application source code has been developed and built on a Docker image and it is available on a Docker repository. We also assume the Kubernetes Cluster has been set up. All the services need to be in a running state. 

## Pod
Containers are not deployed directly on the worker nodes but instead on a Pod. A pod is a single instance of an application. Pods cannot house more than one instance of an application, hence if we want more, we spin up more pods containing the application on the worker node. To expand the clusters physical capacity, we increase the number of worker nodes. 

## Multi-container nodes
A Pod can have multiple containers but they are usually not of the same kind. We can have containers such as helper containers which assist the application container in tasks such as processing user data, processing files etc. When the application dies, so does the helper container as they are part of the same pod. The two containers can communicate to each other directly by referring to each other as localhost since they share the same network space and storage space. 

## How to deploy pods
- `kubectl run nginx --image=nginx`: Deploys specified container on pod.
- `kubectl get pods`: Displays running pods
- `kubectl describe pod nginx`: Displays information on the specified pod.
- `kubectl get pods -o wide`: Provides additional information on running pods.