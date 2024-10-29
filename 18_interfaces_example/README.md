# Simple Go json Note-Taking

### Overview
This Go program demonstrates the use of structs and functions as struct methods stored in a saparated package to get string inputs from the user and save all the information in a json file with pre-determined json field names created by using struct tags.

### Program manual
When run, the program will ask the user to input a note title. After pressing the ```Enter``` key, the user will then be asked to input the note content. Both the title and content can be typed in freely as the program stores them as strings.

After inputting the note content, the program will show a message to confirm the note title and its content as given by the user. Then, the program will attempt to save the information as a ```json``` file. The saved ```json``` file contains three fields: ```title```, ```content```, and ```created_at```. The timestampt in the field ```created_at``` will be generated automatically by the program.

### Code structure
The codes of this projects are organized into the code for the main program (stored in the ```main.go`` file) and the codes for struct methods to display and save user inputs stored in the ```note.go``` file as a separate package.

### Program flow

1. The user inputs the note title as a string
2. The user inputs the note content as a string
3. The program takes those inputs to create a struct which stores the note title, content, and the timestamp at its creation
4. The program displays messages to confirm the title and the content from the inputs
5. The program displays a message to notify the user that it's saving the note  
6. The program attempts to save the inputs as a ```json``` file with json field names according to the struct tags given in the code
7. The program displays the message that it saves the file successfully


