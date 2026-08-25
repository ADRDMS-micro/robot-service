package worker

import (
	"context"
	"time"

	"github.com/b1g-nguyx/ADRDMS/shared/go/pkg/db"
	"github.com/b1g-nguyx/ADRDMS/shared/go/pkg/kafka"
)

type OutboxRelay struct {
	postgres *db.Postgres
	producer *kafka.Producer
}

func NewOutboxRelay(p *db.Postgres, producer *kafka.Producer) *OutboxRelay {
	return &OutboxRelay{postgres: p, producer: producer}
}

func (r *OutboxRelay) Start(ctx context.Context) {
	ticker := time.NewTicker(500 * time.Millisecond)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				r.processOutbox(ctx)
			}
		}
	}()
}

func (r *OutboxRelay) processOutbox(ctx context.Context) {
	// SELECT * FROM outbox_events WHERE status = 'PENDING' FOR UPDATE SKIP LOCKED
	// Kafka.Publish()
	// UPDATE status = 'PUBLISHED'
	
	// TODO: Thực hiện luồng xử lý
}
