package models

type Application struct {
        ID          string `json:"id"`
        Status      string `json:"status"`
        ApplicantID string `json:"applicant_id"`
        SchemeID    string `json:"scheme_id"`
}