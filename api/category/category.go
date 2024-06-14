package category

type Category struct {
	Id   int
	Name string
	Goal float64
}

func NewCategory() Category {
	return Category{}
}
