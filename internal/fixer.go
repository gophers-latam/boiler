package internal

import (
	"fmt"
	"go/parser"
	"go/token"
	"log"
	"path"
	"strconv"
	"strings"

	"golang.org/x/mod/modfile" // dependencia importante
)

// reescribir fuente para reemplazar srcMod con dstMod.
func fix(data []byte, file string, srcMod, dstMod string, isRoot bool) []byte {
	fset := token.NewFileSet()
	pfile, err := parser.ParseFile(fset, file, data, parser.ImportsOnly)
	CheckErr("análisis fuente módulo:\n%s", err)

	bf := NewBuffer(data)
	at := func(pos token.Pos) int {
		return fset.File(pos).Offset(pos)
	}

	srcName := path.Base(srcMod)
	dstName := path.Base(dstMod)
	if isRoot {
		if name := pfile.Name.Name; name == srcName || name == srcName+"_test" {
			dname := dstName + strings.TrimPrefix(name, srcName)
			if !token.IsIdentifier(dname) {
				msg := fmt.Sprintf("no se puede cambiar nombre del paquete %s a paquete %s", name, dname)
				log.Fatalf("%s: "+msg+": nombre de paquete inválido", file)
			}
			bf.Replace(at(pfile.Name.Pos()), at(pfile.Name.End()), dname)
		}
	}

	for _, spec := range pfile.Imports {
		path, err := strconv.Unquote(spec.Path.Value)
		if err != nil {
			continue
		}
		if path == srcMod {
			if srcName != dstName && spec.Name == nil {
				// mejorar insercion del editor
				bf.Insert(at(spec.Path.Pos()), srcName+" ")
			}
			// modificar import path a destino por default
			bf.Replace(at(spec.Path.Pos()), at(spec.Path.End()), strconv.Quote(dstMod))
		}
		if strings.HasPrefix(path, srcMod+"/") {
			// modificar import path a destino
			bf.Replace(at(spec.Path.Pos()), at(spec.Path.End()), strconv.Quote(strings.Replace(path, srcMod, dstMod, 1)))
		}
	}
	return bf.Bytes()
}

// reescribir contenido de go.mod para reemplazar srcMod con dstMod
func fixMod(data []byte, srcMod, dstMod string) []byte {
	f, err := modfile.ParseLax("go.mod", data, nil)
	CheckErr("análisis módulo fuente:\n%s", err)

	_ = f.AddModuleStmt(dstMod)
	new, err := f.Format()
	if err != nil {
		return data
	}
	return new
}

// logs helper
func CheckErr(msg string, err error) {
	if err != nil {
		if msg != "" {
			log.Fatalf(msg, err)
		} else {
			log.Fatal(err)
		}
	}
}
