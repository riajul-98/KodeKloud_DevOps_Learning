# Remote repository

## Initialising a remote repository
To initialise a remote repository, you would first create a remote repository using GitHub, GitLab, BitBucket or a different site. You would then be presented with a connection string. You would then navigate into the directory on your system which you want to initiate the remote repository in and run the `git remote add origin <connection_string>` command. "Origin" in this command is the alias for the repository.

To list remote repositories, you can use the `git remote -v` command. 

## Pushing to remote repository
To push changes to the remote repository, you would use the `git push <alias_name> <branch_name>` command. 

## Cloning a remote repository
To access all the data / files in a remote repository, you would clone the repository. You can do this with the `git clone <git link>`. The link can be of https or ssh. This will create a local folder with the same name as the remote repository.

## Pull requests
When working with a team, it is not a good idea to merge changes directly to the main branch. Instead, it is better to open a pull request. This can be done by first commiting changes and pushing the changes to GitHub or any other Git repository website. You would then see a message under the main branch which says "This branch is 1 commit behind branch_name" with a pull request button next to it. Click that button and then click on "Create pull request". You can add more details on the pull request before submitting. Once submitted, other team members can see these changes and can be approved for merging into the main branch. Merging can only be approved by someone with the right privileges.

## Fetching and pulling
When a change is made in the remote repository, the local repository is not automatically updated. To update the local repository, you would use the `git fetch <alias_name> <branch_name>`. If any changes were made on the local repository before fetching, you can merge these changes using the `git merge <alias_name>/<branch_name>`. You can fetch and merge using one command using the `git pull <alias_name> <branch_name>` command.

## Merge conflicts
In most cases, git is smart enough to know how to modify files when merging branches. But this is not always the case. If two developers are working on the same file in their own branches and one of them merges their branch into the main branch, it could create a merge conflict when the second developer tries to merge their branch into the main branch. Git will display these conflicts and allow you to remove the lines you wish to remove and keep lines you with to keep.

## Fork
There are many open source projects which can be found on sites such as GitHub. In order for us to contribute to projects like these where we do not have permissions to directly make changes to the developers code, we would fork the main project. We can then add our own changes to the forked copy and then send a pull request to the original developers for them to review and potentially add into their project. Forking essentially allows you to have your own copy of the project.