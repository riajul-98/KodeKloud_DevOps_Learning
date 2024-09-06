# File Permissions and Ownership
When using the `ls -l <file_name>` command, you will see something similar to: 
'-rwxrwxr-x bob bob 89 Mar 17 01:35 bash-script.sh'

The initial character, in this case '-' represents the file type. The file types can be seen below;

| File Type         | Identifier |
|-------------------|------------|
| DIRECTORY         | d          |
| REGULAR FILE      | -          |
| CHARACTER DEVICE  | c          |
| LINK              | l          |
| SOCKET FILE       | s          |
| PIPE              | p          |
| BLOCK DEVICE      | b          |

The characters after the file type are known as the file permissions. These can either be 'r', 'w', 'x' or '-'. From left to write, the first 3 characters after the file type are the owner permissions (u), the 3 that follow are group permissions (g) and the remaining 3 are permissions for others (o). You can see what the characters stand for in the below table;

| Bit | Purpose        | Octal Value |
|-----|----------------|-------------|
| r   | Read           | 4           |
| w   | Write          | 2           |
| x   | Execute        | 1           |
| -   | No permissions | 0           |

These are still applicable for directories as well as files.

The Octal values can represent the permissions a a user, group or other has to the file or directory. For example, if the owner has all permissions, the octal value for user will be 4+2+1 = 7.

## Modifying File Permissions
To modify file permissions, you would use chmod, this can be done in two ways as seen below:
- `chmod u+r-x <file_name>`: Symbolic method. Here, we are adding read and execute permissions to the user for the specified file.
- `chmod 770 <filename>`: Numeric method. In this example, we are adding all permissions for user and groups but no permissions for others.

## Modifying File Ownership
To change the ownership and group of a file, you would use the below command;
- `chown michael:developers <file_name>`: Here, we are giving michael ownership with group ownership to the developers group.
- `chown michael <file_name>`: Here, we are giving michael ownership but no groups are changed.
- `chgrp developers <file_name>`: Here, we are giving group ownership to developers but no change in user ownership.