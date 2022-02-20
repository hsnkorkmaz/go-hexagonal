package domain

type Author struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func NewAuthor(id int64, name string) Author {
	return Author{
		Id:   id,
		Name: name,
	}
}