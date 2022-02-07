package main

type Brand struct {
	Id int
	Name string
	Origin string
}

type Car struct {
	Id int `json:"id"`
	Name string `json:"name"`
	BrandId Brand
}



func (c Car) createCar() {
	//1st approach
	//newCar := Car{1,"March","Nissan"}
	//or
	//newCar := Car{Id:1,Name:"March",BrandName: "Nissan"}
}
