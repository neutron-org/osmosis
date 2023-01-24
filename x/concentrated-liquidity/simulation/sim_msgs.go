package simulation

import (
	"errors"
	"fmt"
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"

	legacysimulationtype "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/osmosis-labs/osmosis/osmoutils"
	"github.com/osmosis-labs/osmosis/v14/simulation/simtypes"
	osmosimtypes "github.com/osmosis-labs/osmosis/v14/simulation/simtypes"
	clkeeper "github.com/osmosis-labs/osmosis/v14/x/concentrated-liquidity"
	"github.com/osmosis-labs/osmosis/v14/x/concentrated-liquidity/internal/math"
	clmodeltypes "github.com/osmosis-labs/osmosis/v14/x/concentrated-liquidity/model"
	cltypes "github.com/osmosis-labs/osmosis/v14/x/concentrated-liquidity/types"
)

var PoolCreationFee = sdk.NewInt64Coin("stake", 10_000_000)

func RandomMsgCreateConcentratedPool(k clkeeper.Keeper, sim *osmosimtypes.SimCtx, ctx sdk.Context) (*clmodeltypes.MsgCreateConcentratedPool, error) {
	// generate random values from -15 to 5 (accepted range: -12 to -1)
	exponentAtPriceOne := sdk.NewInt(rand.Int63n(6+15) - 15)

	// get a random sender that contains tokens including feeToksn
	sender, senderExists := sim.RandomSimAccountWithConstraint(createPoolRestriction(k, sim, ctx))
	if !senderExists {
		return nil, fmt.Errorf("no sender with two different denoms & pool creation fee exists")
	}

	// generate 3 coins, use 2 to create pool and 1 for fees. "stake" - doesnot have denom metadata
	poolCoins, ok := sim.GetRandSubsetOfKDenoms(ctx, sender, 3)
	if !ok {
		return nil, fmt.Errorf("provided sender with requested number of denoms does not exist")
	}

	// check if the sender has sufficient amount for fees
	if poolCoins.Add(PoolCreationFee).IsAnyGT(sim.BankKeeper().SpendableCoins(ctx, sender.Address)) {
		return nil, errors.New("chose an account / creation amount that didn't pass fee bar")
	}

	denom0 := poolCoins[0].Denom
	denom1 := poolCoins[1].Denom
	tickSpacing := uint64(rand.Int63n(10)) // random uint64 value from 0 to 100
	precisionFactorAtPriceOne := exponentAtPriceOne

	return &clmodeltypes.MsgCreateConcentratedPool{
		Sender:                    sender.Address.String(),
		Denom0:                    denom0,
		Denom1:                    denom1,
		TickSpacing:               tickSpacing,
		PrecisionFactorAtPriceOne: precisionFactorAtPriceOne,
	}, nil
}

func RandMsgCreatePosition(k clkeeper.Keeper, sim *osmosimtypes.SimCtx, ctx sdk.Context) (*cltypes.MsgCreatePosition, error) {
	// Random exponentAtPriceOne from -10 to 10
	// TODO(Question): Do we want to make this value really extreme?
	exponentAtPriceOne := sdk.NewInt(rand.Int63n(21) - 10)

	// get random pool
	pool_id, poolDenoms, err := getRandCLPool(k, sim, ctx)
	if err != nil {
		return nil, err
	}

	// get address that has all denoms from the randomly selected pool
	sender, tokens, senderExists := sim.SelAddrWithDenoms(ctx, poolDenoms)
	if !senderExists {
		return nil, fmt.Errorf("no sender with denoms %s exists", poolDenoms)
	}

	// Randomize tick values from minTick to maxTick
	minTick, maxTick := getRandMinMaxTicks(exponentAtPriceOne)
	lowerTick := rand.Int63n((maxTick - minTick + 1)) + (minTick) // get random lowerTick between minTick, MaxTick
	upperTick := rand.Int63n((maxTick)-lowerTick+1) + lowerTick   // get random upperTick between lowerTick, MaxTick

	// TODO(maybe): check if these are actually pool denoms
	// tokens that users are trying to create position for
	tokenDesired0 := tokens[0]
	tokenDesired1 := tokens[1]

	// get pool actual amounts
	actualAmount0, actualAmount1, err := getPoolActualAmounts(k, ctx, pool_id, lowerTick, upperTick, exponentAtPriceOne, tokenDesired0, tokenDesired1)

	// generate random tokenMinAmount0 < actualAmount0
	tokenMinAmount0 := rand.Int63n((actualAmount0.Int64()))

	// generate random tokenMinAmount1 < actualAmount1
	tokenMinAmount1 := rand.Int63n((actualAmount1.Int64()))

	return &cltypes.MsgCreatePosition{
		PoolId:          pool_id,
		Sender:          sender.Address.String(),
		LowerTick:       lowerTick,
		UpperTick:       upperTick,
		TokenDesired0:   tokenDesired0,
		TokenDesired1:   tokenDesired1,
		TokenMinAmount0: sdk.NewInt(tokenMinAmount0),
		TokenMinAmount1: sdk.NewInt(tokenMinAmount1),
	}, nil
}

func RandMsgWithdrawPosition(k clkeeper.Keeper, sim *osmosimtypes.SimCtx, ctx sdk.Context) (*cltypes.MsgWithdrawPosition, error) {
	// PseudoCode:
	// get all pool positions for a specific pool (we can randomize which pool to use)
	// - this will include, position lower_tick, upper_tick, liquidityAmt
	// randomly select 1 pool position
	// get random withdraw liquidity from [0 to existing liqudityAmt]
	return nil, nil
}

func RandMsgCollectFees(k clkeeper.Keeper, sim *osmosimtypes.SimCtx, ctx sdk.Context) (*cltypes.MsgCollectFees, error) {
	return nil, nil
}

func createPoolRestriction(k clkeeper.Keeper, sim *simtypes.SimCtx, ctx sdk.Context) simtypes.SimAccountConstraint {
	return func(acc legacysimulationtype.Account) bool {
		accCoins := sim.BankKeeper().SpendableCoins(ctx, acc.Address)
		hasTwoCoins := len(accCoins) >= 2
		hasPoolCreationFee := accCoins.AmountOf(PoolCreationFee.Denom).GT(PoolCreationFee.Amount)
		return hasTwoCoins && hasPoolCreationFee
	}
}

// getRandCLPool gets a concnerated liquidity pool with its pool denoms.
func getRandCLPool(k clkeeper.Keeper, sim *osmosimtypes.SimCtx, ctx sdk.Context) (uint64, []string, error) {
	// get all pools
	clPools, err := k.GetAllPools(ctx)
	if err != nil {
		return 0, nil, err
	}

	numPools := len(clPools)
	if numPools == 0 {
		return 0, nil, fmt.Errorf("no pools created")
	}

	pool_id := uint64(simtypes.RandLTBound(sim, numPools) + 1)

	pool := clPools[pool_id-1]

	//poolCoins := pool.GetTotalPoolLiquidity(ctx)
	poolDenoms := osmoutils.CoinsDenoms(pool.GetTotalPoolLiquidity(ctx))

	return pool_id, poolDenoms, err
}

// getPoolActualAmounts gets the actualAmount for token0 and token1 in a specific pool
func getPoolActualAmounts(k clkeeper.Keeper, ctx sdk.Context, poolId uint64, lowerTick, upperTick int64, exponentAtPriceOne sdk.Int, tokenDesired0, tokenDesired1 sdk.Coin) (sdk.Int, sdk.Int, error) {
	poolI, err := k.GetPool(ctx, poolId)
	if err != nil {
		return sdk.ZeroInt(), sdk.ZeroInt(), err
	}

	concentratedPool := poolI.(cltypes.ConcentratedPoolExtension)

	// Get pool ActualAmounts
	sqrtPriceLowerTick, sqrtPriceUpperTick, err := math.TicksToSqrtPrice(lowerTick, upperTick, exponentAtPriceOne)
	if err != nil {
		return sdk.ZeroInt(), sdk.ZeroInt(), err
	}
	liquidityDelta := math.GetLiquidityFromAmounts(concentratedPool.GetCurrentSqrtPrice(), sqrtPriceLowerTick, sqrtPriceUpperTick, tokenDesired0.Amount, tokenDesired1.Amount)
	actualAmount0, actualAmount1 := concentratedPool.CalcActualAmounts(ctx, lowerTick, upperTick, sqrtPriceLowerTick, sqrtPriceUpperTick, liquidityDelta)

	return actualAmount0.TruncateInt(), actualAmount1.TruncateInt(), nil

}

// getRandMinMaxTicks gets min and max tick range for a specific exponentAtPriceOne.
func getRandMinMaxTicks(exponentAtPriceOne sdk.Int) (minTick, maxTick int64) {
	// Randomize this value
	return math.GetMinAndMaxTicksFromExponentAtPriceOne(exponentAtPriceOne)
}
