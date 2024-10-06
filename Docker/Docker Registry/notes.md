# Docker Registry
Docker images are stored in registries such as DockerHub. When we pull a Docker image, Docker will look at the user/account name, followed by the image/repository name. If no user/account name is provided, it will look at 'library' e.g. `docker run nginx` will look for the nginx image under library/nginx. DockerHub is the default location which Docker looks for an image from. These are public registries.

## Priavte Registry
Should you wish to upload private images, a private registry would be more suitable such as ECR or GCR. To access these images, you would first need to login to your private registry using the `docker login <registry_name>`.

To deploy a private registry, you would run the Docker registry image on port 5000. To push your image to this registry, you would use the `docker image tag <image_name> localhost:5000/<image_name>` command, followed by `docker push localhost:5000/<image_name>`.