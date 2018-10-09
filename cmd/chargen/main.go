package main

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"

	"github.com/ironarachne/chargen"
)

func displayCharacter(char chargen.Character) {
	fmt.Println("\nName: " + char.FirstName + " " + char.LastName)
	fmt.Println("Hair: " + char.HairColor + ", " + char.HairStyle)
	fmt.Println("Eyes: " + char.EyeColor)
	fmt.Println("Face Shape: " + char.FaceShape)
	fmt.Println("Height: " + formatHeight(char.Height))
	fmt.Println("Weight: " + strconv.Itoa(char.Weight) + " lbs")
	fmt.Println("Race: " + char.Race)
	fmt.Println("Gender: " + char.Gender)
	fmt.Println("Orientation: " + char.Orientation)
	fmt.Println("Attitude: " + char.Attitude)
	fmt.Println("Traits: ")
	for _, trait := range char.PsychologicalTraits {
		fmt.Println("- " + trait)
	}
	fmt.Println("Hobby: " + char.Hobby)
	fmt.Println("Motivation: " + char.Motivation)
}

func displayFamily(family chargen.Family) {
	fmt.Println("# The " + family.FamilyName + " Family")
	fmt.Println("\n## Parents")
	displayCharacter(family.Parents.Partner1)
	displayCharacter(family.Parents.Partner2)

	if len(family.Children) > 0 {
		fmt.Println("\n## Children")
		for _, child := range family.Children {
			displayCharacter(child)
		}
	}
}

func formatHeight(height int) string {
	feet := int(math.Floor(float64(height) / 12.0))
	inches := int(math.Mod(float64(height), 12.0))

	return strconv.Itoa(feet) + "'" + strconv.Itoa(inches) + "\""
}

func main() {
	typeOfGeneration := flag.String("t", "individual", "Type to generate: individual or family")
	randomSeed := flag.Int64("s", 0, "Optional random generator seed")

	flag.Parse()

	if *randomSeed == 0 {
		rand.Seed(time.Now().UnixNano())
	} else {
		rand.Seed(*randomSeed)
	}

	if *typeOfGeneration == "individual" {
		character := chargen.Generate()
		displayCharacter(character)
	} else if *typeOfGeneration == "family" {
		family := chargen.GenerateFamily()
		displayFamily(family)
	}
}
