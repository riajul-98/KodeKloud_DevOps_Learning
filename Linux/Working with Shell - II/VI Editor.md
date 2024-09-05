# VI Editor
VI Editor is a console based text editor in Linux which is one of the most popular amongst other text editors. To get started use the vi command followed by the file name e.g. `vi index.html`

## VI Editor Modes
VI Editor has three modes of operations which are:
- Command Mode: Allows you to make commands such as delete line, copy and paste.
- Insert Mode: Allows you to type text. To go into insert mode, you press 'i'. To switch back to command mode, press 'esc' key.
- Last Line Mode: In this mode, you can choose to save the file, discard and quit or save and quit. You would press ':' to switch from command mode to last line mode.

## Commmand mode
The command mode will be the first mode you are in when you open up or create a file using vi. This allows you to highlight, copy and paste text. You can move the cursor using the arrow keys.

## Commands
- `x`: Deletes a character.
- `ZZ`: Saves the file.
- `dd`: Deletes an entire line. You can delete multiple lines by typing the number of lines you wish to delete followed by dd.
- `u`: undo change.
- `ctrl + r`: Redo change after an undo.
- `yy`: Copies a line.
- `p`: Pastes copied content.
- `:w`: Saves changes.
- `:q`: Quits changes.
- `:wq`: Saves and exits VI editor.
- `/of`: Finds all occurances of the word specified, in this case "of". Press 'n' to go to next occurance and 'N' for previous occurance.

## Insert mode
Insert mode is used to write data to the file. The text editor is very similar in functionality to graphical text editors. 
