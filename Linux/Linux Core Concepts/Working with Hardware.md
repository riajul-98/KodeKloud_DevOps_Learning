# Working with Hardware
Suppose you have a USB stick and you attach it to your device, there will be a device driver present in the kernel space which detects the device and sends this event as a uevent to the udev which stands for user space device manager (found in the user space). udev then creates a dynamically creating a device node associated with the newly attached USB stick which will be located in the /dev directory. You will be able to see these devices in the /dev directory.

## Commands
- `dmsg`: Displays messages from the ring buffer which includes logs on any devices.
- `udevadm info --query=path --name=/dev/sda5`: Queries udev database for device information.
- `udevadm monitor`: Listens to kernel uevents and prints details when detected.
- `lspci`: Lists information about all PCI devices configured on the system. PCI devices are devices which connect to the motherboard using the PCI slots.
- `lsblk`: Lists information about block devices. Type disk refers to the whole disk while type part refers to partition. These also give major and minor numbers, the major number identifies the type of device driver associated with the device. Minor number is used to differentiate similar devices
- `lscpu`: Displays information on the CPU architectures.
- `lsmem`: Lists available memory in the system. Using the `lsmem --summary` provides a summary.
- `free -m`: Displays total and used memory (-m displays this in MB, -g in GB and -k in KB).
- `lshw`: Displays detailed information on the entire hardware configuration of the machine.