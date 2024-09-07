# Disk Partition

## Block Device
Block device is a type of file which can be located under the /dev directory. It usually represents a piece of hardware that can store data. It is named block storage as data is read and written in blocks of space. In order to see the block devices in your system, use the below command;
`lsblk`

The example in the course, sda is the whole block storage device and sda1, sda2 and sda3 are the disk partitions.

Each block device has a major and a minor number. The major number is used to identify the type of block device and the minor number is used to identify the device number. 

## Partitions
Partition allow us to segment space and use these partitions for specific purposes. The `lsblk` command also displays the mountpoint which can be used to distinguish what the partition is for. Partitioning is not necessary and the block device can be used as is but it is always recommended due to flexibility. You can use the below commands to inspect the partition tables;
- `lsblk`
- `fdisk -l /dev/sda` 

## Partition Types
- Primary Partition: A type of partition which can be used to boot an operating system. Usually limited to 4 primary partitions per disk.
- Extended Partition: A type of partition that cannot be used on its own but can host logical partitions. As we are limited to 4 primary partitions, we can create extended partitions and carve out logical partitions inside it. An extended partition is similar to a whole disk in the sense that it has a partition table which points to one or more logical partition.
- Logical Partition: Created in an extended partition. 

## Partition Scheme
How a disk is partitioned is defined by a partitioning scheme (partition table). There can only be 4 primary partitions in the Master Boot Record (MBR). The maximum size per disk is two terabytes. In order to create more than 4 partition per disk, the 4th partition would need to be created as an extended partition with logical partitions within it. GPT (GUID Partition Table) is a more recent type of partitioning table and was created to address the limitations of MBR. GPT can have many more partitions per disk (theoretically unlimited). This is usually limited by restrictions imposed by the operating system. 

## Creating Partitions
- `gdisk /dev/sdb`: Allows you to configure partitions on the specified disk. You would enter into a menu and to create a new partition, you would choose the 'n' option. Once you have finished creating the new partition, click 'w' to save the configurations. 