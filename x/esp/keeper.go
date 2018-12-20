package esp

import (
	"bytes"
	"fmt"
	"github.com/btcsuite/btcutil/bech32"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"gitlab.com/regen-network/regen-ledger/x/agent"
	"gitlab.com/regen-network/regen-ledger/x/proposal"

	//"github.com/twpayne/go-geom/encoding/ewkb"
)

type Keeper struct {
	espStoreKey       sdk.StoreKey
	espResultStoreKey sdk.StoreKey

	agentKeeper agent.Keeper

	cdc *codec.Codec
}

func (keeper Keeper) CanHandle(action proposal.ProposalAction) bool {
	switch action.(type) {
	case ActionRegisterESPVersion:
		return true
	case ActionReportESPResult:
		return true
	default:
		return false
	}
}

func (keeper Keeper) Handle(ctx sdk.Context, action proposal.ProposalAction, approvers []sdk.AccAddress) sdk.Result {
	switch action := action.(type) {
	case ActionRegisterESPVersion:
		return keeper.RegisterESPVersion(ctx, action.Curator, action.Name, action.Version, action.Spec, approvers)
	case ActionReportESPResult:
		return keeper.ReportESPResult(ctx, action.Curator, action.Name, action.Version, action.Verifier, action.Result, approvers)
	default:
		errMsg := fmt.Sprintf("Unrecognized data action type: %v", action.Type())
		return sdk.ErrUnknownRequest(errMsg).Result()
	}
}

func NewKeeper(
	espStoreKey sdk.StoreKey,
	espResultStoreKey sdk.StoreKey,
	cdc *codec.Codec) Keeper {
	return Keeper{
		espStoreKey:       espStoreKey,
		espResultStoreKey: espResultStoreKey,
		cdc:               cdc,
	}
}

func espKey(curator agent.AgentId, name string, version string) string {
	k, err := bech32.Encode("", curator)
	if err != nil {
		panic(err)
	}
	return k + "/" + name + "/" + version
}

func (keeper Keeper) RegisterESPVersion(ctx sdk.Context, curator agent.AgentId, name string, version string, spec ESPVersionSpec, signers []sdk.AccAddress) sdk.Result {
	// TODO consume gas

	key := espKey(curator, name, version)
	store := ctx.KVStore(keeper.espStoreKey)
	if store.Has([]byte(key)) {
		return sdk.Result{
			Code: sdk.CodeUnknownRequest,
		}
	}

	if !keeper.agentKeeper.Authorize(ctx, curator, signers) {
		return sdk.Result{
			Code: sdk.CodeUnauthorized,
		}
	}

	bz, err := keeper.cdc.MarshalBinaryBare(spec)
	if err != nil {
		panic(err)
	}

	store.Set([]byte(key), bz)

	return sdk.Result{Code: sdk.CodeOK}
}

func (keeper Keeper) GetESPVersion(ctx sdk.Context, curator agent.AgentId, name string, version string) (spec ESPVersionSpec, err sdk.Error) {
	key := espKey(curator, name, version)
	store := ctx.KVStore(keeper.espStoreKey)
	bz := store.Get([]byte(key))
	marshalErr := keeper.cdc.UnmarshalBinaryBare(bz, &spec)
	if marshalErr != nil {
		return spec, sdk.ErrUnknownRequest(marshalErr.Error())
	}
	return spec, nil
}

func (keeper Keeper) ReportESPResult(ctx sdk.Context, curator agent.AgentId, name string, version string, verifier agent.AgentId, result ESPResult, signers []sdk.AccAddress) sdk.Result {
	// TODO consume gas
	spec, err := keeper.GetESPVersion(ctx, curator, name, version)

	if err != nil {
		return sdk.Result{
			Code: sdk.CodeUnknownRequest,
		}
	}

	canVerify := false

	verifiers := spec.Verifiers

	n := len(verifiers)

	for i := 0; i < n; i++ {
		if bytes.Compare(verifier, verifiers[i]) == 0 {
			canVerify = true
			break
		}
	}

	if !canVerify {
		return sdk.Result{
			Code: sdk.CodeUnauthorized,
		}
	}

	if !keeper.agentKeeper.Authorize(ctx, verifier, signers) {
		return sdk.Result{
			Code: sdk.CodeUnauthorized,
		}
	}

	// TODO verify geometry
	// TODO verify schema

	return sdk.Result{Code: sdk.CodeOK}
}