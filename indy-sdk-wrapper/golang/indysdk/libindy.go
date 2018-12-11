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
	fmt.Println(">>> Create Wallet command")
	channel := wallet.IndyCreateWallet((wallet.Config)(config), (wallet.Credential)(credential))
	result := <-channel
	return result.Error
}

// OpenWallet opens an existing wallet
func OpenWallet(config WalletConfig, credential WalletCredential) (int, error) {
	fmt.Println(">>> Open Wallet command")
	channel := wallet.IndyOpenWallet((wallet.Config)(config), (wallet.Credential)(credential))
	result := <-channel
	return result.Results[0].(int), result.Error
}

// ExportWallet exports an opened wallet
func ExportWallet(walletHandle int, config WalletExportConfig) error {
	fmt.Println(">>> Export Wallet command")
	channel := wallet.IndyExportWallet(walletHandle, (wallet.ExportConfig)(config))
	result := <-channel
	return result.Error
}
