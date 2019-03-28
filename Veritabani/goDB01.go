package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strconv"
)

func main()  {
	db, err := sql.Open("mysql","root:@tcp(127.0.0.1:3306)/demodb")
	if err !=nil{
		panic(err.Error())
	}
	defer db.Close()


	/*createStatement := "CREATE TABLE `users` (`ID` int(11) NOT NULL AUTO_INCREMENT, `Username` varchar(45) NOT NULL, `Email` varchar(45) NOT NULL, `FirstName` varchar(45) NOT NULL, `LastName` varchar(45) NOT NULL, `BirthDate` varchar(45) NOT NULL, `IsActive` varchar(45) NOT NULL, PRIMARY KEY (`ID`), UNIQUE KEY `ID_UNIQUE` (`ID`)) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;"
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS "+ createStatement)
	if err != nil{
		log.Fatal(err)
	}*/

	res, err := db.Exec("INSERT INTO `users` (`Username`, `Email`, `FirstName`, `LastName`, `BirthDate`, `IsActive`) VALUES ( 'hasan.ural', 'info@hasanural.com', 'Hasan', 'URAL', '2017.1.1', '1')")
	if err != nil{
		log.Fatal(err)
	}

	rowCount, err := res.RowsAffected()
	if err != nil{
		log.Fatal(err)
	}

	log.Printf("Inserted %d rows",rowCount)

	lastId, err := res.LastInsertId()
	if err != nil{
		log.Fatal(err)
	}

	log.Printf("Last Inserted Id: %d ",lastId)

	var(
		ID			int
		Username	string
		Email		string
		FirstName	string
		LastName	string
		BirthDate	string
		IsActive	bool
	)
	rows, err := db.Query("SELECT * from users")
	if err != nil{
		log.Fatal(err)
	}

	for rows.Next(){
		err = rows.Scan(&ID,&Username,&Email,&FirstName,&LastName, &BirthDate, &IsActive)
		if err != nil{
			log.Fatal(err)
		}
		log.Printf("Row : %q",strconv.Itoa(ID)+" "+Username+" "+Email+" "+FirstName+" "+LastName+" "+BirthDate+" "+strconv.FormatBool(IsActive))

		/*
		//Aletratif Kullanım
		if err = rows.Err(); err != nil{
			log.Fatal(err)
		}
		*/
	}

	res, derr := db.Exec("DELETE from users where ID=?",5)
	
}

