package main

/*
#cgo CFLAGS: -I${SRCDIR}/include/bridge
#include "bridge.h"
*/
import "C"
import "unsafe"

type Vehicle struct{}

func (Vehicle) Create(model int, x, y, z, rotation float32, color1, color2 uint32, respawnDelay uint32, addsiren bool) int {
	return int(C.Vehicle_Create(C.int(model), C.float(x), C.float(y), C.float(z), C.float(rotation), C.uint32_t(color1), C.uint32_t(color2), C.uint32_t(respawnDelay), C.bool(addsiren)))
}

func (v *Vehicle) ID(vehiclePtr unsafe.Pointer) int {
	return int(C.Vehicle_GetID(vehiclePtr))
}

func (v *Vehicle) FromID(vehicleid int) unsafe.Pointer {
	return C.Vehicle_FromID(C.int(vehicleid))
}
