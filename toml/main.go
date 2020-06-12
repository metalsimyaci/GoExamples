package main

import (
	. "./models"
	"database/sql"
	"fmt"
	"github.com/BurntSushi/toml"
	_ "github.com/lib/pq"
)

func main(){
	var (
		conf Config
		count int
	)
	if _,err := toml.DecodeFile("./configurations/config.toml",&conf);err !=nil{
		panic(err)
	}
	fmt.Printf("%v\n",conf)

	connectionString := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		conf.Database.User,conf.Database.Password, conf.Database.Server,conf.Database.Database)

	db,err := sql.Open("postgres",connectionString)
	if err != nil{
		fmt.Println(err.Error())
	}
	defer db.Close()

	err = db.QueryRow("SELECT 5+5").Scan(&count)
	if err != nil {
		panic(err)
	}
	fmt.Println(count)
}