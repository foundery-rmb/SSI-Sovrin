package main

import "github.com/sasiedu/SSI-Sovrin/indy-sdk-wrapper/golang/indysdk"

func main() {
	config := indysdk.WalletConfig{
		ID: "alice-20",
	}
	credentials := indysdk.WalletCredential{
		Key: "test-alice",
	}

	indysdk.CreateWallet(config, credentials)
}
