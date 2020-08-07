package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"unsafe"
)

// Struct
type mbr2 struct { //3
	N1 uint8
	N2 uint8
	N3 uint8
}

type mbr struct { //22
	Numero   uint8
	Caracter byte
	Cadena   [20]byte
}

func main() {

	writeFile()
	fmt.Println("Reading File: ")
	readFile()
}

func writeFile() {
	file, err := os.Create("test.bin")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	disco := mbr{Numero: 25}
	disco.Caracter = 'D'

	// Igualar cadenas a array de bytes (array de chars)
	cad := "Hola Amigos"
	copy(disco.Cadena[:], cad)

	s := &disco

	var binario bytes.Buffer
	binary.Write(&binario, binary.BigEndian, s)
	writeNextBytes(file, binario.Bytes())

	//primer structsegundostruct

	disco2 := mbr2{}
	disco2.N1 = 14
	disco2.N2 = 222
	disco2.N3 = 30

	s1 := &disco2

	var binario2 bytes.Buffer
	binary.Write(&binario2, binary.BigEndian, s1)
	writeNextBytes(file, binario2.Bytes())

}

func writeNextBytes(file *os.File, bytes []byte) {

	_, err := file.Write(bytes)

	if err != nil {
		log.Fatal(err)
	}

}

func readFile() {

	file, err := os.Open("test.bin")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	m := mbr{}
	var size int = int(unsafe.Sizeof(m))

	data := readNextBytes(file, size)
	buffer := bytes.NewBuffer(data)

	fmt.Println(data)

	err = binary.Read(buffer, binary.BigEndian, &m)
	if err != nil {
		log.Fatal("binary.Read failed", err)
	}

	fmt.Println(m)

	fmt.Printf("Caracter: %c\nCadena: %s\n", m.Caracter, m.Cadena)

	file.Seek(0, 0) // segundo parametro: 0, 1, 2.     0 -> Inicio, 1-> desde donde esta el puntero, 2 -> Del fin para atras
	file.Seek(int64(unsafe.Sizeof(m)), 0)
	//Struct 2
	fmt.Println("Struct 2: ")
	m2 := mbr2{}
	size = int(unsafe.Sizeof(m2))

	data = readNextBytes(file, size)
	buffer = bytes.NewBuffer(data)

	fmt.Println(data)

	err = binary.Read(buffer, binary.BigEndian, &m2)
	if err != nil {
		log.Fatal("binary.Read failed", err)
	}

	fmt.Println(m2)

}

func readNextBytes(file *os.File, number int) []byte {
	bytes := make([]byte, number)

	_, err := file.Read(bytes)
	if err != nil {
		log.Fatal(err)
	}

	return bytes
}
