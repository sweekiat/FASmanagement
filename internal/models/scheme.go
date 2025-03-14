package models

type Scheme struct {
        ID          string `json:"id"`
        Name        string `json:"name"`
        Description string `json:"description"`
		Criterias   Criteria `json:"criteria"`
		Benefits  []Benefit `json:"benefits"`

}