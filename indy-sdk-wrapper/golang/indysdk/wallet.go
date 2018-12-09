package indysdk

/*
#cgo CFLAGS: -I ../Includes
#cgo LDFLAGS: -lindy
#include <indy_types.h>
#include <indy_wallet.h>
*/
import "C"

import (
	"encoding/json"
	"fmt"
	"os"

	//log "github.com/sirupsen/logrus"
)

type WalletStorageConfig struct {
	Path string `json:"path"`
} 

type WalletConfig struct {
	ID            string `json:"id"`
	StorageType   string `json:"storage_type"`
	StorageConfig WalletStorageConfig `json:"storage_config"`
}

type WalletCredential struct {
	Key string `json:"key"`
	//	Rekey                 string `json:"rekey"`
	StorageCredentials string `json:"storage_credentials"`
	//KeyDerivationMethod string `json:"key_derivation_method"`
	//	ReKeyDerivationMethod string `json:"rekey_derivation_method"`
}

func setDefaults(config WalletConfig, credential WalletCredential) (WalletConfig, WalletCredential) {
	if config.StorageType == "" {
		config.StorageType = "default"
	}
	if config.StorageConfig.Path == "" {
		config.StorageConfig.Path = os.Getenv("HOME") + "/.indy_client/wallet"
	}
	return config, credential
}

func CreateWallet(config WalletConfig, credential WalletCredential) error {
	config, credential = setDefaults(config, credential)

	jsonConfig, err := json.Marshal(config)
	if err != nil {
		return err
	}
	jsonCredential, err := json.Marshal(credential)
	if err != nil {
		return err
	}

	res_future := callIndy(string(jsonConfig), string(jsonCredential))
	res := <-res_future
	fmt.Println("Res: ", res)

	return nil
}
