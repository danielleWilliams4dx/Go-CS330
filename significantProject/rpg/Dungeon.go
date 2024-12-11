package rpg

import(
	"fmt"
	"strings"
)

type Dungeon struct{
	entrance *Room
	num int
	name string
	enemyGroup []Enemy
}

//create and return the 1st dungeon
func BuildDungeon1(el []Enemy, il []Item, hil []HealingItem) Dungeon{
	slimes := GetEnemyGroup(el, "Slime")

	entrance := Room{}	
	room1 := Room{}
	room2 := Room{}
	room3 := Room{}
	room4 := Room{}
	room5 := Room{}
	room6 := Room{}
	room7 := Room{}
	room8 := Room{} 

	//building room1
	room1.SetNum(1)
	room1.AssignEnemy(slimes)
	room1.AssignHealingItem(hil)
	room1.SetKey(0,il)
	room1.SetRight(&entrance)
	entrance.SetLeft(&room1)

	//building room 2
	room2.SetNum(2)
	room2.SetBoss(GetEnemyByName(el, "King Slime"))
	room2.AssignHealingItem(hil)
	room2.MakeBoss(true)
	room2.Lock(true)
	room2.SetDown(&entrance)
	entrance.SetUp(&room2)

	//building room 3
	room3.SetNum(3)
	room3.AssignEnemy(slimes)
	room3.AssignHealingItem(hil)
	room3.SetLeft(&entrance)
	entrance.SetRight(&room3)

	//building room 4
	room4.SetNum(4)
	room4.AssignEnemy(slimes)
	room4.AssignHealingItem(hil)
	room4.SetUp(&room3)
	room3.SetDown(&room4)

	//building room 5
	room5.SetNum(5)
	room5.AssignEnemy(slimes)
	room5.AssignHealingItem(hil)
	room5.SetKey(1,il)
	room5.SetLeft(&room4)
	room4.SetRight(&room5)

	//building room 6
	room6.SetNum(6)
	room6.AssignEnemy(slimes)
	room6.AssignHealingItem(hil)
	room6.SetDown(&room5)
	room5.SetUp(&room6)
	room6.SetLeft(&room3)
	room3.SetRight(&room6)

	//building room 7
	room7.SetNum(7)
	room7.AssignEnemy(slimes)
	room7.AssignHealingItem(hil)
	room7.SetKey(2,il)
	room7.SetDown(&room6)
	room6.SetUp(&room7)

	//building room 8
	room8.SetNum(8)
	room8.AssignEnemy(slimes)
	room8.AssignHealingItem(hil)
	room8.SetDown(&room7)
	room7.SetUp(&room8)

	//debugging
	// fmt.Println(room1.ToString())
	// fmt.Println(room2.ToString())
	// fmt.Println(room3.ToString())
	// fmt.Println(room4.ToString())
	// fmt.Println(room5.ToString())
	// fmt.Println(room6.ToString())
	// fmt.Println(room7.ToString())
	// fmt.Println(room8.ToString())

	return Dungeon{&entrance, 1, "Slime Cavern", GetEnemyGroup(el, "Slime")}
}

//create and return the 2nd dungeon
func BuildDungeon2(el []Enemy, il []Item, hil []HealingItem) Dungeon{
	skeletons := GetEnemyGroup(el, "Skeleton")

	entrance := Room{}	
	room1 := Room{}
	room2 := Room{}
	room3 := Room{}
	room4 := Room{}
	room5 := Room{}
	room6 := Room{}
	room7 := Room{}
	room8 := Room{} 
	room9 := Room{}
	room10 := Room{} 

	//building room1
	room1.SetNum(1)
	room1.AssignEnemy(skeletons)
	room1.AssignHealingItem(hil)
	room1.SetRight(&entrance)
	entrance.SetLeft(&room1)

	//building room 2
	room2.SetNum(2)
	room2.AssignEnemy(skeletons)
	room2.AssignHealingItem(hil)
	room2.SetKey(0,il)
	room2.SetDown(&entrance)
	entrance.SetUp(&room2)

	//building room 3
	room3.SetNum(3)
	room3.AssignEnemy(skeletons)
	room3.AssignHealingItem(hil)
	room3.SetLeft(&room2)
	room2.SetRight(&room3)

	//building room 4
	room4.SetNum(4)
	room4.SetBoss(GetEnemyByName(el, "Sketchy Skeleton"))
	room4.AssignHealingItem(hil)
	room4.MakeBoss(true)
	room4.Lock(true)
	room4.SetUp(&room3)
	room3.SetDown(&room4)

	//building room 5
	room5.SetNum(5)
	room5.AssignEnemy(skeletons)
	room5.AssignHealingItem(hil)
	room5.SetLeft(&room3)
	room3.SetRight(&room5)

	//building room 6
	room6.SetNum(6)
	room6.AssignEnemy(skeletons)
	room6.AssignHealingItem(hil)
	room6.SetUp(&room5)
	room5.SetDown(&room6)

	//building room 7
	room7.SetNum(7)
	room7.AssignEnemy(skeletons)
	room7.AssignHealingItem(hil)
	room7.SetKey(1,il)
	room7.SetUp(&room6)
	room6.SetDown(&room7)

	//building room 8
	room8.SetNum(8)
	room8.AssignEnemy(skeletons)
	room8.AssignHealingItem(hil)
	room8.SetKey(2,il)
	room8.SetLeft(&room5)
	room5.SetRight(&room8)

	//building room 9
	room9.SetNum(9)
	room9.AssignEnemy(skeletons)
	room9.AssignHealingItem(hil)
	room9.SetUp(&room8)
	room8.SetDown(&room9)
	room9.SetLeft(&room6)
	room6.SetRight(&room9)

	//building room 10
	room10.SetNum(10)
	room10.AssignEnemy(skeletons)
	room10.AssignHealingItem(hil)
	room10.SetUp(&room9)
	room9.SetDown(&room10)
	room10.SetLeft(&room7)
	room7.SetRight(&room10)

	//debugging
	// fmt.Println(room1.ToString())
	// fmt.Println(room2.ToString())
	// fmt.Println(room3.ToString())
	// fmt.Println(room4.ToString())
	// fmt.Println(room5.ToString())
	// fmt.Println(room6.ToString())
	// fmt.Println(room7.ToString())
	// fmt.Println(room8.ToString())
	// fmt.Println(room9.ToString())
	// fmt.Println(room10.ToString())

	return Dungeon{&entrance, 2, "Labyrinth of Skulls", GetEnemyGroup(el, "Skeleton")}
}

//create and return the 3rd dungeon
func BuildDungeon3(el []Enemy, il []Item, hil []HealingItem) Dungeon{
	zombies := GetEnemyGroup(el, "Zombie")

	entrance := Room{}	
	room1 := Room{}
	room2 := Room{}
	room3 := Room{}
	room4 := Room{}
	room5 := Room{}
	room6 := Room{}
	room7 := Room{}
	room8 := Room{} 
	room9 := Room{}
	room10 := Room{} 

	//building room1
	room1.SetNum(1)
	room1.AssignEnemy(zombies)
	room1.AssignHealingItem(hil)
	room1.SetKey(0,il)

	//building room 2
	room2.SetNum(2)
	room2.AssignEnemy(zombies)
	room2.AssignHealingItem(hil)

	//building room 3
	room3.SetNum(3)
	room3.AssignEnemy(zombies)
	room3.AssignHealingItem(hil)
	room3.SetLeft(&room1)
	room1.SetRight(&room3)
	room3.SetUp(&room2)
	room2.SetDown(&room3)

	//building room 4
	room4.SetNum(4)
	room4.AssignEnemy(zombies)
	room4.AssignHealingItem(hil)
	room4.SetUp(&room3)
	room3.SetDown(&room4)
	room4.SetRight(&entrance)
	entrance.SetLeft(&room4)

	//building room 5
	room5.SetNum(5)
	room5.AssignEnemy(zombies)
	room5.AssignHealingItem(hil)
	room5.SetKey(1,il)
	room5.SetLeft(&room2)
	room2.SetRight(&room5)

	//building room 6
	room6.SetNum(6)
	room6.AssignEnemy(zombies)
	room6.AssignHealingItem(hil)
	room6.SetUp(&room5)
	room5.SetDown(&room6)
	room6.SetLeft(&room3)
	room3.SetRight(&room6)
	room6.SetDown(&entrance)
	entrance.SetUp(&room6)

	//building room 7
	room7.SetNum(7)
	room7.AssignEnemy(zombies)
	room7.AssignHealingItem(hil)
	room7.SetLeft(&room6)
	room6.SetRight(&room7)

	//building room 8
	room8.SetNum(8)
	room8.SetBoss(GetEnemyByName(el, "Undead Queen"))
	room8.AssignHealingItem(hil)
	room8.MakeBoss(true)
	room8.Lock(true)

	//building room 9
	room9.SetNum(9)
	room9.AssignEnemy(zombies)
	room9.AssignHealingItem(hil)
	room9.SetLeft(&room7)
	room7.SetRight(&room9)
	room9.SetUp(&room8)
	room8.SetDown(&room9)

	//building room 10
	room10.SetNum(10)
	room10.AssignEnemy(zombies)
	room10.AssignHealingItem(hil)
	room10.SetKey(2,il)
	room10.SetLeft(&room9)
	room9.SetRight(&room10)

	//debugging
	// fmt.Println(room1.ToString())
	// fmt.Println(room2.ToString())
	// fmt.Println(room3.ToString())
	// fmt.Println(room4.ToString())
	// fmt.Println(room5.ToString())
	// fmt.Println(room6.ToString())
	// fmt.Println(room7.ToString())
	// fmt.Println(room8.ToString())
	// fmt.Println(room9.ToString())
	// fmt.Println(room10.ToString())

	return Dungeon{&entrance, 3, "Royal Crypt", GetEnemyGroup(el, "Zombie")}
}

//simulates a dungeon traversal
//the Hero starts at the entrance
//if the user would like to restart,
//the Hero's HP & PP is restored before restarting
//returns false if the hero has a game over and does not want to restart the dungeon
//returns true if the dungeon is completed
func (d Dungeon) DungeonTraversal(h *Hero, hil []HealingItem) bool{
	fmt.Println("Entering the " + d.name + ".\n")
	d.entrance.Enter(h, d.enemyGroup, hil)

	if(h.GetHP() > 0){
		fmt.Println(h.GetName() + " survived the " + d.name + ".")
	}else{
		if(restartDungeon()){
			h.SetHP(h.GetMaxHP())
			h.FullPPHeal()
			d.entrance.SetEnemy(Enemy{})
			d.entrance.SetHealingItem(HealingItem{})
			h.EmptyInventoryItem()
			d.DungeonTraversal(h, hil)
		}else{
			return false
		}
	}
	return true
}

//prompts the user if they would like to restart the Dungeon
//returns true if they want to restart and false otherwise
func restartDungeon() bool{
	ans := ""
	fmt.Print("\nWould you like to restart the dungeon (Y/N)? ")
	fmt.Scanln(&ans)
	ans = strings.ToUpper(ans)

	for(ans != "Y" && ans != "N"){
		fmt.Print("Invalid input.\nWould you like to restart the dungeon (Y/N)? ")
		fmt.Scanln(&ans)
		ans = strings.ToUpper(ans)
	}
	fmt.Println()

	return ans == "Y"
}

//prompts the user if they would like to continue
//returns true if they want to continue and false otherwise
func continueAdventure() bool{
	ans := ""
	fmt.Print("\nWould you like to continue (Y/N)? ")
	fmt.Scanln(&ans)
	ans = strings.ToUpper(ans)

	for(ans != "Y" && ans != "N"){
		fmt.Print("Invalid input.\nWould you like to continue (Y/N)? ")
		fmt.Scanln(&ans)
		ans = strings.ToUpper(ans)
	}

	return ans == "Y"
}

//controls the traversal of all dungeons
//prompts the user if they would like to continue after each dungeon
//returns false if they would like to stop and true otherwise
func TraverseAllDungeons(dungeons []Dungeon, h *Hero, hil []HealingItem) bool{
	for i := 0; i < len(dungeons); i++{
		fmt.Println()
		completed := dungeons[i].DungeonTraversal(h, hil)

		if(!completed || !continueAdventure()){
			fmt.Println("\nThanks for playing.")
			return false
		}
	}
	return true
}