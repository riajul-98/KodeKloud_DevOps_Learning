# More Linux Commands

## User Accounts
- `whoami`: Displays the current account name.
- `id`: Provides more details about the user such as user ID, group ID and groups.
- `su John`: Switches to specified user, in this case, John.
- `ssh john@196.168.1.2`: Logs into different server with the specified IP address as user John. Can also use hostname instead IP address if hostname and IP are specified in the /etc/hosts file.
- `sudo ls /root`: sudo allows non-root users to do actions which only the root account can do. Users muct be specified in the /etc/sudoers file.

## Download Files
- `curl https://www.some-site.com/some-file.txt -O`: Downloads file from the internet. Need to add -O flag to save results to a file.
- `wget https://www.some-site.com/some-file.txt -O some-file.txt`: Downloads file from the internet. Need to add -O flag to save results to a file.

## Check OS Version
- `ls /etc/*release*`: Displays OS release files.
- `cat /etc/*release*`: Displays OS details.