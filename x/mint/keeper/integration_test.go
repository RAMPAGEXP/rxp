package keeper_test

import (
	"encoding/json"

	"github.com/tendermint/starport/starport/pkg/cosmoscmd"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"

	rxpapp "github.com/RAMPAGEXP/rxp/v7/app"
	"github.com/cosmos/cosmos-sdk/simapp"

	"github.com/RAMPAGEXP/rxp/v7/x/mint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// returns context and an app with updated mint keeper
func createTestApp(isCheckTx bool) (*rxpapp.App, sdk.Context) {
	app := setup(isCheckTx)

	ctx := app.BaseApp.NewContext(isCheckTx, tmproto.Header{})
	app.MintKeeper.SetParams(ctx, types.DefaultParams())
	app.MintKeeper.SetMinter(ctx, types.DefaultInitialMinter())

	return app, ctx
}

func setup(isCheckTx bool) *rxpapp.App {
	app, genesisState := genApp(!isCheckTx, 5)
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
				ConsensusParams: simapp.DefaultConsensusParams,
				AppStateBytes:   stateBytes,
			},
		)
	}

	return app
}

func genApp(withGenesis bool, invCheckPeriod uint) (*rxpapp.App, rxpapp.GenesisState) {
	db := dbm.NewMemDB()
	encCdc := cosmoscmd.MakeEncodingConfig(rxpapp.ModuleBasics)
	app := rxpapp.New(
		log.NewNopLogger(),
		db,
		nil,
		true,
		map[int64]bool{},
		simapp.DefaultNodeHome,
		invCheckPeriod,
		encCdc,
		simapp.EmptyAppOptions{})

	originalApp := app.(*rxpapp.App)
	if withGenesis {
		return originalApp, rxpapp.NewDefaultGenesisState(encCdc.Marshaler)
	}

	return originalApp, rxpapp.GenesisState{}
}
