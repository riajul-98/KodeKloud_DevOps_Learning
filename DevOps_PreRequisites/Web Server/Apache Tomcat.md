# Apache Tomcat
Apache Tomcat provides a web server environment which allows for us to host Java based web applications.

## Installing Apache Tomcat
To install Apache Tomcat on CentOS, run the below commands;
- `yum install java-1.8.0-openjdk-devel -y`
- `wget https://downloads.apache.org/tomcat/tomcat-8/v8.5.53/bin/apache-tomcat-8.5.53.tar.gz`
- `tar xvf apache-tomcat-8.5.53.tar.gz`
- `./apache-tomcat-8.5.53/bin/startup.sh`

Apache Tomcat runs on port 8080.

## Apache Tomcat Directory
- bin: Contains scripts such as the startup and shutdown scripts.
- conf: Contains configuration files for Tomcat. Contains the server.xml file which allows you to change the port number Tomcat is listening on. Also contains the web.xml file which is used to configure web applications.
- logs: This is where the logs are stored.
- webapps: Where web applications hosted by Tomcat are located. This is where you put your web app files. The files should be packaged in .war format which can be done either by the jar command or by using tools such as Maven and Gradle. 