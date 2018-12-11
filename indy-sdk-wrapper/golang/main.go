package main

import (
	"fmt"

	"github.com/sasiedu/SSI-Sovrin/indy-sdk-wrapper/golang/indysdk"
)

func main() {
	config := indysdk.WalletConfig{
		ID: "alice-20",
	}
	credentials := indysdk.WalletCredential{
		Key: "test-alice",
	}

	res := indysdk.CreateWallet(config, credentials)
	if res != nil {
		fmt.Println("Error:", res)
	}
}
