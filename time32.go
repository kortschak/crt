// Copyright 2017 The CRT Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build 386 arm arm64be armbe mips mipsle ppc ppc64le s390 s390x sparc

// +build !windows

package crt

const Ttm = "struct{int32,int32,int32,int32,int32,int32,int32,int32,int32,int32,*int8}"

type Xtm struct {
	X0  int32
	X1  int32
	X2  int32
	X3  int32
	X4  int32
	X5  int32
	X6  int32
	X7  int32
	X8  int32
	X9  int32
	X10 *int8
}

// struct tm *localtime(const time_t *timep);
func Xlocaltime(tls *TLS, timep *int32) *Xtm {
	TODO("")
	panic("TODO")
}

// time_t time(time_t *tloc);
func Xtime(tls *TLS, tloc *int32) int32 {
	panic("TODO")
}