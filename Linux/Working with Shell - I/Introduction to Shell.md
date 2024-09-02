# Introduction to the Shell
The Linux shell allows for a text based interaction between the user and the operating system. When you first login to the shell, you will be in your home directory.

## The Home Directory
The home directory is a directory which contains the home directory of all users. The home directory will be located at /home and the home directory for a user with the name Michael will be found at /home/michael. The home directory allows users to store their files and directories in the system. Other users cannot access files in your directory. Only you can. The users home directory is represented by "~" on the command line. 

## Commands and Arguments
- Commands: `echo` - echo prints out the argument which follows it. This will not provide an output as no argument was given.
- Arguments: `echo Hello` - echo prints out the argument which follows it, in this case "Hello".
- Commands which do not need arguments: `uptime` - Prints information on how long the system has been running for since the last reboot as well as other information.
- Options: `echo -n Hello` - Does the same as echo but the '-n' option removes the new line at the end.

## Command Types
- Internal/Built-in commands: Built into the shell e.g. cd, mkdir, pwd.
- External commands: Binary programs or scripts which either come pre-installed with the distributions manager or can be created or installed by the user.

You can check the command type by using the `type` command followed by the command you wish to check as an argument. e.g. `type echo`. 