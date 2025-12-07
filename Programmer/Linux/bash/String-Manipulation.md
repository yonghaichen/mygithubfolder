## Assigning content to a variable and printing its content

In bash, ‘**$**‘ followed by the variable name is used to print the content of the variable. Shell internally expands the variable with its value. This feature of the shell is also known as **parameter expansion**. Shell does not care about the type of variables and can store strings, integers, or real numbers.

    VariableName='value'
    echo $VariableName

or

    VariableName="value"
    echo ${VariableName}

or

    VariableName=value
    echo "$VariableName"

There should not be any space around the “=” sign in the variable assignment. When you use `VariableName=value`, the shell treats the “=” as an assignment operator and assigns the value to the variable. When you use `VariableName = value`, the shell assumes that **VariableName** is the name of a command and tries to execute it.









## To print length of string inside Bash Shell

‘**#**‘ symbol is used to print the length of a string.

    variableName=value
    echo ${#variablename}








## Extract a substring from a string

`${string:position}`        --> returns a substring starting from `$position` till end.   
`${string:position:length}` --> returns a substring of `$length` characters starting from `$position`.







## Substring matching

In Bash, the shortest and longest possible match of a substring can be found and deleted from either front or back.

`${string#substring}`  To delete the shortest substring match from front of `$string`.   
`${string##substring}`  To delete the longest substring match from front of `$string`.  

`${string%substring}`  To delete the shortest substring match from back of `$string`   
`${string%%substring}` To delete the longest substring match from back of `$string`







　　　　　　　　     