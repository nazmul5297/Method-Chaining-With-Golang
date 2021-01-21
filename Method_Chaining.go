package main

import (
    "fmt"
)

type Cricketer struct {
    Name string
    Jursey_Number  int
    Run int
}

func (crictr *Cricketer) fair() *Cricketer {
    fmt.Printf("fair_score %s\n", crictr.Name)
    if crictr.run>50 {
		crictr.fair_score=crictr.fair_score+10
	}
	else {
		crictr.fair_score=crictr.fair_score-10
	}
    return crictr
}

func (crictr *Cricketer) SetName(name string) *Cricketer {
    fmt.Printf("Set name of new Cricketer to %s\n", name)
    crictr.Name = name
    return crictr
}

func main() {

    Player := new(Cricketer)

	Player.fair_score = 10
	Player.run=60

    Player.SetName("Musfikur").fair()

    fmt.Printf("Here we have %s with rank %d\n",  Player.Name,  Player.fair_score)

}