#ifndef PLAYER_H_
#define PLAYER_H_

#include "bridge.h"

// Player_SetSpawnInfo (bool)
DECL_FN(bool, Player_SetSpawnInfo, Player.SetSpawnInfo, (void *player, uint8_t team, int skin, float x, float y, float z, float angle, uint8_t weapon1, uint32_t ammo1, uint8_t weapon2, uint32_t ammo2, uint8_t weapon3, uint32_t ammo3), (player, team, skin, x, y, z, angle, weapon1, ammo1, weapon2, ammo2, weapon3, ammo3))

// Player_GetSpawnInfo (bool)
DECL_FN(bool, Player_GetSpawnInfo, Player.GetSpawnInfo, (void *player, uint8_t *team, int *skin, float *x, float *y, float *z, float *angle, uint8_t *weapon1, uint32_t *ammo1, uint8_t *weapon2, uint32_t *ammo2, uint8_t *weapon3, uint32_t *ammo3), (player, team, skin, x, y, z, angle, weapon1, ammo1, weapon2, ammo2, weapon3, ammo3))

// Player_GetNetworkStats (int)
DECL_FN(int, Player_GetNetworkStats, Player.GetNetworkStats, (void *player, struct CAPIStringBuffer *output), (player, output))

// Player_NetStatsBytesReceived (int)
DECL_FN(int, Player_NetStatsBytesReceived, Player.NetStatsBytesReceived, (void *player), (player))

// Player_NetStatsBytesSent (int)
DECL_FN(int, Player_NetStatsBytesSent, Player.NetStatsBytesSent, (void *player), (player))

// Player_NetStatsConnectionStatus (int)
DECL_FN(int, Player_NetStatsConnectionStatus, Player.NetStatsConnectionStatus, (void *player), (player))

// Player_NetStatsGetConnectedTime (int)
DECL_FN(int, Player_NetStatsGetConnectedTime, Player.NetStatsGetConnectedTime, (void *player), (player))

// Player_NetStatsGetIpPort (bool)
DECL_FN(bool, Player_NetStatsGetIpPort, Player.NetStatsGetIpPort, (void *player, struct CAPIStringBuffer *output), (player, output))

// Player_NetStatsMessagesReceived (int)
DECL_FN(int, Player_NetStatsMessagesReceived, Player.NetStatsMessagesReceived, (void *player), (player))

// Player_NetStatsMessagesRecvPerSecond (int)
DECL_FN(int, Player_NetStatsMessagesRecvPerSecond, Player.NetStatsMessagesRecvPerSecond, (void *player), (player))

// Player_NetStatsMessagesSent (int)
DECL_FN(int, Player_NetStatsMessagesSent, Player.NetStatsMessagesSent, (void *player), (player))

// Player_NetStatsPacketLossPercent (float)
DECL_FN(float, Player_NetStatsPacketLossPercent, Player.NetStatsPacketLossPercent, (void *player), (player))

// Player_GetCustomSkin (int)
DECL_FN(int, Player_GetCustomSkin, Player.GetCustomSkin, (void *player), (player))

// Player_GetDialog (int)
DECL_FN(int, Player_GetDialog, Player.GetDialog, (void *player), (player))

// Player_GetDialogData (bool)
DECL_FN(bool, Player_GetDialogData, Player.GetDialogData, (void *player, int *dialogid, int *style, struct CAPIStringView *title, struct CAPIStringView *body, struct CAPIStringView *button1, struct CAPIStringView *button2), (player, dialogid, style, title, body, button1, button2))

// Player_GetMenu (void*)
DECL_FN(void *, Player_GetMenu, Player.GetMenu, (void *player), (player))

// Player_GetSurfingPlayerObject (void*)
DECL_FN(void *, Player_GetSurfingPlayerObject, Player.GetSurfingPlayerObject, (void *player), (player))

// Player_GetCameraTargetPlayerObject (void*)
DECL_FN(void *, Player_GetCameraTargetPlayerObject, Player.GetCameraTargetPlayerObject, (void *player), (player))

// Player_FromID (void*)
DECL_FN(void *, Player_FromID, Player.FromID, (int playerid), (playerid))

// Player_GetID (int)
DECL_FN(int, Player_GetID, Player.GetID, (void *player), (player))

// Player_SendClientMessage (bool)
DECL_FN(bool, Player_SendClientMessage, Player.SendClientMessage, (void *player, uint32_t color, const char *text), (player, color, text))

// Player_SetCameraPos (bool)
DECL_FN(bool, Player_SetCameraPos, Player.SetCameraPos, (void *player, float x, float y, float z), (player, x, y, z))

// Player_SetDrunkLevel (bool)
DECL_FN(bool, Player_SetDrunkLevel, Player.SetDrunkLevel, (void *player, int level), (player, level))

// Player_SetInterior (bool)
DECL_FN(bool, Player_SetInterior, Player.SetInterior, (void *player, int interior), (player, interior))

// Player_SetWantedLevel (bool)
DECL_FN(bool, Player_SetWantedLevel, Player.SetWantedLevel, (void *player, int level), (player, level))

// Player_SetWeather (bool)
DECL_FN(bool, Player_SetWeather, Player.SetWeather, (void *player, int weather), (player, weather))

// Player_GetWeather (int)
DECL_FN(int, Player_GetWeather, Player.GetWeather, (void *player), (player))

// Player_SetSkin (bool)
DECL_FN(bool, Player_SetSkin, Player.SetSkin, (void *player, int skin), (player, skin))

// Player_SetShopName (bool)
DECL_FN(bool, Player_SetShopName, Player.SetShopName, (void *player, const char *name), (player, name))

// Player_GiveMoney (bool)
DECL_FN(bool, Player_GiveMoney, Player.GiveMoney, (void *player, int amount), (player, amount))

// Player_SetCameraLookAt (bool)
DECL_FN(bool, Player_SetCameraLookAt, Player.SetCameraLookAt, (void *player, float x, float y, float z, int cutType), (player, x, y, z, cutType))

// Player_SetCameraBehind (bool)
DECL_FN(bool, Player_SetCameraBehind, Player.SetCameraBehind, (void *player), (player))

// Player_CreateExplosion (bool)
DECL_FN(bool, Player_CreateExplosion, Player.CreateExplosion, (void *player, float x, float y, float z, int type, float radius), (player, x, y, z, type, radius))

// Player_PlayAudioStream (bool)
DECL_FN(bool, Player_PlayAudioStream, Player.PlayAudioStream, (void *player, const char *url, float x, float y, float z, float distance, bool usePos), (player, url, x, y, z, distance, usePos))

// Player_StopAudioStream (bool)
DECL_FN(bool, Player_StopAudioStream, Player.StopAudioStream, (void *player), (player))

// Player_ToggleWidescreen (bool)
DECL_FN(bool, Player_ToggleWidescreen, Player.ToggleWidescreen, (void *player, bool enable), (player, enable))

// Player_IsWidescreenToggled (bool)
DECL_FN(bool, Player_IsWidescreenToggled, Player.IsWidescreenToggled, (void *player), (player))

// Player_SetHealth (bool)
DECL_FN(bool, Player_SetHealth, Player.SetHealth, (void *player, float health), (player, health))

// Player_GetHealth (float)
DECL_FN(float, Player_GetHealth, Player.GetHealth, (void *player), (player))

// Player_SetArmor (bool)
DECL_FN(bool, Player_SetArmor, Player.SetArmor, (void *player, float armor), (player, armor))

// Player_GetArmor (float)
DECL_FN(float, Player_GetArmor, Player.GetArmor, (void *player), (player))

// Player_SetTeam (bool)
DECL_FN(bool, Player_SetTeam, Player.SetTeam, (void *player, int team), (player, team))

// Player_GetTeam (int)
DECL_FN(int, Player_GetTeam, Player.GetTeam, (void *player), (player))

// Player_SetScore (bool)
DECL_FN(bool, Player_SetScore, Player.SetScore, (void *player, int score), (player, score))

// Player_GetScore (int)
DECL_FN(int, Player_GetScore, Player.GetScore, (void *player), (player))

// Player_GetSkin (int)
DECL_FN(int, Player_GetSkin, Player.GetSkin, (void *player), (player))

// Player_SetColor (bool)
DECL_FN(bool, Player_SetColor, Player.SetColor, (void *player, uint32_t color), (player, color))

// Player_GetColor (uint32_t)
DECL_FN(uint32_t, Player_GetColor, Player.GetColor, (void *player), (player))

// Player_GetDefaultColor (uint32_t)
DECL_FN(uint32_t, Player_GetDefaultColor, Player.GetDefaultColor, (void *player), (player))

// Player_GetDrunkLevel (int)
DECL_FN(int, Player_GetDrunkLevel, Player.GetDrunkLevel, (void *player), (player))

// Player_GiveWeapon (bool)
DECL_FN(bool, Player_GiveWeapon, Player.GiveWeapon, (void *player, int weapon, int ammo), (player, weapon, ammo))

// Player_RemoveWeapon (bool)
DECL_FN(bool, Player_RemoveWeapon, Player.RemoveWeapon, (void *player, int weapon), (player, weapon))

// Player_GetMoney (int)
DECL_FN(int, Player_GetMoney, Player.GetMoney, (void *player), (player))

// Player_ResetMoney (bool)
DECL_FN(bool, Player_ResetMoney, Player.ResetMoney, (void *player), (player))

// Player_SetName (int)
DECL_FN(int, Player_SetName, Player.SetName, (void *player, const char *name), (player, name))

// Player_GetName (int)
DECL_FN(int, Player_GetName, Player.GetName, (void *player, struct CAPIStringView *name), (player, name))

// Player_GetState (int)
DECL_FN(int, Player_GetState, Player.GetState, (void *player), (player))

// Player_GetPing (int)
DECL_FN(int, Player_GetPing, Player.GetPing, (void *player), (player))

// Player_GetWeapon (int)
DECL_FN(int, Player_GetWeapon, Player.GetWeapon, (void *player), (player))

// Player_SetTime (bool)
DECL_FN(bool, Player_SetTime, Player.SetTime, (void *player, int hour, int minute), (player, hour, minute))

// Player_GetTime (bool)
DECL_FN(bool, Player_GetTime, Player.GetTime, (void *player, int *hour, int *minute), (player, hour, minute))

// Player_ToggleClock (bool)
DECL_FN(bool, Player_ToggleClock, Player.ToggleClock, (void *player, bool enable), (player, enable))

// Player_HasClock (bool)
DECL_FN(bool, Player_HasClock, Player.HasClock, (void *player), (player))

// Player_ForceClassSelection (bool)
DECL_FN(bool, Player_ForceClassSelection, Player.ForceClassSelection, (void *player), (player))

// Player_GetWantedLevel (int)
DECL_FN(int, Player_GetWantedLevel, Player.GetWantedLevel, (void *player), (player))

// Player_SetFightingStyle (bool)
DECL_FN(bool, Player_SetFightingStyle, Player.SetFightingStyle, (void *player, int style), (player, style))

// Player_GetFightingStyle (int)
DECL_FN(int, Player_GetFightingStyle, Player.GetFightingStyle, (void *player), (player))

// Player_SetVelocity (bool)
DECL_FN(bool, Player_SetVelocity, Player.SetVelocity, (void *player, float x, float y, float z), (player, x, y, z))

// Player_GetVelocity (bool)
DECL_FN(bool, Player_GetVelocity, Player.GetVelocity, (void *player, float *x, float *y, float *z), (player, x, y, z))

// Player_GetCameraPos (bool)
DECL_FN(bool, Player_GetCameraPos, Player.GetCameraPos, (void *player, float *x, float *y, float *z), (player, x, y, z))

// Player_GetDistanceFromPoint (float)
DECL_FN(float, Player_GetDistanceFromPoint, Player.GetDistanceFromPoint, (void *player, float x, float y, float z), (player, x, y, z))

// Player_GetInterior (int)
DECL_FN(int, Player_GetInterior, Player.GetInterior, (void *player), (player))

// Player_SetPos (bool)
DECL_FN(bool, Player_SetPos, Player.SetPos, (void *player, float x, float y, float z), (player, x, y, z))

// Player_GetPos (bool)
DECL_FN(bool, Player_GetPos, Player.GetPos, (void *player, float *x, float *y, float *z), (player, x, y, z))

// Player_GetVirtualWorld (int)
DECL_FN(int, Player_GetVirtualWorld, Player.GetVirtualWorld, (void *player), (player))

// Player_IsNPC (bool)
DECL_FN(bool, Player_IsNPC, Player.IsNPC, (void *player), (player))

// Player_IsStreamedIn (bool)
DECL_FN(bool, Player_IsStreamedIn, Player.IsStreamedIn, (void *player, void *other), (player, other))

// Player_PlayGameSound (bool)
DECL_FN(bool, Player_PlayGameSound, Player.PlayGameSound, (void *player, int sound, float x, float y, float z), (player, sound, x, y, z))

// Player_SpectatePlayer (bool)
DECL_FN(bool, Player_SpectatePlayer, Player.SpectatePlayer, (void *player, void *target, int mode), (player, target, mode))

// Player_SpectateVehicle (bool)
DECL_FN(bool, Player_SpectateVehicle, Player.SpectateVehicle, (void *player, void *target, int mode), (player, target, mode))

// Player_SetVirtualWorld (bool)
DECL_FN(bool, Player_SetVirtualWorld, Player.SetVirtualWorld, (void *player, int vw), (player, vw))

// Player_SetWorldBounds (bool)
DECL_FN(bool, Player_SetWorldBounds, Player.SetWorldBounds, (void *player, float xMax, float xMin, float yMax, float yMin), (player, xMax, xMin, yMax, yMin))

// Player_ClearWorldBounds (bool)
DECL_FN(bool, Player_ClearWorldBounds, Player.ClearWorldBounds, (void *player), (player))

// Player_GetWorldBounds (bool)
DECL_FN(bool, Player_GetWorldBounds, Player.GetWorldBounds, (void *player, float *xmax, float *xmin, float *ymax, float *ymin), (player, xmax, xmin, ymax, ymin))

// Player_ClearAnimations (bool)
DECL_FN(bool, Player_ClearAnimations, Player.ClearAnimations, (void *player, int syncType), (player, syncType))

// Player_GetLastShotVectors (bool)
DECL_FN(bool, Player_GetLastShotVectors, Player.GetLastShotVectors, (void *player, float *origin_x, float *origin_y, float *origin_z, float *hit_x, float *hit_y, float *hit_z), (player, origin_x, origin_y, origin_z, hit_x, hit_y, hit_z))

// Player_GetCameraTargetPlayer (void*)
DECL_FN(void *, Player_GetCameraTargetPlayer, Player.GetCameraTargetPlayer, (void *player), (player))

// Player_GetCameraTargetActor (void*)
DECL_FN(void *, Player_GetCameraTargetActor, Player.GetCameraTargetActor, (void *player), (player))

// Player_GetCameraTargetObject (void*)
DECL_FN(void *, Player_GetCameraTargetObject, Player.GetCameraTargetObject, (void *player), (player))

// Player_GetCameraTargetVehicle (void*)
DECL_FN(void *, Player_GetCameraTargetVehicle, Player.GetCameraTargetVehicle, (void *player), (player))

// Player_PutInVehicle (bool)
DECL_FN(bool, Player_PutInVehicle, Player.PutInVehicle, (void *player, void *vehicle, int seat), (player, vehicle, seat))

// Player_RemoveBuilding (bool)
DECL_FN(bool, Player_RemoveBuilding, Player.RemoveBuilding, (void *player, int model, float x, float y, float z, float radius), (player, model, x, y, z, radius))

// Player_GetBuildingsRemoved (int)
DECL_FN(int, Player_GetBuildingsRemoved, Player.GetBuildingsRemoved, (void *player), (player))

// Player_RemoveFromVehicle (bool)
DECL_FN(bool, Player_RemoveFromVehicle, Player.RemoveFromVehicle, (void *player, bool force), (player, force))

// Player_RemoveMapIcon (bool)
DECL_FN(bool, Player_RemoveMapIcon, Player.RemoveMapIcon, (void *player, int icon), (player, icon))

// Player_SetMapIcon (bool)
DECL_FN(bool, Player_SetMapIcon, Player.SetMapIcon, (void *player, int iconID, float x, float y, float z, int type, uint32_t color, int style), (player, iconID, x, y, z, type, color, style))

// Player_ResetWeapons (bool)
DECL_FN(bool, Player_ResetWeapons, Player.ResetWeapons, (void *player), (player))

// Player_SetAmmo (bool)
DECL_FN(bool, Player_SetAmmo, Player.SetAmmo, (void *player, uint8_t id, uint32_t ammo), (player, id, ammo))

// Player_SetArmedWeapon (bool)
DECL_FN(bool, Player_SetArmedWeapon, Player.SetArmedWeapon, (void *player, uint8_t weapon), (player, weapon))

// Player_SetChatBubble (bool)
DECL_FN(bool, Player_SetChatBubble, Player.SetChatBubble, (void *player, const char *text, uint32_t color, float drawdistance, int expiretime), (player, text, color, drawdistance, expiretime))

// Player_SetPosFindZ (bool)
DECL_FN(bool, Player_SetPosFindZ, Player.SetPosFindZ, (void *player, float x, float y, float z), (player, x, y, z))

// Player_SetSkillLevel (bool)
DECL_FN(bool, Player_SetSkillLevel, Player.SetSkillLevel, (void *player, uint8_t weapon, int level), (player, weapon, level))

// Player_SetSpecialAction (bool)
DECL_FN(bool, Player_SetSpecialAction, Player.SetSpecialAction, (void *player, uint32_t action), (player, action))

// Player_ShowNameTagForPlayer (bool)
DECL_FN(bool, Player_ShowNameTagForPlayer, Player.ShowNameTagForPlayer, (void *player, void *other, bool enable), (player, other, enable))

// Player_ToggleControllable (bool)
DECL_FN(bool, Player_ToggleControllable, Player.ToggleControllable, (void *player, bool enable), (player, enable))

// Player_ToggleSpectating (bool)
DECL_FN(bool, Player_ToggleSpectating, Player.ToggleSpectating, (void *player, bool enable), (player, enable))

// Player_ApplyAnimation (bool)
DECL_FN(bool, Player_ApplyAnimation, Player.ApplyAnimation, (void *player, const char *animlib, const char *animname, float delta, bool loop, bool lockX, bool lockY, bool freeze, uint32_t time, int sync), (player, animlib, animname, delta, loop, lockX, lockY, freeze, time, sync))

// Player_GetAnimationName (bool)
DECL_FN(bool, Player_GetAnimationName, Player.GetAnimationName, (int index, struct CAPIStringView *lib, struct CAPIStringView *name), (index, lib, name))

// Player_EditAttachedObject (bool)
DECL_FN(bool, Player_EditAttachedObject, Player.EditAttachedObject, (void *player, int index), (player, index))

// Player_EnableCameraTarget (bool)
DECL_FN(bool, Player_EnableCameraTarget, Player.EnableCameraTarget, (void *player, bool enable), (player, enable))

// Player_EnableStuntBonus (bool)
DECL_FN(bool, Player_EnableStuntBonus, Player.EnableStuntBonus, (void *player, bool enable), (player, enable))

// Player_GetPlayerAmmo (int)
DECL_FN(int, Player_GetPlayerAmmo, Player.GetPlayerAmmo, (void *player), (player))

// Player_GetAnimationIndex (int)
DECL_FN(int, Player_GetAnimationIndex, Player.GetAnimationIndex, (void *player), (player))

// Player_GetFacingAngle (float)
DECL_FN(float, Player_GetFacingAngle, Player.GetFacingAngle, (void *player), (player))

// Player_GetIp (int)
DECL_FN(int, Player_GetIp, Player.GetIp, (void *player, struct CAPIStringBuffer *ip), (player, ip))

// Player_GetSpecialAction (int)
DECL_FN(int, Player_GetSpecialAction, Player.GetSpecialAction, (void *player), (player))

// Player_GetVehicleID (int)
DECL_FN(int, Player_GetVehicleID, Player.GetVehicleID, (void *player), (player))

// Player_GetVehicleSeat (int)
DECL_FN(int, Player_GetVehicleSeat, Player.GetVehicleSeat, (void *player), (player))

// Player_GetWeaponData (bool)
DECL_FN(bool, Player_GetWeaponData, Player.GetWeaponData, (void *player, int slot, int *weaponid, int *ammo), (player, slot, weaponid, ammo))

// Player_GetWeaponState (int)
DECL_FN(int, Player_GetWeaponState, Player.GetWeaponState, (void *player), (player))

// Player_InterpolateCameraPos (bool)
DECL_FN(bool, Player_InterpolateCameraPos, Player.InterpolateCameraPos, (void *player, float from_x, float from_y, float from_z, float to_x, float to_y, float to_z, int time, int cut), (player, from_x, from_y, from_z, to_x, to_y, to_z, time, cut))

// Player_InterpolateCameraLookAt (bool)
DECL_FN(bool, Player_InterpolateCameraLookAt, Player.InterpolateCameraLookAt, (void *player, float from_x, float from_y, float from_z, float to_x, float to_y, float to_z, int time, int cut), (player, from_x, from_y, from_z, to_x, to_y, to_z, time, cut))

// Player_IsPlayerAttachedObjectSlotUsed (bool)
DECL_FN(bool, Player_IsPlayerAttachedObjectSlotUsed, Player.IsPlayerAttachedObjectSlotUsed, (void *player, int index), (player, index))

// Player_AttachCameraToObject (bool)
DECL_FN(bool, Player_AttachCameraToObject, Player.AttachCameraToObject, (void *player, void *object), (player, object))

// Player_AttachCameraToPlayerObject (bool)
DECL_FN(bool, Player_AttachCameraToPlayerObject, Player.AttachCameraToPlayerObject, (void *player, void *object), (player, object))

// Player_GetCameraAspectRatio (float)
DECL_FN(float, Player_GetCameraAspectRatio, Player.GetCameraAspectRatio, (void *player), (player))

// Player_GetCameraFrontVector (bool)
DECL_FN(bool, Player_GetCameraFrontVector, Player.GetCameraFrontVector, (void *player, float *x, float *y, float *z), (player, x, y, z))

// Player_GetCameraMode (int)
DECL_FN(int, Player_GetCameraMode, Player.GetCameraMode, (void *player), (player))

// Player_GetKeys (bool)
DECL_FN(bool, Player_GetKeys, Player.GetKeys, (void *player, int *keys, int *updown, int *leftright), (player, keys, updown, leftright))

// Player_GetSurfingVehicle (void*)
DECL_FN(void *, Player_GetSurfingVehicle, Player.GetSurfingVehicle, (void *player), (player))

// Player_GetSurfingObject (void*)
DECL_FN(void *, Player_GetSurfingObject, Player.GetSurfingObject, (void *player), (player))

// Player_GetTargetPlayer (void*)
DECL_FN(void *, Player_GetTargetPlayer, Player.GetTargetPlayer, (void *player), (player))

// Player_GetTargetActor (void*)
DECL_FN(void *, Player_GetTargetActor, Player.GetTargetActor, (void *player), (player))

// Player_IsInVehicle (bool)
DECL_FN(bool, Player_IsInVehicle, Player.IsInVehicle, (void *player, void *targetVehicle), (player, targetVehicle))

// Player_IsInAnyVehicle (bool)
DECL_FN(bool, Player_IsInAnyVehicle, Player.IsInAnyVehicle, (void *player), (player))

// Player_IsInRangeOfPoint (bool)
DECL_FN(bool, Player_IsInRangeOfPoint, Player.IsInRangeOfPoint, (void *player, float range, float x, float y, float z), (player, range, x, y, z))

// Player_PlayCrimeReport (bool)
DECL_FN(bool, Player_PlayCrimeReport, Player.PlayCrimeReport, (void *player, void *suspect, int crime), (player, suspect, crime))

// Player_RemoveAttachedObject (bool)
DECL_FN(bool, Player_RemoveAttachedObject, Player.RemoveAttachedObject, (void *player, int index), (player, index))

// Player_SetAttachedObject (bool)
DECL_FN(bool, Player_SetAttachedObject, Player.SetAttachedObject, (void *player, int index, int modelid, int bone, float offsetX, float offsetY, float offsetZ, float rotationX, float rotationY, float rotationZ, float scaleX, float scaleY, float scaleZ, uint32_t materialcolor1, uint32_t materialcolor2), (player, index, modelid, bone, offsetX, offsetY, offsetZ, rotationX, rotationY, rotationZ, scaleX, scaleY, scaleZ, materialcolor1, materialcolor2))

// Player_GetAttachedObject (bool)
DECL_FN(bool, Player_GetAttachedObject, Player.GetAttachedObject, (void *player, int index, int *modelid, int *bone, float *offsetX, float *offsetY, float *offsetZ, float *rotationX, float *rotationY, float *rotationZ, float *scaleX, float *scaleY, float *scaleZ, int *materialcolor1, int *materialcolor2), (player, index, modelid, bone, offsetX, offsetY, offsetZ, rotationX, rotationY, rotationZ, scaleX, scaleY, scaleZ, materialcolor1, materialcolor2))

// Player_SetFacingAngle (bool)
DECL_FN(bool, Player_SetFacingAngle, Player.SetFacingAngle, (void *player, float angle), (player, angle))

// Player_SetMarkerForPlayer (bool)
DECL_FN(bool, Player_SetMarkerForPlayer, Player.SetMarkerForPlayer, (void *player, void *other, uint32_t color), (player, other, color))

// Player_GetMarkerForPlayer (uint32_t)
DECL_FN(uint32_t, Player_GetMarkerForPlayer, Player.GetMarkerForPlayer, (void *player, void *other), (player, other))

// Player_AllowTeleport (bool)
DECL_FN(bool, Player_AllowTeleport, Player.AllowTeleport, (void *player, bool allow), (player, allow))

// Player_IsTeleportAllowed (bool)
DECL_FN(bool, Player_IsTeleportAllowed, Player.IsTeleportAllowed, (void *player), (player))

// Player_DisableRemoteVehicleCollisions (bool)
DECL_FN(bool, Player_DisableRemoteVehicleCollisions, Player.DisableRemoteVehicleCollisions, (void *player, bool disable), (player, disable))

// Player_GetCameraZoom (float)
DECL_FN(float, Player_GetCameraZoom, Player.GetCameraZoom, (void *player), (player))

// Player_SelectTextDraw (bool)
DECL_FN(bool, Player_SelectTextDraw, Player.SelectTextDraw, (void *player, uint32_t hoverColour), (player, hoverColour))

// Player_CancelSelectTextDraw (bool)
DECL_FN(bool, Player_CancelSelectTextDraw, Player.CancelSelectTextDraw, (void *player), (player))

// Player_SendClientCheck (bool)
DECL_FN(bool, Player_SendClientCheck, Player.SendClientCheck, (void *player, int actionType, int address, int offset, int count), (player, actionType, address, offset, count))

// Player_Spawn (bool)
DECL_FN(bool, Player_Spawn, Player.Spawn, (void *player), (player))

// Player_GPCI (bool)
DECL_FN(bool, Player_GPCI, Player.GPCI, (void *player, struct CAPIStringView *gpci), (player, gpci))

// Player_IsAdmin (bool)
DECL_FN(bool, Player_IsAdmin, Player.IsAdmin, (void *player), (player))

// Player_Kick (bool)
DECL_FN(bool, Player_Kick, Player.Kick, (void *player), (player))

// Player_ShowGameText (bool)
DECL_FN(bool, Player_ShowGameText, Player.ShowGameText, (void *player, const char *text, int time, int style), (player, text, time, style))

// Player_HideGameText (bool)
DECL_FN(bool, Player_HideGameText, Player.HideGameText, (void *player, int style), (player, style))

// Player_HasGameText (bool)
DECL_FN(bool, Player_HasGameText, Player.HasGameText, (void *player, int style), (player, style))

// Player_GetGameText (bool)
DECL_FN(bool, Player_GetGameText, Player.GetGameText, (void *player, int style, struct CAPIStringView *message, int *time, int *remaining), (player, style, message, time, remaining))

// Player_Ban (bool)
DECL_FN(bool, Player_Ban, Player.Ban, (void *player), (player))

// Player_BanEx (bool)
DECL_FN(bool, Player_BanEx, Player.BanEx, (void *player, const char *reason), (player, reason))

// Player_SendDeathMessage (bool)
DECL_FN(bool, Player_SendDeathMessage, Player.SendDeathMessage, (void *player, void *killer, void *killee, int weapon), (player, killer, killee, weapon))

// Player_SendMessageToPlayer (bool)
DECL_FN(bool, Player_SendMessageToPlayer, Player.SendMessageToPlayer, (void *player, void *sender, const char *message), (player, sender, message))

// Player_GetVersion (int)
DECL_FN(int, Player_GetVersion, Player.GetVersion, (void *player, struct CAPIStringView *version), (player, version))

// Player_GetSkillLevel (int)
DECL_FN(int, Player_GetSkillLevel, Player.GetSkillLevel, (void *player, int skill), (player, skill))

// Player_GetZAim (float)
DECL_FN(float, Player_GetZAim, Player.GetZAim, (void *player), (player))

// Player_GetSurfingOffsets (bool)
DECL_FN(bool, Player_GetSurfingOffsets, Player.GetSurfingOffsets, (void *player, float *offsetX, float *offsetY, float *offsetZ), (player, offsetX, offsetY, offsetZ))

// Player_GetRotationQuat (bool)
DECL_FN(bool, Player_GetRotationQuat, Player.GetRotationQuat, (void *player, float *x, float *y, float *z, float *w), (player, x, y, z, w))

// Player_GetPlayerSpectateID (int)
DECL_FN(int, Player_GetPlayerSpectateID, Player.GetPlayerSpectateID, (void *player), (player))

// Player_GetSpectateType (int)
DECL_FN(int, Player_GetSpectateType, Player.GetSpectateType, (void *player), (player))

// Player_GetRawIp (uint32_t)
DECL_FN(uint32_t, Player_GetRawIp, Player.GetRawIp, (void *player), (player))

// Player_SetGravity (bool)
DECL_FN(bool, Player_SetGravity, Player.SetGravity, (void *player, float gravity), (player, gravity))

// Player_GetGravity (float)
DECL_FN(float, Player_GetGravity, Player.GetGravity, (void *player), (player))

// Player_SetAdmin (bool)
DECL_FN(bool, Player_SetAdmin, Player.SetAdmin, (void *player, bool set), (player, set))

// Player_IsSpawned (bool)
DECL_FN(bool, Player_IsSpawned, Player.IsSpawned, (void *player), (player))

// Player_IsControllable (bool)
DECL_FN(bool, Player_IsControllable, Player.IsControllable, (void *player), (player))

// Player_IsCameraTargetEnabled (bool)
DECL_FN(bool, Player_IsCameraTargetEnabled, Player.IsCameraTargetEnabled, (void *player), (player))

// Player_ToggleGhostMode (bool)
DECL_FN(bool, Player_ToggleGhostMode, Player.ToggleGhostMode, (void *player, bool toggle), (player, toggle))

// Player_GetGhostMode (bool)
DECL_FN(bool, Player_GetGhostMode, Player.GetGhostMode, (void *player), (player))

// Player_AllowWeapons (bool)
DECL_FN(bool, Player_AllowWeapons, Player.AllowWeapons, (void *player, bool allow), (player, allow))

// Player_AreWeaponsAllowed (bool)
DECL_FN(bool, Player_AreWeaponsAllowed, Player.AreWeaponsAllowed, (void *player), (player))

// Player_IsPlayerUsingOfficialClient (bool)
DECL_FN(bool, Player_IsPlayerUsingOfficialClient, Player.IsPlayerUsingOfficialClient, (void *player), (player))

// Player_GetAnimationFlags (int)
DECL_FN(int, Player_GetAnimationFlags, Player.GetAnimationFlags, (void *player), (player))

// Player_IsInDriveByMode (bool)
DECL_FN(bool, Player_IsInDriveByMode, Player.IsInDriveByMode, (void *player), (player))

// Player_IsCuffed (bool)
DECL_FN(bool, Player_IsCuffed, Player.IsCuffed, (void *player), (player))

// Player_IsUsingOmp (bool)
DECL_FN(bool, Player_IsUsingOmp, Player.IsUsingOmp, (void *player), (player))

// Player_IsInModShop (bool)
DECL_FN(bool, Player_IsInModShop, Player.IsInModShop, (void *player), (player))

// Player_GetSirenState (int)
DECL_FN(int, Player_GetSirenState, Player.GetSirenState, (void *player), (player))

// Player_GetLandingGearState (int)
DECL_FN(int, Player_GetLandingGearState, Player.GetLandingGearState, (void *player), (player))

// Player_GetHydraReactorAngle (uint32_t)
DECL_FN(uint32_t, Player_GetHydraReactorAngle, Player.GetHydraReactorAngle, (void *player), (player))

// Player_GetTrainSpeed (float)
DECL_FN(float, Player_GetTrainSpeed, Player.GetTrainSpeed, (void *player), (player))

#endif