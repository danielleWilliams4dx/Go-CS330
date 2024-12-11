package rpg

import(
	"fmt"
	"strings"
)

type Character struct{
	name string
	role string
}

//getters and setters
func (c Character) GetName() string{
	return c.name
}
func (c *Character) SetName(name string){
	c.name = name
}

func (c Character) GetRole() string{
	return c.role
}
func (c *Character) SetRole(role string){
	c.role = role
}

func (c Character) ToString() string{
	return "Name: " + c.name + "\nRole: " + c.role
}

type CharacterFactory struct{
	CharacterFile string
}

//extracts data from csv records to create a Character list
func (cf CharacterFactory) CreateCharacterList() []Character{
	characterList := []Character{}
	records := GetRecords(cf.CharacterFile)
	for i, record := range records{
		name := strings.TrimSpace(record[0])
		role := strings.TrimSpace(record[1])
		characterList = append(characterList, Character{name,role})
		i += 0 //prevents declared and not used error
	}
	return characterList
}

//prints all characters in a list
func PrintAllCharacters(cl []Character){
	fmt.Println("Characters:\n")
	for i := 0; i < len(cl); i++{
		fmt.Println(cl[i].ToString() + "\n")
	}
}

