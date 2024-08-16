package menus

import "cypher/menu"

var MainMenu *menu.Menu

func init() {
	MainMenu = menu.NewMenu("main")
	MainMenu.AddOption("t", "transposition sub menu", func() {
		TranspositionMenu.Start()
	})
	MainMenu.AddOption("m", "monosubstitution sub menu", func() {
		MonoSubstitutionMenu.Start()
	})
}
