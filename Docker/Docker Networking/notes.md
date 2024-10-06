# Docker Networking
When Docker is installed, 3 networks are created automatically. The bridge network, none and host network. The Bridge network is the default network a container gets attached to. In order to use one of the others, you would use the network command line parameter like below;
`docker run Ubuntu --network=none` or `docker run Ubuntu --network=host`

## Bridge Network
A private internal network created by Docker on the host. All containers attach to this network by default and they get an internal IP address, usually in the 172.17 series range. The containers can communicate with each other using this internal IP if required. In order to access these containers from outside, map the ports of these containers to ports on the docker host. 

## Host Network
Another way to access to containers externally is to associate the container with the host network. This takes out any network isolation between the Docker host and the Docker container. If you run a web application in a container on port 5000, it will be accessible externally on the same port of the Docker host. No port mapping required. The issue with this is you will not be able to run multiple containers on the same host, using the same port. 

## None Network
Containers are not attched to any network and does not have any access to the external network or other containers. They are isolated.

## User-defined networks
By default, Docker only creates one internal bridge network. If we wanted to isolate 2 containers in one network and another 2 in another network, we can use the below command to create a new network:
`docker network create --driver <network_type> --subnet <subnet> <custom-network_name>`
e.g.
`docker network create --driver bridge --subnet 182.18.0.0/16 custom-isolated-network`

To view all networks, run the `docker network ls` command.

## Embedded DNS
Containers can reach each other using their names. The DNS server is always running on 127.0.0.11. Docker uses network namespaces that creates different namespaces for each container. It then uses virtual ethernet pairs to connect containers together. 