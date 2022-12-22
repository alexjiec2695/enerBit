package enerbit

import (
	"github.com/go-redis/redis"
	"rest_app/domain/model/entities/entities"
)

type RedisServiceImpl struct {
	db *redis.Client
}

func NewRedisServiceImpl(db *redis.Client) *RedisServiceImpl {
	return &RedisServiceImpl{db: db}
}

func (r *RedisServiceImpl) PublishTicket(data entities.EnerBitEntities) error {
	e := r.db.XAdd(&redis.XAddArgs{
		Stream: "Creation",
		ID:     data.ID,
		Values: map[string]interface{}{
			"ID":               data.ID,
			"Brand":            data.Brand,
			"Address":          data.Address,
			"InstallationDate": data.InstallationDate,
			"RetirementDate":   data.RetirementDate,
			"Serial":           data.Serial,
			"Lines":            data.Lines,
			"IsActive":         data.IsActive,
			"CreatedAt":        data.CreatedAt,
		},
	}).Err()

	return e
}
