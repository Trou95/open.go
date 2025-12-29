package main

/*
#cgo CFLAGS: -I${SRCDIR}/include/bridge
#include "bridge.h"
*/
import "C"

import (
	"fmt"
	"unsafe"
)

const (
	SERVER_MAX_PLAYERS = 100
)

var CopSkins = [14]int{
	265, 266, 267, 280, 281, 282,
	283, 284, 285, 286, 288, 304,
	306, 307,
}

var RobberSkins = [9]int{
	28, 29, 30, 185, 186, 187, 188, 243, 245,
}

type Vector4 struct {
	x float32
	y float32
	z float32
	a float32
}

type SpawnInfo struct {
	Position Vector4
	Skins    map[int]int
	Weapons  [3]struct {
		id   uint8
		ammo uint32
	}
}

type TeamInfo struct {
	Id            int
	Name          string
	Color         uint32
	TextDrawColor string
	SpawnInfo     SpawnInfo
}

var Teams = []TeamInfo{
	{
		Id:            0,
		Name:          "Cops",
		Color:         0x0000FF,
		TextDrawColor: "~b~~h~~h~",
		SpawnInfo: SpawnInfo{
			Position: Vector4{
				x: 1569.0404,
				y: -1691.8381,
				z: 5.8906,
				a: 181.1926,
			},
			Skins: map[int]int{},
			Weapons: [3]struct {
				id   uint8
				ammo uint32
			}{
				{22, 150},
				{29, 500},
				{31, 500},
			},
		},
	},
	{
		Id:            1,
		Name:          "Robbers",
		Color:         0x0000FF,
		TextDrawColor: "~h~",
		SpawnInfo: SpawnInfo{
			Position: Vector4{
				x: 1655.8478,
				y: -1660.6270,
				z: 22.5156,
				a: 180.9567,
			},
			Skins: map[int]int{},
			Weapons: [3]struct {
				id   uint8
				ammo uint32
			}{
				{24, 14},
				{27, 150},
				{30, 200},
			},
		},
	},
}

type ServerInfo struct {
	Pools struct {
		Players [SERVER_MAX_PLAYERS]struct {
			Ptr  *Player
			Team int
		}
	}
}

var Server ServerInfo

//export OnPlayerConnect
func OnPlayerConnect(pPtr unsafe.Pointer) {
	player := (*Player)(pPtr)

	Server.Pools.Players[player.ID()].Ptr = player

	player.SetPos(1548.4720, -1675.3485, 14.5376)
	player.SetFacingAngle(90.9472)
	player.SetCameraPos(1540.1604, -1675.3358, 14.5376)
	player.SetCameraLookAt(1548.4720, -1675.3485, 14.5376, 1)

	text := fmt.Sprintf("%s joined the server", player.GetName())
	Natives.core.SendClientMessageToAll(0xFFFFFFFF, text)
}

//export OnPlayerDisconnect
func OnPlayerDisconnect(pPtr unsafe.Pointer, reason C.int) {
	player := (*Player)(pPtr)

	text := fmt.Sprintf("%s left the server", player.GetName())
	Natives.core.SendClientMessageToAll(0xFFFFFFFF, text)

	Server.Pools.Players[player.ID()].Ptr = nil
}

//export OnPlayerSpawn
func OnPlayerSpawn(pPtr unsafe.Pointer) {}

//export OnPlayerRequestSpawn
func OnPlayerRequestSpawn(pPtr unsafe.Pointer) {
	player := (*Player)(pPtr)

	playerTeamId := Server.Pools.Players[player.ID()].Team
	team := Teams[playerTeamId]

	player.SetSpawnInfo(uint8(team.Id), player.GetSkin(),
		team.SpawnInfo.Position.x, team.SpawnInfo.Position.y,
		team.SpawnInfo.Position.z, team.SpawnInfo.Position.a,
		team.SpawnInfo.Weapons[0].id, team.SpawnInfo.Weapons[0].ammo,
		team.SpawnInfo.Weapons[1].id, team.SpawnInfo.Weapons[1].ammo,
		team.SpawnInfo.Weapons[2].id, team.SpawnInfo.Weapons[2].ammo)

}

//export OnPlayerRequestClass
func OnPlayerRequestClass(pPtr unsafe.Pointer, classId C.int) {
	player := (*Player)(pPtr)

	if team := getTeamByClassId(int(classId)); team != nil {
		player.ShowGameText(team.TextDrawColor+team.Name, 1000, 4)
		Server.Pools.Players[player.ID()].Team = team.Id
	}

}

func getTeamByClassId(classId int) *TeamInfo {
	for _, team := range Teams {
		if _, exists := team.SpawnInfo.Skins[classId]; exists {
			return &team
		}
	}
	return nil
}

//export OnComponentLoad
func OnComponentLoad() {

	var team = &Teams[0]
	for _, skin := range CopSkins {
		classId := Natives.core.AddPlayerClass(uint8(team.Id), skin, 1548.4720, -1675.3485, 14.5376, 90.9472,
			team.SpawnInfo.Weapons[0].id, team.SpawnInfo.Weapons[0].ammo,
			team.SpawnInfo.Weapons[1].id, team.SpawnInfo.Weapons[1].ammo,
			team.SpawnInfo.Weapons[2].id, team.SpawnInfo.Weapons[2].ammo)
		team.SpawnInfo.Skins[int(classId)] = skin
	}

	team = &Teams[1]
	for _, skin := range RobberSkins {
		classId := Natives.core.AddPlayerClass(uint8(team.Id), skin, 1548.4720, -1675.3485, 14.5376, 90.9472,
			team.SpawnInfo.Weapons[0].id, team.SpawnInfo.Weapons[0].ammo,
			team.SpawnInfo.Weapons[1].id, team.SpawnInfo.Weapons[1].ammo,
			team.SpawnInfo.Weapons[2].id, team.SpawnInfo.Weapons[2].ammo)
		team.SpawnInfo.Skins[int(classId)] = skin
	}

	Natives.vehicle.Create(596, 1545.7382, -1676.2213, 5.6177, 90.0534, 0, 1, 300, false)
	Natives.vehicle.Create(596, 1545.7347, -1672.0780, 5.6177, 270.0534, 0, 1, 300, false)
	Natives.vehicle.Create(596, 1570.3313, -1710.9722, 5.6177, 178.844, 0, 1, 300, false)
	Natives.vehicle.Create(596, 1574.4552, -1711.1305, 5.6177, 359.5703, 0, 1, 300, false)
	Natives.vehicle.Create(596, 1602.0367, -1696.1107, 5.6117, 89.3892, 0, 1, 300, false)
	Natives.vehicle.Create(596, 1535.9309, -1672.3579, 13.1044, 181.240, 0, 1, 300, false)
	Natives.vehicle.Create(601, 1534.5891, -1645.1311, 5.6177, 0.1956, 0, 1, 300, false)
	Natives.vehicle.Create(523, 1591.4904, -1710.7095, 5.6177, 178.978, 1, 0, 300, false)

	Natives.vehicle.Create(475, 1671.3212, -1696.1503, 20.2078, 88.8737, 8, 8, 300, false)
	Natives.vehicle.Create(536, 1671.6910, -1703.1172, 20.2115, 89.1441, 11, 11, 300, false)
	Natives.vehicle.Create(547, 1671.8584, -1719.5580, 20.2115, 89.1441, 20, 20, 300, false)
	Natives.vehicle.Create(411, 1668.0576, -1694.8867, 15.3365, 90.9749, 9, 9, 300, false)
	Natives.vehicle.Create(461, 1672.0972, -1708.8220, 20.0553, 90.2679, 8, 8, 300, false)
	Natives.vehicle.Create(461, 1672.1149, -1712.5074, 20.0553, 90.2679, 20, 20, 300, false)
	Natives.vehicle.Create(521, 1668.1301, -1701.6453, 15.2038, 90.267, 20, 20, 300, false)

}

func main() {}
