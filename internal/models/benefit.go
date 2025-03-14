package models

type Benefit struct {
        SchemeID string `json:"scheme_id"`
        ID       string `json:"id"`
        Name     string `json:"name"`
        Amount   int    `json:"amount"`
}