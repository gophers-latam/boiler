package internal

import (
	"fmt"
	"os"
)

// Tabla salida de ayuda
func Help() {
	fmt.Fprintf(
		os.Stderr,
		"[uso]:\n boiler src.modName@version [dst.name/mod [dirName]] \n"+
			"\n[ver]: https://github.com/gophers-latam/boiler \n",
	)
	os.Exit(2)
}
