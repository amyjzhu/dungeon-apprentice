package main

import (
	"math/rand"
)

func findInitialLineup(party []Character, numMonsters int, env Environment) {
	monsters := createEncounter(party, env);
}

func randomizeOrder(party []Character, monsters []Monster) {
	var allCreatures interface{} = append(party, monsters) // hm, type systems
	for i := range allCreatures {
		j := rand.Intn(i + 1)
		allCreatures[i], allCreatures[j] = allCreatures[j], allCreatures[i]
	}
}
