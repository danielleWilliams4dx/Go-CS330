package rpg

import(
	"fmt"
	"strings"
)

//returns a list of the text at the beginning of the story,
//before the User enters their name
func CreateStory() []string{
	story := []string{
		"What’s this?\nPress enter to continue →",
		"There seems to be a commotion outside. →",
		"You leave your modest house and see the same troubling site.\nDry soil, no crops.\nOh where had the rain gone? →",
		"You can’t make a lot of money from farming these days.\nThe months of drought led you to a new occupation. →",
		"VILLAGER:\nHey Mercenary!--\nUm, what’s your name again? →",
	}
	return story
}

//adds the rest of the story to the story list using the Hero's name
func AddRestOfStory(story *[]string, name string){
	restOfStory := []string{
		"VILLAGER:\n" + name + "!\nPrincess Amina is missing! →",
		"VILLAGER:\nShe did not return from her journey to find a cure for this drought.\nWe are fearing the worst. →",
		"VILLAGER:\nYou must go to the haunted catacombs to find her! →",
		"The haunted catacombs is a labyrinth of dungeons filled with terrifying monsters.\nIt is also home to a powerful wizard who was long ago banished from the Kingdom of Ravensborough. →",
		"It is said that this wizard is the guardian of a powerful relic, the Crown of Vitality.\nPrincess Amina had promised that this crown would provide the means to nourish the entire kingdom. →",
		"You begin your journey into the haunted catacombs in search of the princess. →",
		"Entering the Wizard’s Abode.\n\nPRINCESS AMINA:\nHelp me!\nHe’s locked me up!\nPress enter to continue →",
		"You turn around and see the wizard blocking the entrance. →",
		"MARKO THE MISUNDERSTOOD:\nAnother thief, eh?\nYou must be stopped! →",
		"You take the key from the wizard’s robe and unlock Princess Amina’s cell. →",
		"She runs to remove an ornate ivory crown from a pedestal. →",
		"PRINCESS AMINA:\nYou’re a fool, you know that? →",
		"PRINCESS AMINA:\nThe Crown of Vitality is not a relic that will heal our land, but it will allow me to command any living vessel within my sight. →",
		"PRINCESS AMINA:\nYou townspeople can’t be trusted, you were going to betray me.\nIt started with thieves stealing food and gold from the castle.\nYou’re all so angry, I knew it was only time before you would try to take me out or try to replace me with someone who could fix things. →",
		"PRINCESS AMINA:\nI won’t let you!\nNow that I can control everyone, we can become the most powerful kingdom in all of the lands! →",
		"The princess placed the crown on her head and a bright green light shone through her eyes.\nMenacing creatures tore through the walls and surrounded you. →",
		"You swing your sword to knock the crown off of the princess’s head and restrain her.\nPress enter to continue →",
		"PRINCESS AMINA:\nNo!\nNo! NOOO!! →",
		"You help the wizard who is struggling to get up from the ground.\nYou return the crown to him. →",
		"MARKO THE MISUNDERSTOOD:\nThank you, kind mercenary. →",
		"MARKO THE MISUNDERSTOOD:\nI knew that she had evil intentions with that crown.\nI could see that her heart was tainted with darkness. →",
		"You return to the castle with the princess, in shackles, and send her to the dungeon. →",
		"To fill the power vacuum left by the imprisoned princess, the townspeople, grateful for your bravery and service, vote to have you as the new ruler. →",
		"As ruler, you revoke the wizard's banishment and invite him to live in the castle as the new high mage.\nYou also don him a cat to cure his crippling loneliness. →",
		"Marko casts a powerful spell for the kingdom and food becomes bountiful once again.\nPeople come from near and far to admire Marko’s magic and he becomes loved by all. →",
		"The End.",
	}

	*story = append(*story, restOfStory...)
}

//prompts the user to enter a name & returns that name
func EnterName() string{
	name := ""
	fmt.Print("\nPlease enter your name: ")
	fmt.Scanln(&name)
	return name
}

//outputs the story story chunk corresponding to the chunk name
func PrintStoryChunk(story []string, chunkName string){
	chunkToIndex := map[string]([]int){
		"intro" : {0,5},
		"afterName" : {5,11},
		"wizard" : {11,14},
		"rescue" : {14,21},
		"end" : {21, len(story)},
	}

	if(isValidKey(chunkName, chunkToIndex)){
		indices := chunkToIndex[chunkName]
		start := indices[0]
		end := indices[1]

		fmt.Println("\n------------------------------------------------------------------------------------------------------------------------------\n")
		for i := start; i < end; i++{
			fmt.Print(story[i] + " ")
			if(chunkName != "end" || i != len(story)-1){
				ans := ""
				fmt.Scanln(&ans)
			}else{
				fmt.Println()
			}
			fmt.Println()
		}
		fmt.Println("------------------------------------------------------------------------------------------------------------------------------")

	}else{
		fmt.Println("Chunk does not exist in the story.")
	}
}

//returns whether the map contains the key
func isValidKey(key string, m map[string]([]int)) bool{
	for k := range m{
		if(key == k){
			return true
		}
	}
	return false
}

//asks the user if they would like to skip the story
//returns true if they enter "Y" or "y" and false otherwise
func SkipStory() bool{
	ans := ""
	fmt.Print("\nWould you like to skip the story (Y/N)? ")
	fmt.Scanln(&ans)
	ans = strings.ToUpper(ans)

	for(ans != "Y" && ans != "N"){
		fmt.Print("Invalid input.\nWould you like to skip the story (Y/N)? ")
		fmt.Scanln(&ans)
		ans = strings.ToUpper(ans)
	}

	return ans == "Y"
}