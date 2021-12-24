package repository

import (
	"gorm.io/gorm"
	"question-2/model"
)

type mysqlAreaRepository struct {
	DB *gorm.DB
}

func NewMysqlAreaRepository(DB *gorm.DB) model.AreaRepository {
	return &mysqlAreaRepository{DB}
}

func (r *mysqlAreaRepository) InsertArea(param1 int32, param2 int64, content string) (err error) {
	ar := &model.Area{}
	inst := r.DB.Model(ar)
	var area int64
	area = 0

	switch content {
	case "persegi panjang":
		area = int64(param1) * param2
		ar.AreaValue = area
		ar.AreaType = "persegi panjang"
		err = inst.Create(&ar).Error
		if err != nil {
			return err
		}
	case "persegi":
		area = int64(param1) * param2
		ar.AreaValue = area
		ar.AreaType = "persegi"
		err = inst.Create(&ar).Error
		if err != nil {
			return err
		}
	case "segitiga":
		area = int64(param1) * param2
		ar.AreaValue = area
		ar.AreaType = "segitiga"
		err = inst.Create(&ar).Error
		if err != nil {
			return err
		}
	default:
		ar.AreaValue = 0
		ar.AreaType = "undefined data"
		err = inst.Create(&ar).Error
		if err != nil {
			return err
		}
	}

	return nil
}
