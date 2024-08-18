# Remote Repositories
A remote repository is a central location where all the changes made can be pushed to. This allows for all the changes made by the developers to be merged together. If two developers were working on the same file and pushed the changes to the remote repository, Git would most likely be able to merge the changes. However, if both developers change the same line in the same file, a conflict emerges and you would need to tell Git which change needs to be implemented. 

Once the changes have been pushed and merged, all the developers can then pull the latest version of the repository.

You can either install the Git daemon on your own server to host a remote repository or you can use one of the various public remote repositories such a GitH ub and GitLab.

## Git Commands for Pushing to a Remote Repository
To push to a remote repository, you would first need to install git and initialise the directory as a local repository using the `git init` command. You would then stage the files in your repository using the `git add .` command and commit these using `git commit -m "Initial Commit"`. Now you would need to create the remote repository. For example, you can create one on GitHun using the GUI. We would then tell youyr local system there is a remote repository you wish to add using the `git remote add <name> <url>`, where the name can be anything you want. You can then push the changes using `git push <name> <branch>`. The first time you push code, you would need to add the "-u" flag. This would look like `git push -u <name> <branch>`.

Now that the files are on the remote repository, other developers in your team can pull the repository onto their local system using the `git clone <url>`.

If you have pushed more files to the remote repository the other developer wants the most up to date version, they can use the `git pull` command.