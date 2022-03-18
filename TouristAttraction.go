package main

import (
	"fmt"
	"github.com/elliotchance/orderedmap"
)

func main() {
	ListOfTouristAttractionsInNigeria := orderedmap.NewOrderedMap()
	ListOfTouristAttractionsInNigeria.Set("ObuduMountainResort", "CrossRiver")
	ListOfTouristAttractionsInNigeria.Set("ChappalWadi mountains", "Taraba")
	ListOfTouristAttractionsInNigeria.Set("IdanreHill", "Ondo")
	ListOfTouristAttractionsInNigeria.Set("ErinIjeshaWaterfalls", "Osun")
	ListOfTouristAttractionsInNigeria.Set("YankariNationalPark", "Bauchi")
	ListOfTouristAttractionsInNigeria.Set("OgbunikeCaves", "Anambra")
	ListOfTouristAttractionsInNigeria.Set("SurameCulturalLandscape", "Sokoto")

	for el := ListOfTouristAttractionsInNigeria.Front(); el != nil; el = el.Next() {
		fmt.Println(el.Key, el.Value)
	}

}
