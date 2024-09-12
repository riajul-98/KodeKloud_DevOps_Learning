# Conditional Logic
We can give conditions which determine if a line of the script will run by using if statements. The syntax of an if statement in shell scripting can be seen below;

```
if [ condition to check ]
then
    Run this code
fi

```

You can also check additional conditions using elif (else if). The elif condition will only be evaluated if the if statement or a previous elif was not executed due to the conditions not being met. The syntax can be seen below:

```
if [ condition to check ]
then
    Run this code
elif [ another condition to check ]
then
    Run this code instead
fi
```

Lastly, you can add an else statement which only runs when no other conditions have been met. The syntax can be seen below:

```
if [ condition to check ]
then
    Run this code
elif [ another condition to check ]
then
    Run this code instead
else
    No conditions met. Run this code
fi
```

## Conparison Operators
To compare two values, we put the two values in square brackets with spaces between the bracket and the value and spaces between the value and the operand. See below for an example
`[ STRING1 = STRING2 ]`
Some of the operands we can use are:

| Example             | Description                                    |
|---------------------|------------------------------------------------|
| [ "abc" = "abc" ]   | If string1 is exactly equal to string2 (true)  |
| [ "abc" != "abc" ]  | If string1 is not equal to string2 (false)     |
| [ 5 -eq 5 ]         | If number1 is equal to number2 (true)          |
| [ 5 -ne 5 ]         | If number1 is not equal to number2 (false)     |
| [ 6 -gt 5 ]         | If number1 is greater than number2 (true)      |
| [ 5 -lt 6 ]         | If number1 is less than number2 (true)         |


We can also do the same with double square brackets. The difference being is double brackets offer more operations such as matching patterns using expressions (only in BASH). See below for an example
`[[ STRING1 = STRING2 ]]`
Some of the operands we can use are:

| Example                                      | Description                                                                          |
|----------------------------------------------|--------------------------------------------------------------------------------------|
| [[ "abcd" = *bc* ]]                          | If abcd contains bc (true)                                                           |
| [[ "abc" = ab[cd] ]] or [[ "abd" = ab[cd] ]] | If 3rd character of abc is c or d (true)                                             |
| [[ "abe" = "ab[cd]" ]]                       | If 3rd character of abc is c or d (false)                                            |
| [[ "abc" > "bcd" ]]                          | If “abc” comes after “bcd” when sorted in alphabetical (lexographical) order (false) |
| [[ "abc" < "bcd" ]]                          | If “abc” comes before “bcd” when sorted in alphabetical (lexographical) order (true) |


We can also use 'AND' and 'OR' operators to compare values. For AND, we would use &&. The syntax can be seen below;
- [ COND1 ] && [ COND2 ]
- [[ COND1 && COND2 ]]

For OR, we use ||. The syntax can be seen below;
- [ COND1 ] || [ COND2 ]
- [[ COND1 || COND2 ]]

Some examples can be seen below;

| Example                                | Description                                        |
|----------------------------------------|----------------------------------------------------|
| [[ A -gt 4 && A -lt 10 ]]              | If A is greater than 4 and less than 10            |
| [[ A -gt 4 \|\| A -lt 10 ]]            | If A is greater than 4 or less than 10             |


There are also some file level operators which can be seen below;

| Example         | Description                                |
|-----------------|--------------------------------------------|
| `[ -e FILE ]`   | if file exists                             |
| `[ -d FILE ]`   | if file exists and is a directory          |
| `[ -s FILE ]`   | if file exists and has size greater than 0 |
| `[ -x FILE ]`   | If the file is executable                  |
| `[ -w FILE ]`   | If the file is writeable                   |


# Loops

## For Loops
To do a set of tasks multiple times, we can use a for loop. The syntax is as below
```

for <item> in <list>
do
    Run this command
done

```

Rather than using a hard coded list, it is better to have a separate text file which contains all the values in the script. The syntax can be seen below;

```

for <item> in $(cat <list>)
do
    Run this command
done

```

You can also use the for loop for ranges using the below syntax;

```

for <item> in {0..100}
do
    Run this command
done

```

## While loops
For loops are used when you know the number of times you wish to run the commands. If you are not sure how many times this command needs to be run, you can use a while loop. The syntax can be seen below;

```
while [[ <condition> ]]
do
   Run this command
done

```


# Case Statements
Case statements are used to compare an input against a certain choice. It is much simpler than using if statements when making choices. The syntax can be seen below;

```

case $variable in
    1) Command1 ;;
    2) command2 ;;
    3) command3 ;;
    ...
esac

```