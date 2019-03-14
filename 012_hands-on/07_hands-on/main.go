package main

import (
	"log"
	"os"
	"text/template"
)

type menuItem struct {
	Name, Description string
	Price             float64
}

type courses struct {
	Breakfast, Lunch, Dinner []menuItem
}

type restaurant struct {
	Name    string
	Courses courses
}

var samsPlace = restaurant{
	Name: "Sam's Place",
	Courses: courses{
		Breakfast: []menuItem{
			{
				Name:        "Eggs and Bacon",
				Description: "Fluffy eggs and crispy bacon.",
				Price:       5.00,
			},
			{
				Name:        "Pancakes",
				Description: "Pancakes with maple syrup.",
				Price:       5.50,
			},
			{
				Name:        "French Toast",
				Description: "Cinnomon swirl french toast with rasins.",
				Price:       6.00,
			},
		},
		Lunch: []menuItem{
			{
				Name:        "Tuna Melt",
				Description: "Its tuna dude.",
				Price:       6.50,
			},
			{
				Name:        "BLT",
				Description: "Bacon, lettuce, and tomatoes on toast.",
				Price:       5.50,
			},
			{
				Name:        "Turkey Club",
				Description: "Turkey Club sandwich.",
				Price:       6.00,
			},
		},
		Dinner: []menuItem{
			{
				Name:        "Spagetti",
				Description: "Spagetti and red sauce",
				Price:       6.50,
			},
			{
				Name:        "Meatloaf",
				Description: "Just like mom made it.",
				Price:       5.50,
			},
			{
				Name:        "Chicken Pot Pie",
				Description: "Butter crust and gravy.",
				Price:       6.00,
			},
		},
	},
}

var eatAtJoes = restaurant{
	Name: "Eat at Joes",
	Courses: courses{
		Breakfast: []menuItem{
			{
				Name:        "Eggs and Bacon",
				Description: "Fluffy eggs and crispy bacon.",
				Price:       5.00,
			},
			{
				Name:        "Pancakes",
				Description: "Pancakes with maple syrup.",
				Price:       5.50,
			},
			{
				Name:        "French Toast",
				Description: "Cinnomon swirl french toast with rasins.",
				Price:       6.00,
			},
		},
		Lunch: []menuItem{
			{
				Name:        "Tuna Melt",
				Description: "Its tuna dude.",
				Price:       6.50,
			},
			{
				Name:        "BLT",
				Description: "Bacon, lettuce, and tomatoes on toast.",
				Price:       5.50,
			},
			{
				Name:        "Turkey Club",
				Description: "Turkey Club sandwich.",
				Price:       6.00,
			},
		},
		Dinner: []menuItem{
			{
				Name:        "Spagetti",
				Description: "Spagetti and red sauce",
				Price:       6.50,
			},
			{
				Name:        "Meatloaf",
				Description: "Just like mom made it.",
				Price:       5.50,
			},
			{
				Name:        "Chicken Pot Pie",
				Description: "Butter crust and gravy.",
				Price:       6.00,
			},
		},
	},
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	Restaurants := []restaurant{samsPlace, eatAtJoes}

	err := tpl.Execute(os.Stdout, Restaurants)
	if err != nil {
		log.Fatalln(err)
	}
}
