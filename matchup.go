package main

import (
	"strconv"
	"math/rand"
)

var MAX_ENEMY_TYPES = 5 // TODO fix this arbitrary limit

// input argument
func createEncounter(request Encounter) []Monster {
	challengeRating := getChallengeRating(request.Player, request.NumMonsters)
	monsters := getMonstersForEnvironment(request.Env)
	monsters = filterByChallengeRating(monsters, challengeRating)
	scrambleEncounter(request, monsters)

	return monsters
}

func getChallengeRating(party []Character, numEnemies int) int {
	// TODO: implement a curve more like that prescribed in handbook table
	// see http://slyflourish.com/5e_encounter_building.html

	numberRatio := len(party) / numEnemies

	challengeRating := aggregateCharacterLevels(party) * numberRatio // TODO convert to string lol

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

func filterByChallengeRating(candidates []Monster, cr int) []Monster {
	crString := strconv.Itoa(cr)
	crMatches := make([]Monster, MAX_ENEMY_TYPES)
	for _, monster := range candidates {
		if greaterThanEqualChallengeRating(monster.Challenge_Rating, crString) {
			crMatches = append(crMatches, monster)
		}
	}

	return crMatches
}

func scrambleEncounter(request Encounter, candidates []Monster) {
	// alters the encounter
	// TODO implement gaussian distribution
	enemyVarieties := 1
	rand.Seed(100)
	random := rand.Int();
	switch true {
	case random > 70:
		enemyVarieties = 2
	case random > 90:
		enemyVarieties = 3
	}

	enemiesLeftToGenerate := request.NumMonsters
	enemiesChosen := make([]Monster, enemiesLeftToGenerate)

	for ; enemyVarieties > 0; enemyVarieties-- {
		random = rand.Int() % len(candidates)
		selection := candidates[random]

		// decide how many monsters to add
		// if this is the last type of monster, it must finish up
		// the encounter
		random = rand.Int() % enemiesLeftToGenerate
		if enemyVarieties == 1 {
			random = enemiesLeftToGenerate
		}

		// add enemies to roster
		for random > 0 {
			enemiesChosen = append(enemiesChosen, selection)
		}
		enemiesLeftToGenerate -= random
	}

	request.Enemy = enemiesChosen
}
