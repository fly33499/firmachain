package nft

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/firmachain/firmachain/x/nft/keeper"
	"github.com/firmachain/firmachain/x/nft/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	// Set all the nftItem
	for _, elem := range genState.NftItemList {
		k.SetNftItem(ctx, *elem)
	}

	// Set nftItem count
	k.SetNftItemCount(ctx, genState.NftItemCount)

	// this line is used by starport scaffolding # ibc/genesis/init
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// this line is used by starport scaffolding # genesis/module/export
	// Get all nftItem
	nftItemList := k.GetAllNftItem(ctx)
	for _, elem := range nftItemList {
		elem := elem
		genesis.NftItemList = append(genesis.NftItemList, &elem)
	}

	// Set the current count
	genesis.NftItemCount = k.GetNftItemCount(ctx)

	// this line is used by starport scaffolding # ibc/genesis/export

	return genesis
}
