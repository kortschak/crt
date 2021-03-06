// Copyright 2017 The CRT Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crt

import (
	"fmt"
	"syscall"
)

// extern int stat(char *__file, struct stat *__buf);
func Xstat(tls TLS, file, buf uintptr) int32 { return Xstat64(tls, file, buf) }

// extern int stat64(char *__file, struct stat64 *__buf);
func Xstat64(tls TLS, file, buf uintptr) int32 {
	r, _, err := syscall.Syscall(syscall.SYS_STAT, file, buf, 0)
	if strace {
		fmt.Fprintf(TraceWriter, "stat(%q, %#x) %v %v\n", GoString(file), buf, r, err)
	}
	if err != 0 {
		tls.setErrno(err)
	}
	return int32(r)
}

// int fstat(int fd, struct stat *buf);
func Xfstat(tls TLS, fd int32, buf uintptr) int32 { return Xfstat64(tls, fd, buf) }

// int fstat64(int fildes, struct stat64 *buf);
func Xfstat64(tls TLS, fildes int32, buf uintptr) int32 {
	r, _, err := syscall.Syscall(syscall.SYS_FSTAT, uintptr(fildes), buf, 0)
	if strace {
		fmt.Fprintf(TraceWriter, "fstat(%v, %#x) %v %v\n", fildes, buf, r, err)
	}
	if err != 0 {
		tls.setErrno(err)
	}
	return int32(r)
}

// extern int lstat(char *__file, struct stat *__buf);
func Xlstat(tls TLS, file, buf uintptr) int32 { return Xlstat64(tls, file, buf) }

// extern int lstat64(char *__file, struct stat64 *__buf);
func Xlstat64(tls TLS, file, buf uintptr) int32 {
	r, _, err := syscall.Syscall(syscall.SYS_LSTAT, file, buf, 0)
	if strace {
		fmt.Fprintf(TraceWriter, "lstat(%q, %#x) %v %v\n", GoString(file), buf, r, err)
	}
	if err != 0 {
		tls.setErrno(err)
	}
	return int32(r)
}
