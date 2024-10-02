# Git Branches
The default branch is the master / main branch. It is best practice to make changes on separate branches before merging the changes into the main branch once the new changes have been tested thoroughly to avoid introducing bugs. A branch is a pointer to the last commit. 

You can create a new branch by running the `git branch <branch_name>` command or `git checkout -b <branch_name>` command to create and switch to the new branch. This will create a new branch which includes all the previous commits from the main branch. Any new commits from this new branch will only be added to the new branch and not the main until a merge happens. 

You can switch to another branch by using the `git checkout <branch_name>` command.

To view a list of all branches in the repository, you can run the `git branch` command. There will be an asterisk near the name of the branch you are currently on.

To delete a branch, you would run the `git branch -d <branch_name>` command.

HEAD: This is where you currently are in the repository.

## Git Merging Branches
To merge your changes from your branch to the main/master branch, you would first need to switch over to the main branch and then run the `git merge <branch_name>` command with the branch name being the branch which you wish to merge into the main branch. You can also merge into other branches instead of the main/master. 

There are two types of merging:
- Fast-forward merge: Merging when there have been extra no commits in the current branch compared to the branch we are merging. This does not create a new commit, rather git merges the commit from the branch merging into the current branch.
- No fast-forward merge: Merging when there have been extra commits in the current branch compared to the branch we are merging. This creates a new merging commit on the current branch. 