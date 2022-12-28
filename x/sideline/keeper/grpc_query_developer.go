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

func (k Keeper) DeveloperAll(c context.Context, req *types.QueryAllDeveloperRequest) (*types.QueryAllDeveloperResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var developers []types.Developer
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	developerStore := prefix.NewStore(store, types.KeyPrefix(types.DeveloperKeyPrefix))

	pageRes, err := query.Paginate(developerStore, req.Pagination, func(key []byte, value []byte) error {
		var developer types.Developer
		if err := k.cdc.Unmarshal(value, &developer); err != nil {
			return err
		}

		developers = append(developers, developer)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllDeveloperResponse{Developer: developers, Pagination: pageRes}, nil
}

func (k Keeper) Developer(c context.Context, req *types.QueryGetDeveloperRequest) (*types.QueryGetDeveloperResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetDeveloper(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetDeveloperResponse{Developer: val}, nil
}
