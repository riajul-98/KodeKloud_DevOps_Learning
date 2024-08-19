# NodeJS - ExpressJS
ExpressJS is a popular NodeJS based web framework.

## Commands
- `npm install`: Installs all dependant packages which are listed in the package.json file.
- `node app.js`: Runs the application. Must specify the starting code, in this case, app.js.
- `npm run start`: npm run allows you to run a specified script which is defined in the package.json file, in this case, the script is start.

After running these commands, you should be able to view the application from a browser on port 3000.

This is not best practice as if anything fails, node shuts down the application. It is better to use other tools such as supervisord, forever and pm2

## PM2
PM2 is a process manager which runs NodeJS based applications which comes with a built-in load balancer. It keeps the application running and reduces downtime.

## PM2 Commands
- `pm2 start app.js`: Starts the specified application.
- `pm2 start app.js -i 4`: Starts the specified application on multiple instances in cluster mode. In this case, 4 instances.