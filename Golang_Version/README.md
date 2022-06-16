Regarding the project, it is to be run as **go run project.go** (more convinient than the Python command). The project is a terminal-based character creator for the Tabletop Role Playing Game, Dungeons and Dragons, 5th edition.
The program has a few implemented time delays (mostly done for effect and to allow the reader some time to read messages before they disappear).
There will be an input prompt each time one needs to give their input.
This is given in the form of a list (when there are multiple options listed, there is an area with ```[x]```, where the ] and the [ symbols enclose x, which represent an option to be inserted)
Each output must be provided when the time is right, or there might be a bit of confusion if a user is too hasty.
The option ```(ENTER to continue, N to reset, Q to quit)``` expects any type of input, where Q or Quit (not case-sensitive) is a command to quit the program altogether and N or No (not case-sensitive) is a command that allows the user to reset their previous input. ENTER is linient and assumes any other form of input is for ENTER (even black input).
At the end, the program creates a file with a pre-specified string "_sheet.txt" at the end of its title, with the strings before the pre-specified string being dependent on the user's input (the name of the character they gave).
So a brief rundown of what the program does in linear order:
- It greets the user
- It waits for the user to input the name of the character they want to create. (The name is included in the final txt output file)
- It lets the user choose a race (like a species for instance Orc, Dwarf or Human) (it even allows for random selection)
- It rolls 5 "dice" (random number generation from 1 to 6) and adds the 3 highest values, to give an Attribute Score
    - This happens 6 times in total as there are 6 attribures
    - The player is then given the list of 6 scores to assign to each of the 6 attributes
        - The last Attribute score is assigned automatically for user convinience.
    - Based on the scores for each Attribute, the program Sets the Attribute Modifiers (important for Attribute rolls)
- It allows the user to select their class (including random selection)
    - Prompts the user to select their equipment based on their class
    - Prompts the user to select their skills based on their class
- It formats the information neatly in to a ```<name>_sheet.txt``` file, using a character sheet template, and lets the user know that their file is ready.

* New changes:
- I couldn't find operations such as pop from list or sort a list, so I implemented:
    - ```popstr and popint``` that pop value from an array of strings and integers respectively using a given index. The pop command returns a new array and the popped value
    - ```indexint``` which finds the index of a value in a string array.
    - ```listmax``` that takes a list of integers and finds the maximum value
    - ```reverse_bubble_sort```, I needed a sorted array with descending order, so I used a bubble short and reversed the procedure (it normally gives an ascending list, I think)
    - ```pad_right```. All things I found for making a string at least x-length long, involved fmt.Println, which I didn't want, since I had to write everything in a file. This is my replacement for .ljust() in Python
- I realized that the Python version had all the txt files used for the successful running of the program in the same directory as the program, creating chaos. So I decided to put the txt files that are read into a subdirectory for proper organization.