package chargen

import (
	"math"
	"math/rand"

	"github.com/ironarachne/utility"

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
	Age                 int
	AgeCategory         string
	Gender              string
	Orientation         string
	Profession          string
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

func getAppropriateName(gender string, race Race) (string, string) {
	nameCategory := "fantasy"

	if race.Name == "elf" {
		nameCategory = "elf"
	}
	nameGenerator := namegen.NameGeneratorFromType(nameCategory)

	firstName := nameGenerator.FirstName(gender)
	lastName := nameGenerator.LastName()

	return firstName, lastName
}

func getAgeCategoryRange(race Race, ageCategory string) (int, int) {
	minAge := 1
	maxAge := 1

	if ageCategory == "child" {
		minAge = 1
		maxAge = race.FertilityAge - 1
	} else if ageCategory == "young adult" {
		minAge = race.FertilityAge
		maxAge = race.AdultAge - 1
	} else if ageCategory == "adult" {
		minAge = race.AdultAge
		maxAge = int(math.Floor(float64(race.LifeSpan)*0.85)) - 1
	} else if ageCategory == "elderly" {
		minAge = int(math.Floor(float64(race.LifeSpan) * 0.85))
		maxAge = race.LifeSpan
	}

	return minAge, maxAge
}

func getRandomAge(race Race, ageCategory string) int {
	minAge, maxAge := getAgeCategoryRange(race, ageCategory)

	return rand.Intn(maxAge-minAge) + minAge
}

func getAgeCategoryFromAge(age int, race Race) string {
	ageCategory := ""
	minAge := 0
	maxAge := 0

	for category := range ageCategories {
		minAge, maxAge = getAgeCategoryRange(race, category)
		if age >= minAge && age <= maxAge {
			ageCategory = category
		}
	}

	return ageCategory
}

func getAgeFromParents(parents Couple, race Race) (int, string) {
	lowestAge := utility.Min(parents.Partner1.Age, parents.Partner2.Age)

	childAge := rand.Intn(lowestAge-race.FertilityAge) + 1
	childAgeCategory := getAgeCategoryFromAge(childAge, race)

	return childAge, childAgeCategory
}

func getRaceFromParents(parents Couple) Race {
	if parents.Partner1.Race.Name == parents.Partner2.Race.Name {
		return parents.Partner1.Race
	}

	possibleRaces := []string{parents.Partner1.Race.Name, parents.Partner2.Race.Name}

	if utility.ItemInCollection("elf", possibleRaces) && utility.ItemInCollection("human", possibleRaces) {
		return races["half-elf"]
	}

	if utility.ItemInCollection("dwarf", possibleRaces) && utility.ItemInCollection("human", possibleRaces) {
		return races["half-dwarf"]
	}

	if utility.ItemInCollection("half-elf", possibleRaces) {
		return races["half-elf"]
	}

	if utility.ItemInCollection("half-dwarf", possibleRaces) {
		return races["half-dwarf"]
	}

	return races["human"]
}

func randomOrientation() string {
	return utility.RandomItemFromThresholdMap(orientations)
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
	raceName := utility.RandomItemFromThresholdMap(raceData)

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

// GenerateCharacter generates a random character
func GenerateCharacter() Character {
	char := Character{}

	char.Race = randomRace()
	char.Gender = utility.RandomItem(genders)

	char.FirstName, char.LastName = getAppropriateName(char.Gender, char.Race)

	char.AgeCategory = utility.RandomItemFromThresholdMap(ageCategories)
	char.Age = getRandomAge(char.Race, char.AgeCategory)

	char.HairColor = utility.RandomItemFromThresholdMap(char.Race.HairColors)
	char.HairStyle = utility.RandomItemFromThresholdMap(char.Race.HairStyles)
	char.EyeColor = utility.RandomItemFromThresholdMap(char.Race.EyeColors)
	char.FaceShape = utility.RandomItemFromThresholdMap(char.Race.FaceShapes)

	char.Orientation = randomOrientation()
	char.Profession = utility.RandomItem(professions)
	char.Hobby = utility.RandomItem(hobbies)
	char.Motivation = utility.RandomItem(motivations)

	for i := 0; i < 2; i++ {
		trait := utility.RandomItem(traits)
		for utility.ItemInCollection(trait, char.PsychologicalTraits) {
			trait = utility.RandomItem(traits)
		}
		char.PsychologicalTraits = append(char.PsychologicalTraits, trait)
	}

	char.Height = randomHeight(char.Race, char.Gender)
	char.Weight = randomWeight(char.Race, char.Gender)

	return char
}

// GenerateCouple generates a couple
func GenerateCouple() Couple {
	char1 := GenerateCharacter()
	char2 := GenerateCharacter()
	canHaveChildren := false

	if char1.AgeCategory == "child" {
		char1.AgeCategory = "adult"
		char1.Age = getRandomAge(char1.Race, char1.AgeCategory)
	}

	if char2.AgeCategory == "child" {
		char2.AgeCategory = "adult"
		char2.Age = getRandomAge(char2.Race, char2.AgeCategory)
	}

	raceNames := []string{char1.Race.Name, char2.Race.Name}
	orientations := []string{char1.Orientation, char2.Orientation}

	if (char1.Orientation == "gay" && char2.Orientation == "straight") || (char1.Orientation == "straight" && char2.Orientation == "gay") {
		char2.Orientation = char1.Orientation
		orientations = []string{char1.Orientation}
	}

	if char1.Gender == char2.Gender && utility.ItemInCollection("straight", orientations) {
		char2.Gender = getOppositeGender(char1.Gender)
		char2.FirstName, _ = getAppropriateName(char2.Gender, char2.Race)
	} else if char1.Gender != char2.Gender && utility.ItemInCollection("gay", orientations) {
		char2.Gender = char1.Gender
		char2.FirstName, _ = getAppropriateName(char2.Gender, char2.Race)
	}

	if utility.ItemInCollection("dwarf", raceNames) && utility.ItemInCollection("elf", raceNames) {
		char1.Race = char2.Race
		char1.FirstName, _ = getAppropriateName(char1.Gender, char1.Race)
	}

	if utility.ItemInCollection("halfling", raceNames) {
		char1.Race = races["halfling"]
		char2.Race = races["halfling"]
	}

	if char1.Gender != char2.Gender {
		canHaveChildren = true
	}

	couple := Couple{char1, char2, canHaveChildren}

	return couple
}

// GenerateChild generates a child character for a couple
func GenerateChild(couple Couple) Character {
	child := GenerateCharacter()

	child.LastName = couple.Partner1.LastName
	child.Race = getRaceFromParents(couple)
	child.Age, child.AgeCategory = getAgeFromParents(couple, child.Race)

	if child.AgeCategory == "child" {
		child.Profession = "none"
	}

	return child
}

// GenerateFamily generates a random family
func GenerateFamily() Family {
	child := Character{}
	children := []Character{}

	parents := GenerateCouple()

	familyName := parents.Partner1.LastName

	parents.Partner2.LastName = familyName

	if parents.CanHaveChildren {
		for i := 0; i < rand.Intn(6); i++ {
			child = GenerateChild(parents)
			child.LastName = familyName
			children = append(children, child)
		}
	}

	return Family{familyName, parents, children}
}
