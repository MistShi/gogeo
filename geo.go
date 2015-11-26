package gogeo

/*
#include <stdlib.h>
#include <unistd.h>
#include "geolocation.h"
#cgo CFLAGS: -g -O0 -DDEBUG -D_XOPEN_SOURCE -D_BSD_SOURCE
#cgo LDFLAGS: -lm
*/
import "C"
import (
	"errors"
	"sync"
	"unsafe"
)

var mutex sync.Mutex

//ips_t * open_ips(char *filename, uint32_t flags){
func Open_ips(filename string, flags uint32) (ips *C.struct_ips_t, err error) {
	var null *C.struct_ips_t

	ips = C.open_ips(C.CString(filename), C.uint(flags))
	if ips == null {
		return nil, errors.New("can not open " + filename)
	}
	return ips, nil
}

//void clean_ips(ips_t *ips){
func Clean_ips(ips unsafe.Pointer) {
	C.clean_ips(*C.struct_ips_t(ips))
}

/*
typedef struct ip_entry{
#ifdef DEBUG
	uint32_t min;
	uint32_t max;
#endif
    uintptr_t country;
    uintptr_t province;
    uintptr_t city;
    uintptr_t village;
    uintptr_t isp;
*/
//e = (ip_entry *)radix32tree_find(ips->tree, mbuf.u.sip);
func Find_ip(_ips unsafe.Pointer, ip string) (isp, province string, err error) {
	var null *C.struct_ip_entry
	ips := *C.struct_ips_t(_ips)
	p := (*C.struct_ip_entry)(C.radix32tree_find(ips._tree, C.CString(ip)))
	if p != null {
		isp := C.GoString(p._isp)
		province := C.GoString(p._province)
	} else {
		err = errors.New("not found")
	}
}
