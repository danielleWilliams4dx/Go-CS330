package main

import (
	"significantProject/rpg"
	"fmt"
)

func main(){
	// characterTest()
	// moveTest()
	// itemTest()
	// healingItemTest()
	// enemyTest()
	// slimeTest()
	// skeletonTest()
	// zombieTest()
	// dungeonTest()
	// getHealingItemTest()
	// createHeroTest()

	// attackTest()
	// reviveTest()
	// ppHealTest()
	// battleTest()
	// emptyInventoryTest()
	dungeonTraversalTest()
	// storyTest()
}

// create and print character list
func characterTest(){
	cf := rpg.CharacterFactory{"../csvFiles/characters.csv"}
	characterList := cf.CreateCharacterList()
	rpg.PrintAllCharacters(characterList)
}

//create and print move list
func moveTest(){
	mf := rpg.MoveFactory{"../csvFiles/moves.csv"}
	moveList := mf.CreateMoveList()
	rpg.PrintAllMoves(moveList)
}

//create and print item list
func itemTest(){
	itf := rpg.ItemFactory{"../csvFiles/items.csv"}
	itemList := itf.CreateItemList()
	rpg.PrintAllItems(itemList)
}

//create and print healing item list
func healingItemTest(){
	hif := rpg.HealingItemFactory{"../csvFiles/healingItems.csv"}
	healingItemList := hif.CreateHealingItemList()
	rpg.PrintAllHealingItems(healingItemList)
}

//create and print enemy list
func enemyTest(){
	mf := rpg.MoveFactory{"../csvFiles/moves.csv"}
	moveList := mf.CreateMoveList()
	ef := rpg.EnemyFactory{"../csvFiles/enemies.csv"}
	enemyList := ef.CreateEnemyList(moveList)
	rpg.PrintAllEnemies(enemyList)
}

//print all slimes
func slimeTest(){
	mf := rpg.MoveFactory{"../csvFiles/moves.csv"}
	moveList := mf.CreateMoveList()
	ef := rpg.EnemyFactory{"../csvFiles/enemies.csv"}
	enemyList := ef.CreateEnemyList(moveList)

	slimes := rpg.GetEnemyGroup(enemyList, "Slime")
	rpg.PrintAllEnemies(slimes)
}

//print all skeletons
func skeletonTest(){
	mf := rpg.MoveFactory{"../csvFiles/moves.csv"}
	moveList := mf.CreateMoveList()
	ef := rpg.EnemyFactory{"../csvFiles/enemies.csv"}
	enemyList := ef.CreateEnemyList(moveList)

	skeletons := rpg.GetEnemyGroup(enemyList, "Skeleton")
	rpg.PrintAllEnemies(skeletons)
}

//print all zombies
func zombieTest(){
	mf := rpg.MoveFactory{"../csvFiles/moves.csv"}
	moveList := mf.CreateMoveList()
	ef := rpg.EnemyFactory{"../csvFiles/enemies.csv"}
	enemyList := ef.CreateEnemyList(moveList)

	zombies := rpg.GetEnemyGroup(enemyList, "Zombie")
	rpg.PrintAllEnemies(zombies)
}

//print all dungeons (if debugging printing is uncommented in each function)
func dungeonTest(){
	mf := rpg.MoveFactory{"../csvFiles/moves.csv"}
	moveList := mf.CreateMoveList()
	ef := rpg.EnemyFactory{"../csvFiles/enemies.csv"}
	enemyList := ef.CreateEnemyList(moveList)
	itf := rpg.ItemFactory{"../csvFiles/items.csv"}
	itemList := itf.CreateItemList()
	hif := rpg.HealingItemFactory{"../csvFiles/healingItems.csv"}
	healingItemList := hif.CreateHealingItemList()

	rpg.BuildDungeon1(enemyList, itemList, healingItemList)
	rpg.BuildDungeon2(enemyList, itemList, healingItemList)
	rpg.BuildDungeon3(enemyList, itemList, healingItemList)
}

//testing getting healing items by rarity
func getHealingItemTest(){
	hif := rpg.HealingItemFactory{"../csvFiles/healingItems.csv"}
	healingItemList := hif.CreateHealingItemList()

	rpg.PrintAllHealingItems(rpg.GetAllHealingItemsByRarity(healingItemList,1))
	rpg.PrintAllHealingItems(rpg.GetAllHealingItemsByRarity(healingItemList,2))
	rpg.PrintAllHealingItems(rpg.GetAllHealingItemsByRarity(healingItemList,3))
	rpg.PrintAllHealingItems(rpg.GetAllHealingItemsByRarity(healingItemList,4))
	rpg.PrintAllHealingItems(rpg.GetAllHealingItemsByRarity(healingItemList,5))
}

//create and print the hero
func createHeroTest(){
	mf := rpg.MoveFactory{"../csvFiles/moves.csv"}
	moveList := mf.CreateMoveList()
	hif := rpg.HealingItemFactory{"../csvFiles/healingItems.csv"}
	healingItemList := hif.CreateHealingItemList()
	
	p1 := rpg.CreateHero("Danielle", moveList)
	for i := 0; i < 10; i++{
		p1.PickUpHealingItem(rpg.GetRandomHealingItem(healingItemList))
	}
	fmt.Println(p1.ToString())
}

//testing attack
func attackTest(){
	mf := rpg.MoveFactory{"../csvFiles/moves.csv"}
	moveList := mf.CreateMoveList()
	ef := rpg.EnemyFactory{"../csvFiles/enemies.csv"}
	enemyList := ef.CreateEnemyList(moveList)
	
	p1 := rpg.CreateHero("Danielle", moveList)
	op := rpg.GetRandomEnemy(enemyList)

	//before attack
	fmt.Println(p1.ToString())
	fmt.Println(op.ToString() + "\n")

	p1.Attack(&p1.GetMoveset()[2],&op)

	//after attack
	fmt.Println("\n" + p1.ToString())
	fmt.Println(op.ToString())
}

func reviveTest(){
	mf := rpg.MoveFactory{"../csvFiles/moves.csv"}
	moveList := mf.CreateMoveList()
	hif := rpg.HealingItemFactory{"../csvFiles/healingItems.csv"}
	healingItemList := hif.CreateHealingItemList()

	p1 := rpg.CreateHero("Danielle", moveList)
	p1.PickUpHealingItem(rpg.GetAllHealingItemsByRarity(healingItemList, 5)[0])
	p1.SetHP(0)
	fmt.Println(p1.ToString() + "\n")

	p1.UseRevive()
	fmt.Println("\n" + p1.ToString())
}

func ppHealTest(){
	mf := rpg.MoveFactory{"../csvFiles/moves.csv"}
	moveList := mf.CreateMoveList()
	ef := rpg.EnemyFactory{"../csvFiles/enemies.csv"}
	enemyList := ef.CreateEnemyList(moveList)
	hif := rpg.HealingItemFactory{"../csvFiles/healingItems.csv"}
	healingItemList := hif.CreateHealingItemList()
	
	p1 := rpg.CreateHero("Danielle", moveList)
	op := rpg.GetRandomEnemy(enemyList)
	p1.Attack(&p1.GetMoveset()[2],&op)
	p1.PickUpHealingItem(rpg.GetAllHealingItemsByRarity(healingItemList, 2)[0])
	p1.GetInventoryHealingItem()[0].ItemHeal(&p1)

	fmt.Println("\n" + p1.ToString())
}

func battleTest(){
	mf := rpg.MoveFactory{"../csvFiles/moves.csv"}
	moveList := mf.CreateMoveList()
	ef := rpg.EnemyFactory{"../csvFiles/enemies.csv"}
	enemyList := ef.CreateEnemyList(moveList)
	hif := rpg.HealingItemFactory{"../csvFiles/healingItems.csv"}
	healingItemList := hif.CreateHealingItemList()
	
	p1 := rpg.CreateHero("Danielle", moveList)
	p1.IncreaseExperience(200) //a levelup will be triggered when the battle is over
	for i := 0; i < 4; i++{
		p1.PickUpHealingItem(rpg.GetRandomHealingItem(healingItemList))
	}
	op := rpg.GetRandomEnemy(rpg.GetEnemyGroup(enemyList, "Slime"))

	// //Struggle mini test
	// p1.RemoveAllPP()

	rpg.Battle(&op, &p1)
}

func dungeonTraversalTest(){
	mf := rpg.MoveFactory{"../csvFiles/moves.csv"}
	moveList := mf.CreateMoveList()
	ef := rpg.EnemyFactory{"../csvFiles/enemies.csv"}
	enemyList := ef.CreateEnemyList(moveList)
	itf := rpg.ItemFactory{"../csvFiles/items.csv"}
	itemList:= itf.CreateItemList()
	hif := rpg.HealingItemFactory{"../csvFiles/healingItems.csv"}
	healingItemList := hif.CreateHealingItemList()
	p1 := rpg.CreateHero("Danielle", moveList)

	dungeon1 := rpg.BuildDungeon1(enemyList, itemList, healingItemList)
	dungeon1.DungeonTraversal(&p1, healingItemList)
}

func emptyInventoryTest(){
	mf := rpg.MoveFactory{"../csvFiles/moves.csv"}
	moveList := mf.CreateMoveList()
	hif := rpg.HealingItemFactory{"../csvFiles/healingItems.csv"}
	healingItemList := hif.CreateHealingItemList()
	itf := rpg.ItemFactory{"../csvFiles/items.csv"}
	itemList:= itf.CreateItemList()

	p1 := rpg.CreateHero("Danielle", moveList)
	p1.PickUpItem(itemList[0])
	p1.PickUpItem(itemList[1])
	p1.PickUpItem(itemList[2])
	for i := 0; i < 4; i++{
		p1.PickUpHealingItem(rpg.GetRandomHealingItem(healingItemList))
	}

	fmt.Println(p1.ToString() + "\n")
	p1.EmptyInventoryItem()
	fmt.Println(p1.ToString())
}

func storyTest(){
	story := rpg.CreateStory()
	// rpg.PrintStoryChunk(story, "intro")

	name := rpg.EnterName()
	rpg.AddRestOfStory(&story, name)

	rpg.PrintStoryChunk(story, "afterName")
	rpg.PrintStoryChunk(story, "rescue")
	rpg.PrintStoryChunk(story, "end")
}