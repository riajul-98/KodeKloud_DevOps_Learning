# GIT

## Source Code Management (SCM)
SCM allows developers to work with other developers without stepping on each others toes, track modifications in their code as well as who made these changes. Another name for SCM is Version Control Systems (VCS). There are many version control systems but Git is one of the most popular amongst developers.

Git helps to track changes in a set of files within a directory as well as helps multiple developers merge their work together.

## Using Git
- To install Git, run the `yum install git -y` command.
- To check the version installed, run the `git version` command.
- To initiate a directory to be tracked by Git, you would use the `git init` command. This creates a hidden directory called `.git` under the current directory which stores all information on changes relating to this directory and files.
- To check the status of files in the directory, use the `git status` command.
- To stage changes, run the `git add main.py README.md...` command with all the files you wish to track. If you want to add all the changes to staging, you can use `git add .`
- To commit the changes, you would run the `git commit -m "Commit Message"` command. You would need to give a commit message.