package main

type StatusSeed struct {
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
