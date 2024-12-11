package rpg

import (
	"fmt"
	"math/rand"
	"strings"
)

type Room struct{
	num int
	enemy Enemy
	item Item
	healItem HealingItem
	up *Room
	down *Room
	left *Room
	right *Room
	boss bool
	locked bool
	visited bool
}

//getters and setters
func (r *Room) SetNum(num int){
	r.num = num
}
func (r Room) GetNum() int{
	return r.num
}

func (r *Room) SetUp(other *Room){
	r.up = other
}
func (r Room) GetUp() *Room{
	return r.up
}

func (r *Room) SetDown(other *Room){
	r.down = other
}
func (r Room) GetDown() *Room{
	return r.down
}

func (r *Room) SetLeft(other *Room){
	r.left = other
}
func (r Room) GetLeft() *Room{
	return r.left
}

func (r *Room) SetRight(other *Room){
	r.right = other
}
func (r Room) GetRight() *Room{
	return r.right
}

func (r *Room) SetVisited(tf bool){
	r.visited = tf
}
func (r Room) IsVisited() bool{
	return r.visited
}

func (r *Room) SetBoss(e Enemy){
	r.enemy = e
}
func (r *Room) MakeBoss(tf bool){
	r.boss = tf
}

func (r *Room) Lock(tf bool){
	r.locked = tf
}
func(r *Room) IsLocked() bool{
	return r.locked
}

func (r Room) HasEnemy() bool{
	return !r.enemy.IsEmpty()
}
func (r Room) EnemyIsDefeated() bool{
	return r.enemy.GetHP() == 0
}
func (r *Room) SetEnemy(e Enemy){
	r.enemy = e
}

func (r Room) HasItem() bool{
	return !r.item.IsEmpty()
}
func (r *Room) RemoveItem() {
	r.item = Item{}
}
func (r Room) GetItem() Item{
	return r.item
}

func (r Room) HasHealingItem() bool{
	return !r.healItem.IsEmpty()
}
func (r *Room) SetHealingItem(hi HealingItem){
	r.healItem = hi
}

func (r Room) IsBoss() bool{
	return r.boss
}

//assigns a specific key to the room
func(r *Room) SetKey(keyIndex int, il []Item){
	if(keyIndex < len(il)){
		r.item = il[keyIndex]
	}else{
		fmt.Println("Key index out of bounds.")
	}
}

//70% of the time, assign a random enemy from a given group
//otherwise the room will not have an enemy
func (r *Room) AssignEnemy(group []Enemy){
	if(rand.Intn(10) > 2){
		r.enemy = GetRandomEnemy(group)
	}
}

//80% of the time, assign a random HealingItem
//otherwise the room will not have an HealingItem
func (r *Room) AssignHealingItem(hil []HealingItem){
	if(rand.Intn(10) > 1){
		r.healItem = GetRandomHealingItem(hil)
	}
}

func (r Room) ToString() string{
	s := "\nRoom: #" + fmt.Sprint(r.num)

	if(!r.enemy.IsEmpty()){
		s += "\nEnemy:\n" + r.enemy.ToString()
	}

	s += "\nItem(s):\n"
	if(!r.item.IsEmpty()){
		s += r.item.ToString() + "\n"
	}
	if(!r.healItem.IsEmpty()){
		s += r.healItem.ToString() +"\n"
	}

	if(r.up != nil){
		s += "\nUp: Room #" + fmt.Sprint(r.up.GetNum())
	}
	if(r.down != nil){
		s += "\nDown: Room #" + fmt.Sprint(r.down.GetNum())
	}
	if(r.left != nil){
		s += "\nLeft: Room #" + fmt.Sprint(r.left.GetNum())
	}
	if(r.right != nil){
		s += "\nRight: Room #" + fmt.Sprint(r.right.GetNum())
	}

	s += "\nBoss: " + fmt.Sprint(r.boss)
	s += "\nLocked: " + fmt.Sprint(r.locked)
	s += "\nVisited: " + fmt.Sprint(r.visited) + "\n"

	return s
}

//prompts the user which direction they would like to go
//returns a pointer to the selected Room
func (r Room) pickDirection(h *Hero) *Room{
	strToDir := map[string]*Room{
		"U" : r.up,
		"D" : r.down,
		"L" : r.left,
		"R" : r.right,
	}

	dir := ""
	fmt.Print("Please select a direction.\n" + getAvailablePathsAsString(strToDir))
	fmt.Scanln(&dir)
	dir = strings.ToUpper(dir)

	for(!isValidDirection(dir, strToDir, h)){
		fmt.Print("Please select a direction.\n" + getAvailablePathsAsString(strToDir))
		fmt.Scanln(&dir)
		dir = strings.ToUpper(dir)
	}
	fmt.Println()

	return strToDir[dir]
}

//returns the paths that the Hero can take
//up, down, left, right as a string
//adds whether a room has already been visited
func getAvailablePathsAsString(m map[string]*Room) string{
	s := "Options: "

	for key := range m{
		if(m[key] != nil){
			s += key + " "
			if(m[key].IsVisited()){
				s += "(visited) "
			}
		}
	}

	return s
}

//returns true if the inputted direction is valid and false otherwise
//if the Room is locked and the Hero doesn't have all three key,
//the Room is not considered valid
func isValidDirection(dir string, m map[string]*Room, h *Hero) bool{
	for key := range m{
		if(dir == key && m[key] != nil){
			if(m[key].IsLocked()){
				if(len(h.GetInventoryItem()) == 3){
					fmt.Println(h.GetName() + " unlocked the door.")
					h.EmptyInventoryItem()
					return true
				}else{
					fmt.Println("The door is locked.\nFind all three keys to enter.")
					return false
				}
			}
			return true
		}
	}
	fmt.Println("Invalid input.")
	return false
}

//simulates entering the Room
//sets visited to true
//starts a battle if there is an Enemy
//resets the Enemy and HealingItem
//the Hero picks up an Item or HealingItem if there is one
//if the Room is not a boss stage, 
//prompts the user to pick a direction to go
//and enters that Room
func (r *Room) Enter(h* Hero, eg []Enemy, hil []HealingItem){
	r.SetVisited(true)
	if(r.HasEnemy() && !r.EnemyIsDefeated()){
		win := Battle(&r.enemy, h)
		if(!win){
			//reset Enemy
			//exit if the Hero's HP is 0
			//(distinction of a game over vs run away)
			r.enemy.SetHP(r.enemy.GetMaxHP())
			r.enemy.FullPPHeal()
			if(h.GetHP() == 0){
				return
			}
		}
	}
	if(r.HasItem()){
		r.item.FindItem(h)
	}
	if(r.HasHealingItem()){
		r.healItem.FindHealingItem(h)
	}
	if(!r.IsBoss()){
		r.AssignEnemy(eg)
		r.AssignHealingItem(hil)
		if(r.HasItem()){
			r.RemoveItem()
		}
		r.pickDirection(h).Enter(h, eg, hil)
	}
	r.SetVisited(false)
}
