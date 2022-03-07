package core

import (
	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	"github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	ecocreditv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/v1"
	"gotest.tools/v3/assert"
	"testing"
)

func TestUpdateClass_UpdateAdmin(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	addrs := genAddrs(1)
	newAdmin := addrs[0]

	err := s.stateStore.ClassInfoStore().Insert(s.ctx, &ecocreditv1.ClassInfo{
		Name:       "C01",
		Admin:      s.addr,
		Metadata:   nil,
		CreditType: "C",
	})
	assert.NilError(t, err)

	_, err = s.k.UpdateClassAdmin(s.ctx, &v1.MsgUpdateClassAdmin{
		Admin:    s.addr.String(),
		ClassId:  "C01",
		NewAdmin: newAdmin.String(),
	})
	assert.NilError(t, err)

	cInfo, err := s.stateStore.ClassInfoStore().Get(s.ctx, 1)
	assert.NilError(t, err)
	assert.Check(t, newAdmin.Equals(types.AccAddress(cInfo.Admin)))
}

func TestUpdateClass_UpdateAdminErrs(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	addr := genAddrs(1)[0]
	err := s.stateStore.ClassInfoStore().Insert(s.ctx, &ecocreditv1.ClassInfo{
		Name:       "C01",
		Admin:      s.addr,
		Metadata:   nil,
		CreditType: "C",
	})
	assert.NilError(t, err)

	// try to update the admin with addr
	_, err = s.k.UpdateClassAdmin(s.ctx, &v1.MsgUpdateClassAdmin{
		Admin:    addr.String(),
		ClassId:  "C01",
		NewAdmin: addr.String(),
	})
	assert.ErrorContains(t, err, "unauthorized")

	// try to update a non-existent class
	_, err = s.k.UpdateClassAdmin(s.ctx, &v1.MsgUpdateClassAdmin{
		Admin:    addr.String(),
		ClassId:  "FOOBAR",
		NewAdmin: addr.String(),
	})
	assert.ErrorContains(t, err, "not found")

}

func TestUpdateClass_Issuers(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	addrs := genAddrs(3)    // addrs to initially populate the class issuer store
	newAddrs := genAddrs(3) // addrs to add in an update call

	err := s.stateStore.ClassInfoStore().Insert(s.ctx, &ecocreditv1.ClassInfo{
		Name:       "C01",
		Admin:      s.addr,
		Metadata:   nil,
		CreditType: "C",
	})
	assert.NilError(t, err)

	// insert some addrs
	for _, addr := range addrs {
		err := s.stateStore.ClassIssuerStore().Insert(s.ctx, &ecocreditv1.ClassIssuer{
			ClassId: 1,
			Issuer:  addr,
		})
		assert.NilError(t, err)
	}

	// update class with newAddrs
	addrStrs := make([]string, 3)
	for i, newAddr := range newAddrs {
		addrStrs[i] = newAddr.String()
	}
	_, err = s.k.UpdateClassIssuers(s.ctx, &v1.MsgUpdateClassIssuers{
		Admin:         s.addr.String(),
		ClassId:       "C01",
		AddIssuers:    addrStrs,
		RemoveIssuers: nil,
	})
	assert.NilError(t, err)

	// check that the new addrs were added
	count := 0
	it, err := s.stateStore.ClassIssuerStore().List(s.ctx, ecocreditv1.ClassIssuerClassIdIssuerIndexKey{}.WithClassId(1))
	assert.NilError(t, err)
	for it.Next() {
		count++
	}
	assert.Equal(t, len(addrs)+len(newAddrs), count, "expected to get %d address matches, got %d", len(addrs)+len(newAddrs), count)

	// remove the original addrs
	removeAddrs := addrs
	removeAddrStrs := make([]string, 3)
	for i, rmAddr := range removeAddrs {
		removeAddrStrs[i] = rmAddr.String()
	}
	_, err = s.k.UpdateClassIssuers(s.ctx, &v1.MsgUpdateClassIssuers{
		Admin:         s.addr.String(),
		ClassId:       "C01",
		AddIssuers:    nil,
		RemoveIssuers: removeAddrStrs,
	})
	assert.NilError(t, err)

	// check that the removed addrs no longer exist
	it, err = s.stateStore.ClassIssuerStore().List(s.ctx, ecocreditv1.ClassIssuerClassIdIssuerIndexKey{}.WithClassId(1))
	assert.NilError(t, err)
	for it.Next() {
		val, err := it.Value()
		assert.NilError(t, err)
		addr := types.AccAddress(val.Issuer)
		for _, rmAddr := range removeAddrs {
			assert.Check(t, !addr.Equals(rmAddr), "%s was supposed to be deleted", rmAddr.String())
		}
	}
}

func TestUpdateClass_IssuersErrs(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	addr := genAddrs(1)[0]
	classRowId, err := s.stateStore.ClassInfoStore().InsertReturningID(s.ctx, &ecocreditv1.ClassInfo{
		Name:       "C01",
		Admin:      s.addr,
		Metadata:   nil,
		CreditType: "C",
	})
	assert.NilError(t, err)
	err = s.stateStore.ClassIssuerStore().Insert(s.ctx, &ecocreditv1.ClassIssuer{
		ClassId: classRowId,
		Issuer:  s.addr,
	})
	assert.NilError(t, err)

	// try to update without being admin
	_, err = s.k.UpdateClassIssuers(s.ctx, &v1.MsgUpdateClassIssuers{
		Admin:         addr.String(),
		ClassId:       "C01",
		AddIssuers:    nil,
		RemoveIssuers: nil,
	})
	assert.ErrorContains(t, err, sdkerrors.ErrUnauthorized.Error())

	// try to update non-existent class
	_, err = s.k.UpdateClassIssuers(s.ctx, &v1.MsgUpdateClassIssuers{
		Admin:         s.addr.String(),
		ClassId:       "FOO",
		AddIssuers:    nil,
		RemoveIssuers: nil,
	})
	assert.ErrorContains(t, err, sdkerrors.ErrNotFound.Error())

	// try to add an issuer that already exists
	_, err = s.k.UpdateClassIssuers(s.ctx, &v1.MsgUpdateClassIssuers{
		Admin:         s.addr.String(),
		ClassId:       "C01",
		AddIssuers:    []string{s.addr.String()},
		RemoveIssuers: nil,
	})
	assert.ErrorContains(t, err, ormerrors.PrimaryKeyConstraintViolation.Error())
}

func TestUpdateClass_Metadata(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	err := s.stateStore.ClassInfoStore().Insert(s.ctx, &ecocreditv1.ClassInfo{
		Name:       "C01",
		Admin:      s.addr,
		Metadata:   []byte("foobar"),
		CreditType: "C",
	})
	assert.NilError(t, err)

	_, err = s.k.UpdateClassMetadata(s.ctx, &v1.MsgUpdateClassMetadata{
		Admin:    s.addr.String(),
		ClassId:  "C01",
		Metadata: []byte("barfoo"),
	})
	assert.NilError(t, err)

	class, err := s.stateStore.ClassInfoStore().Get(s.ctx, 1)
	assert.NilError(t, err)
	assert.DeepEqual(t, []byte("barfoo"), class.Metadata)
}

func TestUpdateClass_MetadataErrs(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	addr := genAddrs(1)[0]
	err := s.stateStore.ClassInfoStore().Insert(s.ctx, &ecocreditv1.ClassInfo{
		Name:       "C01",
		Admin:      s.addr,
		Metadata:   nil,
		CreditType: "C",
	})
	assert.NilError(t, err)

	// try to update non-existent class
	_, err = s.k.UpdateClassMetadata(s.ctx, &v1.MsgUpdateClassMetadata{
		Admin:    s.addr.String(),
		ClassId:  "FOO",
		Metadata: nil,
	})
	assert.ErrorContains(t, err, sdkerrors.ErrNotFound.Error())

	// try to update class you are not the admin of
	_, err = s.k.UpdateClassMetadata(s.ctx, &v1.MsgUpdateClassMetadata{
		Admin:    addr.String(),
		ClassId:  "C01",
		Metadata: []byte("FOO"),
	})
	assert.ErrorContains(t, err, sdkerrors.ErrUnauthorized.Error())

}

func genAddrs(x int) []types.AccAddress {
	addrs := make([]types.AccAddress, x)
	for i := 0; i < x; i++ {
		_, _, addr := testdata.KeyTestPubAddr()
		addrs[i] = addr
	}
	return addrs
}
