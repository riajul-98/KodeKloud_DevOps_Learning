# Creating a SYSTEMD Service
In order to start a programme or script up during boot up, you would need to add it as a Systemd service. If you want to add one of your scripts or program as a service, you would need to add it to the /etc/systemd/system/ directory as a .service file. For example, in the course, Bob wants to run his bash script when the system boots up. He would need to create a file named /etc/systemd/system/project-mercury.system and configure it like below;

```
[Unit]
Description=Python Django for Project Mercury
Documentation=http://wiki.caleston-dev.ca/reported

# Start Python Application after Postgres DB
After=postgresql.service

[Service]
# Program - /usr/bin/project-mercury.sh
ExecStart=/usr/bin/project-mercury.sh

# Use Service Account project_mercury
User=project_mercury

# Auto Restart on Failure
Restart=on-failure

# Restart Interval 10 seconds
RestartSec=10

# Log Service Events (handled by systemd by default)

[Install]
# Load when booting into Graphical Mode
WantedBy=graphical.target
```