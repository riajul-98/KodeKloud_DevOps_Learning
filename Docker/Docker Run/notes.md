# Docker Run commands
- `docker run redis:4.0`: 4.0 is a tag and is used to tell docker which version of the image to pull and run. In this case, it would be version 4.0. If no tag is provided, Docker will pull and run 'latest'.
- `docker run -it <image_name>`: `-i` runs the docker container in interactive mode, allowing users to input data. `-t` stands for sudo terminal. The combination of `-it` is useful when you would like to prompt your user with a question.
- `docker run -p 80:5000 <image_name>`: Maps host port (80) to container port (5000). 
- `docker run -v /opt/datadir:/var/lib/mysql mysql`: Persists data from the data file location on the container (/var/lib/mysql) to a directory on the host (/opt/datadir).
- `docker inspect <container_name/container_id>`: Provides details on a container. Includes state, mounts, network settings etc
- `docker logs <container_name/container_id>`: Provides logs on the specified container. Good for containers run in the background.