package models

import (
	"github.com/uptrace/bun"
	"time"
)
import "github.com/google/uuid"

type Character struct {
	bun.BaseModel `bun:"table:character,alias:c"`

	ID        uuid.UUID `bun:"id,type:uuid,pk,default:gen_random_uuid()"`
	GameID    uuid.UUID `bun:"type:uuid"`
	Game      Game      `bun:"rel:belongs-to,join:game_id=id"`
	Name      string    `bun:"name,notnull"`
	Slug      string    `bun:"slug,notnull,unique"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt time.Time `bun:",soft_delete,nullzero"`
}
