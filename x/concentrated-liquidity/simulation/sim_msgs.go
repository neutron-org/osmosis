package simulation

import (
	"fmt"
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/osmosis-labs/osmosis/osmoutils"
	"github.com/osmosis-labs/osmosis/v14/simulation/simtypes"
	osmosimtypes "github.com/osmosis-labs/osmosis/v14/simulation/simtypes"
	clkeeper "github.com/osmosis-labs/osmosis/v14/x/concentrated-liquidity"
	"github.com/osmosis-labs/osmosis/v14/x/concentrated-liquidity/internal/math"
	clmodeltypes "github.com/osmosis-labs/osmosis/v14/x/concentrated-liquidity/model"
	cltypes "github.com/osmosis-labs/osmosis/v14/x/concentrated-liquidity/types"
)

func RandomMsgCreateConcentratedPool(k clkeeper.Keeper, sim *osmosimtypes.SimCtx, ctx sdk.Context) (*clmodeltypes.MsgCreateConcentratedPool, error) {
	// fields to randomize:
	// - sender
	// - denom0
	// - denom1
	// - tickSpacing // default 1
	// - PrecisionFactorAtPriceOne // default -4
	return nil, nil
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
