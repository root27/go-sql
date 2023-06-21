## Go Sql

---

### Description

This is simple package to connect sql database

### Usage Example

```go
package main

import (
    "fmt"
    "github.com/root27/go-sql"
)

func main() {
    
    db,err := sql_connection.ConnectDB("sql") // sql is the name in the .env file

    if err != nil {
        fmt.Println(err)
    }

	defer sql_connection.CloseDB(db) //Close db


    //Ping db

    sql_connection.PingDB(db)

    

   

	fmt.Println("Connected to DB")
}