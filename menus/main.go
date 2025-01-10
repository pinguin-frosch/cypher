package menus

import (
	"cypher/menus/monosubstitution"
	"cypher/menus/polisubstitution"
	"cypher/menus/transposition"

	"github.com/pinguin-frosch/menu/pkg/menu"
)

var Main *menu.Menu

func init() {
	Main = menu.NewMenu("main")
	Main.AddOption("t", "transposition sub menu", func() {
		transposition.Main.Start()
	})
	Main.AddOption("m", "monosubstitution sub menu", func() {
		monosubstitution.Main.Start()
	})
	Main.AddOption("p", "polisubstitution sub menu", func() {
		polisubstitution.Main.Start()
	})
}
