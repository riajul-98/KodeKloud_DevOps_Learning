# MongoDB
MongoDB is an open source, NoSQL database which stores data in JSON-like documents. It is a high performance, scalable database.

## Installing on CentOS
To install MongoDB on CentOS, run the following commands;
- Create a `/etc/yum.repos.d/mongodb-org-4.2.repo` file and add the below into the file;

```
[mongodb-org-4.2]
name=MongoDB Repository
baseurl=https://repo.mongodb.org/yum/redhat/$releasever/mongodb-org/4.2/x86_64/
gpgcheck=1
enabled=1
gpgkey=https://www.mongodb.org/static/pgp/server-4.2.asc

```
- Then run `yum install mongodb-org -y`
- `systemctl start mongod`

## Logs
Logs are stored in `/var/log/mongodb/mongod.log`. Here, you can find the database version and the port number which by default is 27017. These can be changed in the configuration settings found in `/etc/mongod.conf`.

## MongoDB Commands
`mongo`: Connects to the database. By default, access control is not restricted.
`show dbs`: Provides a list of databases.
`use school`: Switches to the school database. Or creates a new database with that name if no database present with that name.
`db`: Shows the current database.
`db.createCollection("persons")`: Creates a new collection with the specified name.
`show collections`: Displays a list of collections within the database.
`db.persons.insert({"name": "John Doe", "age": 45....})`: Inserts new data into the database. An example can be seen below;

```

db.persons.insert({
    "name": "John Doe",
    "age": 45,
    "location": "New York",
    "salary": 5000
})

```

`db.persons.find()`: Retrieves data from database.
`db.persons.find({"name": "John Doe"})`: Retrieves data from database with the filter specified.