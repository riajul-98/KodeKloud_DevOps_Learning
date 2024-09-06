# User Management
- `useradd <username>`: Creates a new user with the specified username.
- `passwd <username>`: Sets the password for the specified user.
- `whoami`: Provides the current users username.
- `passwd`: Allows the user to change their username.

## Managing Users
Creating users with custom fields:

```

useradd -u 1009 -g 1009 -d /home/robert -s /bin/bash -c "Mercury Project Member" bob

-c: Custom Comment
-d: Custom Home Directory
-e: Expiry Date
-g: Specific GID
-G: Create user with multiple secondary groups
-s: Specify login shell
-u: Specific UID

```

- `id bob`: Provides UID, GID and groups of a user.
- `userdel <username>`: Deletes specified user.
- `groupadd <groupname>`: Creates a new group with the specified name. You can add -g to specify a specific group ID.
- `groupdel <groupname>`: Deletes specified group.