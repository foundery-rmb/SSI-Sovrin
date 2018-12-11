package main

import (
	"fmt"

	"github.com/sasiedu/SSI-Sovrin/indy-sdk-wrapper/golang/indysdk"
)

func main() {
	config := indysdk.WalletConfig{
		ID: "alice-23",
	}
	credentials := indysdk.WalletCredential{
		Key: "test-alice",
	}

	err := indysdk.CreateWallet(config, credentials)
	if err != nil {
		fmt.Println("Error:", err)
	}

	walletHandle, err := indysdk.OpenWallet(config, credentials)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Wallet handle:", walletHandle)
	}

	exportConfig := indysdk.WalletExportConfig{}

	err = indysdk.ExportWallet(walletHandle, exportConfig)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
