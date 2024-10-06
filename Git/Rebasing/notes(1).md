# Rebasing
Rebasing allows you to pull all the latest changes from the main branch into your own branch instead of merging the master branch into your branch. You can do this by running the `git rebase main`. You can also do this with other branches rather than using the main branch. Rebasing also provides a new commit hash.

## Interactive rebasing
The rebase command also allows us to modify the commit history on a certain branch before rebasing it. We can do this by running the `git rebase -i HEAD~<number_of_commits_to_change>` command. For example, using `git rebase -i HEAD~4` will allow us to modify the last 4 commits.

## Cherry picking
You would like to add a certain commit from one branch to another but don't want all the other commits which have been made on that branch. We can use the `git cherry-pick <commit_hash>` command.