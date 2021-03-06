package Models

import (
	"github.com/z4rx/search-service/Config"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func GetAllDistros(distro *[]Distro) (err error) {
	if err = Config.DB.Find(distro).Error; err != nil {
		return err
	}
	return nil
}

func GetDistroByID(distro *Distro, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(distro).Error; err != nil {
		return err
	}
	return nil
}

func GetDistroByName(distro *Distro, name string) (err error) {
	if err = Config.DB.Where("name = ?", name).First(distro).Error; err != nil {
		return err
	}
	return nil
}
