package data

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
)

func TestMsgAnchorDataRequest_ValidateBasic(t *testing.T) {
	_, _, addr := testdata.KeyTestPubAddr()
	type fields struct {
		Sender string
		Hash   *ContentHash
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr string
	}{
		{
			name: "good",
			fields: fields{
				Sender: addr.String(),
				Hash: &ContentHash{Sum: &ContentHash_Raw_{Raw: &ContentHash_Raw{
					Hash:            make([]byte, 32),
					DigestAlgorithm: DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
					MediaType:       MediaType_MEDIA_TYPE_UNSPECIFIED,
				},
				}},
			},
			wantErr: "",
		},
		{
			name: "nil",
			fields: fields{
				Sender: addr.String(),
				Hash:   nil,
			},
			wantErr: "hash cannot be empty: invalid request",
		},
		{
			name: "bad",
			fields: fields{
				Sender: addr.String(),
				Hash: &ContentHash{Sum: &ContentHash_Raw_{Raw: &ContentHash_Raw{
					Hash:            make([]byte, 31),
					DigestAlgorithm: DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
					MediaType:       MediaType_MEDIA_TYPE_UNSPECIFIED,
				},
				}},
			},
			wantErr: "expected 32 bytes for DIGEST_ALGORITHM_BLAKE2B_256, got 31: invalid request",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MsgAnchorData{
				Sender: tt.fields.Sender,
				Hash:   tt.fields.Hash,
			}
			err := m.ValidateBasic()
			if len(tt.wantErr) != 0 {
				require.EqualError(t, err, tt.wantErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMsgSignDataRequest_ValidateBasic(t *testing.T) {
	_, _, addr := testdata.KeyTestPubAddr()
	type fields struct {
		Signers []string
		Hash    *ContentHash_Graph
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr string
	}{
		{
			"good",
			fields{
				Signers: []string{addr.String()},
				Hash: &ContentHash_Graph{
					Hash:                      make([]byte, 32),
					DigestAlgorithm:           DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
					CanonicalizationAlgorithm: GraphCanonicalizationAlgorithm_GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015,
					MerkleTree:                GraphMerkleTree_GRAPH_MERKLE_TREE_NONE_UNSPECIFIED,
				},
			},
			"",
		},
		{
			"bad",
			fields{
				Signers: nil,
				Hash: &ContentHash_Graph{
					Hash:                      make([]byte, 32),
					DigestAlgorithm:           DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
					CanonicalizationAlgorithm: GraphCanonicalizationAlgorithm_GRAPH_CANONICALIZATION_ALGORITHM_UNSPECIFIED,
					MerkleTree:                GraphMerkleTree_GRAPH_MERKLE_TREE_NONE_UNSPECIFIED,
				},
			},
			"invalid data.GraphCanonicalizationAlgorithm GRAPH_CANONICALIZATION_ALGORITHM_UNSPECIFIED: invalid request",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MsgSignData{
				Signers: tt.fields.Signers,
				Hash:    tt.fields.Hash,
			}
			err := m.ValidateBasic()
			if len(tt.wantErr) != 0 {
				require.EqualError(t, err, tt.wantErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
