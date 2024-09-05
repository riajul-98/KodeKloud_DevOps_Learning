# Searching for Files and Directories
There are many ways of searching for a file or directory using the command line. These include:

- `locate City.txt`: Returns all paths which contain the matching pattern. You would need to install mlocate.db and run the command `updatedb`
- `find /home/michael -name City.txt`: Returns the file specified under the path specified. This will also look in directories located in the path.
- `grep <word> <file_name>`: Searches for the specified word in the specified file. This is case sensitive
- `grep -i <word> <file_name>`: -i makes the search case insensitive.
- `grep -r <word> <directory_path>`: Searches for the given word recursively within a directory.
- `grep -v <word> <file_name>`: Lists all lines that do not have the specified word.
- `grep -w <word> <file_name>`: Searches only for the specified word and nothing more. e.g. if the word was given as "exam", it would only return exam and not "example".
- `grep -vw <word> <file_name>`: Searches for words that contain the specified word but is not exactly that word.
- `grep -A1 <word> <file_name>`: Prints one line after matching the pattern. You can change this to A2, A3.... to print the specified number of lines. It will print the match as well as the line after.
- `grep -B1 <word> <file_name>`: Prints one line before matching the pattern. You can change this to B2, B3.... to print the specified number of lines. It will print the line before as well as the line that matched.