package models

import (
	"github.com/uptrace/bun"
	"time"
)
import "github.com/google/uuid"

type Game struct {
	bun.BaseModel `bun:"table:games,alias:g"`

	ID        uuid.UUID `bun:"id,type:uuid,pk,default:gen_random_uuid()"`
	Name      string    `bun:"name,notnull"`
	Slug      string    `bun:"slug,notnull,unique"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt time.Time `bun:",soft_delete,nullzero"`
}
