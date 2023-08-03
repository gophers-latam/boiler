package internal

import (
	"fmt"
	"sort"
)

// Una cola de ediciones para aplicar a un segmento de byte determinado
type Buffer struct {
	old   []byte
	queue edits
}

// metodos buffer

// Devolver un nuevo búfer para acumular cambios en un segmento de datos inicial
func NewBuffer(old []byte) *Buffer {
	return &Buffer{old: old}
}

// insertar el texto nuevo en old[pos:pos].
func (bf *Buffer) Insert(pos int, new string) {
	if pos < 0 || pos > len(bf.old) {
		panic("posición de edición no válida")
	}
	bf.queue = append(bf.queue, edit{pos, pos, new})
}

// borrar el texto old[start:end].
func (bf *Buffer) Delete(start, end int) {
	if end < start || start < 0 || end > len(bf.old) {
		panic("posición de edición no válida")
	}
	bf.queue = append(bf.queue, edit{start, end, ""})
}

// reemplazar old[start:end] por new.
func (bf *Buffer) Replace(start, end int, new string) {
	if end < start || start < 0 || end > len(bf.old) {
		panic("posición de edición no válida")
	}
	bf.queue = append(bf.queue, edit{start, end, new})
}

// obtener un nuevo segmento de bytes con los datos originales
// y las ediciones en cola aplicadas.
func (bf *Buffer) Bytes() []byte {
	// modificar por posición inicial y luego por posición final
	sort.Stable(bf.queue)

	var new []byte
	offset := 0
	for i, e := range bf.queue {
		if e.start < offset {
			e0 := bf.queue[i-1]
			panic(fmt.Sprintf(
				"ediciones superpuestas: [%d,%d)->%q, [%d,%d)->%q",
				e0.start, e0.end,
				e0.new, e.start,
				e.end, e.new,
			))
		}
		new = append(new, bf.old[offset:e.start]...)
		offset = e.end
		new = append(new, e.new...)
	}
	new = append(new, bf.old[offset:]...)
	return new
}

// devolver cadena de datos originales
// con ediciones en cola aplicadas.
func (bf *Buffer) String() string {
	return string(bf.Bytes())
}
