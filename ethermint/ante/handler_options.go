// Copyright 2021 Evmos Foundation
// This file is part of Evmos' Ethermint library.
//
// The Ethermint library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The Ethermint library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the Ethermint library. If not, see https://github.com/evmos/ethermint/blob/main/LICENSE
package ante

import (
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errortypes "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	ibckeeper "github.com/cosmos/ibc-go/v6/modules/core/keeper"

	evmtypes "gitlab.com/rarimo/rarimo-core/x/evm/types"
)

// HandlerOptions extend the SDK's AnteHandler options by requiring the IBC
// channel keeper, EVM Keeper and Fee Market Keeper.
type HandlerOptions struct {
	AccountKeeper          evmtypes.AccountKeeper
	BankKeeper             evmtypes.BankKeeper
	IBCKeeper              *ibckeeper.Keeper
	FeeMarketKeeper        FeeMarketKeeper
	EvmKeeper              EVMKeeper
	FeegrantKeeper         ante.FeegrantKeeper
	SignModeHandler        authsigning.SignModeHandler
	SigGasConsumer         func(meter sdk.GasMeter, sig signing.SignatureV2, params authtypes.Params) error
	MaxTxGasWanted         uint64
	ExtensionOptionChecker ante.ExtensionOptionChecker
	TxFeeChecker           ante.TxFeeChecker
	DisabledAuthzMsgs      []string
}

func (options HandlerOptions) validate() error {
	if options.AccountKeeper == nil {
		return errorsmod.Wrap(errortypes.ErrLogic, "account keeper is required for AnteHandler")
	}
	if options.BankKeeper == nil {
		return errorsmod.Wrap(errortypes.ErrLogic, "bank keeper is required for AnteHandler")
	}
	if options.SignModeHandler == nil {
		return errorsmod.Wrap(errortypes.ErrLogic, "sign mode handler is required for ante builder")
	}
	if options.FeeMarketKeeper == nil {
		return errorsmod.Wrap(errortypes.ErrLogic, "fee market keeper is required for AnteHandler")
	}
	if options.EvmKeeper == nil {
		return errorsmod.Wrap(errortypes.ErrLogic, "evm keeper is required for AnteHandler")
	}
	return nil
}

func newEthAnteHandler(options HandlerOptions) sdk.AnteHandler {
	return sdk.ChainAnteDecorators(
		NewEthSetUpContextDecorator(options.EvmKeeper),                         // outermost AnteDecorator. SetUpContext must be called first
		NewEthMempoolFeeDecorator(options.EvmKeeper),                           // Check eth effective gas price against minimal-gas-prices
		NewEthMinGasPriceDecorator(options.FeeMarketKeeper, options.EvmKeeper), // Check eth effective gas price against the global MinGasPrice
		NewEthValidateBasicDecorator(options.EvmKeeper),
		NewEthSigVerificationDecorator(options.EvmKeeper),
		NewEthAccountVerificationDecorator(options.AccountKeeper, options.EvmKeeper),
		NewCanTransferDecorator(options.EvmKeeper),
		NewEthGasConsumeDecorator(options.EvmKeeper, options.MaxTxGasWanted),
		NewEthIncrementSenderSequenceDecorator(options.AccountKeeper), // innermost AnteDecorator.
		NewGasWantedDecorator(options.EvmKeeper, options.FeeMarketKeeper),
		NewEthEmitEventDecorator(options.EvmKeeper), // emit eth tx hash and index at the very last ante handler.
	)
}

func newCosmosAnteHandler(options HandlerOptions) sdk.AnteHandler {
	anteHandler, err := ante.NewAnteHandler(
		ante.HandlerOptions{
			AccountKeeper:   options.AccountKeeper,
			BankKeeper:      options.BankKeeper,
			SignModeHandler: options.SignModeHandler,
			FeegrantKeeper:  options.FeegrantKeeper,
			SigGasConsumer:  options.SigGasConsumer,
		},
	)

	if err != nil {
		panic(fmt.Errorf("failed to create AnteHandler: %w", err))
	}

	return anteHandler
}
