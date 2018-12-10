package indysdk

/*
#cgo CFLAGS: -I ../Includes
#cgo LDFLAGS: -lindy
#include <indy_types.h>
*/
import "C"
import "time"

// IndyResult represents callback result from C-call to libindy
type IndyResult struct {
	Error   error
	Results []interface{}
}

var futures = make(map[C.indy_handle_t](chan IndyResult))

func newFutureCommand() (C.indy_handle_t, chan IndyResult) {
	handle := int32(time.Now().Unix())
	commandHandle := (C.indy_handle_t)(handle)
	future := make(chan IndyResult)
	futures[commandHandle] = future
	return commandHandle, future
}

func removeFuture(handle C.indy_handle_t, result IndyResult) chan IndyResult {
	future := futures[handle]
	future <- result
	delete(futures, handle)
	return future
}