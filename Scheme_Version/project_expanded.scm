#!/usr/local/dept/bin/mzscheme
#lang racket

(require racket/system)

;Returns the color code, based on what the user asks
(define Color (lambda (color)       ;color is the color string specified                                 
        (cond   [(string=? color "RESET")           "\033[0m"   ]
                [(string=? color "BLACK")           "\033[0;30m"]
                [(string=? color "RED")             "\033[0;31m"]
                [(string=? color "GREEN")           "\033[0;32m"]
                [(string=? color "YELLOW")          "\033[0;33m"]
                [(string=? color "BLUE")            "\033[0;34m"]
                [(string=? color "PURPLE")          "\033[0;35m"]
                [(string=? color "CYAN")            "\033[0;36m"]
                [(string=? color "WHITE")           "\033[0;37m"]
                [(string=? color "GRAY")            "\033[0;38m"]
                [(string=? color "DARK_GRAY")       "\033[1;30m"]
                [(string=? color "LIGHT_RED")       "\033[1;31m"]
                [(string=? color "LIGHT_GREEN")     "\033[1;32m"]
                [(string=? color "LIGHT_YELLOW")    "\033[1;33m"]
                [(string=? color "LIGHT_BLUE")      "\033[1;34m"]
                [(string=? color "LIGHT_PURPLE")    "\033[1;35m"]    
                [(string=? color "LIGHT_CYAN")      "\033[1;36m"]
                [(string=? color "LIGHT_WHITE")     "\033[1;37m"]
                [(string=? color "LIGHT_GRAY")      "\033[1;38m"]
                [else                               "\033[0m"   ])))

; Empties the screen for nice output
(define empty_screen                                  
	(lambda ()
		(multiDisplay (list "\x1b[2J" "\x1b[H"))))

; Exits the program if something goes wrong, or if user puts incorrect values
(define ABORT (lambda (message)       ; message is the message to notify the user what's wrong              
	(multiDisplay message)
	(exit)))

; Takes in a file and returns a line based on the unique number placed at its start
(define lineParser   
	(lambda (file n) ;file is the given file, n is the number identifier to be looked for
		(let ([line (string-split (read-line file) " ")])
			(if (not (eof-object? line))
				(if (not (string=? (car line) n))
					(lineParser file n)
					line)
				(ABORT (list (Color "LIGHT_RED") "Exception: Unfindable object\n" (Color "RESET") "Exiting...\n"))))))

;Returns the race data based on user input, though it doesn't organize them
(define getRace                            
	(lambda (file selection)       ;file in this case is the file that hold the race data, selection is the user selection
		(make-pseudo-random-generator)	
		(if (number? selection)
			(if (not (or (< selection 0) (> selection 47))) 
				(cond 	[(= selection 0)  (lineParser file  (number->string (+ 1  (random 9 (make-pseudo-random-generator)))))]
						[(= selection 46) (lineParser file (number->string  (+ 10 (random 36 (make-pseudo-random-generator)))))]
						[(= selection 47) (lineParser file (number->string  (+ 1  (random 45 (make-pseudo-random-generator)))))]
						[else (lineParser file (number->string selection))])
				(ABORT (list (Color "LIGHT_RED") "Incorrect numerical value entered.\n" (Color "RESET") "Exiting...\n")))
			(ABORT (list (Color "LIGHT_RED") "Non-numerical value entered.\n" (Color "RESET") "Exiting...\n")))))

; Makes one Stat Attribute: Rolls 3 dice and adds them up
(define makeStat            
	(lambda (a)           ; The number of dice to be rolled
		(sleep 1)
		(if (= a 0)
			(+ 1  (random 6 (make-pseudo-random-generator)))
			(+ (+ 1  (random 6 (make-pseudo-random-generator))) (makeStat (- a 1))))))

; Returns a list of Stats for the player attributes
(define makeStats          
	(lambda (a)      ;the number of stats to be calculated
		(if (= a 0)
			'()
			(cons (makeStat 2) (makeStats (- a 1))))))

; Calculates the Modifiers based on the stats
(define getModifiers      
	(lambda (lis)     ;lis is the list of Stats provided
		(if (null? lis)
			'()
			(cons (- (round (/ (car lis) 2)) 5) (getModifiers (cdr lis))))))

 ;Takes the race data based on user input, though it doesn't organize them
(define getClass            
	(lambda (file selection)       ;file is the file that has the class data, selection is the user selection
		(if (number? selection)
			(if (not (or (< selection 0) (> selection 14))) 
				(cond 	[(= selection 0)  (lineParser file  (number->string (+ 1  (random 14 (make-pseudo-random-generator)))))]
						[else (lineParser file (number->string selection))])
				(ABORT (list (Color "LIGHT_RED") "Incorrect numerical value entered.\n" (Color "RESET") "Exiting...\n")))
			(ABORT (list (Color "LIGHT_RED") "Non-numerical value entered.\n" (Color "RESET") "Exiting...\n")))))

;Part of displaySkill
(define displaySkill
	(lambda (lis a)
		(multiDisplay (list "[" a "] " (car lis) "\n"))
		(displaySkills (cdr lis) (+ a 1)))) ; Recursion for displaySkills

; Displays all of the available Skills for selection
(define displaySkills
	(lambda (lis a)  ;lis is the list of all available skills, a is the skill number 
		(if (not (null? lis))
			(displaySkill lis a)  ; Calls another function since more than 1 thing needs to be done
			(display ""))))

;Gets the length of a list
(define getLength
	(lambda (lis)  ; lis is the list
		(if (null? lis)
			0
			(+ (getLength (cdr lis)) 1))))

; Find a skill within a list and obrain it
(define obtainSkill
	(lambda (num lis in) 
		(if (number? in)
			(if (not (null? lis))
				(if (= in 0)
					(car lis)
					(obtainSkill (- num 1) (cdr lis) (- in 1)))
				(ABORT (list (Color "LIGHT_RED") "Item not found." (Color "RESET") "Exiting.\n")))
			(ABORT (list (Color "LIGHT_RED") "Non-integer value entered" (Color "RESET") "Exiting.\n")))))

; Removes/Pops an element of the list based on the given index
(define removeIndex 
	(lambda (index lis)  ; index is the given index, lis is the list
		(if (number? index)
			(if (not (null? lis))
				(if (= index 0)
					(cdr lis)
					(cons (car lis) (removeIndex (- index 1) (cdr lis))))
				(ABORT (list (Color "LIGHT_RED") "Exception: list index out of range" (Color "RESET") "Exiting.\n")))
			(ABORT (list (Color "LIGHT_RED") "Non-integer value entered" (Color "RESET") "Exiting.\n")))))

; displays the overall message for skill selection screen
(define displaySkillScreen 
	(lambda (lis num) ;lis is the skill list, num is the number of skills in that list
		(multiDisplay (list "You can select " num " of the following items (Type presicely):\n"))
		(displaySkills lis 0)))

;takes in user input once and applies it at 2 instances
(define constructSkills
	(lambda (user_input lis num)
		(cons (obtainSkill (getLength lis) lis user_input) (getSkills (- num 1) (removeIndex user_input lis)))))

; makes a list of the skills the user selects
(define getSkills
	(lambda (num lis) ; num is the amount of skills a user can choose, lis is the total list of skills
		(empty_screen)
		(if (not (= num 0))
			(displaySkillScreen lis num)
			(display ""))
		(if (= num 0)
			'()
			(constructSkills (readInput "\nPlease make your selection: ") lis num)))) ;using this because we can't store the user input without crashing

;prompts the user for input and returns the input, when I realized I could call it prompt it was too late
(define readInput
	(lambda (message) ;message is the prompt
		(display message)
		(read)))

; displays a list of things (more flexible than (display (string-append)), since it can take non-strings)
(define multiDisplay 
    (lambda (lis)
	    (display (car lis))
        (if (not (null? (cdr lis)))
            (multiDisplay (cdr lis))
            (display ""))))

;puts color where it is requested
(define readColor
	(lambda (file str)
		(let ([line (read-line file)])
			(if (not (eof-object? line))
				(let ([wanted_color (car (string-split line "|_|"))] [change_line (cadr (string-split line "|_|"))])
					(readColor file (string-replace str change_line (string-append (Color wanted_color) change_line (Color "RESET")))))
				str))))

;reads the lines of a file one by one for color processing
(define readLines 
	(lambda (file str)
    	(let ([line (read-line file)])
        	(if (not (eof-object? line))
            	(if (not (string=? line "(#COLOR)"))
					(readLines file (string-append str line "\n"))
					(readColor file str))
            	str))))

; reads the lines of a file
(define readFile 
	(lambda (n)
	    (let ((in (open-input-file (string-append "readfiles/" n ".txt"))))
	        (display (readLines in ""))
	        (close-input-port in))))

; like multiDisplay, but puts the output in a file
(define multiOutput
	(lambda (output lis)
		(display (car lis) output)
		(if (not (null? (cdr lis)))
			(multiOutput output (cdr lis))
			(display ""))))

;Writes all of the Attributess in the output file
(define writeAttrs
	(lambda (output scorelis modlis attrlis)
		(multiOutput output (list (car attrlis) "\t|" (car scorelis) "\t\t|" (car modlis) "\n"))
		(if (and (and (not (null? (cdr attrlis))) (not (null? (cdr scorelis)))) (not (null? (cdr modlis))))
			(writeAttrs output (cdr scorelis) (cdr modlis) (cdr attrlis))
			(display ""))))

; if we see "St" get "Strength" etc.
(define elongateAttrShorthand
	(lambda (lis)
		(if (not (null? lis))
			(cond 	[(string=? (car lis) "St")	(cons "Strength" (elongateAttrShorthand (cdr lis)))]
					[(string=? (car lis) "De")	(cons "Dexterity" (elongateAttrShorthand (cdr lis)))]
					[(string=? (car lis) "Co")	(cons "Constitution" (elongateAttrShorthand (cdr lis)))]
					[(string=? (car lis) "Wi")	(cons "Wisdom" (elongateAttrShorthand (cdr lis)))]
					[(string=? (car lis) "In")	(cons "Intelligence" (elongateAttrShorthand (cdr lis)))]
					[(string=? (car lis) "Ch")	(cons "Charisma" (elongateAttrShorthand (cdr lis)))]
					[else (ABORT (list (Color "LIGHT_RED") "Non-Attribute detected" (Color "RESET") "Exiting.\n"))])
			'())))

; Used by print saving Throws
(define printMiddleThrow
	(lambda (output lis)
		(multiOutput output (list (car lis) ", "))
		(printSavingThrows output (cdr lis))))

; Prints all of the saving Throws for the character
(define printSavingThrows
	(lambda (output lis)
		(cond 	[(not (null? (cdr lis))) (printMiddleThrow output lis)] ; we need more than one thing done here, so call function
				[(not (null? lis)) (display (car lis) output)]
				[else (display "")])))

; used by print skills
(define printSkill
	(lambda (output lis)
		(multiOutput output (list "\t" (car lis) "\n"))
		(printSkills output (cdr lis))))

; Prints all of the character's skills
(define printSkills
	(lambda (output lis)
		(cond 	[(not (null? lis)) (printSkill output lis)]))) ; we need more than one thing done here

; Writes the data we collected into a file
(define writeCharacterFile
	(lambda (output)
		(multiOutput output (list "Name: " characterName "\t\tRace: " (cadr race_stuff)  "\t\tClass: " (cadr class_stuff)))
		(multiOutput output (list "\nSize: " (car(cddddr race_stuff)) "\t\tHP: " (+ (string->number (car (string-split (caddr class_stuff) "+"))) (caddr mods)) "\t\t\tSpeed: " (cadr(cddddr race_stuff))))
		(multiOutput output (list "\nLevel: 1\n\nAttributes" "\n---------------------------------" "\nName            |Score  |Modifier" "\n----------------+-------+--------\n"))
		(writeAttrs  output stats mods (list "Strength    " "Dexterity   " "Constitution" "Intelligence" "Wisdom      " "Charisma    "))
		(display "\nSaving Throws: " out)
		(printSavingThrows output (elongateAttrShorthand (string-split (car (cdddr class_stuff)) ",")))
		(multiOutput output (list "\nLanguages: " (string-replace (cadddr race_stuff) "." " ")))
		(multiOutput output (list "\nArmor Proficiences: " (string-replace (car (cddddr class_stuff)) "." " ") "\nWeapon Proficiencies: " (string-replace (cadr (cddddr class_stuff)) "." " ") "\nTool Proficiencies: " (string-replace (caddr (cddddr class_stuff)) "." " ")))
		(multiOutput output (list "\nRace Extras: " (string-replace (cadr (cddddr race_stuff)) "." " ") "\nClass Feats: " (string-replace (car(cddddr (cddddr class_stuff))) "." " ")))
		(multiOutput output (list "\nItems: " (string-replace (cadddr (cddddr class_stuff)) "." " ") "\nBoosted Skills:\n"))
		(printSkills output skills)
		(multiDisplay (list "\nYour character sheet awaits you in the file " characterName "_sheet.txt!\n"))))
	
(empty_screen)
(readFile "welcome")
(define characterName (readInput "\nPlease give your name: "))

(empty_screen)
(readFile "race_selector")
(define race_attributor (open-input-file "readfiles/race_attributor.txt"))
(define race_stuff (getRace race_attributor (readInput "\nPlease select the number that represents your choice: ")))
(close-input-port race_attributor)

(display "Loading Scores. Please Wait. \n")
(display "")
(define stats (makeStats 6))
(define mods (getModifiers stats))

(empty_screen)
(readFile "class_selector")
(define class_attributor (open-input-file "readfiles/class_attributor.txt"))
(define class_stuff (getClass class_attributor (readInput "\nPlease select the number that represents your choice: ")))
(close-input-port class_attributor)
(empty_screen)
(define skills (getSkills (string->number (car (string-split (car(cdr (cddddr (cddddr class_stuff)))) ",."))) (cdr (string-split (car(cdr (cddddr (cddddr class_stuff)))) ",."))))

(define out (open-output-file (string-append (symbol->string characterName) "_sheet.txt")))
(writeCharacterFile out)
(close-output-port out)