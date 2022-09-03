package spec

import (
	"context"
	"github.com/mikejeuga/basketball_club/models"
)

type Club interface {
	Register(ctx context.Context, player models.Player) (models.ID, error)
	FindPlayerBy(ctx context.Context, ID models.ID) (models.Player, error)
}
