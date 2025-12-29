#include "../include/bridge/bridge.h"

struct OMPAPI_t m_api;
int g_Initialized = 0;

#define COMPONENT_UID 0x45D0871CB78E6D72ULL
#define COMPONENT_NAME "Open.GO"

static void onComponentReady(void)
{
    // DebugLog("=== onComponentReady ===");

    if (m_api.Event.AddHandler)
    {
        m_api.Event.AddHandler("onPlayerConnect", EventPriorityType_Default, (void *)onPlayerConnectCB);
        m_api.Event.AddHandler("onPlayerDisconnect", EventPriorityType_Default, (void *)onPlayerDisconnectCB);
        m_api.Event.AddHandler("onPlayerRequestSpawn", EventPriorityType_Default, (void *)onPlayerRequestSpawnCB);
        m_api.Event.AddHandler("onPlayerRequestClass", EventPriorityType_Default, (void *)onPlayerRequestClassCB);
        m_api.Event.AddHandler("onPlayerSpawn", EventPriorityType_Default, (void *)onPlayerSpawnCB);
        // m_api.Event.AddHandler("onPlayerUpdate", EventPriorityType_Highest, (void *)onPlayerUpdateCB);
        // m_api.Event.AddHandler("onTick", EventPriorityType_Lowest, (void *)onTickCB);
    }

    OnComponentLoad();

    // DebugLog("=== onComponentReady complete ===");
}

static void onComponentReset(void) { DebugLog("onComponentReset"); }
static void onComponentFree(void) { DebugLog("onComponentFree"); }

OMP_API_EXPORT void *ComponentEntryPoint(void *arg)
{
    (void)arg;

    if (!omp_initialize_capi(&m_api))
    {
        // DebugLog("FATAL: omp_initialize_capi failed!");
        return NULL;
    }

    // DebugLog("ComponentEntryPoint called successfully.");

    if (m_api.Component.Create)
    {
        struct ComponentVersion version = {1, 0, 0, 0};
        void *component = m_api.Component.Create(
            COMPONENT_UID, COMPONENT_NAME, version,
            (void *)onComponentReady, (void *)onComponentReset, (void *)onComponentFree);
        // DebugLogFormat("Component created: %p", component);
        return component;
    }
    else
    {
        // DebugLog("FATAL: omp->Component.Create is NULL!");
        return NULL;
    }

    return (void *)1;
}