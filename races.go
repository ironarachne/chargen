package chargen

var (
	raceData = map[string]int{
		"human":      30,
		"elf":        7,
		"dwarf":      5,
		"half-elf":   3,
		"half-dwarf": 1,
		"halfling":   3,
	}
	races = map[string]Race{
		"human": Race{
			"human",
			map[string]int{
				"black":  3,
				"brown":  2,
				"blonde": 2,
				"red":    1,
			},
			map[string]int{
				"shoulder-length": 2,
				"waist-length":    1,
				"bald":            1,
				"short":           2,
			},
			map[string]int{
				"blue":  2,
				"brown": 2,
				"amber": 1,
				"green": 2,
				"grey":  1,
			},
			map[string]int{
				"round":    2,
				"chiseled": 2,
				"narrow":   1,
				"broad":    1,
			},
			65,
			78,
			110,
			250,
			90,
			18,
			13,
		},
		"elf": Race{
			"elf",
			map[string]int{
				"black":  3,
				"brown":  2,
				"blonde": 2,
				"auburn": 1,
			},
			map[string]int{
				"shoulder-length": 2,
				"waist-length":    1,
				"topknot":         1,
				"short":           2,
			},
			map[string]int{
				"blue":   2,
				"brown":  2,
				"amber":  1,
				"green":  2,
				"grey":   1,
				"silver": 1,
			},
			map[string]int{
				"chiseled": 2,
				"narrow":   1,
				"oval":     1,
			},
			60,
			70,
			80,
			200,
			500,
			100,
			50,
		},
		"dwarf": Race{
			"dwarf",
			map[string]int{
				"black":  3,
				"brown":  2,
				"red":    2,
				"orange": 1,
			},
			map[string]int{
				"shoulder-length": 3,
				"waist-length":    2,
				"bald":            2,
				"short":           1,
			},
			map[string]int{
				"blue":  2,
				"brown": 3,
				"amber": 1,
				"green": 2,
			},
			map[string]int{
				"round":    2,
				"chiseled": 2,
				"broad":    1,
			},
			40,
			55,
			150,
			300,
			400,
			50,
			25,
		},
		"half-elf": Race{
			"half-elf",
			map[string]int{
				"black":  3,
				"brown":  2,
				"blonde": 2,
				"red":    1,
			},
			map[string]int{
				"shoulder-length": 2,
				"waist-length":    1,
				"bald":            1,
				"short":           2,
			},
			map[string]int{
				"blue":  2,
				"brown": 2,
				"amber": 1,
				"green": 2,
				"grey":  1,
			},
			map[string]int{
				"round":    2,
				"chiseled": 2,
				"narrow":   1,
			},
			65,
			78,
			100,
			250,
			200,
			18,
			13,
		},
		"half-dwarf": Race{
			"half-dwarf",
			map[string]int{
				"black":  3,
				"brown":  2,
				"blonde": 2,
				"red":    1,
			},
			map[string]int{
				"shoulder-length": 2,
				"waist-length":    1,
				"bald":            1,
				"short":           2,
			},
			map[string]int{
				"blue":  2,
				"brown": 2,
				"amber": 1,
				"green": 2,
				"grey":  1,
			},
			map[string]int{
				"round":    2,
				"chiseled": 2,
				"broad":    1,
			},
			60,
			70,
			120,
			250,
			150,
			18,
			13,
		},
		"halfling": Race{
			"halfling",
			map[string]int{
				"black": 3,
				"brown": 2,
			},
			map[string]int{
				"shoulder-length": 3,
				"waist-length":    2,
				"short":           1,
			},
			map[string]int{
				"blue":  2,
				"brown": 3,
				"amber": 1,
				"green": 2,
			},
			map[string]int{
				"round":    2,
				"chiseled": 2,
				"broad":    1,
			},
			40,
			55,
			60,
			200,
			150,
			20,
			14,
		},
	}
)

// Race is a character race
type Race struct {
	Name         string
	HairColors   map[string]int
	HairStyles   map[string]int
	EyeColors    map[string]int
	FaceShapes   map[string]int
	MinHeight    int
	MaxHeight    int
	MinWeight    int
	MaxWeight    int
	LifeSpan     int
	AdultAge     int
	FertilityAge int
}
