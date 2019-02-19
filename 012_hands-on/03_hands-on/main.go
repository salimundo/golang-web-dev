package main

import (
	"log"
	"os"
	"text/template"
)

type hotel struct {
	Name, Address, City, Zip, Region string
}

type region struct {
	Region string
	Hotels []hotel
}

// Regions is a list of hotel regions in California
type Regions []region

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	hotels := Regions{
		region{
			Region: "Southern",
			Hotels: []hotel{
				hotel{
					Name:    "Los Angeles Hotel One",
					Address: "123 Los Angeles Blvd",
					City:    "Los Angeles",
					Zip:     "90210",
				},
				hotel{
					Name:    "San Diego Hotel One",
					Address: "123 San Diego Blvd",
					City:    "San Diego",
					Zip:     "97878",
				},
			},
		},
		region{
			Region: "Central",
			Hotels: []hotel{
				hotel{
					Name:    "El Centro Hotel One",
					Address: "123 El Centro Blvd",
					City:    "El Centro",
					Zip:     "90210",
				},
			},
		},
		region{
			Region: "Northern",
			Hotels: []hotel{
				hotel{
					Name:    "San Fran Hotel One",
					Address: "123 San Fran Blvd",
					City:    "San Fran",
					Zip:     "94210",
				},
				hotel{
					Name:    "Redwoods Hotel One",
					Address: "123 Redwoods Blvd",
					City:    "Redwoods",
					Zip:     "93421",
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, hotels)
	if err != nil {
		log.Fatalln(err)
	}
}
