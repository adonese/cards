package main

import (
	"bufio"
	"flag"
	"io/ioutil"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/skip2/go-qrcode"
)

var fname = flag.String("f", "data.txt", "file to write data to")
var dir = flag.String("dir", "data", "Working directory")
var count = flag.Int("c", 5000, "number of outputted uuids")

func generate(count int) []byte {
	var res []byte

	for i := 1; i <= count; i++ {
		uid := uuid.New().String()
		data := []byte(uid + "\n")
		res = append(res, data...)
	}
	return res
}

func reader(fname string) []string {
	var res []string
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return res
}

func qr(ids []string) error {
	for _, v := range ids {
		var png []byte
		png, err := qrcode.Encode(v, qrcode.Medium, 256)
		if err != nil {
			log.Printf("Error in generating qr")
			return err
		}
		// path := path.Join(*dir, v[1:10]+".txt")
		err = ioutil.WriteFile(v[1:10]+".png", png, 0664)
		if err != nil {
			log.Print(err)
			return err
		}
	}
	return nil
}

func main() {
	flag.Parse()
	// res := generate(*count)
	files := reader(*fname)
	log.Fatal(qr(files))
}
