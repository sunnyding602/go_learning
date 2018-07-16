package main

import (
	"html/template"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	type hotel struct {
		Name, Address, City, Zip, Region string
	}
	type region struct {
		Region string
		Hotels []hotel
	}
	type Regions []region

	h := Regions{
		region{
			Region: "northern",
			Hotels: []hotel{
				hotel{
					Name:    "H",
					Address: "4",
					City:    "L",
					Zip:     "91111",
					Region:  "4",
				},
				hotel{
					Name:    "B",
					Address: "4",
					City:    "L",
					Zip:     "91111",
					Region:  "4",
				},
			},
		},
		region{
			Region: "southren",
			Hotels: []hotel{
				hotel{
					Name:    "H",
					Address: "4",
					City:    "L",
					Zip:     "91111",
					Region:  "4",
				},
				hotel{
					Name:    "B",
					Address: "4",
					City:    "L",
					Zip:     "91111",
					Region:  "4",
				},
			},
		},
	}

	tpl.Execute(os.Stdout, h)
}
