package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "sideline/testutil/keeper"
	"sideline/testutil/nullify"
	"sideline/x/sideline/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestEmployerQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.SidelineKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNEmployer(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetEmployerRequest
		response *types.QueryGetEmployerResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetEmployerRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetEmployerResponse{Employer: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetEmployerRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetEmployerResponse{Employer: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetEmployerRequest{
				Index: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Employer(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestEmployerQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.SidelineKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNEmployer(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllEmployerRequest {
		return &types.QueryAllEmployerRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.EmployerAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Employer), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Employer),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.EmployerAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Employer), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Employer),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.EmployerAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Employer),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.EmployerAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
