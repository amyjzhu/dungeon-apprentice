package main

import (
	"math/rand"
)

func randomizeOrder(party []Character, monsters []Monster) {
	allCreatures := append(party, monsters) // hm, type systems
	for i := range allCreatures {
		j := rand.Intn(i + 1)
		allCreatures[i], allCreatures[j] = allCreatures[j], allCreatures[i]
	}
}
