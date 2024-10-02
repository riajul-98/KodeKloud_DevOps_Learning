# Managing Users and Teams
To create users you would go into: 
Manage Jenkins > Manage Users > Create User

To manage role based access, you would need to first install the "Role-based Authorization Strategy" plugin. Once this is installed, you would see the "Manage and Assign Roles" option under Manage Jenkins. Here, you can create roles and provide different levels of access as well as assign users to these roles. This includes anonymous users.

# Managing the System and Credentials
To configure the system, you would go into: 
Manage Jenkins > Configure System

Here, you can add or remove executers, labels, usage, Jenkins URL, health metrics and many more.

You can also add credentials for other services such as GitHub, DockerHub and AWS. These can be added by navigating to:
Manage Jenkins > Manage Credentials.