package maps

import "fmt"

func ShowsMaps() {
	countries := make(map[string]string)
	fmt.Println()
	countries["Venezuela"] = "Caracas"
	countries["Japan"] = "Tokyo"
	countries["Argentina"] = "Buenos Aires"
	fmt.Println(countries)
	fmt.Println(countries["Venezuela"])

	imdb := map[string]float32{
		"Butterfly Effect":             6.8,
		"Donnie Darko":                 7.2,
		"The killing of a sacred Deer": 7.8,
		"Shanshawk Redemption":         9.8,
	}

	fmt.Println(imdb)
	for movie, pts := range imdb {
		fmt.Printf("The movie %s has an evaluation of %f in imdb\n", movie, pts)
	}

	delete(imdb, "Butterfly Effect")
	fmt.Println(imdb)

	pts, exist := imdb["Butterfly Effect"]
	if exist {
		fmt.Printf("The movie exist and has a evaluation of: %f", pts)
	} else {
		fmt.Println("The movie don't exist.")
	}

}
