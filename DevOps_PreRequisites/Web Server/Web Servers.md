# Web Servers
When you are trying to access a website using a browser, your browser first makes a request to a web server. Once the web server receives the request, it will send a response to your browser which allows you to access the webpage. Server side code is usually written in Java, Python or NodeJS. Client side code are usually written in HTML, CSS and JavaScript.

Apache Tomcat, NGINX and GUnicorn are popular web servers which allow you to host web applications. They run one or more processes which listen on a particular port and performs the needed operations and responds to the client.

Web servers can also host multiple applications at the same time.

## Static vs Dynamic Websites
- *Static websites*: These are websites which only provides static content such as pictures, text, etc. These are usually written in HTML and CSS. There is no need to send requests to the web server after the content is served.
- *Dynamic Websites*: These are websites which perform multiple functions and there are constant changes. They usually have a database which the browser constantly needs to request information from. 