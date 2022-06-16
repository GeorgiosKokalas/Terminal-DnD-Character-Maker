import time, os, sys, random

class Colours:
    colorwheel = {}
    def __init__(self):
        self.colorwheel["RESET"]           = "\033[0m"
        self.colorwheel["BLACK"]           = "\033[0;30m"
        self.colorwheel["RED"]             = "\033[0;31m"
        self.colorwheel["GREEN"]           = "\033[0;32m"
        self.colorwheel["YELLOW"]          = "\033[0;33m"
        self.colorwheel["BLUE"]            = "\033[0;34m"
        self.colorwheel["PURPLE"]          = "\033[0;35m"
        self.colorwheel["CYAN"]            = "\033[0;36m"
        self.colorwheel["WHITE"]           = "\033[0;37m"
        self.colorwheel["GRAY"]            = "\033[0;38m"
        self.colorwheel["DARK_GRAY"]       = "\033[1;30m"
        self.colorwheel["LIGHT_RED"]       = "\033[1;31m"
        self.colorwheel["LIGHT_GREEN"]     = "\033[1;32m"
        self.colorwheel["LIGHT_YELLOW"]    = "\033[1;33m"
        self.colorwheel["LIGHT_BLUE"]      = "\033[1;34m"
        self.colorwheel["LIGHT_PURPLE"]    = "\033[1;35m"
        self.colorwheel["LIGHT_CYAN"]      = "\033[1;36m"
        self.colorwheel["LIGHT_WHITE"]     = "\033[1;37m"
        self.colorwheel["LIGHT_GRAY"]      = "\033[1;38m"

    def creplace(self, text, old, new):
        text = text.replace(old, self.colorwheel[new]+old+self.colorwheel["RESET"])
        return text

class Player:
    class Attribute:
        def __init__(self, at_type, shorthand):
            self.at_type = at_type
            self.val = 0
            self.shorthand = shorthand
            self.Modifier = 0

    def __init__(self):
        self.name, self.age, self.race, self.clas  = ..., ..., ..., ...
        self.attributes = {x[:2] : Player.Attribute(x, x[:2]) for x in ("Strength", "Dexterity", "Consitution", "Intelligence", "Wisdom  ", "Charisma")}
        self.size, self.languages, self.speed, self.extras= ..., ..., ..., ...
        self.HP, self.savingT, self.armorP, self.weaponP, self.toolP, self.equipment, self.feats, self.skills = ..., "", ..., ..., ..., "", ..., []
    
    def race_parser(self, race_int):
        with open("race_attributor.txt") as f:
            race_text = f.read().split("\n")
        if (race_int == 0):
            print("Your race will be selected at random from the Basic Races")
            race_int = random.randrange(1, 10)
        elif (race_int == 46):
            print("Your race will be selected at random from the Races in the Expansions")
            race_int = random.randrange(10, 46)
        elif (race_int == 47):
            print("Your race will be selected at random from the Races in both the Basic game and the Expansions")
            race_int = random.randrange(1, 46)
        for i in race_text:
            if ("#" in i):
                continue
            i = i.split()
            if int(i[0]) == race_int:
                return i

    def select_race(self, race_int):
        parsed_race = self.race_parser(race_int)
        try:
            self.race = parsed_race[1].replace(".", " ")
        except:
            self.race = parsed_race[1]
        selected_attributes=["St", "De", "Co", "In", "Wi", "Ch"]
        at_boost_str = parsed_race[2] 
        try:
            at_boost_str = at_boost_str.split(",")
        except:
            at_boost_str = [at_boost_str]
        for i in at_boost_str:
            selected_attribute = i[0:2]
            if(selected_attribute == "Ra"):
                selected_attribute = selected_attributes.pop(random.randrange(0,len(selected_attributes)-1))
            self.attributes[selected_attribute].val += int(i[3])
        self.size = parsed_race[4]
        self.speed = int(parsed_race[5])
        self.extras = parsed_race[6].replace(".", " ")
        self.languages = parsed_race[3].replace("."," ")   

    def attribute_assigner(self):
        scores = [0, 0, 0, 0, 0, 0]
        for i in range(6):
            dice = []
            for c in range(5):
                dice.append(random.randrange(1,7))
                print(dice)
            for u in range(3):
                scores[i] += max(dice)
                dice.pop(dice.index(max(dice)))
        scores.sort()
        scores.reverse()
        attr = ["St", "De", "Co", "In", "Wi", "Ch"]
        for i in range(5):
            unsure = True
            while (unsure):
                empty_screen()
                print("Remaining values: ", end='')
                for i in scores:
                    print (f"{i}  ", end='')
                print(f"\nPlease assign this: {scores[0]}")
                for p in attr:
                    print(f"[{p}] {self.attributes[p].at_type}")
                selected_attr = input("Your choice: ")
                if (selected_attr in attr):
                    self.attributes[selected_attr].val += scores[0]
                    scores.pop(0)
                    attr.pop(attr.index(selected_attr))
                    unsure = False
                else:
                    print("\033[1;31mNot desired value entered\033[0m")
                    time.sleep(2)
        self.attributes[attr[0]].val = scores[0]
        print(f"Assigned the remaining value {scores[0]} to {self.attributes[attr[0]].at_type}")
        del attr
        del scores
        time.sleep(2)
        empty_screen()
        print("Calculating Modifiers:\n\033[0;32mAttribute\tScore\tModifier\033[0m")
        for i in ["St", "De", "Co", "In", "Wi", "Ch"]:
            self.attributes[i].Modifier = (self.attributes[i].val // 2) - 5
            print(f"{self.attributes[i].at_type}\t{self.attributes[i].val}\t{self.attributes[i].Modifier}")
        get_input()

    def class_parser(self,class_int):
        with open("class_attributor.txt") as f:
            class_text = f.read().split("\n")
        if (class_int == 0):
            print("Your class will be selected at random")
            class_int = random.randrange(1, 15)
        for i in class_text:
            if ("#" in i):
                continue
            i = i.split()
            if int(i[0]) == class_int:
                return i

    def select_class(self,class_int):
        parsed_class = self.class_parser(class_int)
        try:
            self.clas = parsed_class[1].replace(".", " ")
        except:
            self.clas = parsed_class[1]
        if (class_int == 0):
            get_input(reassurance=f"You are a(n) {self.clas}.")
        self.HP = int(parsed_class[2].split("+")[0]) + self.attributes[parsed_class[2].split("+")[1][:2]].Modifier
        for i in parsed_class[3].split(','):
            self.savingT += f"{self.attributes[i].at_type} "
        self.armorP = parsed_class[4].replace("."," ")
        self.weaponP = parsed_class[5].replace("."," ")
        self.toolP = parsed_class[6].replace("."," ")
        self.feats = parsed_class[8].replace("."," ")
        for i in parsed_class[7].replace("."," ").split(", "):
            if ("/ " in i):
                unsure = True
                while (unsure):
                    empty_screen()
                    choice = i.split("/ ")
                    print("Choose which piece of equipment you desire")
                    for o in range(len(choice)):
                        print(f"[{o}] {choice[o]}")
                    choice_input = get_input("Your choice: ", "Are you sure this is the item you want?")
                    try:
                        choice_input = int(choice_input)
                    except:
                        print("\033[1;31mValue entered not of type int. Please try again.\033[0m")
                        time.sleep(2)
                    else:
                        if (choice_input < 0 or choice_input >= len(choice)):
                            print("\033[1;31mInteger entered not within desired bounds. Please try again.\033[0m")
                            time.sleep(2)
                        else:
                            i = choice[choice_input]
                            unsure = False
            self.equipment += i + ", "
        self.equipment = self.equipment[:-2]
        skill_list=parsed_class[9].replace(".", " ").split(", ")[1:]
        for i in range(int(parsed_class[9][0])):
            unsure = True
            while (unsure):
                empty_screen()
                print(f"As a {self.clas}, you have access to {int(parsed_class[9][0])} Skill proficiencies from those below.")
                for o in range(len(skill_list)):
                    current_skill = skill_list[o]
                    print(f"[{o}]\t{current_skill}")
                print(f"Choice(s) remaining: {int(parsed_class[9][0])-i}")
                choice_input = get_input("Your choice: ", "Are you sure this is the skill?")
                try:
                    choice_input = int(choice_input)
                except:
                    print("\033[1;31mValue entered not of type int. Please try again.\033[0m")
                    time.sleep(2)
                else:
                    if (choice_input < 0 or choice_input >= len(skill_list)):
                        print("\033[1;31mInteger entered not within desired bounds. Please try again.\033[0m")
                        time.sleep(2)
                    else:
                        self.skills.append(skill_list[choice_input])
                        skill_list.pop(choice_input)
                        unsure = False                


def empty_screen():
    print('\x1b[' + '?25h', end='')             # Save Current Screen
    print('\x1b[' + '2J', end='')               # Clear Screen
    print('\x1b[' + 'H', end='')                # Move up

def read_from(filename):
    color=Colours()
    with open(filename) as f:
        text = f.read()
        
    if "(#COLOR)" in text:
        textie = text.split("(#COLOR)\n")
        text, colorchange = textie[0], textie[1].split("\n")
        for i in colorchange:
            h = i.split("|_|")
            text = color.creplace(text, h[1], h[0])
    
    texts = text.split("*")
    return texts

def reader(file, interval = 0):
    color = Colours()
    for i in read_from(file):
        print(i)
        time.sleep(interval)

def get_input(prompt = "", reassurance =""):
    unsure = True
    receiver = ""
    while (unsure):
        if (len(prompt) > 0):
            receiver = input(prompt)
        insurance = input(reassurance + " (ENTER to continue, N to reset, Q to quit) ")
        if (insurance.lower() == "q" or insurance.lower() == "quit"):
            print("Farewell for now. Let us hope we meet again.")
            exit()
        elif (insurance.lower() == "n" or insurance.lower() == "no"):
            time.sleep(0)
            unsure = True
        else:
            unsure = False
    return receiver

def Character_Builder(player):
    empty_screen()
    color = Colours()
    reader("welcome.txt", 1)
    get_input()
    
    empty_screen()
    player.name=get_input("Each character has a name, what will yours be? ")
    
    unselected = True
    race_int = ...
    while(unselected):
        empty_screen()
        reader("race_selector.txt",0)
        race_int = get_input("Please Select a number from 0 to 48, representing a race: ", "Are you sure this is your wanted race?")
        try:
            race_int = int(race_int)  
        except:
            print("\033[1;31mNo number detected. Please try again.\033[0m")
            time.sleep(2)
        else:
            if (race_int < 0 or race_int > 47):
                print("\033[1;31mInserted number not within desirable values. Please try again.\033[0m")
                time.sleep(2)
            else:
                unselected =False
    player.select_race(race_int)
    print(f"As a {player.race}, you are {player.size} size, you know {player.languages} and have a speed of {player.speed}, on top of having {player.extras}!")
    get_input()

    empty_screen()
    print("Time to Calculate your Attributes and their Modifiers!\nFor this, 5 random dice will be rolled and you will get the maximum 3 for each Attribute.")
    time.sleep(5)
    player.attribute_assigner()

    unselected = True
    class_int = ...
    while(unselected):
        empty_screen()
        reader("class_selector.txt",0)
        class_int = get_input("Please Select a number from 0 to 14, representing a class: ", "Are you sure this is your wanted class?")
        try:
            class_int = int(class_int)  
        except:
            print("\033[1;31mNo number detected. Please try again.\033[0m")
            time.sleep(2)
        else:
            if (class_int < 0 or class_int > 14):
                print("\033[1;31mInserted number not within desirable values. Please try again.\033[0m")
                time.sleep(2)
            else:
                unselected =False
    player.select_class(class_int)

def Sheet_Former(player):
    empty_screen()
    print(f"The character creation process has finished and {player.name}_sheet.txt is waiting for you!")
    with open("character_template.txt") as f:
        sheet = f.read()

    sheet = sheet.replace("Name:  Race:  Class:", f"Name:{player.name}\tRace:{player.race}\tClass:{player.clas} ")
    sheet = sheet.replace("Size:    HP:   Speed:", f"Size:{player.size}\tHP:{player.HP}\tSpeed:{player.speed}")
    
    for i in player.attributes:
        sheet = sheet.replace(f"{player.attributes[i].at_type}",str(player.attributes[i].at_type).ljust(13) + "|  "+str(player.attributes[i].val).ljust(4)  + "|  "+str(player.attributes[i].Modifier).ljust(4))
    sheet = sheet.replace("Saving Throws:", f"Saving Throws: {player.savingT}")

    sheet = sheet.replace("Languages:", f"Languages: {player.languages}")
    sheet = sheet.replace("Armor Proficiencies: ", f"Armor Proficiencies: {player.armorP}")
    sheet = sheet.replace("Weapon Proficiencies: ", f"Weapon Proficiencies: {player.weaponP}")
    sheet = sheet.replace("Tool Proficiencies: ", f"Tool Proficiencies: {player.toolP}")
    sheet = sheet.replace("Extras:", f"Extras: {player.extras}")
    sheet = sheet.replace("Feats: ", f"Feats: {player.feats}")
    sheet = sheet.replace("Items: ", f"Items: {player.equipment}")

    for i in ["Athletics", "Acrobatics", "Sleight of Hand", "Stealth", "Arcana", "History", "Investigation", "Nature", "Religion", "Animal Handling", "Insight", "Medicine", "Perception", "Survival", "Deception", "Intimidation", "Performance", "Persuasion"]:
        if i in player.skills:
            sheet = sheet.replace(f"{i}", str(str(i)+" * ").ljust(19))
        else:
            sheet = sheet.replace(f"{i}", str(i).ljust(19))

    with open(player.name+"_sheet.txt",'w') as f:
        f.write(sheet)

def main():
    empty_screen()
    color = Colours()
    player = Player()
    Character_Builder(player)
    Sheet_Former(player)
    print('\x1b[' + '?47l', end='')             # Restore Screen

if __name__ == '__main__':
    main()
