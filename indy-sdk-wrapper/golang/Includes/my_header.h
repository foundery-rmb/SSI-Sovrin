#include "indy_core.h"

//#define LAMBDA(c_) ({ c_ _;})

/*static void (*get_callback(int i)) (indy_handle_t command_handle, indy_error_t err)
{
	if (i == 0){
		return LAMDA(void _(indy_handle_t h, indy_error_t e)){
			walletCallback(h, e, null);
		};
	}
	return LAMDA(void _(indy_handle_t h, indy_error_t e, const char * const val)){
		walletCallback(h, e, val);
	};
}*/