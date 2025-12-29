# Open.GO

A bare TDM gamemode for [open.mp](https://open.mp) servers with Go. Write your game mode logic in Go while leveraging the open.mp C API.


## Project Structure

```
├── bridge/          # C bridge code (callbacks, entry point)
├── include/
│   ├── bridge/      # Bridge headers (player.h, vehicle.h, natives.h)
│   └── capi/        # open.mp C API headers
├── *.go             # Go source files
└── CMakeLists.txt
```

## Usage

1. Build the project
2. Copy `open.go.dll` to your open.mp server's `components/` folder
3. Start the server

**Note:** This project is experimental. Do not use in production.

## License

MIT
