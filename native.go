package main

/*
#cgo CFLAGS: -I${SRCDIR}/include/bridge
#include "bridge.h"
*/
import "C"

import "unsafe"

type Native struct {
	core    Core
	vehicle Vehicle
}

type Core struct{}

func (Core) Log(message string) {
	cMessage := C.CString(message)
	defer C.free(unsafe.Pointer(cMessage))
	C.Native_Core_Log(cMessage)
}

func (Core) SendClientMessageToAll(color uint32, message string) bool {
	cMessage := C.CString(message)
	defer C.free(unsafe.Pointer(cMessage))
	result := C.SendClientMessageToAll(C.uint32_t(color), cMessage)
	return bool(result)
}

func (Core) AddPlayerClass(team uint8, skin int, x, y, z, angle float32, weapon1 uint8, ammo1 uint32, weapon2 uint8, ammo2 uint32, weapon3 uint8, ammo3 uint32) int {
	return int(C.Native_AddPlayerClass(C.uint8_t(team), C.int(skin), C.float(x), C.float(y), C.float(z), C.float(angle), C.uint8_t(weapon1), C.uint32_t(ammo1), C.uint8_t(weapon2), C.uint32_t(ammo2), C.uint8_t(weapon3), C.uint32_t(ammo3)))
}

var Natives Native
