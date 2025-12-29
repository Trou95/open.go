#include "../include/bridge/bridge.h"

bool onPlayerConnectCB(struct EventArgs_onPlayerConnect *args)
{

    if (args == NULL || args->size == 0 || args->list == NULL || args->list->player == NULL)
        return true;

    void *player = *args->list->player;

    OnPlayerConnect(player);

    return true;
}

bool onPlayerSpawnCB(struct EventArgs_onPlayerSpawn *args)
{

    if (args == NULL || args->size == 0 || args->list == NULL || args->list->player == NULL)
        return true;

    void *player = *args->list->player;

    OnPlayerSpawn(player);

    return true;
}

bool onPlayerRequestSpawnCB(struct EventArgs_onPlayerRequestSpawn *args)
{

    if (args == NULL || args->size == 0 || args->list == NULL || args->list->player == NULL)
        return true;

    void *player = *args->list->player;

    OnPlayerRequestSpawn(player);

    return true;
}

bool onPlayerRequestClassCB(struct EventArgs_onPlayerRequestClass *args)
{
    if (args == NULL || args->size == 0 || args->list == NULL || args->list->player == NULL || args->list->classId == NULL)
        return true;

    void *player = *args->list->player;

    OnPlayerRequestClass(player, *args->list->classId);

    return true;
}