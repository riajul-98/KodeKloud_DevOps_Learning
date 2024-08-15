# Services
Some packages such as web servers or database servers may need to be running and need to stay running after servers are rebooted. You may also need to make sure the services are started in the right order if some services are dependant on others. Services aid in allowing you to make sure software are running in the background and are automatically running in the correct order after start up.

## Commands
- `service httpd start`: Starts the specified service (older method)
- `systemctl start httpd`: Starts the specified service (newer method)
- `systemctl stop httpd`: Stops 
- `systemctl status https`: Checks the status of the given service.
- `systemctl enable httpd`: Enables the service to start at system boot.
- `systemctl disable httpd`: Disables service from starting up at system boot.

## How to configure a programme or application as a service
Say we have a python programme located at /opt/code/my_app.py, running on port. You can run the code by using the following command:
`/usr/bin/python3 /opt/code/my_app.py`
And then access the application with:
`curl http://localhost:5000`

To configure this as a service, you need to configure the programme as a systemd service. To do this, you need to add a systemd unit file at /etc/systemd/system. The file needs to be named as something which you wish for the application to be known as e.g. my_app.service. An example file can be seen below;

```
[Service]
ExecStart=/usr/bin/python3 /opt/code/my_app.py
```

You would then use `systemctl daemon-reload` command to let the system know there is a new service which has been configured. You can then start the service with `systemctl start my_app`.

To configure the service to automatically start on boot up you would need to configure the unit file to add an "Install" section with "WantedBy-multi-user.target". After this, you would run the command `systemctl enable my_app` See below for an example of the updated unit file;

```
[Service]
ExecStart=/usr/bin/python3 /opt/code/my_app.py

[Install]
WantedBy=multi-user.target
```

Additional information can also be added to the unit file such as a description to tell other users about the service. An example of this can be seen below;

```
[Unit]
Description=My python web application

[Service]
ExecStart=/usr/bin/python3 /opt/code/my_app.py

[Install]
WantedBy=multi-user.target
```

If your application has other dependancies such as scripts to be run before or after the application is run, you can add these to the unit file as can be seen below;

```
[Unit]
Description=My python web application

[Service]
ExecStart=/usr/bin/python3 /opt/code/my_app.py
ExecStartPre=/opt/code/configure_db.sh
ExecStartPost=/opt/code/email_status.sh

[Install]
WantedBy=multi-user.target
```

For the application to automatically restart in the event of it crashing, you would add the 'restart=always' flag under the service section, as can be seen below;

```
[Unit]
Description=My python web application

[Service]
ExecStart=/usr/bin/python3 /opt/code/my_app.py
ExecStartPre=/opt/code/configure_db.sh
ExecStartPost=/opt/code/email_status.sh
Restart=always

[Install]
WantedBy=multi-user.target
```

There are many more configuration flags which can be added to a file. The below shows the unit file for the Docker service;

```
[Unit]
Description=Docker Application Container Engine
Documentation=https://docs.docker.com
BindsTo=containerd.service
After=network-online.target firewalld.service containerd.service
Wants=network-online.target
Requires=docker.socket

[Service]
Type=notify
ExecStart=/usr/bin/dockerd -H fd:// --containerd=/run/containerd/containerd.sock
ExecReload=/bin/kill -s HUP $MAINPID
Restart=always
StartLimitBurst=3
StartLimitInterval=60s
LimitNOFILE=infinity
LimitNPROC=infinity
LimitCORE=infinity

[Install]
WantedBy=multi-user.target
```