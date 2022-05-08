package main

type rider struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Nationality string `json:"nationaliy"`
	Age         int    `json:"age"`
	Team        string `json:"team"`
}
