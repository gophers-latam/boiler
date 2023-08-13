package internal

import (
	"fmt"
	"log"
	"os"
)

const (
	cmdUse      = "boiler src.modName@version [dst.name/mod [dirName]]"
	aboutUse    = "https://github.com/gophers-latam/boiler"
	templateUse = `
	~ minimal: https://github.com/gophers-latam/minimal
	~ minimalchi: https://github.com/gophers-latam/minimalchi
	~ small-template: https://github.com/gophers-latam/small-template
	`
)

// Tabla salida de ayuda
func Help() {
	fmt.Fprintf(os.Stderr,
		"[uso]: %s \n\n[ver]: %s \n\n[plantillas]: %s \n",
		cmdUse, aboutUse, templateUse,
	)
	if os.Stderr != nil {
		log.Println("faltan argumentos")
	}
	os.Exit(2)
}
