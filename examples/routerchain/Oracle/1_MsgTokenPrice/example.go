package main

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"time"

	chainclient "github.com/router-protocol/sdk-go/client/routerchain"
	"github.com/router-protocol/sdk-go/client/routerchain/common"

	sdk "github.com/cosmos/cosmos-sdk/types"
	oracleTypes "github.com/router-protocol/sdk-go/routerchain/oracle/types"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
)

func main() {
	network := common.LoadNetwork("devnet-alpha", "k8s")
	tmRPC, err := rpchttp.New(network.TmEndpoint, "/websocket")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Network", network)
	senderAddress, cosmosKeyring, err := chainclient.InitCosmosKeyring(
		os.Getenv("HOME")+"/.routerd",
		"routerd",
		"file",
		"gaurav",
		"12345678",
		"35F804E20869F9150029FC2A70AAB602B5E4606251420083B3D36200A9B10C52", // keyring will be used if pk not provided
		false,
	)

	if err != nil {
		panic(err)
	}

	// initialize grpc client
	clientCtx, err := chainclient.NewClientContext(network.ChainId, senderAddress.String(), cosmosKeyring)
	if err != nil {
		fmt.Println(err)
	}
	clientCtx = clientCtx.WithNodeURI(network.TmEndpoint).WithClient(tmRPC)

	ctx, _ := context.WithTimeout(context.Background(), time.Minute)

	routerChainClient, err := chainclient.NewChainClient(
		clientCtx,
		network.ChainGrpcEndpoint,
		// common.OptionTLSCert(network.ChainTlsCert),
		common.OptionGasPrices("100000000000000route"),
	)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	allChainConfig, err := routerChainClient.GetAllChainConfig(ctx)
	if err != nil {
		fmt.Println("Error fetching all meta info")
		panic(err)
	}

	fmt.Println("allChainConfig", allChainConfig)

	tokenPrices := make([]oracleTypes.TokenPrice, 0)
	for _, chainConfig := range allChainConfig.ChainConfig {
		tokenPrice := oracleTypes.TokenPrice{
			Symbol:     chainConfig.Symbol,
			TokenPrice: sdk.NewDecFromBigInt(big.NewInt(100000)),
			Decimals:   18,
		}
		tokenPrices = append(tokenPrices, tokenPrice)
	}

	fmt.Println("tokenPrices", tokenPrices)
	// prepare tx msg
	msg := oracleTypes.NewMsgTokenPrices(routerChainClient.FromAddress().String(), tokenPrices)

	//AsyncBroadcastMsg, SyncBroadcastMsg, QueueBroadcastMsg
	err = routerChainClient.QueueBroadcastMsg(msg)
	if err != nil {
		fmt.Println(err)
	}

	time.Sleep(time.Second * 5)
	gasFee, err := routerChainClient.GetGasFee()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("gas fee:", gasFee)

}
