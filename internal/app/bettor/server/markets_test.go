package server_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/bufbuild/connect-go"
	api "github.com/elh/bettor/api/bettor/v1alpha"
	"github.com/elh/bettor/internal/app/bettor/repo/mem"
	"github.com/elh/bettor/internal/app/bettor/server"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestCreateMarket(t *testing.T) {
	user := &api.User{
		Id:          uuid.NewString(),
		Username:    "rusty",
		Centipoints: 100,
	}
	testCases := []struct {
		desc      string
		market    *api.Market
		expectErr bool
	}{
		{
			desc: "basic case",
			market: &api.Market{
				Title:   "Will I PB?",
				Creator: user.Id,
				Type: &api.Market_Pool{
					Pool: &api.Pool{
						Outcomes: []*api.Outcome{
							{Title: "Yes"},
							{Title: "No"},
						},
					},
				},
			},
		},
		{
			desc: "fails if title not set",
			market: &api.Market{
				Creator: user.Id,
				Type: &api.Market_Pool{
					Pool: &api.Pool{
						Outcomes: []*api.Outcome{
							{Title: "Yes"},
							{Title: "No"},
						},
					},
				},
			},
			expectErr: true,
		},
		{
			desc: "fails if creator is not an existing user",
			market: &api.Market{
				Title:   "Will I PB?",
				Creator: "other",
				Type: &api.Market_Pool{
					Pool: &api.Pool{
						Outcomes: []*api.Outcome{
							{Title: "Yes"},
							{Title: "No"},
						},
					},
				},
			},
			expectErr: true,
		},
		{
			desc: "fails if type not implemented",
			market: &api.Market{
				Title:   "Will I PB?",
				Creator: user.Id,
			},
			expectErr: true,
		},
		{
			desc: "fails if pool has less than 2 outcomes",
			market: &api.Market{
				Title:   "Will I PB?",
				Creator: user.Id,
				Type: &api.Market_Pool{
					Pool: &api.Pool{
						Outcomes: []*api.Outcome{
							{Title: "Yes"},
						},
					},
				},
			},
			expectErr: true,
		},
	}
	for _, tC := range testCases {
		tC := tC
		t.Run(tC.desc, func(t *testing.T) {
			s := server.New(&mem.Repo{Users: []*api.User{user}})
			out, err := s.CreateMarket(context.Background(), connect.NewRequest(&api.CreateMarketRequest{Market: tC.market}))
			if tC.expectErr {
				require.NotNil(t, err)
				return
			}
			require.Nil(t, err)

			assert.NotEmpty(t, out)
		})
	}
}

func TestGetMarket(t *testing.T) {
	market := &api.Market{
		Id: uuid.NewString(),
	}
	testCases := []struct {
		desc      string
		marketID  string
		expected  *api.Market
		expectErr bool
	}{
		{
			desc:     "basic case",
			marketID: market.Id,
			expected: market,
		},
		{
			desc:      "fails if market does not exist",
			marketID:  "does-not-exist",
			expectErr: true,
		},
		{
			desc:      "fails if id is empty",
			marketID:  "",
			expectErr: true,
		},
	}
	for _, tC := range testCases {
		tC := tC
		t.Run(tC.desc, func(t *testing.T) {
			s := server.New(&mem.Repo{Markets: []*api.Market{market}})
			out, err := s.GetMarket(context.Background(), connect.NewRequest(&api.GetMarketRequest{MarketId: tC.marketID}))
			if tC.expectErr {
				require.NotNil(t, err)
				return
			}
			require.Nil(t, err)
			assert.Equal(t, tC.expected, out.Msg.GetMarket())
		})
	}
}

func TestLockMarket(t *testing.T) {
	user := &api.User{
		Id:          uuid.NewString(),
		Username:    "rusty",
		Centipoints: 100,
	}
	market := &api.Market{
		Id:      uuid.NewString(),
		Title:   "Will I PB?",
		Creator: user.Id,
		Status:  api.Market_STATUS_OPEN,
		Type: &api.Market_Pool{
			Pool: &api.Pool{
				Outcomes: []*api.Outcome{
					{Title: "Yes"},
					{Title: "No"},
				},
			},
		},
	}
	lockedMarket := &api.Market{
		Id:      uuid.NewString(),
		Title:   "Will I PB?",
		Creator: user.Id,
		Status:  api.Market_STATUS_BETS_LOCKED,
		Type: &api.Market_Pool{
			Pool: &api.Pool{
				Outcomes: []*api.Outcome{
					{Title: "Yes"},
					{Title: "No"},
				},
			},
		},
	}
	testCases := []struct {
		desc      string
		marketID  string
		expectErr bool
	}{
		{
			desc:     "basic case",
			marketID: market.Id,
		},
		{
			desc:      "fails if market does not exist",
			marketID:  "other",
			expectErr: true,
		},
		{
			desc:      "fails if market is not open",
			marketID:  lockedMarket.Id,
			expectErr: true,
		},
	}
	for _, tC := range testCases {
		tC := tC
		t.Run(tC.desc, func(t *testing.T) {
			s := server.New(&mem.Repo{Users: []*api.User{user}, Markets: []*api.Market{market, lockedMarket}})
			out, err := s.LockMarket(context.Background(), connect.NewRequest(&api.LockMarketRequest{MarketId: tC.marketID}))
			if tC.expectErr {
				require.NotNil(t, err)
				return
			}
			require.Nil(t, err)
			assert.Equal(t, api.Market_STATUS_BETS_LOCKED, out.Msg.GetMarket().GetStatus())

			got, err := s.GetMarket(context.Background(), connect.NewRequest(&api.GetMarketRequest{MarketId: tC.marketID}))
			require.Nil(t, err)
			assert.Equal(t, api.Market_STATUS_BETS_LOCKED, got.Msg.GetMarket().GetStatus())
		})
	}
}

func TestSettleMarket(t *testing.T) {
	user1 := &api.User{
		Id:          uuid.NewString(),
		Username:    "rusty",
		Centipoints: 1000,
	}
	user2 := &api.User{
		Id:          uuid.NewString(),
		Username:    "danny",
		Centipoints: 1000,
	}
	user3 := &api.User{
		Id:          uuid.NewString(),
		Username:    "linus",
		Centipoints: 1000,
	}
	settledMarket := &api.Market{
		Id:      uuid.NewString(),
		Title:   "Will I PB?",
		Creator: user1.Id,
		Status:  api.Market_STATUS_SETTLED,
		Type: &api.Market_Pool{
			Pool: &api.Pool{
				Outcomes: []*api.Outcome{
					{Id: "outcome-1", Title: "Yes"},
					{Id: "outcome-2", Title: "No"},
				},
			},
		},
	}
	testCases := []struct {
		desc                          string
		marketID                      string
		winnerID                      string
		markets                       []*api.Market
		bets                          []*api.Bet
		expectedBetSettledCentipoints map[string]uint64
		expectedUserCentipoints       map[string]uint64
		expectErr                     bool
	}{
		{
			desc:      "fails if market does not exist",
			marketID:  "other",
			winnerID:  "outcome-1",
			expectErr: true,
		},
		{
			desc:      "fails if market is not locked",
			markets:   []*api.Market{settledMarket},
			marketID:  settledMarket.Id,
			winnerID:  "outcome-1",
			expectErr: true,
		},
		{
			desc: "fails if winner does not exist",
			markets: []*api.Market{
				{
					Id:      "z",
					Title:   "Will I PB?",
					Creator: user1.Id,
					Status:  api.Market_STATUS_BETS_LOCKED,
					Type: &api.Market_Pool{
						Pool: &api.Pool{
							Outcomes: []*api.Outcome{
								{Id: "outcome-1", Title: "Yes", Centipoints: 100},
								{Id: "outcome-2", Title: "No", Centipoints: 100},
							},
						},
					},
				},
			},
			marketID: "z",
			winnerID: "other",
			bets: []*api.Bet{
				{Id: "a", UserId: user1.Id, MarketId: "z", Centipoints: 100, Type: &api.Bet_OutcomeId{OutcomeId: "outcome-1"}},
				{Id: "b", UserId: user2.Id, MarketId: "z", Centipoints: 100, Type: &api.Bet_OutcomeId{OutcomeId: "outcome-2"}},
			},
			expectErr: true,
		},
		{
			markets: []*api.Market{
				{
					Id:      "z",
					Title:   "Will I PB?",
					Creator: user1.Id,
					Status:  api.Market_STATUS_BETS_LOCKED,
					Type: &api.Market_Pool{
						Pool: &api.Pool{
							Outcomes: []*api.Outcome{
								{Id: "outcome-1", Title: "Yes", Centipoints: 100},
								{Id: "outcome-2", Title: "No", Centipoints: 100},
							},
						},
					},
				},
			},
			marketID: "z",
			winnerID: "outcome-1",
			bets: []*api.Bet{
				{Id: "a", UserId: user1.Id, MarketId: "z", Centipoints: 100, Type: &api.Bet_OutcomeId{OutcomeId: "outcome-1"}},
				{Id: "b", UserId: user2.Id, MarketId: "z", Centipoints: 100, Type: &api.Bet_OutcomeId{OutcomeId: "outcome-2"}},
			},
			expectedBetSettledCentipoints: map[string]uint64{
				"a": 200,
				"b": 0,
			},
			expectedUserCentipoints: map[string]uint64{
				user1.Id: 1200,
				user2.Id: 1000,
			},
		},
		{
			markets: []*api.Market{
				{
					Id:      "z",
					Title:   "Will I PB?",
					Creator: user1.Id,
					Status:  api.Market_STATUS_BETS_LOCKED,
					Type: &api.Market_Pool{
						Pool: &api.Pool{
							Outcomes: []*api.Outcome{
								{Id: "outcome-1", Title: "Yes", Centipoints: 100},
								{Id: "outcome-2", Title: "No", Centipoints: 150},
							},
						},
					},
				},
			},
			marketID: "z",
			winnerID: "outcome-1",
			bets: []*api.Bet{
				{Id: "a", UserId: user1.Id, MarketId: "z", Centipoints: 100, Type: &api.Bet_OutcomeId{OutcomeId: "outcome-1"}},
				{Id: "b", UserId: user2.Id, MarketId: "z", Centipoints: 100, Type: &api.Bet_OutcomeId{OutcomeId: "outcome-2"}},
				{Id: "c", UserId: user3.Id, MarketId: "z", Centipoints: 50, Type: &api.Bet_OutcomeId{OutcomeId: "outcome-2"}},
			},
			expectedBetSettledCentipoints: map[string]uint64{
				"a": 250,
				"b": 0,
				"c": 0,
			},
			expectedUserCentipoints: map[string]uint64{
				user1.Id: 1250,
				user2.Id: 1000,
				user3.Id: 1000,
			},
		},
		{
			markets: []*api.Market{
				{
					Id:      "z",
					Title:   "Will I PB?",
					Creator: user1.Id,
					Status:  api.Market_STATUS_BETS_LOCKED,
					Type: &api.Market_Pool{
						Pool: &api.Pool{
							Outcomes: []*api.Outcome{
								{Id: "outcome-1", Title: "Yes", Centipoints: 100},
								{Id: "outcome-2", Title: "No", Centipoints: 200},
							},
						},
					},
				},
			},
			marketID: "z",
			winnerID: "outcome-1",
			bets: []*api.Bet{
				{Id: "a", UserId: user1.Id, MarketId: "z", Centipoints: 25, Type: &api.Bet_OutcomeId{OutcomeId: "outcome-1"}},
				{Id: "b", UserId: user2.Id, MarketId: "z", Centipoints: 75, Type: &api.Bet_OutcomeId{OutcomeId: "outcome-1"}},
				{Id: "c", UserId: user3.Id, MarketId: "z", Centipoints: 200, Type: &api.Bet_OutcomeId{OutcomeId: "outcome-2"}},
			},
			expectedBetSettledCentipoints: map[string]uint64{
				"a": 75,
				"b": 225,
				"c": 0,
			},
			expectedUserCentipoints: map[string]uint64{
				user1.Id: 1075,
				user2.Id: 1225,
				user3.Id: 1000,
			},
		},
	}
	for _, tC := range testCases {
		tC := tC
		t.Run(tC.desc, func(t *testing.T) {
			s := server.New(&mem.Repo{Users: []*api.User{proto.Clone(user1).(*api.User), proto.Clone(user2).(*api.User), proto.Clone(user3).(*api.User)}, Markets: tC.markets, Bets: tC.bets})
			out, err := s.SettleMarket(context.Background(), connect.NewRequest(&api.SettleMarketRequest{MarketId: tC.marketID, Type: &api.SettleMarketRequest_WinnerId{WinnerId: tC.winnerID}}))
			if tC.expectErr {
				require.NotNil(t, err)
				return
			}
			require.Nil(t, err)
			assert.Equal(t, out.Msg.GetMarket().GetStatus(), out.Msg.GetMarket().GetStatus())
			assert.NotEmpty(t, out.Msg.GetMarket().GetSettledAt())
			assert.Equal(t, tC.winnerID, out.Msg.GetMarket().GetPool().GetWinnerId())

			got, err := s.GetMarket(context.Background(), connect.NewRequest(&api.GetMarketRequest{MarketId: tC.marketID}))
			require.Nil(t, err)
			assert.Equal(t, out.Msg.GetMarket().GetStatus(), got.Msg.GetMarket().GetStatus())
			assert.NotEmpty(t, got.Msg.GetMarket().GetSettledAt())
			assert.Equal(t, tC.winnerID, got.Msg.GetMarket().GetPool().GetWinnerId())

			for betID, cp := range tC.expectedBetSettledCentipoints {
				gotBet, err := s.GetBet(context.Background(), connect.NewRequest(&api.GetBetRequest{BetId: betID}))
				require.Nil(t, err)
				assert.NotEmpty(t, gotBet.Msg.GetBet().GetSettledAt())
				assert.Equal(t, cp, gotBet.Msg.GetBet().GetSettledCentipoints(), betID)
				fmt.Println(cp, gotBet.Msg.GetBet().GetSettledCentipoints())
			}

			for userID, cp := range tC.expectedUserCentipoints {
				gotUser, err := s.GetUser(context.Background(), connect.NewRequest(&api.GetUserRequest{UserId: userID}))
				require.Nil(t, err)
				assert.Equal(t, cp, gotUser.Msg.GetUser().GetCentipoints())
				fmt.Println(user1.Id, user2.Id, user3.Id)
				fmt.Println(gotUser.Msg.GetUser().GetCentipoints())
			}
		})
	}
}

func TestCreateBet(t *testing.T) {
	user := &api.User{
		Id:          uuid.NewString(),
		Username:    "rusty",
		Centipoints: 1000,
	}
	poolMarket := &api.Market{
		Id:      uuid.NewString(),
		Title:   "Will I PB?",
		Creator: user.Id,
		Status:  api.Market_STATUS_OPEN,
		Type: &api.Market_Pool{
			Pool: &api.Pool{
				Outcomes: []*api.Outcome{
					{Id: uuid.NewString(), Title: "Yes"},
					{Id: uuid.NewString(), Title: "No"},
				},
			},
		},
	}
	lockedPoolMarket := &api.Market{
		Id:      uuid.NewString(),
		Title:   "Will I PB?",
		Creator: user.Id,
		Status:  api.Market_STATUS_BETS_LOCKED,
		Type: &api.Market_Pool{
			Pool: &api.Pool{
				Outcomes: []*api.Outcome{
					{Id: uuid.NewString(), Title: "Yes"},
					{Id: uuid.NewString(), Title: "No"},
				},
			},
		},
	}
	settledPoolMarket := &api.Market{
		Id:      uuid.NewString(),
		Title:   "Will I PB?",
		Creator: user.Id,
		Status:  api.Market_STATUS_SETTLED,
		Type: &api.Market_Pool{
			Pool: &api.Pool{
				Outcomes: []*api.Outcome{
					{Id: uuid.NewString(), Title: "Yes"},
					{Id: uuid.NewString(), Title: "No"},
				},
			},
		},
	}
	testCases := []struct {
		desc                  string
		bet                   *api.Bet
		expectUserCentipoints uint64
		expectErr             bool
	}{
		// pool bets
		{
			desc: "basic case - pool bet",
			bet: &api.Bet{
				UserId:      user.Id,
				MarketId:    poolMarket.Id,
				Centipoints: 100,
				Type:        &api.Bet_OutcomeId{OutcomeId: poolMarket.GetPool().Outcomes[0].Id},
			},
			expectUserCentipoints: 900,
		},
		{
			desc: "fails if user does not exist",
			bet: &api.Bet{
				UserId:      "other",
				MarketId:    poolMarket.Id,
				Centipoints: 100,
				Type:        &api.Bet_OutcomeId{OutcomeId: poolMarket.GetPool().Outcomes[0].Id},
			},
			expectErr: true,
		},
		{
			desc: "fails if market does not exist",
			bet: &api.Bet{
				UserId:      user.Id,
				MarketId:    "other",
				Centipoints: 100,
				Type:        &api.Bet_OutcomeId{OutcomeId: poolMarket.GetPool().Outcomes[0].Id},
			},
			expectErr: true,
		},
		{
			desc: "fails if type not provided",
			bet: &api.Bet{
				UserId:      user.Id,
				MarketId:    poolMarket.Id,
				Centipoints: 100,
			},
			expectErr: true,
		},
		{
			desc: "fails if outcome does not exist",
			bet: &api.Bet{
				UserId:      user.Id,
				MarketId:    poolMarket.Id,
				Centipoints: 100,
				Type:        &api.Bet_OutcomeId{OutcomeId: "other"},
			},
			expectErr: true,
		},
		{
			desc: "fails if creating a bet on a locked market",
			bet: &api.Bet{
				UserId:      user.Id,
				MarketId:    lockedPoolMarket.Id,
				Centipoints: 100,
				Type:        &api.Bet_OutcomeId{OutcomeId: lockedPoolMarket.GetPool().Outcomes[0].Id},
			},
			expectErr: true,
		},
		{
			desc: "fails if creating a bet on a settled market",
			bet: &api.Bet{
				UserId:      user.Id,
				MarketId:    settledPoolMarket.Id,
				Centipoints: 100,
				Type:        &api.Bet_OutcomeId{OutcomeId: settledPoolMarket.GetPool().Outcomes[0].Id},
			},
			expectErr: true,
		},
		{
			desc: "fails if betting more points than user has",
			bet: &api.Bet{
				UserId:      user.Id,
				MarketId:    poolMarket.Id,
				Centipoints: 2000,
				Type:        &api.Bet_OutcomeId{OutcomeId: poolMarket.GetPool().Outcomes[0].Id},
			},
			expectErr: true,
		},
	}
	for _, tC := range testCases {
		tC := tC
		t.Run(tC.desc, func(t *testing.T) {
			s := server.New(&mem.Repo{Users: []*api.User{user}, Markets: []*api.Market{poolMarket, lockedPoolMarket, settledPoolMarket}})
			out, err := s.CreateBet(context.Background(), connect.NewRequest(&api.CreateBetRequest{Bet: tC.bet}))
			if tC.expectErr {
				require.NotNil(t, err)
				return
			}
			require.Nil(t, err)

			assert.NotEmpty(t, out)

			u, err := s.GetUser(context.Background(), connect.NewRequest(&api.GetUserRequest{UserId: tC.bet.GetUserId()}))
			require.Nil(t, err)
			assert.Equal(t, tC.expectUserCentipoints, u.Msg.User.GetCentipoints())

			m, err := s.GetMarket(context.Background(), connect.NewRequest(&api.GetMarketRequest{MarketId: tC.bet.GetMarketId()}))
			require.Nil(t, err)
			assert.Equal(t, tC.bet.GetCentipoints(), m.Msg.Market.GetPool().GetOutcomes()[0].GetCentipoints())
		})
	}
}

func TestGetBet(t *testing.T) {
	bet := &api.Bet{
		Id: uuid.NewString(),
	}
	testCases := []struct {
		desc      string
		betID     string
		expected  *api.Bet
		expectErr bool
	}{
		{
			desc:     "basic case",
			betID:    bet.Id,
			expected: bet,
		},
		{
			desc:      "fails if bet does not exist",
			betID:     "does-not-exist",
			expectErr: true,
		},
		{
			desc:      "fails if id is empty",
			betID:     "",
			expectErr: true,
		},
	}
	for _, tC := range testCases {
		tC := tC
		t.Run(tC.desc, func(t *testing.T) {
			s := server.New(&mem.Repo{Bets: []*api.Bet{bet}})
			out, err := s.GetBet(context.Background(), connect.NewRequest(&api.GetBetRequest{BetId: tC.betID}))
			if tC.expectErr {
				require.NotNil(t, err)
				return
			}
			require.Nil(t, err)
			assert.Equal(t, tC.expected, out.Msg.GetBet())
		})
	}
}
