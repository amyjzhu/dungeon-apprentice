package main

var numChars int
var numEnemies int
var difficulty Difficulty
var party []Character

// TODO: move environment enum to better location


func createEncounter(numEnemiesDesired int, env Environment) []Monster {
	numEnemies = numEnemiesDesired

	challengeRating := getChallengeRating(difficulty, party)
	return getMonsters(env, challengeRating)
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
	length := len(party)
	sum := 0
	for i := 0; i < length; i++ {
		sum += party[i].level
	}

	return sum / length
}
