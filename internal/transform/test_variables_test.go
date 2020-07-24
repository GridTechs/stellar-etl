package transform

import (
	ingestio "github.com/stellar/go/exp/ingest/io"
	"github.com/stellar/go/xdr"
	"github.com/stellar/stellar-etl/internal/utils"
)

var genericSourceAccount, _ = xdr.NewMuxedAccount(xdr.CryptoKeyTypeKeyTypeEd25519, xdr.Uint256([32]byte{}))
var genericAccountID, _ = xdr.NewAccountId(xdr.PublicKeyTypePublicKeyTypeEd25519, xdr.Uint256([32]byte{}))
var genericAccountAddress, _ = genericAccountID.GetAddress()
var genericOperation = xdr.Operation{
	SourceAccount: &genericSourceAccount,
	Body: xdr.OperationBody{
		Type:           xdr.OperationTypeBumpSequence,
		BumpSequenceOp: &xdr.BumpSequenceOp{},
	},
}
var genericEnvelope = xdr.TransactionV1Envelope{
	Tx: xdr.Transaction{
		SourceAccount: genericSourceAccount,
		Memo:          xdr.Memo{},
		Operations: []xdr.Operation{
			genericOperation,
		},
	},
}
var genericLedgerTransaction = ingestio.LedgerTransaction{
	Index: 1,
	Envelope: xdr.TransactionEnvelope{
		Type: xdr.EnvelopeTypeEnvelopeTypeTx,
		V1:   &genericEnvelope,
	},
	Result: utils.CreateSampleResultMeta(true, 10).Result,
}
var genericLedgerHeaderHistoryEntry = xdr.LedgerHeaderHistoryEntry{}

// a selection of hardcoded accounts with their IDs and addresses

var testAccount1, _ = xdr.NewMuxedAccount(xdr.CryptoKeyTypeKeyTypeEd25519, xdr.Uint256([32]byte{0x88, 0xe1, 0xa6, 0xb4, 0xa5, 0x71, 0x52, 0x8f, 0xaa, 0x9f, 0x26, 0xe2, 0xe, 0x6c, 0xbc, 0x3f, 0xe7, 0x3c, 0xae, 0x80, 0xb4, 0x46, 0x63, 0xc, 0x5b, 0xef, 0x1, 0x52, 0xaf, 0x70, 0x7d, 0x78}))
var testAccount1ID = testAccount1.ToAccountId()
var testAccount1Address = "GCEODJVUUVYVFD5KT4TOEDTMXQ76OPFOQC2EMYYMLPXQCUVPOB6XRWPQ"

var testAccount2, _ = xdr.NewMuxedAccount(xdr.CryptoKeyTypeKeyTypeEd25519, xdr.Uint256([32]byte{0x1c, 0x47, 0x41, 0x97, 0x18, 0xee, 0xfa, 0xa4, 0x5b, 0x38, 0xcb, 0x7f, 0x2f, 0x25, 0x50, 0x1a, 0x9e, 0x39, 0xcb, 0x83, 0x87, 0xa6, 0x36, 0xe9, 0xfb, 0xcc, 0xc, 0x74, 0xa4, 0x77, 0x3, 0x18}))
var testAccount2ID = testAccount2.ToAccountId()
var testAccount2Address = "GAOEOQMXDDXPVJC3HDFX6LZFKANJ4OOLQOD2MNXJ7PGAY5FEO4BRRAQU"

var testAccount3, _ = xdr.NewMuxedAccount(xdr.CryptoKeyTypeKeyTypeEd25519, xdr.Uint256([32]byte{0x67, 0xcc, 0x0, 0x86, 0x4c, 0x3b, 0x89, 0x16, 0x8c, 0x6a, 0xaf, 0xe0, 0xbc, 0x34, 0x70, 0x9e, 0xd0, 0x21, 0xc5, 0x5, 0x72, 0xe2, 0xf9, 0x88, 0x61, 0x34, 0x22, 0x8, 0x2c, 0x22, 0x29, 0x72}))
var testAccount3ID = testAccount3.ToAccountId()
var testAccount3Address = "GBT4YAEGJQ5YSFUMNKX6BPBUOCPNAIOFAVZOF6MIME2CECBMEIUXFZZN"

var testAccount4, _ = xdr.NewMuxedAccount(xdr.CryptoKeyTypeKeyTypeEd25519, xdr.Uint256([32]byte{0x6b, 0x58, 0xdd, 0x6c, 0x68, 0x93, 0xb, 0x6d, 0x3, 0x1d, 0xc5, 0xbb, 0xfa, 0xe2, 0x3e, 0xfa, 0x1e, 0xc3, 0xf0, 0xbb, 0x58, 0xc2, 0xbc, 0x8d, 0x93, 0x8d, 0x47, 0xc1, 0xdf, 0xb3, 0xbe, 0x73}))
var testAccount4ID = testAccount4.ToAccountId()
var testAccount4Address = "GBVVRXLMNCJQW3IDDXC3X6XCH35B5Q7QXNMMFPENSOGUPQO7WO7HGZPA"

// a selection of hardcoded assets and their AssetOutput representations

var usdtAsset = xdr.Asset{
	Type: xdr.AssetTypeAssetTypeCreditAlphanum4,
	AlphaNum4: &xdr.AssetAlphaNum4{
		AssetCode: xdr.AssetCode4([4]byte{0x55, 0x53, 0x44, 0x54}),
		Issuer:    testAccount4ID,
	},
}
var usdtAssetOutput = AssetOutput{
	AssetType:   "credit_alphanum4",
	AssetCode:   "USDT",
	AssetIssuer: testAccount4Address,
}

var ethAsset = xdr.Asset{
	Type: xdr.AssetTypeAssetTypeCreditAlphanum4,
	AlphaNum4: &xdr.AssetAlphaNum4{
		AssetCode: xdr.AssetCode4([4]byte{0x45, 0x54, 0x48}),
		Issuer:    testAccount3ID,
	},
}
var ethAssetOutput = AssetOutput{
	AssetType:   "credit_alphanum4",
	AssetCode:   "ETH",
	AssetIssuer: testAccount1Address,
}

var nativeAsset = xdr.MustNewNativeAsset()
var nativeAssetOutput = AssetOutput{
	AssetType: "native",
}
