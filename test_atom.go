package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"code.google.com/p/go-charset/charset"
	"github.com/mjibson/goread/atom"
)

func main() {
	//resp, _ := http.Get("http://feeds.feedburner.com/newsyc100")
	resp, _ := http.Get("http://feeds.feedburner.com/Savoir-inutilecom-ToutCeQuiNeSertRienPourBrillerEnSocit")

	reader := io.LimitReader(resp.Body, 1<<21)
	b, _ := ioutil.ReadAll(reader)

	fmt.Println(string(b))
	r := atom.Feed{}
	d := xml.NewDecoder(bytes.NewReader(b))
	d.CharsetReader = charset.NewReader
	if err := d.Decode(&r); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", r)
}
