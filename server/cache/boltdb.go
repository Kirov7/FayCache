package cache

import (
	"github.com/boltdb/bolt"
	"log"
)

type boltCache struct {
	db *bolt.DB
}

func (b *boltCache) Set(k string, v []byte) error {
	err := b.db.Update(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte("cache"))

		if b != nil {
			err := b.Put([]byte(k), v)
			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

func (b *boltCache) Get(k string) ([]byte, error) {
	var data []byte
	err := b.db.View(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte("cache"))

		if b != nil {
			data = b.Get([]byte(k))
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (b *boltCache) Del(k string) error {
	err := b.db.Update(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte("cache"))

		if b != nil {
			err := b.Delete([]byte(k))
			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

func (b *boltCache) GetStat() Stat {
	var stat Stat
	err := b.db.Update(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte("cache"))

		if b != nil {
			bucketStats := b.Stats()
			stat.Count = int64(bucketStats.KeyN)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return stat
}

func newBoltCache() *boltCache {
	db, err := bolt.Open("./ondisk/my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Update(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte("cache"))
		if b == nil {

			_, err := tx.CreateBucket([]byte("cache"))
			if err != nil {
				log.Fatal(err)
			}
		}
		return nil
	})

	//更新数据库失败
	if err != nil {
		log.Fatal(err)
	}
	return &boltCache{
		db: db,
	}
}
