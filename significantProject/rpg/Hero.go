package rpg

import(
	"fmt"
	"strings"
	"math/rand"
)

type Hero struct{
	baseChar Character
	level int64
	maxHP int64 //max hp
	hp int64
	moveset []Move
	learnableMoves map[int64]Move
	experience int64 //100 exp = +1 level
	inventoryItem []Item
	inventoryHealingItem []HealingItem
}

//getters and setters

func (h Hero) GetName() string{
	return h.baseChar.name
}

func (h Hero) GetRole() string{
	return h.baseChar.role
}

func (h Hero) GetMoveset() []Move{
	return h.moveset
}

func (h Hero) GetInventoryHealingItem() []HealingItem{
	return h.inventoryHealingItem
}

func (h Hero) GetInventoryItem() []Item{
	return h.inventoryItem
}
func (h *Hero) EmptyInventoryItem(){
	h.inventoryItem = []Item{}
}

func (h Hero) GetHP() int64{
	return h.hp
}

func (h *Hero) SetHP(num int64){
	h.hp = num
}

func (h Hero) GetMaxHP() int64{
	return h.maxHP
}

//other functions

//Sets the hero's stats to their starting values
//Called after the character is named
func CreateHero (name string, ml []Move) Hero{
	baseChar := Character{name, "Mercenary"}
	var level int64 = 10
	var maxHP int64 = 75
	hp := maxHP
	moveset := []Move{GetMoveFromName("Snack Time", ml), GetMoveFromName("Rotten Tomato Chuck", ml), GetMoveFromName("Glitter Bomb", ml)}
	learnableMoves := CompileLearnableMoves(ml)
	var experience int64 = 0
	inventoryItem := []Item{}
	inventoryHealingItem := []HealingItem{}
	return Hero{baseChar, level, maxHP, hp, moveset, learnableMoves, experience, inventoryItem, inventoryHealingItem}
}

//uses the Hero's level to calculate their maxHP
func (h Hero) calcMaxHp() int64{
	var num float32 = 3.125
	return int64(43.75 + float32(h.level)*num)
}

//calculates and sets the maxHP
func (h *Hero) updateMaxHP(){
	h.maxHP = h.calcMaxHp()
}

//increases the Hero's experience
//if the experience exceeds the threshhold (100), a levelUp is triggered
func (h *Hero) IncreaseExperience(exp int64){
	h.experience += exp
	fmt.Println(h.GetName() + " gained " + fmt.Sprint(exp) + " EXP.")
	for(h.experience >= 100){
		h.levelUp()
		h.experience -= 100
	}
	fmt.Println("\n" + h.GetAllStats())
}

//increases the Hero's level by 1 and updates their maxHP
//if there is a Move that can be learned, the Hero learns the new Move
func (h *Hero) levelUp(){
	h.level++
	h.updateMaxHP()
	fmt.Println(h.GetName() + " leveled up.")
	h.learnMove()
}

//if a Move can be learned and the Hero already has 4 moves,
//the user is prompted which Move they would like to replace
func (h *Hero) learnMove(){
	if(h.isLearnableMove()){
		newMove := h.learnableMoves[h.level]
		if(len(h.moveset) < 4){
			h.moveset = append(h.moveset, newMove)
		}else{
			fmt.Println(h.GetName() + " would like to learn " + newMove.GetName() + ".\n\n" + newMove.TabbedToString())
			fmt.Println("\n" + h.GetName() + "'s " + GetMovesetAsString(h.moveset))
			fmt.Print("\nWhat move would you like to replace (1, 2, etc.)? ")
			index := ChooseMoveReplaceIndex(*h)
			h.moveset[index] = newMove
			fmt.Println()
		}
		fmt.Println(h.GetName() + " learned " + newMove.GetName())
	}
}

//returns whether the Hero's level has a corresponding learnable Move
func (h *Hero) isLearnableMove() bool{
	for key := range h.learnableMoves{
		if(key == h.level){
			return true
		}
	}
	return false
}

func (h Hero) GetAllStats() string{
	return h.baseChar.ToString() + "\nLevel: " + fmt.Sprint(h.level) + "\nHP: " + fmt.Sprint(h.hp) + "/" + fmt.Sprint(h.maxHP) + "\nExperience: " + fmt.Sprint(h.experience) + "/100\n"
}

func (h Hero) ToString() string{
	return h.baseChar.ToString() + "\nLevel: " + fmt.Sprint(h.level) + "\nHP: " + fmt.Sprint(h.hp) + "/" + fmt.Sprint(h.maxHP) + "\nExperience: " + fmt.Sprint(h.experience) + "/100\n" + GetMovesetAsString(h.moveset) + h.GetInventoryAsString()
}

//returns all Items and HealingItems in the inventory as a string
func (h Hero) GetInventoryAsString() string{
	index := 1
	s := "\nInventory:\n"
	for i := 0; i < len(h.inventoryItem); i++{
		s += "     " + fmt.Sprint(index) + "\n" + h.inventoryItem[i].TabbedToString()
		if(i < len(h.inventoryItem)-1){
			s += "\n\n"
		}
		index++
	}
	if(len(h.inventoryItem) > 0){
		s += "\n\n"
	}

	for j := 0; j < len(h.inventoryHealingItem); j++{
		s += "     " + fmt.Sprint(index) + "\n" + h.inventoryHealingItem[j].TabbedToString()
		if(j < len(h.inventoryHealingItem)-1){
			s += "\n\n"
		}
		index++
	}
	return s
}

//adds a regular Item to the inventory
func (h *Hero) PickUpItem(i Item){
	h.inventoryItem = append(h.inventoryItem, i)
}

//adds a HealingItem to the inventory
func (h *Hero) PickUpHealingItem(hi HealingItem){
	h.inventoryHealingItem = append(h.inventoryHealingItem, hi)
}

//checks the inventory for a revive
//returns the index if it exists, -1 otherwise
func (h Hero) hasRevive() int{
	for i, item := range(h.inventoryHealingItem){
		if(item.GetName() == "Life Gumdrop"){
			return i
		}
	}
	return -1
}

//restores the PP of all of the Hero's moves
func (h *Hero) FullPPHeal(){
	for i := 0; i < len(h.moveset); i++ {
		move := &h.moveset[i]
		move.RestorePP()
	}
}

//finds the specified HealingItem in the Hero's inventory
//and removes it
func (h *Hero) RemoveHealingItem(hi HealingItem){
	for j, item := range h.inventoryHealingItem{
		if(item.GetName() == hi.GetName()){
			//delete the item from the inventory
			h.inventoryHealingItem = append(h.inventoryHealingItem[:j], h.inventoryHealingItem[j+1:]...)
			return //break out of the function
		}
	}
	fmt.Println("The item could not be removed.")
}

//revives the Hero, if they have a revive and choose to be revived
//removes Life Gumdrop from the inventory
func (h *Hero) UseRevive() bool{
	i := h.hasRevive()
	if(i != -1){
		ans := ""
		fmt.Print("\nYou hear a faint voice...\n" + h.GetName() + " don't give up!\n\nUse a Life Gumdrop (y/n)? ")
		fmt.Scanln(&ans)
		if(ans == "Y" || ans == "y"){
			h.inventoryHealingItem[i].ItemHeal(h)
			return true
		}
	}
	return false
}

//Updates the Enemy's HP
//If the Enemy's HP is 0, call Enemy's defeated method
func (h *Hero) doDamage(damage int64, e *Enemy){
	if(e.GetHP() - damage <= 0){
		e.SetHP(0)
	}else{
		e.SetHP(e.GetHP() - damage)
	}
}

//calculates whether the move hits
//calculates whether the move is a crit (damage x1.5)
//does damage and subtracts the Move's PP 
func (h Hero) Attack(m *Move, e *Enemy){
	m.DecreasePP()
	fmt.Println("\n— " + h.GetName() + " used " + m.GetName() + ". —")
	if(m.HitOrMiss()){
		if(m.IsCrit()){
			h.doDamage((m.GetDamage()*5)/4, e)
			fmt.Println("It was a critical hit!")
		}else{
			h.doDamage(m.GetDamage(), e)
		}
	}else{
		fmt.Println("The attack missed!")
	}
}

//calculates whether the move hits
//heals the user and subtracts the Move's PP 
func (h *Hero) HealMove(m *Move){
	m.DecreasePP()
	fmt.Println("— " + h.GetName() + " used " + m.GetName() + ". —")
	if(m.HitOrMiss()){
		//if the move will restore more than the maxHP, max out the hp
		if(-1*m.GetDamage() + h.hp > h.maxHP){
			h.SetHP(h.maxHP)
			fmt.Println(h.GetName() + " restored their HP.")
		}else{
			h.SetHP(h.GetHP() + -1*m.GetDamage())
			fmt.Println(h.GetName() + " restored some of their HP.")
		}
	}else{
		fmt.Println("The move failed!")
	}
}

//the Hero struggles if they have no attacking moves left
//it does a little damage to the Enemy, and more damage to itself
func (h *Hero) Struggle(e *Enemy){
	h.doDamage(10, e)
	if(h.GetHP() - 25 > 0){
		h.SetHP(e.GetHP()-25)
	}else{
		h.SetHP(0)
	}
	fmt.Println("\n" + h.GetName() + " has no moves left!\n— " + h.GetName() + " used Struggle. —")
}

//get rid of all PP
//for testing purposes only
func (h *Hero) RemoveAllPP(){
	for i := 0; i < len(h.moveset); i++{
		move := &h.moveset[i]
		move.SetPP(0)
	}
}

//if the user selects run, there's a 30% that it is successful
//if the enemy is a boss, this always fails
func (h Hero) Run(e Enemy) bool{
	if(!strings.Contains(e.GetRole(), "Boss") && !strings.Contains(e.GetRole(), "Possessed") && !strings.Contains(e.GetRole(), "Wizard") && rand.Intn(3) == 0){
		fmt.Println("\n"+ h.GetName() + " ran away.\n")
		return true
	}else{
		fmt.Println("\n"+ h.GetName() + " couldn't get away.")
		return false
	}
}