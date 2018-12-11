package indysdk

/*
#cgo CFLAGS: -I ../Includes
#cgo LDFLAGS: -lindy
#include <wrapper.h>
*/
import "C"
import (
	"fmt"

	"github.com/sasiedu/SSI-Sovrin/indy-sdk-wrapper/golang/indysdk/wallet"
)

// CreateWallet creates a new secure wallet with the given unique name
func CreateWallet(config WalletConfig, credential WalletCredential) error {
	fmt.Println("Create Wallet command")
	channel := wallet.IndyCreateWallet((wallet.Config)(config), (wallet.Credential)(credential))
	result := <-channel
	return result.Error
}
