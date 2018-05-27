package main

import (
    "fmt"
    "log"
    "path/filepath"
    "bobbie/cmd"
    "bobbie/db"
    homedir "github.com/mitchellh/go-homedir"
)

func check(err error) {

    if err != nil {
        log.Fatal(err.Error())
    }
}

func main() {

    homeDir, _ := homedir.Dir()
    dbPath := filepath.Join(homeDir, "bobbie.db")
    err := db.Init(dbPath)

    check(err)

    fmt.Println("DB connection established.")

    cmd.Execute()
}
