package spec

import (
	"context"
	"github.com/google/uuid"
	"github.com/mikejeuga/basketball_club/models"
)

type Club interface {
	Register(ctx context.Context, player models.Player) (uuid.UUID, error)
	FindPlayerBy(ctx context.Context, ID uuid.UUID) (models.Player, error)
}
