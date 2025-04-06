package storage

import "github.com/redis/go-redis/v9"

func SetURL(id string, url string) error {
	return RDB.Set(Ctx, id, url, 0).Err()
}

func GetURL(id string) (string, error) {
	url, err := RDB.Get(Ctx, id).Result()
	if err == redis.Nil {
		return "", err
	}
	return url, err
}
