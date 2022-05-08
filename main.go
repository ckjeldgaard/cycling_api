package main

type rider struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Nationality string `json:"nationaliy"`
	Age         int    `json:"age"`
	Team        string `json:"team"`
}

// riders slice to seed rider data.
var riders = []rider{
	{ID: "1", Name: "Wout van Aert", Nationality: "Belgium", Age: 27, Team: "Jumbo-Visma"},
	{ID: "2", Name: "Mathieu van der Poel", Nationality: "Netherlands", Age: 27, Team: "Alpecin-Fenix"},
	{ID: "3", Name: "Julian Alaphilippe", Nationality: "France", Age: 29, Team: "Quick-Step Alpha Vinyl Team"},
	{ID: "4", Name: "Kasper Asgreen", Nationality: "Denmark", Age: 27, Team: "Quick-Step Alpha Vinyl Team"},
	{ID: "5", Name: "Tadej Pogačar", Nationality: "Slovenia", Age: 23, Team: "UAE Team Emirates"},
}
