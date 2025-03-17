package services

import (
	"FASManagementSystem/internal/database"
	"FASManagementSystem/internal/models"
	"database/sql"
	"encoding/json"
)

func GetAllSchemes() ([]models.Scheme, error) {
	db := database.GetDB()

	rows, err := db.Query(`
                SELECT
                    s.id, s.name,s.description, s.criteria,
                    b.id AS benefit_id, b.name AS benefit_name, b.amount AS benefit_amount
                FROM
                    schemes s
                LEFT JOIN
                    benefits b ON s.id = b.scheme_id;
        `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	schemeMap := make(map[string]*models.Scheme)

	for rows.Next() {
		var (
			schemeID, schemeName, description, criteriaSQL, benefitID, benefitName sql.NullString
			benefitAmount                                                          sql.NullInt32
		)

		err := rows.Scan(
			&schemeID, &schemeName, &description, &criteriaSQL, &benefitID, &benefitName, &benefitAmount,
		)
		if err != nil {
			return nil, err
		}

		scheme, exists := schemeMap[schemeID.String]
		if !exists {
			scheme = &models.Scheme{
				ID:          schemeID.String,
				Name:        schemeName.String,
				Description: description.String,
				Benefits:    []models.Benefit{},
			}

			if criteriaSQL.Valid {
				err := json.Unmarshal([]byte(criteriaSQL.String), &scheme.Criterias)
				if err != nil {
					return nil, err
				}
			}

			schemeMap[schemeID.String] = scheme
		}

		if benefitID.Valid {
			scheme.Benefits = append(scheme.Benefits, models.Benefit{
				ID:     benefitID.String,
				Name:   benefitName.String,
				Amount: int(benefitAmount.Int32),
			})
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	schemes := make([]models.Scheme, 0, len(schemeMap))
	for _, scheme := range schemeMap {
		schemes = append(schemes, *scheme)
	}

	return schemes, nil
}

func GetElligibleSchemes(applicant models.Applicant) ([]models.Scheme, error) {
	schemes, err := GetAllSchemes()
	if err != nil {
		return nil, err
	}
	var elligibleSchemes []models.Scheme
	for _, scheme := range schemes {
		if scheme.IsElligible(applicant) {
			elligibleSchemes = append(elligibleSchemes, scheme)
		}
	}
	return elligibleSchemes, nil
}

