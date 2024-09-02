# Bash Shell

## Shell Types
- Bourne Shell (sh)
- C Shell (csh or tcsh)
- Korn Shell (ksh)
- Z Shell (zsh)
- Bourne Again Shell (bash)

- To check the shell being used, use the `echo $SHELL` command.
- To change the shell, use the `chsh` command which will prompt the password and then the new shell. You will then need to login to a new session to see the change.

## Bash Shell Features
- Bash Auto-Complete: You can write a partial command and bash will attempt to complete this when you press the tab key.
- Alias: We can set custom aliases for commands. For example to set an alias, dt for the date command, we can use `alias dt=date`.
- History: Using the `history` command lets you see the list of commands you have run.

## Bash Environment Variables
- Used to store information. $SHELL contains the type of shell used. 
- To see all environment variables, use the `env` command.
- To set a new variable, you would use the export command. An example for this is `export OFFICE=caleston`. You can also use `OFFICE=caleston` but this will only apply on the shell and will not be used in any future process.
- To make these variables persistant, add these to the ~/.profile or ~/.pam_environment file.

## Path Variables
- When a user issues an external command into the shell, the shell uses a path variable to search for these external commands. You can see these using `echo $PATH`.
- To check if the location of a command can be identified, use the which command. For example `which obs-studio`.
- To add a path variable, we would use the export command in the following way `export PATH=$PATH:/opt/obs/bin`, where $PATH includes all the previous paths added.

## Bash Prompt
- When you login to your Linux command line, the line which you see is called the command prompt. This can provide useful information such as the logged in user and the name of the system.
- The bash prompt is set and controlled by a set of special shell environment variables, the most common is PS1.
- To view the value assigned to PS1, run the `echo $PS1` command.
- To change the prompt, you can change the value of the variable using the `PS1="ubuntu-server:"` command.
- The documentation provides more special options which can be used to customise the prompt.