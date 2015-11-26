package gogeo

/*
#include <stdlib.h>
#include <unistd.h>
#include "geolocation.h"
#cgo CFLAGS: -g -O0 -DDEBUG -D_XOPEN_SOURCE -D_BSD_SOURCE
#cgo LDFLAGS: -lm
*/
import "C"
import "errors"

const (
	GEO_F_ALIAS = 1
)

func Open_ips(filename string, flags uint32) (ips *C.struct_ips_t, err error) {
	var null *C.struct_ips_t

	ips = C.open_ips(C.CString(filename), C.uint32_t(flags))
	if ips == null {
		return nil, errors.New("can not open " + filename)
	}
	return ips, nil
}

func Clean_ips(ips *C.struct_ips_t) {
	C.clean_ips(ips)
}

func Find_ip(ips *C.struct_ips_t, ip string) (isp, province string, err error) {
	var null *C.struct_out_entry
	p := C.find_ip(ips, C.CString(ip))
	if p != null {
		isp = C.GoString(p.isp)
		if len(isp) > 2 && isp[:2] == "i_" {
			isp = isp[2:]
		}
		province = C.GoString(p.province)
		if len(province) > 2 && province[:2] == "p_" {
			province = province[2:]
		}
	} else {
		err = errors.New("not found")
	}
	return
}
