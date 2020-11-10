package repository

import (
	"go-shorterer/model"
	"strings"
)

func (db *dbm) SaveShortlink(shortlink *model.ShortLink, flag bool) error {
	done := false

	for !done {
		err := db.PsqlDB.Model(&model.ShortLink{}).Save(shortlink).Error
		if err != nil && !flag {
			return err
		} else if err != nil && !strings.Contains(err.Error(), "duplicate key value") && flag {
			return err
		}

		done = err == nil
	}
	return nil
}

func (db *dbm) DeleteShortlink(sourceKey string) error {
	return nil
}

func (db *dbm) GetDestination(key string) (string, error) {
	sl := &model.ShortLink{}
	err := db.PsqlDB.Model(sl).Where(&model.ShortLink{SourceKey: key}).Find(sl).Error
	return sl.DestinationValue, err
}
