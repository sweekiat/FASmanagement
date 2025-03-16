package models

type Benefit struct {
        ID       string `json:"id"`
        Name     string `json:"name"`
        Amount   int    `json:"amount"`
}