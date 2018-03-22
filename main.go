package main

import (
	"fmt"
	"os"

	"github.com/ameech/copyql/cmd"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
