package spec

import (
	"context"
	"github.com/adamluzsi/testcase"
	"github.com/google/uuid"
	"github.com/mikejeuga/basketball_club/models"
	"testing"
)

type Club interface {
	Register(ctx context.Context, player models.Player) (uuid.UUID, error)
	FindPlayerBy(ctx context.Context, ID uuid.UUID) (models.Player, error)
}

type RegisterPlayerSpec struct {
	Club Club
}

func NewRegisterPlayerSpec(club Club) *RegisterPlayerSpec {
	return &RegisterPlayerSpec{Club: club}
}

func (spec RegisterPlayerSpec) AddPlayer(t *testing.T) {
	s := testcase.NewSpec(t)
	ctx := context.Background()

	s.Describe("Happy Path", func(s *testcase.Spec) {
		player1 := testcase.Var[*models.Player]{
			ID: "First Player",
			Init: func(t *testcase.T) *models.Player {
				player := models.NewPlayer("Michael", "Jordan", models.Birthday{
					Day:   30,
					Month: 10,
					Year:  1991,
				})
				return &player
			},
		}

		s.When("the club wants to register the player,", func(s *testcase.Spec) {
			act := func(t *testcase.T) (uuid.UUID, error) {
				return spec.Club.Register(ctx, *player1.Get(t))
			}
			s.Then("the registration goes smoothly.", func(t *testcase.T) {
				id, err := act(t)
				t.Must.NoError(err)

				s.And("The player is now part of the club", func(s *testcase.Spec) {
					act := func(t *testcase.T) (models.Player, error) {
						return spec.Club.FindPlayerBy(ctx, id)
					}
					returnedPlayer, err := act(t)
					t.Must.NoError(err)
					t.Must.Equal(player1.Get(t).FirstName, returnedPlayer.FirstName)
					t.Must.Equal(player1.Get(t).LastName, returnedPlayer.LastName)
					t.Must.Equal(player1.Get(t).Dob, returnedPlayer.Dob)
				})
			})

		})

	})

	s.Describe("Unhappy Path", func(s *testcase.Spec) {
		player2 := testcase.Var[*models.Player]{
			ID: "First Player",
			Init: func(t *testcase.T) *models.Player {
				player := models.NewPlayer("Kevin", "Durant", models.Birthday{
					Day:   30,
					Month: 10,
					Year:  2010,
				})
				return &player
			},
		}
		s.When("the club wants to register the player,", func(s *testcase.Spec) {
			s.And("the player is not old enough to be part of the team", func(s *testcase.Spec) {
				act := func(t *testcase.T) (uuid.UUID, error) {
					return spec.Club.Register(ctx, *player2.Get(t))
				}
				s.Then("the player cannot be part of the team", func(t *testcase.T) {
					_, err := act(t)
					t.Must.NotNil(err)
					t.Must.Equal(models.PlayerAgeError, err)
				})
			})
		})

	})
}
