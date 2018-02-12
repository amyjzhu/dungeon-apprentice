package main

var MAX_ENEMY_TYPES = 5 // TODO fix this arbitrary limit

func createEncounter(request Encounter) []Monster {
	challengeRating := getChallengeRating(request.Player, request.NumMonsters)
	monsters := getMonstersForEnvironment(request.Env)


	return getMonsters(request.Env, challengeRating)
	// create a number of monsters...
}

func getChallengeRating(party []Character, numEnemies int) int {
	// TODO: implement a curve more like that prescribed in handbook table
	// see http://slyflourish.com/5e_encounter_building.html

	numberRatio := len(party) / numEnemies

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

func filterByChallengeRating(candidates []Monster, cr int) {
	crMatches := make([]Monster, MAX_ENEMY_TYPES)
	for _, monster := range candidates {
		if monster.Challenge_Rating >= cr {
			crMatches = append(crMatches, monster)
		}
	}
}