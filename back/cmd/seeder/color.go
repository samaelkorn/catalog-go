package main

type ColorSeed struct {
	Name string
	Code string
}

func (app *application) seederColors() {
	array := []ColorSeed{
		{Name: "Черный", Code: "black"},
		{Name: "Синий", Code: "blue"},
		{Name: "Зеленый", Code: "green"},
		{Name: "Желтый", Code: "yellow"},
	}

	for _, item := range array {
		app.db.InsertColor(item.Name, item.Code)
	}
}
