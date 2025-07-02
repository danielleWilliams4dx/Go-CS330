# Significant Project Write-up
My signifcant project can be found [here](https://github.com/danielleWilliams4dx/Go-CS330/tree/main/significantProject).

**Play the game in a terminal by downloading the entire ```significantProject``` folder, and running the ```rpg.go``` file in the ```app``` folder.**

## UML Diagram
![RPG_UML_Labeled](https://github.com/user-attachments/assets/25237cc2-f608-43b6-a683-b99f751168d7)

## Program Description

### Overview

My program is a text-based RPG written in Go. It has a story, which is skippable; ```Dungeons```, which are labyrinths of ```Rooms``` containing ```Enemies```, ```Items```, and ```HealingItems```; and a sequence of final bosses at the end. As the game progresses, levels are scaled, and the ```Hero``` can learn more powerful ```Moves```. In order to keep the game at a reasonable difficulty, the ```Hero``` should limit the number of battles that they run away from.

### Simulated Inheritance

This project is composed of many isolated structures that communicate by sending messages to each other. Since Go values composition over inheritance, all **IS-A** relationships are simulated by **HAS-A** relationships. For example, both the ```Enemy``` and ```Hero``` structures embed a ```Character``` and inherit all of its attributes and methods. I made these attributes more accessible and seem like parent attributes by writing getter methods in both “children.”

### Structures

Factory structures are used to create all of the ```Enemies```, ```Moves```, ```Items```, and ```HealingItems```. As for compositional structures, ```Enemies``` and the ```Hero``` embed a ```Character``` and a ```moveset``` that is a slice of ```Moves```; the ```Hero``` also embeds slices of ```Items``` and ```HealingItems``` for the inventory; ```HealingItems``` embed an ```Item```; ```Rooms``` embed an ```enemyGroup``` that is a slice of ```Enemies```, an ```Item```, a ```HealingItem```, and pointers to ```Rooms``` that it is connected to; and ```Dungeons``` embed an entrance that is a pointer to a ```Room``` and an ```enemyGroup``` that is a slice of ```Enemies```.

### Other Files

In addition to structures, I wrote files that assist in reading from a CSV file (_CSVInteraction.go_); building and running the story (_Story.go_); simulate a battle and all of its mechanics including menu selection, Hero and Enemy turn logic, game overs, and boss fights (_Battle.go_).

### Key Functions

Some of my most important functions are ```runGame()``` (_rpg.go_), ```Battle(e *Enemy, h *Hero) bool``` (_rpg.go_), ```DungeonTraversal(h *Hero, hil []HealingItem) bool``` (_Dungeon.go_), and ```Enter(h* Hero, eg []Enemy, hil []HealingItem)``` (_Room.go_). 

#### ```runGame``` 
Instantiates all of the game elements by calling factory structure methods, controls story skipping, calls all dungeon traversals and standalone battles, and allows the user to choose whether to stop after each dungeon and battle loss. 

#### ```Battle```
Simulates a battle between the ```Hero``` and an ```Enemy```, outputting encounter messages and opponent stats between turns and handling defeats and game overs such that losing or running away returns ```false``` and winning returns ```true```. 

#### ```DungeonTraversal```
Is a recursive method that simulates a dungeon traversal, starting at an entrance, allowing the ```Hero``` to restart if they game over, and returning whether the dungeon was completed. 

#### ```Enter```
Is a recursive method that performs the inner workings of each ```DungeonTraversal```. When a ```Room``` is entered, its ```visited``` attribute is set to true, a battle is started if there is an ```Enemy```, the ```Hero``` picks up an ```Item``` or ```HealingItem``` if there is one, and the user can pick a direction to go and ```Enter``` that ```Room```.
