package indy

// #cgo LDFLAGS: -lindy
// #include "include/indy_wallet.h"
import (
	"C"
	"encoding/json"
)

type WalletConfig struct {
	ID            string `json:"id"`
	StorageType   string `json:"storage_type"`
	StorageConfig string `json:"storage_config"`
}

type WalletCredential struct {
	Key                   string `json:"key"`
	Rekey                 string `json:"rekey"`
	StorageCredentials    string `json:"storage_credentials"`
	KeyDerivationMethod   string `json:"key_derivation_method"`
	ReKeyDerivationMethod string `json:"rekey_derivation_method"`
}

func CreateWallet(config WalletConfig, credential WalletCredential) error {
	jsonConfig, err := json.Marshal(config)
	if err != nil {
		return err
	}
	jsonCredential, err := json.Marshal(credential)
	if err != nil {
		return err
	}

	commandHandle := C.int(1)
	indyError := C.indy_create_wallet(commandHandle, C.CString(string(jsonConfig)), C.CString(string(jsonCredential)))
	//fmt.Println(indyError)

	return nil
}
