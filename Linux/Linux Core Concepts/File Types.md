# File Types
Everything is a file in Linux. There are three types of files;

- Regular File: These include files which contain text, data, configuration and images.
- Directory: A type of file which stores other files and directories.
- Special Files: Special files can be subdivided into 5 other file types:
    - Character files: Represent devices under the /dev file system that allows the OS to talk to IO devices such as keyboard and mouse.
    - Block files: Files representing block devices which are located under /dev. A block device reads from and writes to the device in blocks or chunks of data.
    - Links: Allow 2 or more files to be associated to the same set of file data. There are two types, soft links (similar to a shortcut in Windows, deleting one does not delete the other) and hard links (same block of data on the physical disk, deleting one, deletes the other)
    - Socket files: A file which enables communication between two processes.
    - Named Pipes: Special type of file that allows connecting one process as an input to another.

You can check the file types in the following ways:

- `file <file name>`: Displays the file type.
- `ls -ld <file_name>`: Long lists the files and provides file type as well as permissions.