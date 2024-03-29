package models

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Element struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

var Elements = []Element{
	{Name: "Airbending", Description: "Manipulation of air currents", Image: "https://example.com/airbending.jpg"},
	{Name: "Waterbending", Description: "Manipulation of water in various forms", Image: "https://example.com/waterbending.jpg"},
	{Name: "Earthbending", Description: "Manipulation earth and rock", Image: "https://example.com/earthbending.jpg"},
	{Name: "Firebending", Description: "Manipulation of fire and lightning", Image: "https://example.com/firebending.jpg"},
}
