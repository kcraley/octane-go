package callofduty

import (
	"errors"
	"strings"
)

// WeaponMapping defines a weapon and the various
// mapping between the friendly and backend names
// and categories.
type WeaponMapping struct {
	BackendName      string
	FriendlyName     string
	BackendCategory  string
	FriendlyCategory string
}

var (
	// weaponMappings is a mapping between friendly
	// weapon name to backend names used by the Call
	// of Duty APIs
	weaponMappings = map[string]WeaponMapping{
		"ak-47": WeaponMapping{
			BackendName:      "iw8_ar_akilo47",
			FriendlyName:     "AK-47",
			BackendCategory:  "weapon_assult_rifle",
			FriendlyCategory: "assault rifle",
		},
		"kilo-141": WeaponMapping{
			BackendName:      "iw8_ar_kilo433",
			FriendlyName:     "Kilo-141",
			BackendCategory:  "weapon_assult_rifle",
			FriendlyCategory: "assault rifle",
		},
		"m13": WeaponMapping{
			BackendName:      "iw8_ar_mcharlie",
			FriendlyName:     "M13",
			BackendCategory:  "weapon_assult_rifle",
			FriendlyCategory: "assault rifle",
		},
		"fal": WeaponMapping{
			BackendName:      "iw8_ar_falima",
			FriendlyName:     "FAL",
			BackendCategory:  "weapon_assult_rifle",
			FriendlyCategory: "assault rifle",
		},
		"oden": WeaponMapping{
			BackendName:      "iw8_ar_asierra12",
			FriendlyName:     "Oden",
			BackendCategory:  "weapon_assult_rifle",
			FriendlyCategory: "assault rifle",
		},
		"mp7": WeaponMapping{
			BackendName:      "iw8_sm_mpapa7",
			FriendlyName:     "MP7",
			BackendCategory:  "weapon_smg",
			FriendlyCategory: "submachine guns",
		},
		"aug": WeaponMapping{
			BackendName:      "iw8_sm_augolf",
			FriendlyName:     "AUG",
			BackendCategory:  "weapon_smg",
			FriendlyCategory: "submachine guns",
		},
		"uzi": WeaponMapping{
			BackendName:      "iw8_sm_uzulu",
			FriendlyName:     "Uzi",
			BackendCategory:  "weapon_smg",
			FriendlyCategory: "submachine guns",
		},
		"model 680": WeaponMapping{
			BackendName:      "iw8_sh_romeo870",
			FriendlyName:     "Model 680",
			BackendCategory:  "weapon_shotgun",
			FriendlyCategory: "shotgun",
		},
		"725": WeaponMapping{
			BackendName:      "iw8_sh_charlie725",
			FriendlyName:     "725",
			BackendCategory:  "weapon_shotgun",
			FriendlyCategory: "shotgun",
		},
		"origin 12": WeaponMapping{
			BackendName:      "iw8_sh_oscar12",
			FriendlyName:     "Origin 12",
			BackendCategory:  "weapon_shotgun",
			FriendlyCategory: "shotgun",
		},
		"pkg": WeaponMapping{
			BackendName:      "iw8_lm_pkilo",
			FriendlyName:     "PKM",
			BackendCategory:  "weapon_lmg",
			FriendlyCategory: "light machine guns",
		},
		"mg34": WeaponMapping{
			BackendName:      "iw8_lm_mgolf34",
			FriendlyName:     "MG34",
			BackendCategory:  "weapon_lmg",
			FriendlyCategory: "light machine guns",
		},
		"sa87": WeaponMapping{
			BackendName:      "iw8_lm_lima86",
			FriendlyName:     "SA87",
			BackendCategory:  "weapon_lmg",
			FriendlyCategory: "light machine guns",
		},
		"mo juggernaut": WeaponMapping{
			BackendName:      "iw8_lm_dblmg",
			FriendlyName:     "MP Juggernaut",
			BackendCategory:  "",
			FriendlyCategory: "",
		},
		"ebr-14": WeaponMapping{
			BackendName:      "iw8_sn_mike14",
			FriendlyName:     "EBR-14",
			BackendCategory:  "weapon_marksman",
			FriendlyCategory: "marksman rifles",
		},
		"dragunov": WeaponMapping{
			BackendName:      "iw8_sn_delta",
			FriendlyName:     "Dragunov",
			BackendCategory:  "weapon_sniper",
			FriendlyCategory: "sniper",
		},
		"ax-50": WeaponMapping{
			BackendName:      "iw8_sn_alpha50",
			FriendlyName:     "AX-50",
			BackendCategory:  "weapon_sniper",
			FriendlyCategory: "sniper",
		},
		"hdr": WeaponMapping{
			BackendName:      "iw8_sn_hdromeo",
			FriendlyName:     "HDR",
			BackendCategory:  "weapon_sniper",
			FriendlyCategory: "sniper",
		},
		"mk2 carbine": WeaponMapping{
			BackendName:      "iw8_sn_sbeta",
			FriendlyName:     "Mk2 Carbine",
			BackendCategory:  "weapon_marksman",
			FriendlyCategory: "marksman rifles",
		},
		"m19": WeaponMapping{
			BackendName:      "iw8_pi_papa320",
			FriendlyName:     "M19",
			BackendCategory:  "weapon_pistol",
			FriendlyCategory: "pistol",
		},
		".357": WeaponMapping{
			BackendName:      "iw8_pi_cpapa",
			FriendlyName:     ".357",
			BackendCategory:  "weapon_pistol",
			FriendlyCategory: "pistol",
		},
		"rpg-7": WeaponMapping{
			BackendName:      "iw8_la_rpapa7",
			FriendlyName:     "RPG-7",
			BackendCategory:  "weapon_launcher",
			FriendlyCategory: "launcher",
		},
		"jokr": WeaponMapping{
			BackendName:      "iw8_la_juliet",
			FriendlyName:     "JOKR",
			BackendCategory:  "weapon_launcher",
			FriendlyCategory: "launcher",
		},
		"pila": WeaponMapping{
			BackendName:      "iw8_la_gromeo",
			FriendlyName:     "PILA",
			BackendCategory:  "weapon_launcher",
			FriendlyCategory: "launcher",
		},
		"strela-p": WeaponMapping{
			BackendName:      "iw8_la_kgolf",
			FriendlyName:     "Strela-P",
			BackendCategory:  "weapon_launcher",
			FriendlyCategory: "launcher",
		},
		"riot shield": WeaponMapping{
			BackendName:      "iw8_me_riotshield",
			FriendlyName:     "Riot Shield",
			BackendCategory:  "weapon_other",
			FriendlyCategory: "primary Melee",
		},
		"gas grenade": WeaponMapping{
			BackendName:      "equip_gas_grenade",
			FriendlyName:     "Gas Grenade",
			BackendCategory:  "tacticals",
			FriendlyCategory: "tacticals",
		},
		"snapshot grenade": WeaponMapping{
			BackendName:      "equip_snapshot_grenade",
			FriendlyName:     "Snapshot Grenade",
			BackendCategory:  "tacticals",
			FriendlyCategory: "tacticals",
		},
		"decoy grenade": WeaponMapping{
			BackendName:      "equip_decoy",
			FriendlyName:     "Decoy Grenade",
			BackendCategory:  "tacticals",
			FriendlyCategory: "tacticals",
		},
		"smoke grenade": WeaponMapping{
			BackendName:      "equip_smoke",
			FriendlyName:     "Smoke Grenade",
			BackendCategory:  "tacticals",
			FriendlyCategory: "tacticals",
		},
		"stun grenade": WeaponMapping{
			BackendName:      "equip_concussion",
			FriendlyName:     "Stun Grenade",
			BackendCategory:  "tacticals",
			FriendlyCategory: "tacticals",
		},
		"heartbeat sensor": WeaponMapping{
			BackendName:      "equip_hb_sensor",
			FriendlyName:     "Heartbeat Sensor",
			BackendCategory:  "tacticals",
			FriendlyCategory: "tacticals",
		},
		"flash grenade": WeaponMapping{
			BackendName:      "equip_flash",
			FriendlyName:     "Flash Grenade",
			BackendCategory:  "tacticals",
			FriendlyCategory: "tacticals",
		},
		"stim": WeaponMapping{
			BackendName:      "equip_adrenaline",
			FriendlyName:     "Stim",
			BackendCategory:  "tacticals",
			FriendlyCategory: "tacticals",
		},
		"frag grenade": WeaponMapping{
			BackendName:      "equip_frag",
			FriendlyName:     "Frag Grenade",
			BackendCategory:  "lethals",
			FriendlyCategory: "lethals",
		},
		"thermite": WeaponMapping{
			BackendName:      "equip_thermite",
			FriendlyName:     "Thermite",
			BackendCategory:  "lethals",
			FriendlyCategory: "lethals",
		},
		"semtex": WeaponMapping{
			BackendName:      "equip_semtex",
			FriendlyName:     "Semtex",
			BackendCategory:  "lethals",
			FriendlyCategory: "lethals",
		},
		"claymore": WeaponMapping{
			BackendName:      "equip_claymore",
			FriendlyName:     "Claymore",
			BackendCategory:  "lethals",
			FriendlyCategory: "lethals",
		},
		"c4": WeaponMapping{
			BackendName:      "equip_c4",
			FriendlyName:     "C4",
			BackendCategory:  "lethals",
			FriendlyCategory: "lethals",
		},
		"proximity mine": WeaponMapping{
			BackendName:      "equip_at_mine",
			FriendlyName:     "Proximity Mine",
			BackendCategory:  "lethals",
			FriendlyCategory: "lethals",
		},
		"throwing knife": WeaponMapping{
			BackendName:      "equip_throwing_knife",
			FriendlyName:     "Throwing Knife",
			BackendCategory:  "lethals",
			FriendlyCategory: "lethals",
		},
		"molotov cocktail": WeaponMapping{
			BackendName:      "equip_molotov",
			FriendlyName:     "Molotov Cocktail",
			BackendCategory:  "lethals",
			FriendlyCategory: "lethals",
		},
		"combat knife": WeaponMapping{
			BackendName:      "iw8_knife",
			FriendlyName:     "Combat Knife",
			BackendCategory:  "weapon_melee",
			FriendlyCategory: "melee",
		},
		"ram-7": WeaponMapping{
			BackendName:      "iw8_ar_tango21",
			FriendlyName:     "RAM-7",
			BackendCategory:  "weapon_assult_rifle",
			FriendlyCategory: "assault rifle",
		},
		"fr 5.56": WeaponMapping{
			BackendName:      "iw8_ar_falpha",
			FriendlyName:     "FR 5.56",
			BackendCategory:  "weapon_assult_rifle",
			FriendlyCategory: "assault rifle",
		},
		"m4a1": WeaponMapping{
			BackendName:      "iw8_ar_mike4",
			FriendlyName:     "M4A1",
			BackendCategory:  "weapon_assult_rifle",
			FriendlyCategory: "assault rifle",
		},
		"p90": WeaponMapping{
			BackendName:      "iw8_sm_papa90",
			FriendlyName:     "P90",
			BackendCategory:  "weapon_smg",
			FriendlyCategory: "submachine guns",
		},
		"mp5": WeaponMapping{
			BackendName:      "iw8_sm_mpapa5",
			FriendlyName:     "MP5",
			BackendCategory:  "weapon_smg",
			FriendlyCategory: "submachine guns",
		},
		"pp19 bizon": WeaponMapping{
			BackendName:      "iw8_sm_beta",
			FriendlyName:     "PP19 Bizon",
			BackendCategory:  "weapon_smg",
			FriendlyCategory: "submachine guns",
		},
		"r9-0": WeaponMapping{
			BackendName:      "iw8_sh_dpapa12",
			FriendlyName:     "R9-0",
			BackendCategory:  "weapon_shotgun",
			FriendlyCategory: "shotgun",
		},
		"holger-26": WeaponMapping{
			BackendName:      "iw8_lm_mgolf36",
			FriendlyName:     "Holger-26",
			BackendCategory:  "weapon_lmg",
			FriendlyCategory: "light machine guns",
		},
		"kar98k": WeaponMapping{
			BackendName:      "iw8_sn_kilo98",
			FriendlyName:     "Kar98k",
			BackendCategory:  "weapon_marksman",
			FriendlyCategory: "marksman rifles",
		},
		"1911": WeaponMapping{
			BackendName:      "iw8_pi_mike1911",
			FriendlyName:     "1911",
			BackendCategory:  "weapon_pistol",
			FriendlyCategory: "pistol",
		},
		"x16": WeaponMapping{
			BackendName:      "iw8_pi_golf21",
			FriendlyName:     "X16",
			BackendCategory:  "weapon_pistol",
			FriendlyCategory: "pistol",
		},
		".50 gs": WeaponMapping{
			BackendName:      "iw8_pi_decho",
			FriendlyName:     ".50 GS",
			BackendCategory:  "weapon_pistol",
			FriendlyCategory: "pistol",
		},
		"m91": WeaponMapping{
			BackendName:      "iw8_lm_kilo121",
			FriendlyName:     "M91",
			BackendCategory:  "weapon_lmg",
			FriendlyCategory: "light machine guns",
		},
		"fn scar 17": WeaponMapping{
			BackendName:      "iw8_ar_scharlie",
			FriendlyName:     "FN Scar 17",
			BackendCategory:  "weapon_assault_rifle",
			FriendlyCategory: "assault rifle",
		},
		"grau 5.56": WeaponMapping{
			BackendName:      "iw8_ar_sierra552",
			FriendlyName:     "Grau 5.56",
			BackendCategory:  "weapon_assault_rifle",
			FriendlyCategory: "assault rifle",
		},
		"striker 45": WeaponMapping{
			BackendName:      "iw8_sm_smgolf45",
			FriendlyName:     "Striker 45",
			BackendCategory:  "weapon_smg",
			FriendlyCategory: "submachine guns",
		},
		"bruen mk9": WeaponMapping{
			BackendName:      "iw8_lm_mkilo3",
			FriendlyName:     "Bruen MK9",
			BackendCategory:  "weapon_lmg",
			FriendlyCategory: "light machine guns",
		},
		"vlk rogue": WeaponMapping{
			BackendName:      "iw8_sh_mike26",
			FriendlyName:     "VLK Rogue",
			BackendCategory:  "weapon_shotgun",
			FriendlyCategory: "shotgun",
		},
		"crossbow": WeaponMapping{
			BackendName:      "iw8_sn_crossbow",
			FriendlyName:     "Crossbow",
			BackendCategory:  "weapon_marksman",
			FriendlyCategory: "marksman rifles",
		},
		"sks": WeaponMapping{
			BackendName:      "iw8_sn_sksierra",
			FriendlyName:     "SKS",
			BackendCategory:  "weapon_marksman",
			FriendlyCategory: "marksman rifles",
		},
		"cr-56 amax": WeaponMapping{
			BackendName:      "iw8_ar_galima",
			FriendlyName:     "CR-56 AMAX",
			BackendCategory:  "weapon_assault_rifle",
			FriendlyCategory: "assault rifle",
		},
		"fennec mk9": WeaponMapping{
			BackendName:      "iw8_sm_victor",
			FriendlyName:     "Fennec Mk9",
			BackendCategory:  "weapon_smg",
			FriendlyCategory: "submachine guns",
		},
		"rytec amr": WeaponMapping{
			BackendName:      "iw8_sn_xmike109",
			FriendlyName:     "Rytec AMR",
			BackendCategory:  "weapon_sniper",
			FriendlyCategory: "sniper",
		},
		"renetti": WeaponMapping{
			BackendName:      "iw8_pi_mike9",
			FriendlyName:     "Renetti",
			BackendCategory:  "weapon_pistol",
			FriendlyCategory: "pistol",
		},
		"kali sticks": WeaponMapping{
			BackendName:      "iw8_me_akimboblunt",
			FriendlyName:     "Kali Sticks",
			BackendCategory:  "weapon_melee",
			FriendlyCategory: "melee",
		},
		"an-94": WeaponMapping{
			BackendName:      "iw8_ar_anovember94",
			FriendlyName:     "AN-94",
			BackendCategory:  "weapon_assault_rifle",
			FriendlyCategory: "assault rifle",
		},
		"iso": WeaponMapping{
			BackendName:      "iw8_sm_charlie9",
			FriendlyName:     "ISO",
			BackendCategory:  "weapon_smg",
			FriendlyCategory: "submachine guns",
		},
		"dual kodachis": WeaponMapping{
			BackendName:      "iw8_me_akimboblades",
			FriendlyName:     "Dual Kodachis",
			BackendCategory:  "weapon_melee",
			FriendlyCategory: "melee",
		},
	}
)

// GetWeaponMapping returns the weapon mapping given the
// friendly name of the weapon
func GetWeaponMapping(weapon string) (WeaponMapping, error) {
	if mapping, ok := weaponMappings[strings.ToLower(weapon)]; ok {
		return mapping, nil
	}
	return WeaponMapping{}, errors.New("Unable to find weapon mapping")
}

// GetBackendName returns the backend name of the weapon
func (w WeaponMapping) GetBackendName() string {
	return w.BackendName
}

// GetFriendlyName returns the friendly name of the weapon
func (w WeaponMapping) GetFriendlyName() string {
	return w.FriendlyName
}

// GetBackendCategory returns the backend category of the weapon
func (w WeaponMapping) GetBackendCategory() string {
	return w.BackendCategory
}

// GetFriendlyCategory returns the friendly category of the weapon
func (w WeaponMapping) GetFriendlyCategory() string {
	return w.FriendlyCategory
}
