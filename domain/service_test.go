package domain_test

import (
	"github.com/adamluzsi/testcase"
	"github.com/mikejeuga/basketball_club/domain"
	"github.com/mikejeuga/basketball_club/spec"
	"testing"
)

func TestService(t *testing.T) {
	s := testcase.NewSpec(t)
	club := testcase.Let(s, func(t *testcase.T) *domain.Club {
		return domain.NewClub()
	})

	s.Test("Player Registration", func(tc *testcase.T) {
		registerPlayerSpec := spec.NewRegisterPlayerSpec(club.Get(tc))
		registerPlayerSpec.Add_Player(t)
	})
}
