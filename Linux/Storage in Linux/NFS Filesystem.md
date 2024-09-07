# NFS Filesystem
NFS does not store data as blocks. It saves data in the form of files. NFS works on a server-client model. The data on your laptop does not phtsically reside on any of the attached disks, rather, it is delivered over the network but once mounted, it can be used like any other file system in the operating system. Directory sharing in NFS is known as exporting. The NFS server maintains an export configuration file at /etc/exports that defines the clients which should be able to access the directories in the server. 

Ideally there should be a network firewall between the NFS server and the clients. As a result, specific ports need to be open between the server and the clients for the NFS solution to work. Once the exports configuration file has been updated, the directory is shared to the clients by using the exportfs command. 

- `exportfs -a`: Exports all the mounts defined in /etc/exports.
- `exportsfs -o 10.61.35.201:/software/repos`: Manually exports the specified directory to the specified IP address. 

You can then mount it on a local directory using the below command;
- `mount 10.61.112.101:/software/repos /mnt/software/repos`: Mounts the directory from the NFS server (10.61.112.101) to the chosen local directory.