#ifndef VEHICLE_H
#define VEHICLE_H

#include "bridge.h"

// Vehicle_Create - returns vehicle ID
static inline int Vehicle_Create(int model, float x, float y, float z, float rotation, int color1, int color2, int respawnDelay, bool addsiren)
{
    int id;
    m_api.Vehicle.Create(model, x, y, z, rotation, color1, color2, respawnDelay, addsiren, &id);
    return id;
}

// Vehicle_Destroy
DECL_FN(bool, Vehicle_Destroy, Vehicle.Destroy, (void *vehicle), (vehicle))

// Vehicle_FromID
DECL_FN(void *, Vehicle_FromID, Vehicle.FromID, (int vehicleid), (vehicleid))

// Vehicle_GetID
DECL_FN(int, Vehicle_GetID, Vehicle.GetID, (void *vehicle), (vehicle))

// Vehicle_GetMaxPassengerSeats (static - no vehicle pointer)
DECL_FN(int, Vehicle_GetMaxPassengerSeats, Vehicle.GetMaxPassengerSeats, (int modelid), (modelid))

// Vehicle_IsStreamedIn
DECL_FN(bool, Vehicle_IsStreamedIn, Vehicle.IsStreamedIn, (void *vehicle, void *player), (vehicle, player))

// Vehicle_GetPos
DECL_FN(bool, Vehicle_GetPos, Vehicle.GetPos, (void *vehicle, float *x, float *y, float *z), (vehicle, x, y, z))

// Vehicle_SetPos
DECL_FN(bool, Vehicle_SetPos, Vehicle.SetPos, (void *vehicle, float x, float y, float z), (vehicle, x, y, z))

// Vehicle_GetZAngle
DECL_FN(float, Vehicle_GetZAngle, Vehicle.GetZAngle, (void *vehicle), (vehicle))

// Vehicle_GetRotationQuat
DECL_FN(bool, Vehicle_GetRotationQuat, Vehicle.GetRotationQuat, (void *vehicle, float *w, float *x, float *y, float *z), (vehicle, w, x, y, z))

// Vehicle_GetDistanceFromPoint
DECL_FN(float, Vehicle_GetDistanceFromPoint, Vehicle.GetDistanceFromPoint, (void *vehicle, float x, float y, float z), (vehicle, x, y, z))

// Vehicle_SetZAngle
DECL_FN(bool, Vehicle_SetZAngle, Vehicle.SetZAngle, (void *vehicle, float angle), (vehicle, angle))

// Vehicle_SetParamsForPlayer
DECL_FN(bool, Vehicle_SetParamsForPlayer, Vehicle.SetParamsForPlayer, (void *vehicle, void *player, int objective, int doors), (vehicle, player, objective, doors))

// Vehicle_UseManualEngineAndLights (static - no params)
DECL_FN(bool, Vehicle_UseManualEngineAndLights, Vehicle.UseManualEngineAndLights, (void), ())

// Vehicle_SetParamsEx
DECL_FN(bool, Vehicle_SetParamsEx, Vehicle.SetParamsEx, (void *vehicle, int engine, int lights, int alarm, int doors, int bonnet, int boot, int objective), (vehicle, engine, lights, alarm, doors, bonnet, boot, objective))

// Vehicle_GetParamsEx
DECL_FN(bool, Vehicle_GetParamsEx, Vehicle.GetParamsEx, (void *vehicle, int *engine, int *lights, int *alarm, int *doors, int *bonnet, int *boot, int *objective), (vehicle, engine, lights, alarm, doors, bonnet, boot, objective))

// Vehicle_GetParamsSirenState
DECL_FN(int, Vehicle_GetParamsSirenState, Vehicle.GetParamsSirenState, (void *vehicle), (vehicle))

// Vehicle_SetParamsCarDoors
DECL_FN(bool, Vehicle_SetParamsCarDoors, Vehicle.SetParamsCarDoors, (void *vehicle, int frontLeft, int frontRight, int rearLeft, int rearRight), (vehicle, frontLeft, frontRight, rearLeft, rearRight))

// Vehicle_GetParamsCarDoors
DECL_FN(bool, Vehicle_GetParamsCarDoors, Vehicle.GetParamsCarDoors, (void *vehicle, int *frontLeft, int *frontRight, int *rearLeft, int *rearRight), (vehicle, frontLeft, frontRight, rearLeft, rearRight))

// Vehicle_SetParamsCarWindows
DECL_FN(bool, Vehicle_SetParamsCarWindows, Vehicle.SetParamsCarWindows, (void *vehicle, int frontLeft, int frontRight, int rearLeft, int rearRight), (vehicle, frontLeft, frontRight, rearLeft, rearRight))

// Vehicle_GetParamsCarWindows
DECL_FN(bool, Vehicle_GetParamsCarWindows, Vehicle.GetParamsCarWindows, (void *vehicle, int *frontLeft, int *frontRight, int *rearLeft, int *rearRight), (vehicle, frontLeft, frontRight, rearLeft, rearRight))

// Vehicle_SetToRespawn
DECL_FN(bool, Vehicle_SetToRespawn, Vehicle.SetToRespawn, (void *vehicle), (vehicle))

// Vehicle_LinkToInterior
DECL_FN(bool, Vehicle_LinkToInterior, Vehicle.LinkToInterior, (void *vehicle, int interiorid), (vehicle, interiorid))

// Vehicle_AddComponent
DECL_FN(bool, Vehicle_AddComponent, Vehicle.AddComponent, (void *vehicle, int componentid), (vehicle, componentid))

// Vehicle_RemoveComponent
DECL_FN(bool, Vehicle_RemoveComponent, Vehicle.RemoveComponent, (void *vehicle, int componentid), (vehicle, componentid))

// Vehicle_ChangeColor
DECL_FN(bool, Vehicle_ChangeColor, Vehicle.ChangeColor, (void *vehicle, int color1, int color2), (vehicle, color1, color2))

// Vehicle_ChangePaintjob
DECL_FN(bool, Vehicle_ChangePaintjob, Vehicle.ChangePaintjob, (void *vehicle, int paintjobid), (vehicle, paintjobid))

// Vehicle_SetHealth
DECL_FN(bool, Vehicle_SetHealth, Vehicle.SetHealth, (void *vehicle, float health), (vehicle, health))

// Vehicle_GetHealth
DECL_FN(float, Vehicle_GetHealth, Vehicle.GetHealth, (void *vehicle), (vehicle))

// Vehicle_AttachTrailer
DECL_FN(bool, Vehicle_AttachTrailer, Vehicle.AttachTrailer, (void *trailer, void *vehicle), (trailer, vehicle))

// Vehicle_DetachTrailer
DECL_FN(bool, Vehicle_DetachTrailer, Vehicle.DetachTrailer, (void *vehicle), (vehicle))

// Vehicle_IsTrailerAttached
DECL_FN(bool, Vehicle_IsTrailerAttached, Vehicle.IsTrailerAttached, (void *vehicle), (vehicle))

// Vehicle_GetTrailer
DECL_FN(void *, Vehicle_GetTrailer, Vehicle.GetTrailer, (void *vehicle), (vehicle))

// Vehicle_SetNumberPlate
DECL_FN(bool, Vehicle_SetNumberPlate, Vehicle.SetNumberPlate, (void *vehicle, const char *numberPlate), (vehicle, numberPlate))

// Vehicle_GetModel
DECL_FN(int, Vehicle_GetModel, Vehicle.GetModel, (void *vehicle), (vehicle))

// Vehicle_GetComponentInSlot
DECL_FN(int, Vehicle_GetComponentInSlot, Vehicle.GetComponentInSlot, (void *vehicle, int slot), (vehicle, slot))

// Vehicle_GetComponentType (static - no vehicle pointer)
DECL_FN(int, Vehicle_GetComponentType, Vehicle.GetComponentType, (int componentid), (componentid))

// Vehicle_CanHaveComponent (static - no vehicle pointer)
DECL_FN(bool, Vehicle_CanHaveComponent, Vehicle.CanHaveComponent, (int modelid, int componentid), (modelid, componentid))

// Vehicle_GetRandomColorPair (static)
DECL_FN(bool, Vehicle_GetRandomColorPair, Vehicle.GetRandomColorPair, (int modelid, int *color1, int *color2, int *color3, int *color4), (modelid, color1, color2, color3, color4))

// Vehicle_ColorIndexToColor (static)
DECL_FN(int, Vehicle_ColorIndexToColor, Vehicle.ColorIndexToColor, (int colorIndex, int alpha), (colorIndex, alpha))

// Vehicle_Repair
DECL_FN(bool, Vehicle_Repair, Vehicle.Repair, (void *vehicle), (vehicle))

// Vehicle_GetVelocity
DECL_FN(bool, Vehicle_GetVelocity, Vehicle.GetVelocity, (void *vehicle, float *x, float *y, float *z), (vehicle, x, y, z))

// Vehicle_SetVelocity
DECL_FN(bool, Vehicle_SetVelocity, Vehicle.SetVelocity, (void *vehicle, float x, float y, float z), (vehicle, x, y, z))

// Vehicle_SetAngularVelocity
DECL_FN(bool, Vehicle_SetAngularVelocity, Vehicle.SetAngularVelocity, (void *vehicle, float x, float y, float z), (vehicle, x, y, z))

// Vehicle_GetDamageStatus
DECL_FN(bool, Vehicle_GetDamageStatus, Vehicle.GetDamageStatus, (void *vehicle, int *panels, int *doors, int *lights, int *tires), (vehicle, panels, doors, lights, tires))

// Vehicle_UpdateDamageStatus
DECL_FN(bool, Vehicle_UpdateDamageStatus, Vehicle.UpdateDamageStatus, (void *vehicle, int panels, int doors, int lights, int tires), (vehicle, panels, doors, lights, tires))

// Vehicle_GetModelInfo (static)
DECL_FN(bool, Vehicle_GetModelInfo, Vehicle.GetModelInfo, (int vehiclemodel, int infotype, float *x, float *y, float *z), (vehiclemodel, infotype, x, y, z))

// Vehicle_SetVirtualWorld
DECL_FN(bool, Vehicle_SetVirtualWorld, Vehicle.SetVirtualWorld, (void *vehicle, int virtualWorld), (vehicle, virtualWorld))

// Vehicle_GetVirtualWorld
DECL_FN(int, Vehicle_GetVirtualWorld, Vehicle.GetVirtualWorld, (void *vehicle), (vehicle))

// Vehicle_GetLandingGearState
DECL_FN(int, Vehicle_GetLandingGearState, Vehicle.GetLandingGearState, (void *vehicle), (vehicle))

// Vehicle_IsValid
DECL_FN(bool, Vehicle_IsValid, Vehicle.IsValid, (void *vehicle), (vehicle))

// Vehicle_AddStatic - returns vehicle ID
static inline int Vehicle_AddStatic(int modelid, float x, float y, float z, float angle, int color1, int color2)
{
    int id;
    m_api.Vehicle.AddStatic(modelid, x, y, z, angle, color1, color2, &id);
    return id;
}

// Vehicle_AddStaticEx - returns vehicle ID
static inline int Vehicle_AddStaticEx(int modelid, float x, float y, float z, float angle, int color1, int color2, int respawnDelay, bool addSiren)
{
    int id;
    m_api.Vehicle.AddStaticEx(modelid, x, y, z, angle, color1, color2, respawnDelay, addSiren, &id);
    return id;
}

// Vehicle_EnableFriendlyFire (static - no params)
DECL_FN(bool, Vehicle_EnableFriendlyFire, Vehicle.EnableFriendlyFire, (void), ())

// Vehicle_GetSpawnInfo
DECL_FN(bool, Vehicle_GetSpawnInfo, Vehicle.GetSpawnInfo, (void *vehicle, float *x, float *y, float *z, float *rotation, int *color1, int *color2), (vehicle, x, y, z, rotation, color1, color2))

// Vehicle_SetSpawnInfo
DECL_FN(bool, Vehicle_SetSpawnInfo, Vehicle.SetSpawnInfo, (void *vehicle, int modelid, float x, float y, float z, float rotation, int color1, int color2, int respawn_time, int interior), (vehicle, modelid, x, y, z, rotation, color1, color2, respawn_time, interior))

// Vehicle_GetModelCount (static)
DECL_FN(int, Vehicle_GetModelCount, Vehicle.GetModelCount, (int modelid), (modelid))

// Vehicle_GetModelsUsed (static - no params)
DECL_FN(int, Vehicle_GetModelsUsed, Vehicle.GetModelsUsed, (void), ())

// Vehicle_GetPaintjob
DECL_FN(int, Vehicle_GetPaintjob, Vehicle.GetPaintjob, (void *vehicle), (vehicle))

// Vehicle_GetColor
DECL_FN(bool, Vehicle_GetColor, Vehicle.GetColor, (void *vehicle, int *color1, int *color2), (vehicle, color1, color2))

// Vehicle_GetInterior
DECL_FN(int, Vehicle_GetInterior, Vehicle.GetInterior, (void *vehicle), (vehicle))

// Vehicle_GetNumberPlate
DECL_FN(bool, Vehicle_GetNumberPlate, Vehicle.GetNumberPlate, (void *vehicle, struct CAPIStringView *numberPlate), (vehicle, numberPlate))

// Vehicle_SetRespawnDelay
DECL_FN(bool, Vehicle_SetRespawnDelay, Vehicle.SetRespawnDelay, (void *vehicle, int respawn_delay), (vehicle, respawn_delay))

// Vehicle_GetRespawnDelay
DECL_FN(int, Vehicle_GetRespawnDelay, Vehicle.GetRespawnDelay, (void *vehicle), (vehicle))

// Vehicle_GetCab
DECL_FN(void *, Vehicle_GetCab, Vehicle.GetCab, (void *vehicle), (vehicle))

// Vehicle_GetTower
DECL_FN(void *, Vehicle_GetTower, Vehicle.GetTower, (void *vehicle), (vehicle))

// Vehicle_GetOccupiedTick
DECL_FN(int, Vehicle_GetOccupiedTick, Vehicle.GetOccupiedTick, (void *vehicle), (vehicle))

// Vehicle_GetRespawnTick
DECL_FN(int, Vehicle_GetRespawnTick, Vehicle.GetRespawnTick, (void *vehicle), (vehicle))

// Vehicle_HasBeenOccupied
DECL_FN(bool, Vehicle_HasBeenOccupied, Vehicle.HasBeenOccupied, (void *vehicle), (vehicle))

// Vehicle_IsOccupied
DECL_FN(bool, Vehicle_IsOccupied, Vehicle.IsOccupied, (void *vehicle), (vehicle))

// Vehicle_IsDead
DECL_FN(bool, Vehicle_IsDead, Vehicle.IsDead, (void *vehicle), (vehicle))

// Vehicle_SetParamsSirenState
DECL_FN(bool, Vehicle_SetParamsSirenState, Vehicle.SetParamsSirenState, (void *vehicle, bool siren_state), (vehicle, siren_state))

// Vehicle_ToggleSirenEnabled
DECL_FN(bool, Vehicle_ToggleSirenEnabled, Vehicle.ToggleSirenEnabled, (void *vehicle, bool status), (vehicle, status))

// Vehicle_IsSirenEnabled
DECL_FN(bool, Vehicle_IsSirenEnabled, Vehicle.IsSirenEnabled, (void *vehicle), (vehicle))

// Vehicle_GetLastDriver
DECL_FN(void *, Vehicle_GetLastDriver, Vehicle.GetLastDriver, (void *vehicle), (vehicle))

// Vehicle_GetDriver
DECL_FN(void *, Vehicle_GetDriver, Vehicle.GetDriver, (void *vehicle), (vehicle))

// Vehicle_GetSirenState
DECL_FN(int, Vehicle_GetSirenState, Vehicle.GetSirenState, (void *vehicle), (vehicle))

// Vehicle_GetHydraReactorAngle
DECL_FN(uint32_t, Vehicle_GetHydraReactorAngle, Vehicle.GetHydraReactorAngle, (void *vehicle), (vehicle))

// Vehicle_GetTrainSpeed
DECL_FN(float, Vehicle_GetTrainSpeed, Vehicle.GetTrainSpeed, (void *vehicle), (vehicle))

// Vehicle_GetMatrix
DECL_FN(bool, Vehicle_GetMatrix, Vehicle.GetMatrix, (void *vehicle, float *rightX, float *rightY, float *rightZ, float *upX, float *upY, float *upZ, float *atX, float *atY, float *atZ), (vehicle, rightX, rightY, rightZ, upX, upY, upZ, atX, atY, atZ))

// Vehicle_GetOccupant
DECL_FN(void *, Vehicle_GetOccupant, Vehicle.GetOccupant, (void *vehicle, int seat), (vehicle, seat))

// Vehicle_CountOccupants
DECL_FN(int, Vehicle_CountOccupants, Vehicle.CountOccupants, (void *vehicle), (vehicle))

#endif
