package main

import (
	"restgo/database"
	"restgo/routers"
	"fmt"
)

//var db *gorm.DB



func main() {

	//if err := config.Load("config/config.yaml"); err != nil {
	//	fmt.Println("加载配置失败")
	//	return
	//}

	_, err := database.InitPostgre()
	if err != nil {
		fmt.Println("数据库连接失败")
		return
	}

	router := routers.Routers()
	router.Run(":8081")


	////defer db.Close()
	//dsn := "host=192.168.25.130 user=postgres password=123456 dbname=crudapi port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	//db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	////gorm.Open("postgres", "host=192.168.25.130 user=postgres password=123456 dbname=crudapi port=5432 sslmode=disable TimeZone=Asia/Shanghai")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//
	//db.AutoMigrate(&User{})
	//
	//
	//user := User{
	//	Id: 5,
	//	Name: "aa",
	//}
	//db.Create(&user)
	//db.First(&user)
	//fmt.Println(user)
	////defer db.Close()
	//
	//db.First(&user, 1)
	//fmt.Println(user)



	//r := Routers()
	//
	//
	//
	//r.Run(":8085")

}
