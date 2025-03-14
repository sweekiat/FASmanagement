package models

type Criteria struct {
        SchemeID string `json:"scheme_id"`
        Criteria string `json:"criteria"` // find a way to parse from json to a map, inside db is a json
}