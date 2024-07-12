package main

import (
	"log"
	"os"
	"text/template"
)

var templateMain *template.Template
var templateAggregate *template.Template
var templateAggMap *template.Template
var templateDataAgg *template.Template

type userDetails struct {
	Name     string
	Hint     string
	Password string
}

type userJob struct {
	Position    string
	YearStarted int
	Location    string
}

type employee struct {
	Person   []userDetails
	Position []userJob
}

func init() {
	templateMain = template.Must(template.ParseFiles("templateMain.gohtml"))
	templateAggregate = template.Must(template.ParseFiles("templateAggregate.gohtml"))
	templateAggMap = template.Must(template.ParseFiles("templateAggMap.gohtml"))
	templateDataAgg = template.Must(template.ParseFiles("templateDataAgg.gohtml"))
}

func main() {
	// we are passing a single variable into the template
	// in Go you can only pass in one variable, but can pass in aggregates (lists, array, etc)
	errMain := templateMain.ExecuteTemplate(os.Stdout, "templateMain.gohtml", "Happy people look for things to be happy about.")
	if errMain != nil {
		log.Fatalln(errMain)
	}

	groceries := []string{"Apples", "Bananas", "Carrots"}
	errAgg := templateAggregate.ExecuteTemplate(os.Stdout, "templateAggregate.gohtml", groceries)
	if errAgg != nil {
		log.Fatalln(errAgg)
	}

	// passing in a map (list of; key:value)
	passwords := map[string]string{
		"Board Game":   "Jumunji",
		"Drink Bottle": "Yeti",
		"Desk":         "Matrix",
	}
	errMap := templateAggMap.ExecuteTemplate(os.Stdout, "templateAggMap.gohtml", passwords)
	if errMap != nil {
		log.Fatalln(errMap)
	}

	// a structure holding slices that contain structures
	james := userDetails{
		Name:     "James",
		Hint:     "Board Game",
		Password: "Jumunji",
	}

	laura := userDetails{
		Name:     "Laura",
		Hint:     "Drink Bottle",
		Password: "Yeti",
	}

	users := []userDetails{james, laura}

	salesManager := userJob{
		Position:    "Sales Manager",
		YearStarted: 2020,
		Location:    "Melbourne",
	}
	financeManager := userJob{
		Position:    "Finance Manager",
		YearStarted: 2023,
		Location:    "Sydney",
	}

	userPositions := []userJob{salesManager, financeManager}

	employees := employee{
		Person:   users,
		Position: userPositions,
	}

	errUsers := templateDataAgg.ExecuteTemplate(os.Stdout, "templateDataAgg.gohtml", employees)
	if errUsers != nil {
		log.Fatalln(errUsers)
	}
}
