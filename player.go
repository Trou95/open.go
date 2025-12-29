package main

/*
#cgo CFLAGS: -I${SRCDIR}/include/bridge
#include "bridge.h"
*/
import "C"

import (
	"unsafe"
)

type Player struct{}

// =====================================================================
// Player Functions - Following order from capi/ompcapi.h
// =====================================================================

func (p *Player) ID() int {
	return int(C.Player_GetID(unsafe.Pointer(p)))
}

// SetSpawnInfo sets player spawn information
func (p *Player) SetSpawnInfo(team uint8, skin int, x, y, z, angle float32, weapon1 uint8, ammo1 uint32, weapon2 uint8, ammo2 uint32, weapon3 uint8, ammo3 uint32) bool {
	return bool(C.Player_SetSpawnInfo(unsafe.Pointer(p), C.uint8_t(team), C.int(skin), C.float(x), C.float(y), C.float(z), C.float(angle), C.uint8_t(weapon1), C.uint32_t(ammo1), C.uint8_t(weapon2), C.uint32_t(ammo2), C.uint8_t(weapon3), C.uint32_t(ammo3)))
}

// GetSpawnInfo gets player spawn information
func (p *Player) GetSpawnInfo() (team uint8, skin int, x, y, z, angle float32, weapon1 uint8, ammo1 uint32, weapon2 uint8, ammo2 uint32, weapon3 uint8, ammo3 uint32, ok bool) {
	var cTeam C.uint8_t
	var cSkin C.int
	var cX, cY, cZ, cAngle C.float
	var cWeapon1, cWeapon2, cWeapon3 C.uint8_t
	var cAmmo1, cAmmo2, cAmmo3 C.uint32_t
	ok = bool(C.Player_GetSpawnInfo(unsafe.Pointer(p), &cTeam, &cSkin, &cX, &cY, &cZ, &cAngle, &cWeapon1, &cAmmo1, &cWeapon2, &cAmmo2, &cWeapon3, &cAmmo3))
	return uint8(cTeam), int(cSkin), float32(cX), float32(cY), float32(cZ), float32(cAngle), uint8(cWeapon1), uint32(cAmmo1), uint8(cWeapon2), uint32(cAmmo2), uint8(cWeapon3), uint32(cAmmo3), ok
}

// NetStatsBytesReceived returns bytes received
func (p *Player) NetStatsBytesReceived() int {
	return int(C.Player_NetStatsBytesReceived(unsafe.Pointer(p)))
}

// NetStatsBytesSent returns bytes sent
func (p *Player) NetStatsBytesSent() int {
	return int(C.Player_NetStatsBytesSent(unsafe.Pointer(p)))
}

// NetStatsConnectionStatus returns connection status
func (p *Player) NetStatsConnectionStatus() int {
	return int(C.Player_NetStatsConnectionStatus(unsafe.Pointer(p)))
}

// NetStatsGetConnectedTime returns connected time
func (p *Player) NetStatsGetConnectedTime() int {
	return int(C.Player_NetStatsGetConnectedTime(unsafe.Pointer(p)))
}

// NetStatsMessagesReceived returns messages received
func (p *Player) NetStatsMessagesReceived() int {
	return int(C.Player_NetStatsMessagesReceived(unsafe.Pointer(p)))
}

// NetStatsMessagesRecvPerSecond returns messages received per second
func (p *Player) NetStatsMessagesRecvPerSecond() int {
	return int(C.Player_NetStatsMessagesRecvPerSecond(unsafe.Pointer(p)))
}

// NetStatsMessagesSent returns messages sent
func (p *Player) NetStatsMessagesSent() int {
	return int(C.Player_NetStatsMessagesSent(unsafe.Pointer(p)))
}

// NetStatsPacketLossPercent returns packet loss percentage
func (p *Player) NetStatsPacketLossPercent() float32 {
	return float32(C.Player_NetStatsPacketLossPercent(unsafe.Pointer(p)))
}

// GetNetworkStats returns network stats as string
func (p *Player) GetNetworkStats() string {
	buf := make([]byte, 512)
	var output C.struct_CAPIStringBuffer
	output.data = (*C.char)(unsafe.Pointer(&buf[0]))
	output.capacity = C.uint(len(buf))
	C.Player_GetNetworkStats(unsafe.Pointer(p), &output)
	return string(buf[:output.len])
}

// NetStatsGetIpPort returns IP and port as string
func (p *Player) NetStatsGetIpPort() (string, bool) {
	buf := make([]byte, 64)
	var output C.struct_CAPIStringBuffer
	output.data = (*C.char)(unsafe.Pointer(&buf[0]))
	output.capacity = C.uint(len(buf))
	ok := bool(C.Player_NetStatsGetIpPort(unsafe.Pointer(p), &output))
	return string(buf[:output.len]), ok
}

// GetCustomSkin returns custom skin
func (p *Player) GetCustomSkin() int {
	return int(C.Player_GetCustomSkin(unsafe.Pointer(p)))
}

// GetDialog returns dialog ID
func (p *Player) GetDialog() int {
	return int(C.Player_GetDialog(unsafe.Pointer(p)))
}

// GetDialogData returns dialog data
func (p *Player) GetDialogData() (dialogid, style int, title, body, button1, button2 string, ok bool) {
	var cDialogid, cStyle C.int
	var cTitle, cBody, cButton1, cButton2 C.struct_CAPIStringView
	ok = bool(C.Player_GetDialogData(unsafe.Pointer(p), &cDialogid, &cStyle, &cTitle, &cBody, &cButton1, &cButton2))
	return int(cDialogid), int(cStyle), C.GoStringN(cTitle.data, C.int(cTitle.len)), C.GoStringN(cBody.data, C.int(cBody.len)), C.GoStringN(cButton1.data, C.int(cButton1.len)), C.GoStringN(cButton2.data, C.int(cButton2.len)), ok
}

// GetMenu returns menu pointer
func (p *Player) GetMenu() unsafe.Pointer {
	return unsafe.Pointer(C.Player_GetMenu(unsafe.Pointer(p)))
}

// GetSurfingPlayerObject returns surfing player object pointer
func (p *Player) GetSurfingPlayerObject() unsafe.Pointer {
	return unsafe.Pointer(C.Player_GetSurfingPlayerObject(unsafe.Pointer(p)))
}

// GetCameraTargetPlayerObject returns camera target player object pointer
func (p *Player) GetCameraTargetPlayerObject() unsafe.Pointer {
	return unsafe.Pointer(C.Player_GetCameraTargetPlayerObject(unsafe.Pointer(p)))
}

// FromID returns player pointer from ID
func PlayerFromID(playerid int) *Player {
	ptr := C.Player_FromID(C.int(playerid))
	return (*Player)(ptr)
}

// GetAnimationName returns animation library and name by index (static function)
func GetAnimationName(index int) (lib, name string, ok bool) {
	var cLib, cName C.struct_CAPIStringView
	ok = bool(C.Player_GetAnimationName(C.int(index), &cLib, &cName))
	return C.GoStringN(cLib.data, C.int(cLib.len)), C.GoStringN(cName.data, C.int(cName.len)), ok
}

// GetID returns player ID
func (p *Player) GetID() int {
	return int(C.Player_GetID(unsafe.Pointer(p)))
}

// SendClientMessage sends a client message
func (p *Player) SendClientMessage(color uint32, text string) bool {
	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))
	return bool(C.Player_SendClientMessage(unsafe.Pointer(p), C.uint32_t(color), cText))
}

// SetCameraPos sets camera position
func (p *Player) SetCameraPos(x, y, z float32) bool {
	return bool(C.Player_SetCameraPos(unsafe.Pointer(p), C.float(x), C.float(y), C.float(z)))
}

// SetDrunkLevel sets drunk level
func (p *Player) SetDrunkLevel(level int) bool {
	return bool(C.Player_SetDrunkLevel(unsafe.Pointer(p), C.int(level)))
}

// SetInterior sets interior
func (p *Player) SetInterior(interior int) bool {
	return bool(C.Player_SetInterior(unsafe.Pointer(p), C.int(interior)))
}

// SetWantedLevel sets wanted level
func (p *Player) SetWantedLevel(level int) bool {
	return bool(C.Player_SetWantedLevel(unsafe.Pointer(p), C.int(level)))
}

// SetWeather sets weather
func (p *Player) SetWeather(weather int) bool {
	return bool(C.Player_SetWeather(unsafe.Pointer(p), C.int(weather)))
}

// GetWeather returns weather
func (p *Player) GetWeather() int {
	return int(C.Player_GetWeather(unsafe.Pointer(p)))
}

// SetSkin sets skin
func (p *Player) SetSkin(skin int) bool {
	return bool(C.Player_SetSkin(unsafe.Pointer(p), C.int(skin)))
}

// SetShopName sets shop name
func (p *Player) SetShopName(name string) bool {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	return bool(C.Player_SetShopName(unsafe.Pointer(p), cName))
}

// GiveMoney gives money
func (p *Player) GiveMoney(amount int) bool {
	return bool(C.Player_GiveMoney(unsafe.Pointer(p), C.int(amount)))
}

// SetCameraLookAt sets camera look at position
func (p *Player) SetCameraLookAt(x, y, z float32, cutType int) bool {
	return bool(C.Player_SetCameraLookAt(unsafe.Pointer(p), C.float(x), C.float(y), C.float(z), C.int(cutType)))
}

// SetCameraBehind sets camera behind player
func (p *Player) SetCameraBehind() bool {
	return bool(C.Player_SetCameraBehind(unsafe.Pointer(p)))
}

// CreateExplosion creates explosion for player
func (p *Player) CreateExplosion(x, y, z float32, explosionType int, radius float32) bool {
	return bool(C.Player_CreateExplosion(unsafe.Pointer(p), C.float(x), C.float(y), C.float(z), C.int(explosionType), C.float(radius)))
}

// PlayAudioStream plays audio stream
func (p *Player) PlayAudioStream(url string, x, y, z, distance float32, usePos bool) bool {
	cUrl := C.CString(url)
	defer C.free(unsafe.Pointer(cUrl))
	return bool(C.Player_PlayAudioStream(unsafe.Pointer(p), cUrl, C.float(x), C.float(y), C.float(z), C.float(distance), C.bool(usePos)))
}

// StopAudioStream stops audio stream
func (p *Player) StopAudioStream() bool {
	return bool(C.Player_StopAudioStream(unsafe.Pointer(p)))
}

// ToggleWidescreen toggles widescreen
func (p *Player) ToggleWidescreen(enable bool) bool {
	return bool(C.Player_ToggleWidescreen(unsafe.Pointer(p), C.bool(enable)))
}

// IsWidescreenToggled checks if widescreen is toggled
func (p *Player) IsWidescreenToggled() bool {
	return bool(C.Player_IsWidescreenToggled(unsafe.Pointer(p)))
}

// SetHealth sets health
func (p *Player) SetHealth(health float32) bool {
	return bool(C.Player_SetHealth(unsafe.Pointer(p), C.float(health)))
}

// GetHealth returns health
func (p *Player) GetHealth() float32 {
	return float32(C.Player_GetHealth(unsafe.Pointer(p)))
}

// SetArmor sets armor
func (p *Player) SetArmor(armor float32) bool {
	return bool(C.Player_SetArmor(unsafe.Pointer(p), C.float(armor)))
}

// GetArmor returns armor
func (p *Player) GetArmor() float32 {
	return float32(C.Player_GetArmor(unsafe.Pointer(p)))
}

// SetTeam sets team
func (p *Player) SetTeam(team int) bool {
	return bool(C.Player_SetTeam(unsafe.Pointer(p), C.int(team)))
}

// GetTeam returns team
func (p *Player) GetTeam() int {
	return int(C.Player_GetTeam(unsafe.Pointer(p)))
}

// SetScore sets score
func (p *Player) SetScore(score int) bool {
	return bool(C.Player_SetScore(unsafe.Pointer(p), C.int(score)))
}

// GetScore returns score
func (p *Player) GetScore() int {
	return int(C.Player_GetScore(unsafe.Pointer(p)))
}

// GetSkin returns skin
func (p *Player) GetSkin() int {
	return int(C.Player_GetSkin(unsafe.Pointer(p)))
}

// SetColor sets color
func (p *Player) SetColor(color uint32) bool {
	return bool(C.Player_SetColor(unsafe.Pointer(p), C.uint32_t(color)))
}

// GetColor returns color
func (p *Player) GetColor() uint32 {
	return uint32(C.Player_GetColor(unsafe.Pointer(p)))
}

// GetDefaultColor returns default color
func (p *Player) GetDefaultColor() uint32 {
	return uint32(C.Player_GetDefaultColor(unsafe.Pointer(p)))
}

// GetDrunkLevel returns drunk level
func (p *Player) GetDrunkLevel() int {
	return int(C.Player_GetDrunkLevel(unsafe.Pointer(p)))
}

// GiveWeapon gives weapon
func (p *Player) GiveWeapon(weapon, ammo int) bool {
	return bool(C.Player_GiveWeapon(unsafe.Pointer(p), C.int(weapon), C.int(ammo)))
}

// RemoveWeapon removes weapon
func (p *Player) RemoveWeapon(weapon int) bool {
	return bool(C.Player_RemoveWeapon(unsafe.Pointer(p), C.int(weapon)))
}

// GetMoney returns money
func (p *Player) GetMoney() int {
	return int(C.Player_GetMoney(unsafe.Pointer(p)))
}

// ResetMoney resets money
func (p *Player) ResetMoney() bool {
	return bool(C.Player_ResetMoney(unsafe.Pointer(p)))
}

// SetName sets name
func (p *Player) SetName(name string) int {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	return int(C.Player_SetName(unsafe.Pointer(p), cName))
}

// GetName returns player name
func (p *Player) GetName() string {
	var nameView C.struct_CAPIStringView
	C.Player_GetName(unsafe.Pointer(p), &nameView)
	return C.GoStringN(nameView.data, C.int(nameView.len))
}

// GetState returns state
func (p *Player) GetState() int {
	return int(C.Player_GetState(unsafe.Pointer(p)))
}

// GetPing returns ping
func (p *Player) GetPing() int {
	return int(C.Player_GetPing(unsafe.Pointer(p)))
}

// GetWeapon returns current weapon
func (p *Player) GetWeapon() int {
	return int(C.Player_GetWeapon(unsafe.Pointer(p)))
}

// SetTime sets time
func (p *Player) SetTime(hour, minute int) bool {
	return bool(C.Player_SetTime(unsafe.Pointer(p), C.int(hour), C.int(minute)))
}

// GetTime returns time
func (p *Player) GetTime() (hour, minute int, ok bool) {
	var cHour, cMinute C.int
	ok = bool(C.Player_GetTime(unsafe.Pointer(p), &cHour, &cMinute))
	return int(cHour), int(cMinute), ok
}

// ToggleClock toggles clock
func (p *Player) ToggleClock(enable bool) bool {
	return bool(C.Player_ToggleClock(unsafe.Pointer(p), C.bool(enable)))
}

// HasClock checks if player has clock
func (p *Player) HasClock() bool {
	return bool(C.Player_HasClock(unsafe.Pointer(p)))
}

// ForceClassSelection forces class selection
func (p *Player) ForceClassSelection() bool {
	return bool(C.Player_ForceClassSelection(unsafe.Pointer(p)))
}

// GetWantedLevel returns wanted level
func (p *Player) GetWantedLevel() int {
	return int(C.Player_GetWantedLevel(unsafe.Pointer(p)))
}

// SetFightingStyle sets fighting style
func (p *Player) SetFightingStyle(style int) bool {
	return bool(C.Player_SetFightingStyle(unsafe.Pointer(p), C.int(style)))
}

// GetFightingStyle returns fighting style
func (p *Player) GetFightingStyle() int {
	return int(C.Player_GetFightingStyle(unsafe.Pointer(p)))
}

// SetVelocity sets velocity
func (p *Player) SetVelocity(x, y, z float32) bool {
	return bool(C.Player_SetVelocity(unsafe.Pointer(p), C.float(x), C.float(y), C.float(z)))
}

// GetVelocity returns velocity
func (p *Player) GetVelocity() (x, y, z float32, ok bool) {
	var cX, cY, cZ C.float
	ok = bool(C.Player_GetVelocity(unsafe.Pointer(p), &cX, &cY, &cZ))
	return float32(cX), float32(cY), float32(cZ), ok
}

// GetCameraPos returns camera position
func (p *Player) GetCameraPos() (x, y, z float32, ok bool) {
	var cX, cY, cZ C.float
	ok = bool(C.Player_GetCameraPos(unsafe.Pointer(p), &cX, &cY, &cZ))
	return float32(cX), float32(cY), float32(cZ), ok
}

// GetDistanceFromPoint returns distance from point
func (p *Player) GetDistanceFromPoint(x, y, z float32) float32 {
	return float32(C.Player_GetDistanceFromPoint(unsafe.Pointer(p), C.float(x), C.float(y), C.float(z)))
}

// GetInterior returns interior
func (p *Player) GetInterior() int {
	return int(C.Player_GetInterior(unsafe.Pointer(p)))
}

// SetPos sets position
func (p *Player) SetPos(x, y, z float32) bool {
	return bool(C.Player_SetPos(unsafe.Pointer(p), C.float(x), C.float(y), C.float(z)))
}

// GetPos returns position
func (p *Player) GetPos() (x, y, z float32, ok bool) {
	var cX, cY, cZ C.float
	ok = bool(C.Player_GetPos(unsafe.Pointer(p), &cX, &cY, &cZ))
	return float32(cX), float32(cY), float32(cZ), ok
}

// GetVirtualWorld returns virtual world
func (p *Player) GetVirtualWorld() int {
	return int(C.Player_GetVirtualWorld(unsafe.Pointer(p)))
}

// IsNPC checks if player is NPC
func (p *Player) IsNPC() bool {
	return bool(C.Player_IsNPC(unsafe.Pointer(p)))
}

// IsStreamedIn checks if player is streamed in for other
func (p *Player) IsStreamedIn(other *Player) bool {
	return bool(C.Player_IsStreamedIn(unsafe.Pointer(p), unsafe.Pointer(other)))
}

// PlayGameSound plays game sound
func (p *Player) PlayGameSound(sound int, x, y, z float32) bool {
	return bool(C.Player_PlayGameSound(unsafe.Pointer(p), C.int(sound), C.float(x), C.float(y), C.float(z)))
}

// SpectatePlayer spectates player
func (p *Player) SpectatePlayer(target *Player, mode int) bool {
	return bool(C.Player_SpectatePlayer(unsafe.Pointer(p), unsafe.Pointer(target), C.int(mode)))
}

// SpectateVehicle spectates vehicle
func (p *Player) SpectateVehicle(target unsafe.Pointer, mode int) bool {
	return bool(C.Player_SpectateVehicle(unsafe.Pointer(p), target, C.int(mode)))
}

// SetVirtualWorld sets virtual world
func (p *Player) SetVirtualWorld(vw int) bool {
	return bool(C.Player_SetVirtualWorld(unsafe.Pointer(p), C.int(vw)))
}

// SetWorldBounds sets world bounds
func (p *Player) SetWorldBounds(xMax, xMin, yMax, yMin float32) bool {
	return bool(C.Player_SetWorldBounds(unsafe.Pointer(p), C.float(xMax), C.float(xMin), C.float(yMax), C.float(yMin)))
}

// ClearWorldBounds clears world bounds
func (p *Player) ClearWorldBounds() bool {
	return bool(C.Player_ClearWorldBounds(unsafe.Pointer(p)))
}

// GetWorldBounds returns world bounds
func (p *Player) GetWorldBounds() (xmax, xmin, ymax, ymin float32, ok bool) {
	var cXMax, cXMin, cYMax, cYMin C.float
	ok = bool(C.Player_GetWorldBounds(unsafe.Pointer(p), &cXMax, &cXMin, &cYMax, &cYMin))
	return float32(cXMax), float32(cXMin), float32(cYMax), float32(cYMin), ok
}

// ClearAnimations clears animations
func (p *Player) ClearAnimations(syncType int) bool {
	return bool(C.Player_ClearAnimations(unsafe.Pointer(p), C.int(syncType)))
}

// GetLastShotVectors returns last shot vectors
func (p *Player) GetLastShotVectors() (originX, originY, originZ, hitX, hitY, hitZ float32, ok bool) {
	var cOriginX, cOriginY, cOriginZ, cHitX, cHitY, cHitZ C.float
	ok = bool(C.Player_GetLastShotVectors(unsafe.Pointer(p), &cOriginX, &cOriginY, &cOriginZ, &cHitX, &cHitY, &cHitZ))
	return float32(cOriginX), float32(cOriginY), float32(cOriginZ), float32(cHitX), float32(cHitY), float32(cHitZ), ok
}

// GetCameraTargetPlayer returns camera target player
func (p *Player) GetCameraTargetPlayer() *Player {
	ptr := C.Player_GetCameraTargetPlayer(unsafe.Pointer(p))
	return (*Player)(ptr)
}

// GetCameraTargetActor returns camera target actor
func (p *Player) GetCameraTargetActor() unsafe.Pointer {
	return unsafe.Pointer(C.Player_GetCameraTargetActor(unsafe.Pointer(p)))
}

// GetCameraTargetObject returns camera target object
func (p *Player) GetCameraTargetObject() unsafe.Pointer {
	return unsafe.Pointer(C.Player_GetCameraTargetObject(unsafe.Pointer(p)))
}

// GetCameraTargetVehicle returns camera target vehicle
func (p *Player) GetCameraTargetVehicle() unsafe.Pointer {
	return unsafe.Pointer(C.Player_GetCameraTargetVehicle(unsafe.Pointer(p)))
}

// PutInVehicle puts player in vehicle
func (p *Player) PutInVehicle(vehicle unsafe.Pointer, seat int) bool {
	return bool(C.Player_PutInVehicle(unsafe.Pointer(p), vehicle, C.int(seat)))
}

// RemoveBuilding removes building
func (p *Player) RemoveBuilding(model int, x, y, z, radius float32) bool {
	return bool(C.Player_RemoveBuilding(unsafe.Pointer(p), C.int(model), C.float(x), C.float(y), C.float(z), C.float(radius)))
}

// GetBuildingsRemoved returns buildings removed count
func (p *Player) GetBuildingsRemoved() int {
	return int(C.Player_GetBuildingsRemoved(unsafe.Pointer(p)))
}

// RemoveFromVehicle removes player from vehicle
func (p *Player) RemoveFromVehicle(force bool) bool {
	return bool(C.Player_RemoveFromVehicle(unsafe.Pointer(p), C.bool(force)))
}

// RemoveMapIcon removes map icon
func (p *Player) RemoveMapIcon(icon int) bool {
	return bool(C.Player_RemoveMapIcon(unsafe.Pointer(p), C.int(icon)))
}

// SetMapIcon sets map icon
func (p *Player) SetMapIcon(iconID int, x, y, z float32, iconType int, color uint32, style int) bool {
	return bool(C.Player_SetMapIcon(unsafe.Pointer(p), C.int(iconID), C.float(x), C.float(y), C.float(z), C.int(iconType), C.uint32_t(color), C.int(style)))
}

// ResetWeapons resets weapons
func (p *Player) ResetWeapons() bool {
	return bool(C.Player_ResetWeapons(unsafe.Pointer(p)))
}

// SetAmmo sets ammo
func (p *Player) SetAmmo(id uint8, ammo uint32) bool {
	return bool(C.Player_SetAmmo(unsafe.Pointer(p), C.uint8_t(id), C.uint32_t(ammo)))
}

// SetArmedWeapon sets armed weapon
func (p *Player) SetArmedWeapon(weapon uint8) bool {
	return bool(C.Player_SetArmedWeapon(unsafe.Pointer(p), C.uint8_t(weapon)))
}

// SetChatBubble sets chat bubble
func (p *Player) SetChatBubble(text string, color uint32, drawdistance float32, expiretime int) bool {
	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))
	return bool(C.Player_SetChatBubble(unsafe.Pointer(p), cText, C.uint32_t(color), C.float(drawdistance), C.int(expiretime)))
}

// SetPosFindZ sets position and finds Z
func (p *Player) SetPosFindZ(x, y, z float32) bool {
	return bool(C.Player_SetPosFindZ(unsafe.Pointer(p), C.float(x), C.float(y), C.float(z)))
}

// SetSkillLevel sets skill level
func (p *Player) SetSkillLevel(weapon uint8, level int) bool {
	return bool(C.Player_SetSkillLevel(unsafe.Pointer(p), C.uint8_t(weapon), C.int(level)))
}

// SetSpecialAction sets special action
func (p *Player) SetSpecialAction(action uint32) bool {
	return bool(C.Player_SetSpecialAction(unsafe.Pointer(p), C.uint32_t(action)))
}

// ShowNameTagForPlayer shows name tag for player
func (p *Player) ShowNameTagForPlayer(other *Player, enable bool) bool {
	return bool(C.Player_ShowNameTagForPlayer(unsafe.Pointer(p), unsafe.Pointer(other), C.bool(enable)))
}

// ToggleControllable toggles controllable
func (p *Player) ToggleControllable(enable bool) bool {
	return bool(C.Player_ToggleControllable(unsafe.Pointer(p), C.bool(enable)))
}

// ToggleSpectating toggles spectating
func (p *Player) ToggleSpectating(enable bool) bool {
	return bool(C.Player_ToggleSpectating(unsafe.Pointer(p), C.bool(enable)))
}

// ApplyAnimation applies animation
func (p *Player) ApplyAnimation(animlib, animname string, delta float32, loop, lockX, lockY, freeze bool, time uint32, sync int) bool {
	cAnimlib := C.CString(animlib)
	cAnimname := C.CString(animname)
	defer C.free(unsafe.Pointer(cAnimlib))
	defer C.free(unsafe.Pointer(cAnimname))
	return bool(C.Player_ApplyAnimation(unsafe.Pointer(p), cAnimlib, cAnimname, C.float(delta), C.bool(loop), C.bool(lockX), C.bool(lockY), C.bool(freeze), C.uint32_t(time), C.int(sync)))
}

// EditAttachedObject edits attached object
func (p *Player) EditAttachedObject(index int) bool {
	return bool(C.Player_EditAttachedObject(unsafe.Pointer(p), C.int(index)))
}

// EnableCameraTarget enables camera target
func (p *Player) EnableCameraTarget(enable bool) bool {
	return bool(C.Player_EnableCameraTarget(unsafe.Pointer(p), C.bool(enable)))
}

// EnableStuntBonus enables stunt bonus
func (p *Player) EnableStuntBonus(enable bool) bool {
	return bool(C.Player_EnableStuntBonus(unsafe.Pointer(p), C.bool(enable)))
}

// GetPlayerAmmo returns player ammo
func (p *Player) GetPlayerAmmo() int {
	return int(C.Player_GetPlayerAmmo(unsafe.Pointer(p)))
}

// GetAnimationIndex returns animation index
func (p *Player) GetAnimationIndex() int {
	return int(C.Player_GetAnimationIndex(unsafe.Pointer(p)))
}

// GetFacingAngle returns facing angle
func (p *Player) GetFacingAngle() float32 {
	return float32(C.Player_GetFacingAngle(unsafe.Pointer(p)))
}

// GetIp returns player IP address
func (p *Player) GetIp() string {
	buf := make([]byte, 46)
	var output C.struct_CAPIStringBuffer
	output.data = (*C.char)(unsafe.Pointer(&buf[0]))
	output.capacity = C.uint(len(buf))
	C.Player_GetIp(unsafe.Pointer(p), &output)
	return string(buf[:output.len])
}

// GetSpecialAction returns special action
func (p *Player) GetSpecialAction() int {
	return int(C.Player_GetSpecialAction(unsafe.Pointer(p)))
}

// GetVehicleID returns vehicle ID
func (p *Player) GetVehicleID() int {
	return int(C.Player_GetVehicleID(unsafe.Pointer(p)))
}

// GetVehicleSeat returns vehicle seat
func (p *Player) GetVehicleSeat() int {
	return int(C.Player_GetVehicleSeat(unsafe.Pointer(p)))
}

// GetWeaponData returns weapon data
func (p *Player) GetWeaponData(slot int) (weaponid, ammo int, ok bool) {
	var cWeaponid, cAmmo C.int
	ok = bool(C.Player_GetWeaponData(unsafe.Pointer(p), C.int(slot), &cWeaponid, &cAmmo))
	return int(cWeaponid), int(cAmmo), ok
}

// GetWeaponState returns weapon state
func (p *Player) GetWeaponState() int {
	return int(C.Player_GetWeaponState(unsafe.Pointer(p)))
}

// InterpolateCameraPos interpolates camera position
func (p *Player) InterpolateCameraPos(fromX, fromY, fromZ, toX, toY, toZ float32, time, cut int) bool {
	return bool(C.Player_InterpolateCameraPos(unsafe.Pointer(p), C.float(fromX), C.float(fromY), C.float(fromZ), C.float(toX), C.float(toY), C.float(toZ), C.int(time), C.int(cut)))
}

// InterpolateCameraLookAt interpolates camera look at
func (p *Player) InterpolateCameraLookAt(fromX, fromY, fromZ, toX, toY, toZ float32, time, cut int) bool {
	return bool(C.Player_InterpolateCameraLookAt(unsafe.Pointer(p), C.float(fromX), C.float(fromY), C.float(fromZ), C.float(toX), C.float(toY), C.float(toZ), C.int(time), C.int(cut)))
}

// IsPlayerAttachedObjectSlotUsed checks if attached object slot is used
func (p *Player) IsPlayerAttachedObjectSlotUsed(index int) bool {
	return bool(C.Player_IsPlayerAttachedObjectSlotUsed(unsafe.Pointer(p), C.int(index)))
}

// AttachCameraToObject attaches camera to object
func (p *Player) AttachCameraToObject(object unsafe.Pointer) bool {
	return bool(C.Player_AttachCameraToObject(unsafe.Pointer(p), object))
}

// AttachCameraToPlayerObject attaches camera to player object
func (p *Player) AttachCameraToPlayerObject(object unsafe.Pointer) bool {
	return bool(C.Player_AttachCameraToPlayerObject(unsafe.Pointer(p), object))
}

// GetCameraAspectRatio returns camera aspect ratio
func (p *Player) GetCameraAspectRatio() float32 {
	return float32(C.Player_GetCameraAspectRatio(unsafe.Pointer(p)))
}

// GetCameraFrontVector returns camera front vector
func (p *Player) GetCameraFrontVector() (x, y, z float32, ok bool) {
	var cX, cY, cZ C.float
	ok = bool(C.Player_GetCameraFrontVector(unsafe.Pointer(p), &cX, &cY, &cZ))
	return float32(cX), float32(cY), float32(cZ), ok
}

// GetCameraMode returns camera mode
func (p *Player) GetCameraMode() int {
	return int(C.Player_GetCameraMode(unsafe.Pointer(p)))
}

// GetKeys returns keys
func (p *Player) GetKeys() (keys, updown, leftright int, ok bool) {
	var cKeys, cUpdown, cLeftright C.int
	ok = bool(C.Player_GetKeys(unsafe.Pointer(p), &cKeys, &cUpdown, &cLeftright))
	return int(cKeys), int(cUpdown), int(cLeftright), ok
}

// GetSurfingVehicle returns surfing vehicle
func (p *Player) GetSurfingVehicle() unsafe.Pointer {
	return unsafe.Pointer(C.Player_GetSurfingVehicle(unsafe.Pointer(p)))
}

// GetSurfingObject returns surfing object
func (p *Player) GetSurfingObject() unsafe.Pointer {
	return unsafe.Pointer(C.Player_GetSurfingObject(unsafe.Pointer(p)))
}

// GetTargetPlayer returns target player
func (p *Player) GetTargetPlayer() *Player {
	ptr := C.Player_GetTargetPlayer(unsafe.Pointer(p))
	return (*Player)(ptr)
}

// GetTargetActor returns target actor
func (p *Player) GetTargetActor() unsafe.Pointer {
	return unsafe.Pointer(C.Player_GetTargetActor(unsafe.Pointer(p)))
}

// IsInVehicle checks if player is in vehicle
func (p *Player) IsInVehicle(targetVehicle unsafe.Pointer) bool {
	return bool(C.Player_IsInVehicle(unsafe.Pointer(p), targetVehicle))
}

// IsInAnyVehicle checks if player is in any vehicle
func (p *Player) IsInAnyVehicle() bool {
	return bool(C.Player_IsInAnyVehicle(unsafe.Pointer(p)))
}

// IsInRangeOfPoint checks if player is in range of point
func (p *Player) IsInRangeOfPoint(rangeVal, x, y, z float32) bool {
	return bool(C.Player_IsInRangeOfPoint(unsafe.Pointer(p), C.float(rangeVal), C.float(x), C.float(y), C.float(z)))
}

// PlayCrimeReport plays crime report
func (p *Player) PlayCrimeReport(suspect *Player, crime int) bool {
	return bool(C.Player_PlayCrimeReport(unsafe.Pointer(p), unsafe.Pointer(suspect), C.int(crime)))
}

// RemoveAttachedObject removes attached object
func (p *Player) RemoveAttachedObject(index int) bool {
	return bool(C.Player_RemoveAttachedObject(unsafe.Pointer(p), C.int(index)))
}

// SetAttachedObject sets attached object
func (p *Player) SetAttachedObject(index, modelid, bone int, offsetX, offsetY, offsetZ, rotationX, rotationY, rotationZ, scaleX, scaleY, scaleZ float32, materialcolor1, materialcolor2 uint32) bool {
	return bool(C.Player_SetAttachedObject(unsafe.Pointer(p), C.int(index), C.int(modelid), C.int(bone), C.float(offsetX), C.float(offsetY), C.float(offsetZ), C.float(rotationX), C.float(rotationY), C.float(rotationZ), C.float(scaleX), C.float(scaleY), C.float(scaleZ), C.uint32_t(materialcolor1), C.uint32_t(materialcolor2)))
}

// GetAttachedObject gets attached object
func (p *Player) GetAttachedObject(index int) (modelid, bone int, offsetX, offsetY, offsetZ, rotationX, rotationY, rotationZ, scaleX, scaleY, scaleZ float32, materialcolor1, materialcolor2 int, ok bool) {
	var cModelid, cBone, cMaterialcolor1, cMaterialcolor2 C.int
	var cOffsetX, cOffsetY, cOffsetZ, cRotationX, cRotationY, cRotationZ, cScaleX, cScaleY, cScaleZ C.float
	ok = bool(C.Player_GetAttachedObject(unsafe.Pointer(p), C.int(index), &cModelid, &cBone, &cOffsetX, &cOffsetY, &cOffsetZ, &cRotationX, &cRotationY, &cRotationZ, &cScaleX, &cScaleY, &cScaleZ, &cMaterialcolor1, &cMaterialcolor2))
	return int(cModelid), int(cBone), float32(cOffsetX), float32(cOffsetY), float32(cOffsetZ), float32(cRotationX), float32(cRotationY), float32(cRotationZ), float32(cScaleX), float32(cScaleY), float32(cScaleZ), int(cMaterialcolor1), int(cMaterialcolor2), ok
}

// SetFacingAngle sets facing angle
func (p *Player) SetFacingAngle(angle float32) bool {
	return bool(C.Player_SetFacingAngle(unsafe.Pointer(p), C.float(angle)))
}

// SetMarkerForPlayer sets marker for player
func (p *Player) SetMarkerForPlayer(other *Player, color uint32) bool {
	return bool(C.Player_SetMarkerForPlayer(unsafe.Pointer(p), unsafe.Pointer(other), C.uint32_t(color)))
}

// GetMarkerForPlayer gets marker for player
func (p *Player) GetMarkerForPlayer(other *Player) uint32 {
	return uint32(C.Player_GetMarkerForPlayer(unsafe.Pointer(p), unsafe.Pointer(other)))
}

// AllowTeleport allows teleport
func (p *Player) AllowTeleport(allow bool) bool {
	return bool(C.Player_AllowTeleport(unsafe.Pointer(p), C.bool(allow)))
}

// IsTeleportAllowed checks if teleport is allowed
func (p *Player) IsTeleportAllowed() bool {
	return bool(C.Player_IsTeleportAllowed(unsafe.Pointer(p)))
}

// DisableRemoteVehicleCollisions disables remote vehicle collisions
func (p *Player) DisableRemoteVehicleCollisions(disable bool) bool {
	return bool(C.Player_DisableRemoteVehicleCollisions(unsafe.Pointer(p), C.bool(disable)))
}

// GetCameraZoom returns camera zoom
func (p *Player) GetCameraZoom() float32 {
	return float32(C.Player_GetCameraZoom(unsafe.Pointer(p)))
}

// SelectTextDraw selects text draw
func (p *Player) SelectTextDraw(hoverColour uint32) bool {
	return bool(C.Player_SelectTextDraw(unsafe.Pointer(p), C.uint32_t(hoverColour)))
}

// CancelSelectTextDraw cancels select text draw
func (p *Player) CancelSelectTextDraw() bool {
	return bool(C.Player_CancelSelectTextDraw(unsafe.Pointer(p)))
}

// SendClientCheck sends client check
func (p *Player) SendClientCheck(actionType, address, offset, count int) bool {
	return bool(C.Player_SendClientCheck(unsafe.Pointer(p), C.int(actionType), C.int(address), C.int(offset), C.int(count)))
}

// Spawn spawns player
func (p *Player) Spawn() bool {
	return bool(C.Player_Spawn(unsafe.Pointer(p)))
}

// IsAdmin checks if player is admin
func (p *Player) IsAdmin() bool {
	return bool(C.Player_IsAdmin(unsafe.Pointer(p)))
}

// Kick kicks player
func (p *Player) Kick() bool {
	return bool(C.Player_Kick(unsafe.Pointer(p)))
}

// ShowGameText shows game text
func (p *Player) ShowGameText(text string, time, style int) bool {
	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))
	return bool(C.Player_ShowGameText(unsafe.Pointer(p), cText, C.int(time), C.int(style)))
}

// HideGameText hides game text
func (p *Player) HideGameText(style int) bool {
	return bool(C.Player_HideGameText(unsafe.Pointer(p), C.int(style)))
}

// HasGameText checks if player has game text
func (p *Player) HasGameText(style int) bool {
	return bool(C.Player_HasGameText(unsafe.Pointer(p), C.int(style)))
}

// GetGameText returns game text data
func (p *Player) GetGameText(style int) (message string, time, remaining int, ok bool) {
	var cMessage C.struct_CAPIStringView
	var cTime, cRemaining C.int
	ok = bool(C.Player_GetGameText(unsafe.Pointer(p), C.int(style), &cMessage, &cTime, &cRemaining))
	return C.GoStringN(cMessage.data, C.int(cMessage.len)), int(cTime), int(cRemaining), ok
}

// Ban bans player
func (p *Player) Ban() bool {
	return bool(C.Player_Ban(unsafe.Pointer(p)))
}

// BanEx bans player with reason
func (p *Player) BanEx(reason string) bool {
	cReason := C.CString(reason)
	defer C.free(unsafe.Pointer(cReason))
	return bool(C.Player_BanEx(unsafe.Pointer(p), cReason))
}

// SendDeathMessage sends death message
func (p *Player) SendDeathMessage(killer, killee *Player, weapon int) bool {
	return bool(C.Player_SendDeathMessage(unsafe.Pointer(p), unsafe.Pointer(killer), unsafe.Pointer(killee), C.int(weapon)))
}

// SendMessageToPlayer sends message to player
func (p *Player) SendMessageToPlayer(sender *Player, message string) bool {
	cMessage := C.CString(message)
	defer C.free(unsafe.Pointer(cMessage))
	return bool(C.Player_SendMessageToPlayer(unsafe.Pointer(p), unsafe.Pointer(sender), cMessage))
}

// GetVersion returns player client version
func (p *Player) GetVersion() string {
	var version C.struct_CAPIStringView
	C.Player_GetVersion(unsafe.Pointer(p), &version)
	return C.GoStringN(version.data, C.int(version.len))
}

// GPCI returns player GPCI (serial)
func (p *Player) GPCI() (string, bool) {
	var gpci C.struct_CAPIStringView
	ok := bool(C.Player_GPCI(unsafe.Pointer(p), &gpci))
	return C.GoStringN(gpci.data, C.int(gpci.len)), ok
}

// GetSkillLevel returns skill level
func (p *Player) GetSkillLevel(skill int) int {
	return int(C.Player_GetSkillLevel(unsafe.Pointer(p), C.int(skill)))
}

// GetZAim returns Z aim
func (p *Player) GetZAim() float32 {
	return float32(C.Player_GetZAim(unsafe.Pointer(p)))
}

// GetSurfingOffsets returns surfing offsets
func (p *Player) GetSurfingOffsets() (offsetX, offsetY, offsetZ float32, ok bool) {
	var cOffsetX, cOffsetY, cOffsetZ C.float
	ok = bool(C.Player_GetSurfingOffsets(unsafe.Pointer(p), &cOffsetX, &cOffsetY, &cOffsetZ))
	return float32(cOffsetX), float32(cOffsetY), float32(cOffsetZ), ok
}

// GetRotationQuat returns rotation quaternion
func (p *Player) GetRotationQuat() (x, y, z, w float32, ok bool) {
	var cX, cY, cZ, cW C.float
	ok = bool(C.Player_GetRotationQuat(unsafe.Pointer(p), &cX, &cY, &cZ, &cW))
	return float32(cX), float32(cY), float32(cZ), float32(cW), ok
}

// GetPlayerSpectateID returns player spectate ID
func (p *Player) GetPlayerSpectateID() int {
	return int(C.Player_GetPlayerSpectateID(unsafe.Pointer(p)))
}

// GetSpectateType returns spectate type
func (p *Player) GetSpectateType() int {
	return int(C.Player_GetSpectateType(unsafe.Pointer(p)))
}

// GetRawIp returns raw IP
func (p *Player) GetRawIp() uint32 {
	return uint32(C.Player_GetRawIp(unsafe.Pointer(p)))
}

// SetGravity sets gravity
func (p *Player) SetGravity(gravity float32) bool {
	return bool(C.Player_SetGravity(unsafe.Pointer(p), C.float(gravity)))
}

// GetGravity returns gravity
func (p *Player) GetGravity() float32 {
	return float32(C.Player_GetGravity(unsafe.Pointer(p)))
}

// SetAdmin sets admin
func (p *Player) SetAdmin(set bool) bool {
	return bool(C.Player_SetAdmin(unsafe.Pointer(p), C.bool(set)))
}

// IsSpawned checks if player is spawned
func (p *Player) IsSpawned() bool {
	return bool(C.Player_IsSpawned(unsafe.Pointer(p)))
}

// IsControllable checks if player is controllable
func (p *Player) IsControllable() bool {
	return bool(C.Player_IsControllable(unsafe.Pointer(p)))
}

// IsCameraTargetEnabled checks if camera target is enabled
func (p *Player) IsCameraTargetEnabled() bool {
	return bool(C.Player_IsCameraTargetEnabled(unsafe.Pointer(p)))
}

// ToggleGhostMode toggles ghost mode
func (p *Player) ToggleGhostMode(toggle bool) bool {
	return bool(C.Player_ToggleGhostMode(unsafe.Pointer(p), C.bool(toggle)))
}

// GetGhostMode returns ghost mode
func (p *Player) GetGhostMode() bool {
	return bool(C.Player_GetGhostMode(unsafe.Pointer(p)))
}

// AllowWeapons allows weapons
func (p *Player) AllowWeapons(allow bool) bool {
	return bool(C.Player_AllowWeapons(unsafe.Pointer(p), C.bool(allow)))
}

// AreWeaponsAllowed checks if weapons are allowed
func (p *Player) AreWeaponsAllowed() bool {
	return bool(C.Player_AreWeaponsAllowed(unsafe.Pointer(p)))
}

// IsPlayerUsingOfficialClient checks if player is using official client
func (p *Player) IsPlayerUsingOfficialClient() bool {
	return bool(C.Player_IsPlayerUsingOfficialClient(unsafe.Pointer(p)))
}

// GetAnimationFlags returns animation flags
func (p *Player) GetAnimationFlags() int {
	return int(C.Player_GetAnimationFlags(unsafe.Pointer(p)))
}

// IsInDriveByMode checks if player is in drive-by mode
func (p *Player) IsInDriveByMode() bool {
	return bool(C.Player_IsInDriveByMode(unsafe.Pointer(p)))
}

// IsCuffed checks if player is cuffed
func (p *Player) IsCuffed() bool {
	return bool(C.Player_IsCuffed(unsafe.Pointer(p)))
}

// IsUsingOmp checks if player is using OMP
func (p *Player) IsUsingOmp() bool {
	return bool(C.Player_IsUsingOmp(unsafe.Pointer(p)))
}

// IsInModShop checks if player is in mod shop
func (p *Player) IsInModShop() bool {
	return bool(C.Player_IsInModShop(unsafe.Pointer(p)))
}

// GetSirenState returns siren state
func (p *Player) GetSirenState() int {
	return int(C.Player_GetSirenState(unsafe.Pointer(p)))
}

// GetLandingGearState returns landing gear state
func (p *Player) GetLandingGearState() int {
	return int(C.Player_GetLandingGearState(unsafe.Pointer(p)))
}

// GetHydraReactorAngle returns hydra reactor angle
func (p *Player) GetHydraReactorAngle() uint32 {
	return uint32(C.Player_GetHydraReactorAngle(unsafe.Pointer(p)))
}

// GetTrainSpeed returns train speed
func (p *Player) GetTrainSpeed() float32 {
	return float32(C.Player_GetTrainSpeed(unsafe.Pointer(p)))
}
