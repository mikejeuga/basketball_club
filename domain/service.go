package domain

import (
	"context"
	"github.com/google/uuid"
	"github.com/mikejeuga/basketball_club/models"
	"github.com/mikejeuga/basketball_club/spec"
	"time"
)

type Service struct {
	Club spec.Club
}

func NewService(club spec.Club) *Service {
	return &Service{Club: club}
}

func (s *Service) Register(ctx context.Context, player models.Player) (uuid.UUID, error) {
	if !isOldEnough(player, parseDate(time.Now())) {
		return [16]byte{}, models.PlayerAgeError
	}
	return s.Club.Register(ctx, player)
}

func (s *Service) FindPlayerBy(ctx context.Context, ID uuid.UUID) (models.Player, error) {
	return s.Club.FindPlayerBy(ctx, ID)
}

func isOldEnough(player models.Player, today models.Birthday) bool {
	if today.Year-player.Dob.Year > 16 {
		return true
	}

	if (today.Year-player.Dob.Year >= 16) && (today.Month >= player.Dob.Month) && (today.Day >= player.Dob.Day) {
		return true
	}
	return false
}

func parseDate(date time.Time) models.Birthday {
	b := models.Birthday{}
	b.Day = date.Day()
	b.Month = int(date.Month())
	b.Year = date.Year()

	return b
}
