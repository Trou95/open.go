
#include <stdio.h>
#include <time.h>
#include <stdarg.h>

void DebugLog(const char *msg)
{
#ifdef DEBUG
    FILE *f = fopen("debug.log", "a");
    if (f)
    {
        time_t now = time(NULL);
        struct tm *t = localtime(&now);
        fprintf(f, "[%02d:%02d:%02d] %s\n", t->tm_hour, t->tm_min, t->tm_sec, msg);
        fflush(f);
        fclose(f);
    }
#endif
}

void DebugLogFormat(const char *format, ...)
{
#ifdef DEBUG
    char buffer[1024];
    va_list args;
    va_start(args, format);
    vsnprintf(buffer, sizeof(buffer), format, args);
    va_end(args);
    DebugLog(buffer);
#endif
}