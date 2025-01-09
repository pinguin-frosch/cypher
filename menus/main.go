package menus

import "github.com/pinguin-frosch/menu/pkg/menu"

var MainMenu *menu.Menu

func init() {
	MainMenu = menu.NewMenu("main")
	MainMenu.AddOption("t", "transposition sub menu", func() {
		TranspositionMenu.Start()
	})
	MainMenu.AddOption("m", "monosubstitution sub menu", func() {
		MonoSubstitutionMenu.Start()
	})
	MainMenu.AddOption("p", "polisubtitution sub menu", func() {
		PolisubstitutionMenu.Start()
	})
}
