# DAS, NAS and SAN
It is not feasible to use internal storage for enterprice server environments such as a production database or a web server storing a lot of data. This is where external storage comes in. These external storage devices offer high availability and high capacity. The course discusses about 3 different types of storage. These are:
- Direct Attached Storage (DAS)
- Network Attached Storage (NAS)
- Storage Area Network (SAN)

## DAS
With DAS, external storage is attached directly to the host system that requires the storage. The host OS sees a connected DAS device as a block device. As this is attached directly, there is no network or firewall between the storage and the host. This results in excellent performance at an affordable cost. DAS usually has a faster response than a NAS device. The negative in this setup is as the device is attached directly, it can only be dedicated to a single server. This is not ideal for enterprise servers.

## NAS
NAS provides storage over the network, meaning the device is separate from the hosts which need storage. The storage is provided to the host in the form of a directory which is present in the NAS device but exported to the host device over the network. This type of storage is ideal for data which simultaneously needs to be accessed by multiple different host machines. 

## SAN
SAN provides block storage used by enterprises for business critical applications that need to deliver high throughput and low latency. Storage is allocated to hosts in the form of a LUN (Logical Unit Number). A LUN is a range of blocks provisioned from a pool of shared storage and presented to the server as a logical disk. The host system will detect the storage as a raw disk. We can then create partitions and file systems on top of it and mount it on the system to store data. SAN can also be ethernet based but it mainly makes use of the FCP (Fiber Chanel Protocol).