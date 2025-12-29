package main

const (
	SAMP_INVALID_PLAYER_ID  = 0xFFFF
	SAMP_INVALID_VEHICLE_ID = 0xFFFF
	SAMP_INVALID_MENU       = 0xFF
	SAMP_INVALID_TEXT_DRAW  = 0xFFFF
	SAMP_INVALID_GANG_ZONE  = -1
	SAMP_INVALID_3DTEXT_ID  = 0xFFFF
	SAMP_NO_TEAM            = 255

	SAMP_MAX_VEHICLES                = 2000
	SAMP_MAX_ACTORS                  = 1000
	SAMP_MAX_PLAYER_ATTACHED_OBJECTS = 10
	SAMP_MAX_PLAYER_TEXT_DRAWS       = 2048
	SAMP_MAX_PLAYER_3DTEXTS          = 1024
)

const (
	PLAYER_STATE_NONE = iota
	PLAYER_STATE_ONFOOT
	PLAYER_STATE_DRIVER
	PLAYER_STATE_PASSENGER
	PLAYER_STATE_EXIT_VEHICLE
	PLAYER_STATE_ENTER_VEHICLE_DRVER
	PLAYER_STATE_ENTER_VEHICLE_PASSENGER
	PLAYER_STATE_WASTED
	PLAYER_STATE_SPAWNED
	PLAYER_STATE_SPECTATING
)

const (
	PLAYER_MARKERS_MODE_OFF = iota
	PLAYER_MARKERS_MODE_GLOBAL
	PLAYER_MARKERS_MODE_STREAMED
)

const (
	SPECIAL_ACTION_DUCK = iota
	SPECIAL_ACTION_USEJETPACK
	SPECIAL_ACTION_ENTER_VEHICLE
	SPECIAL_ACTION_EXIT_VEHICLE
	SPECIAL_ACTION_DANCE1
	SPECIAL_ACTION_DANCE2
	SPECIAL_ACTION_DANCE3
	SPECIAL_ACTION_DANCE4
	SPECIAL_ACTION_HANDSUP
	SPECIAL_ACTION_USECELLPHONE
	SPECIAL_ACTION_SITTING
	SPECIAL_ACTION_STOPUSECELLPHONE
	SPECIAL_ACTION_DRINK_BEER
	SPECIAL_ACTION_SMOKE_CIGGY
	SPECIAL_ACTION_DRINK_WINE
	SPECIAL_ACTION_DRINK_SPRUNK
	SPECIAL_ACTION_CUFFED
	SPECIAL_ACTION_CARRY
)

const (
	FIGHT_STYLE_NORMAL   = 4
	FIGHT_STYLE_BOXING   = 5
	FIGHT_STYLE_KUNGFU   = 6
	FIGHT_STYLE_KNEEHEAD = 7
	FIGHT_STYLE_GRABKICK = 15
	FIGHT_STYLE_ELBOW    = 16
)

const (
	WEAPONSKILL_PISTOL = iota
	WEAPONSKILL_PISTOL_SILENCED
	EAPONSKILL_DESERT_EAGLE
	EAPONSKILL_SHOTGUN
	EAPONSKILL_SAWNOFF_SHOTGUN
	EAPONSKILL_SPAS12_SHOTGUN
	EAPONSKILL_MICRO_UZI
	EAPONSKILL_MP5
	EAPONSKILL_AK47
	EAPONSKILL_M4
	EAPONSKILL_SNIPERRIFLE
)

const (
	WEAPONSTATE_UNKNOWN      = -1
	WEAPONSTATE_NO_BULLETS   = 0
	WEAPONSTATE_LAST_BULLET  = 1
	WEAPONSTATE_MORE_BULLETS = 2
	WEAPONSTATE_RELOADING    = 3
)

const (
	WEAPON_BRASSKNUCKLE = iota + 1
	WEAPON_GOLFCLUB
	WEAPON_NITESTICK
	WEAPON_KNIFE
	WEAPON_BAT
	WEAPON_SHOVEL
	WEAPON_POOLSTICK
	WEAPON_KATANA
	WEAPON_CHAINSAW
	WEAPON_DILDO
	WEAPON_DILDO2
	WEAPON_VIBRATOR
	WEAPON_VIBRATOR2
	WEAPON_FLOWER
	WEAPON_CANE
	WEAPON_GRENADE
	WEAPON_TEARGAS
	WEAPON_MOLTOV
	WEAPON_COLT45
	WEAPON_SILENCED
	WEAPON_DEAGLE
	WEAPON_SHOTGUN
	WEAPON_SAWEDOFF
	WEAPON_SHOTGSPA
	WEAPON_UZI
	WEAPON_MP5
	WEAPON_AK47
	WEAPON_M4
	WEAPON_TEC9
	WEAPON_RIFLE
	WEAPON_SNIPER
	WEAPON_ROCKETLAUNCHER
	WEAPON_HEATSEEKER
	WEAPON_FLAMETHROWER
	WEAPON_MINIGUN
	WEAPON_SATCHEL
	WEAPON_BOMB
	WEAPON_SPRAYCAN
	WEAPON_FIREEXTINGUISHER
	WEAPON_CAMERA
	WEAPON_PARACHUTE
	WEAPON_VEHICLE
	WEAPON_DROWN
	WEAPON_COLLISION
)

const (
	KEY_ACTION = 1 << iota
	KEY_CROUCH
	KEY_FIRE
	KEY_SPRINT
	KEY_SECONDARY_ATTACK
	KEY_JUMP
	KEY_LOOK_RIGHT
	KEY_HANDBRAKE
	KEY_LOOK_LEFT
	KEY_SUBMISSION
	KEY_LOOK_BEHIND
	KEY_WALK
	KEY_ANALOG_UP
	KEY_ANALOG_DOWN
	KEY_ANALOG_LEFT
	KEY_ANALOG_RIGHT
	KEY_YES
	KEY_NO
	KEY_CTRL_BACK

	KEY_uP    = -128
	KEY_DOWN  = 128
	KEY_LEFT  = -128
	KEY_RIGHT = 128
)

const (
	MAPICON_LOCAL = iota
	MAPICON_GLOBAL
	MAPICON_LOCAL_CHECKPOINT
	MAPICON_GLOBAL_CHECKPOINT
)

const (
	CAMERA_CUT  = 2
	CAMERA_MOVE = 1
)

const (
	SPECTATE_MODE_NORMAL = iota
	SPECTATE_MODE_FIXED
	SPECTATE_MODE_SIDE
)

const (
	PLAYER_RECORDING_TYPE_NONE = iota
	PLAYER_RECORDING_TYPE_DRIVER
	PLAYER_RECORDING_TYPE_ONFOOT
)

const (
	CLICK_SOURCE_SCOREBOARD = 0
)

const (
	SELECT_OBJECT_GLOBAL = 1
	SELECT_OBJECT_PLAYER = 2
)

const (
	CARMODTYPE_SPOILER = iota
	CARMODTYPE_HOOD
	CARMODTYPE_ROOF
	CARMODTYPE_SIDESKIRT
	CARMODTYPE_LAMPS
	CARMODTYPE_NITRO
	CARMODTYPE_EXHAUST
	CARMODTYPE_WHEELS
	CARMODTYPE_STEREO
	CARMODTYPE_HYDRAULICS
	CARMODTYPE_FRONT_BUMPER
	CARMODTYPE_REAR_BUMPER
	CARMODTYPE_VENT_RIGHT
	CARMODTYPE_VENT_LEFT
)

const (
	VEHICLE_PARAMS_UNSET = -1
	VEHICLE_PARAMS_OFF   = 0
	VEHICLE_PARAMS_ON    = 1
)

const (
	VEHICLE_MODEL_INFO_SIZE = iota
	VEHICLE_MODEL_INFO_FRONTSEAT
	VEHICLE_MODEL_INFO_REARSEAT
	VEHICLE_MODEL_INFO_PETROLCAP
	VEHICLE_MODEL_INFO_WHEELSFRONT
	VEHICLE_MODEL_INFO_WHEELSREAR
	VEHICLE_MODEL_INFO_WHEELSMID
	VEHICLE_MODEL_INFO_FRONT_BUMPER_Z
	VEHICLE_MODEL_INFO_REAR_BUMPER_Z
)

const (
	OBJECT_MATERIAL_SIZE_32x32   = 10
	OBJECT_MATERIAL_SIZE_64x32   = 20
	OBJECT_MATERIAL_SIZE_64x64   = 30
	OBJECT_MATERIAL_SIZE_128x32  = 40
	OBJECT_MATERIAL_SIZE_128x64  = 50
	OBJECT_MATERIAL_SIZE_128x128 = 60
	OBJECT_MATERIAL_SIZE_256x32  = 70
	OBJECT_MATERIAL_SIZE_256x64  = 80
	OBJECT_MATERIAL_SIZE_256x128 = 90
	OBJECT_MATERIAL_SIZE_256x256 = 100
	OBJECT_MATERIAL_SIZE_512x64  = 110
	OBJECT_MATERIAL_SIZE_512x128 = 120
	OBJECT_MATERIAL_SIZE_512x256 = 130
	OBJECT_MATERIAL_SIZE_512x512 = 140
)

const (
	OBJECT_MATERIAL_TEXT_ALIGN_LEFT = iota
	OBJECT_MATERIAL_TEXT_ALIGN_CENTER
	OBJECT_MATERIAL_TEXT_ALIGN_RIGHT
)

var weaponNames = map[int]string{
	0:  "Fist",
	1:  "Brass Knuckles",
	2:  "Golf Club",
	3:  "Nightstick",
	4:  "Knife",
	5:  "Baseball Bat",
	6:  "Shovel",
	7:  "Pool Cue",
	8:  "Katana",
	9:  "Chainsaw",
	10: "Dildo",
	11: "Dildo",
	12: "Vibrator",
	13: "Vibrator",
	14: "Flower",
	15: "Cane",
	16: "Grenade",
	17: "Tear Gas",
	18: "Molotov",
	22: "Colt 45",
	23: "Silenced Pistol",
	24: "Desert Eagle",
	25: "Shotgun",
	26: "Sawn-off Shotgun",
	27: "Spas-12 Shotgun",
	28: "Micro Uzi",
	29: "MP5",
	30: "AK-47",
	31: "M4",
	32: "Tec-9",
	33: "Rifle",
	34: "Sniper Rifle",
	35: "Rocket Launcher",
	36: "Heat-Seeking Rocket Launcher",
	37: "Flamethrower",
	38: "Minigun",
	39: "Satchel Charge",
	40: "Bomb",
	41: "Spray Can",
	42: "Fire Extinguisher",
	43: "Camera",
	44: "Parachute",
}

var vehicleModelNames = []string{
	"Landstalker", "Bravura", "Buffalo", "Linerunner", "Perennial", "Sentinel", "Dumper", "Fire Truck", "Trashmaster", "Stretch", "Manana",
	"Infernus", "Voodoo", "Pony", "Mule", "Cheetah", "Ambulance", "Leviathan", "Moonbeam", "Esperanto", "Taxi", "Washington", "Bobcat",
	"Mr. Whoopee", "BF Injection", "Hunter", "Premier", "Enforcer", "Securicar", "Banshee", "Predator", "Bus", "Rhino", "Barracks", "Hotknife",
	"Articulated Trailer", "Previon", "Coach", "Cabbie", "Stallion", "Rumpo", "RC Bandit", "Romero", "Packer", "Monster", "Admiral", "Squalo",
	"Seasparrow", "Pizzaboy", "Tram", "Articulated Trailer 2", "Turismo", "Speeder", "Reefer", "Tropic", "Flatbed", "Yankee", "Caddy", "Solair",
	"Berkley's RC Van", "Skimmer", "PCJ-600", "Faggio", "Freeway", "RC Baron", "RC Raider", "Glendale", "Oceanic", "Sanchez", "Sparrow",
	"Patriot", "Quad", "Coastguard", "Dinghy", "Hermes", "Sabre", "Rustler", "ZR-350", "Walton", "Regina", "Comet", "BMX", "Burrito",
	"Camper", "Marquis", "Baggage", "Dozer", "Maverick", "News Chopper", "Rancher", "FBI Rancher", "Virgo", "Greenwood", "Jetmax", "Hotring Racer",
	"Sandking", "Blista Compact", "Police Maverick", "Boxville", "Benson", "Mesa", "RC Goblin", "Hotring Racer A", "Hotring Racer B",
	"Bloodring Banger", "Rancher", "Super GT", "Elegant", "Journey", "Bike", "Mountain Bike", "Beagle", "Cropduster", "Stuntplane", "Tanker",
	"Roadtrain", "Nebula", "Majestic", "Buccaneer", "Shamal", "Hydra", "FCR-900", "NRG-500", "HPV1000", "Cement Truck", "Towtruck", "Fortune",
	"Cadrona", "FBI Truck", "Willard", "Forklift", "Tractor", "Combine Harvester", "Feltzer", "Remington", "Slamvan", "Blade", "Freight", "Brown Streak",
	"Vortex", "Vincent", "Bullet", "Clover", "Sadler", "Fire Truck Ladder", "Hustler", "Intruder", "Primo", "Cargobob", "Tampa", "Sunrise", "Merit",
	"Utility Van", "Nevada", "Yosemite", "Windsor", "Monster A", "Monster B", "Uranus", "Jester", "Sultan", "Stratum", "Elegy", "Raindance",
	"RC Tiger", "Flash", "Tahoma", "Savanna", "Bandito", "Freight Flat", "Streak Carriage", "Kart", "Mower", "Dune", "Sweeper", "Broadway",
	"Tornado", "AT-400", "DFT-30", "Huntley", "Stafford", "BF-400", "Newsvan", "Tug", "Tanker Trailer", "Emperor", "Wayfarer", "Euros", "Hotdog",
	"Club", "Freight Box", "Articulated Trailer 3", "Andromada", "Dodo", "RC Cam", "Launch", "Police (LS)", "Police (SF)",
	"Police (LV)", "Ranger", "Picador", "S.W.A.T.", "Alpha", "Phoenix", "Glendale (Damaged)", "Sadler (Damaged)", "Baggage Box A",
	"Baggage Box B", "Tug Stairs", "Boxville", "Farm Trailer", "Utility Trailer",
}

type ZoneCords struct {
	minX float32
	minY float32
	minZ float32
	maxX float32
	maxY float32
	maxZ float32
}

type Zone struct {
	Name  string
	Cords ZoneCords
}

var zones = []Zone{
	{Name: "The Big Ear", Cords: ZoneCords{minX: -410, minY: 1403.3, minZ: -3, maxX: -137.9, maxY: 1681.2, maxZ: 200}},
	{Name: "Aldea Malvada", Cords: ZoneCords{minX: -1372.1, minY: 2498.5, minZ: 0, maxX: -1277.5, maxY: 2615.3, maxZ: 200}},
	{Name: "Angel Pine", Cords: ZoneCords{minX: -2324.9, minY: -2584.2, minZ: -6.1, maxX: -1964.2, maxY: -2212.1, maxZ: 200}},
	{Name: "Arco del Oeste", Cords: ZoneCords{minX: -901.1, minY: 2221.8, minZ: 0, maxX: -592, maxY: 2571.9, maxZ: 200}},
	{Name: "Avispa Country Club", Cords: ZoneCords{minX: -2646.4, minY: -355.4, minZ: 0, maxX: -2270, maxY: -222.5, maxZ: 200}},
	{Name: "Avispa Country Club", Cords: ZoneCords{minX: -2831.8, minY: -430.2, minZ: -6.1, maxX: -2646.4, maxY: -222.5, maxZ: 200}},
	{Name: "Avispa Country Club", Cords: ZoneCords{minX: -2361.5, minY: -417.1, minZ: 0, maxX: -2270, maxY: -355.4, maxZ: 200}},
	{Name: "Avispa Country Club", Cords: ZoneCords{minX: -2667.8, minY: -302.1, minZ: -28.8, maxX: -2646.4, maxY: -262.3, maxZ: 71.1}},
	{Name: "Avispa Country Club", Cords: ZoneCords{minX: -2470, minY: -355.4, minZ: 0, maxX: -2270, maxY: -318.4, maxZ: 46.1}},
	{Name: "Avispa Country Club", Cords: ZoneCords{minX: -2550, minY: -355.4, minZ: 0, maxX: -2470, maxY: -318.4, maxZ: 39.7}},
	{Name: "Back o Beyond", Cords: ZoneCords{minX: -1166.9, minY: -2641.1, minZ: 0, maxX: -321.7, maxY: -1856, maxZ: 200}},
	{Name: "Battery Point", Cords: ZoneCords{minX: -2741, minY: 1268.4, minZ: -4.5, maxX: -2533, maxY: 1490.4, maxZ: 200}},
	{Name: "Bayside", Cords: ZoneCords{minX: -2741, minY: 2175.1, minZ: 0, maxX: -2353.1, maxY: 2722.7, maxZ: 200}},
	{Name: "Bayside Marina", Cords: ZoneCords{minX: -2353.1, minY: 2275.7, minZ: 0, maxX: -2153.1, maxY: 2475.7, maxZ: 200}},
	{Name: "Beacon Hill", Cords: ZoneCords{minX: -399.6, minY: -1075.5, minZ: -1.4, maxX: -319, maxY: -977.5, maxZ: 198.5}},
	{Name: "Blackfield", Cords: ZoneCords{minX: 964.3, minY: 1203.2, minZ: -89, maxX: 1197.3, maxY: 1403.2, maxZ: 110.9}},
	{Name: "Blackfield", Cords: ZoneCords{minX: 964.3, minY: 1403.2, minZ: -89, maxX: 1197.3, maxY: 1726.2, maxZ: 110.9}},
	{Name: "Blackfield Chapel", Cords: ZoneCords{minX: 1375.6, minY: 596.3, minZ: -89, maxX: 1558, maxY: 823.2, maxZ: 110.9}},
	{Name: "Blackfield Chapel", Cords: ZoneCords{minX: 1325.6, minY: 596.3, minZ: -89, maxX: 1375.6, maxY: 795, maxZ: 110.9}},
	{Name: "Blackfield Intersection", Cords: ZoneCords{minX: 1197.3, minY: 1044.6, minZ: -89, maxX: 1277, maxY: 1163.3, maxZ: 110.9}},
	{Name: "Blackfield Intersection", Cords: ZoneCords{minX: 1166.5, minY: 795, minZ: -89, maxX: 1375.6, maxY: 1044.6, maxZ: 110.9}},
	{Name: "Blackfield Intersection", Cords: ZoneCords{minX: 1277, minY: 1044.6, minZ: -89, maxX: 1315.3, maxY: 1087.6, maxZ: 110.9}},
	{Name: "Blackfield Intersection", Cords: ZoneCords{minX: 1375.6, minY: 823.2, minZ: -89, maxX: 1457.3, maxY: 919.4, maxZ: 110.9}},
	{Name: "Blueberry", Cords: ZoneCords{minX: 104.5, minY: -220.1, minZ: 2.3, maxX: 349.6, maxY: 152.2, maxZ: 200}},
	{Name: "Blueberry", Cords: ZoneCords{minX: 19.6, minY: -404.1, minZ: 3.8, maxX: 349.6, maxY: -220.1, maxZ: 200}},
	{Name: "Blueberry Acres", Cords: ZoneCords{minX: -319.6, minY: -220.1, minZ: 0, maxX: 104.5, maxY: 293.3, maxZ: 200}},
	{Name: "Caligula's Palace", Cords: ZoneCords{minX: 2087.3, minY: 1543.2, minZ: -89, maxX: 2437.3, maxY: 1703.2, maxZ: 110.9}},
	{Name: "Caligula's Palace", Cords: ZoneCords{minX: 2137.4, minY: 1703.2, minZ: -89, maxX: 2437.3, maxY: 1783.2, maxZ: 110.9}},
	{Name: "Calton Heights", Cords: ZoneCords{minX: -2274.1, minY: 744.1, minZ: -6.1, maxX: -1982.3, maxY: 1358.9, maxZ: 200}},
	{Name: "Chinatown", Cords: ZoneCords{minX: -2274.1, minY: 578.3, minZ: -7.6, maxX: -2078.6, maxY: 744.1, maxZ: 200}},
	{Name: "City Hall", Cords: ZoneCords{minX: -2867.8, minY: 277.4, minZ: -9.1, maxX: -2593.4, maxY: 458.4, maxZ: 200}},
	{Name: "Come-A-Lot", Cords: ZoneCords{minX: 2087.3, minY: 943.2, minZ: -89, maxX: 2623.1, maxY: 1203.2, maxZ: 110.9}},
	{Name: "Commerce", Cords: ZoneCords{minX: 1323.9, minY: -1842.2, minZ: -89, maxX: 1701.9, maxY: -1722.2, maxZ: 110.9}},
	{Name: "Commerce", Cords: ZoneCords{minX: 1323.9, minY: -1722.2, minZ: -89, maxX: 1440.9, maxY: -1577.5, maxZ: 110.9}},
	{Name: "Commerce", Cords: ZoneCords{minX: 1370.8, minY: -1577.5, minZ: -89, maxX: 1463.9, maxY: -1384.9, maxZ: 110.9}},
	{Name: "Commerce", Cords: ZoneCords{minX: 1463.9, minY: -1577.5, minZ: -89, maxX: 1667.9, maxY: -1430.8, maxZ: 110.9}},
	{Name: "Commerce", Cords: ZoneCords{minX: 1583.5, minY: -1722.2, minZ: -89, maxX: 1758.9, maxY: -1577.5, maxZ: 110.9}},
	{Name: "Commerce", Cords: ZoneCords{minX: 1667.9, minY: -1577.5, minZ: -89, maxX: 1812.6, maxY: -1430.8, maxZ: 110.9}},
	{Name: "Conference Center", Cords: ZoneCords{minX: 1046.1, minY: -1804.2, minZ: -89, maxX: 1323.9, maxY: -1722.2, maxZ: 110.9}},
	{Name: "Conference Center", Cords: ZoneCords{minX: 1073.2, minY: -1842.2, minZ: -89, maxX: 1323.9, maxY: -1804.2, maxZ: 110.9}},
	{Name: "Cranberry Station", Cords: ZoneCords{minX: -2007.8, minY: 56.3, minZ: 0, maxX: -1922, maxY: 224.7, maxZ: 100}},
	{Name: "Creek", Cords: ZoneCords{minX: 2749.9, minY: 1937.2, minZ: -89, maxX: 2921.6, maxY: 2669.7, maxZ: 110.9}},
	{Name: "Dillimore", Cords: ZoneCords{minX: 580.7, minY: -674.8, minZ: -9.5, maxX: 861, maxY: -404.7, maxZ: 200}},
	{Name: "Doherty", Cords: ZoneCords{minX: -2270, minY: -324.1, minZ: 0, maxX: -1794.9, maxY: -222.5, maxZ: 200}},
	{Name: "Doherty", Cords: ZoneCords{minX: -2173, minY: -222.5, minZ: 0, maxX: -1794.9, maxY: 265.2, maxZ: 200}},
	{Name: "Downtown", Cords: ZoneCords{minX: -1982.3, minY: 744.1, minZ: -6.1, maxX: -1871.7, maxY: 1274.2, maxZ: 200}},
	{Name: "Downtown", Cords: ZoneCords{minX: -1871.7, minY: 1176.4, minZ: -4.5, maxX: -1620.3, maxY: 1274.2, maxZ: 200}},
	{Name: "Downtown", Cords: ZoneCords{minX: -1700, minY: 744.2, minZ: -6.1, maxX: -1580, maxY: 1176.5, maxZ: 200}},
	{Name: "Downtown", Cords: ZoneCords{minX: -1580, minY: 744.2, minZ: -6.1, maxX: -1499.8, maxY: 1025.9, maxZ: 200}},
	{Name: "Downtown", Cords: ZoneCords{minX: -2078.6, minY: 578.3, minZ: -7.6, maxX: -1499.8, maxY: 744.2, maxZ: 200}},
	{Name: "Downtown", Cords: ZoneCords{minX: -1993.2, minY: 265.2, minZ: -9.1, maxX: -1794.9, maxY: 578.3, maxZ: 200}},
	{Name: "Downtown Los Santos", Cords: ZoneCords{minX: 1463.9, minY: -1430.8, minZ: -89, maxX: 1724.7, maxY: -1290.8, maxZ: 110.9}},
	{Name: "Downtown Los Santos", Cords: ZoneCords{minX: 1724.7, minY: -1430.8, minZ: -89, maxX: 1812.6, maxY: -1250.9, maxZ: 110.9}},
	{Name: "Downtown Los Santos", Cords: ZoneCords{minX: 1463.9, minY: -1290.8, minZ: -89, maxX: 1724.7, maxY: -1150.8, maxZ: 110.9}},
	{Name: "Downtown Los Santos", Cords: ZoneCords{minX: 1370.8, minY: -1384.9, minZ: -89, maxX: 1463.9, maxY: -1170.8, maxZ: 110.9}},
	{Name: "Downtown Los Santos", Cords: ZoneCords{minX: 1724.7, minY: -1250.9, minZ: -89, maxX: 1812.6, maxY: -1150.8, maxZ: 110.9}},
	{Name: "Downtown Los Santos", Cords: ZoneCords{minX: 1370.8, minY: -1170.8, minZ: -89, maxX: 1463.9, maxY: -1130.8, maxZ: 110.9}},
	{Name: "Downtown Los Santos", Cords: ZoneCords{minX: 1378.3, minY: -1130.8, minZ: -89, maxX: 1463.9, maxY: -1026.3, maxZ: 110.9}},
	{Name: "Downtown Los Santos", Cords: ZoneCords{minX: 1391, minY: -1026.3, minZ: -89, maxX: 1463.9, maxY: -926.9, maxZ: 110.9}},
	{Name: "Downtown Los Santos", Cords: ZoneCords{minX: 1507.5, minY: -1385.2, minZ: 110.9, maxX: 1582.5, maxY: -1325.3, maxZ: 335.9}},
	{Name: "East Beach", Cords: ZoneCords{minX: 2632.8, minY: -1852.8, minZ: -89, maxX: 2959.3, maxY: -1668.1, maxZ: 110.9}},
	{Name: "East Beach", Cords: ZoneCords{minX: 2632.8, minY: -1668.1, minZ: -89, maxX: 2747.7, maxY: -1393.4, maxZ: 110.9}},
	{Name: "East Beach", Cords: ZoneCords{minX: 2747.7, minY: -1668.1, minZ: -89, maxX: 2959.3, maxY: -1498.6, maxZ: 110.9}},
	{Name: "East Beach", Cords: ZoneCords{minX: 2747.7, minY: -1498.6, minZ: -89, maxX: 2959.3, maxY: -1120, maxZ: 110.9}},
	{Name: "East Los Santos", Cords: ZoneCords{minX: 2421, minY: -1628.5, minZ: -89, maxX: 2632.8, maxY: -1454.3, maxZ: 110.9}},
	{Name: "East Los Santos", Cords: ZoneCords{minX: 2222.5, minY: -1628.5, minZ: -89, maxX: 2421, maxY: -1494, maxZ: 110.9}},
	{Name: "East Los Santos", Cords: ZoneCords{minX: 2266.2, minY: -1494, minZ: -89, maxX: 2381.6, maxY: -1372, maxZ: 110.9}},
	{Name: "East Los Santos", Cords: ZoneCords{minX: 2381.6, minY: -1494, minZ: -89, maxX: 2421, maxY: -1454.3, maxZ: 110.9}},
	{Name: "East Los Santos", Cords: ZoneCords{minX: 2281.4, minY: -1372, minZ: -89, maxX: 2381.6, maxY: -1135, maxZ: 110.9}},
	{Name: "East Los Santos", Cords: ZoneCords{minX: 2381.6, minY: -1454.3, minZ: -89, maxX: 2462.1, maxY: -1135, maxZ: 110.9}},
	{Name: "East Los Santos", Cords: ZoneCords{minX: 2462.1, minY: -1454.3, minZ: -89, maxX: 2581.7, maxY: -1135, maxZ: 110.9}},
	{Name: "Easter Basin", Cords: ZoneCords{minX: -1794.9, minY: 249.9, minZ: -9.1, maxX: -1242.9, maxY: 578.3, maxZ: 200}},
	{Name: "Easter Basin", Cords: ZoneCords{minX: -1794.9, minY: -50, minZ: 0, maxX: -1499.8, maxY: 249.9, maxZ: 200}},
	{Name: "Easter Bay Airport", Cords: ZoneCords{minX: -1499.8, minY: -50, minZ: 0, maxX: -1242.9, maxY: 249.9, maxZ: 200}},
	{Name: "Easter Bay Airport", Cords: ZoneCords{minX: -1794.9, minY: -730.1, minZ: -3, maxX: -1213.9, maxY: -50, maxZ: 200}},
	{Name: "Easter Bay Airport", Cords: ZoneCords{minX: -1213.9, minY: -730.1, minZ: 0, maxX: -1132.8, maxY: -50, maxZ: 200}},
	{Name: "Easter Bay Airport", Cords: ZoneCords{minX: -1242.9, minY: -50, minZ: 0, maxX: -1213.9, maxY: 578.3, maxZ: 200}},
	{Name: "Easter Bay Airport", Cords: ZoneCords{minX: -1213.9, minY: -50, minZ: -4.5, maxX: -947.9, maxY: 578.3, maxZ: 200}},
	{Name: "Easter Bay Airport", Cords: ZoneCords{minX: -1315.4, minY: -405.3, minZ: 15.4, maxX: -1264.4, maxY: -209.5, maxZ: 25.4}},
	{Name: "Easter Bay Airport", Cords: ZoneCords{minX: -1354.3, minY: -287.3, minZ: 15.4, maxX: -1315.4, maxY: -209.5, maxZ: 25.4}},
	{Name: "Easter Bay Airport", Cords: ZoneCords{minX: -1490.3, minY: -209.5, minZ: 15.4, maxX: -1264.4, maxY: -148.3, maxZ: 25.4}},
	{Name: "Easter Bay Chemicals", Cords: ZoneCords{minX: -1132.8, minY: -768, minZ: 0, maxX: -956.4, maxY: -578.1, maxZ: 200}},
	{Name: "Easter Bay Chemicals", Cords: ZoneCords{minX: -1132.8, minY: -787.3, minZ: 0, maxX: -956.4, maxY: -768, maxZ: 200}},
	{Name: "El Castillo del Diablo", Cords: ZoneCords{minX: -464.5, minY: 2217.6, minZ: 0, maxX: -208.5, maxY: 2580.3, maxZ: 200}},
	{Name: "El Castillo del Diablo", Cords: ZoneCords{minX: -208.5, minY: 2123, minZ: -7.6, maxX: 114, maxY: 2337.1, maxZ: 200}},
	{Name: "El Castillo del Diablo", Cords: ZoneCords{minX: -208.5, minY: 2337.1, minZ: 0, maxX: 8.4, maxY: 2487.1, maxZ: 200}},
	{Name: "El Corona", Cords: ZoneCords{minX: 1812.6, minY: -2179.2, minZ: -89, maxX: 1970.6, maxY: -1852.8, maxZ: 110.9}},
	{Name: "El Corona", Cords: ZoneCords{minX: 1692.6, minY: -2179.2, minZ: -89, maxX: 1812.6, maxY: -1842.2, maxZ: 110.9}},
	{Name: "El Quebrados", Cords: ZoneCords{minX: -1645.2, minY: 2498.5, minZ: 0, maxX: -1372.1, maxY: 2777.8, maxZ: 200}},
	{Name: "Esplanade East", Cords: ZoneCords{minX: -1620.3, minY: 1176.5, minZ: -4.5, maxX: -1580, maxY: 1274.2, maxZ: 200}},
	{Name: "Esplanade East", Cords: ZoneCords{minX: -1580, minY: 1025.9, minZ: -6.1, maxX: -1499.8, maxY: 1274.2, maxZ: 200}},
	{Name: "Esplanade East", Cords: ZoneCords{minX: -1499.8, minY: 578.3, minZ: -79.6, maxX: -1339.8, maxY: 1274.2, maxZ: 20.3}},
	{Name: "Esplanade North", Cords: ZoneCords{minX: -2533, minY: 1358.9, minZ: -4.5, maxX: -1996.6, maxY: 1501.2, maxZ: 200}},
	{Name: "Esplanade North", Cords: ZoneCords{minX: -1996.6, minY: 1358.9, minZ: -4.5, maxX: -1524.2, maxY: 1592.5, maxZ: 200}},
	{Name: "Esplanade North", Cords: ZoneCords{minX: -1982.3, minY: 1274.2, minZ: -4.5, maxX: -1524.2, maxY: 1358.9, maxZ: 200}},
	{Name: "Fallen Tree", Cords: ZoneCords{minX: -792.2, minY: -698.5, minZ: -5.3, maxX: -452.4, maxY: -380, maxZ: 200}},
	{Name: "Fallow Bridge", Cords: ZoneCords{minX: 434.3, minY: 366.5, minZ: 0, maxX: 603, maxY: 555.6, maxZ: 200}},
	{Name: "Fern Ridge", Cords: ZoneCords{minX: 508.1, minY: -139.2, minZ: 0, maxX: 1306.6, maxY: 119.5, maxZ: 200}},
	{Name: "Financial", Cords: ZoneCords{minX: -1871.7, minY: 744.1, minZ: -6.1, maxX: -1701.3, maxY: 1176.4, maxZ: 300}},
	{Name: "Fisher's Lagoon", Cords: ZoneCords{minX: 1916.9, minY: -233.3, minZ: -100, maxX: 2131.7, maxY: 13.8, maxZ: 200}},
	{Name: "Flint Intersection", Cords: ZoneCords{minX: -187.7, minY: -1596.7, minZ: -89, maxX: 17, maxY: -1276.6, maxZ: 110.9}},
	{Name: "Flint Range", Cords: ZoneCords{minX: -594.1, minY: -1648.5, minZ: 0, maxX: -187.7, maxY: -1276.6, maxZ: 200}},
	{Name: "Fort Carson", Cords: ZoneCords{minX: -376.2, minY: 826.3, minZ: -3, maxX: 123.7, maxY: 1220.4, maxZ: 200}},
	{Name: "Foster Valley", Cords: ZoneCords{minX: -2270, minY: -430.2, minZ: 0, maxX: -2178.6, maxY: -324.1, maxZ: 200}},
	{Name: "Foster Valley", Cords: ZoneCords{minX: -2178.6, minY: -599.8, minZ: 0, maxX: -1794.9, maxY: -324.1, maxZ: 200}},
	{Name: "Foster Valley", Cords: ZoneCords{minX: -2178.6, minY: -1115.5, minZ: 0, maxX: -1794.9, maxY: -599.8, maxZ: 200}},
	{Name: "Foster Valley", Cords: ZoneCords{minX: -2178.6, minY: -1250.9, minZ: 0, maxX: -1794.9, maxY: -1115.5, maxZ: 200}},
	{Name: "Frederick Bridge", Cords: ZoneCords{minX: 2759.2, minY: 296.5, minZ: 0, maxX: 2774.2, maxY: 594.7, maxZ: 200}},
	{Name: "Gant Bridge", Cords: ZoneCords{minX: -2741.4, minY: 1659.6, minZ: -6.1, maxX: -2616.4, maxY: 2175.1, maxZ: 200}},
	{Name: "Gant Bridge", Cords: ZoneCords{minX: -2741, minY: 1490.4, minZ: -6.1, maxX: -2616.4, maxY: 1659.6, maxZ: 200}},
	{Name: "Ganton", Cords: ZoneCords{minX: 2222.5, minY: -1852.8, minZ: -89, maxX: 2632.8, maxY: -1722.3, maxZ: 110.9}},
	{Name: "Ganton", Cords: ZoneCords{minX: 2222.5, minY: -1722.3, minZ: -89, maxX: 2632.8, maxY: -1628.5, maxZ: 110.9}},
	{Name: "Garcia", Cords: ZoneCords{minX: -2411.2, minY: -222.5, minZ: 0, maxX: -2173, maxY: 265.2, maxZ: 200}},
	{Name: "Garcia", Cords: ZoneCords{minX: -2395.1, minY: -222.5, minZ: -5.3, maxX: -2354, maxY: -204.7, maxZ: 200}},
	{Name: "Garver Bridge", Cords: ZoneCords{minX: -1339.8, minY: 828.1, minZ: -89, maxX: -1213.9, maxY: 1057, maxZ: 110.9}},
	{Name: "Garver Bridge", Cords: ZoneCords{minX: -1213.9, minY: 950, minZ: -89, maxX: -1087.9, maxY: 1178.9, maxZ: 110.9}},
	{Name: "Garver Bridge", Cords: ZoneCords{minX: -1499.8, minY: 696.4, minZ: -179.6, maxX: -1339.8, maxY: 925.3, maxZ: 20.3}},
	{Name: "Glen Park", Cords: ZoneCords{minX: 1812.6, minY: -1449.6, minZ: -89, maxX: 1996.9, maxY: -1350.7, maxZ: 110.9}},
	{Name: "Glen Park", Cords: ZoneCords{minX: 1812.6, minY: -1100.8, minZ: -89, maxX: 1994.3, maxY: -973.3, maxZ: 110.9}},
	{Name: "Glen Park", Cords: ZoneCords{minX: 1812.6, minY: -1350.7, minZ: -89, maxX: 2056.8, maxY: -1100.8, maxZ: 110.9}},
	{Name: "Green Palms", Cords: ZoneCords{minX: 176.5, minY: 1305.4, minZ: -3, maxX: 338.6, maxY: 1520.7, maxZ: 200}},
	{Name: "Greenglass College", Cords: ZoneCords{minX: 964.3, minY: 1044.6, minZ: -89, maxX: 1197.3, maxY: 1203.2, maxZ: 110.9}},
	{Name: "Greenglass College", Cords: ZoneCords{minX: 964.3, minY: 930.8, minZ: -89, maxX: 1166.5, maxY: 1044.6, maxZ: 110.9}},
	{Name: "Hampton Barns", Cords: ZoneCords{minX: 603, minY: 264.3, minZ: 0, maxX: 761.9, maxY: 366.5, maxZ: 200}},
	{Name: "Hankypanky Point", Cords: ZoneCords{minX: 2576.9, minY: 62.1, minZ: 0, maxX: 2759.2, maxY: 385.5, maxZ: 200}},
	{Name: "Harry Gold Parkway", Cords: ZoneCords{minX: 1777.3, minY: 863.2, minZ: -89, maxX: 1817.3, maxY: 2342.8, maxZ: 110.9}},
	{Name: "Hashbury", Cords: ZoneCords{minX: -2593.4, minY: -222.5, minZ: 0, maxX: -2411.2, maxY: 54.7, maxZ: 200}},
	{Name: "Hilltop Farm", Cords: ZoneCords{minX: 967.3, minY: -450.3, minZ: -3, maxX: 1176.7, maxY: -217.9, maxZ: 200}},
	{Name: "Hunter Quarry", Cords: ZoneCords{minX: 337.2, minY: 710.8, minZ: -115.2, maxX: 860.5, maxY: 1031.7, maxZ: 203.7}},
	{Name: "Idlewood", Cords: ZoneCords{minX: 1812.6, minY: -1852.8, minZ: -89, maxX: 1971.6, maxY: -1742.3, maxZ: 110.9}},
	{Name: "Idlewood", Cords: ZoneCords{minX: 1812.6, minY: -1742.3, minZ: -89, maxX: 1951.6, maxY: -1602.3, maxZ: 110.9}},
	{Name: "Idlewood", Cords: ZoneCords{minX: 1951.6, minY: -1742.3, minZ: -89, maxX: 2124.6, maxY: -1602.3, maxZ: 110.9}},
	{Name: "Idlewood", Cords: ZoneCords{minX: 1812.6, minY: -1602.3, minZ: -89, maxX: 2124.6, maxY: -1449.6, maxZ: 110.9}},
	{Name: "Idlewood", Cords: ZoneCords{minX: 2124.6, minY: -1742.3, minZ: -89, maxX: 2222.5, maxY: -1494, maxZ: 110.9}},
	{Name: "Idlewood", Cords: ZoneCords{minX: 1971.6, minY: -1852.8, minZ: -89, maxX: 2222.5, maxY: -1742.3, maxZ: 110.9}},
	{Name: "Jefferson", Cords: ZoneCords{minX: 1996.9, minY: -1449.6, minZ: -89, maxX: 2056.8, maxY: -1350.7, maxZ: 110.9}},
	{Name: "Jefferson", Cords: ZoneCords{minX: 2124.6, minY: -1494, minZ: -89, maxX: 2266.2, maxY: -1449.6, maxZ: 110.9}},
	{Name: "Jefferson", Cords: ZoneCords{minX: 2056.8, minY: -1372, minZ: -89, maxX: 2281.4, maxY: -1210.7, maxZ: 110.9}},
	{Name: "Jefferson", Cords: ZoneCords{minX: 2056.8, minY: -1210.7, minZ: -89, maxX: 2185.3, maxY: -1126.3, maxZ: 110.9}},
	{Name: "Jefferson", Cords: ZoneCords{minX: 2185.3, minY: -1210.7, minZ: -89, maxX: 2281.4, maxY: -1154.5, maxZ: 110.9}},
	{Name: "Jefferson", Cords: ZoneCords{minX: 2056.8, minY: -1449.6, minZ: -89, maxX: 2266.2, maxY: -1372, maxZ: 110.9}},
	{Name: "Julius Thruway East", Cords: ZoneCords{minX: 2623.1, minY: 943.2, minZ: -89, maxX: 2749.9, maxY: 1055.9, maxZ: 110.9}},
	{Name: "Julius Thruway East", Cords: ZoneCords{minX: 2685.1, minY: 1055.9, minZ: -89, maxX: 2749.9, maxY: 2626.5, maxZ: 110.9}},
	{Name: "Julius Thruway East", Cords: ZoneCords{minX: 2536.4, minY: 2442.5, minZ: -89, maxX: 2685.1, maxY: 2542.5, maxZ: 110.9}},
	{Name: "Julius Thruway East", Cords: ZoneCords{minX: 2625.1, minY: 2202.7, minZ: -89, maxX: 2685.1, maxY: 2442.5, maxZ: 110.9}},
	{Name: "Julius Thruway North", Cords: ZoneCords{minX: 2498.2, minY: 2542.5, minZ: -89, maxX: 2685.1, maxY: 2626.5, maxZ: 110.9}},
	{Name: "Julius Thruway North", Cords: ZoneCords{minX: 2237.4, minY: 2542.5, minZ: -89, maxX: 2498.2, maxY: 2663.1, maxZ: 110.9}},
	{Name: "Julius Thruway North", Cords: ZoneCords{minX: 2121.4, minY: 2508.2, minZ: -89, maxX: 2237.4, maxY: 2663.1, maxZ: 110.9}},
	{Name: "Julius Thruway North", Cords: ZoneCords{minX: 1938.8, minY: 2508.2, minZ: -89, maxX: 2121.4, maxY: 2624.2, maxZ: 110.9}},
	{Name: "Julius Thruway North", Cords: ZoneCords{minX: 1534.5, minY: 2433.2, minZ: -89, maxX: 1848.4, maxY: 2583.2, maxZ: 110.9}},
	{Name: "Julius Thruway North", Cords: ZoneCords{minX: 1848.4, minY: 2478.4, minZ: -89, maxX: 1938.8, maxY: 2553.4, maxZ: 110.9}},
	{Name: "Julius Thruway North", Cords: ZoneCords{minX: 1704.5, minY: 2342.8, minZ: -89, maxX: 1848.4, maxY: 2433.2, maxZ: 110.9}},
	{Name: "Julius Thruway North", Cords: ZoneCords{minX: 1377.3, minY: 2433.2, minZ: -89, maxX: 1534.5, maxY: 2507.2, maxZ: 110.9}},
	{Name: "Julius Thruway South", Cords: ZoneCords{minX: 1457.3, minY: 823.2, minZ: -89, maxX: 2377.3, maxY: 863.2, maxZ: 110.9}},
	{Name: "Julius Thruway South", Cords: ZoneCords{minX: 2377.3, minY: 788.8, minZ: -89, maxX: 2537.3, maxY: 897.9, maxZ: 110.9}},
	{Name: "Julius Thruway West", Cords: ZoneCords{minX: 1197.3, minY: 1163.3, minZ: -89, maxX: 1236.6, maxY: 2243.2, maxZ: 110.9}},
	{Name: "Julius Thruway West", Cords: ZoneCords{minX: 1236.6, minY: 2142.8, minZ: -89, maxX: 1297.4, maxY: 2243.2, maxZ: 110.9}},
	{Name: "Juniper Hill", Cords: ZoneCords{minX: -2533, minY: 578.3, minZ: -7.6, maxX: -2274.1, maxY: 968.3, maxZ: 200}},
	{Name: "Juniper Hollow", Cords: ZoneCords{minX: -2533, minY: 968.3, minZ: -6.1, maxX: -2274.1, maxY: 1358.9, maxZ: 200}},
	{Name: "K.A.C.C. Military Fuels", Cords: ZoneCords{minX: 2498.2, minY: 2626.5, minZ: -89, maxX: 2749.9, maxY: 2861.5, maxZ: 110.9}},
	{Name: "Kincaid Bridge", Cords: ZoneCords{minX: -1339.8, minY: 599.2, minZ: -89, maxX: -1213.9, maxY: 828.1, maxZ: 110.9}},
	{Name: "Kincaid Bridge", Cords: ZoneCords{minX: -1213.9, minY: 721.1, minZ: -89, maxX: -1087.9, maxY: 950, maxZ: 110.9}},
	{Name: "Kincaid Bridge", Cords: ZoneCords{minX: -1087.9, minY: 855.3, minZ: -89, maxX: -961.9, maxY: 986.2, maxZ: 110.9}},
	{Name: "King's", Cords: ZoneCords{minX: -2329.3, minY: 458.4, minZ: -7.6, maxX: -1993.2, maxY: 578.3, maxZ: 200}},
	{Name: "King's", Cords: ZoneCords{minX: -2411.2, minY: 265.2, minZ: -9.1, maxX: -1993.2, maxY: 373.5, maxZ: 200}},
	{Name: "King's", Cords: ZoneCords{minX: -2253.5, minY: 373.5, minZ: -9.1, maxX: -1993.2, maxY: 458.4, maxZ: 200}},
	{Name: "LVA Freight Depot", Cords: ZoneCords{minX: 1457.3, minY: 863.2, minZ: -89, maxX: 1777.4, maxY: 1143.2, maxZ: 110.9}},
	{Name: "LVA Freight Depot", Cords: ZoneCords{minX: 1375.6, minY: 919.4, minZ: -89, maxX: 1457.3, maxY: 1203.2, maxZ: 110.9}},
	{Name: "LVA Freight Depot", Cords: ZoneCords{minX: 1277, minY: 1087.6, minZ: -89, maxX: 1375.6, maxY: 1203.2, maxZ: 110.9}},
	{Name: "LVA Freight Depot", Cords: ZoneCords{minX: 1315.3, minY: 1044.6, minZ: -89, maxX: 1375.6, maxY: 1087.6, maxZ: 110.9}},
	{Name: "LVA Freight Depot", Cords: ZoneCords{minX: 1236.6, minY: 1163.4, minZ: -89, maxX: 1277, maxY: 1203.2, maxZ: 110.9}},
	{Name: "Las Barrancas", Cords: ZoneCords{minX: -926.1, minY: 1398.7, minZ: -3, maxX: -719.2, maxY: 1634.6, maxZ: 200}},
	{Name: "Las Brujas", Cords: ZoneCords{minX: -365.1, minY: 2123, minZ: -3, maxX: -208.5, maxY: 2217.6, maxZ: 200}},
	{Name: "Las Colinas", Cords: ZoneCords{minX: 1994.3, minY: -1100.8, minZ: -89, maxX: 2056.8, maxY: -920.8, maxZ: 110.9}},
	{Name: "Las Colinas", Cords: ZoneCords{minX: 2056.8, minY: -1126.3, minZ: -89, maxX: 2126.8, maxY: -920.8, maxZ: 110.9}},
	{Name: "Las Colinas", Cords: ZoneCords{minX: 2185.3, minY: -1154.5, minZ: -89, maxX: 2281.4, maxY: -934.4, maxZ: 110.9}},
	{Name: "Las Colinas", Cords: ZoneCords{minX: 2126.8, minY: -1126.3, minZ: -89, maxX: 2185.3, maxY: -934.4, maxZ: 110.9}},
	{Name: "Las Colinas", Cords: ZoneCords{minX: 2747.7, minY: -1120, minZ: -89, maxX: 2959.3, maxY: -945, maxZ: 110.9}},
	{Name: "Las Colinas", Cords: ZoneCords{minX: 2632.7, minY: -1135, minZ: -89, maxX: 2747.7, maxY: -945, maxZ: 110.9}},
	{Name: "Las Colinas", Cords: ZoneCords{minX: 2281.4, minY: -1135, minZ: -89, maxX: 2632.7, maxY: -945, maxZ: 110.9}},
	{Name: "Las Payasadas", Cords: ZoneCords{minX: -354.3, minY: 2580.3, minZ: 2, maxX: -133.6, maxY: 2816.8, maxZ: 200}},
	{Name: "Las Venturas Airport", Cords: ZoneCords{minX: 1236.6, minY: 1203.2, minZ: -89, maxX: 1457.3, maxY: 1883.1, maxZ: 110.9}},
	{Name: "Las Venturas Airport", Cords: ZoneCords{minX: 1457.3, minY: 1203.2, minZ: -89, maxX: 1777.3, maxY: 1883.1, maxZ: 110.9}},
	{Name: "Las Venturas Airport", Cords: ZoneCords{minX: 1457.3, minY: 1143.2, minZ: -89, maxX: 1777.4, maxY: 1203.2, maxZ: 110.9}},
	{Name: "Las Venturas Airport", Cords: ZoneCords{minX: 1515.8, minY: 1586.4, minZ: -12.5, maxX: 1729.9, maxY: 1714.5, maxZ: 87.5}},
	{Name: "Last Dime Motel", Cords: ZoneCords{minX: 1823, minY: 596.3, minZ: -89, maxX: 1997.2, maxY: 823.2, maxZ: 110.9}},
	{Name: "Leafy Hollow", Cords: ZoneCords{minX: -1166.9, minY: -1856, minZ: 0, maxX: -815.6, maxY: -1602, maxZ: 200}},
	{Name: "Liberty City", Cords: ZoneCords{minX: -1000, minY: 400, minZ: 1300, maxX: -700, maxY: 600, maxZ: 1400}},
	{Name: "Lil' Probe Inn", Cords: ZoneCords{minX: -90.2, minY: 1286.8, minZ: -3, maxX: 153.8, maxY: 1554.1, maxZ: 200}},
	{Name: "Linden Side", Cords: ZoneCords{minX: 2749.9, minY: 943.2, minZ: -89, maxX: 2923.3, maxY: 1198.9, maxZ: 110.9}},
	{Name: "Linden Station", Cords: ZoneCords{minX: 2749.9, minY: 1198.9, minZ: -89, maxX: 2923.3, maxY: 1548.9, maxZ: 110.9}},
	{Name: "Linden Station", Cords: ZoneCords{minX: 2811.2, minY: 1229.5, minZ: -39.5, maxX: 2861.2, maxY: 1407.5, maxZ: 60.4}},
	{Name: "Little Mexico", Cords: ZoneCords{minX: 1701.9, minY: -1842.2, minZ: -89, maxX: 1812.6, maxY: -1722.2, maxZ: 110.9}},
	{Name: "Little Mexico", Cords: ZoneCords{minX: 1758.9, minY: -1722.2, minZ: -89, maxX: 1812.6, maxY: -1577.5, maxZ: 110.9}},
	{Name: "Los Flores", Cords: ZoneCords{minX: 2581.7, minY: -1454.3, minZ: -89, maxX: 2632.8, maxY: -1393.4, maxZ: 110.9}},
	{Name: "Los Flores", Cords: ZoneCords{minX: 2581.7, minY: -1393.4, minZ: -89, maxX: 2747.7, maxY: -1135, maxZ: 110.9}},
	{Name: "Los Santos International", Cords: ZoneCords{minX: 1249.6, minY: -2394.3, minZ: -89, maxX: 1852, maxY: -2179.2, maxZ: 110.9}},
	{Name: "Los Santos International", Cords: ZoneCords{minX: 1852, minY: -2394.3, minZ: -89, maxX: 2089, maxY: -2179.2, maxZ: 110.9}},
	{Name: "Los Santos International", Cords: ZoneCords{minX: 1382.7, minY: -2730.8, minZ: -89, maxX: 2201.8, maxY: -2394.3, maxZ: 110.9}},
	{Name: "Los Santos International", Cords: ZoneCords{minX: 1974.6, minY: -2394.3, minZ: -39, maxX: 2089, maxY: -2256.5, maxZ: 60.9}},
	{Name: "Los Santos International", Cords: ZoneCords{minX: 1400.9, minY: -2669.2, minZ: -39, maxX: 2189.8, maxY: -2597.2, maxZ: 60.9}},
	{Name: "Los Santos International", Cords: ZoneCords{minX: 2051.6, minY: -2597.2, minZ: -39, maxX: 2152.4, maxY: -2394.3, maxZ: 60.9}},
	{Name: "Marina", Cords: ZoneCords{minX: 647.7, minY: -1804.2, minZ: -89, maxX: 851.4, maxY: -1577.5, maxZ: 110.9}},
	{Name: "Marina", Cords: ZoneCords{minX: 647.7, minY: -1577.5, minZ: -89, maxX: 807.9, maxY: -1416.2, maxZ: 110.9}},
	{Name: "Marina", Cords: ZoneCords{minX: 807.9, minY: -1577.5, minZ: -89, maxX: 926.9, maxY: -1416.2, maxZ: 110.9}},
	{Name: "Market", Cords: ZoneCords{minX: 787.4, minY: -1416.2, minZ: -89, maxX: 1072.6, maxY: -1310.2, maxZ: 110.9}},
	{Name: "Market", Cords: ZoneCords{minX: 952.6, minY: -1310.2, minZ: -89, maxX: 1072.6, maxY: -1130.8, maxZ: 110.9}},
	{Name: "Market", Cords: ZoneCords{minX: 1072.6, minY: -1416.2, minZ: -89, maxX: 1370.8, maxY: -1130.8, maxZ: 110.9}},
	{Name: "Market", Cords: ZoneCords{minX: 926.9, minY: -1577.5, minZ: -89, maxX: 1370.8, maxY: -1416.2, maxZ: 110.9}},
	{Name: "Market Station", Cords: ZoneCords{minX: 787.4, minY: -1410.9, minZ: -34.1, maxX: 866, maxY: -1310.2, maxZ: 65.8}},
	{Name: "Martin Bridge", Cords: ZoneCords{minX: -222.1, minY: 293.3, minZ: 0, maxX: -122.1, maxY: 476.4, maxZ: 200}},
	{Name: "Missionary Hill", Cords: ZoneCords{minX: -2994.4, minY: -811.2, minZ: 0, maxX: -2178.6, maxY: -430.2, maxZ: 200}},
	{Name: "Montgomery", Cords: ZoneCords{minX: 1119.5, minY: 119.5, minZ: -3, maxX: 1451.4, maxY: 493.3, maxZ: 200}},
	{Name: "Montgomery", Cords: ZoneCords{minX: 1451.4, minY: 347.4, minZ: -6.1, maxX: 1582.4, maxY: 420.8, maxZ: 200}},
	{Name: "Montgomery Intersection", Cords: ZoneCords{minX: 1546.6, minY: 208.1, minZ: 0, maxX: 1745.8, maxY: 347.4, maxZ: 200}},
	{Name: "Montgomery Intersection", Cords: ZoneCords{minX: 1582.4, minY: 347.4, minZ: 0, maxX: 1664.6, maxY: 401.7, maxZ: 200}},
	{Name: "Mulholland", Cords: ZoneCords{minX: 1414, minY: -768, minZ: -89, maxX: 1667.6, maxY: -452.4, maxZ: 110.9}},
	{Name: "Mulholland", Cords: ZoneCords{minX: 1281.1, minY: -452.4, minZ: -89, maxX: 1641.1, maxY: -290.9, maxZ: 110.9}},
	{Name: "Mulholland", Cords: ZoneCords{minX: 1269.1, minY: -768, minZ: -89, maxX: 1414, maxY: -452.4, maxZ: 110.9}},
	{Name: "Mulholland", Cords: ZoneCords{minX: 1357, minY: -926.9, minZ: -89, maxX: 1463.9, maxY: -768, maxZ: 110.9}},
	{Name: "Mulholland", Cords: ZoneCords{minX: 1318.1, minY: -910.1, minZ: -89, maxX: 1357, maxY: -768, maxZ: 110.9}},
	{Name: "Mulholland", Cords: ZoneCords{minX: 1169.1, minY: -910.1, minZ: -89, maxX: 1318.1, maxY: -768, maxZ: 110.9}},
	{Name: "Mulholland", Cords: ZoneCords{minX: 768.6, minY: -954.6, minZ: -89, maxX: 952.6, maxY: -860.6, maxZ: 110.9}},
	{Name: "Mulholland", Cords: ZoneCords{minX: 687.8, minY: -860.6, minZ: -89, maxX: 911.8, maxY: -768, maxZ: 110.9}},
	{Name: "Mulholland", Cords: ZoneCords{minX: 737.5, minY: -768, minZ: -89, maxX: 1142.2, maxY: -674.8, maxZ: 110.9}},
	{Name: "Mulholland", Cords: ZoneCords{minX: 1096.4, minY: -910.1, minZ: -89, maxX: 1169.1, maxY: -768, maxZ: 110.9}},
	{Name: "Mulholland", Cords: ZoneCords{minX: 952.6, minY: -937.1, minZ: -89, maxX: 1096.4, maxY: -860.6, maxZ: 110.9}},
	{Name: "Mulholland", Cords: ZoneCords{minX: 911.8, minY: -860.6, minZ: -89, maxX: 1096.4, maxY: -768, maxZ: 110.9}},
	{Name: "Mulholland", Cords: ZoneCords{minX: 861, minY: -674.8, minZ: -89, maxX: 1156.5, maxY: -600.8, maxZ: 110.9}},
	{Name: "Mulholland Intersection", Cords: ZoneCords{minX: 1463.9, minY: -1150.8, minZ: -89, maxX: 1812.6, maxY: -768, maxZ: 110.9}},
	{Name: "North Rock", Cords: ZoneCords{minX: 2285.3, minY: -768, minZ: 0, maxX: 2770.5, maxY: -269.7, maxZ: 200}},
	{Name: "Ocean Docks", Cords: ZoneCords{minX: 2373.7, minY: -2697, minZ: -89, maxX: 2809.2, maxY: -2330.4, maxZ: 110.9}},
	{Name: "Ocean Docks", Cords: ZoneCords{minX: 2201.8, minY: -2418.3, minZ: -89, maxX: 2324, maxY: -2095, maxZ: 110.9}},
	{Name: "Ocean Docks", Cords: ZoneCords{minX: 2324, minY: -2302.3, minZ: -89, maxX: 2703.5, maxY: -2145.1, maxZ: 110.9}},
	{Name: "Ocean Docks", Cords: ZoneCords{minX: 2089, minY: -2394.3, minZ: -89, maxX: 2201.8, maxY: -2235.8, maxZ: 110.9}},
	{Name: "Ocean Docks", Cords: ZoneCords{minX: 2201.8, minY: -2730.8, minZ: -89, maxX: 2324, maxY: -2418.3, maxZ: 110.9}},
	{Name: "Ocean Docks", Cords: ZoneCords{minX: 2703.5, minY: -2302.3, minZ: -89, maxX: 2959.3, maxY: -2126.9, maxZ: 110.9}},
	{Name: "Ocean Docks", Cords: ZoneCords{minX: 2324, minY: -2145.1, minZ: -89, maxX: 2703.5, maxY: -2059.2, maxZ: 110.9}},
	{Name: "Ocean Flats", Cords: ZoneCords{minX: -2994.4, minY: 277.4, minZ: -9.1, maxX: -2867.8, maxY: 458.4, maxZ: 200}},
	{Name: "Ocean Flats", Cords: ZoneCords{minX: -2994.4, minY: -222.5, minZ: 0, maxX: -2593.4, maxY: 277.4, maxZ: 200}},
	{Name: "Ocean Flats", Cords: ZoneCords{minX: -2994.4, minY: -430.2, minZ: 0, maxX: -2831.8, maxY: -222.5, maxZ: 200}},
	{Name: "Octane Springs", Cords: ZoneCords{minX: 338.6, minY: 1228.5, minZ: 0, maxX: 664.3, maxY: 1655, maxZ: 200}},
	{Name: "Old Venturas Strip", Cords: ZoneCords{minX: 2162.3, minY: 2012.1, minZ: -89, maxX: 2685.1, maxY: 2202.7, maxZ: 110.9}},
	{Name: "Palisades", Cords: ZoneCords{minX: -2994.4, minY: 458.4, minZ: -6.1, maxX: -2741, maxY: 1339.6, maxZ: 200}},
	{Name: "Palomino Creek", Cords: ZoneCords{minX: 2160.2, minY: -149, minZ: 0, maxX: 2576.9, maxY: 228.3, maxZ: 200}},
	{Name: "Paradiso", Cords: ZoneCords{minX: -2741, minY: 793.4, minZ: -6.1, maxX: -2533, maxY: 1268.4, maxZ: 200}},
	{Name: "Pershing Square", Cords: ZoneCords{minX: 1440.9, minY: -1722.2, minZ: -89, maxX: 1583.5, maxY: -1577.5, maxZ: 110.9}},
	{Name: "Pilgrim", Cords: ZoneCords{minX: 2437.3, minY: 1383.2, minZ: -89, maxX: 2624.4, maxY: 1783.2, maxZ: 110.9}},
	{Name: "Pilgrim", Cords: ZoneCords{minX: 2624.4, minY: 1383.2, minZ: -89, maxX: 2685.1, maxY: 1783.2, maxZ: 110.9}},
	{Name: "Pilson Intersection", Cords: ZoneCords{minX: 1098.3, minY: 2243.2, minZ: -89, maxX: 1377.3, maxY: 2507.2, maxZ: 110.9}},
	{Name: "Pirates in Men's Pants", Cords: ZoneCords{minX: 1817.3, minY: 1469.2, minZ: -89, maxX: 2027.4, maxY: 1703.2, maxZ: 110.9}},
	{Name: "Playa del Seville", Cords: ZoneCords{minX: 2703.5, minY: -2126.9, minZ: -89, maxX: 2959.3, maxY: -1852.8, maxZ: 110.9}},
	{Name: "Prickle Pine", Cords: ZoneCords{minX: 1534.5, minY: 2583.2, minZ: -89, maxX: 1848.4, maxY: 2863.2, maxZ: 110.9}},
	{Name: "Prickle Pine", Cords: ZoneCords{minX: 1117.4, minY: 2507.2, minZ: -89, maxX: 1534.5, maxY: 2723.2, maxZ: 110.9}},
	{Name: "Prickle Pine", Cords: ZoneCords{minX: 1848.4, minY: 2553.4, minZ: -89, maxX: 1938.8, maxY: 2863.2, maxZ: 110.9}},
	{Name: "Prickle Pine", Cords: ZoneCords{minX: 1938.8, minY: 2624.2, minZ: -89, maxX: 2121.4, maxY: 2861.5, maxZ: 110.9}},
	{Name: "Queens", Cords: ZoneCords{minX: -2533, minY: 458.4, minZ: 0, maxX: -2329.3, maxY: 578.3, maxZ: 200}},
	{Name: "Queens", Cords: ZoneCords{minX: -2593.4, minY: 54.7, minZ: 0, maxX: -2411.2, maxY: 458.4, maxZ: 200}},
	{Name: "Queens", Cords: ZoneCords{minX: -2411.2, minY: 373.5, minZ: 0, maxX: -2253.5, maxY: 458.4, maxZ: 200}},
	{Name: "Randolph Industrial Estate", Cords: ZoneCords{minX: 1558, minY: 596.3, minZ: -89, maxX: 1823, maxY: 823.2, maxZ: 110.9}},
	{Name: "Redsands East", Cords: ZoneCords{minX: 1817.3, minY: 2011.8, minZ: -89, maxX: 2106.7, maxY: 2202.7, maxZ: 110.9}},
	{Name: "Redsands East", Cords: ZoneCords{minX: 1817.3, minY: 2202.7, minZ: -89, maxX: 2011.9, maxY: 2342.8, maxZ: 110.9}},
	{Name: "Redsands East", Cords: ZoneCords{minX: 1848.4, minY: 2342.8, minZ: -89, maxX: 2011.9, maxY: 2478.4, maxZ: 110.9}},
	{Name: "Redsands West", Cords: ZoneCords{minX: 1236.6, minY: 1883.1, minZ: -89, maxX: 1777.3, maxY: 2142.8, maxZ: 110.9}},
	{Name: "Redsands West", Cords: ZoneCords{minX: 1297.4, minY: 2142.8, minZ: -89, maxX: 1777.3, maxY: 2243.2, maxZ: 110.9}},
	{Name: "Redsands West", Cords: ZoneCords{minX: 1377.3, minY: 2243.2, minZ: -89, maxX: 1704.5, maxY: 2433.2, maxZ: 110.9}},
	{Name: "Redsands West", Cords: ZoneCords{minX: 1704.5, minY: 2243.2, minZ: -89, maxX: 1777.3, maxY: 2342.8, maxZ: 110.9}},
	{Name: "Regular Tom", Cords: ZoneCords{minX: -405.7, minY: 1712.8, minZ: -3, maxX: -276.7, maxY: 1892.7, maxZ: 200}},
	{Name: "Richman", Cords: ZoneCords{minX: 647.5, minY: -1118.2, minZ: -89, maxX: 787.4, maxY: -954.6, maxZ: 110.9}},
	{Name: "Richman", Cords: ZoneCords{minX: 647.5, minY: -954.6, minZ: -89, maxX: 768.6, maxY: -860.6, maxZ: 110.9}},
	{Name: "Richman", Cords: ZoneCords{minX: 225.1, minY: -1369.6, minZ: -89, maxX: 334.5, maxY: -1292, maxZ: 110.9}},
	{Name: "Richman", Cords: ZoneCords{minX: 225.1, minY: -1292, minZ: -89, maxX: 466.2, maxY: -1235, maxZ: 110.9}},
	{Name: "Richman", Cords: ZoneCords{minX: 72.6, minY: -1404.9, minZ: -89, maxX: 225.1, maxY: -1235, maxZ: 110.9}},
	{Name: "Richman", Cords: ZoneCords{minX: 72.6, minY: -1235, minZ: -89, maxX: 321.3, maxY: -1008.1, maxZ: 110.9}},
	{Name: "Richman", Cords: ZoneCords{minX: 321.3, minY: -1235, minZ: -89, maxX: 647.5, maxY: -1044, maxZ: 110.9}},
	{Name: "Richman", Cords: ZoneCords{minX: 321.3, minY: -1044, minZ: -89, maxX: 647.5, maxY: -860.6, maxZ: 110.9}},
	{Name: "Richman", Cords: ZoneCords{minX: 321.3, minY: -860.6, minZ: -89, maxX: 687.8, maxY: -768, maxZ: 110.9}},
	{Name: "Richman", Cords: ZoneCords{minX: 321.3, minY: -768, minZ: -89, maxX: 700.7, maxY: -674.8, maxZ: 110.9}},
	{Name: "Robada Intersection", Cords: ZoneCords{minX: -1119, minY: 1178.9, minZ: -89, maxX: -862, maxY: 1351.4, maxZ: 110.9}},
	{Name: "Roca Escalante", Cords: ZoneCords{minX: 2237.4, minY: 2202.7, minZ: -89, maxX: 2536.4, maxY: 2542.5, maxZ: 110.9}},
	{Name: "Roca Escalante", Cords: ZoneCords{minX: 2536.4, minY: 2202.7, minZ: -89, maxX: 2625.1, maxY: 2442.5, maxZ: 110.9}},
	{Name: "Rockshore East", Cords: ZoneCords{minX: 2537.3, minY: 676.5, minZ: -89, maxX: 2902.3, maxY: 943.2, maxZ: 110.9}},
	{Name: "Rockshore West", Cords: ZoneCords{minX: 1997.2, minY: 596.3, minZ: -89, maxX: 2377.3, maxY: 823.2, maxZ: 110.9}},
	{Name: "Rockshore West", Cords: ZoneCords{minX: 2377.3, minY: 596.3, minZ: -89, maxX: 2537.3, maxY: 788.8, maxZ: 110.9}},
	{Name: "Rodeo", Cords: ZoneCords{minX: 72.6, minY: -1684.6, minZ: -89, maxX: 225.1, maxY: -1544.1, maxZ: 110.9}},
	{Name: "Rodeo", Cords: ZoneCords{minX: 72.6, minY: -1544.1, minZ: -89, maxX: 225.1, maxY: -1404.9, maxZ: 110.9}},
	{Name: "Rodeo", Cords: ZoneCords{minX: 225.1, minY: -1684.6, minZ: -89, maxX: 312.8, maxY: -1501.9, maxZ: 110.9}},
	{Name: "Rodeo", Cords: ZoneCords{minX: 225.1, minY: -1501.9, minZ: -89, maxX: 334.5, maxY: -1369.6, maxZ: 110.9}},
	{Name: "Rodeo", Cords: ZoneCords{minX: 334.5, minY: -1501.9, minZ: -89, maxX: 422.6, maxY: -1406, maxZ: 110.9}},
	{Name: "Rodeo", Cords: ZoneCords{minX: 312.8, minY: -1684.6, minZ: -89, maxX: 422.6, maxY: -1501.9, maxZ: 110.9}},
	{Name: "Rodeo", Cords: ZoneCords{minX: 422.6, minY: -1684.6, minZ: -89, maxX: 558, maxY: -1570.2, maxZ: 110.9}},
	{Name: "Rodeo", Cords: ZoneCords{minX: 558, minY: -1684.6, minZ: -89, maxX: 647.5, maxY: -1384.9, maxZ: 110.9}},
	{Name: "Rodeo", Cords: ZoneCords{minX: 466.2, minY: -1570.2, minZ: -89, maxX: 558, maxY: -1385, maxZ: 110.9}},
	{Name: "Rodeo", Cords: ZoneCords{minX: 422.6, minY: -1570.2, minZ: -89, maxX: 466.2, maxY: -1406, maxZ: 110.9}},
	{Name: "Rodeo", Cords: ZoneCords{minX: 466.2, minY: -1385, minZ: -89, maxX: 647.5, maxY: -1235, maxZ: 110.9}},
	{Name: "Rodeo", Cords: ZoneCords{minX: 334.5, minY: -1406, minZ: -89, maxX: 466.2, maxY: -1292, maxZ: 110.9}},
	{Name: "Royal Casino", Cords: ZoneCords{minX: 2087.3, minY: 1383.2, minZ: -89, maxX: 2437.3, maxY: 1543.2, maxZ: 110.9}},
	{Name: "San Andreas Sound", Cords: ZoneCords{minX: 2450.3, minY: 385.5, minZ: -100, maxX: 2759.2, maxY: 562.3, maxZ: 200}},
	{Name: "Santa Flora", Cords: ZoneCords{minX: -2741, minY: 458.4, minZ: -7.6, maxX: -2533, maxY: 793.4, maxZ: 200}},
	{Name: "Santa Maria Beach", Cords: ZoneCords{minX: 342.6, minY: -2173.2, minZ: -89, maxX: 647.7, maxY: -1684.6, maxZ: 110.9}},
	{Name: "Santa Maria Beach", Cords: ZoneCords{minX: 72.6, minY: -2173.2, minZ: -89, maxX: 342.6, maxY: -1684.6, maxZ: 110.9}},
	{Name: "Shady Cabin", Cords: ZoneCords{minX: -1632.8, minY: -2263.4, minZ: -3, maxX: -1601.3, maxY: -2231.7, maxZ: 200}},
	{Name: "Shady Creeks", Cords: ZoneCords{minX: -1820.6, minY: -2643.6, minZ: -8, maxX: -1226.7, maxY: -1771.6, maxZ: 200}},
	{Name: "Shady Creeks", Cords: ZoneCords{minX: -2030.1, minY: -2174.8, minZ: -6.1, maxX: -1820.6, maxY: -1771.6, maxZ: 200}},
	{Name: "Sobell Rail Yards", Cords: ZoneCords{minX: 2749.9, minY: 1548.9, minZ: -89, maxX: 2923.3, maxY: 1937.2, maxZ: 110.9}},
	{Name: "Spinybed", Cords: ZoneCords{minX: 2121.4, minY: 2663.1, minZ: -89, maxX: 2498.2, maxY: 2861.5, maxZ: 110.9}},
	{Name: "Starfish Casino", Cords: ZoneCords{minX: 2437.3, minY: 1783.2, minZ: -89, maxX: 2685.1, maxY: 2012.1, maxZ: 110.9}},
	{Name: "Starfish Casino", Cords: ZoneCords{minX: 2437.3, minY: 1858.1, minZ: -39, maxX: 2495, maxY: 1970.8, maxZ: 60.9}},
	{Name: "Starfish Casino", Cords: ZoneCords{minX: 2162.3, minY: 1883.2, minZ: -89, maxX: 2437.3, maxY: 2012.1, maxZ: 110.9}},
	{Name: "Temple", Cords: ZoneCords{minX: 1252.3, minY: -1130.8, minZ: -89, maxX: 1378.3, maxY: -1026.3, maxZ: 110.9}},
	{Name: "Temple", Cords: ZoneCords{minX: 1252.3, minY: -1026.3, minZ: -89, maxX: 1391, maxY: -926.9, maxZ: 110.9}},
	{Name: "Temple", Cords: ZoneCords{minX: 1252.3, minY: -926.9, minZ: -89, maxX: 1357, maxY: -910.1, maxZ: 110.9}},
	{Name: "Temple", Cords: ZoneCords{minX: 952.6, minY: -1130.8, minZ: -89, maxX: 1096.4, maxY: -937.1, maxZ: 110.9}},
	{Name: "Temple", Cords: ZoneCords{minX: 1096.4, minY: -1130.8, minZ: -89, maxX: 1252.3, maxY: -1026.3, maxZ: 110.9}},
	{Name: "Temple", Cords: ZoneCords{minX: 1096.4, minY: -1026.3, minZ: -89, maxX: 1252.3, maxY: -910.1, maxZ: 110.9}},
	{Name: "The Camel's Toe", Cords: ZoneCords{minX: 2087.3, minY: 1203.2, minZ: -89, maxX: 2640.4, maxY: 1383.2, maxZ: 110.9}},
	{Name: "The Clown's Pocket", Cords: ZoneCords{minX: 2162.3, minY: 1783.2, minZ: -89, maxX: 2437.3, maxY: 1883.2, maxZ: 110.9}},
	{Name: "The Emerald Isle", Cords: ZoneCords{minX: 2011.9, minY: 2202.7, minZ: -89, maxX: 2237.4, maxY: 2508.2, maxZ: 110.9}},
	{Name: "The Farm", Cords: ZoneCords{minX: -1209.6, minY: -1317.1, minZ: 114.9, maxX: -908.1, maxY: -787.3, maxZ: 251.9}},
	{Name: "The Four Dragons Casino", Cords: ZoneCords{minX: 1817.3, minY: 863.2, minZ: -89, maxX: 2027.3, maxY: 1083.2, maxZ: 110.9}},
	{Name: "The High Roller", Cords: ZoneCords{minX: 1817.3, minY: 1283.2, minZ: -89, maxX: 2027.3, maxY: 1469.2, maxZ: 110.9}},
	{Name: "The Mako Span", Cords: ZoneCords{minX: 1664.6, minY: 401.7, minZ: 0, maxX: 1785.1, maxY: 567.2, maxZ: 200}},
	{Name: "The Panopticon", Cords: ZoneCords{minX: -947.9, minY: -304.3, minZ: -1.1, maxX: -319.6, maxY: 327, maxZ: 200}},
	{Name: "The Pink Swan", Cords: ZoneCords{minX: 1817.3, minY: 1083.2, minZ: -89, maxX: 2027.3, maxY: 1283.2, maxZ: 110.9}},
	{Name: "The Sherman Dam", Cords: ZoneCords{minX: -968.7, minY: 1929.4, minZ: -3, maxX: -481.1, maxY: 2155.2, maxZ: 200}},
	{Name: "The Strip", Cords: ZoneCords{minX: 2027.4, minY: 863.2, minZ: -89, maxX: 2087.3, maxY: 1703.2, maxZ: 110.9}},
	{Name: "The Strip", Cords: ZoneCords{minX: 2106.7, minY: 1863.2, minZ: -89, maxX: 2162.3, maxY: 2202.7, maxZ: 110.9}},
	{Name: "The Strip", Cords: ZoneCords{minX: 2027.4, minY: 1783.2, minZ: -89, maxX: 2162.3, maxY: 1863.2, maxZ: 110.9}},
	{Name: "The Strip", Cords: ZoneCords{minX: 2027.4, minY: 1703.2, minZ: -89, maxX: 2137.4, maxY: 1783.2, maxZ: 110.9}},
	{Name: "The Visage", Cords: ZoneCords{minX: 1817.3, minY: 1863.2, minZ: -89, maxX: 2106.7, maxY: 2011.8, maxZ: 110.9}},
	{Name: "The Visage", Cords: ZoneCords{minX: 1817.3, minY: 1703.2, minZ: -89, maxX: 2027.4, maxY: 1863.2, maxZ: 110.9}},
	{Name: "Unity Station", Cords: ZoneCords{minX: 1692.6, minY: -1971.8, minZ: -20.4, maxX: 1812.6, maxY: -1932.8, maxZ: 79.5}},
	{Name: "Valle Ocultado", Cords: ZoneCords{minX: -936.6, minY: 2611.4, minZ: 2, maxX: -715.9, maxY: 2847.9, maxZ: 200}},
	{Name: "Verdant Bluffs", Cords: ZoneCords{minX: 930.2, minY: -2488.4, minZ: -89, maxX: 1249.6, maxY: -2006.7, maxZ: 110.9}},
	{Name: "Verdant Bluffs", Cords: ZoneCords{minX: 1073.2, minY: -2006.7, minZ: -89, maxX: 1249.6, maxY: -1842.2, maxZ: 110.9}},
	{Name: "Verdant Bluffs", Cords: ZoneCords{minX: 1249.6, minY: -2179.2, minZ: -89, maxX: 1692.6, maxY: -1842.2, maxZ: 110.9}},
	{Name: "Verdant Meadows", Cords: ZoneCords{minX: 37, minY: 2337.1, minZ: -3, maxX: 435.9, maxY: 2677.9, maxZ: 200}},
	{Name: "Verona Beach", Cords: ZoneCords{minX: 647.7, minY: -2173.2, minZ: -89, maxX: 930.2, maxY: -1804.2, maxZ: 110.9}},
	{Name: "Verona Beach", Cords: ZoneCords{minX: 930.2, minY: -2006.7, minZ: -89, maxX: 1073.2, maxY: -1804.2, maxZ: 110.9}},
	{Name: "Verona Beach", Cords: ZoneCords{minX: 851.4, minY: -1804.2, minZ: -89, maxX: 1046.1, maxY: -1577.5, maxZ: 110.9}},
	{Name: "Verona Beach", Cords: ZoneCords{minX: 1161.5, minY: -1722.2, minZ: -89, maxX: 1323.9, maxY: -1577.5, maxZ: 110.9}},
	{Name: "Verona Beach", Cords: ZoneCords{minX: 1046.1, minY: -1722.2, minZ: -89, maxX: 1161.5, maxY: -1577.5, maxZ: 110.9}},
	{Name: "Vinewood", Cords: ZoneCords{minX: 787.4, minY: -1310.2, minZ: -89, maxX: 952.6, maxY: -1130.8, maxZ: 110.9}},
	{Name: "Vinewood", Cords: ZoneCords{minX: 787.4, minY: -1130.8, minZ: -89, maxX: 952.6, maxY: -954.6, maxZ: 110.9}},
	{Name: "Vinewood", Cords: ZoneCords{minX: 647.5, minY: -1227.2, minZ: -89, maxX: 787.4, maxY: -1118.2, maxZ: 110.9}},
	{Name: "Vinewood", Cords: ZoneCords{minX: 647.7, minY: -1416.2, minZ: -89, maxX: 787.4, maxY: -1227.2, maxZ: 110.9}},
	{Name: "Whitewood Estates", Cords: ZoneCords{minX: 883.3, minY: 1726.2, minZ: -89, maxX: 1098.3, maxY: 2507.2, maxZ: 110.9}},
	{Name: "Whitewood Estates", Cords: ZoneCords{minX: 1098.3, minY: 1726.2, minZ: -89, maxX: 1197.3, maxY: 2243.2, maxZ: 110.9}},
	{Name: "Willowfield", Cords: ZoneCords{minX: 1970.6, minY: -2179.2, minZ: -89, maxX: 2089, maxY: -1852.8, maxZ: 110.9}},
	{Name: "Willowfield", Cords: ZoneCords{minX: 2089, minY: -2235.8, minZ: -89, maxX: 2201.8, maxY: -1989.9, maxZ: 110.9}},
	{Name: "Willowfield", Cords: ZoneCords{minX: 2089, minY: -1989.9, minZ: -89, maxX: 2324, maxY: -1852.8, maxZ: 110.9}},
	{Name: "Willowfield", Cords: ZoneCords{minX: 2201.8, minY: -2095, minZ: -89, maxX: 2324, maxY: -1989.9, maxZ: 110.9}},
	{Name: "Willowfield", Cords: ZoneCords{minX: 2541.7, minY: -1941.4, minZ: -89, maxX: 2703.5, maxY: -1852.8, maxZ: 110.9}},
	{Name: "Willowfield", Cords: ZoneCords{minX: 2324, minY: -2059.2, minZ: -89, maxX: 2541.7, maxY: -1852.8, maxZ: 110.9}},
	{Name: "Willowfield", Cords: ZoneCords{minX: 2541.7, minY: -2059.2, minZ: -89, maxX: 2703.5, maxY: -1941.4, maxZ: 110.9}},
	{Name: "Yellow Bell Station", Cords: ZoneCords{minX: 1377.4, minY: 2600.4, minZ: -21.9, maxX: 1492.4, maxY: 2687.3, maxZ: 78}},
	{Name: "Los Santos", Cords: ZoneCords{minX: 44.6, minY: -2892.9, minZ: -242.9, maxX: 2997, maxY: -768, maxZ: 900}},
	{Name: "Las Venturas", Cords: ZoneCords{minX: 869.4, minY: 596.3, minZ: -242.9, maxX: 2997, maxY: 2993.8, maxZ: 900}},
	{Name: "Bone County", Cords: ZoneCords{minX: -480.5, minY: 596.3, minZ: -242.9, maxX: 869.4, maxY: 2993.8, maxZ: 900}},
	{Name: "Tierra Robada", Cords: ZoneCords{minX: -2997.4, minY: 1659.6, minZ: -242.9, maxX: -480.5, maxY: 2993.8, maxZ: 900}},
	{Name: "Tierra Robada", Cords: ZoneCords{minX: -1213.9, minY: 596.3, minZ: -242.9, maxX: -480.5, maxY: 1659.6, maxZ: 900}},
	{Name: "San Fierro", Cords: ZoneCords{minX: -2997.4, minY: -1115.5, minZ: -242.9, maxX: -1213.9, maxY: 1659.6, maxZ: 900}},
	{Name: "Red County", Cords: ZoneCords{minX: -1213.9, minY: -768, minZ: -242.9, maxX: 2997, maxY: 596.3, maxZ: 900}},
	{Name: "Flint County", Cords: ZoneCords{minX: -1213.9, minY: -2892.9, minZ: -242.9, maxX: 44.6, maxY: -768, maxZ: 900}},
	{Name: "Whetstone", Cords: ZoneCords{minX: -2997.4, minY: -2892.9, minZ: -242.9, maxX: -1213.9, maxY: -1115.5, maxZ: 900}},
}

type Global struct {
	Weapons  map[int]string
	Vehicles []string
	Zones    []Zone
}

var Globals Global = Global{
	Weapons:  weaponNames,
	Vehicles: vehicleModelNames,
	Zones:    zones,
}
