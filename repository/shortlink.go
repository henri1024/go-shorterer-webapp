package repository

import "shorterer/model"

func (db *dbm) Save(shortlink *model.ShortLink, flag bool) error {
	done := false

	for !done {
		err := db.PsqlDB.Model(&model.ShortLink{}).Save(shortlink).Error
		if flag {
			return err
		}
		done = err == nil
	}
	return nil
}

func (db *dbm) Delete(sourceKey string) error {
	return nil
}

func (db *dbm) GetDestination(key string) (string, error) {
	sl := &model.ShortLink{}
	err := db.PsqlDB.Model(sl).Where(&model.ShortLink{SourceKey: key}).Find(sl).Error
	return sl.DestinationValue, err
}
