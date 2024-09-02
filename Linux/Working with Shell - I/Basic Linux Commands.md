# Basic Linux Commands
- `pwd` - Displays the present working directory.
- `mkdir Asia` - Creates a new directory with the name provided in the argument. In this case, Asia.
- `mkdir Europe Africa America` - Creates multiple directories with the names provided in the argument.
- `ls` - Displays files and directories in the present working directory.
- `cd Asia` - Changes directory to the one specified in the argument, in this case, Asia.
- `mkdir India/Mumbai` - Creates a directory inside the India directory with the name Mumbai without having to change directories from the Asia directory.
- `mkdir -p India/Mumbai` - Creates a directory inside the India directory with the name Mumbai without having to change directories from the Asia directory or create the India directory.
- `cd ..` - Goes up a directory.
- `cd` - Running cd without any arguments takes you to your home directory.

## Absolute Path and Relative Path
Suppose you are in your home directory and wish to change directory to the Asia directory, you can do this in two different ways. You can use the relative path which would use the command `cd Asia` or you can use the absolute path which uses the whole path from the root directory (/) to the chosen directory. You can do this using the command `cd /home/michael/Asia`.

The absolute path is useful when you are in a directory and wish to change directory into a directory which is in a completely different path. For example, moving from the /etc path to the /var/www/html path would be as easy as `cd /var/www/html`.

## Pushd/Popd
Pushd remembers the most previous directory before the last change in directory command. You would then use the popd command to move back to that directory. For example, you are in the /home/michael directory and you wish to move to the /etc directory, you would use `pushd /etc`. You can then move to as many different directories as you wish using cd. To then go back to the original directory when the pushd command was made, you would use `popd` to go back to /home/michael.

## More Basic Linux Commands
- `mv /home/michael/Europe/Morocco /home/michael/Africa` - Moves specified file to the specified path. In this case, we are moving Morocco to Africa from Europe. Can also use relative path.
- `mv Asia/India/Munbai Asia/India/Mumbai` - The mv command can also be used to rename files/directories. In this example, we are renaming Munbai to Mumbai.
- `cp Asia/India/Mumbai/city.txt Africa/Egypt/Cairo` - Copies file from one directory to another. In this case from Mumbai to Cairo.
- `rm Eutope/UK/London/Tottenham.txt` - Deletes the specified file.
- `cp -r Europe/UK Europe/UnitedKingdom` - Copies recursively i.e. the folder and everything in it.
- `cat Asia/India/Mumbai/City.txt` - Allows you to read contents of the specified file.
- `cat > Africa/Egypt/Cairo/City.txt` - Allows you to write text to a file. Once done, use ctrl + d.
- `touch Asia/China/Country.txt` - Creates a file with the specified name. In this case, Country.txt

## Pagers
- `more new_file.txt` - View specified file in a scrollable manner. Space bar scrolls down a screen, Enter scrolls one line, 'b' scrolls back one page and / searches text.
- `less new_file.txt` - Allows you to view the contents of a file and navigate through the file. Up arrow scrolls the display up one line, down arrow scrolls the display down by one line, / searches text.

## LS (Long List)
- `ls -l` - Provides more details about the files and directories such as read, write and execute access, ownership, the last access time etc.
- `ls -a` - Displays all files including hidden files. 
- `ls -lt` - Displays files in order of creation (newest first).
- `ls -ltr` - Does the same as above but in reverse.