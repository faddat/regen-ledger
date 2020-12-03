package types

import (
	"fmt"
	"time"

	"github.com/cockroachdb/apd/v2"
	proto "github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"
	"github.com/regen-network/regen-ledger/math"
	"github.com/regen-network/regen-ledger/orm"
	"gopkg.in/yaml.v2"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

type GroupID uint64

func (g GroupID) Uint64() uint64 {
	return uint64(g)
}

func (g GroupID) Empty() bool {
	return g == 0
}

func (g GroupID) Bytes() []byte {
	return orm.EncodeSequence(uint64(g))
}

type ProposalID uint64

func (p ProposalID) Bytes() []byte {
	return orm.EncodeSequence(uint64(p))
}

func (p ProposalID) Uint64() uint64 {
	return uint64(p)
}
func (p ProposalID) Empty() bool {
	return p == 0
}

type DecisionPolicyResult struct {
	Allow bool
	Final bool
}

// DecisionPolicy is the persistent set of rules to determine the result of election on a proposal.
type DecisionPolicy interface {
	codec.ProtoMarshaler

	orm.Validateable
	GetTimeout() types.Duration
	Allow(tally Tally, totalPower string, votingDuration time.Duration) (DecisionPolicyResult, error)
	Validate(g GroupInfo) error
}

// Implements DecisionPolicy Interface
var _ DecisionPolicy = &ThresholdDecisionPolicy{}

// NewThresholdDecisionPolicy creates a threshold DecisionPolicy
func NewThresholdDecisionPolicy(threshold string, timeout types.Duration) DecisionPolicy {
	return &ThresholdDecisionPolicy{threshold, timeout}
}

// Allow allows a proposal to pass when the tally of yes votes equals or exceeds the threshold before the timeout.
func (p ThresholdDecisionPolicy) Allow(tally Tally, totalPower string, votingDuration time.Duration) (DecisionPolicyResult, error) {
	timeout, err := types.DurationFromProto(&p.Timeout)
	if err != nil {
		return DecisionPolicyResult{}, err
	}
	if timeout <= votingDuration {
		return DecisionPolicyResult{Allow: false, Final: true}, nil
	}

	threshold, err := math.ParsePositiveDecimal(p.Threshold)
	if err != nil {
		return DecisionPolicyResult{}, err
	}
	yesCount, err := math.ParseNonNegativeDecimal(tally.YesCount)
	if err != nil {
		return DecisionPolicyResult{}, err
	}
	if yesCount.Cmp(threshold) >= 0 {
		return DecisionPolicyResult{Allow: true, Final: true}, nil
	}

	totalPowerDec, err := math.ParseNonNegativeDecimal(totalPower)
	if err != nil {
		return DecisionPolicyResult{}, err
	}
	totalCounts, err := tally.TotalCounts()
	if err != nil {
		return DecisionPolicyResult{}, err
	}
	var undecided apd.Decimal
	err = math.SafeSub(&undecided, totalPowerDec, totalCounts)
	if err != nil {
		return DecisionPolicyResult{}, err
	}
	var sum apd.Decimal
	err = math.Add(&sum, yesCount, &undecided)
	if err != nil {
		return DecisionPolicyResult{}, err
	}
	if sum.Cmp(threshold) < 0 {
		return DecisionPolicyResult{Allow: false, Final: true}, nil
	}
	return DecisionPolicyResult{Allow: false, Final: false}, nil
}

// Validate returns an error if policy threshold is greater than the total group weight
func (p *ThresholdDecisionPolicy) Validate(g GroupInfo) error {
	threshold, err := math.ParsePositiveDecimal(p.Threshold)
	if err != nil {
		return sdkerrors.Wrap(err, "threshold")
	}
	totalWeight, err := math.ParseNonNegativeDecimal(g.TotalWeight)
	if err != nil {
		return sdkerrors.Wrap(err, "group total weight")
	}
	if threshold.Cmp(totalWeight) > 0 {
		return sdkerrors.Wrap(ErrInvalid, "policy threshold should not be greater than the total group weight")
	}
	return nil
}

func (p ThresholdDecisionPolicy) ValidateBasic() error {
	if _, err := math.ParsePositiveDecimal(p.Threshold); err != nil {
		return sdkerrors.Wrap(err, "threshold")
	}

	timeout, err := types.DurationFromProto(&p.Timeout)
	if err != nil {
		return sdkerrors.Wrap(err, "timeout")
	}

	if timeout <= time.Nanosecond {
		return sdkerrors.Wrap(ErrInvalid, "timeout")
	}
	return nil
}

func (g GroupMember) NaturalKey() []byte {
	result := make([]byte, 8, 8+len(g.Member))
	copy(result[0:8], g.GroupId.Bytes())
	result = append(result, g.Member...)
	return result
}

func (g GroupAccountInfo) NaturalKey() []byte {
	return g.GroupAccount
}

var _ orm.Validateable = GroupAccountInfo{}

// NewGroupAccountInfo creates a new GroupAccountInfo instance
func NewGroupAccountInfo(groupAccount sdk.AccAddress, group GroupID, admin sdk.AccAddress, comment string, version uint64, decisionPolicy DecisionPolicy) (GroupAccountInfo, error) {
	p := GroupAccountInfo{
		GroupAccount: groupAccount,
		GroupId:      group,
		Admin:        admin,
		Comment:      comment,
		Version:      version,
	}

	msg, ok := decisionPolicy.(proto.Message)
	if !ok {
		return GroupAccountInfo{}, fmt.Errorf("%T does not implement proto.Message", decisionPolicy)
	}

	any, err := codectypes.NewAnyWithValue(msg)
	if err != nil {
		return GroupAccountInfo{}, err
	}

	p.DecisionPolicy = any
	return p, nil
}

func (g GroupAccountInfo) GetDecisionPolicy() DecisionPolicy {
	decisionPolicy, ok := g.DecisionPolicy.GetCachedValue().(DecisionPolicy)
	if !ok {
		return nil
	}
	return decisionPolicy
}

func (g GroupAccountInfo) ValidateBasic() error {
	if g.Admin.Empty() {
		return sdkerrors.Wrap(ErrEmpty, "admin")
	}
	if err := sdk.VerifyAddressFormat(g.Admin); err != nil {
		return sdkerrors.Wrap(err, "admin")
	}
	if g.GroupAccount.Empty() {
		return sdkerrors.Wrap(ErrEmpty, "group account")
	}
	if err := sdk.VerifyAddressFormat(g.GroupAccount); err != nil {
		return sdkerrors.Wrap(err, "group account")
	}

	if g.GroupId == 0 {
		return sdkerrors.Wrap(ErrEmpty, "group")
	}
	if g.Version == 0 {
		return sdkerrors.Wrap(ErrEmpty, "version")
	}
	policy := g.GetDecisionPolicy()

	if policy == nil {
		return sdkerrors.Wrap(ErrEmpty, "policy")
	}
	if err := policy.ValidateBasic(); err != nil {
		return sdkerrors.Wrap(err, "policy")
	}
	return nil
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (g GroupAccountInfo) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	var decisionPolicy DecisionPolicy
	return unpacker.UnpackAny(g.DecisionPolicy, &decisionPolicy)
}

func (v Vote) NaturalKey() []byte {
	result := make([]byte, 8, 8+len(v.Voter))
	copy(result[0:8], v.ProposalId.Bytes())
	result = append(result, v.Voter...)
	return result
}

var _ orm.Validateable = Vote{}

func (v Vote) ValidateBasic() error {
	if len(v.Voter) == 0 {
		return sdkerrors.Wrap(ErrEmpty, "voter")
	}
	if err := sdk.VerifyAddressFormat(v.Voter); err != nil {
		return sdkerrors.Wrap(err, "voter")
	}
	if v.ProposalId == 0 {
		return sdkerrors.Wrap(ErrEmpty, "proposal")
	}
	if v.Choice == Choice_CHOICE_UNSPECIFIED {
		return sdkerrors.Wrap(ErrEmpty, "choice")
	}
	if _, ok := Choice_name[int32(v.Choice)]; !ok {
		return sdkerrors.Wrap(ErrInvalid, "choice")
	}
	t, err := types.TimestampFromProto(&v.SubmittedAt)
	if err != nil {
		return sdkerrors.Wrap(err, "submitted at")
	}
	if t.IsZero() {
		return sdkerrors.Wrap(ErrEmpty, "submitted at")
	}
	return nil
}

const defaultMaxCommentLength = 255

// Parameter keys
var (
	ParamMaxCommentLength = []byte("MaxCommentLength")
)

// DefaultParams returns the default parameters for the group module.
func DefaultParams() Params {
	return Params{
		MaxCommentLength: defaultMaxCommentLength,
	}
}

func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// ParamSetPairs returns the parameter set pairs.
func (p *Params) ParamSetPairs() paramstypes.ParamSetPairs {
	return paramstypes.ParamSetPairs{
		paramstypes.NewParamSetPair(ParamMaxCommentLength, &p.MaxCommentLength, noopValidator()),
	}
}

func (p Params) Validate() error {
	return nil
}

func noopValidator() paramstypes.ValueValidatorFn {
	return func(value interface{}) error { return nil }
}

var _ orm.Validateable = GroupInfo{}

func (m GroupInfo) ValidateBasic() error {
	if m.GroupId.Empty() {
		return sdkerrors.Wrap(ErrEmpty, "group")
	}
	if m.Admin.Empty() {
		return sdkerrors.Wrap(ErrEmpty, "admin")
	}
	if err := sdk.VerifyAddressFormat(m.Admin); err != nil {
		return sdkerrors.Wrap(err, "admin")
	}
	if _, err := math.ParseNonNegativeDecimal(m.TotalWeight); err != nil {
		return sdkerrors.Wrap(err, "total weight")
	}
	if m.Version == 0 {
		return sdkerrors.Wrap(ErrEmpty, "version")
	}
	return nil
}

var _ orm.Validateable = GroupMember{}

func (g GroupMember) ValidateBasic() error {
	if g.GroupId.Empty() {
		return sdkerrors.Wrap(ErrEmpty, "group")
	}
	if g.Member.Empty() {
		return sdkerrors.Wrap(ErrEmpty, "address")
	}
	if _, err := math.ParsePositiveDecimal(g.Weight); err != nil {
		return sdkerrors.Wrap(err, "power")
	}
	if err := sdk.VerifyAddressFormat(g.Member); err != nil {
		return sdkerrors.Wrap(err, "address")
	}
	return nil
}

func (t *Tally) Sub(vote Vote, weight string) error {
	if err := t.operation(vote, weight, math.SafeSub); err != nil {
		return err
	}
	return nil
}

func (t *Tally) Add(vote Vote, weight string) error {
	if err := t.operation(vote, weight, math.Add); err != nil {
		return err
	}
	return nil
}

type operation func(res, x, y *apd.Decimal) error

func (t *Tally) operation(vote Vote, weight string, op operation) error {
	weightDec, err := math.ParsePositiveDecimal(weight)
	if err != nil {
		return err
	}

	yesCount, err := t.GetYesCount()
	if err != nil {
		return sdkerrors.Wrap(err, "yes count")
	}
	noCount, err := t.GetNoCount()
	if err != nil {
		return sdkerrors.Wrap(err, "no count")
	}
	abstainCount, err := t.GetAbstainCount()
	if err != nil {
		return sdkerrors.Wrap(err, "abstain count")
	}
	vetoCount, err := t.GetVetoCount()
	if err != nil {
		return sdkerrors.Wrap(err, "veto count")
	}

	switch vote.Choice {
	case Choice_CHOICE_YES:
		err := op(yesCount, yesCount, weightDec)
		if err != nil {
			return sdkerrors.Wrap(err, "yes count")
		}
		t.YesCount = math.DecimalString(yesCount)
	case Choice_CHOICE_NO:
		err := op(noCount, noCount, weightDec)
		if err != nil {
			return sdkerrors.Wrap(err, "no count")
		}
		t.NoCount = math.DecimalString(noCount)
	case Choice_CHOICE_ABSTAIN:
		err := op(abstainCount, abstainCount, weightDec)
		if err != nil {
			return sdkerrors.Wrap(err, "abstain count")
		}
		t.AbstainCount = math.DecimalString(abstainCount)
	case Choice_CHOICE_VETO:
		err := op(vetoCount, vetoCount, weightDec)
		if err != nil {
			return sdkerrors.Wrap(err, "veto count")
		}
		t.VetoCount = math.DecimalString(vetoCount)
	default:
		return sdkerrors.Wrapf(ErrInvalid, "unknown choice %s", vote.Choice.String())
	}
	return nil
}

// TotalCounts is the sum of all weights.
func (t Tally) TotalCounts() (*apd.Decimal, error) {
	yesCount, err := t.GetYesCount()
	if err != nil {
		return nil, sdkerrors.Wrap(err, "yes count")
	}
	noCount, err := t.GetNoCount()
	if err != nil {
		return nil, sdkerrors.Wrap(err, "no count")
	}
	abstainCount, err := t.GetAbstainCount()
	if err != nil {
		return nil, sdkerrors.Wrap(err, "abstain count")
	}
	vetoCount, err := t.GetVetoCount()
	if err != nil {
		return nil, sdkerrors.Wrap(err, "veto count")
	}

	totalCounts := apd.New(0, 0)
	err = math.Add(totalCounts, totalCounts, yesCount)
	if err != nil {
		return nil, err
	}
	err = math.Add(totalCounts, totalCounts, noCount)
	if err != nil {
		return nil, err
	}
	err = math.Add(totalCounts, totalCounts, abstainCount)
	if err != nil {
		return nil, err
	}
	err = math.Add(totalCounts, totalCounts, vetoCount)
	if err != nil {
		return nil, err
	}
	return totalCounts, nil
}

func (t Tally) GetYesCount() (*apd.Decimal, error) {
	yesCount, err := math.ParseNonNegativeDecimal(t.YesCount)
	if err != nil {
		return nil, err
	}
	return yesCount, nil
}

func (t Tally) GetNoCount() (*apd.Decimal, error) {
	noCount, err := math.ParseNonNegativeDecimal(t.NoCount)
	if err != nil {
		return nil, err
	}
	return noCount, nil
}

func (t Tally) GetAbstainCount() (*apd.Decimal, error) {
	abstainCount, err := math.ParseNonNegativeDecimal(t.AbstainCount)
	if err != nil {
		return nil, err
	}
	return abstainCount, nil
}

func (t Tally) GetVetoCount() (*apd.Decimal, error) {
	vetoCount, err := math.ParseNonNegativeDecimal(t.VetoCount)
	if err != nil {
		return nil, err
	}
	return vetoCount, nil
}

func (t Tally) ValidateBasic() error {
	if _, err := t.GetYesCount(); err != nil {
		return sdkerrors.Wrap(err, "yes count")
	}
	if _, err := t.GetNoCount(); err != nil {
		return sdkerrors.Wrap(err, "no count")
	}
	if _, err := t.GetAbstainCount(); err != nil {
		return sdkerrors.Wrap(err, "abstain count")
	}
	if _, err := t.GetVetoCount(); err != nil {
		return sdkerrors.Wrap(err, "veto count")
	}
	return nil
}

func (g GenesisState) String() string {
	out, _ := yaml.Marshal(g)
	return string(out)
}
