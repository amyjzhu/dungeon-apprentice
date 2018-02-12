package main

import (
	"strconv"
	"log"
	"encoding/json"
)

// TODO these probably don't have to be exportabl
type Encounter struct {
	Player []Character
	Npc []Character
	Enemy []Monster
	Env Environment
	NumMonsters int
	Difficulty Difficulty
	Condition []Condition
}

func (encounter *Encounter) getString() string {
	b, err := json.Marshal(encounter)
	if err != nil {
		log.Fatal("Couldn't serialize encounter!")
	}

	return string(b[:])
}

type Condition struct {

}

// TODO abstract implementing function pointers? too annoying?
func greaterThanEqualChallengeRating(cr string, otherCr string) bool {
	crInt, err := strconv.Atoi(cr);
	otherCrInt, err := strconv.Atoi(otherCr)

	if err != nil {
		log.Fatal("challenge ratings are non-numeric")
	}

	return crInt >= otherCrInt
}


type Monster struct {
	Name string
	Alignment string
	Challenge_Rating string // challenge rating
	env []Environment // places usually found
	exp int // experience points award for defeat
	Hit_Dice DamageRoll
	Base_Stats BaseStats
	Skills Skills
	//Saves Saves
	Damage_Vulnerabilities string
	Damage_Immunities string
	Special_Abilities []Action
	Legendary_Actions []Action
	Actions[] Action
}

type Action struct {
	Name string
	Desc string
	Attack_Bonus int
	Damage_Dice DamageRoll
	Damage_Bonus int
}

type Skills struct {
	// there has to be a better way to do this?
	// at least with a java class, enum constructors...
	Athletics int
	Acrobatics int
	Sleight_of_Hand int
	Stealth int
	Arcana int
	History int
	Investigation int
	Nature int
	Religion int
	Animal_Handling int
	Insight int
	Medicine int
	Perception int
	Survival int
	Deception int
	Intimidation int
	Performance int
	Persuasion int
}



type DamageRoll struct {
	rolls int
	die int
}

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

var Environments = []Environment {
	arctic, coastal, desert, forest, grassland,
	hill, mountain, swamp, underdark, underwater, urban}

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

