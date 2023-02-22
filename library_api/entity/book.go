package entity

type Book struct {
	Name   string `json:"name"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

type ByYear []Book

func (a ByYear) Len() int           { return len(a) }
func (a ByYear) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByYear) Less(i, j int) bool { return a[i].Year < a[j].Year }
