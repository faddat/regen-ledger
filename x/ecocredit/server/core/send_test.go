package core

import (
	"testing"

	"github.com/golang/mock/gomock"
	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"

	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

func TestSend_Valid(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	_, _, recipient := testdata.KeyTestPubAddr()
	s.setupClassProjectBatch(t)
	any := gomock.Any()
	s.paramsKeeper.EXPECT().GetParamSet(any, any).Do(func(_ interface{}, p *ecocredit.Params) {
		p.CreditTypes = []*ecocredit.CreditType{{Name: "carbon", Abbreviation: "C", Unit: "tonne", Precision: 6}}
	}).Times(2) // this will be called for each batchDenom we send

	// s.Addr starting balance -> 10.5 tradable, 10.5 retired

	_, err := s.k.Send(s.ctx, &core.MsgSend{
		Sender:    s.addr.String(),
		Recipient: recipient.String(),
		Credits: []*core.MsgSend_SendCredits{
			{BatchDenom: "C01-20200101-20210101-01", TradableAmount: "2.51"},
			{BatchDenom: "C01-20200101-20210101-01", RetiredAmount: "1.30", RetirementLocation: "US-OR"},
		},
	})
	assert.NilError(t, err)

	senderBal, err := s.stateStore.BatchBalanceTable().Get(s.ctx, s.addr, 1)
	assert.NilError(t, err)

	recipientBal, err := s.stateStore.BatchBalanceTable().Get(s.ctx, recipient, 1)
	assert.NilError(t, err)

	// sender tradable -> 10.5 - 2.51 = 7.99
	// recipient now has 2.51
	// sender tradable -> 7.99 retires 1.30 = 6.69
	// recipient now has 1.30 retired

	assert.Equal(t, "6.69", senderBal.Tradable)
	assert.Equal(t, "2.51", recipientBal.Tradable)
	assert.Equal(t, "1.30", recipientBal.Retired)
	assert.Equal(t, "10.5", senderBal.Retired) // retired credits should be untouched

	sup, err := s.stateStore.BatchSupplyTable().Get(s.ctx, 1)
	assert.NilError(t, err)

	// retired -> 10.5 + 1.30(retired) = 11.80
	// tradable -> 10.5 - 1.30 = 9.20

	assert.Equal(t, "9.20", sup.TradableAmount)
	assert.Equal(t, "11.80", sup.RetiredAmount)
}

func TestSend_Errors(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	_, _, recipient := testdata.KeyTestPubAddr()
	s.setupClassProjectBatch(t)
	any := gomock.Any()
	s.paramsKeeper.EXPECT().GetParamSet(any, any).Do(func(_ interface{}, p *ecocredit.Params) {
		p.CreditTypes = []*ecocredit.CreditType{{Name: "carbon", Abbreviation: "C", Unit: "tonne", Precision: 6}}
	}).Times(2)

	// test sending more than user balance
	_, err := s.k.Send(s.ctx, &core.MsgSend{
		Sender:    s.addr.String(),
		Recipient: recipient.String(),
		Credits: []*core.MsgSend_SendCredits{
			{BatchDenom: "C01-20200101-20210101-01", TradableAmount: "1000000"},
		},
	})
	assert.ErrorContains(t, err, "insufficient funds")

	// test sending more precise than the credit type
	_, err = s.k.Send(s.ctx, &core.MsgSend{
		Sender:    s.addr.String(),
		Recipient: recipient.String(),
		Credits: []*core.MsgSend_SendCredits{
			{BatchDenom: "C01-20200101-20210101-01", TradableAmount: "10.325092385"},
		},
	})
	assert.ErrorContains(t, err, "exceeds maximum decimal places")
}