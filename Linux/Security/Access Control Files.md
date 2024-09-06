# Access Control Files
Most of the access control files are stored in the /etc directory. This directory can be read by any user by default but can only be written by the root user. The access control files are designed in a way that they should never be modified using a text editor. Instead, use the built-in commands to add or modify access as needed. 

## Files
- /etc/passwd: Contains basic information about users in the system including username, UID, GID and home directory
- /etc/shadow: Contains password information for users. Contents of this file are hashed.
- /etc/group: Stores information about all user groups on the system.

## File fields
- /etc/passwd: USERNAME:PASSWORD:UID:GID:GECOS:HOMEDIR:SHELL   GECOS is a csv format of user information such as fullname, location, etc
- /etc/shadow: USERNAME:PASSWORD:LASTCHANGE:MINAGE:MAXAGE:WARN:INACTIVE:EXPDATE
- /etc/group: NAME:PASSWORD:GID:MEMBERS