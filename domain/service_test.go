package domain_test

import (
	"github.com/adamluzsi/testcase"
	"github.com/mikejeuga/basketball_club/domain"
	"github.com/mikejeuga/basketball_club/models"
	"github.com/mikejeuga/basketball_club/spec"
	"testing"
)

func TestService(t *testing.T) {
	s := testcase.NewSpec(t)
	club := testcase.Let(s, func(t *testcase.T) *models.Club {
		return models.NewClub()
	})

	service := testcase.Let(s, func(t *testcase.T) *domain.Service {
		return domain.NewService(club.Get(t))
	})

	s.Test("Player Registration", func(tc *testcase.T) {
		registerPlayerSpec := spec.NewRegisterPlayerSpec(service.Get(tc))
		registerPlayerSpec.Add_Player(t)
	})
}
