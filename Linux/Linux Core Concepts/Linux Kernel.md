# Linux Kernel
The kernel is the core interface between hardware and its processes. It communicates between the two managing resources as efficiently as possible. In simple terms, it allocates the correct amount of hardware resources to the processes and then removes these resources when the processes are back so other processes can utilise them.

The kernel is responsible for the following tasks:
- Memory management: Keeping track on how much memory is store what and where.
- Process management: Determines which process can use the CPU and for how long.
- Device drivers: Acts as a mediator between the hardware and processes
- System calls and security: Receive requests for service from the processes.

The Linux kernel is monolithic, meaning the kernel carries out all of the above by itself. The kernel is also modular which means it can increase its capabilities using dynamically loaded kernel modules.

## Kernel Versions
- `uname`: Displays information about the kernel.
- `uname -r`: Prints kernel version. For version 4.15.0.72-generic, 4 is the kernel version, 15 is the major version, 0 is the minor version, 72 i the patch release and Generic is the distro specific information.

## Kernel and User Space
Memory is divided into two areas, the kernel space and the user space. Kernel space is the portion of memory in which the kernel executes and provides its services. A process running in the kernel space has unrestricted access to hardware resources and therefore it is strictly reserved for tasks such as running the kernel code, kernel extensions and most device drivers.

All processed running outside the kernel reside in the user space which has restricted access to CPU and memory. Most UNIX-like operating systems, including Linux come pre-packaged with different utilities, graphical tools and programming languages. These are called user space applications. Examples include programmes written in Python or Java.

All user programmes function by manipulating data which is stored in memory and on disk. These programmes get access to data by making system calls to the kernel. These include asking for memory or opening a file. 