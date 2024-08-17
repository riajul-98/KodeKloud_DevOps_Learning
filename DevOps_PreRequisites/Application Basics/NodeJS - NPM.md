# NodeJS
NodeJS is a form of JavaScript which is used for backend/server-side development. It is free, open source and cross platform compatible. What makes it popular is its ability to handle a large number of concurrent connections.

## Installing NodeJS
- `curl -sL https://rpm.nodesource.com/setup_13.x | bash -`
- `yum install nodejs -y`

## NodeJS Commands
- `node -v`: Displays installed version of NodeJS.
- `node add.js`: Runs the specified NodeJS file.

# Node Package Manager (NPM)
NodeJS supports a large number of packages and libraries that are developed by either the core team or by the community. These include files, web servers, databases, security updates and many more. These are located in a public repository called npmjs.com. NPM allows developers to develop new reusable packages and modules and share them on the public repository which can then be downloaded by other developers and used. 

## NPM Commands
- `npm -v`: Displays installed version of NPM.
- `npm search file`: Searches for the specified package, in this case 'file'.
- `npm install file`: Installs npm package. Package will be installed in your present working directory under ./node_modules/. Code is stored in ./bode_modules/lib. There will also be a package.json file which contains the metadata of the package. 
- `node -e "console.log(module.paths)"`: Displays paths which node looks at while importing a package to an application.
- `npm install file -g`: Installs package globally.

## Common Modules
Built in modules are installed automatically when NodeJS is installed. They are located in /usr/lib/node_modules/npm/node_modules/. Some example built-in modules are fs, http, os, events, tls and url.

External modules are modules which are installed by the user. These include express, react, debug, async, lodash. They can be located at /usr/lib/node_modules/.

## Application Dependancies
Applications usually come with dependancies. In NodeJS and NPM, these and their version numbers can be found in the metadata file called package.json