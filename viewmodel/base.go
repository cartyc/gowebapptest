package viewmodel

type Base struct {
	Title  string
	Active string
}

func NewBase() Base {
	return Base{
		Active: "home",
		Title:  "Lemonade Stand Supply",
	}
}
