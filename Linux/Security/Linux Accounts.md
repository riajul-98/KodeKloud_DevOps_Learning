# Linux Accounts
Every user in Linux has an associated account. The user account maintains information such as username and password as well as UID. The information about a user account is stored in the /etc/passwd file. 

A Linux group is a collection of users. It is used to organise users based on common attributes such as role or function. Information on groups can be seen in the /etc/group file. Each group has a group ID (GID). Users in the same group can be granted access to the same files and directories. A user can be part of multiple groups.

Each user has attributes such as the username, UID, GID, the default shell and the home directory. You can get information on the user by running the below command;
`id <username>`

## Types of Accounts
- User Account: These are user accounts for people using the system.
- Superuser Account: The superuser or root user has the UID 0. The root user has unrestricted and control over the system including over other users. 
- System Accounts: These are usually created during OS installation. These are used by software and services that will not run as the super user. UIDs for these accounts are usually below 100 or between 500 - 1000. They usually do not have a home directory but if they do, it will not be located under /home.
- Service Accounts: Similar to ssytem accounts but are created when services are installed in Linux. For example, Nginx makes use of a service account called Nginx.

## Commands
- `who`: Displays users currently logged into the system.
- `last`: Displays records of all logged in users. Also shows date and time when the system was rebooted.

## Switching Users
- `su -`: Switch to any user in the system, including root.
- `su -c "<command>"`: Allows you to run commands from another user without switching to that user.

## Sudo
Sudo allows non-root users to make root privileged commands without having to login as the root user. Only users who have been added to the /etc/sudoers file can use sudo commands. An example of these commands are;
`sudo yum install httpd -y`

