package main

import "log"

func PreparationTime(layers []string, time int) int {
    if time == 0 {
        time = 2
    }

    return len(layers)*time
}

func Quantities(layers []string) (int, float64) {

    noddles, sauce := 0, 0

    for _, v := range layers {
        switch v {
            case "noodles":
                noddles++
            case "sauce":
                sauce++
        }
    }

    return noddles*50, float64(sauce)*0.2
}

func AddSecretIngredient(friendsList []string, myList []string) {
    myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

func ScaleRecipe(quantities []float64, portions int) []float64 {

    out := make([]float64, len(quantities))

    for i, v := range quantities {
        out[i] = v*(float64(portions)/2)
    }

    return out
}

func main() {
    log.Println(PreparationTime([]string{"sauce", "noodles", "sauce", "meat", "mozzarella", "noodles"}, 3))
    log.Println(PreparationTime([]string{"sauce", "noodles", "sauce", "meat", "mozzarella", "noodles"}, 0))
    log.Println(Quantities([]string{"sauce", "noodles", "sauce", "meat", "mozzarella", "noodles"}))
    myList := []string{"noodles", "meat", "sauce", "mozzarella","?"}
    AddSecretIngredient([]string{"noodles", "sauce", "mozzarella", "kampot pepper"}, myList)
    log.Println(myList)
    log.Println(ScaleRecipe([]float64{ 1.2, 3.6, 10.5 }, 4))
}