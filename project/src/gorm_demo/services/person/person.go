package person

import "gorm_demo/models"

type UserService struct {
	Name string
	Age int
}



func (us *UserService) CreateTable(){
	models.DB.CreateTable()
}







