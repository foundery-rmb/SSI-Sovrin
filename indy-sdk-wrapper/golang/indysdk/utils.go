package indysdk

/*
#cgo CFLAGS: -I ../Includes
#cgo LDFLAGS: -lindy
#include <indy_types.h>
#include <indy_mod.h>
#include <indy_wallet.h>

extern void goIndyCallback(indy_handle_t, indy_error_t);

static void indy_callback(indy_handle_t command_handle, indy_error_t err)
{
	goIndyCallback(command_handle, err);
}

static void (*get_indy_callback(void)) (indy_handle_t command_handle, indy_error_t err)
{
	return indy_callback;
}
*/
import "C"
import (
	"fmt"

	//log "github.com/sirupsen/logrus"
)

var future = make (chan C.indy_error_t)

//export goIndyCallback
func goIndyCallback(commandHandle C.indy_handle_t, indyError C.indy_error_t) {
	//log.Debug("This is the indy callback")
	future <- indyError
}

func callIndy(config string, credential string) (chan C.indy_error_t) {
	commandHandle := (C.indy_handle_t)(C.int(1))

	indyError := C.indy_create_wallet(commandHandle, C.CString(config), C.CString(credential), C.get_indy_callback())
	fmt.Println("First Indy Error: ", indyError)
	return future
}