# Logical Volume Metric (LVM)
LVM allows grouping of multiple physical volomes which are hard disks or partitions into a volume group. From this volume group, you can carve out logical columes. LVM allow the logical volumes to be resized dynamically as long as there is sufficient space in the volume group. On Linux systems, you can have file systems such as /home/var/tmp created on top of LVM managed volumes. This would allow them to be easily resized when needed. 

## Using LVM
- `sudo apt-get install lvm2`: Installs LVM package.
- `pvcreate /dev/sdb`: Creates a physical volume for the specified disk or partition (the disk must be free). The physical volume is how LVM identifies a disk or partition. 
- `vgcreate caleston_vg /dev/sdb`: Creates a volume group called caleston_vg using the specified volume or partition.
- `pvdisplay`: Lists all the physical volumes as well as their names, sizes and the volume group they are a part of.
- `vgdisplay`: Provides more details on the volume group including the name, the total size and the amount of disk space used.
- `lvcreate -L 1G -n vol1 caleston_vg`: Creates 1GB logical volume called vol1 in the caleston_vg group. -L stands for linear volume which is a common type of logical volume. This allows us to make use of multiple physical volumes if available.
- `lvdisplay`: Displays the logical volumes.
- `lvs`: Lists the logical volumes including their names, volume group and size.
- `mkfs.ext4 /dev/caleston_vg/vol1`: Creates a file system on the logical volume.
- `mount -t ext4 /dev/caleston_vg/vol1 /mnt/vol1`: Mounts the file system onto the chosen directory.
- `vgs`: Lists the VGs and their details.
- `lvresize -L +1G -n /dev/caleston_vg/vol1`: Adds 1GB storage to the logical volume at the given path. Only resizes the logical volume and not the file system.
- `resize2fs /dev/caleston_vg/vol1`: Resizes the file system to match the size of the LV.