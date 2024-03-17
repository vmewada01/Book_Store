package config

import "gorm.io/gorm"


var (
	db *gorm.DB
)
// dsn := "root:password@tcp(127.0.0.1:3306)/learningGo?charset=utf8mb4&parseTime=True&loc=Local"

func Connect(){
	d, err :=  gorm.Open("mysql", "root:password/bookStore?charset=utf8mb4&parseTime=True&loc=Local")

	if err!= nil{
        panic(err)
    }

	db = d 

}

func GetDB() *gorm.DB{
return db 
}