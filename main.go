// Based on gonew experimental implementation
package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"github.com/gophers-latam/boiler/internal"
)

func main() {
	log.SetPrefix("boiler: ")
	log.SetFlags(0)
	flag.Usage = internal.Help
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 || len(args) > 3 {
		// Mostrar salida de ayuda
		internal.Help()
	}

	// Comprobar fuente
	srcMod := args[0]
	srcModVers := srcMod
	if !strings.Contains(srcModVers, "@") {
		srcModVers += "@latest"
	}
	srcMod, _, _ = strings.Cut(srcMod, "@")
	internal.VerifyPath(srcMod, "src")

	// Comprobar destino
	dstMod := srcMod
	if len(args) >= 2 {
		dstMod = args[1]
		internal.VerifyPath(dstMod, "dst")
	}

	// Comprobar directorio en ruta
	var dir string
	if len(args) == 3 {
		dir = args[2]
	} else {
		dir = "." + string(filepath.Separator) + path.Base(dstMod)
	}

	// "dir" no debe existir o debe ser un directorio vacío.
	dest, err := os.ReadDir(dir)
	if err == nil && len(dest) > 0 {
		msg := fmt.Sprintf("directorio destino %s existe y no está vacío", dir)
		internal.CheckErr(msg+"\n%v", err)
	}
	needMkdir := err != nil

	// Descargar fuente
	var stdOut, stdErr bytes.Buffer
	cmd := exec.Command("go", "mod", "download", "-json", srcModVers)
	cmd.Stdout, cmd.Stderr = &stdOut, &stdErr
	if err := cmd.Run(); err != nil {
		msg := fmt.Sprintf("go mod download -json %s: %v\n%s%s", srcModVers, err, stdErr.Bytes(), stdOut.Bytes())
		internal.CheckErr(msg+"\n%v", err)
	}

	info := &internal.Info{}
	if err := json.Unmarshal(stdOut.Bytes(), &info); err != nil {
		msg := fmt.Sprintf("go mod download -json %s: salida JSON no válida: %v\n%s%s", srcMod, err, stdErr.Bytes(), stdOut.Bytes())
		internal.CheckErr(msg+"\n%v", err)
	}

	if needMkdir {
		err := os.MkdirAll(dir, 0o777)
		internal.CheckErr("", err)
	}

	// Editar destino
	err = internal.Edit(info, dir, srcMod, dstMod)
	internal.CheckErr("", err)

	// Salida exitosa
	log.Printf("Se inicializó: %s en %s", dstMod, dir)
}
