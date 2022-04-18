package cipher

import (
	"strings"
)

type Cipher interface {
	Encode(string) string
	Decode(string) string
}

//-----------------------------Caesar-----------------------------

type Caesar struct {
}

func NewCaesar() Cipher {
	return Caesar{}
}

func (c Caesar) Encode(s string) string {
	m := []int32(s)
	var buff []int32
	for i, _ := range s {
		if m[i] > 64 && m[i] < 91 {
			buff = append(buff, m[i]+32)
		} else if m[i] > 96 && m[i] < 123 {
			buff = append(buff, m[i])
		}
	}

	for i, _ := range buff {
		if buff[i]+3 > 122 {
			buff[i] = buff[i] - 23
		} else {
			buff[i] = buff[i] + 3
		}
	}

	return string(buff)
}

func (c Caesar) Decode(s string) string {
	m := []int32(s)
	for i, _ := range s {
		if m[i]-3 < 97 {
			m[i] = m[i] + 23
		} else {
			m[i] = m[i] - 3
		}
	}

	return string(m)
}

//-------------------------------------Shift-----------------------------------

type Shift struct {
	I int
}

func Module(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func NewShift(i int) Cipher {
	if Module(i) < 25 && i != 0 {
		return Shift{i}
	}

	return nil
}

func (c Shift) Encode(s string) string {
	m := []int32(s)
	var buff []int32
	for i, _ := range s {
		if m[i] > 64 && m[i] < 91 {
			buff = append(buff, m[i]+32)
		} else if m[i] > 96 && m[i] < 123 {
			buff = append(buff, m[i])
		}
	}

	for i, _ := range buff {
		if buff[i]+int32(c.I) > 122 {
			buff[i] = buff[i] - int32(26) + int32(c.I)
		} else if buff[i]+int32(c.I) < 97 {
			buff[i] = buff[i] + int32(26) + int32(c.I)
		} else {
			buff[i] = buff[i] + int32(c.I)
		}
	}

	return string(buff)
}

func (c Shift) Decode(s string) string {
	m := []int32(s)
	for i, _ := range s {
		if m[i]-int32(c.I) < 97 {
			m[i] = m[i] + int32(26) - int32(c.I)
		} else if m[i]-int32(c.I) > 122 {
			m[i] = m[i] - int32(26) - int32(c.I)
		} else {
			m[i] = m[i] - int32(c.I)
		}
	}

	return string(m)
}

//-------------------------------------Vingenere-----------------------------------

type Vigenere struct {
	S string
}

func NewVigenere(s string) Cipher {
	m := []int32(s)
	for i, _ := range s {
		if m[i] < 97 || m[i] > 122 {
			return nil
		}
	}

	for i, _ := range s {
		if s[i] != 'a' {
			return Vigenere{s}
		}
	}
	return nil
}

func (v Vigenere) Encode(s string) string {
	m := []int32(s)
	code := []int32(v.S)
	var buff []string
	for i, _ := range s {
		if m[i] > 64 && m[i] < 91 {
			buff = append(buff, string(m[i]+32))
		} else if m[i] > 96 && m[i] < 123 {
			buff = append(buff, string(m[i]))
		}
	}

	for i, _ := range buff {
		var x int32
		if i >= len(code) {
			x = code[i%len(code)] - 97
		} else {
			x = code[i] - 97
		}

		buff[i] = NewShift(int(x)).Encode(buff[i])
	}

	return strings.Join(buff, "")
}

func (v Vigenere) Decode(s string) string {
	m := []int32(s)
	code := []int32(v.S)
	var buff []string
	for i, _ := range s {
		buff = append(buff, string(m[i]))
	}

	for i, _ := range buff {
		var x int32
		if i >= len(code) {
			x = code[i%len(code)] - 97
		} else {
			x = code[i] - 97
		}

		buff[i] = NewShift(int(x)).Decode(buff[i])
	}

	return strings.Join(buff, "")
}
