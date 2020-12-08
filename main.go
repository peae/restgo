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

}
