package rpg

import (
	"fmt"
	"strconv"
	"strings"
	"math/rand"
)

type Move struct{
	name string
	category string
	damage int64
	maxPP int64
	pp int64
	accurracy int64
	levelLearned int64
}

func (m Move) GetName() string{
	return m.name
}

func (m Move) GetPP() int64{
	return m.pp
}

func (m *Move) SetPP(num int64){
	m.pp = num
}

func (m Move) GetDamage() int64{
	return m.damage
}

func (m Move) GetLevelLearned() int64{
	return m.levelLearned
}

func (m Move) GetCategory() string{
	return m.category
}

//decreases the Move's PP by 10
func (m *Move) DecreasePP(){
	if(m.pp > 0){
		m.pp --
	}
}

//restores the Move's PP
func (m *Move) RestorePP(){
	m.pp = m.maxPP
}

//returns whether the Move is critical
//rolls 1 in 20
func (m Move) IsCrit() bool{
	num := rand.Intn(20)
	// fmt.Println(fmt.Sprint(num))
	return num == 1
}  

//returns whether the Move hits
func (m Move) HitOrMiss() bool{
	num := int64(rand.Intn(100) + 1)
	// fmt.Println(fmt.Sprint(num))
	return num <= m.accurracy
}

//does not include levelLearned
func (m Move) ToString() string{
	return "Name: " + m.name + "\nCategory: " + m.category + "\nDamage: " + fmt.Sprint(m.damage) + "\nPP: " + fmt.Sprint(m.pp) + "/" + fmt.Sprint(m.maxPP) + "\nAccuracy: " + fmt.Sprint(m.accurracy)
}

//is tabbed
//does not include levelLearned
func (m Move) TabbedToString() string{
	return "\tName: " + m.name + "\n\tCategory: " + m.category + "\n\tDamage: " + fmt.Sprint(m.damage) + "\n\tPP: " + fmt.Sprint(m.pp) + "/" + fmt.Sprint(m.maxPP) + "\n\tAccuracy: " + fmt.Sprint(m.accurracy)
}

type MoveFactory struct{
	MoveFile string
}

//extracts data from csv records to create a move list
func (mf MoveFactory) CreateMoveList() []Move{
	moveList := []Move{}
	records := GetRecords(mf.MoveFile)
	for i, record := range records{
		name := strings.TrimSpace(record[0])
		category := strings.TrimSpace(record[1])
		damage, error := strconv.ParseInt(record[2],10,64)
		PrintError(error)
		maxPP, error := strconv.ParseInt(record[3],10,64)
		PrintError(error)
		pp := maxPP
		accurracy, error := strconv.ParseInt(record[4],10,64)
		PrintError(error)
		levelLearned, error := strconv.ParseInt(record[5],10,64)
		PrintError(error)
		moveList = append(moveList, Move{name, category, damage, maxPP, pp, accurracy, levelLearned})
		i += 0 //prevents declared and not used error
	}
	return moveList
}

//prints all moves
func PrintAllMoves(ml []Move){
	fmt.Println("Moves:\n")
	for i := 0; i < len(ml); i++{
		fmt.Println(ml[i].ToString() + "\n")
	}
}

//returns an enemy or hero's moveset as a string
func GetMovesetAsString(moveset []Move) string{
	var s string = "Moves:\n" 
	index := 1
	for i := 0; i < len(moveset); i++{
		s += "     " + fmt.Sprint(index) + "\n" + moveset[i].TabbedToString()
		if(i < len(moveset)-1){
			s += "\n\n"
		}
		index++
	}
	return s
}

//extracts the moveset from a string
//retrieves each Move by its name
func BuildMoveset(movesStr string, moveList []Move) []Move{
	moveset := []Move{}
	allMoves := strings.Split(movesStr[1:len(movesStr)-1],", ") //cuts of {} and separates string at commas
	
	for i, str := range allMoves{
		moveset = append(moveset, GetMoveFromName(str, moveList))
		i += 0
	}
	return moveset
}

//goes through the master list of moves and creates a map
//for learnable moves indexed by the level learned
func CompileLearnableMoves(ml []Move) map[int64]Move{
	lm := make(map[int64]Move)
	for i := 0; i < len(ml); i++{
		move := ml[i]
		level := move.GetLevelLearned()
		if(level > 0){
			lm[level] = move
		}
	}
	return lm
}

//returns the corresponding Move from its name
func GetMoveFromName(name string, moveList []Move) Move{
	for i, m := range moveList{
		if(name == m.GetName()){
			return m
		}
		i += 0
	}
	return Move{} //return empty Move if not found
}

//returns true if all of the Moves in a list are out of PP
func NoMovesLeft(moveList []Move) bool{
	for i, m := range moveList{
		if(m.GetPP() > 0){
			return false
		}
		i += 0
	}
	return true
}