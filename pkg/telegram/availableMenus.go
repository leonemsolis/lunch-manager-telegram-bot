package telegram

type CafeMenus struct {
	Cafe string `json:"cafe"`
	Menus []AvailableMenu `json:"menus"`
}

type AvailableMenu struct {
	Name string `json:"name"`
	Items []string `json:"items"`
}
