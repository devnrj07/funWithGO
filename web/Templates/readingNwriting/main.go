//we will parse template and write into it a composite literal, maps, struct, []struct, and more data structures

package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

type car struct {
	Model  string
	Brand  string
	Rating int
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	//composite literal
	people := []string{"abc", "xyz", "wtf"}

	err := tpl.ExecuteTemplate(os.Stdout, "compositeliteral.gohtml", people)
	if err != nil {
		log.Fatalln(err)
	}

	//map
	guitarist := map[string]string{
		"salsh":      "fender",
		"gilmour":    "twinguitar",
		"jimmy page": "yamaha",
	}

	err = tpl.ExecuteTemplate(os.Stdout, "maps.gohtml", guitarist)
	if err != nil {
		log.Fatalln(err)
	}

	//struct anonymous
	fiction := struct {
		Name  string
		FName string
	}{
		"van diesel",
		"doom",
	}

	err = tpl.ExecuteTemplate(os.Stdout, "structs.gohtml", fiction)
	if err != nil {
		log.Fatalln(err)
	}

	//slice of structs
	buddha := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	gandhi := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}

	mlk := sage{
		Name:  "Martin Luther King",
		Motto: "Hatred never ceases with hatred but with love alone is healed.",
	}

	jesus := sage{
		Name:  "Jesus",
		Motto: "Love all",
	}

	muhammad := sage{
		Name:  "Muhammad",
		Motto: "To overcome evil with good is good, to resist evil by evil is evil.",
	}

	//taking structs and creating a slice of them
	sages := []sage{buddha, gandhi, mlk, jesus, muhammad}

	err = tpl.ExecuteTemplate(os.Stdout, "stucts.gohtml", sages)

	if err != nil {
		log.Fatalln(err)
	}

	//struct-slices-struct
	x := car{
		"Model X",
		"Tesla",
		5,
	}

	y := car{
		"Harrier",
		"Tata",
		3,
	}

	z := car{
		"i8",
		"BMW",
		4,
	}

	cars := []car{x, y, z}

	collection := struct {
		Sages []sage
		Cars  []car
	}{
		sages,
		cars,
	}

	err = tpl.ExecuteTemplate(os.Stdout, "collection.gohtml", collection)

	if err != nil {
		log.Fatalln(err)
	}
}
