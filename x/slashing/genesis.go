package slashing

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GenesisState - all slashing state that must be provided at genesis
type GenesisState struct {
	Params DefaultParams
}

// HubDefaultGenesisState - default GenesisState used by Cosmos Hub
func HubDefaultGenesisState() GenesisState {
	return GenesisState{
		Params: HubDefaultParams(),
	}
}

// InitGenesis initialize default parameter
func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) error {
	p := data.Params
	keeper.paramstore.Set(ctx, maxEvidenceAgeKey, p.MaxEvidenceAge)
	keeper.paramstore.Set(ctx, signedBlocksWindowKey, p.SignedBlocksWindow)
	keeper.paramstore.Set(ctx, minSignedPerWindowKey, p.MinSignedPerWindow)
	keeper.paramstore.Set(ctx, doubleSignUnbondDurationKey, p.DoubleSignUnbondDuration)
	keeper.paramstore.Set(ctx, downtimeUnbondDurationKey, p.DowntimeUnbondDuration)
	keeper.paramstore.Set(ctx, slashFractionDoubleSignKey, p.SlashFractionDoubleSign)
	keeper.paramstore.Set(ctx, slashFractionDowntimeKey, p.SlashFractionDowntime)
	return nil
}