package main

import (
	"html/template"
	"os"
)

const tax = 6.75 / 100

type Product struct {
	Name  string
	Price float32
}

// func (p Product) PriceWithTax() float32 {
// 	return p.Price * (1 + tax)
// }

const templateString = `
{{"Item Information"}}

Name {{ .Name }}

Price {{ .Price }}

Price with tax: {{ calctax .Price | printf "$%.2f" }}
`

func main() {
	product := Product{
		Name:  "Lemonade",
		Price: 2.16,
	}

	fm := template.FuncMap{}
	fm["calctax"] = func(price float32) float32 {
		return price * (1 + tax)
	}

	t := template.Must(template.New("").Funcs(fm).Parse(templateString))
	t.Execute(os.Stdout, product)
}
