# Working Through The Linux CLI
According to research conducted by StackOverflow, Linux is one of the most commonly used and most demanded platform used in the DevOps field. Most tools used in the DevOps fields such as Docker and Ansible are first developed in Linux and then rolled out to other platforms such as Windows. Certifications such as the CKA are also done on the Linux system. This course will mainly be covering CentOS but may also touch on Ubuntu.

## Shell Types
Linux systems have both CLI and GUI interfaces. Most IT roles only make use of the CLI. The Linux CLI is also known as the Linux shell. There are different types of shells. These include:
- Bourne Shell (Sh Shell)
- C Shell (csh or tcsh)
- Z Shell (Z Shell)
- Bourne Again Shell (bash)

You can see the shell you are using by using the command
`echo $SHELL`

## Basic Commands
- `echo Hi`: prints the text that follows echo. In this case "Hi". Can be used to print environment variables.
- `ls`: Lists files and directories.
- `cd my_dir1`: changes the directory to the one specified after cv.
- `pwd`: Displays the present working directory.
- `mkdir new_directory`: Creates a new directory with the name specified.
- `cd new_directory; mkdir www; pwd`: The semicolon separating the commands allows you to run multiple commands, one after the other.

## Commands - Directories
- `mkdir -p /tmp/asia/india/bangalore`: Allows you to create directories within directories even if the previous directories are not present.
- `rm -r /tmp/my_dir`: Removes the specified directory
- `cp -r mydir1 /tmp/mydir1`: Allows you to copy mydir1 and its contents to /tmp/mydir1.

## Commands - Files
- `touch new_file.txt`: Creates an empty file with the specified name.
- `cat > newfile.txt`: Allows you to add content to the specified file.
- `cat new_file.txt`: Displays the contents of the specified file.
- `cp new_file.txt copy_file.txt`: Copies file (can be to a new directory or the same directory under a different name).
- `mv new_file.txt sample_file.txt`: Moves file to a new location or renames file if in the same location.
- `rm new_file.txt`: Removes specified file.