# SSH and SCP

## SSH
SSH is used to login to and execute commands on a remote computer. You do this with the below commands:
- `ssh <hostname or IP address>`
- `ssh <user>@<hostname or IP address>`
- `ssh -l <user> <hostname or IP address>`

For the connection to work, the server should have port 22 open. You would also need either the username and password or an SSH key which is a key which allows you to login to the host without a password.

## Passwordless SSH
To login using passwordless SSH, you would need to first generate a keypair on the client. These would be a public key and private key. The public key would be added to the remote server and the private key would be something you have which authenticates you to login to the server. 

To set up passwordless SSH do the following:
- `ssh-keygen -t rsa`: Creates a new keypair. The public key is stored in /home/<user>/.ssh/id_rsa.pub. The private key is stored in /home/<user>/.ssh/id_rsa
- You now need to copy the public key to the remote server. You can do this with the below command:
`ssh-copy-id <user>@<hostname>`
You would need to enter the password for your user on the remote server.
- Now you should be able to SSH into the remote server without a password. The public key will be stored in /home/<user>/.ssh/authorized_keys

## SCP
SCP allows you to copy data over SSH, meaning you can copy files from your laptop or a remote server to another remote server in which you have SSH access. To do this, use the below command;
- `scp <location of file to be copied> <hostname>:<location you wish file to be copied to>`
- e.g. `scp /home/bob/caleston-code.tar.gz devapp01:/home/bob`
- For directories with files: `scp -pr /home/bob/media/ devapp01:/home/bob`: The -r flag means recursive and -p means preserve ownership details.