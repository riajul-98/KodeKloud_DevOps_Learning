# Docker Image
To build a docker image using dockerfile, run the `docker build Dockerfile -t <docker_image_name:tags>` command.

To publish the docker image on DockerHub, run the `docker push <docker_image_name:tags>` command. You would first need to login to your DockerHub account using the `docker login` command.

You can see the size of the created docker image by typing `docker history <docker_image_name:tags>`.

## Environment Variables
To set an environment variable, you would run the docker container with the `-e` flag, followed by the variable_name=variable_value. e.g. `docker run -e APP_COLOR=blue simple-webapp-color`.

To check the environment variables set on a running container, you can use the `docker inspect <container_name/container_id>` command, under the config section.

## Entrypoint vs CMD
- CMD: Specifies the default command to run when a container starts. It can be overridden by the user when they run the container. For example: `CMD ["sleep", "5"]` or `CMD sleep 5` will run if no command line arguments were given but will be overridden if another command line argument is provided at the time of docker run command such as `docker run Ubuntu sleep 10`.

- Entrypoint: Defines the main command that always runs when the container starts. It is more rigid than CMD and typically used when you want to enforce a particular command to always run. For example, `ENTRYPOINT ["sleep"]` would have the command sleep and you would just add a parameter to tell Docker how long to sleep for. e.g. `docker run Ubuntu 10`. The problem with this is, an error will be given if no value is provided as a parameter. Hence, it is better to use ENTRYPOINT + CMD.

- CMD and Entrypoint: When used together, CMD is treated as additional arguments for ENTRYPOINT. The combination allows you to specify default arguments that can be overridden

To overwrite the Entrypoint, you would add the --entrypoint flag at the time of running the container. E.g. `docker run --entrypoint sleep 2.0 Ubuntu 10`