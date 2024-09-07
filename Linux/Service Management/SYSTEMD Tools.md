# SYSTEMD Tools
The main tools which are used with systemd are systemctl and journalctl. 

Systemctl is the main tool used to manage services on a systemd managed server. It can be used to start, stop, restart or reload a service as well as enable or disable services during the boot process. You can also use it to view information about and manage the system state. 

Journalctl is used to query the systemd journal and is a convenient troubleshooting tool to figure out issues such as service failures.

## Systemctl Commands
- `systemctl start httpd`: Starts the specified service (newer method)
- `systemctl stop httpd`: Stops the service.
- `systemctl restart httpd`: Restarts the service.
- `systemctl reload httpd`: Reloads the service.
- `systemctl status https`: Checks the status of the given service.
- `systemctl enable httpd`: Enables the service to start at system boot.
- `systemctl disable httpd`: Disables service from starting up at system boot.
- `systemctl daemon-reload`: This command should be run when changes are made to a service unit file to make systemd aware of the changes.
- `systemctl edit project-mercury.service --full`: Allows you to make changes to the system unit file specified. Doing it this way will allow you to make the changes immediately without having to run the daemon-reload command.
- `systemctl list-units --all`: Lists all units which have been loaded by systemd.

## Journalctl Commands
- `journalctl`: Prints all the entries from the oldest to the newest.
- `journalctl -b`: Prints logs from the current boot.
- `journalctl -u <unit_name>`: Prints entries for the specified unit. Good for debugging.