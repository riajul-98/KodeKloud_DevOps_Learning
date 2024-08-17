# Java
Java is one of the most popular programming language, used to build desktops and web applications. It is free, open-source and has a huge community. Version 9 introduced changes which made many popular tools and libraries incompatible and therefore, many companies opted not to upgrade to version 9. Therefore, when working in an IT company, you may notice them still using version 8 or below.

## Installing Java on Linux
- `wget httpd://download.java.net...`
- `tar xvf openjdk-13.0.2_linux-x64_bin.tar.gz`
- `jdk-13.0.2/bin/java -version`

Java 8 and below will display a version like 1.8.0. From Java 9 and above, these will be similar to 9.0.2.

## Java Development Kit (JDK)
JDK is a set of tools that will assist in developing, building and running applications in Java. Some of the tools include;
- jdb: Helps in debugging code.
- javadoc: Assists in writing documentation for your application.
- javac: Compiles code.
- jar: Archives source code and libraries into a single file.
- JRE: Java Runtime Environment. Needed to run a java application. Before version 9, JRE is separate to the JDK. After version 8, they are installed together.

These commands and many more can be found in /bin/jdk-{version number} directory.


# Java - Build and Packaging

## Build Process
- Develop Source Code
- Compile: `javac MyClass.java` - Compiles code
- Run: `java MyClass` - Runs application (don't specify .class extension)

## Compile
Compiling code means to convert human readable code to machine readable code. After compiling, the "byte code" is run in Java virtual machines (JVM). This allows us to run the application anywhere where there is a supported JVM available.

## Package
A typical application has many files which may be dependant on each other or have dependancies on external libraries. In order to distribute the application to end users, these files and libraries need to be packaged together. To do this, we can use an archiver such as JAR (Java Archive) which helps compress and archive Java class files and dependant libraries into a single package.

Some applications also have static content such as images or HTML files. In this case, they would need to be packaged into a WAR (web archive) file. To create an archive using JAR, use the below command;
`jar cf MyApp.jar MyClass.class Service1.class Service2.class ...`
In this case, MyApp.jar will be the resultant jar file and the other files named are the files to be archived. A manifest file will also be produced, located at META-INF/MANIFEST.MF, which contains information about the archived files and any other metadata (including the entrypoint of the application). 

The application can then be run using the below command;
`java -jar MyApp.jar`

## Documentation
Documentation is important to tell users the functionalities of the application. You would need to run the `javadoc` command and specify the source code as the input. An example of this can be seen below;
`javadoc -d doc MyClass.java`

## Build Tools
Build processes can get very complex when there are large applications and multiple developers are working on the same application at the same time. Different developers would be pushing different updates and these would then need to be compiled, packaged and documented. Build tools such as Maven, Gradle and ANT assist in automating these processes. These tools use configuration files which can be used to specify which steps need to be run and in what order.

## ANT
The below is an ANT configuration file;

```
<xml version="1.0">
<project name="Ant" default="main" basedir=".">
    <!-- Compiles the Java code (including the usage of library for JUnit) -->
    <target name="compile">
        <javac srcdir="/app/src" destdir="/app/build">
        </javac>
    </target>
    <!-- Creates Javadoc -->
    <target name="docs" depends="compile">
        <javadoc packagenames="src" sourcepath="/app/src" destdir="/app/docs">
            <!-- Define which files / directory should get included, we include all -->
            <fileset dir="/app/src">
                <include name="**"/>
            </fileset>
        </javadoc>
    </target>
    <!-- Creates the deployable jar file -->
    <target name="jar" depends="compile">
        <jar basedir="/app/build" destfile="/app/dist/MyClass.jar" >
            <manifest>
                <attribute name="Main-Class" value="MyClass" />
            </manifest>
        </jar>
    </target>
    <target name="main" depends="compile, jar, docs">
        <description>Main target</description>
    </target>
</project>
```
After creating the configuration file, we would run the build using the `ant` command. This will then run the steps one after the other and create the JAR files and documentation files.

If you wish to only run a specific section from the configuration file, you can specify this by adding the target names after the ANT command. For example;
`ant compile jar`
This would compile the code and then archive the source code into a JAR file.

