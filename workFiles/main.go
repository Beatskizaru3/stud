package main

import (
	"fmt"
	"io"
	"log"
)

func SimpeReader() {
	buf := make([]byte, 10)

	reader := NumsReader{nums: "1,23,4.2.325,4;sdk!"}

	count, err := reader.Read(buf)
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
	fmt.Println(string(buf), count) // выводим фильтрованую строку приводя слайс байтов и количество элементов
}

type NumsReader struct {
	nums string
}

func (r NumsReader) Read(p []byte) (n int, err error) {
	var count int
	for i := 0; i < len(r.nums); i++ {
		if r.nums[i] >= '0' && r.nums[i] <= '9' {
			p[count] = r.nums[i] // записуем байт, если он нам подоходит в слайсБайтов(буфер)
			count++
		}
	}
	return count, io.EOF //EOF - end of file
}

func SimpleWriter() {
	nums := []byte{'1', 's', '5', 'x', '2', '3', '!'}

	writer := NumsWriter{storedNums: make([]byte, 10)}
	count, err := writer.Write(nums)
	if err != nil && err != io.EOF { // <== исправлено условие обработки ошибки
		log.Fatal(err)
	}
	fmt.Println(string(writer.storedNums), count)
}

type NumsWriter struct {
	storedNums []byte
}

func (w *NumsWriter) Write(p []byte) (n int, err error) {
	var count int

	for i := 0; i < len(p); i++ {
		if p[i] >= '0' && p[i] <= '9' {
			w.storedNums[count] = p[i]
			count++
		}
	}
	return count, io.EOF
}

func main() {
	SimpleWriter()
}
