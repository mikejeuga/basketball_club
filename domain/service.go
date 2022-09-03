package domain

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/mikejeuga/basketball_club/models"
)

type Roster map[uuid.UUID]models.Player

type Club struct {
	Players Roster
}

func (c *Club) Register(ctx context.Context, player models.Player) (uuid.UUID, error) {
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return [16]byte{}, err
	}
	player.ID = newUUID
	c.Players[player.ID] = player
	return player.ID, nil
}

func (c *Club) FindPlayerBy(ctx context.Context, ID uuid.UUID) (models.Player, error) {
	player, found := c.Players[ID]
	if !found {
		return models.Player{}, fmt.Errorf("no player registered with this ID: %v", ID)
	}

	return player, nil
}

func NewClub() *Club {
	roster := make(Roster)
	return &Club{Players: roster}
}
