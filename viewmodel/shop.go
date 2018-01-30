package viewmodel

type Shop struct {
	Title 		string
	Active 		string
	Categories 	[]Category
}

type Category struct {
	URL 		string
	ImageURL 	string
	Title 		string
	Description string
}

func NewShop() Shop {
	result := Shop{
		Title: "Lemonade Stand Supply - Shop",
		Active: "shop",
	}

	juiceCategory := Category {
		URL: "/show_details",
		Title: "Shop",
		ImageURL: "lemon.png",
		Description: `Explore our wide assortment of juices and mixes expected by
			today's lemonade stand clientelle. Now Featuring a full line of organic juices
		`,
	}
	supplyCategory := Category{
		URL: ".",
		ImageURL: "kiwi.png",
		Title: "Cups, Straws, and Other Supplies",
		Description: "Cups!!!!!",
	}

	advertiseCategory := Category{
		URL: ".",
		ImageURL: "pineapple.png",
		Title: "Signs and Advertising",
		Description: "Advertising!!!",
	}

	result.Categories = []Category{juiceCategory, supplyCategory, advertiseCategory}
	return result
}