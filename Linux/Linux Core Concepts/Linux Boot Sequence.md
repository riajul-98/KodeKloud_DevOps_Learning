# Linux Boot Sequence
The boot process is broken down into 4 stages;
- BIOS POST
- Boot loader (GRUB2)
- Kernel Initialisation
- INIT Process (systemd)

## BIOS POST
Very little to do with Linux itself. POST stands for Power-On Self-Test. BIOS runs a POST test to make sure hardware components attached are functioning correctly. If faile, the computer may not be able to operate and the system will not proceed to the boot loader stage.

## Boot Loader
After a successful POST test, the BIOS loads and executes the boot code from the boot device which is located in the first sector of the hard disk. In Linux, this is located in /boot. This is where we see the boot screen which gives us different options to boot into if multiple OS's are available. Once selection is made, the boot loader loads the kernel into memory, provides some parameters and allows the kernel to take control. GRUB2 is a popular example and is used for almost all Linu distributions. 

## Kernel Initialisation
The kernel is then usually decompressed to save space. It is then loaded into memory and starts executing. During this phase, the kernel starts carrying out tasks such as initialising hardware and memory management tasks. The kernel then looks for an init process to run which is used to set up the user space and the processes needed for the user environment.

## INIT Process (systemd)
The INIT function usually calls the systemd daemon. Systemd is responsible for bringing the Linux host to a usable state. Systemd is responsible for starting and managing system services. 

There are other INIT systems and you can check the init system used by running the `ls -l /sbin/init` command.