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

// Family is a family of characters
type Family struct {
	Father   Character
	Mother   Character
	Children []Character
}

func raceFromParents(father Character, mother Character) string {
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
	minHeight := 0
	maxHeight := 1
	genderRatio := 1.0
	if gender == "male" {
		genderRatio = 1.08
	}

	if race == "human" {
		minHeight = 65
		maxHeight = 78
	} else if race == "elf" {
		minHeight = 60
		maxHeight = 70
	} else if race == "dwarf" {
		minHeight = 40
		maxHeight = 55
	}

	height := rand.Intn(maxHeight-minHeight) + minHeight

	return int(float64(height) * genderRatio)
}

func randomRace() string {
	return randomItemFromThresholdMap(raceData)
}

func randomWeight(race string, gender string) int {
	minWeight := 0
	maxWeight := 1
	genderRatio := 1.0
	if gender == "male" {
		genderRatio = 1.16
	}

	if race == "human" {
		minWeight = 140
		maxWeight = 230
	} else if race == "elf" {
		minWeight = 70
		maxWeight = 160
	} else if race == "dwarf" {
		minWeight = 180
		maxWeight = 300
	}

	weight := rand.Intn(maxWeight-minWeight) + minWeight

	return int(float64(weight) * genderRatio)
}

// Generate generates a random character
func Generate() Character {
	char := Character{}
	nameGenerator := namegen.NameGeneratorFromType("anglosaxon")

	char.FirstName = nameGenerator.FirstName()
	char.LastName = nameGenerator.LastName()
	char.HairColor = randomItem(hairColors)
	char.HairStyle = randomItem(hairStyles)
	char.EyeColor = randomItem(eyeColors)
	char.FaceShape = randomItem(faceShapes)
	char.Race = randomRace()
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

// GenerateFamily generates a random family
func GenerateFamily() Family {
	mother := Character{}
	father := Character{}
	child := Character{}
	children := []Character{}

	mother = Generate()
	mother.Gender = "female"

	father = Generate()
	father.Gender = "male"

	mother.LastName = father.LastName

	for i := 0; i < rand.Intn(6); i++ {
		child = Generate()
		child.LastName = father.LastName
		child.Race = raceFromParents(father, mother)
		children = append(children, child)
	}

	return Family{father, mother, children}
}
