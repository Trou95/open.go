#ifndef NATIVES_H
#define NATIVES_H

#include "bridge.h"

// This project still experimental

static inline void Native_Core_Log(const char *message)
{
    if (g_Initialized)
    {
        m_api.Core.Log(message);
    }
}

static inline int Native_AddPlayerClass(uint8_t team, int skin, float x, float y, float z, float angle, uint8_t weapon1, uint32_t ammo1, uint8_t weapon2, uint32_t ammo2, uint8_t weapon3, uint32_t ammo3)
{
    int id;
    m_api.Class.Add(team, skin, x, y, z, angle, weapon1, ammo1, weapon2, ammo2, weapon3, ammo3, &id);
    return id;
}

DECL_FN(bool, SendClientMessageToAll, All.SendClientMessage, (uint32_t color, const char *message), (color, message))

#endif