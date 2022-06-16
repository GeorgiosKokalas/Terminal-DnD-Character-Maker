Regarding the project, it can be run either as **./project.scm** or **./project_expanded.scm**. Both files contain the same code, but **project_expanded** also contains comments and a slightly different format.
The project is a terminal-based character creator for the Tabletop Role Playing Game, Dungeons and Dragons, 5th edition.
The program had to be rebuilt from the ground up, so some features had to be removed or simplified. For the most part it functions pretty similar to the previous versions of the project as far as the user is concerned.
The program has a few implemented time delays (mostly done for effect and to allow the reader some time to read messages before they disappear).
There will be an input prompt each time one needs to give their input.
This is given in the form of a list (when there are multiple options listed, there is an area with ```[x]```, where the ] and the [ symbols enclose x, which represent an option to be inserted)
Each output must be provided when the time is right, or there might be a bit of confusion if a user is too hasty.
If the user enters the wrong input, the program decides to end (implemented, not just an Error)
At the end, the program creates a file with a pre-specified string "_sheet.txt" at the end of its title, with the strings before the pre-specified string being dependent on the user's input (the name of the character they gave).
So a brief rundown of what the program does in linear order:
- It greets the user
- It waits for the user to input the name of the character they want to create. (The name is included in the final txt output file)
- It lets the user choose a race (like a species for instance Orc, Dwarf or Human) (it even allows for random selection)
- It gets the Attribute:
    - Instead of rolling 5 dice and adding the highest 3 for each attribute, the program simply rolls 3 dice and adds them.
    - Instead of allowing the user to assign their attributes, the attributes needed to be pre-assigned
    - Based on the scores for each Attribute, the program Sets the Attribute Modifiers (important for Attribute rolls)
- It allows the user to select their class (including random selection)
    - No longer prompts the user to select their equipment based on their class
    - Prompts the user to select their skills based on their class
- It formats the information neatly in to a ```<name>_sheet.txt``` file, using a character sheet template, and lets the user know that their file is ready.
