package internal

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/mod/module" // dependencia importante
)

type (
	// guardar informacion de directorio inicial
	Info struct {
		Dir string
	}
	// registrar una modificación de texto: cambiar los bytes en [start,end] a new
	edit struct {
		start int
		end   int
		new   string
	}
	// lista de ediciones que se pueden ordenar por desplazamiento inicial
	edits []edit
)

// metodos edits
func (e edits) Len() int {
	return len(e)
}

func (e edits) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func (e edits) Less(i, j int) bool {
	if e[i].start != e[j].start {
		return e[i].start < e[j].start
	}
	return e[i].end < e[j].end
}

// Comprobar rutas a usar
func VerifyPath(path, modType string) {
	var msg string
	switch modType {
	case "src":
		msg = "nombre módulo origen no válido: %v"
	case "dst":
		msg = "nombre módulo destino no válido: %v"
	}
	err := module.CheckPath(path)
	Fatalf(msg, err)
}

// Copiar de caché del módulo en un nuevo directorio y realizar edicion.
func Edit(info *Info, dir, srcMod, dstMod string) error {
	err := filepath.WalkDir(info.Dir, func(src string, ent fs.DirEntry, err error) error {
		Fatal(err)

		rel, err := filepath.Rel(info.Dir, src)
		Fatal(err)
		dst := filepath.Join(dir, rel)
		if ent.IsDir() {
			err := os.MkdirAll(dst, 0o777)
			Fatal(err)
			return nil
		}

		data, err := os.ReadFile(src)
		Fatal(err)

		isRoot := !strings.Contains(rel, string(filepath.Separator))
		if strings.HasSuffix(rel, ".go") {
			data = fix(data, rel, srcMod, dstMod, isRoot)
		}
		if rel == "go.mod" {
			data = fixMod(data, srcMod, dstMod)
		}

		//#nosec gosec:G306
		err = os.WriteFile(dst, data, 0o666)
		Fatal(err)

		return nil
	})

	return err
}
