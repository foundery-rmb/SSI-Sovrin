package utils

/*
#cgo CFLAGS: -I ../../Includes
#cgo LDFLAGS: -lindy
#include <wrapper.h>
*/
import "C"
import (
	"time"
)

// IndyResult represents callback result from C-call to libindy
type IndyResult struct {
	Error   error
	Results []interface{}
}

var futures = make(map[C.indy_handle_t](chan IndyResult))

// NewFutureCommand creates a new future command
func NewFutureCommand() (C.indy_handle_t, chan IndyResult) {
	handle := int32(time.Now().Unix())
	commandHandle := (C.indy_handle_t)(handle)
	future := make(chan IndyResult)
	futures[commandHandle] = future
	return commandHandle, future
}

// RemoveFuture removes a future from the futures map
func RemoveFuture(handle int, result IndyResult) chan IndyResult {
	future := futures[(C.indy_handle_t)(handle)]
	future <- result
	delete(futures, (C.indy_handle_t)(handle))
	return future
}
