package main

/*
#cgo CFLAGS: -I${SRCDIR}/include/bridge
#include "bridge.h"
*/
import "C"

import (
	"unsafe"
)

type Vehicle struct{}

// =====================================================================
// Vehicle Functions
// =====================================================================

// Create creates a new vehicle and returns vehicle ID
func (Vehicle) Create(model int, x, y, z, rotation float32, color1, color2 int, respawnDelay int, addsiren bool) int {
	return int(C.Vehicle_Create(C.int(model), C.float(x), C.float(y), C.float(z), C.float(rotation), C.int(color1), C.int(color2), C.int(respawnDelay), C.bool(addsiren)))
}

// FromID returns vehicle pointer from ID
func VehicleFromID(vehicleid int) *Vehicle {
	ptr := C.Vehicle_FromID(C.int(vehicleid))
	return (*Vehicle)(ptr)
}

// GetID returns vehicle ID
func (v *Vehicle) GetID() int {
	return int(C.Vehicle_GetID(unsafe.Pointer(v)))
}

// Destroy destroys the vehicle
func (v *Vehicle) Destroy() bool {
	return bool(C.Vehicle_Destroy(unsafe.Pointer(v)))
}

// GetMaxPassengerSeats returns max passenger seats for model (static)
func VehicleGetMaxPassengerSeats(modelid int) int {
	return int(C.Vehicle_GetMaxPassengerSeats(C.int(modelid)))
}

// IsStreamedIn checks if vehicle is streamed in for player
func (v *Vehicle) IsStreamedIn(player *Player) bool {
	return bool(C.Vehicle_IsStreamedIn(unsafe.Pointer(v), unsafe.Pointer(player)))
}

// GetPos returns vehicle position
func (v *Vehicle) GetPos() (x, y, z float32, ok bool) {
	var cX, cY, cZ C.float
	ok = bool(C.Vehicle_GetPos(unsafe.Pointer(v), &cX, &cY, &cZ))
	return float32(cX), float32(cY), float32(cZ), ok
}

// SetPos sets vehicle position
func (v *Vehicle) SetPos(x, y, z float32) bool {
	return bool(C.Vehicle_SetPos(unsafe.Pointer(v), C.float(x), C.float(y), C.float(z)))
}

// GetZAngle returns vehicle Z angle
func (v *Vehicle) GetZAngle() float32 {
	return float32(C.Vehicle_GetZAngle(unsafe.Pointer(v)))
}

// GetRotationQuat returns vehicle rotation quaternion
func (v *Vehicle) GetRotationQuat() (w, x, y, z float32, ok bool) {
	var cW, cX, cY, cZ C.float
	ok = bool(C.Vehicle_GetRotationQuat(unsafe.Pointer(v), &cW, &cX, &cY, &cZ))
	return float32(cW), float32(cX), float32(cY), float32(cZ), ok
}

// GetDistanceFromPoint returns distance from point
func (v *Vehicle) GetDistanceFromPoint(x, y, z float32) float32 {
	return float32(C.Vehicle_GetDistanceFromPoint(unsafe.Pointer(v), C.float(x), C.float(y), C.float(z)))
}

// SetZAngle sets vehicle Z angle
func (v *Vehicle) SetZAngle(angle float32) bool {
	return bool(C.Vehicle_SetZAngle(unsafe.Pointer(v), C.float(angle)))
}

// SetParamsForPlayer sets params for player
func (v *Vehicle) SetParamsForPlayer(player *Player, objective, doors int) bool {
	return bool(C.Vehicle_SetParamsForPlayer(unsafe.Pointer(v), unsafe.Pointer(player), C.int(objective), C.int(doors)))
}

// UseManualEngineAndLights enables manual engine and lights (static)
func VehicleUseManualEngineAndLights() bool {
	return bool(C.Vehicle_UseManualEngineAndLights())
}

// SetParamsEx sets extended params
func (v *Vehicle) SetParamsEx(engine, lights, alarm, doors, bonnet, boot, objective int) bool {
	return bool(C.Vehicle_SetParamsEx(unsafe.Pointer(v), C.int(engine), C.int(lights), C.int(alarm), C.int(doors), C.int(bonnet), C.int(boot), C.int(objective)))
}

// GetParamsEx returns extended params
func (v *Vehicle) GetParamsEx() (engine, lights, alarm, doors, bonnet, boot, objective int, ok bool) {
	var cEngine, cLights, cAlarm, cDoors, cBonnet, cBoot, cObjective C.int
	ok = bool(C.Vehicle_GetParamsEx(unsafe.Pointer(v), &cEngine, &cLights, &cAlarm, &cDoors, &cBonnet, &cBoot, &cObjective))
	return int(cEngine), int(cLights), int(cAlarm), int(cDoors), int(cBonnet), int(cBoot), int(cObjective), ok
}

// GetParamsSirenState returns siren state
func (v *Vehicle) GetParamsSirenState() int {
	return int(C.Vehicle_GetParamsSirenState(unsafe.Pointer(v)))
}

// SetParamsCarDoors sets car doors params
func (v *Vehicle) SetParamsCarDoors(frontLeft, frontRight, rearLeft, rearRight int) bool {
	return bool(C.Vehicle_SetParamsCarDoors(unsafe.Pointer(v), C.int(frontLeft), C.int(frontRight), C.int(rearLeft), C.int(rearRight)))
}

// GetParamsCarDoors returns car doors params
func (v *Vehicle) GetParamsCarDoors() (frontLeft, frontRight, rearLeft, rearRight int, ok bool) {
	var cFrontLeft, cFrontRight, cRearLeft, cRearRight C.int
	ok = bool(C.Vehicle_GetParamsCarDoors(unsafe.Pointer(v), &cFrontLeft, &cFrontRight, &cRearLeft, &cRearRight))
	return int(cFrontLeft), int(cFrontRight), int(cRearLeft), int(cRearRight), ok
}

// SetParamsCarWindows sets car windows params
func (v *Vehicle) SetParamsCarWindows(frontLeft, frontRight, rearLeft, rearRight int) bool {
	return bool(C.Vehicle_SetParamsCarWindows(unsafe.Pointer(v), C.int(frontLeft), C.int(frontRight), C.int(rearLeft), C.int(rearRight)))
}

// GetParamsCarWindows returns car windows params
func (v *Vehicle) GetParamsCarWindows() (frontLeft, frontRight, rearLeft, rearRight int, ok bool) {
	var cFrontLeft, cFrontRight, cRearLeft, cRearRight C.int
	ok = bool(C.Vehicle_GetParamsCarWindows(unsafe.Pointer(v), &cFrontLeft, &cFrontRight, &cRearLeft, &cRearRight))
	return int(cFrontLeft), int(cFrontRight), int(cRearLeft), int(cRearRight), ok
}

// SetToRespawn sets vehicle to respawn
func (v *Vehicle) SetToRespawn() bool {
	return bool(C.Vehicle_SetToRespawn(unsafe.Pointer(v)))
}

// LinkToInterior links vehicle to interior
func (v *Vehicle) LinkToInterior(interiorid int) bool {
	return bool(C.Vehicle_LinkToInterior(unsafe.Pointer(v), C.int(interiorid)))
}

// AddComponent adds component to vehicle
func (v *Vehicle) AddComponent(componentid int) bool {
	return bool(C.Vehicle_AddComponent(unsafe.Pointer(v), C.int(componentid)))
}

// RemoveComponent removes component from vehicle
func (v *Vehicle) RemoveComponent(componentid int) bool {
	return bool(C.Vehicle_RemoveComponent(unsafe.Pointer(v), C.int(componentid)))
}

// ChangeColor changes vehicle color
func (v *Vehicle) ChangeColor(color1, color2 int) bool {
	return bool(C.Vehicle_ChangeColor(unsafe.Pointer(v), C.int(color1), C.int(color2)))
}

// ChangePaintjob changes vehicle paintjob
func (v *Vehicle) ChangePaintjob(paintjobid int) bool {
	return bool(C.Vehicle_ChangePaintjob(unsafe.Pointer(v), C.int(paintjobid)))
}

// SetHealth sets vehicle health
func (v *Vehicle) SetHealth(health float32) bool {
	return bool(C.Vehicle_SetHealth(unsafe.Pointer(v), C.float(health)))
}

// GetHealth returns vehicle health
func (v *Vehicle) GetHealth() float32 {
	return float32(C.Vehicle_GetHealth(unsafe.Pointer(v)))
}

// AttachTrailer attaches trailer to vehicle
func (v *Vehicle) AttachTrailer(trailer *Vehicle) bool {
	return bool(C.Vehicle_AttachTrailer(unsafe.Pointer(trailer), unsafe.Pointer(v)))
}

// DetachTrailer detaches trailer from vehicle
func (v *Vehicle) DetachTrailer() bool {
	return bool(C.Vehicle_DetachTrailer(unsafe.Pointer(v)))
}

// IsTrailerAttached checks if trailer is attached
func (v *Vehicle) IsTrailerAttached() bool {
	return bool(C.Vehicle_IsTrailerAttached(unsafe.Pointer(v)))
}

// GetTrailer returns attached trailer
func (v *Vehicle) GetTrailer() *Vehicle {
	ptr := C.Vehicle_GetTrailer(unsafe.Pointer(v))
	return (*Vehicle)(ptr)
}

// SetNumberPlate sets vehicle number plate
func (v *Vehicle) SetNumberPlate(numberPlate string) bool {
	cNumberPlate := C.CString(numberPlate)
	defer C.free(unsafe.Pointer(cNumberPlate))
	return bool(C.Vehicle_SetNumberPlate(unsafe.Pointer(v), cNumberPlate))
}

// GetModel returns vehicle model
func (v *Vehicle) GetModel() int {
	return int(C.Vehicle_GetModel(unsafe.Pointer(v)))
}

// GetComponentInSlot returns component in slot
func (v *Vehicle) GetComponentInSlot(slot int) int {
	return int(C.Vehicle_GetComponentInSlot(unsafe.Pointer(v), C.int(slot)))
}

// GetComponentType returns component type (static)
func VehicleGetComponentType(componentid int) int {
	return int(C.Vehicle_GetComponentType(C.int(componentid)))
}

// CanHaveComponent checks if model can have component (static)
func VehicleCanHaveComponent(modelid, componentid int) bool {
	return bool(C.Vehicle_CanHaveComponent(C.int(modelid), C.int(componentid)))
}

// GetRandomColorPair returns random color pair for model (static)
func VehicleGetRandomColorPair(modelid int) (color1, color2, color3, color4 int, ok bool) {
	var cColor1, cColor2, cColor3, cColor4 C.int
	ok = bool(C.Vehicle_GetRandomColorPair(C.int(modelid), &cColor1, &cColor2, &cColor3, &cColor4))
	return int(cColor1), int(cColor2), int(cColor3), int(cColor4), ok
}

// ColorIndexToColor converts color index to color (static)
func VehicleColorIndexToColor(colorIndex, alpha int) int {
	return int(C.Vehicle_ColorIndexToColor(C.int(colorIndex), C.int(alpha)))
}

// Repair repairs vehicle
func (v *Vehicle) Repair() bool {
	return bool(C.Vehicle_Repair(unsafe.Pointer(v)))
}

// GetVelocity returns vehicle velocity
func (v *Vehicle) GetVelocity() (x, y, z float32, ok bool) {
	var cX, cY, cZ C.float
	ok = bool(C.Vehicle_GetVelocity(unsafe.Pointer(v), &cX, &cY, &cZ))
	return float32(cX), float32(cY), float32(cZ), ok
}

// SetVelocity sets vehicle velocity
func (v *Vehicle) SetVelocity(x, y, z float32) bool {
	return bool(C.Vehicle_SetVelocity(unsafe.Pointer(v), C.float(x), C.float(y), C.float(z)))
}

// SetAngularVelocity sets vehicle angular velocity
func (v *Vehicle) SetAngularVelocity(x, y, z float32) bool {
	return bool(C.Vehicle_SetAngularVelocity(unsafe.Pointer(v), C.float(x), C.float(y), C.float(z)))
}

// GetDamageStatus returns vehicle damage status
func (v *Vehicle) GetDamageStatus() (panels, doors, lights, tires int, ok bool) {
	var cPanels, cDoors, cLights, cTires C.int
	ok = bool(C.Vehicle_GetDamageStatus(unsafe.Pointer(v), &cPanels, &cDoors, &cLights, &cTires))
	return int(cPanels), int(cDoors), int(cLights), int(cTires), ok
}

// UpdateDamageStatus updates vehicle damage status
func (v *Vehicle) UpdateDamageStatus(panels, doors, lights, tires int) bool {
	return bool(C.Vehicle_UpdateDamageStatus(unsafe.Pointer(v), C.int(panels), C.int(doors), C.int(lights), C.int(tires)))
}

// GetModelInfo returns model info (static)
func VehicleGetModelInfo(vehiclemodel, infotype int) (x, y, z float32, ok bool) {
	var cX, cY, cZ C.float
	ok = bool(C.Vehicle_GetModelInfo(C.int(vehiclemodel), C.int(infotype), &cX, &cY, &cZ))
	return float32(cX), float32(cY), float32(cZ), ok
}

// SetVirtualWorld sets vehicle virtual world
func (v *Vehicle) SetVirtualWorld(virtualWorld int) bool {
	return bool(C.Vehicle_SetVirtualWorld(unsafe.Pointer(v), C.int(virtualWorld)))
}

// GetVirtualWorld returns vehicle virtual world
func (v *Vehicle) GetVirtualWorld() int {
	return int(C.Vehicle_GetVirtualWorld(unsafe.Pointer(v)))
}

// GetLandingGearState returns landing gear state
func (v *Vehicle) GetLandingGearState() int {
	return int(C.Vehicle_GetLandingGearState(unsafe.Pointer(v)))
}

// IsValid checks if vehicle is valid
func (v *Vehicle) IsValid() bool {
	return bool(C.Vehicle_IsValid(unsafe.Pointer(v)))
}

// AddStatic adds static vehicle (static)
func VehicleAddStatic(modelid int, x, y, z, angle float32, color1, color2 int) int {
	return int(C.Vehicle_AddStatic(C.int(modelid), C.float(x), C.float(y), C.float(z), C.float(angle), C.int(color1), C.int(color2)))
}

// AddStaticEx adds static vehicle with extra options (static)
func VehicleAddStaticEx(modelid int, x, y, z, angle float32, color1, color2, respawnDelay int, addSiren bool) int {
	return int(C.Vehicle_AddStaticEx(C.int(modelid), C.float(x), C.float(y), C.float(z), C.float(angle), C.int(color1), C.int(color2), C.int(respawnDelay), C.bool(addSiren)))
}

// EnableFriendlyFire enables friendly fire (static)
func VehicleEnableFriendlyFire() bool {
	return bool(C.Vehicle_EnableFriendlyFire())
}

// GetSpawnInfo returns vehicle spawn info
func (v *Vehicle) GetSpawnInfo() (x, y, z, rotation float32, color1, color2 int, ok bool) {
	var cX, cY, cZ, cRotation C.float
	var cColor1, cColor2 C.int
	ok = bool(C.Vehicle_GetSpawnInfo(unsafe.Pointer(v), &cX, &cY, &cZ, &cRotation, &cColor1, &cColor2))
	return float32(cX), float32(cY), float32(cZ), float32(cRotation), int(cColor1), int(cColor2), ok
}

// SetSpawnInfo sets vehicle spawn info
func (v *Vehicle) SetSpawnInfo(modelid int, x, y, z, rotation float32, color1, color2, respawnTime, interior int) bool {
	return bool(C.Vehicle_SetSpawnInfo(unsafe.Pointer(v), C.int(modelid), C.float(x), C.float(y), C.float(z), C.float(rotation), C.int(color1), C.int(color2), C.int(respawnTime), C.int(interior)))
}

// GetModelCount returns model count (static)
func VehicleGetModelCount(modelid int) int {
	return int(C.Vehicle_GetModelCount(C.int(modelid)))
}

// GetModelsUsed returns models used count (static)
func VehicleGetModelsUsed() int {
	return int(C.Vehicle_GetModelsUsed())
}

// GetPaintjob returns vehicle paintjob
func (v *Vehicle) GetPaintjob() int {
	return int(C.Vehicle_GetPaintjob(unsafe.Pointer(v)))
}

// GetColor returns vehicle colors
func (v *Vehicle) GetColor() (color1, color2 int, ok bool) {
	var cColor1, cColor2 C.int
	ok = bool(C.Vehicle_GetColor(unsafe.Pointer(v), &cColor1, &cColor2))
	return int(cColor1), int(cColor2), ok
}

// GetInterior returns vehicle interior
func (v *Vehicle) GetInterior() int {
	return int(C.Vehicle_GetInterior(unsafe.Pointer(v)))
}

// GetNumberPlate returns vehicle number plate
func (v *Vehicle) GetNumberPlate() string {
	var numberPlate C.struct_CAPIStringView
	C.Vehicle_GetNumberPlate(unsafe.Pointer(v), &numberPlate)
	return C.GoStringN(numberPlate.data, C.int(numberPlate.len))
}

// SetRespawnDelay sets vehicle respawn delay
func (v *Vehicle) SetRespawnDelay(respawnDelay int) bool {
	return bool(C.Vehicle_SetRespawnDelay(unsafe.Pointer(v), C.int(respawnDelay)))
}

// GetRespawnDelay returns vehicle respawn delay
func (v *Vehicle) GetRespawnDelay() int {
	return int(C.Vehicle_GetRespawnDelay(unsafe.Pointer(v)))
}

// GetCab returns vehicle cab
func (v *Vehicle) GetCab() *Vehicle {
	ptr := C.Vehicle_GetCab(unsafe.Pointer(v))
	return (*Vehicle)(ptr)
}

// GetTower returns vehicle tower
func (v *Vehicle) GetTower() *Vehicle {
	ptr := C.Vehicle_GetTower(unsafe.Pointer(v))
	return (*Vehicle)(ptr)
}

// GetOccupiedTick returns vehicle occupied tick
func (v *Vehicle) GetOccupiedTick() int {
	return int(C.Vehicle_GetOccupiedTick(unsafe.Pointer(v)))
}

// GetRespawnTick returns vehicle respawn tick
func (v *Vehicle) GetRespawnTick() int {
	return int(C.Vehicle_GetRespawnTick(unsafe.Pointer(v)))
}

// HasBeenOccupied checks if vehicle has been occupied
func (v *Vehicle) HasBeenOccupied() bool {
	return bool(C.Vehicle_HasBeenOccupied(unsafe.Pointer(v)))
}

// IsOccupied checks if vehicle is occupied
func (v *Vehicle) IsOccupied() bool {
	return bool(C.Vehicle_IsOccupied(unsafe.Pointer(v)))
}

// IsDead checks if vehicle is dead
func (v *Vehicle) IsDead() bool {
	return bool(C.Vehicle_IsDead(unsafe.Pointer(v)))
}

// SetParamsSirenState sets siren state
func (v *Vehicle) SetParamsSirenState(sirenState bool) bool {
	return bool(C.Vehicle_SetParamsSirenState(unsafe.Pointer(v), C.bool(sirenState)))
}

// ToggleSirenEnabled toggles siren enabled
func (v *Vehicle) ToggleSirenEnabled(status bool) bool {
	return bool(C.Vehicle_ToggleSirenEnabled(unsafe.Pointer(v), C.bool(status)))
}

// IsSirenEnabled checks if siren is enabled
func (v *Vehicle) IsSirenEnabled() bool {
	return bool(C.Vehicle_IsSirenEnabled(unsafe.Pointer(v)))
}

// GetLastDriver returns last driver
func (v *Vehicle) GetLastDriver() *Player {
	ptr := C.Vehicle_GetLastDriver(unsafe.Pointer(v))
	return (*Player)(ptr)
}

// GetDriver returns current driver
func (v *Vehicle) GetDriver() *Player {
	ptr := C.Vehicle_GetDriver(unsafe.Pointer(v))
	return (*Player)(ptr)
}

// GetSirenState returns siren state
func (v *Vehicle) GetSirenState() int {
	return int(C.Vehicle_GetSirenState(unsafe.Pointer(v)))
}

// GetHydraReactorAngle returns hydra reactor angle
func (v *Vehicle) GetHydraReactorAngle() uint32 {
	return uint32(C.Vehicle_GetHydraReactorAngle(unsafe.Pointer(v)))
}

// GetTrainSpeed returns train speed
func (v *Vehicle) GetTrainSpeed() float32 {
	return float32(C.Vehicle_GetTrainSpeed(unsafe.Pointer(v)))
}

// GetMatrix returns vehicle matrix
func (v *Vehicle) GetMatrix() (rightX, rightY, rightZ, upX, upY, upZ, atX, atY, atZ float32, ok bool) {
	var cRightX, cRightY, cRightZ, cUpX, cUpY, cUpZ, cAtX, cAtY, cAtZ C.float
	ok = bool(C.Vehicle_GetMatrix(unsafe.Pointer(v), &cRightX, &cRightY, &cRightZ, &cUpX, &cUpY, &cUpZ, &cAtX, &cAtY, &cAtZ))
	return float32(cRightX), float32(cRightY), float32(cRightZ), float32(cUpX), float32(cUpY), float32(cUpZ), float32(cAtX), float32(cAtY), float32(cAtZ), ok
}

// GetOccupant returns vehicle occupant in seat
func (v *Vehicle) GetOccupant(seat int) *Player {
	ptr := C.Vehicle_GetOccupant(unsafe.Pointer(v), C.int(seat))
	return (*Player)(ptr)
}

// CountOccupants returns vehicle occupant count
func (v *Vehicle) CountOccupants() int {
	return int(C.Vehicle_CountOccupants(unsafe.Pointer(v)))
}
