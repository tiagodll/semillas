package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

func Migrate(migrations []Migration, sqler Sqler) {
	config := &Config{}
	config.Load()
	db, err := sql.Open(config.Db.Type, config.Db.ConnectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	for _, m := range migrations {
		fmt.Printf("%s", sqler.ToSql(&m))
		result, err := db.Exec(sqler.ToSql(&m))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v", result)
	}
}
