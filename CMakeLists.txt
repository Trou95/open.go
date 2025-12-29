cmake_minimum_required(VERSION 3.15)
project(open.go C)

set(OUTPUT_NAME "open.go")
set(BRIDGE_DIR "${CMAKE_SOURCE_DIR}/bridge")
set(INCLUDE_DIR "${CMAKE_SOURCE_DIR}/include")

# Build options
option(DEBUG "Enable debug logging" 0)

# Force 32-bit build
set(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} -m32")
set(CMAKE_EXE_LINKER_FLAGS "${CMAKE_EXE_LINKER_FLAGS} -m32")

file(GLOB BRIDGE_SOURCES "${BRIDGE_DIR}/*.c")

add_library(bridge STATIC ${BRIDGE_SOURCES})
target_include_directories(bridge PRIVATE
    ${INCLUDE_DIR}
    ${INCLUDE_DIR}/bridge
    ${INCLUDE_DIR}/capi
)

if(DEBUG)
    target_compile_definitions(bridge PRIVATE DEBUG)
endif()

set(GO_OUTPUT "${CMAKE_SOURCE_DIR}/${OUTPUT_NAME}.dll")

file(GLOB GO_SOURCES "${CMAKE_SOURCE_DIR}/*.go")

add_custom_command(
    OUTPUT ${GO_OUTPUT}
    COMMAND ${CMAKE_COMMAND} -E env
            "CGO_ENABLED=1"
            "GOOS=windows"
            "GOARCH=386"
            "CGO_CFLAGS=-I${INCLUDE_DIR} -I${INCLUDE_DIR}/bridge"
            "CGO_LDFLAGS=-L${CMAKE_BINARY_DIR} -lbridge"
            go build -buildmode=c-shared -o ${GO_OUTPUT} .
    WORKING_DIRECTORY ${CMAKE_SOURCE_DIR}
    DEPENDS bridge ${GO_SOURCES}
    COMMENT "Building Go shared library..."
)

add_custom_target(go_dll ALL DEPENDS ${GO_OUTPUT})

set_property(DIRECTORY PROPERTY ADDITIONAL_CLEAN_FILES
    ${GO_OUTPUT}
    "${CMAKE_SOURCE_DIR}/${OUTPUT_NAME}.h"
)
