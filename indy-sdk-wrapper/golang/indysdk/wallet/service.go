package wallet

/*
#cgo CFLAGS: -I ../../Includes
#cgo LDFLAGS: -lindy
#include <wrapper.h>
*/
import "C"

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/sasiedu/SSI-Sovrin/indy-sdk-wrapper/golang/indysdk/utils"
	//log "github.com/sirupsen/logrus"
)

//export defaultCallback
func defaultCallback(commandHandle C.indy_handle_t, indyError C.indy_error_t) {
	if indyError == 0 {
		utils.RemoveFuture((int)(commandHandle), utils.IndyResult{Error: nil})
	} else {
		utils.RemoveFuture((int)(commandHandle), utils.IndyResult{Error: fmt.Errorf("%v", int(indyError))})
	}
}

func setDefaults(config Config, credential Credential) (Config, Credential) {
	if config.StorageType == "" {
		config.StorageType = "default"
	}
	if config.StorageConfig.Path == "" {
		config.StorageConfig.Path = os.Getenv("HOME") + "/.indy_client/wallet"
	}
	if credential.KeyDerivationMethod == "" {
		credential.KeyDerivationMethod = "ARGON2I_MOD"
	}
	if credential.ReKeyDerivationMethod == "" {
		credential.ReKeyDerivationMethod = "ARGON2I_MOD"
	}
	return config, credential
}

// IndyCreateWallet creates a new secure wallet with the given unique name
func IndyCreateWallet(config Config, credential Credential) chan utils.IndyResult {
	config, credential = setDefaults(config, credential)
	handle, future := utils.NewFutureCommand()

	jsonConfig, err := json.Marshal(config)
	if err != nil {
		go func() { utils.RemoveFuture((int)(handle), utils.IndyResult{Error: err}) }()
		return future
	}
	jsonCredential, err := json.Marshal(credential)
	if err != nil {
		go func() { utils.RemoveFuture((int)(handle), utils.IndyResult{Error: err}) }()
		return future
	}

	configString := string(jsonConfig)
	credentialString := string(jsonCredential)
	commandHandle := (C.indy_handle_t)(handle)
	fmt.Println("Creating wallet with commandHandle", commandHandle, "with config", configString, "and credential", credentialString)
	res := C.indy_create_wallet(commandHandle, C.CString(configString), C.CString(credentialString), C.get_default_callback())
	if res != 0 {
		go func() { utils.RemoveFuture((int)(handle), utils.IndyResult{Error: fmt.Errorf("%v", int(res))}) }()
		return future
	}

	return future
}
