package repositories

import (
	"log"

	"github.com/jmoiron/sqlx"
)

type IAdminRepository interface {
	PostAdmin(uniqname string) error
	IsAdmin(uniqname string) bool
	DeleteAdmin(uniqname string) error
}

type AdminRepository struct {
	Database *sqlx.DB
}

func DefaultAdminRepository(db *sqlx.DB) *AdminRepository {
	return &AdminRepository{
		Database: db,
	}
}

func (ar *AdminRepository) PostAdmin(uniqname string) error {
	_, err := ar.Database.Exec(`INSERT INTO admins=?`, uniqname)
	if err != nil {
		log.Println("Error in PostAdmin: ", err)
		return err
	}
	return nil
}

func (ar *AdminRepository) IsAdmin(uniqname string) bool {
	_, err := ar.Database.Exec(`SELECT admin_uniqname from admins WHERE admin_uniqname=?`, uniqname)
	if err == nil {
		return true
	}
	return false
}

func (ar *AdminRepository) DeleteAdmin(uniqname string) error {
	_, err := ar.Database.Exec(`DELETE from admins WHERE admin_uniqname=?`, uniqname)
	if err != nil {
		log.Println("Error in DeleteAdmin: ", err)
		return err
	}
	return nil
}
