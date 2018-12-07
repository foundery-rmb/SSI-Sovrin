package indysdk

/*
#cgo CFLAGS: -I ../Includes
#cgo LDFLAGS: -lindy
#include <indy_types.h>
#include <indy_mod.h>
*/
import "C"
import (
	"unsafe"
)

func indyCallback(commandHandle C.indy_handle_t, indyError C.indy_error_t) {
	logger.Debug("This is the indy callback")
}

func createIndyCallback() *[0]byte {
	cbFunc := indyCallback
	cb := (*[0]byte)(unsafe.Pointer(&cbFunc))
	return cb
}
