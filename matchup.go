package main

var numChars int
var numEnemies int
var difficulty Difficulty
var party []Character

// TODO: move environment enum to better location

type Environment int

const (
	arctic Environment = iota
	coastal
	desert
	forest
	grassland
	hill
	mountain
	swamp
	underdark
	underwater
	urban
)

type Difficulty int // encounter difficulty
const (
	easy Difficulty = iota
	medium
	hard
	deadly
)

type BaseStats struct {
	Strength int
	Charisma int
	Constitution int
	Wisdom int
	Intelligence int
	Dexterity int
}

type Saves struct { // statistics related to saving throws
	Strength_Save int
	Charisma_Save int
	Constitution_Save int
	Wisdom_Save int
	Intelligence_Save int
	Dexterity_Save int
}

type Character struct {
	name string
	level int
	stats BaseStats
}

func createEncounter(numEnemiesDesired int, env Environment) {
	numEnemies = numEnemiesDesired

	challengeRating := getChallengeRating(difficulty, party)
	getMonsters(env, challengeRating)
	// create a number of monsters...

}


func getChallengeRating(difficulty Difficulty, party []Character) int {
	// TODO: implement a curve more like that prescribed in handbook table
	// see http://slyflourish.com/5e_encounter_building.html

	numberRatio := numChars / numEnemies

	challengeRating := aggregateCharacterLevels(party) * numberRatio

	return challengeRating

}


// TODO: later be able to specify round up or down
func aggregateCharacterLevels(party []Character) int {
	len := len(party)
	sum := 0
	for i := 0; i < len; i++ {
		sum += party[i].level
	}

	return sum / len
}
