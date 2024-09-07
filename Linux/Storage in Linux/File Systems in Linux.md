# File Systems in Linux
Partitioning alone does not make a disk usable in the Linux OS. The disks and partitions are seen as raw disks by the Linux kernel. To write to a disk or partition, we must first create a file system. The file system defines how data is stored on a disk. After creating the file system, we need to mount it to a directory and then we can read or write data to it.

## Linux Filesystems
- EXT2:
    - 2TB maximum file size
    - Allows 4TB volume size
    - Supports compression
    - Supports Linux permissions
    - Long crash recovery
- EXT3:
    - 2TB maximum file size
    - Allows 4TB volume size
    - Uses journal
    - Backwards compatible
    - Advantage over EXT2: Does not have long crash recovery (allows quicker startup time after a crash)
- EXT4:
    - 16TB maximum file size
    - 1 Exabyte
    - Uses journal
    - Uses chksum for journal
    - Backwards compatible

## Working with EXT4
- `mkfs.ext4 /dev/sdb1`: Creates an EXT4 file system using the specified disk.
- `mount /dev/sdb1 /mnt/ext4`: Mounts the disk in the specified mount point.

To make sure the mount is available after the system reboots, you would need to add an entry to the /etc/fstab file. See below for an example;

```

<file system>   <mount point>   <type>  <option>     <dump>      <pass>
/dev/sda1       /               ext4    RW              0           0 

```
The value meanings can be seen in the table below;

| FIELD       | Purpose                                             |
|-------------|-----------------------------------------------------|
| Filesystem  | Such as /dev/vdb1 to be mounted                     |
| Mountpoint  | Directory to be mounted on                          |
| Type        | Example ext2, ext3, ext4                            |
| Options     | Such as RW = Read-write, RO = Read Only             |
| Dump        | 0 = Ignore, 1 = take backup                         |
| Pass        | 0 = ignore, 1 or 2 = FSCK filesystem check enforced |


You can check file systems created using `df -h` command.