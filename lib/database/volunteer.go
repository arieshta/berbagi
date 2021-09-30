package libdb

import (
	"berbagi/config"
	"berbagi/models"
)

func ListVolunteers() (interface{}, int, error) {
	volunteerRes := []models.VolunteerAPI{}

	query := `SELECT volunteers.id, users.name, users.email 
				FROM volunteers
				INNER JOIN users ON volunteers.user_id = users.id;`

	tx := config.Db.Raw(query).Scan(&volunteerRes)
	if tx.Error != nil {
		return nil, 0, tx.Error
	}

	if tx.RowsAffected > 0 {
		return volunteerRes, 1, nil
	}
	return nil, 0, nil
}

func GetVolunteerProfile(volunteerId int) (interface{}, int, error) {
	volunteer := models.Volunteer{}
	tx := config.Db.Find(&volunteer, volunteerId)
	if tx.Error != nil {
		return nil, 0, tx.Error
	}

	if tx.RowsAffected > 0 {
		resVolunteer := models.ProfileVolunteerAPI{}
		query := `SELECT users.name, users.email, users.nik, volunteers.birth_date
				FROM volunteers
				INNER JOIN users ON volunteers.user_id = users.id
				WHERE volunteers.id = ?`

		findUser := config.Db.Raw(query, volunteerId).Scan(&resVolunteer)
		if findUser.Error != nil {
			return nil, 0, findUser.Error
		}

		address := models.Address{}
		findAddress := config.Db.Find(&address, volunteer.AddressID)
		if findAddress.Error != nil {
			return nil, 0, findAddress.Error
		}

		city := models.City{}
		findCity := config.Db.Find(&city, address.CityID)
		if findCity.Error != nil {
			return nil, 0, findCity.Error
		}

		provinces := models.Province{}
		findProvince := config.Db.Find(&provinces, city.ProvinceID)
		if findProvince.Error != nil {
			return nil, 0, findProvince.Error
		}

		proficiency := models.Proficiency{}
		findProficiency := config.Db.Find(&proficiency, volunteer.ProficiencyID)
		if findProficiency.Error != nil {
			return nil, 0, findProficiency.Error
		}

		resVolunteer.ProficiencyName = proficiency.Name
		resVolunteer.AddressName = address.Name
		resVolunteer.CityName = city.Name
		resVolunteer.ProvinceName = provinces.Name

		return resVolunteer, 1, nil
	}
	return nil, 0, nil
}
