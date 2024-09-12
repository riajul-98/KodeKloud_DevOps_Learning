# Shebang
Shebang is a line at the top of the shell scripts which denotes the type of script this is. There are different types of shells such as the bash shell, bourne shell and C shell. The Shebang will tell others what kind of script it is. For example, for bash shell, we denote this using the #!/bin/bash shebang. 

# Exit codes
When you run a command on Linux, it either succeeds or fails. The command will either return the expected output or an error message as well as an exit code. To view the exit code, you can run the `echo $?`. Exit status 0 means success and anything other than 0 is a failure. 

You can add exit codes to your script. For example, if you want exit code 1 when there are any failures, you can use the code `exit 1`. It is best practice for scripts to return appropriate exit codes.

# Functions
Functions in shell scripting are similar to that of Python where you write a function and can use it in your script multiple times. The syntax is as below;

```

function func_name() {
    function code
    .
    .
    .
    .
}

```

To call the function, you would type the function name in your code, followed by any arguments if any. When using functions, rather than using `exit <exit-code>`, you would use `return <exit-code>`. If we use an exit code, it will exit the whole script but using return will only exit the function and carry out the rest of the script. Command line arguments in a function are defined in the script and not in the command line.

To save the output of a function to a variable, you would assign it to the variable using the command substitution method. An example can be seen below;

```

function add(){
    echo $(( $1 + $2 ))
}

sum=$( add 3 5 )

```