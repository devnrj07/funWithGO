package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

var funcMap = template.FuncMap{
	"formatDate": dateFormat,
	"fmtdate":    monthDayYear,
}

func init() {
	tpl = template.Must(template.New("").Funcs(funcMap).ParseFiles("index.gohtml"))
}

func dateFormat(t time.Time) string {
	return t.Format(time.Kitchen)
}

func monthDayYear(t time.Time) string {
	return t.Format("01-02-2006")
}

func main() {

	dateTime := time.Now()
	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", dateTime)

	if err != nil {
		log.Fatalln(err)
	}
}
