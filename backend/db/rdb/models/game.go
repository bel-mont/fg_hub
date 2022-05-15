package models

import (
	"github.com/uptrace/bun"
	"time"
)
import "github.com/google/uuid"

type Game struct {
	bun.BaseModel `bun:"table:game,alias:g"`

	ID         uuid.UUID    `bun:"id,type:uuid,pk,default:gen_random_uuid()"`
	SeriesID   uuid.UUID    `bun:"type:uuid"`
	Series     Series       `bun:"rel:belongs-to,join:series_id=id"`
	Name       string       `bun:"name,notnull"`
	Slug       string       `bun:"slug,notnull,unique"`
	Characters []*Character `bun:"rel:has-many,join:id=game_id"`
	CreatedAt  time.Time    `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt  time.Time    `bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt  time.Time    `bun:",soft_delete,nullzero"`
}
