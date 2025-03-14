package models

type HouseholdMember struct {
        ID              string `json:"id"`
        ApplicantID     string `json:"applicant_id"`
        Name            string `json:"name"`
        EmploymentStatus string `json:"employment_status"`
        Sex             string `json:"sex"`
        DateOfBirth     string `json:"date_of_birth"` 
        Relation        string `json:"relation"`
}