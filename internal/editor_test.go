package internal

import "testing"

func TestEditor(t *testing.T) {
	bf := NewBuffer([]byte("0123456789"))
	bf.Insert(4, ",hola,")
	bf.Replace(6, 7, "go")
	want := "0123,hola,45go789"

	s := bf.String()
	if s != want {
		t.Errorf("got b.String() = %q, want %q", s, want)
	}
	sb := bf.Bytes()
	if string(sb) != want {
		t.Errorf("got b.Bytes() = %q, want %q", sb, want)
	}
}
