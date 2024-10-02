# Resetting and Reverting
There are several options which can help in undoing changes which we did not intend to make. One of these is reverting using the `git revert <commit_hash>`. This will create a new commit which removes the changes made from the commit hash specified. This is useful if you wish to undo changes but also want to keep the changes in your git history.

You can also use the reset command using `git reset --soft HEAD~<number_of_commits_wished_to_be_removed>` to keep all the changes that were made (back to staging) or `git reset --soft HEAD~<number_of_commits_wished_to_be_removed>` to lose all the changes made on that commit.

## Stashing
If you are working on a file but need to switch to another branch for whatever reason, to have a clean working area before changing branches, you can use the `git stash` command to save these changes without committing them. When you want to access the file again, you would use the `git stash pop` command to bring the file back to the working area.

To view all stashes, you can use the `git stash list` command. You can see the contents of the stash using the `git stash show stash@{stash_number}` command. 

If you have a specific stash you wish to bring back rather than the last file stashed, you would use the `git stash pop stash@{stash_number}` command.

## Reflog
The `git reflog` command shows all the actions which have taken plave on the repository. This would be useful if you have accidentally committed a hard reset and need the files back.