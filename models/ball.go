package models

import (
	"context"
	"fmt"
	"github.com/google/uuid"
)

type Roster map[uuid.UUID]Player

type Club struct {
	Roster Roster
}

func (c *Club) Register(ctx context.Context, player Player) (uuid.UUID, error) {
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return [16]byte{}, err
	}
	player.ID = newUUID
	c.Roster[player.ID] = player
	return player.ID, nil
}

func (c *Club) FindPlayerBy(ctx context.Context, ID uuid.UUID) (Player, error) {
	player, found := c.Roster[ID]
	if !found {
		return Player{}, fmt.Errorf("no player registered with this ID: %v", ID)
	}

	return player, nil
}

func NewClub() *Club {
	roster := make(Roster)
	return &Club{Roster: roster}
}
