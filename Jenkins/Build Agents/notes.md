# Build Agents
Build agents are used to run the CICD pipeline. They are responsible for testing, building, deploying and whatever other tasks you have defined in your pipeline script. Agents are essentially a task executer. Build agents can be Linux servers, Windows servers, Mac servers, Docker images etc. They just need to be able to run Java.

Jenkins tasks can be run directly by the Jenkins server, however, this is not recommended due to security and performance concerns. The Jenkins server may crash due to the resources needed to run the workload. 

## Adding a new build agent (Linux)
- Add a new user to the Jenkins server
- Add sudo privileges to the new user
- Manage Jenkins > Manage credentials > Add credentials
- Add new user details (username and password)
- Manage Jenkins > Manage nodes and clouds > new node
- Permanent agent
- Follow instructions, Launch method is "Launch agents via SSH" and give the public IP address for the host section
- For Host Key Verification Strategy, choose "Manually trusted key verification strategy"
