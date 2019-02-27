package main

import (
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/ironarachne/chargen"
)

var (
	jsonTemplate = `
	{{ $traitsLength := (len .PsychologicalTraits) }}
	{
		"firstName": "{{ .FirstName }}",
		"lastName": "{{ .LastName }}",
		"race": "{{ .Race.Name }}",
		"age": {{ .Age }},
		"ageCategory": "{{ .AgeCategory }}",
		"gender": "{{ .Gender }}",
		"orientation": "{{ .Orientation }}",
		"profession": "{{ .Profession }}",
		"height": {{ .Height }},
		"weight": {{ .Weight }},
		"hairColor": "{{ .HairColor }}",
		"hairStyle": "{{ .HairStyle }}",
		"eyeColor": "{{ .EyeColor }}",
		"faceShape": "{{ .FaceShape }}",
		"motivation": "{{ .Motivation }}",
		"hobby": "{{ .Hobby }}",

		"traits": [
			{{ range $i, $t := .PsychologicalTraits }}
			"{{ $t }}"{{ if ne (add $i 1) $traitsLength }},{{ end }}
			{{ end }}
		]
	}
	`
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	rootHandler := func(w http.ResponseWriter, req *http.Request) {
		rand.Seed(time.Now().UnixNano())

		tmpl, err := template.New("webpage").Funcs(template.FuncMap{"add": func(a, b int) int { return a + b }}).Parse(jsonTemplate)
		check(err)

		character := chargen.GenerateCharacter()

		tmpl.Execute(w, character)
	}

	http.HandleFunc("/", rootHandler)
	log.Fatal(http.ListenAndServe(":9798", nil))
}
