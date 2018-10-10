package chargen

import (
	"math/rand"

	"github.com/ironarachne/namegen"
)

// Character is a character
type Character struct {
	FirstName           string
	LastName            string
	HairColor           string
	HairStyle           string
	EyeColor            string
	FaceShape           string
	Height              int
	Weight              int
	Race                Race
	Gender              string
	Orientation         string
	Attitude            string
	Hobby               string
	PsychologicalTraits []string
	Motivation          string
}

// Couple is a pair of partners
type Couple struct {
	Partner1        Character
	Partner2        Character
	CanHaveChildren bool
}

// Family is a family of characters
type Family struct {
	FamilyName string
	Parents    Couple
	Children   []Character
}

func getOppositeGender(gender string) string {
	if gender == "male" {
		return "female"
	}

	return "male"
}

func getRaceFromParents(father Character, mother Character) Race {
	if father.Race.Name == mother.Race.Name {
		return father.Race
	}

	possibleRaces := []string{father.Race.Name, mother.Race.Name}

	if itemInCollection("elf", possibleRaces) && itemInCollection("human", possibleRaces) {
		return races["half-elf"]
	}

	if itemInCollection("dwarf", possibleRaces) && itemInCollection("human", possibleRaces) {
		return races["half-dwarf"]
	}

	if itemInCollection("half-elf", possibleRaces) {
		return races["half-elf"]
	}

	if itemInCollection("half-dwarf", possibleRaces) {
		return races["half-dwarf"]
	}

	return races["human"]
}

func randomOrientation() string {
	return randomItemFromThresholdMap(orientations)
}

func randomHeight(race Race, gender string) int {
	minHeight := race.MinHeight
	maxHeight := race.MaxHeight
	genderRatio := 1.0
	if gender == "male" {
		genderRatio = 1.08
	}

	height := rand.Intn(maxHeight-minHeight) + minHeight

	return int(float64(height) * genderRatio)
}

func randomRace() Race {
	raceName := randomItemFromThresholdMap(raceData)

	race := races[raceName]

	return race
}

func randomWeight(race Race, gender string) int {
	minWeight := race.MinWeight
	maxWeight := race.MaxWeight
	genderRatio := 1.0
	if gender == "male" {
		genderRatio = 1.16
	}

	weight := rand.Intn(maxWeight-minWeight) + minWeight

	return int(float64(weight) * genderRatio)
}

// Generate generates a random character
func Generate() Character {
	char := Character{}
	nameGenerator := namegen.NameGeneratorFromType("anglosaxon")

	char.Race = randomRace()
	char.FirstName = nameGenerator.FirstName()
	char.LastName = nameGenerator.LastName()
	char.HairColor = randomItemFromThresholdMap(char.Race.HairColors)
	char.HairStyle = randomItemFromThresholdMap(char.Race.HairStyles)
	char.EyeColor = randomItemFromThresholdMap(char.Race.EyeColors)
	char.FaceShape = randomItemFromThresholdMap(char.Race.FaceShapes)

	char.Gender = randomItem(genders)
	char.Orientation = randomOrientation()
	char.Attitude = randomItem(attitudes)
	char.Hobby = randomItem(hobbies)
	char.Motivation = randomItem(motivations)

	for i := 0; i < 2; i++ {
		trait := randomItem(traits)
		for itemInCollection(trait, char.PsychologicalTraits) {
			trait = randomItem(traits)
		}
		char.PsychologicalTraits = append(char.PsychologicalTraits, trait)
	}

	char.Height = randomHeight(char.Race, char.Gender)
	char.Weight = randomWeight(char.Race, char.Gender)

	return char
}

// GenerateCouple generates a couple
func GenerateCouple() Couple {
	char1 := Generate()
	char2 := Generate()
	canHaveChildren := false

	raceNames := []string{char1.Race.Name, char2.Race.Name}
	orientations := []string{char1.Orientation, char2.Orientation}

	if (char1.Orientation == "gay" && char2.Orientation == "straight") || (char1.Orientation == "straight" && char2.Orientation == "gay") {
		char2.Orientation = char1.Orientation
		orientations = []string{char1.Orientation}
	}

	if char1.Gender == char2.Gender && itemInCollection("straight", orientations) {
		char2.Gender = getOppositeGender(char1.Gender)
	} else if char1.Gender != char2.Gender && itemInCollection("gay", orientations) {
		char2.Gender = char1.Gender
	}

	if itemInCollection("dwarf", raceNames) && itemInCollection("elf", raceNames) {
		char1.Race = char2.Race
	}

	if itemInCollection("halfling", raceNames) {
		char1.Race = races["halfling"]
		char2.Race = races["halfling"]
	}

	if char1.Gender != char2.Gender {
		canHaveChildren = true
	}

	couple := Couple{char1, char2, canHaveChildren}

	return couple
}

// GenerateFamily generates a random family
func GenerateFamily() Family {
	child := Character{}
	children := []Character{}

	parents := GenerateCouple()

	parents.Partner1.LastName = parents.Partner2.LastName
	familyName := parents.Partner2.LastName

	if parents.CanHaveChildren {
		for i := 0; i < rand.Intn(6); i++ {
			child = Generate()
			child.LastName = familyName
			child.Race = getRaceFromParents(parents.Partner2, parents.Partner1)
			children = append(children, child)
		}
	}

	return Family{familyName, parents, children}
}
