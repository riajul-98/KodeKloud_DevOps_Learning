# 2 Tier Application
The Kodekloud e-commerce website is a LAMP stack application which sells electrical products. LAMP stack is an application which is deployed on Linux, with Apache server, MySQL database and written in PHP. This project uses MariaDB instead of MySQL. 

## Installing and configuring Firewall
- `sudo yum install firewalld -y`: Installs firewalld.
- `sudo service firewalld start`: Starts the firewalld service.
- `sudo systemctl enable firewalld`: Allows the service to start at reboot.

## Installing and configuring MariaDB database
- `sudo yum install mariadb-server -y`: Installs MariaDB.
- `sudo vi /etc/my.conf`: Add the necessary information to the mariaDB config file.
- `sudo service mariadb start`: Starts the MariaDB service.
- `sudo systemctl enable mariadb`: Allows the service to start at reboot.
- `sudo firewall-cmd --permanent --zone=public --add-port=3306/tcp`: Configures firewall for this service.
- `sudo firewall-cmd --reload`: Reloads the firewall.
- Configure database with the below commands:

```

mysql
MariaDB > CREATE DATABASE ecomdb;
MariaDB > CREATE USER 'ecomuser'@'localhost' IDENTIFIED BY 'ecompassword';
MARIADB > GRANT ALL PRIVILEGES ON *.* TO 'ecomuser'@'localhost';
MARIADB > FLUSH PRIVILEGES;

```

- `mysql < db-load-script.sql`: Loads data into database.

## Installing and Configuring Apache Web Server
- `sudo yum install httpd -y`: Installs HTTPD service.
- `sudo firewall-cmd --permanent --zone=public --add-port=90/tcp`: Adds firewall rule for this service.
- `sudo firewall-cmd --reload`: Reloads the firewall.

- `sudo vi /etc/httpd/conf/httpd.conf`: Configure this file to use index.php file instead of index.html.
- `sudo systemctl start httpd`: Starts HTTPD service.
- `sudo systemctl enable httpd`: Allows the service to start at reboot.

## Download code
- `git clone https://github.com/kodekloudhub/learning-app-ecommerce > /var/www/html`: Clones git repository

## Multi-node Setup
To do the same on a multi-node setup, the database will be separate to the Apache web server and PHP code. Most instructions are the same but the index.php file would need the database IP changing to the server of the database as well as in the database, instead of using 'ecomuser'@'localhost', you would need to use 'ecomuser'@'{ip-address}'