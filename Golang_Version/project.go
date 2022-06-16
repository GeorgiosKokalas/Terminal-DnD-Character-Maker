package main

import("fmt"
	"strings"
	"io/ioutil"
	"os"
	"time"
	"strconv"
	"math/rand"
)

// GENERIC FUNCTIONS
//------------------
// read_from:
//   - Input: the string containing the name of a file 
//   - Output: an array of strings containing the color-modified contents of the file
//   - Takes in a file and reads it, then assigns the colors to the text (putting them in the txt file didn't work), and then splits the file string into multiple lesser strings
func read_from(filename string) []string{
	file_read, _ := ioutil.ReadFile("readfiles/"+filename)
	file_data := string(file_read)
	if strings.Contains(file_data, "(#COLOR)"){
		textie := strings.Split(file_data, "(#COLOR)")
		colorchange :=strings.Split(textie[1], "\n")
		for _, i := range colorchange{  
			if len(i) < 2 {continue}
			splitter := strings.Split(i, "|_|")
			textie[0] = color.Give_Colors(textie[0],splitter[1], splitter[0])
		}
		file_data = textie[0]
	}
	file_lines := strings.Split(file_data,"*")
	return file_lines
}

// reader 
//   - Input: the string containing the name of a file, an integer containing the second interval to occur between each line 
//   - Output: None
//   - Takes in a file, gives it to read_from(), then prints the contents in a delayed manner based on the integer given
func reader(file string, interval int){   
	file_lines := read_from(file)
	for i := range file_lines{              
		fmt.Println(file_lines[i])
		switch interval {
		case 1:
			time.Sleep(1 * time.Second)
		case 2:
			time.Sleep(2 * time.Second)
		case 3:
			time.Sleep(3 * time.Second)
		default:
			time.Sleep(0 * time.Second)
		}
	}
}

// get_input:
//   - Input: 2 strings: the message to be printed before the input, and the message to be printed after the input asking for confirmation
//   - Output: 1 string: the input of the user
//   - Prints the input prompt, waits for user input, then asks for confirmation. If the user gives n or no, it resets and asks again, if they use q or quit, it exits the program. Anything else and it continues
func get_input(prompt, reassurance string) string {
	unsure := true 
	var receiver string
	for unsure {
		if len(prompt) > 0{
			fmt.Printf(prompt)
			fmt.Scanln(&receiver)
			fmt.Println()
		}
		fmt.Printf(reassurance+" (ENTER to continue, N to reset, Q to quit) ")
		var insurance string
		fmt.Scanln(&insurance)
		insurance = strings.ToLower(insurance)
		if insurance == "q" || insurance == "quit" {
			fmt.Println("Farewell for now. Let us hope we meet again.")
			os.Exit(0)
		} else if insurance == "n" || insurance == "no" {
			time.Sleep(1 * time.Second)
			unsure = true
		} else {
			unsure = false
		}
	}
	return receiver
}

// empty_screen
//   - Input: None
//   - Output: None
//   - Literally just empties the screen
func empty_screen(){
	fmt.Println("\x1b[2J")
	fmt.Println("\x1b[H")
}

// popstr
//   - Input: 1 string array: the array to pop the element from, 1 int: the index to pop from
//   - Output: 1 string array: the array minus the popped item, 1 string: the popped item
//   - Takes a string array and an index to know which element to remove, then returns an array that is identical to the initial one minus the popped string that is returned seperately
func popstr(list []string, index int) ([]string, string){
	if (index >= len(list)) {index = len(list)-1}
	if (index < 0 && index >= -1 * len(list)) {
		index += len(list)
	} else if (index < -1 * len(list)) { index = 0 }
	new_slice := list[:index]
	remainder := list[index+1:]
	ret := list[index]
	for _, i := range remainder{
		new_slice = append(new_slice, i)
	}
	return new_slice, ret
}

// popint
//   - Input: 1 integer array: the array to pop the element from, 1 int: the index to pop from
//   - Output: 1 integer array: the array minus the popped item, 1 int: the popped item
//   - Takes an integer array and an index to know which element to remove, then returns an array that is identical to the initial one minus the popped integer that is returned seperately
func popint(list []int, index int) ([]int, int) {
	if (index >= len(list)) {index = len(list)-1}
	if (index < 0 && index >= -1 * len(list)) {
		index += len(list)
	} else if (index < -1 * len(list)) { index = 0 }
	
	new_slice := list[:index]
	remainder := list[index+1:]
	ret := list[index]
	for _, i := range remainder{
		new_slice = append(new_slice, i)
	}
	return new_slice, ret
}

// indexint
//   - Input: 1 integer array: the array to inspect, 1 int: the value to locate
//   - Output: 1 integer: the index of the given value
//   - Takes an integer array and a value to locate, then returns the index having the first instance of that value
func indexint(list []int, value int) int{
	var index int = -9223372036854775808
	for pos, i := range list{
		if i == value {
			index = pos
			break
		}
	}
	return index
}

// listmax:
//   - Input: 1 integer array: the array to inspect
//   - Output: 1 integer: maximum value in that array
//   - Takes in an integer array and finds the highest value in it
func listmax(list []int) int {
	var output int = -9223372036854775808
	for _, i := range list {
		if i > output {output = i}
	}
	return output
}

// reverse_bubble_sort
//   - Input: 1 integer array: the array to sort
//   - Output: 1 integer array: the sorted array
//	 - Uses bubble sorting to sort the values of the array in descendign order
func reverse_bubble_sort(list []int) []int{
	for bubbles := true; bubbles; {
		bubbles = false
		for i:=0; i < len(list)-1; i++{
			if (list[i] < list[i+1]) {
				bubbles = true
				temp := list[i]
				list[i] = list[i+1]
				list[i+1] = temp
			}
		}
	}
	return list
}

// pad_right
//   - Input: 1 string: the string to modify, 1 int: how long the string should be
//   - Output: 1 string: the expanded string
//	 - Takes in a string and if the integer is higher than the length of the string, the string is given whitespaces at the end until it reaches that length
func pad_right(str string, length int) string{
	if length < len(str) { length = len(str) }
	for len(str) < length {
		str += " "
	}
	return str
}

//DEDICATED TO COLOR STRUCT
//-------------------------
type Color struct {color map[string]string}
var color Color
func (c *Color) Assign_Colors(){                // Specify all colors I could find
	c.color = make(map[string]string)
	c.color["RESET"]           = "\033[0m"
	c.color["BLACK"]           = "\033[0;30m"
	c.color["RED"]             = "\033[0;31m"
	c.color["GREEN"]           = "\033[0;32m"
	c.color["YELLOW"]          = "\033[0;33m"
	c.color["BLUE"]            = "\033[0;34m"
	c.color["PURPLE"]          = "\033[0;35m"
	c.color["CYAN"]            = "\033[0;36m"
	c.color["WHITE"]           = "\033[0;37m"
	c.color["GRAY"]            = "\033[0;38m"
	c.color["DARK_GRAY"]       = "\033[1;30m"
	c.color["LIGHT_RED"]       = "\033[1;31m"
	c.color["LIGHT_GREEN"]     = "\033[1;32m"
	c.color["LIGHT_YELLOW"]    = "\033[1;33m"
	c.color["LIGHT_BLUE"]      = "\033[1;34m"
	c.color["LIGHT_PURPLE"]    = "\033[1;35m"
	c.color["LIGHT_CYAN"]      = "\033[1;36m"
	c.color["LIGHT_WHITE"]     = "\033[1;37m"
	c.color["LIGHT_GRAY"]      = "\033[1;38m"
}

// Give_Colors:  (part of Color)
//   - Input: 3 strings: the text to edit, what part of the text to edit, what color to add
//   - Output: 1 string: the string with the color codes in it
//   - Takes in a string and places colors in the substring that is specified
func (c *Color) Give_Colors(text, old, new_color string) string{
	text = strings.ReplaceAll(text, old,c.color[new_color]+old+c.color["RESET"])
	return text
}

//DEDICATED TO THE PLAYER AND THE ATTRIBUTE STRUCTS
//-------------------------------------------------
type Attribute struct {
	shortname, fullname string
	value, modifier int
} 

type Player struct{
	attrs map[string]*Attribute
	name, race, class string
    size, languages, extras string
	speed, HP int
    savingT, armorP, weaponP, toolP, equipment, feats string
	skills []string	
}

// race_parser:  (part of Player)
//   - Input: 1 int: the selection of the user
//   - Output: 1 string array: the race that is specified by the player
//   - Takes in the player's choice as an integer and finds the appropriate race to take data from and returns it in the form of a string array
func (player *Player) race_parser(race_int int) []string{
	race_file, _ := ioutil.ReadFile("readfiles/race_attributor.txt")
	race_lines := strings.Split(string(race_file), "\n")
	rand.Seed(time.Now().UnixNano())
	switch race_int{
	case 0:
		fmt.Println("Your race will be selected randomly from the Base Game Races")
		race_int = rand.Intn(9) + 1
	case 46:
		fmt.Println("Your race will be selected randomly from the Expansion Races")
		race_int = rand.Intn(36) + 10
	case 47:
		fmt.Println("Your race will be selected randomly")
		race_int = rand.Intn(45) + 1		
	}
	var race_line []string
	for _, i := range race_lines{
		if strings.Contains(i, "#"){ continue }
		judge := strings.Fields(i)
		race_int_getter, _ := strconv.Atoi(judge[0])
		if race_int_getter == race_int{
			race_line = strings.Fields(i)
		}
	}
	return race_line
}

// select_race:  (part of Player)
//   - Input: 1 int: the selection of the user
//   - Output: None
//   - Takes in the player's choice, gives it to race_parser() to get the string array to read data from and assigns the data appropriately
func (player *Player) select_race(race_int int) {
	parsed_race := player.race_parser(race_int)
	player.race = parsed_race[1]
	if strings.Contains(player.race, ".") {
		player.race = strings.ReplaceAll(player.race, ".", " ")
	}
	selected_attributes := []string{"St", "De", "Co", "In", "Wi", "Ch"}
	var at_boost_str []string
	if strings.Contains(parsed_race[2], ",") {
		at_boost_str = strings.Split(parsed_race[2], ",")
	} else {
		at_boost_str = append(at_boost_str, parsed_race[2])
	}
	for _, i := range at_boost_str {
		selected_attribute := string(i[0:2])
		if selected_attribute == "Ra" {
			rand.Seed(time.Now().UnixNano())
			selected_attributes, selected_attribute = popstr(selected_attributes, rand.Intn(len(selected_attributes)))
		}
		increase_value, _ := strconv.Atoi(string(i[3]))
		player.attrs[selected_attribute].value += increase_value
	}
	player.size = parsed_race[4]
    player.speed,_ = strconv.Atoi(parsed_race[5])
    player.extras = strings.ReplaceAll(parsed_race[6],".", " ")
    player.languages = strings.ReplaceAll(parsed_race[3],".", " ")
}

// attribute_assigner:  (part of Player)
//   - Input: None
//   - Output: None
//   - Repeat 6 times (roll 5 dice, add the 3 maximum values and put the result in a list), sort the array, allow the user to assign each attribute, calculate the modifier from the results, then assign to appropriate place
func (player *Player) attribute_assigner(){
	var scores []int
	for i := 0; i<6; i++{
		var score int = 0
		var dice []int
		for c := 0; c<6; c++{
			rand.Seed(time.Now().UnixNano())
			dice = append(dice, rand.Intn(6)+1)
		}
		for u := 0; u<3; u++ {
			var local_maximum int
			dice, local_maximum = popint(dice, indexint(dice, listmax(dice)))
			score += local_maximum
		}
		scores = append(scores, score)
	}
	scores = reverse_bubble_sort(scores)
	var attribute_shorts []string = []string{"St", "De", "Co", "In", "Wi", "Ch"}

	for i:=0; i<5; i++{
        unsure := true
        for unsure{
            empty_screen()
            fmt.Printf("Remaining values: ")
            for _,i := range scores{
				fmt.Printf(strconv.Itoa(i))
				fmt.Printf("  ")
			}
            fmt.Println("\nPlease assign this:", scores[0])
            for _, p := range attribute_shorts{
				fmt.Println("[" + p + "] " + player.attrs[p].fullname)
			}
            selected_attr := get_input("Your choice: ", "Are you sure?")
			is_there := false
			for i_pos, i := range attribute_shorts{
				if i == selected_attr {
					is_there = true
					attribute_shorts, _ = popstr(attribute_shorts, i_pos)
					break
				}
			}
            if (is_there){
				var cur_score int
				scores, cur_score = popint(scores, 0)
                player.attrs[selected_attr].value += cur_score
                unsure = false
			} else {
                print("\033[1;31mNot desired value entered\033[0m")
                time.Sleep(2 * time.Second)
			}
		}
	}
	fmt.Println("Assigned the remaining value", scores[0], "to", player.attrs[attribute_shorts[0]].fullname)
	player.attrs[attribute_shorts[0]].value += scores[0]
	time.Sleep(2 * time.Second)
	empty_screen()
	fmt.Println("Calculating Modifiers:\n"+color.color["GREEN"]+"Attribute\tScore\tModifier"+color.color["RESET"])
	for _, i := range []string{"St", "De", "Co", "In", "Wi", "Ch"} {
		player.attrs[i].modifier = (player.attrs[i].value / 2) - 5
		fmt.Println(player.attrs[i].fullname,"\t",player.attrs[i].value,"\t",player.attrs[i].modifier)
	}
	get_input("","")
}

// class_parser:  (part of Player)
//   - Input: 1 int: the selection of the user
//   - Output: 1 string array: the class that is specified by the player
//   - Takes in the player's choice as an integer and finds the appropriate class to take data from and returns it in the form of a string array
func (player *Player) class_parser(class_int int) []string {
	class_file, _ := ioutil.ReadFile("readfiles/class_attributor.txt")
	class_lines := strings.Split(string(class_file), "\n")
	var announce bool = false
	rand.Seed(time.Now().UnixNano())
	switch class_int{
	case 0:
		class_int = rand.Intn(13) + 1
		announce = true
	}
	var class_line []string
	for _, i := range class_lines{
		if strings.Contains(i, "#"){ continue }
		judge := strings.Fields(i)
		class_int_getter, _ := strconv.Atoi(judge[0])
		if class_int_getter == class_int{
			class_line = strings.Fields(i)
		}
	}
	empty_screen()
	if announce { 
		fmt.Println("Random Selection made you a", strings.ReplaceAll(class_line[1],".", " ")) 
		get_input("", "")
	}
	return class_line
}

// select_class:  (part of Player)
//   - Input: 1 int: the selection of the user
//   - Output: None
//   - Takes in the player's choice, gives it to class_parser() to get the string array to read data from and assigns the data appropriately
func (player *Player) select_class(class_int int) {
	parsed_class := player.class_parser(class_int)
	player.class = strings.ReplaceAll(parsed_class[1],".", " ")

	HPsep := strings.Split(parsed_class[2], "+")
	baseHP, _ := strconv.Atoi(HPsep[0])
	player.HP = baseHP + player.attrs[HPsep[1][:2]].modifier

	for _, i := range strings.Split(parsed_class[3], ",") { player.savingT += player.attrs[i].fullname + ", " }
	player.savingT = player.savingT[:len(player.savingT)-2]
	
	player.armorP  = strings.ReplaceAll(parsed_class[4],"."," ")
	player.weaponP = strings.ReplaceAll(parsed_class[5],"."," ")
	player.toolP   = strings.ReplaceAll(parsed_class[6],"."," ")
	player.feats   = strings.ReplaceAll(parsed_class[8],"."," ")
	
	var unsure bool
	for _,i := range strings.Split(strings.ReplaceAll(parsed_class[7],"."," "), ", "){
		if strings.Contains(i, "/"){
			unsure = true
			for unsure {
				empty_screen()
				choice := strings.Split(i, "/ ")
				fmt.Println("Choose which piece of equipment you desire")
				for o:= 0; o < len(choice); o++ { fmt.Println("["+strconv.Itoa(o)+"] " + choice[o]) }
				choice_input := get_input("Your choice: ", "Are you sure this is the item you want?")
				choice_int, choice_int_err := strconv.Atoi(choice_input)	
				if choice_int_err != nil{
					fmt.Println(color.color["LIGHT_RED"] + "Value entered not of type int. Please try again." + color.color["RESET"])
					time.Sleep(2 * time.Second)
				} else if choice_int_err != nil && (choice_int < 0 || choice_int >= len(choice)){
					fmt.Println(color.color["LIGHT_RED"] + "Integer entered not within desired bounds. Please try again." + color.color["RESET"])
					time.Sleep(2 * time.Second)
				} else{
					i = choice[choice_int]
					unsure = false
				}
			}
		}
		player.equipment += i + ", "
	}
	player.equipment = player.equipment[:len(player.equipment)-2]

	var skill_list []string = strings.Split(strings.ReplaceAll(parsed_class[9],"."," "), ", ")[1:]
	var skills_to_choose int
	skills_to_choose, _ = strconv.Atoi(string(parsed_class[9][0]))

	for i := 0; i < skills_to_choose; i++{
		unsure = true
		for unsure {
			empty_screen()
			fmt.Println("As a " + player.class + ", you have access to " + string(parsed_class[9][0]) + " Skill proficiencies from those below.")
			for o_pos, o := range skill_list { fmt.Println("[" + strconv.Itoa(o_pos) + "]\t" + o) }
			fmt.Println("Choices remaining:", skills_to_choose - i)
			choice_input := get_input("Your choice: ", "Are you sure this is the skill you want?")
			choice_int, choice_int_err := strconv.Atoi(choice_input)
			if choice_int_err != nil{
				fmt.Println(color.color["LIGHT_RED"] + "Value entered not of type int. Please try again." + color.color["RESET"])
				time.Sleep(2 * time.Second)
			} else if choice_int_err != nil && (choice_int < 0 || choice_int >= len(skill_list)){
				fmt.Println(color.color["LIGHT_RED"] + "Integer entered not within desired bounds. Please try again." + color.color["RESET"])
				time.Sleep(2 * time.Second)
			} else{
				var skill_selected string
				skill_list, skill_selected = popstr(skill_list, choice_int)
				player.skills = append(player.skills, skill_selected)
				unsure = false
			}
		}
	}
}

// Character_Builder:  (part of Player)
//   - Input: None
//   - Output: None
//   - Asks the player for information to build a DnD character
func (player *Player) Character_Builder() {
	reader("welcome.txt", 1)
	get_input("", "")
	empty_screen()

	player.name = get_input("Each character has a name, what will yours be? ", "Are you sure this is the name you desire?")
	
	var unsure bool = true
	var race_int int
	var race_int_err error
	for unsure{
		empty_screen()
		reader("race_selector.txt",0)
		race_int, race_int_err = strconv.Atoi(get_input("Please Select a number from 0 to 47, representing a race: ", "Are you sure this is your wanted race?"))
		if race_int_err != nil {
			fmt.Println(color.color["LIGHT_RED"] + "You didn't give an integer, please try again" + color.color["RESET"])
			time.Sleep(2 * time.Second)
		} else if race_int_err == nil && (race_int < 0 || race_int > 47){
			fmt.Println(color.color["LIGHT_RED"] + "Inserted number not within desirable values. Please try again." + color.color["RESET"])
			time.Sleep(2 * time.Second)
		} else {
			unsure = false
		}
	}
	player.select_race(race_int)
	fmt.Println("As a", player.race, ", you are",player.size,"size, you know ", player.languages,"and have a speed of",player.speed,", on top of having",player.extras,"!")
	get_input("", "")

	empty_screen()
    fmt.Println("Time to Calculate your Attributes and their Modifiers!\nFor this, 5 random dice will be rolled and you will get the maximum 3 for each Attribute.")
    time.Sleep(5 * time.Second)
	player.attribute_assigner()

	unsure = true
	var class_int int
	for unsure{
		empty_screen()
        reader("class_selector.txt",0)
        class_int_str := get_input("Please Select a number from 0 to 14, representing a race: ", "Are you sure this is your wanted class?")
		var class_int_err error
		class_int, class_int_err = strconv.Atoi(class_int_str)
		if class_int_err != nil {
			fmt.Println(color.color["LIGHT_RED"] + "No number detected. Please try again." + color.color["RESET"])
			time.Sleep(2 * time.Second)
		} else if class_int_err == nil && (class_int < 0 || class_int > 14) {
			fmt.Println(color.color["LIGHT_RED"] + "Inserted number not within desirable values. Please try again." + color.color["RESET"])
			time.Sleep(2 * time.Second)
		} else {
			unsure =false
		}
	}
	player.select_class(class_int)
}

// Sheet_Builder:  (part of Player)
//   - Input: None
//   - Output: None
//   - Formats the data given by the user into a readable tt file with the DnD character name on it
func (player *Player) Sheet_Builder(){
	sheet_b, _ := ioutil.ReadFile("readfiles/character_template.txt")
	sheet := string(sheet_b)

	sheet = strings.ReplaceAll(sheet, "Name:  Race:  Class:", pad_right("Name:" + player.name, 25) + pad_right("Race:" + player.race, 25) + "Class:" + player.class)
	sheet = strings.ReplaceAll(sheet, "Size:    HP:   Speed:", pad_right("Size:" + player.size, 25) + pad_right("HP:" + strconv.Itoa(player.HP),25) + "Speed:" + strconv.Itoa(player.speed))

	for _, i := range player.attrs { sheet = strings.ReplaceAll(sheet, i.fullname, pad_right(i.fullname, 13) + "| " + pad_right(strconv.Itoa(i.value), 5) + "| " + pad_right(strconv.Itoa(i.modifier), 4)) }
	sheet = strings.ReplaceAll(sheet, "Saving Throws:", "Saving Throws: " + player.savingT)

    sheet = strings.ReplaceAll(sheet, "Languages:", "Languages: " + player.languages)
    sheet = strings.ReplaceAll(sheet, "Armor Proficiencies: ", "Armor Proficiencies: " + player.armorP)
    sheet = strings.ReplaceAll(sheet, "Weapon Proficiencies: ", "Weapon Proficiencies: " + player.weaponP)
    sheet = strings.ReplaceAll(sheet, "Tool Proficiencies: ", "Tool Proficiencies: " + player.toolP)
    sheet = strings.ReplaceAll(sheet, "Extras:", "Extras: " + player.extras)
    sheet = strings.ReplaceAll(sheet, "Feats: ", "Feats: " + player.feats)
    sheet = strings.ReplaceAll(sheet, "Items: ", "Items: " + player.equipment)

	for _, i := range []string{"Athletics", "Acrobatics", "Sleight of Hand", "Stealth", "Arcana", "History", "Investigation", "Nature", "Religion", "Animal Handling", "Insight", "Medicine", "Perception", "Survival", "Deception", "Intimidation", "Performance", "Persuasion"} {
		var is_there bool = false
		for _, o := range player.skills {
			if i == o {
				is_there = true
				break
			}
		}
		var skill_changer string
		if is_there {
			skill_changer = pad_right(i + " * ", 19)
		} else {
			skill_changer = pad_right(i, 19)
		}
		sheet = strings.ReplaceAll(sheet, i,skill_changer)
	}

	sheet_write, _ := os.Create(player.name + "_sheet.txt")
	defer sheet_write.Close()

	sheet_write.WriteString(sheet)
	
	empty_screen()
	fmt.Println("The character creation process is finished. " + player.name + "_sheet.txt awaits you!")
}

// main:
//   - Input: None
//   - Output: None
//   - The main function.
func main() {
	empty_screen()
	color.Assign_Colors()
	var player Player
	player.attrs = make(map[string]*Attribute)
	a := []string{"Strength", "Dexterity", "Constitution", "Intelligence", "Wisdom  ", "Charisma"}
	for _,i := range a{ 
		attribute := Attribute{string(i[:2]), string(i), 0, 0}
		player.attrs[string(i[:2])] = &attribute
	}
	player.Character_Builder()
	player.Sheet_Builder()
}
