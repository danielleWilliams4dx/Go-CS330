package rpg

import(
	"fmt"
	"strings"
	"strconv"
	"regexp"
	"math/rand"
	"time"
)

//simulates a battle between the Enemy and Hero
//outputs both opponents stats between attacks/item usage
//if either opponent has no moves left, they struggle
//handles defeats, gave overs, and increasing EXP
//returns true if the Hero wins and false otherwise
func Battle(e *Enemy, h *Hero) bool{
	e.PrintEncounterMsg(*h)
	for(e.hp > 0){
		PrintOpponentStats(*e, *h)

		//game over breaks out of for loop
		continueBattle := heroTurnLogic(h,e)
		if(!continueBattle){
			return false
		}

		enemyTurnLogic(e, h)

		if(h.GetHP() == 0 && !h.UseRevive()){
			GameOver(*h)
			return false
		}
	}
	return true
}

//prompts the user whether they want to use a Move or an Item
//returns true if the user selects Move and false otherwise
//if the inventory is empty or the input is invalid ask for reinput
func moveOrItemOrRun(h Hero) string{
	ans := ""
	fmt.Print("FIGHT (F)\tINVENTORY (I)\tRUN (R) ")
	fmt.Scanln(&ans)
	ans = strings.ToUpper(ans)

	for((ans != "F" && ans != "I" && ans != "R")||(ans == "I" && len(h.inventoryHealingItem) == 0)){
		if(ans == "I" && len(h.inventoryHealingItem) == 0){
			fmt.Print("You hear a faint voice...\n" + h.GetName() + ", you have no useable healing items.\nFIGHT (F)\tINVENTORY (I)\tRUN (R) ")
		}else{
			fmt.Print("Invalid input.\nFIGHT (F)\tINVENTORY (I)\tRUN (R) ")
		}
		fmt.Scanln(&ans)
		ans = strings.ToUpper(ans)
	}

	return ans
}

//simulates the Hero's turn in battle
//if all Moves are out of PP, the Hero struggles to prevent soft-locking
//if the Hero runs out of HP and they don't have a revive, there's a game over
//returns false during a game over, otherwise returns true
func heroTurnLogic(h *Hero, e *Enemy) bool{
	ans := moveOrItemOrRun(*h)
	if(ans == "F"){
		if(!NoMovesLeft(h.GetMoveset())){
			return chooseHeroMove(h, e)
		}else{
			h.Struggle(e)
			if(h.GetHP() == 0 && !h.UseRevive()){
				GameOver(*h)
				return false
			}
		}
	}else if(ans == "I"){
		return h.chooseItem(e)
	}else{
		return !h.Run(*e)
	}
	return true
}

//simulates the Enemy's turn in battle
//if all Moves are out of PP, the Enemy struggles to prevent soft-locking
//if the Enemy runs out of HP, they are defeated and the Hero's experience increases
func enemyTurnLogic(e *Enemy, h *Hero){
	if(e.hp > 0){
		PrintOpponentStats(*e, *h)
		time.Sleep(time.Second) //wait a second
		if(!NoMovesLeft(e.GetMoveset())){
			chooseEnemyMove(e, h)
		}else{
			e.Struggle(h)
			if(e.hp == 0){
				e.Defeated(*h)
				h.IncreaseExperience(e.GetExperience())
			}
		}
	}else{
		e.Defeated(*h)
		h.IncreaseExperience(e.GetExperience())
	}
}

//prompts the user which Move they would like to restore
//returns that Move's pointer
func ChooseMoveRestore(h *Hero) *Move{
	//print moveset
	fmt.Print(GetMovesetAsString(h.moveset) + "\n\nWhich move would you like to restore (1, 2, etc.)? ")

	index := chooseMoveIndex(*h)
	return &h.moveset[index]
}

//randomly chooses a move from the Enemy's moveset and calls it
func chooseEnemyMove(e *Enemy, h *Hero){
	move := &e.moveset[rand.Intn(len(e.moveset))]

	if(move.GetCategory() == "Heal"){
		e.HealMove(move)
	}else{
		e.Attack(move, h)
	}
}

//prompts for the user to select a Move and calls it
func chooseHeroMove(h *Hero, e *Enemy) bool{
	//print moveset
	fmt.Print("\n" + GetMovesetAsString(h.moveset) + "\n\nGo Back (B)\nPlease select a move (1, 2, etc.) ")
	index := chooseMoveIndex(*h)

	//if -1 is returned, go back to choosing fight or inventory
	if(index == -1){
		return heroTurnLogic(h, e)
	}else{
		move := &h.moveset[index]

		if(move.GetCategory() == "Heal"){
			h.HealMove(move)
		}else{
			h.Attack(move, e)
		}
		return true
	}
}

//prompts the user to select a move and returns the index
func chooseMoveIndex(h Hero) int64{
	ans := ""
	var ansAsInt int64
	var e error
	fmt.Scanln(&ans)

	//if the user wants to go back, return -1
	if(strings.ToUpper(ans) == "B"){
		return -1
	}

	//check if the input is a number
	re := regexp.MustCompile(`[0-9]+`)
	req := re.FindString(ans) != ""
	if(req){
		ansAsInt, e = strconv.ParseInt(ans,10,64)
		PrintError(e)
	}

	//if the answer is not valid ask for reinput
	for(!req || int(ansAsInt) > len(h.moveset) || ansAsInt <= 0){
		fmt.Println("Invalid input for " + h.GetName() + "'s moveset.")
		fmt.Print("Go back (B)\nPlease select a move (1, 2, etc.) ")
		fmt.Scanln(&ans)

		//if the user wants to go back, return -1
		if(strings.ToUpper(ans) == "B"){
			return -1
		}

		req = re.FindString(ans) != ""
		if(req){
			ansAsInt, e = strconv.ParseInt(ans,10,64)
			PrintError(e)
		}
	}

	return ansAsInt-1
}

//prompts the user to select a move to replace and returns the index
func ChooseMoveReplaceIndex(h Hero) int64{
	ans := ""
	var ansAsInt int64
	var e error
	fmt.Scanln(&ans)

	//check if the input is a number
	re := regexp.MustCompile(`[0-9]+`)
	req := re.FindString(ans) != ""
	if(req){
		ansAsInt, e = strconv.ParseInt(ans,10,64)
		PrintError(e)
	}

	//if the answer is not valid ask for reinput
	for(!req || int(ansAsInt) > len(h.moveset) || ansAsInt <= 0){
		fmt.Println("Invalid input for " + h.GetName() + "'s moveset.")
		fmt.Print("\nWhat move would you like to replace (1, 2, etc.)? ")
		fmt.Scanln(&ans)

		req = re.FindString(ans) != ""
		if(req){
			ansAsInt, e = strconv.ParseInt(ans,10,64)
			PrintError(e)
		}
	}

	return ansAsInt-1
}

//prompts for the user to select an Item and uses it
func (h *Hero) chooseItem(e *Enemy) bool{
	//print inventory
	fmt.Print(h.GetInventoryAsString() + "\n\nGo back (B)\nPlease select an item (1, 2, etc.) ")
	index := chooseItemIndex(h)

	if(index == -1){
		return heroTurnLogic(h, e)
	}else{
		item := h.inventoryHealingItem[index]
		item.ItemHeal(h)
		return true
	}
}

//prompts the user to select an item and returns the index
//if the Item is not a HealingItem or the input is invalid, ask for reinput
func chooseItemIndex(h *Hero) int64{
	ans := ""
	var ansAsInt int64
	var e error
	fmt.Scanln(&ans)

	//if the user wants to go back, return -1
	if(strings.ToUpper(ans) == "B"){
		return -1
	}

	//check if the input is a number
	re := regexp.MustCompile(`[0-9]+`)
	req := re.FindString(ans) != ""
	if(req){
		ansAsInt, e = strconv.ParseInt(ans,10,64)
		PrintError(e)
	}

	//if the answer is not valid ask for reinput
	for(!req || int(ansAsInt) > len(h.inventoryHealingItem) + len(h.inventoryItem) || int(ansAsInt) <= len(h.inventoryItem)){
		if(int(ansAsInt) > 0 && int(ansAsInt) <= len(h.inventoryItem)){
			fmt.Print("You hear a faint voice...\n" + h.GetName() + ", now is not the time to use this item.")
		}
		fmt.Println("Invalid input for " + h.GetName() + "'s inventory")
		fmt.Print("Go back (B)\nPlease select an item (1, 2, etc.) ")
		fmt.Scanln(&ans)
		
		//if the user wants to go back, return -1
		if(strings.ToUpper(ans) == "B"){
			return -1
		}

		req = re.FindString(ans) != ""
		if(req){
			ansAsInt, e = strconv.ParseInt(ans,10,64)
			PrintError(e)
		}
	}

	return ansAsInt - int64(len(h.inventoryItem)) - 1
}

//prints the stats of the Enemy and the Hero
func PrintOpponentStats(e Enemy, h Hero){
	fmt.Println()
	fmt.Println(e.GetAllStats())
	fmt.Println(h.GetAllStats())
}

//outputs that the Hero has a Game Over
func GameOver(h Hero){
	fmt.Println("\n"+ h.GetName() + " fell.\n\nGAME OVER.")
}

//the battle between the Hero and the Wizard
func WizardBattle(h *Hero, el []Enemy) bool{
	wizard :=  GetEnemyByName(el,"Markos the Misunderstood")
	return standaloneBattle(h, &wizard)
}

//a battle outside of a Dungeon
//returns true if the Hero wins and false if they lose and give up
//each restart resets the enemy's and the Hero's hp & pp
func standaloneBattle(h *Hero, e *Enemy) bool{
	fmt.Println()
	win := Battle(e, h)

	for(!win && restartBattle()){
		h.SetHP(h.GetMaxHP())
		h.FullPPHeal()
		e.SetHP(e.GetMaxHP())
		e.FullPPHeal()
		win = Battle(e, h)
	}
	return win
}

//the final boss of the game, has 3 stages
//possessed skeleton, zombie, and then wizard
func FinalBoss(h *Hero, el []Enemy) bool{
	finalBosses := []Enemy{
		GetEnemyByName(el,"Sketchy Skeleton (Possessed)"),
		GetEnemyByName(el,"Undead Queen (Possessed)"),
		GetEnemyByName(el,"Markos the Misunderstood (Possessed)"),
	}

	for i := 2; i < len(finalBosses); i++{
		if(!standaloneBattle(h, &finalBosses[i])){
			return false
		}
	}
	return true
}

//prompts the user if they would like to restart the battle
//returns true if they want to restart and false otherwise
func restartBattle() bool{
	ans := ""
	fmt.Print("\nWould you like to restart the battle (Y/N)? ")
	fmt.Scanln(&ans)
	ans = strings.ToUpper(ans)

	for(ans != "Y" && ans != "N"){
		fmt.Print("Invalid input.\nWould you like to restart the battle (Y/N)? ")
		fmt.Scanln(&ans)
		ans = strings.ToUpper(ans)
	}
	fmt.Println()

	return ans == "Y"
}