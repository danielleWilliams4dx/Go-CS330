package rpg

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

type HealingItem struct{
	baseItem Item
	hp int64
	rarity int64 //a number from 1-5
	prefix string //to address healing item in story if found
}

func (hi HealingItem) GetName() string{
	return hi.baseItem.name
}

func (hi HealingItem) GetCategory() string{
	return hi.baseItem.category
}

func (hi HealingItem) GetDescription() string{
	return hi.baseItem.desc
}

func (hi HealingItem) GetBaseItem() Item{
	return hi.baseItem
}

func (hi HealingItem) GetHP() int64{
	return hi.hp
}

func (hi HealingItem) GetPrefix() string{
	return hi.prefix
}

//applies the HealingItem's effects to the hero
func (hi HealingItem) ItemHeal(h *Hero){
	fmt.Println()
	if(hi.GetCategory() == "Full Revive"){
		h.FullPPHeal()
		h.SetHP(h.maxHP)
		fmt.Println("All of " + h.GetName() + "'s HP and PP was restored.")
	}else if(hi.GetCategory() == "Full PP Heal"){
		h.FullPPHeal()
		fmt.Println("All of " + h.GetName() + "'s PP was restored.")
	}else if(hi.GetCategory() == "Full HP Heal"){
		h.SetHP(h.maxHP)
		fmt.Println("All of " + h.GetName() + "'s HP was restored.")
	}else if(hi.GetCategory() == "HP Heal"){
		if(h.GetHP() + hi.GetHP() < h.GetMaxHP()){
			h.SetHP(h.GetHP() + hi.GetHP())
		}else{
			h.SetHP(h.GetMaxHP())
		}

		if(h.hp < h.maxHP){
			fmt.Print("Some of ")
		}
		fmt.Println(h.GetName() + "'s HP was restored.")
	}else{
		ChooseMoveRestore(h).RestorePP()
		fmt.Println("\nThe move's PP was restored.")
	}
	fmt.Println("The " + hi.GetName() + " was used up.")
	h.RemoveHealingItem(hi) //remove the healing item from the inventory
}

//does not include the prefix
//only includes hp if it is not 0 or full (200)
func (hi HealingItem) ToString() string{
	if(hi.hp != 0 && hi.hp != 200){
		return "Name: " + hi.GetName() + "\nCategory: " + hi.GetCategory() + "\nHP: " + fmt.Sprint(hi.hp) + "\nRarity: " + fmt.Sprint(hi.rarity) + "\nDescription:\n" + hi.GetDescription()
	}
	return "Name: " + hi.GetName() + "\nCategory: " + hi.GetCategory() + "\nRarity: " + fmt.Sprint(hi.rarity) + "\nDescription:\n" + hi.GetDescription()
}

func (hi HealingItem) TabbedToString() string{
	if(hi.hp != 0 && hi.hp != 200){
		return "\tName: " + hi.GetName() + "\n\tCategory: " + hi.GetCategory() + "\n\tHP: " + fmt.Sprint(hi.hp) + "\n\tRarity: " + fmt.Sprint(hi.rarity) + "\n\tDescription:\n" + hi.GetDescription()
	}
	return "\tName: " + hi.GetName() + "\n\tCategory: " + hi.GetCategory() + "\n\tRarity: " + fmt.Sprint(hi.rarity) + "\n\tDescription:\n" + hi.GetDescription()
}

type HealingItemFactory struct{
	HealingItemFile string
}

//extracts data from csv records to create a healingItem list
func (hif HealingItemFactory) CreateHealingItemList() []HealingItem{
	healingItemList := []HealingItem{}
	records := GetRecords(hif.HealingItemFile)
	for i, record := range records{
		name := strings.TrimSpace(record[0])
		category := strings.TrimSpace(record[1])
		desc := addLineBreak(record[2])
		hp, error := strconv.ParseInt(record[3],10,64)
		PrintError(error)
		rarity, error := strconv.ParseInt(record[4],10,64)
		PrintError(error)
		prefix := strings.TrimSpace(record[5])
		healingItemList = append(healingItemList, HealingItem{Item{name, category, desc}, hp, rarity, prefix})
		i += 0 //prevents declared and not used error
	}
	return healingItemList
}

//prints all healing items in a list
func PrintAllHealingItems(hi []HealingItem){
	fmt.Println("Healing Items:\n")
	for i := 0; i < len(hi); i++{
		fmt.Println(hi[i].TabbedToString() + "\n")
	}
}

//returns a list of HealingItems with a specified rarity
func GetAllHealingItemsByRarity(hil []HealingItem, rarity int64) []HealingItem{
	withRarity := []HealingItem{}

	for i, item := range hil{
		if(item.rarity == rarity){
			withRarity = append(withRarity, item)
		}
		i += 0
	}

	return withRarity
}

//returns a random HealingItem based on its rarity
func GetRandomHealingItem(hil []HealingItem) HealingItem{
	randNum := rand.Intn(25) + 1
	itemRange := []HealingItem{}

	if(randNum < 10){
		itemRange = GetAllHealingItemsByRarity(hil, int64(1))
	}else if(randNum < 17){
		itemRange = GetAllHealingItemsByRarity(hil, int64(2))
	}else if(randNum < 21){
		itemRange = GetAllHealingItemsByRarity(hil, int64(3))
	}else if(randNum < 24){
		itemRange = GetAllHealingItemsByRarity(hil, int64(4))
	}else{
		itemRange = GetAllHealingItemsByRarity(hil, int64(5))
	}

	return itemRange[rand.Intn(len(itemRange))]
}

//returns whether a HealingItem is empty
//an HealingItem is considered empty if it doesn't have a name
func (hi HealingItem) IsEmpty() bool{
	return hi.GetName() == ""
}

//when a HealingItem is found, its description is outputted and it is added to the inventory
func (hi HealingItem) FindHealingItem (h *Hero){
	fmt.Println(h.GetName() + " found " + hi.GetPrefix() + " " + hi.GetName() + ".\n\n" + hi.TabbedToString() + "\n")
	h.PickUpHealingItem(hi)
}
 
// //returns the corresponding HealingItem from the base Item
// func GetHealingItemFromBaseItem(item Item, hil []HealingItem) HealingItem{
// 	for i, hi := range hil{
// 		if(item.Equals(hi.GetBaseItem())){
// 			return hi
// 		}
// 		i += 0
// 	}
// 	fmt.Println("The HealingItem does not exist.")
// 	return HealingItem{}
// }