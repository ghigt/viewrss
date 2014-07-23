package main

type Feed struct {
	ID          in64     `json:"id"`
	Title       string   `json:"title"`
	Link        string   `json:"link"`
	Description string   `json:"description"`
	Date        string   `json:"date"`
	Stories     []*Story `json:"stories"`
}

type Story struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Link        string `json:"link"`
	Description string `json:"description"`
	Date        string `json:"date"`
	Content     string `json:"content"`
}

func main() {
}
