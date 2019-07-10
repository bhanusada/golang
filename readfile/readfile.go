package main

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	//myslice := []string{}
	file, err := os.Open("interline-HD856.pdf")
	if err != nil {
		log.Fatal(err)
	}
	//fi, _ := file.Stat()
	buf := make([]byte, 10)
	n, _ := file.ReadAt(buf, 0)
	fmt.Println("#######################")
	fmt.Println(string(buf[:n]))
	fmt.Println("***********************")
	/*buf = make([]byte, 100)
	file.ReadAt(buf, fi.Size()-100)
	fmt.Println("#######################")
	//fmt.Printf("%s\n", buf)
	fmt.Println("***********************")
	for len(buf) > 0 && buf[len(buf)-1] == '\n' || buf[len(buf)-1] == '\r' {
		buf = buf[:len(buf)-1]
	}
	buf = bytes.TrimRight(buf, "\r\n\t ")
	if !bytes.HasSuffix(buf, []byte("%%EOF")) {
		fmt.Println("not a PDF file: missing %%%%EOF")
	}

	i := findLastLine(buf, "startxref")

	fmt.Printf("number of startxref : %d\n ", i)*/

	buf1 := make([]byte, 406)
	file.ReadAt(buf1, 220)
	fmt.Printf("%s\n", buf1)
	defer file.Close()

	b := bytes.NewReader(buf1)

	r, err := zlib.NewReader(b)
	io.Copy(os.Stdout, r)
	r.Close()
	//buf = bytes.TrimRight(buf, "\r\n\t ")
	//fmt.Printf("%v", buf)
	//for _, cont := range buf {
	//		fmt.Println(cont)
	//	}
	/*scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// fmt.Println("len", len(myslice))
		myslice = append(myslice, scanner.Text())
	}
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}*/
	//fmt.Println(myslice)

	/*file, _ := ioutil.ReadFile("interline-HD856.pdf")
	fmt.Printf("%s", file) */
}

func findLastLine(buf []byte, s string) int {
	bs := []byte(s)
	max := len(buf)
	for {
		i := bytes.LastIndex(buf[:max], bs)
		if i <= 0 || i+len(bs) >= len(buf) {
			return -1
		}
		if (buf[i-1] == '\n' || buf[i-1] == '\r') && (buf[i+len(bs)] == '\n' || buf[i+len(bs)] == '\r') {
			return i
		}
		max = i
	}
}
