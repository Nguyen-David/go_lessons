package entity

import (
	"encoding/json"
	"strconv"
)

type Book struct {
	Name   string `csv:"name"`
	Author string `csv:"author"`
	Year   int    `csv:"year"`
}

type BookRequestAndResponse struct {
	Name   string `json:"name"`
	Author string `json:"author"`
	Year   string `json:"year"`
}

type BookList struct {
	Books []Book `json:"books"`
	Date BookTime `json:"date"`
}

type SortedBooks []Book

func (a SortedBooks) Len() int           { return len(a) }
func (a SortedBooks) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortedBooks) Less(i, j int) bool { return a[i].Year < a[j].Year }

func (b *Book) MarshalJSON() ([]byte, error) {
	return json.Marshal(&BookRequestAndResponse{
		Name:   b.Name,
		Author: b.Author,
		Year:   strconv.Itoa(b.Year),
	})
}

func (b *Book) UnmarshalJSON(data []byte) error {
	book := &BookRequestAndResponse{}

	if err := json.Unmarshal(data, book); err != nil {
		return err
	}

	var err error
	b.Name = book.Name
	b.Author = book.Author
	b.Year, err = strconv.Atoi(book.Year)


	return err
}
