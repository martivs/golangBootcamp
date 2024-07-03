package myLib

import (
	"fmt"
)

type cakes struct {
	name string
	oi   int
	ni   int
	ings []ingredients
}

type ingredients struct {
	name string
	oi   int
	ni   int
}

func CompareJson(old, new *JsonReader) {
	// cakes
	checkAddCake(old, new)
	checkRemovedCake(old, new)
	// tyime, ingridients etc
	cakes := getEqCake(old, new)
	checkTime(cakes, old, new)
	checkAddIng(cakes, old, new)
	checkRemovedIng(cakes, old, new)
	// ingridirnts
	for i := range cakes {
		cakes[i].ings = getEqIngs(&cakes[i], old, new)
	}
	checkIngs(cakes, old, new)

}

func checkAddCake(old, new *JsonReader) {
	for _, newCake := range new.Cake {
		for j, oldCake := range old.Cake {
			if newCake.Name == oldCake.Name {
				break
			} else {
				if j == len(old.Cake)-1 {
					fmt.Printf("ADDED cake \"%s\"\n", newCake.Name)
				}
			}
		}
	}
}

func checkRemovedCake(old, new *JsonReader) {
	for _, oldCake := range old.Cake {
		for j, newCake := range new.Cake {
			if oldCake.Name == newCake.Name {
				break
			} else {
				if j == len(new.Cake)-1 {
					fmt.Printf("REMOVED cake \"%s\"\n", oldCake.Name)
				}
			}
		}
	}
}

func getEqCake(old, new *JsonReader) []cakes {
	var res []cakes
	for i, oldCake := range old.Cake {
		for j, newCake := range new.Cake {
			if oldCake.Name == newCake.Name {
				res = append(res, cakes{name: oldCake.Name, oi: i, ni: j})
				break
			}
		}
	}
	return res
}

func checkTime(cakes []cakes, old, new *JsonReader) {
	for _, cake := range cakes {
		if new.Cake[cake.ni].Time != old.Cake[cake.oi].Time {
			fmt.Printf("CHANGED cooking time for cake \"%s\" - \"%s\" instead \"%s\"\n", cake.name, new.Cake[cake.ni].Time, old.Cake[cake.oi].Time)
		}
	}
}

func checkAddIng(cakes []cakes, old, new *JsonReader) {
	for _, value := range cakes {
		for _, newIng := range new.Cake[value.ni].Ingredients {
			for j, oldIng := range old.Cake[value.oi].Ingredients {
				if newIng.IngredientName == oldIng.IngredientName {
					break
				} else {
					if j == len(old.Cake[value.oi].Ingredients)-1 {
						fmt.Printf("ADDED ingredient \"%s\" for cake \"%s\"\n", newIng.IngredientName, value.name)
					}
				}
			}
		}
	}
}

func checkRemovedIng(cakes []cakes, old, new *JsonReader) {
	for _, value := range cakes {
		for _, oldIng := range old.Cake[value.oi].Ingredients {
			for j, newIng := range new.Cake[value.ni].Ingredients {
				if newIng.IngredientName == oldIng.IngredientName {
					break
				} else {
					if j == len(new.Cake[value.oi].Ingredients)-1 {
						fmt.Printf("REMOVED ingredient \"%s\" for cake \"%s\"\n", oldIng.IngredientName, value.name)
					}
				}
			}
		}
	}
}

func getEqIngs(cake *cakes, old, new *JsonReader) []ingredients {
	var res []ingredients
	for i, oldIng := range old.Cake[cake.oi].Ingredients {
		for j, newIng := range new.Cake[cake.ni].Ingredients {
			if oldIng.IngredientName == newIng.IngredientName {
				res = append(res, ingredients{name: oldIng.IngredientName, oi: i, ni: j})
				break
			}
		}
	}
	return res
}

func checkIngs(cakes []cakes, old, new *JsonReader) {
	for _, cake := range cakes {
		for _, ings := range cake.ings {
			ingNew := new.Cake[cake.ni].Ingredients[ings.ni].IngredientUnit
			ingOld := old.Cake[cake.oi].Ingredients[ings.oi].IngredientUnit
			if ingNew != ingOld {
				if ingNew == "" {
					fmt.Printf("REMOVED unit \"%s\" for ingredient \"%s\" for cake  \"%s\"\n", old.Cake[cake.oi].Ingredients[ings.oi].IngredientUnit, ings.name, cake.name)
				} else {
					if ingOld == "" {
						fmt.Printf("ADDED unit \"%s\" for ingredient \"%s\" for cake  \"%s\"\n", new.Cake[cake.oi].Ingredients[ings.oi].IngredientUnit, ings.name, cake.name)
					} else {
						fmt.Printf("CHANGED unit for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n", ings.name, cake.name, new.Cake[cake.ni].Ingredients[ings.ni].IngredientUnit, old.Cake[cake.oi].Ingredients[ings.oi].IngredientUnit)
					}
				}
			}
			ingNew = new.Cake[cake.ni].Ingredients[ings.ni].IngredientCount
			ingOld = old.Cake[cake.oi].Ingredients[ings.oi].IngredientCount
			if ingNew != ingOld {
				if ingNew == "" {
					fmt.Printf("REMOVED unit count \"%s\" for ingredient \"%s\" for cake  \"%s\"\n", old.Cake[cake.oi].Ingredients[ings.oi].IngredientUnit, ings.name, cake.name)
				} else {
					if ingOld == "" {
						fmt.Printf("ADDED unit count \"%s\" for ingredient \"%s\" for cake  \"%s\"\n", new.Cake[cake.oi].Ingredients[ings.oi].IngredientUnit, ings.name, cake.name)
					} else {
						fmt.Printf("CHANGED unit count for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n", ings.name, cake.name, new.Cake[cake.ni].Ingredients[ings.ni].IngredientCount, old.Cake[cake.oi].Ingredients[ings.oi].IngredientCount)
					}
				}
			}
		}
	}
}
