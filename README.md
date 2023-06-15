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
    url := sql_connection.LoadUri("sql") // sql is the name in the .env file

	db := sql_connection.ConnectDB(url)

	defer db.Close()

	fmt.Println("Connected to DB")
}