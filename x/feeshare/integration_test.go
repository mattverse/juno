package feeshare_test

import (
	"encoding/json"

	"github.com/CosmWasm/wasmd/x/wasm"
	dbm "github.com/cometbft/cometbft-db"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/libs/log"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"

	junoapp "github.com/CosmosContracts/juno/v15/app"

	"github.com/CosmosContracts/juno/v15/x/mint/types"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// returns context and an app with updated mint keeper
func CreateTestApp(isCheckTx bool) (*junoapp.App, sdk.Context) {
	app := Setup(isCheckTx)

	ctx := app.BaseApp.NewContext(isCheckTx, tmproto.Header{})
	app.AppKeepers.MintKeeper.SetParams(ctx, types.DefaultParams())
	app.AppKeepers.MintKeeper.SetMinter(ctx, types.DefaultInitialMinter())

	return app, ctx
}

func Setup(isCheckTx bool) *junoapp.App {
	app, genesisState := GenApp(!isCheckTx)
	if !isCheckTx {
		// init chain must be called to stop deliverState from being nil
		stateBytes, err := json.MarshalIndent(genesisState, "", " ")
		if err != nil {
			panic(err)
		}

		// Initialize the chain
		app.InitChain(
			abci.RequestInitChain{
				Validators:      []abci.ValidatorUpdate{},
				ConsensusParams: simtestutil.DefaultConsensusParams,
				AppStateBytes:   stateBytes,
			},
		)
	}

	return app
}

func GenApp(withGenesis bool) (*junoapp.App, junoapp.GenesisState) {
	db := dbm.NewMemDB()
	encCdc := junoapp.MakeEncodingConfig()

	var emptyWasmOpts []wasm.Option

	app := junoapp.New(
		log.NewNopLogger(),
		db,
		nil,
		true,
		wasm.EnableAllProposals,
		simtestutil.EmptyAppOptions{},
		emptyWasmOpts,
	)

	if withGenesis {
		return app, junoapp.NewDefaultGenesisState(encCdc.Marshaler)
	}

	return app, junoapp.GenesisState{}
}
