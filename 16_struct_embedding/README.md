# Embedded Struct & Struct Package Example

### Overview
This example of a Go program demonstrates the use of an embedded struct and a struct package. The embeded struct acts as a scaffold for a newly created struct to build upon and use its existing fields.

### Program manual
When run, the user will be presented with choices of action to choose from by inputting a number preceeding the action of choice. The choices include editing data, displaying data, and clearinng data for both student and teacher. Once and operation is performed, the user will be aske to confirm to continue performing another operation or exiting the program by inputting "Y" or "N" (case-insensitive).

Each step of choice input is validated, the program will keep asking the user until it gets the correct type of input. For the data input, no validation apart from checking if the input is a blank string. No data in any field of the structs in this program can be left blank.

### Code structure
The code is divided into main program code stored in the ```main.go``` file the code for functions stored in the ```opn.go``` (operations) file in the folder named ```opn```. The code in the ```opn.go``` file is another package that we have to import into the main program code.

The codes in this program also feature the use of pointers and struct embedding technique. The pointers are used to pass and returned through functions as arguments in order to directly modify the original structs without creating their copies, a technique to save on machine's memory space. The struct embedding technique is also used to create a struct called ```teacher``` based on a previously created struct called ```student```. This struct embedding technique allows to create a new struct (based on an exsiting one) quickly with some additional new fields.

### Program flow
The program flow is shown as the following numbered list:

1. The program presents choices of action to the user
2. The user chooses one of the choices which include editing, displaying, or erasing (clearing) the student and teacher's data
3. The program asks if the user wants to perform another action

NOTE: All choices and sub-choices are validate at the time of input. If not provided at first, the program will keep asking the user for a correct input.
