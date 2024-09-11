The Shell Scripting for Beginners section of the courses uses rocket launches as examples on how to write scripts. The rocket launch process is as below:

1. Start Auxiilary Power
2. Switch to Internal Power
3. Auto Sequence Start
4. Main Engine Start
5. Lift Off

The Linux environment has been configured to use the below commands to run the rocket launch process;

1. `rocket-start-power`
2. `rocket-internal-power`
3. `rocket-start-sequence`
4. `rocket-start-engine`
5. `rocket-lift-off`

There will also be other rocket related commands. These are:
1. `rocket-ls`: Lists all rockets
2. `rocket-add`: Creates a new rocket
3. `rocket-del`: Deletes a rocket
4. `rocket-status`: Gives the status of a rocket
5. `rocket-debug`: Debugs issues with rocket in case of failure

I have added all the scripts which they have used to create these commands in the "Rocket command scripts" folder. To add these scripts as a command, you will need to add the path to these scripts to $PATH by running the following command;
`export PATH=$PATH:<Path to scripts>`