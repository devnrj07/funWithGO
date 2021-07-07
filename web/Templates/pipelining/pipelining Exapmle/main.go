package main

import (
	"log"
	"math"
	"os"
	"text/template"
)

var tpl *template.Template

var funcMap = template.FuncMap{
	"fdouble": doubleIt,
	"fcube":   cubeIt,
	"fcqrt":   cubeqrtIt,
}

func init() {
	tpl = template.Must(template.New("").Funcs(funcMap).ParseFiles("index.gohtml"))
}

func doubleIt(x int) float64 {
	return float64(2 * x)
}

func cubeIt(x int) float64 {
	return math.Pow(float64(x), 3)
}

func cubeqrtIt(x int) float64 {
	return math.Cbrt(float64(x))
}

func main() {

	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", 6)

	if err != nil {
		log.Fatalln(err)
	}
}
