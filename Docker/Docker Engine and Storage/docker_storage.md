# Docker Storage
When docker is installed, a directory is created in /var/lib/docker which contains all its data by default.

Docker images are created in layered architecture, meaning when you create a Dockerfile, each command in that dockerfile is a layer and each layer will be cached. If you create a new dockerfile with some of the same commands and arguments, it will bring these from the cache rather than building new layers. This saves disk space and results in images being built faster and more efficiently. Can also be used when updating application code. These are the image layers and will be read-only.

When we create a container using the new image we have built, a new layer is created called the container layer which will have read/write permissions. This layer is used to store data created by the container such as log files created by the application, temp files created by the container, or any file modified by the user on that container. This layer is only alive whilst the container is alive.

If any changes are made to a file located in the image layer using the container layer, a copy of that file is created in the container layer and this is what is edited, not the file in the image layer. This is known as copy-on-write. The image layer and any files in the image layer will remain the same unless rebuilt. When the container is destroyed, the copy of the file will also be destroyed. If we wish to persist the file, we can add a persistant volume to the container. We can do this by running the below command;
`docker volume create <volume_name>`
A new volume directory will be created in the /var/lib/docker/volumes directory with the same volume name. You would then use the -v flag and specify the volume as well as where it is to be mounted when running the docker image. E.g.
- `docker volume create data_volume`
- `docker run -v data_volume:/var/lib/mysql mysql`
The first step is not required as if you run the second command, docker will automatically create a new volume and mount it to the specified location. Mounting a volume using the /var/lib/docker/volumes directory is known as a volume mount.

You can also mount volumes which are in separate locations to the /var/lib/docker/volumes directory by adding the absolute path of the chosen directory. This is known as bind mount. For example, we can do:
`docker run -v /data/mysql:/var/lib/mysql mysql`

The newer method of mounting volumes is using the --mount flag as can be seen below;
`docker run --mount type=bind,source=/data/mysql,target=/var/lib/mysql mysql`

The storage drivers are responsible for maintaining the layered architecture, maintaining a writable layer, moving files across layers to enable copy-and-write, etc. Some storage drivers include:
- AUFS
- ZFS
- BTRFS
- Device Mapper
- Overlay
- Overlay2

The type of storage driver depends on the underlying OS.