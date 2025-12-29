#ifndef VEHICLE_H
#define VEHICLE_H

#include "bridge.h"

static inline int Vehicle_Create(int model, float x, float y, float z, float rotationX, uint32_t color1, uint32_t color2, uint32_t respawnDelay, bool addsiren)
{
    int id;
    m_api.Vehicle.Create(model, x, y, z, rotationX, color1, color2, respawnDelay, addsiren, &id);
    return id;
}

static inline int Vehicle_GetID(void *vehicle)
{
    return m_api.Vehicle.GetID(vehicle);
}

static inline void *Vehicle_FromID(int vehicleid)
{
    return m_api.Vehicle.FromID(vehicleid);
}

#endif