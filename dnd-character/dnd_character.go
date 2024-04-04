package dndcharacter

import (
	"math"
	"math/rand"
)

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	floatScore := float64(score)
	return int(math.Floor((floatScore - 10) / 2))
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	sum, min := 0, 20
	for i := 0; i < 4; i++ {
		roll := rand.Intn(6) + 1
		sum += roll
		if roll < min {
			min = roll
		}
	}
	return sum - min
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	cons := Ability()
	return Character{
		Strength:     Ability(),
		Dexterity:    Ability(),
		Constitution: cons,
		Intelligence: Ability(),
		Wisdom:       Ability(),
		Charisma:     Ability(),
		Hitpoints:    10 + Modifier(cons),
	}
}
