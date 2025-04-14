package toncenterzp

import (
	"fmt"
)

// Version is the current version of the library
const Version = "1.0.0"

// PrintVersion prints the current version of the library
func PrintVersion() {
	fmt.Printf("toncenterzp version %s\n", Version)
}

// Constants for API endpoints
const (
	EndpointDetectAddress              = "/detectAddress"
	EndpointEstimateFee                = "/estimateFee"
	EndpointGetAddressBalance          = "/getAddressBalance"
	EndpointGetAddressInformation      = "/getAddressInformation"
	EndpointGetAddressState            = "/getAddressState"
	EndpointGetBlockHeader             = "/getBlockHeader"
	EndpointGetBlockTransactions       = "/getBlockTransactions"
	EndpointGetConsensusBlock          = "/getConsensusBlock"
	EndpointGetExtendedAddressInformation = "/getExtendedAddressInformation"
	EndpointGetMasterchainBlockSignatures = "/getMasterchainBlockSignatures"
	EndpointGetMasterchainInfo         = "/getMasterchainInfo"
	EndpointGetShardBlockProof         = "/getShardBlockProof"
	EndpointGetTokenData               = "/getTokenData"
	EndpointGetTransactions            = "/getTransactions"
	EndpointGetWalletInformation       = "/getWalletInformation"
	EndpointJSONRPC                    = "/jsonRPC"
	EndpointLookupBlock                = "/lookupBlock"
	EndpointPackAddress                = "/packAddress"
	EndpointRunGetMethod               = "/runGetMethod"
	EndpointSendBoc                    = "/sendBoc"
	EndpointSendBocReturnHash          = "/sendBocReturnHash"
	EndpointSendQuery                  = "/sendQuery"
	EndpointShards                     = "/shards"
	EndpointTryLocateResultTx          = "/tryLocateResultTx"
	EndpointTryLocateSourceTx          = "/tryLocateSourceTx"
	EndpointTryLocateTx                = "/tryLocateTx"
	EndpointUnpackAddress              = "/unpackAddress"
)
