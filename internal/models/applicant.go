package models

type Applicant struct {
        ID              string    `json:"id"`
        Name            string    `json:"name"`
        EmploymentStatus string    `json:"employment_status"`
        MaritalStatus   string    `json:"marital_status"`
        Sex             string    `json:"sex"`
        DateOfBirth     string    `json:"date_of_birth"` 
		HouseholdMembers []HouseholdMember `json:"household"`
}