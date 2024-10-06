# Docker Engine
Docker Engine is the host on which Docker is installed. When Docker is installed on Linux, you are actually installing 3 different components;
- The Docker Daemon: A background process which manages Docker objects such as images, networks and containers
- The Docker REST API: The API interface that programs can use to talk to the daemon to create instructions
- The Docker CLI: The command line interface to perform actions such as running containers, stopping containers etc. Uses the REST API to talk to the Daemon. Does not need to be on the same host. If you want to run a command on a separate host, you would use the below command;
`docker -H=<remote_docker_address:port> <command>`
e.g.
`docker -H=10.123.2.1:2375 run nginx`

Docker uses namespaces to isolate workspaces. Process ID, network, InterProcesses, mounts and Unix Timesharing systems are all created in their own namespace, providing isolation between containers.

Cgroups (control groups) are used to restrict the amount of hardware allocated to each container. This is needed as if a container uses too much resources, it might take up all the resources on the host OS. You can use this by running the command below;
`docker run --cpus=0.5 ubuntu` or `docker run --memory=100m ubuntu`