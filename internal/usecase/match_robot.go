package usecase

import (
	"context"
	"fmt"

	"github.com/b1g-nguyx/ADRDMS/shared/go/pkg/db"
)

type MatchRobotUseCase struct {
	postgres *db.Postgres
}

func NewMatchRobotUseCase(p *db.Postgres) *MatchRobotUseCase {
	return &MatchRobotUseCase{postgres: p}
}

func (u *MatchRobotUseCase) MatchRobot(ctx context.Context, orderID string, lat, lng float64) (string, error) {
	// Luồng xử lý:
	// 1. Chạy Query SQL:
	//    SELECT id, location, current_battery, ST_Distance(...) as distance
	//    FROM robots
	//    WHERE ST_DWithin(location, ST_MakePoint(:lng, :lat), :radius)
	//      AND status IN ('IDLE', 'EN_ROUTE')
	// 2. Chuyển kết quả về Go.
	// 3. Lặp qua các ứng viên, tính pin: BatteryConsumed = (D*2.5) + (D*S*0.3) + 5
	// 4. Lọc ra robot có (current_battery - BatteryConsumed) >= 10.
	// 5. Sort theo distance và trả về robot tốt nhất.
	// Xử lý 3 chiến lược: Thử bán kính 5km (IDLE) -> 3km (EN_ROUTE) -> 10km.
	
	// TODO: Thực hiện các bước trên.
	return "robot-id-placeholder", nil
}
