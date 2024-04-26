package main

type StatusSeed struct {
	Name string
	Code string
}

type ColorSeed struct {
	Name string
	Code string
}
type ProductSeed struct {
	Name     string
	Image    string
	Price    int
	ColorId  int
	StatusId int
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
func (app *application) seederProduct() {
	urlCont := app.config.nameContainer
	array := []ProductSeed{
		{Name: "Yamaha R1", Image: urlCont + "/storage/img/yamaha.jpg", Price: 5000, ColorId: 2, StatusId: 2},
		{Name: "BMW S1000RR", Image: urlCont + "/storage/img/bmw.jpg", Price: 6000, ColorId: 4, StatusId: 1},
	}

	for _, item := range array {
		app.db.InsertProduct(item.Name, item.Image, item.Price, item.ColorId, item.StatusId)
	}
}

func (app *application) seeder() {
	list, err := app.db.GetColors()
	if err == nil && len(list) == 0 {
		app.seederStatus()
		app.seederColors()
		app.seederProduct()
	}
}
