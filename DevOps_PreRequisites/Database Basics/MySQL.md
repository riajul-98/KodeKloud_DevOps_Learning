# MySQL
MySQL is an opensource database that is fast and reliable. It stores data as SQL format.

## Installing MySQL on CentOS
To install MySQL on CentOS, run the below commands;
- `wget https://dev.mysql.com/get/mysql80-community-release-el7-3.noarch.rpm`
- `rpm -ivh mysql80-community-release-el7-3.noarch.rpm`
- `yum install mysql-server -y`
- `systemctl start mysqld`

## Logs
The MySQL logs can be found in `/var/log/mysqld.log`
This file will display the bootup, the version of MySQL as well as the port which it listens on. The default port will be 3306. You would also need to view this file to access the temporary password to login to MySQL.

## MySQL Commands
`mysql -u root -pg/io%pFlE77m`: Logging into the MySQL command line utility. -u allows you to state the user and -p is the temporary password from the above file, no space between the -p and the password.
`ALTER USER 'root'@'localhost' IDENTIFIED BY 'MyNewPass4!';`: Changes the password for the specified user. Needs to be done the first time logged in.
`SHOW DATABASES;`: Displays a list of all databases.
`CREATE DATABASE school;`: Creates a new database with the specified name, in this case, school.
`USE school;`: Allows you to select the specified database.
`CREATE TABLE persons (Name varchar(255), Age....);` Creates a new table within the selected database. An example can be seen below
`SHOW TABLES;`: Provides a list of tables within the selected database.

```
CREATE TABLE persons
(
    Name varchar(255),
    Age int,
    Location varchar(255)
);

```

`INSERT INTO persons values ("John Doe", 45, "New York");`: Inserts new values into the specified table.
`SELECT * FROM persons;`: Displays table data.
`CREATE USER 'john'@'localhost' IDENTIFIED BY 'MyNewPass4!';`: Creates a new user for MySQL on the local system. For other systems, you can replace localhost with the IP address of the system needed to be accessed. If you want a user to access all systems, you can use % instead of localhost.
`GRANT <PERMISSION> ON <DB.TABLE> TO 'john'@'%';`: Grants specified permissions on the specified table. Permissions can be seen below;

| Privileges     |                         |
| -------------- | ----------------------- |
| ALL PRIVILEGES | Grant all access        |
| CREATE         | Create Databases        |
| DROP           | Delete Databases        |
| DELETE         | Delete rows from tables |
| INSERT         | Insert rows into tables |
| SELECT         | Read/query tables       |
| UPDATE         | Update rows in tables   |

Some example grant permission commands can be seen below;

```

GRANT SELECT ON school.persons TO 'john'@'%';

GRANT SELECT, UPDATE ON school.persons TO 'john'@'%';

GRANT SELECT, UPDATE ON school.* TO 'john'@'%';

GRANT ALL PRIVILEGES ON *.* TO 'john'@'%';


```

`SHOW GRANTS FOR 'john'@'localhost';`: Shows permissions for the specified user.