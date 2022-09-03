package domain_test

import (
	"context"
	"github.com/adamluzsi/testcase"
	"github.com/google/uuid"
	"github.com/mikejeuga/basketball_club/domain"
	"github.com/mikejeuga/basketball_club/models"
	"testing"
)

func TestService(t *testing.T) {
	s := testcase.NewSpec(t)
	ctx := context.Background()
	club := domain.NewClub()

	player := testcase.Let(s, func(t *testcase.T) *models.Player {
		return models.NewPlayer("Michael", "Jordan", models.Birthday{
			Day:   5,
			Month: 7,
			Year:  1989,
		})
	})

	s.Describe("Given the club wants to sign a player", func(s *testcase.Spec) {
		s.When("the club registers a player", func(s *testcase.Spec) {
			act := func(t *testcase.T) (uuid.UUID, error) {
				return club.Register(ctx, *player.Get(t))
			}
			s.Then("there is no error", func(t *testcase.T) {
				id, err := act(t)
				t.Must.NoError(err)

				act2 := func(t *testcase.T) (models.Player, error) {
					return club.FindPlayerBy(ctx, id)
				}

				returnedPlayer, err := act2(t)
				t.Must.NoError(err)
				t.Must.Equal(player.Get(t).FirstName, returnedPlayer.FirstName)
				t.Must.Equal(player.Get(t).LastName, returnedPlayer.LastName)
			})
		})
	})
}
