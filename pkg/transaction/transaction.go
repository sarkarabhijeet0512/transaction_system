package transaction

import (
	"time"

	"go.uber.org/fx"
)

// Module provides all constructor and invocation methods to facilitate credits module
var Module = fx.Options(
	fx.Provide(
		NewDBRepository,
		NewService,
	),
)

type (
	Transaction struct {
		ID            int64     `json:"id" pg:"id"`
		TransactionID int64     `json:"transaction_id" pg:"transaction_id,unique"`
		Amount        float64   `json:"amount" pg:"amount"`
		Type          string    `json:"type" pg:"type"`
		ParentID      *int64    `json:"parent_id,omitempty" pg:"parent_id"`
		IsActive      bool      `json:"is_active" pg:"is_active"`
		CreatedAt     time.Time `json:"created_at" pg:"created_at"`
		UpdatedAt     time.Time `json:"updated_at" pg:"updated_at"`
	}
)
