//go:build itest
// +build itest

package itest

var testCases = []*testCase{
	{
		name: "mint assets",
		test: testMintAssets,
	},
	{
		name: "asset name collision raises mint error",
		test: testMintAssetNameCollisionError,
	},
	{
		name: "addresses",
		test: testAddresses,
	},
	{
		name: "multi address",
		test: testMultiAddress,
	},
	{
		name:           "basic send",
		test:           testBasicSend,
		enableHashMail: true,
	},
	{
		name:           "reattempt failed asset send",
		test:           testReattemptFailedAssetSend,
		enableHashMail: true,
	},
	{
		name:           "offline receiver eventually receives",
		test:           testOfflineReceiverEventuallyReceives,
		enableHashMail: true,
	},
	{
		name: "basic send passive asset",
		test: testBasicSendPassiveAsset,
	},
	{
		name: "multi input send non-interactive single ID",
		test: testMultiInputSendNonInteractiveSingleID,
	},
	{
		name: "round trip send",
		test: testRoundTripSend,
	},
	{
		name: "full value send",
		test: testFullValueSend,
	},
	{
		name: "collectible send",
		test: testCollectibleSend,
	},
	{
		name: "re-issuance",
		test: testReIssuance,
	},
	{
		name: "minting multi asset groups",
		test: testMintMultiAssetGroups,
	},
	{
		name: "re-issuance amount overflow",
		test: testReIssuanceAmountOverflow,
	},
	{
		name: "minting multi asset groups errors",
		test: testMintMultiAssetGroupErrors,
	},
	{
		name: "mint with group key errors",
		test: testMintWithGroupKeyErrors,
	},
	{
		name: "psbt script hash lock send",
		test: testPsbtScriptHashLockSend,
	},
	{
		name: "psbt script check sig send",
		test: testPsbtScriptCheckSigSend,
	},
	{
		name: "psbt normal interactive full value send",
		test: testPsbtNormalInteractiveFullValueSend,
	},
	{
		name: "psbt grouped interactive full value send",
		test: testPsbtGroupedInteractiveFullValueSend,
	},
	{
		name: "psbt normal interactive split send",
		test: testPsbtNormalInteractiveSplitSend,
	},
	{
		name: "psbt grouped interactive split send",
		test: testPsbtGroupedInteractiveSplitSend,
	},
	{
		name: "psbt interactive tapscript sibling",
		test: testPsbtInteractiveTapscriptSibling,
	},
	{
		name: "psbt multi send",
		test: testPsbtMultiSend,
	},
	{
		name: "universe REST API",
		test: testUniverseREST,
	},
	{
		name: "universe sync",
		test: testUniverseSync,
	},
	{
		name: "universe federation",
		test: testUniverseFederation,
	},
	{
		name: "get info",
		test: testGetInfo,
	},
}
