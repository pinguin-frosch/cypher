package menus

import "cypher/menu"

var MainMenu *menu.Menu

func init() {
	MainMenu = menu.NewMenu("main")
	MainMenu.AddOption("t", "transposition sub menu", func() {
		TranspositionMenu.Start()
	})
}
