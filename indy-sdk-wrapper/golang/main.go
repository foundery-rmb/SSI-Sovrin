package main

import "github.com/sasiedu/SSI-Sovrin/indy-sdk-wrapper/golang/indysdk/wallet"

func main() {
	config := wallet.WalletConfig{
		ID: "alice-20",
	}
	credentials := wallet.WalletCredential{
		Key: "test-alice",
	}

	wallet.CreateWallet(config, credentials)
}
