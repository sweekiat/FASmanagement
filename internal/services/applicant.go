package services

import (
	"FASManagementSystem/internal/database"
	"FASManagementSystem/internal/models"
	"database/sql"
)

func GetAllApplicantsWithHousehold() ([]models.Applicant, error) {
	db := database.GetDB()
	rows, err := db.Query(`
        SELECT
            a.id, a.name, a.employment_status, a.marital_status, a.sex, a.date_of_birth,
            h.id, h.applicant_id, h.name, h.employment_status, h.sex, h.date_of_birth, h.relation
        FROM
            applicants a
        LEFT JOIN
            household_members h ON a.id = h.applicant_id;
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	applicantMap := make(map[string]*models.Applicant)

	for rows.Next() {
		var (
			applicantID, applicantName, applicantEmploymentStatus, applicantMaritalStatus, applicantSex, applicantDateOfBirth                  string
			householdID, householdApplicantID, householdName, householdEmploymentStatus, householdSex, householdDateOfBirth, householdRelation sql.NullString
		)

		err := rows.Scan(
			&applicantID, &applicantName, &applicantEmploymentStatus, &applicantMaritalStatus, &applicantSex, &applicantDateOfBirth,
			&householdID, &householdApplicantID, &householdName, &householdEmploymentStatus, &householdSex, &householdDateOfBirth, &householdRelation,
		)
		if err != nil {
			return nil, err
		}

		applicant, exists := applicantMap[applicantID]
		if !exists {
			applicant = &models.Applicant{
				ID:               applicantID,
				Name:             applicantName,
				EmploymentStatus: applicantEmploymentStatus,
				MaritalStatus:    applicantMaritalStatus,
				Sex:              applicantSex,
				DateOfBirth:      applicantDateOfBirth,
			}
			applicantMap[applicantID] = applicant
		}

		if householdID.Valid {
			applicant.HouseholdMembers = append(applicant.HouseholdMembers, models.HouseholdMember{
				ID:               householdID.String,
				ApplicantID:      householdApplicantID.String,
				Name:             householdName.String,
				EmploymentStatus: householdEmploymentStatus.String,
				Sex:              householdSex.String,
				DateOfBirth:      householdDateOfBirth.String,
				Relation:         householdRelation.String,
			})
		} else {
			applicant.HouseholdMembers = []models.HouseholdMember{}
		}
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	applicants := make([]models.Applicant, 0, len(applicantMap))
	for _, applicant := range applicantMap {
		applicants = append(applicants, *applicant)
	}
	return applicants, nil
}


func CreateNewApplicant(applicant models.Applicant) error {
	db := database.GetDB()
	tx, dbErr := db.Begin()
	if dbErr != nil {
		return dbErr
	}

	householdMembers := applicant.HouseholdMembers
	if applicant.MaritalStatus == "" {
		applicant.MaritalStatus = "Single"
	}
	_, err := tx.Exec(`
                INSERT INTO applicants (
						id,
                        name,
                        employment_status,
                        marital_status,
                        sex,
                        date_of_birth
                ) VALUES (?, ?, ?, ?, ?, ?)
        `, applicant.ID, applicant.Name, applicant.EmploymentStatus, applicant.MaritalStatus, applicant.Sex, applicant.DateOfBirth)
	if err != nil {
		tx.Rollback()
		return err
	}
	for _, householdMember := range householdMembers {
		_, insertErr := tx.Exec(`
                INSERT INTO household_members (
						id,
                        applicant_id,
                        name,
                        employment_status,
                        sex,
                        date_of_birth,
                        relation
                ) VALUES (?,?, ?, ?, ?, ?, ?)
        `, householdMember.ID, applicant.ID, householdMember.Name, householdMember.EmploymentStatus, householdMember.Sex, householdMember.DateOfBirth, householdMember.Relation)
		if insertErr != nil {
			tx.Rollback()
			return insertErr
		}
		txErr := tx.Commit()
		if txErr != nil {
			return txErr
		}
	}

	return nil
}
