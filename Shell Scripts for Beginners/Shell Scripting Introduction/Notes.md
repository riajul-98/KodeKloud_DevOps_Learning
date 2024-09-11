# Shell Scripting Introduction
Shell scripts can be run as commands. To do this, you need to add the path of the scripts to $PATH by running the below command;
- `export PATH=$PATH:<Path to scripts>`
Generally, its not a good idea to add the .sh extension to shell scripts which you want to export as commands.

To run a shell script, you would first need to add the execute permissions to the shell script using the `chmod +x <file_name>` command. After that you can run the shell script using the below methods;
- `bash <script name>`
- `./<scriptname>`

# Variables
Variables should be defined at the top of the file under the shebang. To call a variable, you would use $ followed by the variable name. The variable name can only contain alphanumberic characters and underscores. Varible names are also case sensitive. 

You can also use variables to store the results of another command such as below;

```

rocket-status=$(rocket-status $mission_name)
echo "Status of launch: $rocket_status"

```

You should only name variables with lowercase letters as this is best practice.


# Command Line Arguments
You are able to use arguments on the command line when running a script. To do this, you would add $1-$9 to the part of the script which you wish to use the command line argument for. $0 is the name of the script.

# Arithmetic Operations
There are a number of ways to perform arithmetic operations in shell scripts. One such way is using "expr" before the operation. For example:
- `expr 6 + 3`
- `expr 6 - 3`
- `expr 6 / 3`
- `expr 6 \* 3`
- `A=6 B=3      expr $A + $B`
The operator and values must be separated by a space. In the case of multiplication, you need to add a backwards slash before the multiplication operator. This is because the * symbol is a reserved regex character in shell.

You can also perform arithmetic operations using the double parentheses. Some examples can be seen below;
- `echo $(( A + B))`
- `echo $(( A - B))`
- `echo $(( A / B))`
- `echo $(( A * B))`
- `echo $(( ++A ))`: Increments a variable (Adds 1)
- `echo $(( --A ))`: Decrements a variable (Minuses 1)
- `echo $(( ++A ))`: Post-increments a variable (Current value of A is used first, then increased by 1)
- `echo $(( ++A ))`: Post-decrements a variable (Current value of A is used first, then decreased by 1)

Both the expr and double parentheses methods return whole numbers and not floating point values. To produce float values as the output, you can use the bc -l command. For example:
`echo $A / $B | bc -l`