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
	Race                string
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

func getRaceFromParents(father Character, mother Character) string {
	if father.Race == mother.Race {
		return father.Race
	}

	races := []string{father.Race, mother.Race}

	if itemInCollection("elf", races) && itemInCollection("human", races) {
		return "half-elf"
	}

	if itemInCollection("dwarf", races) && itemInCollection("human", races) {
		return "half-dwarf"
	}

	if itemInCollection("dwarf", races) && itemInCollection("elf", races) {
		return "dwelf"
	}

	if itemInCollection("half-elf", races) {
		return "half-elf"
	}

	if itemInCollection("half-dwarf", races) {
		return "half-dwarf"
	}

	return "abomination"
}

func randomOrientation() string {
	thresholdBi := 5
	thresholdGay := 15

	result := rand.Intn(100)
	if result <= thresholdBi {
		return "bi"
	} else if result <= thresholdGay {
		return "gay"
	} else {
		return "straight"
	}
}

func randomHeight(race string, gender string) int {
	minHeight := races[race].MinHeight
	maxHeight := races[race].MaxHeight
	genderRatio := 1.0
	if gender == "male" {
		genderRatio = 1.08
	}

	height := rand.Intn(maxHeight-minHeight) + minHeight

	return int(float64(height) * genderRatio)
}

func randomRace() string {
	return randomItemFromThresholdMap(raceData)
}

func randomWeight(race string, gender string) int {
	minWeight := races[race].MinWeight
	maxWeight := races[race].MaxWeight
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
	char.HairColor = randomItemFromThresholdMap(races[char.Race].HairColors)
	char.HairStyle = randomItemFromThresholdMap(races[char.Race].HairStyles)
	char.EyeColor = randomItemFromThresholdMap(races[char.Race].EyeColors)
	char.FaceShape = randomItemFromThresholdMap(races[char.Race].FaceShapes)

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

	races := []string{char1.Race, char2.Race}
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

	if itemInCollection("dwarf", races) && itemInCollection("elf", races) {
		char1.Race = char2.Race
	}

	if itemInCollection("halfling", races) {
		char1.Race = "halfling"
		char2.Race = "halfling"
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
