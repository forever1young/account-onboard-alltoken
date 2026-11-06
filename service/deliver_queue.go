package service

import "github.com/redis/go-redis/v9"

// DeliverQueue holds accounts pending delivery to account-pool in Redis,
// tracks the pool's pull cursor, and clears entries once ack'd.
type DeliverQueue struct {
	rdb *redis.Client
}

func NewDeliverQueue(rdb *redis.Client) *DeliverQueue {
	return &DeliverQueue{rdb: rdb}
}

func (q *DeliverQueue) Enqueue(accountID uint) error {
	panic("not implemented")
}

func (q *DeliverQueue) PullBatch(cursor string, limit int) (accountIDs []uint, nextCursor string, err error) {
	panic("not implemented")
}

func (q *DeliverQueue) Ack(accountIDs []uint) error {
	panic("not implemented")
}
