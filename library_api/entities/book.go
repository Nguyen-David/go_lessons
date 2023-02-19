package entities

type Book struct {
	Name   string `json:"name"`
	Author string `json:"author"`
	Date   uint `json:"date"`
}

type ByDate []Book

func (a ByDate) Len() int           { return len(a) }
func (a ByDate) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByDate) Less(i, j int) bool { return a[i].Date < a[j].Date }
