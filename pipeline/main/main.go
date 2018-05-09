package main

import(
	"fmt"

	"github.com/cloudnativego/go-primer/npcs"
)

func main()  {
	fmt.Println("Hello cloud")

	mob := npcs.NonPlayerCharacter{Name: "Alfred"}
	fmt.Println(mob)

	hortense := npcs.NonPlayerCharacter{Name: "Hortense",
		Loc: npcs.Location{X: 10.0, Y: 10.0, Z: 10.0}}

	fmt.Println(hortense)


	fmt.Printf("Alfred is %f units from Hortense.\n",
		mob.DistinceTo(hortense))
}

