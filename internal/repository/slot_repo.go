package repository

import (
	"context"

	"github.com/b1g-nguyx/ADRDMS/shared/go/pkg/db"
	"github.com/jackc/pgx/v5"
)

type SlotRepo struct {
	postgres *db.Postgres
}

func NewSlotRepo(p *db.Postgres) *SlotRepo {
	return &SlotRepo{postgres: p}
}

func (r *SlotRepo) ReserveSlot(ctx context.Context, robotID string, eventPayload []byte) error {
	return r.postgres.WithTransaction(ctx, func(tx pgx.Tx) error {
		// Bắt đầu Transaction (tx)
		// SQL: SELECT available_slots FROM robot_slots WHERE robot_id = $1 FOR UPDATE;
		// Nếu slots > 0: UPDATE robot_slots SET available_slots = available_slots - 1
		// Insert vào bảng outbox_events (RobotAssignedEvent) trong CÙNG tx.
		
		// TODO: Thực hiện logic query ở trên
		return nil
	})
}
