package data

import (
	"errors"

	"git.hotmilk.space/maki/foxlib/foxcache"
	"go.etcd.io/bbolt"
)

var (
	Database *bbolt.DB

	CACHE_BUCKET = []byte("cache")
)

func Init() error {
	var err error
	Database, err = bbolt.Open("data.db", 0600, nil)
	if err != nil {
		return errors.New("failed to open database: " + err.Error())
	}

	// ensure bucket
	err = Database.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(CACHE_BUCKET)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return errors.New("failed to create bucket: " + err.Error())
	}

	return foxcache.Init(Database, CACHE_BUCKET, []foxcache.DataInterface{
		&Blahaj, &GitHubStars,
	})
}
