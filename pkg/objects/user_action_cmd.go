package objects

import (
	"time"
	"worker-service/pkg/constant"
)

type UserActionCmd struct {
	ID        int                 `json:"id"`
	Username  string              `json:"username"`
	Action    constant.UserAction `json:"action"`
	CreatedAt time.Time           `json:"createdAt"`
	RequestID string              `json:"requestId"`
	Email     string              `json:"email"`
}
