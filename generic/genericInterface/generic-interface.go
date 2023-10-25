package genericInterface

import "fmt"

func Run() {
	// mock movie
	movie := []Movie{
		{
			Name:  "movie1",
			Price: 30,
		},
		{
			Name:  "movie2",
			Price: 32,
		},
		{
			Name:  "movie3",
			Price: 30,
		},
	}

	// movie2 := &[]Movie{
	// 	{
	// 		Name:  "movie2_1",
	// 		Price: 30,
	// 	},
	// 	{
	// 		Name:  "movie2_2",
	// 		Price: 233,
	// 	},
	// }
	// mock game
	game := []Game{
		{
			Name:  "game1",
			Price: 12,
		},
		{
			Name:  "game2",
			Price: 38,
		},
	}

	fmt.Println("sum price movie: ", sum(movie))
	fmt.Println("sum price game:", sum(game))

	// fmt.Println("rest name movie: ", resetName(movie))
	// fmt.Println("rest name movie: ", resetName2(movie2[0]))

}

type Movie struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func (m Movie) GetPrice() int {
	return m.Price
}

func (m *Movie) SetName(name string) {
	m.Name = name
}

type Game struct {
	Name     string `json:"name"`
	Platform string `json:"platform"`
	Price    int    `json:"price"`
}

func (g Game) GetPrice() int {
	return g.Price
}

func (g *Game) SetName(name string) {
	g.Name = name
}

type IPrice interface {
	GetPrice() int
}

func sum[T IPrice](objs []T) int {
	var result int
	for _, obj := range objs {
		result += obj.GetPrice()
	}
	return result
}

type IName interface {
	*Movie | *Game
	SetName(name string)
}

func resetName[T IName](objs *[]T) {
	for _, obj := range *objs {
		obj.SetName("")
	}
}

// func resetName2[T IName](objs *T) {
// 	obj.SetName("")
// }
