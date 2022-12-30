package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"gitlab.com/rarimo/rarimo-core/x/tokenmanager/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) InfoAll(c context.Context, req *types.QueryAllInfoRequest) (*types.QueryAllInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var infos []types.Info
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	infoStore := prefix.NewStore(store, types.KeyPrefix(types.InfoKeyPrefix))

	pageRes, err := query.Paginate(infoStore, req.Pagination, func(key []byte, value []byte) error {
		var info types.Info
		if err := k.cdc.Unmarshal(value, &info); err != nil {
			return err
		}

		infos = append(infos, info)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllInfoResponse{Info: infos, Pagination: pageRes}, nil
}

func (k Keeper) Info(c context.Context, req *types.QueryGetInfoRequest) (*types.QueryGetInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetInfo(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetInfoResponse{Info: val}, nil
}