package wallet

/*
#cgo CFLAGS: -I ../../Includes
#cgo LDFLAGS: -lindy
#include <indy_types.h>
#include <indy_wallet.h>
#include <my_header.h>

#define LAMBDA(c_) ({ c_ _;})

extern void walletCallback(indy_handle_t, indy_error_t);

static void (*get_void_callback(void)) (indy_handle_t command_handle, indy_error_t err)
{
	return LAMBDA(void _(indy_handle_t h, indy_error_t e){
		walletCallback(h, e);
	});
}
*/
import "C"

import (
	"encoding/json"
	"errors"
	"os"

	_ "github.com/sasiedu/SSI-Sovrin/indy-sdk-wrapper/golang/indysdk"
	//log "github.com/sirupsen/logrus"
)

func setDefaults(config WalletConfig, credential WalletCredential) (WalletConfig, WalletCredential) {
	if config.StorageType == "" {
		config.StorageType = "default"
	}
	if config.StorageConfig.Path == "" {
		config.StorageConfig.Path = os.Getenv("HOME") + "/.indy_client/wallet"
	}
	return config, credential
}

// export walletCallback
func walletCallback(commandHandle C.indy_handle_t, indyError C.indy_error_t) {

}

func createWallet(config WalletConfig, credential WalletCredential) chan IndyResult {
	config, credential = setDefaults(config, credential)
	commandHandle, future := newFutureCommand()

	jsonConfig, err := json.Marshal(config)
	if err != nil {
		return removeFuture(commandHandle, IndyResult{Error: err})
	}
	jsonCredential, err := json.Marshal(credential)
	if err != nil {
		return removeFuture(commandHandle, IndyResult{Error: err})
	}
	res := C.indy_create_wallet(commandHandle, C.CString(config), C.CString(credential), C.get_void_callback())
	if res != 0 {
		return removeFuture(commandHandle, IndyResult{Error: errors.New(string(res))})
	}

	return future
}
