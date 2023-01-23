package main

import (
	"fmt"
	"github.com/Sraik25/task/cmd"
	"github.com/Sraik25/task/db"
	"github.com/mitchellh/go-homedir"
	"os"
	"path/filepath"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	must(db.Init(dbPath))

	fmt.Println("db stuff worked!")
	must(cmd.RootCmd.Execute())
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
