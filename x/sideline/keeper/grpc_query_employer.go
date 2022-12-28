package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"sideline/x/sideline/types"
)

func (k Keeper) EmployerAll(c context.Context, req *types.QueryAllEmployerRequest) (*types.QueryAllEmployerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var employers []types.Employer
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	employerStore := prefix.NewStore(store, types.KeyPrefix(types.EmployerKeyPrefix))

	pageRes, err := query.Paginate(employerStore, req.Pagination, func(key []byte, value []byte) error {
		var employer types.Employer
		if err := k.cdc.Unmarshal(value, &employer); err != nil {
			return err
		}

		employers = append(employers, employer)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllEmployerResponse{Employer: employers, Pagination: pageRes}, nil
}

func (k Keeper) Employer(c context.Context, req *types.QueryGetEmployerRequest) (*types.QueryGetEmployerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetEmployer(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetEmployerResponse{Employer: val}, nil
}
