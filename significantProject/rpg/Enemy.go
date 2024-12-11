package rpg

import(
	"fmt"
	"strconv"
	"strings"
	"math/rand"
)

type Enemy struct{
	baseChar Character
	level int64
	maxHP int64 //max hp
	hp int64
	experience int64
	moveset []Move
	prefix string //to address the enemy in story
}

func (e Enemy) GetName() string{
	return e.baseChar.name
}

func (e Enemy) GetRole() string{
	return e.baseChar.role
}

func (e Enemy) GetHP() int64{
	return e.hp
}

func (e Enemy) GetMaxHP() int64{
	return e.maxHP
}

func (e *Enemy) SetHP(num int64){
	e.hp = num
}

func (e Enemy) GetExperience() int64{
	return e.experience
}

func (e Enemy) GetMoveset() []Move{
	return e.moveset
}

func (e Enemy) GetAllStats() string{
	return e.baseChar.ToString() + "\nLevel: " + fmt.Sprint(e.level) + "\nHP: " + fmt.Sprint(e.hp) + "/" + fmt.Sprint(e.maxHP) + "\n"
}

//don't include experience
func (e Enemy) ToString() string{
	return e.baseChar.ToString() + "\nLevel: " + fmt.Sprint(e.level) + "\nHP: " + fmt.Sprint(e.hp) + "/" + fmt.Sprint(e.maxHP) + "\n" + GetMovesetAsString(e.moveset)
}

//prints that the Hero encountered the Enemy
func (e Enemy) PrintEncounterMsg(h Hero) {
	fmt.Println(h.GetName() + " encountered " + e.prefix + " " + e.GetName())
}

//prints that the Enemy is defeated
func (e Enemy) Defeated(h Hero){
	fmt.Println("\n" + h.GetName() + " defeated " + e.prefix + " " + e.GetName())
}

//returns whether an Enemy is empty
//an Enemy is considered empty if it doesn't have a name
func (e Enemy) IsEmpty() bool{
	return e.GetName() == ""
}

//Updates the Hero's HP
//If the Hero's HP is 0, call GameOver method
func (e *Enemy) doDamage(damage int64, h *Hero){
	if(h.GetHP() - damage <= 0){
		h.SetHP(0)
	}else{
		h.SetHP(h.GetHP() - damage)
	}
}

//calculates whether the move hits
//calculates whether the move is a crit (damage x1.5)
//does damage and subtracts the Move's PP 
func (e Enemy) Attack(m *Move, h *Hero){
	m.DecreasePP()
	fmt.Println("— " + e.GetName() + " used " + m.GetName() + ". —")
	if(m.HitOrMiss()){
		if(m.IsCrit()){
			e.doDamage((m.GetDamage()*5)/4, h)
			fmt.Println("It was a critical hit!")
		}else{
			e.doDamage(m.GetDamage(), h)
		}
	}else{
		fmt.Println("The attack missed!")
	}
}

//calculates whether the move hits
//heals the user and subtracts the Move's PP 
func (e *Enemy) HealMove(m *Move){
	m.DecreasePP()
	fmt.Println("— " + e.GetName() + " used " + m.GetName() + ". —")
	if(m.HitOrMiss()){
		//if the move will restore more than the maxHP, max out the hp
		if(-1*m.GetDamage() + e.hp > e.maxHP){
			e.SetHP(e.maxHP)
			fmt.Println(e.GetName() + " restored its HP.")
		}else{
			e.SetHP(e.GetHP() + -1*m.GetDamage())
			fmt.Println(e.GetName() + " restored some of its HP.")
		}
	}else{
		fmt.Println("The move failed!")
	}
}

//the Enemy struggles if they have no attacking moves left
//it does a little damage to the Hero, and more damage to itself
func (e *Enemy) Struggle(h *Hero){
	e.doDamage(10, h)
	if(e.GetHP() - 25 > 0){
		e.SetHP(e.GetHP()-25)
	}else{
		e.SetHP(0)
	}
	fmt.Println(e.GetName() + " has no moves left!\n— " + e.GetName() + " used Struggle. —")
}

type EnemyFactory struct{
	EnemyFile string
}

//extracts data from csv records to create a Character list
func (ef EnemyFactory) CreateEnemyList(moveList []Move) []Enemy{
	enemyList := []Enemy{}
	records := GetRecords(ef.EnemyFile)
	for i, record := range records{
		name := strings.TrimSpace(record[0])
		role := strings.TrimSpace(record[1])
		level, error := strconv.ParseInt(record[2],10,64)
		PrintError(error)
		maxHP, error := strconv.ParseInt(record[3],10,64)
		PrintError(error)
		hp := maxHP
		experience, error := strconv.ParseInt(record[4],10,64)
		PrintError(error)
		moveset := BuildMoveset(record[5], moveList)
		prefix := ""
		if(len(records) > 6){
			prefix = strings.TrimSpace(record[6])
		}
		enemyList = append(enemyList, Enemy{Character{name,role}, level, maxHP, hp, experience, moveset, prefix})
		i += 0 //prevents declared and not used error
	}
	return enemyList
}

//prints all enemies in a list
func PrintAllEnemies(el []Enemy){
	fmt.Println("Enemies:\n")
	for i := 0; i < len(el); i++{
		fmt.Println(el[i].ToString() + "\n")
	}
}

//returns the enemy with the same name
func GetEnemyByName(el []Enemy, name string) Enemy{
	for i, e := range el{
		if(e.GetName() == name){
			return e
		}
		i += 0
	}
	fmt.Println("No enemies found named " + name)
	return Enemy{}
}

//returns all enemies with a role that contains the substring
//excludes standard bosses and possessed bosses
func GetEnemyGroup(el []Enemy, group string) []Enemy{
	enemyGroup := []Enemy{}
	for i, e := range el{
		if(!strings.Contains(e.GetRole(), "Possessed") && !strings.Contains(e.GetRole(), "Boss") && strings.Contains(e.GetRole(), group)){
			enemyGroup = append(enemyGroup, e)
		}
		i += 0
	}
	return enemyGroup
}

//returns a random Enemy from a group
func GetRandomEnemy(group []Enemy) Enemy{
	return group[rand.Intn(len(group))]
}

//restores the PP of all of the Enemies's moves
func (e *Enemy) FullPPHeal(){
	for i := 0; i < len(e.moveset); i++ {
		move := &e.moveset[i]
		move.RestorePP()
	}
}

//get rid of all PP
//for testing purposes only
func (e *Enemy) RemoveAllPP(){
	for i := 0; i < len(e.moveset); i++{
		move := &e.moveset[i]
		move.SetPP(0)
	}
}