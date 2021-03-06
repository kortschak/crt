// Copyright 2017 The CRT Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crt

import (
	"fmt"
	"syscall"
)

// int open(const char *pathname, int flags, ...);
func Xopen(tls TLS, pathname uintptr, flags int32, args ...interface{}) int32 {
	return Xopen64(tls, pathname, flags, args...)
}

// int open(const char *pathname, int flags, ...);
func Xopen64(tls TLS, pathname uintptr, flags int32, args ...interface{}) int32 {
	var mode uintptr
	switch len(args) {
	case 0:
		// nop
	case 1:
		switch x := args[0].(type) {
		case int32:
			mode = uintptr(x)
		case uint32:
			mode = uintptr(x)
		default:
			panic(fmt.Errorf("crt.Xopen64 %T", x))
		}
	default:
		panic("TODO")
	}
	r, _, err := syscall.Syscall(syscall.SYS_OPEN, pathname, uintptr(flags), mode)
	if strace {
		fmt.Fprintf(TraceWriter, "open(%q, %v, %#o) %v %v\n", GoString(pathname), modeString(flags), mode, r, err)
	}
	if err != 0 {
		tls.setErrno(err)
	}
	return int32(r)
}
