#ifndef BRIDGE_H
#define BRIDGE_H

#include <stdbool.h>
#include <stdint.h>

#include "../capi/include/ompcapi.h"

#define DECL_VOID_FN(WRAP_NAME, MEMBER, PARAMS, ARGS) \
    static inline void WRAP_NAME PARAMS { m_api.MEMBER ARGS; }

#define DECL_FN(RET_TYPE, WRAP_NAME, MEMBER, PARAMS, ARGS) \
    static inline RET_TYPE WRAP_NAME PARAMS { return m_api.MEMBER ARGS; }

extern struct OMPAPI_t m_api;
extern int g_Initialized;

// Debug functions (debug.c)
void DebugLog(const char *msg);
void DebugLogFormat(const char *format, ...);

// Event callbacks (bridge.c)
bool onPlayerConnectCB(struct EventArgs_onPlayerConnect *args);
bool onPlayerDisconnectCB(struct EventArgs_onPlayerDisconnect *args);
bool onPlayerSpawnCB(struct EventArgs_onPlayerSpawn *args);
bool onPlayerRequestSpawnCB(struct EventArgs_onPlayerRequestSpawn *args);
bool onPlayerRequestClassCB(struct EventArgs_onPlayerRequestClass *args);

// GO Event Handlers (Go tarafindan export ediliyor)
extern void OnComponentLoad(void);
extern void OnPlayerConnect(void *player);
extern void OnPlayerDisconnect(void *player, int reason);
extern void OnPlayerSpawn(void *player);
extern void OnPlayerRequestSpawn(void *player);
extern void OnPlayerRequestClass(void *player, int classId);

#include "player.h"
#include "vehicle.h"
#include "natives.h"

#endif
