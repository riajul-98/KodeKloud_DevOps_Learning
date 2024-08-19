# Apache Web Server
The Apache Web Server is an open-source HTTP server, created by Apache. It is usually used to host web content such as JavaScript, HTML and CSS files. It can be used with an application server that acts as a backend.

## Install Apache Web Server
To install Apache Web Server on CentOS:
- `yum install httpd -y`
- `systemctl start httpd`

If you have a firewall in your system, you would need to add a rule to allow https traffic by using the following command;

`firewall-cmd --permanent --add-service=http`

## View Logs 
Every service comes with a way to view logs. To view the access logs for Apache Web Server, run the below command;

`cat /var/log/httpd/access_log`

And for error logs;

`cat /var/log/httpd/error_log`


## Config File
With the configuration file, you are able to make many changes. Some of these include changing the port which the service is listening on, DNS name, location of the logs and changing the path for where the web application files are located. The configuration file can be found at `/etc/httpd/conf/httpd.conf`.

You can also use the same configuration file to configure multiple web applications using a virtual host block. See below for an example;

```
<VirtualHost *:80>
    ServerName www.houses.com
    DocumentRoot /var/www/houses
</VirtualHost>

<VirtualHost *:80>
    ServerName www.oranges.com
    DocumentRoot /var/www/oranges
</VirtualHost>
```

You would need to add the same server name to the /etc/hosts file. If you have made any changes to the configuration file, you would need to restart the service using the below command;

`systemctl restart httpd`

You can also move the virtual hosts to separate configuration files. For example, the www.houses.com configuration would be in the `/etc/httpd/conf/houses.conf` file and would look like the below;

```
<VirtualHost *:80>
    ServerName www.houses.com
    DocumentRoot /var/www/houses
</VirtualHost>
```

Similarly, the www.oranges.com configuration would be in the `/etc/httpd/conf/houses.conf` file and would look like the below;

```
<VirtualHost *:80>
    ServerName www.oranges.com
    DocumentRoot /var/www/oranges
</VirtualHost>
```

You would then add `Include conf/houses.conf` and `Include conf/oranges.conf` statements into the original configuration file.