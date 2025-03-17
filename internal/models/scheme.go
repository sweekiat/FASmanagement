package models

import (
	"time"
)

type Scheme struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Criterias   Criteria  `json:"criteria"`
	Benefits    []Benefit `json:"benefits"`
}

func (s *Scheme) IsElligible(applicant Applicant) bool {
	if s.Criterias["employment_status"] != nil {
		if s.Criterias["employment_status"] != applicant.EmploymentStatus {
			return false
		}
	}
	if s.Criterias["marital_status"] != nil {
		if s.Criterias["marital_status"] != applicant.MaritalStatus {
			return false
		}
	}
	hhMembers := applicant.HouseholdMembers
	if s.Criterias["has_children"] != nil {
		if len(hhMembers) == 0 {
			return false
		}
		var hasChildren bool
		childrenJson := s.Criterias["has_children"].(map[string]interface{})
		if childrenJson["school_level"] != "" {
			for _, member := range hhMembers {
				if member.Relation == "daughter" || member.Relation == "son" {
					hasChildren = true
					if checkAges(member, childrenJson["school_level"].(string)) {
						return true
					}
				}
			}
			return false
		}
		return hasChildren

	}
	return true

}

func checkAges(hhMember HouseholdMember, level string) bool {
	dob := hhMember.DateOfBirth
	dobTime, _ := time.Parse(time.DateOnly, dob)
	age := time.Now().Year() - dobTime.Year()
	if level == "primary" {
		if age < 6 || age > 12 {
			return false
		}
	}
	if level == "secondary" {
		if age < 13 || age > 16 {
			return false
		}
		if level == "tertiary" {
			if age < 17 || age > 18 {
				return false
			}
		}

	}
	return true
}
