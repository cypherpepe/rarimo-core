package keeper

import (
	"context"
	"math/big"
	"strconv"

	cosmostypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/crypto/operation/origin"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/types"
	"gitlab.com/rarify-protocol/rarimo-core/x/tokenmanager/saver"
	savermsg "gitlab.com/rarify-protocol/saver-grpc-lib/grpc"
)

func (k msgServer) CreateTransferOperation(goCtx context.Context, msg *types.MsgCreateTransferOp) (*types.MsgCreateTransferOpResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	origin := origin.NewDefaultOriginBuilder().
		SetTxHash(msg.Tx).
		SetOpId(msg.EventId).
		SetCurrentNetwork(msg.FromChain).
		Build().
		GetOrigin()

	index := hexutil.Encode(origin[:])

	_, isFound := k.GetOperation(
		ctx,
		index,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	infoResp, err := k.getDepositInfo(ctx, msg)
	if err != nil {
		return nil, err
	}

	if _, ok := k.tm.GetNetwork(ctx, infoResp.TargetNetwork); !ok {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "network not found: %s", infoResp.TargetNetwork)
	}

	currentItem, ok := k.tm.GetItem(ctx, infoResp.TokenAddress, infoResp.TokenId, msg.FromChain)
	if !ok {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "current token not found")
	}

	targetItem, ok := k.tm.GetItemByNetwork(ctx, currentItem.Index, infoResp.TargetNetwork)
	if !ok {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "target token not found")
	}

	var transferOp = types.Transfer{
		Origin:     index,
		Tx:         msg.Tx,
		EventId:    msg.EventId,
		FromChain:  msg.FromChain,
		ToChain:    infoResp.TargetNetwork,
		Receiver:   infoResp.Receiver,
		Amount:     castAmount(infoResp.Amount, uint8(currentItem.Decimals), uint8(targetItem.Decimals)),
		BundleData: infoResp.BundleData,
		BundleSalt: getSalt(infoResp),
		TokenIndex: currentItem.Index,
	}

	details, err := cosmostypes.NewAnyWithValue(&transferOp)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "error parsing details %s", err.Error())
	}

	var operation = types.Operation{
		Index:         index,
		OperationType: types.OpType_TRANSFER,
		Details:       details,
		Signed:        false,
		Creator:       msg.Creator,
		Timestamp:     ctx.BlockTime().Unix(),
	}

	k.SetOperation(
		ctx,
		operation,
	)

	ctx.EventManager().EmitEvent(sdk.NewEvent(types.EventTypeNewOperation,
		sdk.NewAttribute(types.AttributeKeyOperationId, operation.Index),
		sdk.NewAttribute(types.AttributeKeyOperationType, types.OpType_TRANSFER.String()),
	))
	return &types.MsgCreateTransferOpResponse{}, nil
}

func castAmount(currentAmount string, currentDecimals uint8, targetDecimals uint8) string {
	if currentDecimals == targetDecimals {
		return currentAmount
	}

	value, _ := new(big.Int).SetString(currentAmount, 10)

	if currentDecimals < targetDecimals {
		for i := uint8(0); i < targetDecimals-currentDecimals; i++ {
			value.Mul(value, new(big.Int).SetInt64(10))
		}

		return value.String()
	}

	for i := uint8(0); i < targetDecimals-currentDecimals; i++ {
		value.Div(value, new(big.Int).SetInt64(10))
	}

	return value.String()
}

func (k *Keeper) getDepositInfo(ctx sdk.Context, msg *types.MsgCreateTransferOp) (*savermsg.MsgDepositResponse, error) {
	network, ok := k.tm.GetNetwork(ctx, msg.FromChain)
	if !ok {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "network not found: %s", msg.FromChain)
	}

	infoRequest := &savermsg.MsgTransactionInfoRequest{
		Hash:    msg.Tx,
		EventId: msg.EventId,
		Type:    network.Types[strconv.Itoa(int(msg.TokenType))],
	}

	saverClient, err := saver.GetClient(msg.FromChain)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "error getting saver connection", err.Error())
	}

	infoResp, err := saverClient.GetDepositInfo(ctx.Context(), infoRequest)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "error searching deposit %s", err.Error())
	}

	return infoResp, nil
}

func getSalt(response *savermsg.MsgDepositResponse) string {
	if response.BundleData == "" || response.BundleSalt == "" {
		return ""
	}

	return hexutil.Encode(crypto.Keccak256(hexutil.MustDecode(response.BundleSalt), hexutil.MustDecode(response.Receiver)))
}
