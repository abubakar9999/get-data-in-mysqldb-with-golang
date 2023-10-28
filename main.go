package main

import (
	"database/sql"
	"encoding/json"
	f "fmt"

	model "example.com/mysql_connect/model"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:A@bu936943@tcp(localhost:3306)/testdb1") //connect db

	if err != nil {//error
		f.Println("err validatin sql arguments")
		panic(err.Error())
	}
	defer db.Close() //close db

	err = db.Ping()
	if err != nil {
		f.Println("error connection with db.ping")
		panic(err.Error())
	}

	rows, err := db.Query(`SELECT * FROM student`) //select db row--- 
	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()
 




	var result[] model.Mymodel// model
	
	
	
	if err = rows.Err(); err != nil {
		panic(err.Error())
	}
	for rows.Next(){ //get rows data  
	   var res model.Mymodel
		if err := rows.Scan(&res.Id, &res.Fname,&res.Lname); err != nil {// get rows info
            panic(err.Error())
        }
		f.Printf("ID: %d, Name: %s,LastName:%s\n", res.Id, res.Fname,res.Lname)

		result=append(result,res)

	}
	 
	jsonData, err := json.Marshal(result)
	if err != nil {
		panic(err.Error())
	}

	// Print JSON data
	f.Println(string(jsonData))



}
