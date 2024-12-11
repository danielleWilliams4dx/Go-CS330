package rpg
import(
	"fmt"
	"strings"
)

type Item struct{
	name string
	category string
	desc string
}

func (i Item) GetName() string{
	return i.name
}

func (i Item) ToString() string{
	return "Name: " + i.name + "\nCategory: " + i.category + "\nDescription:\n" + i.desc
}

func (i Item) TabbedToString() string{
	return "\tName: " + i.name + "\n\tCategory: " + i.category + "\n\tDescription:\n" + i.desc
}

type ItemFactory struct{
	ItemFile string
}

//extracts data from csv records to create an item list
func (itf ItemFactory) CreateItemList() []Item{
	itemList := []Item{}
	records := GetRecords(itf.ItemFile)
	for i, record := range records{
		name := strings.TrimSpace(record[0])
		category := strings.TrimSpace(record[1])
		desc := addLineBreak(record[2])
		itemList = append(itemList, Item{name, category, desc})
		i += 0 //prevents declared and not used error
	}
	return itemList
}

//prints all Items in a list
//if the Item is a HealingItem, retrieves tand prints that HealingItem
func PrintAllItems(il []Item){
	fmt.Println("Items:\n")
	for i := 0; i < len(il); i++{	
		fmt.Println(il[i].ToString() + "\n")
	}
}

//returns whether an Item is empty
//an Item is considered empty if it doesn't have a name
func (i Item) IsEmpty() bool{
	return i.name == ""
}

//returns whether two Items are the same
func (i Item) Equals(other Item) bool{
	return i.name == other.name && i.category == other.category && i.desc == other.desc
}

//when an Item is found, its description is outputted and it is added to the inventory
func (i Item) FindItem (h *Hero){
	fmt.Println(h.GetName() + " found " + i.GetName() + ".\n\n" + i.TabbedToString() + "\n")
	h.PickUpItem(i)

}

// //returns whether the Item is a HealingItem
// func (i Item) isHealingItem(hil []HealingItem) bool{
// 	for j, hi := range hil{
// 		if(i.Equals(hi.GetBaseItem())){
// 			return true
// 		}
// 		j += 0
// 	}
// 	return false
// }