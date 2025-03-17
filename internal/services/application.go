package services

import (
	"FASManagementSystem/internal/database"
	"FASManagementSystem/internal/models"
)

func GetAllApplications() ([]models.Application, error) {
	db := database.GetDB()
	rows, err := db.Query(`
        SELECT
            id,status,applicant_id,scheme_id
        FROM
            applications
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var applications []models.Application
	for rows.Next() {
		var (
			id, status, applicant_id, scheme_id string
		)
		err := rows.Scan(
			&id, &status, &applicant_id, &scheme_id,
		)
		if err != nil {
			return nil, err
		}
		applications = append(applications, models.Application{
			ID:          id,
			Status:      status,
			ApplicantID: applicant_id,
			SchemeID:    scheme_id,
		})
	}
	return applications, nil
}

func CreateNewApplication(application models.Application) error {
	db := database.GetDB()
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(`
		INSERT INTO applications (id, status, applicant_id, scheme_id)
		VALUES (?, ?, ?, ?)
	`, application.ID, application.Status, application.ApplicantID, application.SchemeID)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}
