package main

type StatusSeed struct {
	Name string
	Code string
}

type ColorSeed struct {
	Name string
	Code string
}

func (app *application) seederStatus() {
	array := []StatusSeed{
		{Name: "В наличии", Code: "stock"},
		{Name: "В пути", Code: "way"},
		{Name: "Под заказ", Code: "order"},
	}

	for _, item := range array {
		app.db.InsertStatus(item.Name, item.Code)
	}
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

func (app *application) seeder() {
	list, err := app.db.GetColors()
	if err == nil && len(list) == 0 {
		app.seederStatus()
		app.seederColors()
	}
}