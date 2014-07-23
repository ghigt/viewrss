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
	_ "code.google.com/p/go-charset/data"
	"github.com/mjibson/goread/atom"
	"github.com/mjibson/goread/rss"
)

func get(s string) []byte {
	resp, err := http.Get(s)
	if err != nil {
		log.Fatal(err)
	} else if resp.StatusCode != http.StatusOK {
		log.Fatal("resp != StatusOK")
	}

	reader := io.LimitReader(resp.Body, 1<<21)
	b, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}

	return b
}

func parseAtom(s string) {

	b := get(s)
	a := atom.Feed{}
	d := xml.NewDecoder(bytes.NewReader(b))
	d.CharsetReader = charset.NewReader
	if err := d.Decode(&a); err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%#v\n", a)
	//fmt.Println("XMLName:", a.XMLName)
	//fmt.Println("title:", a.Title)
	//fmt.Println("ID:", a.ID)
	//fmt.Println("Links:", a.Link)
	//fmt.Println("Updated:", a.Updated)
	//fmt.Println("XMLBase:", a.XMLBase)

	//fmt.Println("Entry:", a.Entry)
	//for _, e := range a.Entry {
	//	fmt.Println("title:", e.Title)
	//	fmt.Println("ID:", e.ID)
	//	fmt.Println("Link:", e.Link)
	//	fmt.Println("Published:", e.Published)
	//	fmt.Println("Updated:", e.Updated)
	//	//fmt.Println("Summary:", e.Summary)
	//	fmt.Println()
	//}
	fmt.Println("Parsing Atom OK ->", s)
}

func parseRss(s string) {
	b := get(s)

	r := rss.Rss{}
	d := xml.NewDecoder(bytes.NewReader(b))
	d.CharsetReader = charset.NewReader
	d.DefaultSpace = "DefaultSpace"
	if err := d.Decode(&r); err != nil {
		log.Fatal(err)
	}
	//fmt.Println("")

	//fmt.Println("r.XMLName:", r.XMLName)
	//fmt.Println("r.Title:", r.Title)
	//fmt.Println("r.Link:", r.Link)
	//fmt.Println("r.Description:", r.Description)
	//fmt.Println("r.PubDate:", r.PubDate)
	//fmt.Println("r.LastBuildDate:", r.LastBuildDate)
	//fmt.Println("r.Items:", r.Items)
	//for _, i := range r.Items {
	//	fmt.Println("title:", i.Title)
	//	fmt.Println("link:", i.Link)
	//	fmt.Println("description:", i.Description)
	//	fmt.Println("author:", i.Author)
	//	fmt.Println("enclosure:", i.Enclosure)
	//	fmt.Println("guid:", i.Guid)
	//	fmt.Println("pubdate:", i.PubDate)
	//	fmt.Println("source:", i.Source)
	//	//fmt.Println(i.Content    )
	//	fmt.Println("date:", i.Date)
	//	fmt.Println("published:", i.Published)
	//	fmt.Println("media:", i.Media)
	//	fmt.Println()
	//}
	fmt.Println("Parsing RSS OK ->", s)
}

func main() {
	//parseRss("http://feeds.feedburner.com/newsyc100")
	//parseRss("http://feeds.feedburner.com/Savoir-inutilecom-ToutCeQuiNeSertRienPourBrillerEnSocit")
	//parseRss("http://feeds.feedburner.com/Grafikart")
	//parseRss("http://feeds2.feedburner.com/KorbensBlog-UpgradeYourMind")
	//parseRss("http://www.alexedwards.net/static/feed.rss")
	//parseRss("http://areyoufuckingcoding.me/feed/")
	//parseRss("http://blog.atom.io/feed.xml")
	//parseAtom("http://codegangsta.io/atom.xml")
	//parseRss("http://feeds.feedburner.com/codinghorror/")
	//parseRss("http://www.commitstrip.com/en/feed/")
	//parseAtom("http://cyrilmottier.com/atom.xml")
	//parseRss("http://dave.cheney.net/feed")
	//parseAtom("http://dominik.honnef.co/atom.xml")
	//parseRss("https://egghead.io/feed")
	//parseRss("http://fabiensanglard.net/rss.xml")
	//parseAtom("http://featherweightmusings.blogspot.com/feeds/posts/default")
	//parseRss("http://feeds.feedburner.com/Frandroid")
	//parseRss("http://rss.futura-sciences.com/packfs")
	//parseRss("http://www.geek.com/feed/")
	//parseRss("http://gizmodo.com/rss")
	//parseRss("http://golang-examples.tumblr.com/rss")
	//parseAtom("http://nathanleclaire.com/atom.xml")
	//parseRss("http://www.pheelicks.com/feed/")
	//parseRss("http://english.blogoverflow.com/feed/")
	//parseRss("http://thechangelog.com/feed/")
	//parseAtom("http://blog.golang.org/feeds/posts/default?alt=rss")
	//parseAtom("http://blog.gopheracademy.com/feed.atom")
	parseRss("http://blog.newsblur.com/rss")
	parseAtom("http://www.theverge.com/rss/index.xml")
	parseRss("http://www.tutos-android.com/feed")
	parseRss("http://blog.unrolled.ca/rss/")
	parseAtom("https://blog.uraniborg.net/atom.xml")
	parseRss("http://whizdumb.me/feed/")
	parseAtom("http://xkcd.com/atom.xml")
	parseAtom("http://feeds.feedburner.com/yannespositocomfr")
}
