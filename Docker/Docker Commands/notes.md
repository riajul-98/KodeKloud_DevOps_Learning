# Commands
- `docker run <docker_image>`: Runs an instance of the specified docker image. If image is not present, docker will pull from DockerHub.
- `docker ps`: Lists all running containers and basic information such as container ID, image, port and name of the container.
- `docker ps -a`: Lists all running and exited containers as well as basic information such as container ID, image and name of the container.
- `docker stop <container_id/container_name>`: Stops the running container specified.
- `docker images`: Lists all docker images on system as well as their sizes.
- `docker rm <container_id/container_name>`: Removes container.
- `docker rmi <image_name>`: Removes Docker image.
- `docker pull <docker_image>`: Pulls docker image but does not run it.
- `docker run ubuntu sleep 5`: Runs an Ubuntu container and instructs it to run the sleep command for 5 seconds.
- `docker exec <container_name> <command>`: Allows you to run commands on a running container. E.g. 
`docker exec distracted_mcclintock cat /etc/hosts`
- `docker run -d <docker_image>`: Runs container in detached mode (background).
- `docker attach <container_name/container_id>`: Brings container to the foreground.