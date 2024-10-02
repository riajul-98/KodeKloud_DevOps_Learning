# Local and Remote Repositories

## Git repository types
- Local: Repository on your own machine. 
- Remote: A centralised server where changes made locally can be pushed to. Useful for backups and collaboration with other developers.

## Stages of a local repository
- Working area: Contain all your active changes.
- Staging area: Contains new changes which are ready to be commited.
- Commited files: These are files which have been commited.

# Initialising a GIT repository
To initialise a directory as a git repository, you would navigate to the directory and nrun the below command:
`git init`

This would create a .git directory.

To check the status of any changes, you can run the `git status` command. Untracked files are files which have not yet been added to the staging area. You can then add the changes to the staging area by running the `git add <filename>` for a specific file changes or `git add .` for all changes. To then commit the changes in the staging area, you run the `git commit -m "<comment on this commit>"` command. The comment should be useful to allow other developers know what you have changed.

Before committing, you would need to configure git to add a name and email address to allow others to know these changes have been made by you. To do this, you would run `git config user.name <name>` and `git config user.email <email address>`.

# Git Log
To view previous commits, you can run the `git log` command. This provides the git hash, the author name, the date commited and the commit message. You can get less information using the `git commit --oneline` command. You can also list the files chnaged using the `git log --name-only` command.