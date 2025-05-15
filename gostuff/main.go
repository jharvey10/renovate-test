package main

import (
	"fmt"

	"github.com/burningalchemist/sql_exporter"
)

func main() {
	exporter, err := sql_exporter.NewExporter("asdf")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(exporter.Config())
}
