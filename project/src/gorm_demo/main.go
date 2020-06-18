package main

import "gorm_demo/models"

func main() {
	models.CloseDB()
	
	models.Init()
	
	models.DB.CreateTable()
	
	
}



