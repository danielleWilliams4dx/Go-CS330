package main

import (
	"significantProject/rpg"
	"fmt"
)

func main(){
	runGame()
}

//creates all used game elements,
//outputs the story if not skipped,
//calls all dungeon traversals and standalone battles,
//allows the user choose whether to stop after each dungeon and battle loss
func runGame(){
	mf := rpg.MoveFactory{"../csvFiles/moves.csv"}
	moveList := mf.CreateMoveList()

	ef := rpg.EnemyFactory{"../csvFiles/enemies.csv"}
	enemyList := ef.CreateEnemyList(moveList)

	itf := rpg.ItemFactory{"../csvFiles/items.csv"}
	itemList:= itf.CreateItemList()

	hif := rpg.HealingItemFactory{"../csvFiles/healingItems.csv"}
	healingItemList := hif.CreateHealingItemList()

	dungeon1 := rpg.BuildDungeon1(enemyList, itemList, healingItemList)
	dungeon2 := rpg.BuildDungeon2(enemyList, itemList, healingItemList)
	dungeon3 := rpg.BuildDungeon3(enemyList, itemList, healingItemList)
	dungeons := []rpg.Dungeon{dungeon1, dungeon2, dungeon3}

	story := []string{}
	skipStory := rpg.SkipStory()

	if(!skipStory){
		story = rpg.CreateStory()
		rpg.PrintStoryChunk(story, "intro")
	}

	name := rpg.EnterName()
	p1 := rpg.CreateHero(name, moveList)

	if(!skipStory){
		rpg.AddRestOfStory(&story, name)
		rpg.PrintStoryChunk(story, "afterName")
	}

	continueAdventure := rpg.TraverseAllDungeons(dungeons, &p1, healingItemList)

	if(continueAdventure){
		if(!skipStory){
			rpg.PrintStoryChunk(story, "wizard")
		}else{
			fmt.Println("\nEntering the Wizard’s Abode.\n")
		}
		
		if(!rpg.WizardBattle(&p1, enemyList)){
			return
		}

		if(!skipStory){
			rpg.PrintStoryChunk(story, "rescue")
		}

		if(!rpg.FinalBoss(&p1, enemyList)){
			return
		}

		if(!skipStory){
			rpg.PrintStoryChunk(story, "end")
		}else{
			fmt.Println("You Win!\nThanks for playing.")
		}
	}	
}