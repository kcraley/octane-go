package callofduty

// Profile is a simplified profile of a player
type Profile struct {
	Title            string  `json:"title"`
	Platform         string  `json:"platform"`
	Username         string  `json:"username"`
	Type             string  `json:"type"`
	Level            float64 `json:"level"`
	MaxLevel         float64 `json:"maxLevel"`
	LevelXpRemainder float64 `json:"levelXpRemainder"`
	LevelXpGained    float64 `json:"levelXpGained"`
	Prestige         float64 `json:"prestige"`
	PrestigeID       float64 `json:"prestigeId"`
	MaxPrestige      float64 `json:"maxPrestige"`
	TotalXp          float64 `json:"totalXp"`
}

// ProfileResponse is the entire response object
// returned when getting a player's profile
type ProfileResponse struct {
	Status string `json:"status"`
	Data   struct {
		Title            string  `json:"title"`
		Platform         string  `json:"platform"`
		Username         string  `json:"username"`
		Type             string  `json:"type"`
		Level            float64 `json:"level"`
		MaxLevel         float64 `json:"maxLevel"`
		LevelXpRemainder float64 `json:"levelXpRemainder"`
		LevelXpGained    float64 `json:"levelXpGained"`
		Prestige         float64 `json:"prestige"`
		PrestigeID       float64 `json:"prestigeId"`
		MaxPrestige      float64 `json:"maxPrestige"`
		TotalXp          float64 `json:"totalXp"`
		ParagonRank      float64 `json:"paragonRank"`
		ParagonID        float64 `json:"paragonId"`
		S                float64 `json:"s"`
		P                float64 `json:"p"`
		Lifetime         struct {
			All struct {
				Properties struct {
					RecordLongestWinStreak float64 `json:"recordLongestWinStreak"`
					RecordXpInAMatch       float64 `json:"recordXpInAMatch"`
					Accuracy               float64 `json:"accuracy"`
					Losses                 float64 `json:"losses"`
					TotalGamesPlayed       float64 `json:"totalGamesPlayed"`
					Score                  float64 `json:"score"`
					WinLossRatio           float64 `json:"winLossRatio"`
					TotalShots             float64 `json:"totalShots"`
					BestScoreXp            float64 `json:"bestScoreXp"`
					GamesPlayed            float64 `json:"gamesPlayed"`
					BestSquardKills        float64 `json:"bestSquardKills"`
					BestSguardWave         float64 `json:"bestSguardWave"`
					BestConfirmed          float64 `json:"bestConfirmed"`
					Deaths                 float64 `json:"deaths"`
					Wins                   float64 `json:"wins"`
					BestSquardCrates       float64 `json:"bestSquardCrates"`
					KdRatio                float64 `json:"kdRatio"`
					BestAssists            float64 `json:"bestAssists"`
					BestFieldgoals         float64 `json:"bestFieldgoals"`
					BestScore              float64 `json:"bestScore"`
					RecordDeathsInAMatch   float64 `json:"recordDeathsInAMatch"`
					ScorePerGame           float64 `json:"scorePerGame"`
					BestSPM                float64 `json:"bestSPM"`
					BestKillChains         float64 `json:"bestKillChains"`
					RecordKillsInAMatch    float64 `json:"recordKillsInAMatch"`
					Suicides               float64 `json:"suicides"`
					WlRatio                float64 `json:"wlRatio"`
					CurrentWinStreak       float64 `json:"currentWinStreak"`
					BestMatchBonusXp       float64 `json:"bestMatchBonusXp"`
					BestMatchXp            float64 `json:"bestMatchXp"`
					BestSguardWeaponLevel  float64 `json:"bestSguardWeaponLevel"`
					BestKD                 float64 `json:"bestKD"`
					Kills                  float64 `json:"kills"`
					BestKillsAsInfected    float64 `json:"bestKillsAsInfected"`
					BestReturns            float64 `json:"bestReturns"`
					BestStabs              float64 `json:"bestStabs"`
					BestKillsAsSurvivor    float64 `json:"bestKillsAsSurvivor"`
					TimePlayedTotal        float64 `json:"timePlayedTotal"`
					BestDestructions       float64 `json:"bestDestructions"`
					Headshots              float64 `json:"headshots"`
					BestRescues            float64 `json:"bestRescues"`
					Assists                float64 `json:"assists"`
					Ties                   float64 `json:"ties"`
					RecordKillStreak       float64 `json:"recordKillStreak"`
					BestPlants             float64 `json:"bestPlants"`
					Misses                 float64 `json:"misses"`
					BestDamage             float64 `json:"bestDamage"`
					BestSetbacks           float64 `json:"bestSetbacks"`
					BestTouchdowns         float64 `json:"bestTouchdowns"`
					ScorePerMinute         float64 `json:"scorePerMinute"`
					BestDeaths             float64 `json:"bestDeaths"`
					BestMedalXp            float64 `json:"bestMedalXp"`
					BestDefends            float64 `json:"bestDefends"`
					BestSquardRevives      float64 `json:"bestSquardRevives"`
					BestKills              float64 `json:"bestKills"`
					BestDefuses            float64 `json:"bestDefuses"`
					BestCaptures           float64 `json:"bestCaptures"`
					Hits                   float64 `json:"hits"`
					BestKillStreak         float64 `json:"bestKillStreak"`
					BestDenied             float64 `json:"bestDenied"`
				} `json:"properties"`
			} `json:"all"`
			Mode struct {
				Gun struct {
					Properties struct {
						Kills          float64 `json:"kills"`
						Score          float64 `json:"score"`
						TimePlayed     float64 `json:"timePlayed"`
						KdRatio        float64 `json:"kdRatio"`
						SetBacks       float64 `json:"setBacks"`
						ScorePerMinute float64 `json:"scorePerMinute"`
						Stabs          float64 `json:"stabs"`
						Deaths         float64 `json:"deaths"`
					} `json:"properties"`
				} `json:"gun"`
				Dom struct {
					Properties struct {
						Kills          float64 `json:"kills"`
						Score          float64 `json:"score"`
						TimePlayed     float64 `json:"timePlayed"`
						KdRatio        float64 `json:"kdRatio"`
						Captures       float64 `json:"captures"`
						Defends        float64 `json:"defends"`
						ScorePerMinute float64 `json:"scorePerMinute"`
						Deaths         float64 `json:"deaths"`
					} `json:"properties"`
				} `json:"dom"`
				War struct {
					Properties struct {
						Kills          float64 `json:"kills"`
						Score          float64 `json:"score"`
						TimePlayed     float64 `json:"timePlayed"`
						KdRatio        float64 `json:"kdRatio"`
						Assists        float64 `json:"assists"`
						ScorePerMinute float64 `json:"scorePerMinute"`
						Deaths         float64 `json:"deaths"`
					} `json:"properties"`
				} `json:"war"`
				Hq struct {
					Properties struct {
						Kills          float64 `json:"kills"`
						Score          float64 `json:"score"`
						TimePlayed     float64 `json:"timePlayed"`
						KdRatio        float64 `json:"kdRatio"`
						Captures       float64 `json:"captures"`
						Defends        float64 `json:"defends"`
						ScorePerMinute float64 `json:"scorePerMinute"`
						Deaths         float64 `json:"deaths"`
					} `json:"properties"`
				} `json:"hq"`
				HcDom struct {
					Properties struct {
						Kills          float64 `json:"kills"`
						Score          float64 `json:"score"`
						TimePlayed     float64 `json:"timePlayed"`
						KdRatio        float64 `json:"kdRatio"`
						Captures       float64 `json:"captures"`
						Defends        float64 `json:"defends"`
						ScorePerMinute float64 `json:"scorePerMinute"`
						Deaths         float64 `json:"deaths"`
					} `json:"properties"`
				} `json:"hc_dom"`
				HcConf struct {
					Properties struct {
						Kills          float64 `json:"kills"`
						Score          float64 `json:"score"`
						TimePlayed     float64 `json:"timePlayed"`
						KdRatio        float64 `json:"kdRatio"`
						Confirms       float64 `json:"confirms"`
						ScorePerMinute float64 `json:"scorePerMinute"`
						Denies         float64 `json:"denies"`
						Deaths         float64 `json:"deaths"`
					} `json:"properties"`
				} `json:"hc_conf"`
				Koth struct {
					Properties struct {
						Kills          float64 `json:"kills"`
						Score          float64 `json:"score"`
						TimePlayed     float64 `json:"timePlayed"`
						KdRatio        float64 `json:"kdRatio"`
						Defends        float64 `json:"defends"`
						ObjTime        float64 `json:"objTime"`
						ScorePerMinute float64 `json:"scorePerMinute"`
						Deaths         float64 `json:"deaths"`
					} `json:"properties"`
				} `json:"koth"`
				Conf struct {
					Properties struct {
						Kills          float64 `json:"kills"`
						Score          float64 `json:"score"`
						TimePlayed     float64 `json:"timePlayed"`
						KdRatio        float64 `json:"kdRatio"`
						Confirms       float64 `json:"confirms"`
						ScorePerMinute float64 `json:"scorePerMinute"`
						Denies         float64 `json:"denies"`
						Deaths         float64 `json:"deaths"`
					} `json:"properties"`
				} `json:"conf"`
				HcHq struct {
					Properties struct {
						Kills          float64 `json:"kills"`
						Score          float64 `json:"score"`
						TimePlayed     float64 `json:"timePlayed"`
						KdRatio        float64 `json:"kdRatio"`
						Captures       float64 `json:"captures"`
						Defends        float64 `json:"defends"`
						ScorePerMinute float64 `json:"scorePerMinute"`
						Deaths         float64 `json:"deaths"`
					} `json:"properties"`
				} `json:"hc_hq"`
				Arena struct {
					Properties struct {
						Kills          float64 `json:"kills"`
						Score          float64 `json:"score"`
						TimePlayed     float64 `json:"timePlayed"`
						Damage         float64 `json:"damage"`
						KdRatio        float64 `json:"kdRatio"`
						Assists        float64 `json:"assists"`
						ScorePerMinute float64 `json:"scorePerMinute"`
						Deaths         float64 `json:"deaths"`
					} `json:"properties"`
				} `json:"arena"`
				BrDmz struct {
					Properties struct {
						Wins           float64 `json:"wins"`
						Kills          float64 `json:"kills"`
						KdRatio        float64 `json:"kdRatio"`
						Downs          float64 `json:"downs"`
						TopTwentyFive  float64 `json:"topTwentyFive"`
						TopTen         float64 `json:"topTen"`
						Contracts      float64 `json:"contracts"`
						Revives        float64 `json:"revives"`
						TopFive        float64 `json:"topFive"`
						Score          float64 `json:"score"`
						TimePlayed     float64 `json:"timePlayed"`
						GamesPlayed    float64 `json:"gamesPlayed"`
						Tokens         float64 `json:"tokens"`
						ScorePerMinute float64 `json:"scorePerMinute"`
						Cash           float64 `json:"cash"`
						Deaths         float64 `json:"deaths"`
					} `json:"properties"`
				} `json:"br_dmz"`
				Br struct {
					Properties struct {
						Wins           float64 `json:"wins"`
						Kills          float64 `json:"kills"`
						KdRatio        float64 `json:"kdRatio"`
						Downs          float64 `json:"downs"`
						TopTwentyFive  float64 `json:"topTwentyFive"`
						TopTen         float64 `json:"topTen"`
						Contracts      float64 `json:"contracts"`
						Revives        float64 `json:"revives"`
						TopFive        float64 `json:"topFive"`
						Score          float64 `json:"score"`
						TimePlayed     float64 `json:"timePlayed"`
						GamesPlayed    float64 `json:"gamesPlayed"`
						Tokens         float64 `json:"tokens"`
						ScorePerMinute float64 `json:"scorePerMinute"`
						Cash           float64 `json:"cash"`
						Deaths         float64 `json:"deaths"`
					} `json:"properties"`
				} `json:"br"`
				Sd struct {
					Properties struct {
						Kills          float64 `json:"kills"`
						Score          float64 `json:"score"`
						TimePlayed     float64 `json:"timePlayed"`
						KdRatio        float64 `json:"kdRatio"`
						Plants         float64 `json:"plants"`
						ScorePerMinute float64 `json:"scorePerMinute"`
						Defuses        float64 `json:"defuses"`
						Deaths         float64 `json:"deaths"`
					} `json:"properties"`
				} `json:"sd"`
				Grnd struct {
					Properties struct {
						Kills          float64 `json:"kills"`
						Score          float64 `json:"score"`
						TimePlayed     float64 `json:"timePlayed"`
						KdRatio        float64 `json:"kdRatio"`
						Defends        float64 `json:"defends"`
						ObjTime        float64 `json:"objTime"`
						ScorePerMinute float64 `json:"scorePerMinute"`
						Deaths         float64 `json:"deaths"`
					} `json:"properties"`
				} `json:"grnd"`
				Cyber struct {
					Properties struct {
						Kills          float64 `json:"kills"`
						Score          float64 `json:"score"`
						TimePlayed     float64 `json:"timePlayed"`
						KdRatio        float64 `json:"kdRatio"`
						Plants         float64 `json:"plants"`
						ScorePerMinute float64 `json:"scorePerMinute"`
						Revives        float64 `json:"revives"`
						Deaths         float64 `json:"deaths"`
					} `json:"properties"`
				} `json:"cyber"`
				HcWar struct {
					Properties struct {
						Kills          float64 `json:"kills"`
						Score          float64 `json:"score"`
						TimePlayed     float64 `json:"timePlayed"`
						KdRatio        float64 `json:"kdRatio"`
						Assists        float64 `json:"assists"`
						ScorePerMinute float64 `json:"scorePerMinute"`
						Deaths         float64 `json:"deaths"`
					} `json:"properties"`
				} `json:"hc_war"`
				BrAll struct {
					Properties struct {
						Wins           float64 `json:"wins"`
						Kills          float64 `json:"kills"`
						KdRatio        float64 `json:"kdRatio"`
						Downs          float64 `json:"downs"`
						TopTwentyFive  float64 `json:"topTwentyFive"`
						TopTen         float64 `json:"topTen"`
						Contracts      float64 `json:"contracts"`
						Revives        float64 `json:"revives"`
						TopFive        float64 `json:"topFive"`
						Score          float64 `json:"score"`
						TimePlayed     float64 `json:"timePlayed"`
						GamesPlayed    float64 `json:"gamesPlayed"`
						Tokens         float64 `json:"tokens"`
						ScorePerMinute float64 `json:"scorePerMinute"`
						Cash           float64 `json:"cash"`
						Deaths         float64 `json:"deaths"`
					} `json:"properties"`
				} `json:"br_all"`
				HcSd struct {
					Properties struct {
						Kills          float64 `json:"kills"`
						Score          float64 `json:"score"`
						TimePlayed     float64 `json:"timePlayed"`
						KdRatio        float64 `json:"kdRatio"`
						Plants         float64 `json:"plants"`
						ScorePerMinute float64 `json:"scorePerMinute"`
						Defuses        float64 `json:"defuses"`
						Deaths         float64 `json:"deaths"`
					} `json:"properties"`
				} `json:"hc_sd"`
				Arm struct {
					Properties struct {
						Kills          float64 `json:"kills"`
						Score          float64 `json:"score"`
						TimePlayed     float64 `json:"timePlayed"`
						KdRatio        float64 `json:"kdRatio"`
						Captures       float64 `json:"captures"`
						Defends        float64 `json:"defends"`
						ScorePerMinute float64 `json:"scorePerMinute"`
						Deaths         float64 `json:"deaths"`
					} `json:"properties"`
				} `json:"arm"`
				HcCyber struct {
					Properties struct {
						Kills          float64 `json:"kills"`
						Score          float64 `json:"score"`
						TimePlayed     float64 `json:"timePlayed"`
						KdRatio        float64 `json:"kdRatio"`
						Plants         float64 `json:"plants"`
						ScorePerMinute float64 `json:"scorePerMinute"`
						Revives        float64 `json:"revives"`
						Deaths         float64 `json:"deaths"`
					} `json:"properties"`
				} `json:"hc_cyber"`
				Infect struct {
					Properties struct {
						Kills          float64 `json:"kills"`
						Score          float64 `json:"score"`
						Infected       float64 `json:"infected"`
						TimePlayed     float64 `json:"timePlayed"`
						KdRatio        float64 `json:"kdRatio"`
						ScorePerMinute float64 `json:"scorePerMinute"`
						Time           float64 `json:"time"`
						Deaths         float64 `json:"deaths"`
					} `json:"properties"`
				} `json:"infect"`
			} `json:"mode"`
			Map struct {
			} `json:"map"`
			ItemData struct {
				WeaponSniper struct {
					Iw8SnAlpha50 struct {
						Properties struct {
							Hits      float64 `json:"hits"`
							Kills     float64 `json:"kills"`
							KdRatio   float64 `json:"kdRatio"`
							Headshots float64 `json:"headshots"`
							Accuracy  float64 `json:"accuracy"`
							Shots     float64 `json:"shots"`
							Deaths    float64 `json:"deaths"`
						} `json:"properties"`
					} `json:"iw8_sn_alpha50"`
					Iw8SnHdromeo struct {
						Properties struct {
							Hits      float64 `json:"hits"`
							Kills     float64 `json:"kills"`
							KdRatio   float64 `json:"kdRatio"`
							Headshots float64 `json:"headshots"`
							Accuracy  float64 `json:"accuracy"`
							Shots     float64 `json:"shots"`
							Deaths    float64 `json:"deaths"`
						} `json:"properties"`
					} `json:"iw8_sn_hdromeo"`
					Iw8SnDelta struct {
						Properties struct {
							Hits      float64 `json:"hits"`
							Kills     float64 `json:"kills"`
							KdRatio   float64 `json:"kdRatio"`
							Headshots float64 `json:"headshots"`
							Accuracy  float64 `json:"accuracy"`
							Shots     float64 `json:"shots"`
							Deaths    float64 `json:"deaths"`
						} `json:"properties"`
					} `json:"iw8_sn_delta"`
					Iw8SnXmike109 struct {
						Properties struct {
							Hits      float64 `json:"hits"`
							Kills     float64 `json:"kills"`
							KdRatio   float64 `json:"kdRatio"`
							Headshots float64 `json:"headshots"`
							Accuracy  float64 `json:"accuracy"`
							Shots     float64 `json:"shots"`
							Deaths    float64 `json:"deaths"`
						} `json:"properties"`
					} `json:"iw8_sn_xmike109"`
				} `json:"weapon_sniper"`
				Tacticals struct {
					EquipGasGrenade struct {
						Properties struct {
							Hits      float64 `json:"hits"`
							Kills     float64 `json:"kills"`
							Shots     float64 `json:"shots"`
							Deaths    float64 `json:"deaths"`
							HeadShots float64 `json:"headShots"`
						} `json:"properties"`
					} `json:"equip_gas_grenade"`
					EquipSnapshotGrenade struct {
						Properties struct {
							Hits      float64 `json:"hits"`
							Kills     float64 `json:"kills"`
							Shots     float64 `json:"shots"`
							Deaths    float64 `json:"deaths"`
							HeadShots float64 `json:"headShots"`
						} `json:"properties"`
					} `json:"equip_snapshot_grenade"`
					EquipDecoy struct {
						Properties struct {
							Hits      float64 `json:"hits"`
							Kills     float64 `json:"kills"`
							Shots     float64 `json:"shots"`
							Deaths    float64 `json:"deaths"`
							HeadShots float64 `json:"headShots"`
						} `json:"properties"`
					} `json:"equip_decoy"`
					EquipSmoke struct {
						Properties struct {
							Hits      float64 `json:"hits"`
							Kills     float64 `json:"kills"`
							Shots     float64 `json:"shots"`
							Deaths    float64 `json:"deaths"`
							HeadShots float64 `json:"headShots"`
						} `json:"properties"`
					} `json:"equip_smoke"`
					EquipConcussion struct {
						Properties struct {
							Hits      float64 `json:"hits"`
							Kills     float64 `json:"kills"`
							Shots     float64 `json:"shots"`
							Deaths    float64 `json:"deaths"`
							HeadShots float64 `json:"headShots"`
						} `json:"properties"`
					} `json:"equip_concussion"`
					EquipHbSensor struct {
						Properties struct {
							Hits      float64 `json:"hits"`
							Kills     float64 `json:"kills"`
							Shots     float64 `json:"shots"`
							Deaths    float64 `json:"deaths"`
							HeadShots float64 `json:"headShots"`
						} `json:"properties"`
					} `json:"equip_hb_sensor"`
					EquipFlash struct {
						Properties struct {
							Hits      float64 `json:"hits"`
							Kills     float64 `json:"kills"`
							Shots     float64 `json:"shots"`
							Deaths    float64 `json:"deaths"`
							HeadShots float64 `json:"headShots"`
						} `json:"properties"`
					} `json:"equip_flash"`
					EquipAdrenaline struct {
						Properties struct {
							Hits      float64 `json:"hits"`
							Kills     float64 `json:"kills"`
							Shots     float64 `json:"shots"`
							Deaths    float64 `json:"deaths"`
							HeadShots float64 `json:"headShots"`
						} `json:"properties"`
					} `json:"equip_adrenaline"`
				} `json:"tacticals"`
				Lethals struct {
					EquipFrag struct {
						Properties struct {
							Hits      float64 `json:"hits"`
							Kills     float64 `json:"kills"`
							Shots     float64 `json:"shots"`
							Deaths    float64 `json:"deaths"`
							HeadShots float64 `json:"headShots"`
						} `json:"properties"`
					} `json:"equip_frag"`
					EquipThermite struct {
						Properties struct {
							Hits      float64 `json:"hits"`
							Kills     float64 `json:"kills"`
							Shots     float64 `json:"shots"`
							Deaths    float64 `json:"deaths"`
							HeadShots float64 `json:"headShots"`
						} `json:"properties"`
					} `json:"equip_thermite"`
					EquipSemtex struct {
						Properties struct {
							Hits      float64 `json:"hits"`
							Kills     float64 `json:"kills"`
							Shots     float64 `json:"shots"`
							Deaths    float64 `json:"deaths"`
							HeadShots float64 `json:"headShots"`
						} `json:"properties"`
					} `json:"equip_semtex"`
					EquipClaymore struct {
						Properties struct {
							Hits      float64 `json:"hits"`
							Kills     float64 `json:"kills"`
							Shots     float64 `json:"shots"`
							Deaths    float64 `json:"deaths"`
							HeadShots float64 `json:"headShots"`
						} `json:"properties"`
					} `json:"equip_claymore"`
					EquipC4 struct {
						Properties struct {
							Hits      float64 `json:"hits"`
							Kills     float64 `json:"kills"`
							Shots     float64 `json:"shots"`
							Deaths    float64 `json:"deaths"`
							HeadShots float64 `json:"headShots"`
						} `json:"properties"`
					} `json:"equip_c4"`
					EquipAtMine struct {
						Properties struct {
							Hits      float64 `json:"hits"`
							Kills     float64 `json:"kills"`
							Shots     float64 `json:"shots"`
							Deaths    float64 `json:"deaths"`
							HeadShots float64 `json:"headShots"`
						} `json:"properties"`
					} `json:"equip_at_mine"`
					EquipThrowingKnife struct {
						Properties struct {
							Hits      float64 `json:"hits"`
							Kills     float64 `json:"kills"`
							Shots     float64 `json:"shots"`
							Deaths    float64 `json:"deaths"`
							HeadShots float64 `json:"headShots"`
						} `json:"properties"`
					} `json:"equip_throwing_knife"`
					EquipMolotov struct {
						Properties struct {
							Hits      float64 `json:"hits"`
							Kills     float64 `json:"kills"`
							Shots     float64 `json:"shots"`
							Deaths    float64 `json:"deaths"`
							HeadShots float64 `json:"headShots"`
						} `json:"properties"`
					} `json:"equip_molotov"`
				} `json:"lethals"`
				WeaponLmg struct {
					Iw8LmKilo121 struct {
						Properties struct {
							Hits      float64 `json:"hits"`
							Kills     float64 `json:"kills"`
							KdRatio   float64 `json:"kdRatio"`
							Headshots float64 `json:"headshots"`
							Accuracy  float64 `json:"accuracy"`
							Shots     float64 `json:"shots"`
							Deaths    float64 `json:"deaths"`
						} `json:"properties"`
					} `json:"iw8_lm_kilo121"`
					Iw8LmMkilo3 struct {
						Properties struct {
							Hits      float64 `json:"hits"`
							Kills     float64 `json:"kills"`
							KdRatio   float64 `json:"kdRatio"`
							Headshots float64 `json:"headshots"`
							Accuracy  float64 `json:"accuracy"`
							Shots     float64 `json:"shots"`
							Deaths    float64 `json:"deaths"`
						} `json:"properties"`
					} `json:"iw8_lm_mkilo3"`
					Iw8LmMgolf34 struct {
						Properties struct {
							Hits      float64 `json:"hits"`
							Kills     float64 `json:"kills"`
							KdRatio   float64 `json:"kdRatio"`
							Headshots float64 `json:"headshots"`
							Accuracy  float64 `json:"accuracy"`
							Shots     float64 `json:"shots"`
							Deaths    float64 `json:"deaths"`
						} `json:"properties"`
					} `json:"iw8_lm_mgolf34"`
					Iw8LmLima86 struct {
						Properties struct {
							Hits      float64 `json:"hits"`
							Kills     float64 `json:"kills"`
							KdRatio   float64 `json:"kdRatio"`
							Headshots float64 `json:"headshots"`
							Accuracy  float64 `json:"accuracy"`
							Shots     float64 `json:"shots"`
							Deaths    float64 `json:"deaths"`
						} `json:"properties"`
					} `json:"iw8_lm_lima86"`
					Iw8LmPkilo struct {
						Properties struct {
							Hits      float64 `json:"hits"`
							Kills     float64 `json:"kills"`
							KdRatio   float64 `json:"kdRatio"`
							Headshots float64 `json:"headshots"`
							Accuracy  float64 `json:"accuracy"`
							Shots     float64 `json:"shots"`
							Deaths    float64 `json:"deaths"`
						} `json:"properties"`
					} `json:"iw8_lm_pkilo"`
					Iw8LmMgolf36 struct {
						Properties struct {
							Hits      float64 `json:"hits"`
							Kills     float64 `json:"kills"`
							KdRatio   float64 `json:"kdRatio"`
							Headshots float64 `json:"headshots"`
							Accuracy  float64 `json:"accuracy"`
							Shots     float64 `json:"shots"`
							Deaths    float64 `json:"deaths"`
						} `json:"properties"`
					} `json:"iw8_lm_mgolf36"`
				} `json:"weapon_lmg"`
				WeaponLauncher struct {
					Iw8LaGromeo struct {
						Properties struct {
							Hits      float64 `json:"hits"`
							Kills     float64 `json:"kills"`
							KdRatio   float64 `json:"kdRatio"`
							Headshots float64 `json:"headshots"`
							Accuracy  float64 `json:"accuracy"`
							Shots     float64 `json:"shots"`
							Deaths    float64 `json:"deaths"`
						} `json:"properties"`
					} `json:"iw8_la_gromeo"`
					Iw8LaRpapa7 struct {
						Properties struct {
							Hits      float64 `json:"hits"`
							Kills     float64 `json:"kills"`
							KdRatio   float64 `json:"kdRatio"`
							Headshots float64 `json:"headshots"`
							Accuracy  float64 `json:"accuracy"`
							Shots     float64 `json:"shots"`
							Deaths    float64 `json:"deaths"`
						} `json:"properties"`
					} `json:"iw8_la_rpapa7"`
					Iw8LaJuliet struct {
						Properties struct {
							Hits      float64 `json:"hits"`
							Kills     float64 `json:"kills"`
							KdRatio   float64 `json:"kdRatio"`
							Headshots float64 `json:"headshots"`
							Accuracy  float64 `json:"accuracy"`
							Shots     float64 `json:"shots"`
							Deaths    float64 `json:"deaths"`
						} `json:"properties"`
					} `json:"iw8_la_juliet"`
					Iw8LaKgolf struct {
						Properties struct {
							Hits      float64 `json:"hits"`
							Kills     float64 `json:"kills"`
							KdRatio   float64 `json:"kdRatio"`
							Headshots float64 `json:"headshots"`
							Accuracy  float64 `json:"accuracy"`
							Shots     float64 `json:"shots"`
							Deaths    float64 `json:"deaths"`
						} `json:"properties"`
					} `json:"iw8_la_kgolf"`
					Iw8LaMike32 struct {
						Properties struct {
							Hits      float64 `json:"hits"`
							Kills     float64 `json:"kills"`
							KdRatio   float64 `json:"kdRatio"`
							Headshots float64 `json:"headshots"`
							Accuracy  float64 `json:"accuracy"`
							Shots     float64 `json:"shots"`
							Deaths    float64 `json:"deaths"`
						} `json:"properties"`
					} `json:"iw8_la_mike32"`
				} `json:"weapon_launcher"`
				WeaponPistol struct {
					Iw8PiCpapa struct {
						Properties struct {
							Hits      float64 `json:"hits"`
							Kills     float64 `json:"kills"`
							KdRatio   float64 `json:"kdRatio"`
							Headshots float64 `json:"headshots"`
							Accuracy  float64 `json:"accuracy"`
							Shots     float64 `json:"shots"`
							Deaths    float64 `json:"deaths"`
						} `json:"properties"`
					} `json:"iw8_pi_cpapa"`
					Iw8PiMike9 struct {
						Properties struct {
							Hits      float64 `json:"hits"`
							Kills     float64 `json:"kills"`
							KdRatio   float64 `json:"kdRatio"`
							Headshots float64 `json:"headshots"`
							Accuracy  float64 `json:"accuracy"`
							Shots     float64 `json:"shots"`
							Deaths    float64 `json:"deaths"`
						} `json:"properties"`
					} `json:"iw8_pi_mike9"`
					Iw8PiMike1911 struct {
						Properties struct {
							Hits      float64 `json:"hits"`
							Kills     float64 `json:"kills"`
							KdRatio   float64 `json:"kdRatio"`
							Headshots float64 `json:"headshots"`
							Accuracy  float64 `json:"accuracy"`
							Shots     float64 `json:"shots"`
							Deaths    float64 `json:"deaths"`
						} `json:"properties"`
					} `json:"iw8_pi_mike1911"`
					Iw8PiGolf21 struct {
						Properties struct {
							Hits      float64 `json:"hits"`
							Kills     float64 `json:"kills"`
							KdRatio   float64 `json:"kdRatio"`
							Headshots float64 `json:"headshots"`
							Accuracy  float64 `json:"accuracy"`
							Shots     float64 `json:"shots"`
							Deaths    float64 `json:"deaths"`
						} `json:"properties"`
					} `json:"iw8_pi_golf21"`
					Iw8PiDecho struct {
						Properties struct {
							Hits      float64 `json:"hits"`
							Kills     float64 `json:"kills"`
							KdRatio   float64 `json:"kdRatio"`
							Headshots float64 `json:"headshots"`
							Accuracy  float64 `json:"accuracy"`
							Shots     float64 `json:"shots"`
							Deaths    float64 `json:"deaths"`
						} `json:"properties"`
					} `json:"iw8_pi_decho"`
					Iw8PiPapa320 struct {
						Properties struct {
							Hits      float64 `json:"hits"`
							Kills     float64 `json:"kills"`
							KdRatio   float64 `json:"kdRatio"`
							Headshots float64 `json:"headshots"`
							Accuracy  float64 `json:"accuracy"`
							Shots     float64 `json:"shots"`
							Deaths    float64 `json:"deaths"`
						} `json:"properties"`
					} `json:"iw8_pi_papa320"`
				} `json:"weapon_pistol"`
				WeaponAssaultRifle struct {
					Iw8ArFalima struct {
						Properties struct {
							Hits      float64 `json:"hits"`
							Kills     float64 `json:"kills"`
							KdRatio   float64 `json:"kdRatio"`
							Headshots float64 `json:"headshots"`
							Accuracy  float64 `json:"accuracy"`
							Shots     float64 `json:"shots"`
							Deaths    float64 `json:"deaths"`
						} `json:"properties"`
					} `json:"iw8_ar_falima"`
					Iw8ArTango21 struct {
						Properties struct {
							Hits      float64 `json:"hits"`
							Kills     float64 `json:"kills"`
							KdRatio   float64 `json:"kdRatio"`
							Headshots float64 `json:"headshots"`
							Accuracy  float64 `json:"accuracy"`
							Shots     float64 `json:"shots"`
							Deaths    float64 `json:"deaths"`
						} `json:"properties"`
					} `json:"iw8_ar_tango21"`
					Iw8ArMike4 struct {
						Properties struct {
							Hits      float64 `json:"hits"`
							Kills     float64 `json:"kills"`
							KdRatio   float64 `json:"kdRatio"`
							Headshots float64 `json:"headshots"`
							Accuracy  float64 `json:"accuracy"`
							Shots     float64 `json:"shots"`
							Deaths    float64 `json:"deaths"`
						} `json:"properties"`
					} `json:"iw8_ar_mike4"`
					Iw8ArFalpha struct {
						Properties struct {
							Hits      float64 `json:"hits"`
							Kills     float64 `json:"kills"`
							KdRatio   float64 `json:"kdRatio"`
							Headshots float64 `json:"headshots"`
							Accuracy  float64 `json:"accuracy"`
							Shots     float64 `json:"shots"`
							Deaths    float64 `json:"deaths"`
						} `json:"properties"`
					} `json:"iw8_ar_falpha"`
					Iw8ArMcharlie struct {
						Properties struct {
							Hits      float64 `json:"hits"`
							Kills     float64 `json:"kills"`
							KdRatio   float64 `json:"kdRatio"`
							Headshots float64 `json:"headshots"`
							Accuracy  float64 `json:"accuracy"`
							Shots     float64 `json:"shots"`
							Deaths    float64 `json:"deaths"`
						} `json:"properties"`
					} `json:"iw8_ar_mcharlie"`
					Iw8ArAkilo47 struct {
						Properties struct {
							Hits      float64 `json:"hits"`
							Kills     float64 `json:"kills"`
							KdRatio   float64 `json:"kdRatio"`
							Headshots float64 `json:"headshots"`
							Accuracy  float64 `json:"accuracy"`
							Shots     float64 `json:"shots"`
							Deaths    float64 `json:"deaths"`
						} `json:"properties"`
					} `json:"iw8_ar_akilo47"`
					Iw8ArKilo433 struct {
						Properties struct {
							Hits      float64 `json:"hits"`
							Kills     float64 `json:"kills"`
							KdRatio   float64 `json:"kdRatio"`
							Headshots float64 `json:"headshots"`
							Accuracy  float64 `json:"accuracy"`
							Shots     float64 `json:"shots"`
							Deaths    float64 `json:"deaths"`
						} `json:"properties"`
					} `json:"iw8_ar_kilo433"`
					Iw8ArScharlie struct {
						Properties struct {
							Hits      float64 `json:"hits"`
							Kills     float64 `json:"kills"`
							KdRatio   float64 `json:"kdRatio"`
							Headshots float64 `json:"headshots"`
							Accuracy  float64 `json:"accuracy"`
							Shots     float64 `json:"shots"`
							Deaths    float64 `json:"deaths"`
						} `json:"properties"`
					} `json:"iw8_ar_scharlie"`
					Iw8ArAsierra12 struct {
						Properties struct {
							Hits      float64 `json:"hits"`
							Kills     float64 `json:"kills"`
							KdRatio   float64 `json:"kdRatio"`
							Headshots float64 `json:"headshots"`
							Accuracy  float64 `json:"accuracy"`
							Shots     float64 `json:"shots"`
							Deaths    float64 `json:"deaths"`
						} `json:"properties"`
					} `json:"iw8_ar_asierra12"`
					Iw8ArGalima struct {
						Properties struct {
							Hits      float64 `json:"hits"`
							Kills     float64 `json:"kills"`
							KdRatio   float64 `json:"kdRatio"`
							Headshots float64 `json:"headshots"`
							Accuracy  float64 `json:"accuracy"`
							Shots     float64 `json:"shots"`
							Deaths    float64 `json:"deaths"`
						} `json:"properties"`
					} `json:"iw8_ar_galima"`
					Iw8ArSierra552 struct {
						Properties struct {
							Hits      float64 `json:"hits"`
							Kills     float64 `json:"kills"`
							KdRatio   float64 `json:"kdRatio"`
							Headshots float64 `json:"headshots"`
							Accuracy  float64 `json:"accuracy"`
							Shots     float64 `json:"shots"`
							Deaths    float64 `json:"deaths"`
						} `json:"properties"`
					} `json:"iw8_ar_sierra552"`
				} `json:"weapon_assault_rifle"`
				WeaponOther struct {
					Iw8MeRiotshield struct {
						Properties struct {
							Hits      float64 `json:"hits"`
							Kills     float64 `json:"kills"`
							KdRatio   float64 `json:"kdRatio"`
							Headshots float64 `json:"headshots"`
							Accuracy  float64 `json:"accuracy"`
							Shots     float64 `json:"shots"`
							Deaths    float64 `json:"deaths"`
						} `json:"properties"`
					} `json:"iw8_me_riotshield"`
				} `json:"weapon_other"`
				WeaponShotgun struct {
					Iw8ShMike26 struct {
						Properties struct {
							Hits      float64 `json:"hits"`
							Kills     float64 `json:"kills"`
							KdRatio   float64 `json:"kdRatio"`
							Headshots float64 `json:"headshots"`
							Accuracy  float64 `json:"accuracy"`
							Shots     float64 `json:"shots"`
							Deaths    float64 `json:"deaths"`
						} `json:"properties"`
					} `json:"iw8_sh_mike26"`
					Iw8ShCharlie725 struct {
						Properties struct {
							Hits      float64 `json:"hits"`
							Kills     float64 `json:"kills"`
							KdRatio   float64 `json:"kdRatio"`
							Headshots float64 `json:"headshots"`
							Accuracy  float64 `json:"accuracy"`
							Shots     float64 `json:"shots"`
							Deaths    float64 `json:"deaths"`
						} `json:"properties"`
					} `json:"iw8_sh_charlie725"`
					Iw8ShOscar12 struct {
						Properties struct {
							Hits      float64 `json:"hits"`
							Kills     float64 `json:"kills"`
							KdRatio   float64 `json:"kdRatio"`
							Headshots float64 `json:"headshots"`
							Accuracy  float64 `json:"accuracy"`
							Shots     float64 `json:"shots"`
							Deaths    float64 `json:"deaths"`
						} `json:"properties"`
					} `json:"iw8_sh_oscar12"`
					Iw8ShRomeo870 struct {
						Properties struct {
							Hits      float64 `json:"hits"`
							Kills     float64 `json:"kills"`
							KdRatio   float64 `json:"kdRatio"`
							Headshots float64 `json:"headshots"`
							Accuracy  float64 `json:"accuracy"`
							Shots     float64 `json:"shots"`
							Deaths    float64 `json:"deaths"`
						} `json:"properties"`
					} `json:"iw8_sh_romeo870"`
					Iw8ShDpapa12 struct {
						Properties struct {
							Hits      float64 `json:"hits"`
							Kills     float64 `json:"kills"`
							KdRatio   float64 `json:"kdRatio"`
							Headshots float64 `json:"headshots"`
							Accuracy  float64 `json:"accuracy"`
							Shots     float64 `json:"shots"`
							Deaths    float64 `json:"deaths"`
						} `json:"properties"`
					} `json:"iw8_sh_dpapa12"`
				} `json:"weapon_shotgun"`
				WeaponSmg struct {
					Iw8SmMpapa7 struct {
						Properties struct {
							Hits      float64 `json:"hits"`
							Kills     float64 `json:"kills"`
							KdRatio   float64 `json:"kdRatio"`
							Headshots float64 `json:"headshots"`
							Accuracy  float64 `json:"accuracy"`
							Shots     float64 `json:"shots"`
							Deaths    float64 `json:"deaths"`
						} `json:"properties"`
					} `json:"iw8_sm_mpapa7"`
					Iw8SmAugolf struct {
						Properties struct {
							Hits      float64 `json:"hits"`
							Kills     float64 `json:"kills"`
							KdRatio   float64 `json:"kdRatio"`
							Headshots float64 `json:"headshots"`
							Accuracy  float64 `json:"accuracy"`
							Shots     float64 `json:"shots"`
							Deaths    float64 `json:"deaths"`
						} `json:"properties"`
					} `json:"iw8_sm_augolf"`
					Iw8SmPapa90 struct {
						Properties struct {
							Hits      float64 `json:"hits"`
							Kills     float64 `json:"kills"`
							KdRatio   float64 `json:"kdRatio"`
							Headshots float64 `json:"headshots"`
							Accuracy  float64 `json:"accuracy"`
							Shots     float64 `json:"shots"`
							Deaths    float64 `json:"deaths"`
						} `json:"properties"`
					} `json:"iw8_sm_papa90"`
					Iw8SmMpapa5 struct {
						Properties struct {
							Hits      float64 `json:"hits"`
							Kills     float64 `json:"kills"`
							KdRatio   float64 `json:"kdRatio"`
							Headshots float64 `json:"headshots"`
							Accuracy  float64 `json:"accuracy"`
							Shots     float64 `json:"shots"`
							Deaths    float64 `json:"deaths"`
						} `json:"properties"`
					} `json:"iw8_sm_mpapa5"`
					Iw8SmSmgolf45 struct {
						Properties struct {
							Hits      float64 `json:"hits"`
							Kills     float64 `json:"kills"`
							KdRatio   float64 `json:"kdRatio"`
							Headshots float64 `json:"headshots"`
							Accuracy  float64 `json:"accuracy"`
							Shots     float64 `json:"shots"`
							Deaths    float64 `json:"deaths"`
						} `json:"properties"`
					} `json:"iw8_sm_smgolf45"`
					Iw8SmBeta struct {
						Properties struct {
							Hits      float64 `json:"hits"`
							Kills     float64 `json:"kills"`
							KdRatio   float64 `json:"kdRatio"`
							Headshots float64 `json:"headshots"`
							Accuracy  float64 `json:"accuracy"`
							Shots     float64 `json:"shots"`
							Deaths    float64 `json:"deaths"`
						} `json:"properties"`
					} `json:"iw8_sm_beta"`
					Iw8SmVictor struct {
						Properties struct {
							Hits      float64 `json:"hits"`
							Kills     float64 `json:"kills"`
							KdRatio   float64 `json:"kdRatio"`
							Headshots float64 `json:"headshots"`
							Accuracy  float64 `json:"accuracy"`
							Shots     float64 `json:"shots"`
							Deaths    float64 `json:"deaths"`
						} `json:"properties"`
					} `json:"iw8_sm_victor"`
					Iw8SmUzulu struct {
						Properties struct {
							Hits      float64 `json:"hits"`
							Kills     float64 `json:"kills"`
							KdRatio   float64 `json:"kdRatio"`
							Headshots float64 `json:"headshots"`
							Accuracy  float64 `json:"accuracy"`
							Shots     float64 `json:"shots"`
							Deaths    float64 `json:"deaths"`
						} `json:"properties"`
					} `json:"iw8_sm_uzulu"`
				} `json:"weapon_smg"`
				WeaponMarksman struct {
					Iw8SnSbeta struct {
						Properties struct {
							Hits      float64 `json:"hits"`
							Kills     float64 `json:"kills"`
							KdRatio   float64 `json:"kdRatio"`
							Headshots float64 `json:"headshots"`
							Accuracy  float64 `json:"accuracy"`
							Shots     float64 `json:"shots"`
							Deaths    float64 `json:"deaths"`
						} `json:"properties"`
					} `json:"iw8_sn_sbeta"`
					Iw8SnCrossbow struct {
						Properties struct {
							Hits      float64 `json:"hits"`
							Kills     float64 `json:"kills"`
							KdRatio   float64 `json:"kdRatio"`
							Headshots float64 `json:"headshots"`
							Accuracy  float64 `json:"accuracy"`
							Shots     float64 `json:"shots"`
							Deaths    float64 `json:"deaths"`
						} `json:"properties"`
					} `json:"iw8_sn_crossbow"`
					Iw8SnKilo98 struct {
						Properties struct {
							Hits      float64 `json:"hits"`
							Kills     float64 `json:"kills"`
							KdRatio   float64 `json:"kdRatio"`
							Headshots float64 `json:"headshots"`
							Accuracy  float64 `json:"accuracy"`
							Shots     float64 `json:"shots"`
							Deaths    float64 `json:"deaths"`
						} `json:"properties"`
					} `json:"iw8_sn_kilo98"`
					Iw8SnMike14 struct {
						Properties struct {
							Hits      float64 `json:"hits"`
							Kills     float64 `json:"kills"`
							KdRatio   float64 `json:"kdRatio"`
							Headshots float64 `json:"headshots"`
							Accuracy  float64 `json:"accuracy"`
							Shots     float64 `json:"shots"`
							Deaths    float64 `json:"deaths"`
						} `json:"properties"`
					} `json:"iw8_sn_mike14"`
					Iw8SnSksierra struct {
						Properties struct {
							Hits      float64 `json:"hits"`
							Kills     float64 `json:"kills"`
							KdRatio   float64 `json:"kdRatio"`
							Headshots float64 `json:"headshots"`
							Accuracy  float64 `json:"accuracy"`
							Shots     float64 `json:"shots"`
							Deaths    float64 `json:"deaths"`
						} `json:"properties"`
					} `json:"iw8_sn_sksierra"`
				} `json:"weapon_marksman"`
				WeaponMelee struct {
					Iw8MeAkimboblunt struct {
						Properties struct {
							Hits      float64 `json:"hits"`
							Kills     float64 `json:"kills"`
							KdRatio   float64 `json:"kdRatio"`
							Headshots float64 `json:"headshots"`
							Accuracy  float64 `json:"accuracy"`
							Shots     float64 `json:"shots"`
							Deaths    float64 `json:"deaths"`
						} `json:"properties"`
					} `json:"iw8_me_akimboblunt"`
					Iw8Knife struct {
						Properties struct {
							Hits      float64 `json:"hits"`
							Kills     float64 `json:"kills"`
							KdRatio   float64 `json:"kdRatio"`
							Headshots float64 `json:"headshots"`
							Accuracy  float64 `json:"accuracy"`
							Shots     float64 `json:"shots"`
							Deaths    float64 `json:"deaths"`
						} `json:"properties"`
					} `json:"iw8_knife"`
				} `json:"weapon_melee"`
			} `json:"itemData"`
			ScorestreakData struct {
				LethalScorestreakData struct {
					PrecisionAirstrike struct {
						Properties struct {
							ExtraStat1   float64 `json:"extraStat1"`
							Uses         float64 `json:"uses"`
							AwardedCount float64 `json:"awardedCount"`
						} `json:"properties"`
					} `json:"precision_airstrike"`
					CruisePredator struct {
						Properties struct {
							ExtraStat1   float64 `json:"extraStat1"`
							Uses         float64 `json:"uses"`
							AwardedCount float64 `json:"awardedCount"`
						} `json:"properties"`
					} `json:"cruise_predator"`
					ManualTurret struct {
						Properties struct {
							ExtraStat1   float64 `json:"extraStat1"`
							Uses         float64 `json:"uses"`
							AwardedCount float64 `json:"awardedCount"`
						} `json:"properties"`
					} `json:"manual_turret"`
					WhitePhosphorus struct {
						Properties struct {
							ExtraStat1   float64 `json:"extraStat1"`
							Uses         float64 `json:"uses"`
							AwardedCount float64 `json:"awardedCount"`
						} `json:"properties"`
					} `json:"white_phosphorus"`
					HoverJet struct {
						Properties struct {
							ExtraStat1   float64 `json:"extraStat1"`
							Uses         float64 `json:"uses"`
							AwardedCount float64 `json:"awardedCount"`
						} `json:"properties"`
					} `json:"hover_jet"`
					ChopperGunner struct {
						Properties struct {
							ExtraStat1   float64 `json:"extraStat1"`
							Uses         float64 `json:"uses"`
							AwardedCount float64 `json:"awardedCount"`
						} `json:"properties"`
					} `json:"chopper_gunner"`
					Gunship struct {
						Properties struct {
							ExtraStat1   float64 `json:"extraStat1"`
							Uses         float64 `json:"uses"`
							AwardedCount float64 `json:"awardedCount"`
						} `json:"properties"`
					} `json:"gunship"`
					SentryGun struct {
						Properties struct {
							ExtraStat1   float64 `json:"extraStat1"`
							Uses         float64 `json:"uses"`
							AwardedCount float64 `json:"awardedCount"`
						} `json:"properties"`
					} `json:"sentry_gun"`
					TomaStrike struct {
						Properties struct {
							ExtraStat1   float64 `json:"extraStat1"`
							Uses         float64 `json:"uses"`
							AwardedCount float64 `json:"awardedCount"`
						} `json:"properties"`
					} `json:"toma_strike"`
					Nuke struct {
						Properties struct {
							ExtraStat1   float64 `json:"extraStat1"`
							Uses         float64 `json:"uses"`
							AwardedCount float64 `json:"awardedCount"`
						} `json:"properties"`
					} `json:"nuke"`
					Juggernaut struct {
						Properties struct {
							ExtraStat1   float64 `json:"extraStat1"`
							Uses         float64 `json:"uses"`
							AwardedCount float64 `json:"awardedCount"`
						} `json:"properties"`
					} `json:"juggernaut"`
					PacSentry struct {
						Properties struct {
							ExtraStat1   float64 `json:"extraStat1"`
							Uses         float64 `json:"uses"`
							AwardedCount float64 `json:"awardedCount"`
						} `json:"properties"`
					} `json:"pac_sentry"`
					ChopperSupport struct {
						Properties struct {
							ExtraStat1   float64 `json:"extraStat1"`
							Uses         float64 `json:"uses"`
							AwardedCount float64 `json:"awardedCount"`
						} `json:"properties"`
					} `json:"chopper_support"`
					Bradley struct {
						Properties struct {
							ExtraStat1   float64 `json:"extraStat1"`
							Uses         float64 `json:"uses"`
							AwardedCount float64 `json:"awardedCount"`
						} `json:"properties"`
					} `json:"bradley"`
				} `json:"lethalScorestreakData"`
				SupportScorestreakData struct {
					Airdrop struct {
						Properties struct {
							ExtraStat1   float64 `json:"extraStat1"`
							Uses         float64 `json:"uses"`
							AwardedCount float64 `json:"awardedCount"`
						} `json:"properties"`
					} `json:"airdrop"`
					RadarDroneOverwatch struct {
						Properties struct {
							ExtraStat1   float64 `json:"extraStat1"`
							Uses         float64 `json:"uses"`
							AwardedCount float64 `json:"awardedCount"`
						} `json:"properties"`
					} `json:"radar_drone_overwatch"`
					ScramblerDroneGuard struct {
						Properties struct {
							ExtraStat1   float64 `json:"extraStat1"`
							Uses         float64 `json:"uses"`
							AwardedCount float64 `json:"awardedCount"`
						} `json:"properties"`
					} `json:"scrambler_drone_guard"`
					Uav struct {
						Properties struct {
							ExtraStat1   float64 `json:"extraStat1"`
							Uses         float64 `json:"uses"`
							AwardedCount float64 `json:"awardedCount"`
						} `json:"properties"`
					} `json:"uav"`
					AirdropMultiple struct {
						Properties struct {
							ExtraStat1   float64 `json:"extraStat1"`
							Uses         float64 `json:"uses"`
							AwardedCount float64 `json:"awardedCount"`
						} `json:"properties"`
					} `json:"airdrop_multiple"`
					DirectionalUav struct {
						Properties struct {
							ExtraStat1   float64 `json:"extraStat1"`
							Uses         float64 `json:"uses"`
							AwardedCount float64 `json:"awardedCount"`
						} `json:"properties"`
					} `json:"directional_uav"`
				} `json:"supportScorestreakData"`
			} `json:"scorestreakData"`
			AccoladeData struct {
				Properties struct {
					ClassChanges                           float64 `json:"classChanges"`
					HighestAvgAltitude                     float64 `json:"highestAvgAltitude"`
					KillsFromBehind                        float64 `json:"killsFromBehind"`
					LmgDeaths                              float64 `json:"lmgDeaths"`
					RiotShieldDamageAbsorbed               float64 `json:"riotShieldDamageAbsorbed"`
					FlashbangHits                          float64 `json:"flashbangHits"`
					MeleeKills                             float64 `json:"meleeKills"`
					TagsLargestBank                        float64 `json:"tagsLargestBank"`
					ShotgunKills                           float64 `json:"shotgunKills"`
					SniperDeaths                           float64 `json:"sniperDeaths"`
					TimeProne                              float64 `json:"timeProne"`
					KillstreakWhitePhosphorousKillsAssists float64 `json:"killstreakWhitePhosphorousKillsAssists"`
					ShortestLife                           float64 `json:"shortestLife"`
					DeathsFromBehind                       float64 `json:"deathsFromBehind"`
					HigherRankedKills                      float64 `json:"higherRankedKills"`
					MostAssists                            float64 `json:"mostAssists"`
					LeastKills                             float64 `json:"leastKills"`
					TagsDenied                             float64 `json:"tagsDenied"`
					KillstreakWheelsonKills                float64 `json:"killstreakWheelsonKills"`
					SniperHeadshots                        float64 `json:"sniperHeadshots"`
					KillstreakJuggernautKills              float64 `json:"killstreakJuggernautKills"`
					SmokesUsed                             float64 `json:"smokesUsed"`
					AvengerKills                           float64 `json:"avengerKills"`
					DecoyHits                              float64 `json:"decoyHits"`
					KillstreakCarePackageUsed              float64 `json:"killstreakCarePackageUsed"`
					MolotovKills                           float64 `json:"molotovKills"`
					GasHits                                float64 `json:"gasHits"`
					ComebackKills                          float64 `json:"comebackKills"`
					LmgHeadshots                           float64 `json:"lmgHeadshots"`
					SmgDeaths                              float64 `json:"smgDeaths"`
					CarrierKills                           float64 `json:"carrierKills"`
					DeployableCoverUsed                    float64 `json:"deployableCoverUsed"`
					ThermiteKills                          float64 `json:"thermiteKills"`
					ArKills                                float64 `json:"arKills"`
					C4Kills                                float64 `json:"c4Kills"`
					Suicides                               float64 `json:"suicides"`
					Clutch                                 float64 `json:"clutch"`
					SurvivorKills                          float64 `json:"survivorKills"`
					KillstreakGunshipKills                 float64 `json:"killstreakGunshipKills"`
					TimeSpentAsPassenger                   float64 `json:"timeSpentAsPassenger"`
					Returns                                float64 `json:"returns"`
					SmgHeadshots                           float64 `json:"smgHeadshots"`
					LauncherDeaths                         float64 `json:"launcherDeaths"`
					OneShotOneKills                        float64 `json:"oneShotOneKills"`
					AmmoBoxUsed                            float64 `json:"ammoBoxUsed"`
					SpawnSelectSquad                       float64 `json:"spawnSelectSquad"`
					WeaponPickups                          float64 `json:"weaponPickups"`
					PointBlankKills                        float64 `json:"pointBlankKills"`
					TagsCaptured                           float64 `json:"tagsCaptured"`
					KillstreakGroundKills                  float64 `json:"killstreakGroundKills"`
					DistanceTraveledInVehicle              float64 `json:"distanceTraveledInVehicle"`
					LongestLife                            float64 `json:"longestLife"`
					StunHits                               float64 `json:"stunHits"`
					SpawnSelectFlag                        float64 `json:"spawnSelectFlag"`
					ShotgunHeadshots                       float64 `json:"shotgunHeadshots"`
					BombDefused                            float64 `json:"bombDefused"`
					SnapshotHits                           float64 `json:"snapshotHits"`
					NoKillsWithDeath                       float64 `json:"noKillsWithDeath"`
					KillstreakAUAVAssists                  float64 `json:"killstreakAUAVAssists"`
					KillstreakPersonalUAVKills             float64 `json:"killstreakPersonalUAVKills"`
					TacticalInsertionSpawns                float64 `json:"tacticalInsertionSpawns"`
					LauncherKills                          float64 `json:"launcherKills"`
					SpawnSelectVehicle                     float64 `json:"spawnSelectVehicle"`
					MostKillsLeastDeaths                   float64 `json:"mostKillsLeastDeaths"`
					MostKills                              float64 `json:"mostKills"`
					Defends                                float64 `json:"defends"`
					TimeSpentAsDriver                      float64 `json:"timeSpentAsDriver"`
					BombDetonated                          float64 `json:"bombDetonated"`
					ArHeadshots                            float64 `json:"arHeadshots"`
					TimeOnPoint                            float64 `json:"timeOnPoint"`
					LmgKills                               float64 `json:"lmgKills"`
					KillstreakUAVAssists                   float64 `json:"killstreakUAVAssists"`
					CarepackagesCaptured                   float64 `json:"carepackagesCaptured"`
					MostKillsLongestStreak                 float64 `json:"mostKillsLongestStreak"`
					KillstreakCruiseMissileKills           float64 `json:"killstreakCruiseMissileKills"`
					LongestStreak                          float64 `json:"longestStreak"`
					DestroyedKillstreaks                   float64 `json:"destroyedKillstreaks"`
					HipfireKills                           float64 `json:"hipfireKills"`
					StimDamageHealed                       float64 `json:"stimDamageHealed"`
					SkippedKillcams                        float64 `json:"skippedKillcams"`
					LeastAssists                           float64 `json:"leastAssists"`
					MostMultikills                         float64 `json:"mostMultikills"`
					HighestRankedKills                     float64 `json:"highestRankedKills"`
					KillstreakAirstrikeKills               float64 `json:"killstreakAirstrikeKills"`
					DistanceTravelled                      float64 `json:"distanceTravelled"`
					KillstreakKills                        float64 `json:"killstreakKills"`
					SemtexKills                            float64 `json:"semtexKills"`
					PenetrationKills                       float64 `json:"penetrationKills"`
					ExplosionsSurvived                     float64 `json:"explosionsSurvived"`
					HighestMultikill                       float64 `json:"highestMultikill"`
					ArDeaths                               float64 `json:"arDeaths"`
					LongshotKills                          float64 `json:"longshotKills"`
					ProximityMineKills                     float64 `json:"proximityMineKills"`
					TagsMegaBanked                         float64 `json:"tagsMegaBanked"`
					MostKillsMostHeadshots                 float64 `json:"mostKillsMostHeadshots"`
					FirstInfected                          float64 `json:"firstInfected"`
					KillstreakCUAVAssists                  float64 `json:"killstreakCUAVAssists"`
					ThrowingKnifeKills                     float64 `json:"throwingKnifeKills"`
					ExecutionKills                         float64 `json:"executionKills"`
					LastSurvivor                           float64 `json:"lastSurvivor"`
					ReconDroneMarks                        float64 `json:"reconDroneMarks"`
					DeadSilenceKills                       float64 `json:"deadSilenceKills"`
					RevengeKills                           float64 `json:"revengeKills"`
					InfectedKills                          float64 `json:"infectedKills"`
					KillEnemyTeam                          float64 `json:"killEnemyTeam"`
					SniperKills                            float64 `json:"sniperKills"`
					KillstreakCluserStrikeKills            float64 `json:"killstreakCluserStrikeKills"`
					MeleeDeaths                            float64 `json:"meleeDeaths"`
					TimeWatchingKillcams                   float64 `json:"timeWatchingKillcams"`
					KillstreakTankKills                    float64 `json:"killstreakTankKills"`
					NoKillNoDeath                          float64 `json:"noKillNoDeath"`
					ShotgunDeaths                          float64 `json:"shotgunDeaths"`
					KillstreakChopperGunnerKills           float64 `json:"killstreakChopperGunnerKills"`
					ShotsFired                             float64 `json:"shotsFired"`
					StoppingPowerKills                     float64 `json:"stoppingPowerKills"`
					PistolPeaths                           float64 `json:"pistolPeaths"`
					KillstreakShieldTurretKills            float64 `json:"killstreakShieldTurretKills"`
					TimeCrouched                           float64 `json:"timeCrouched"`
					NoDeathsFromBehind                     float64 `json:"noDeathsFromBehind"`
					BombPlanted                            float64 `json:"bombPlanted"`
					Setbacks                               float64 `json:"setbacks"`
					SmgKills                               float64 `json:"smgKills"`
					ClaymoreKills                          float64 `json:"claymoreKills"`
					Kills10NoDeaths                        float64 `json:"kills10NoDeaths"`
					PistolHeadshots                        float64 `json:"pistolHeadshots"`
					KillstreakVTOLJetKills                 float64 `json:"killstreakVTOLJetKills"`
					Headshots                              float64 `json:"headshots"`
					MostDeaths                             float64 `json:"mostDeaths"`
					AdsKills                               float64 `json:"adsKills"`
					EmpDroneHits                           float64 `json:"empDroneHits"`
					DefenderKills                          float64 `json:"defenderKills"`
					LauncherHeadshots                      float64 `json:"launcherHeadshots"`
					TimesSelectedAsSquadLeader             float64 `json:"timesSelectedAsSquadLeader"`
					KillstreakAirKills                     float64 `json:"killstreakAirKills"`
					Assaults                               float64 `json:"assaults"`
					FragKills                              float64 `json:"fragKills"`
					KillstreakEmergencyAirdropUsed         float64 `json:"killstreakEmergencyAirdropUsed"`
					Captures                               float64 `json:"captures"`
					KillstreakChopperSupportKills          float64 `json:"killstreakChopperSupportKills"`
					SpawnSelectBase                        float64 `json:"spawnSelectBase"`
					NoKill10Deaths                         float64 `json:"noKill10Deaths"`
					LeastDeaths                            float64 `json:"leastDeaths"`
					KillstreakSentryGunKills               float64 `json:"killstreakSentryGunKills"`
					LongestTimeSpentOnWeapon               float64 `json:"longestTimeSpentOnWeapon"`
					LowerRankedKills                       float64 `json:"lowerRankedKills"`
					TrophySystemHits                       float64 `json:"trophySystemHits"`
					ClutchRevives                          float64 `json:"clutchRevives"`
					LowestAvgAltitude                      float64 `json:"lowestAvgAltitude"`
					Pickups                                float64 `json:"pickups"`
					PistolKills                            float64 `json:"pistolKills"`
					Reloads                                float64 `json:"reloads"`
				} `json:"properties"`
			} `json:"accoladeData"`
		} `json:"lifetime"`
		Weekly struct {
			All struct {
				Properties struct {
					Kills                          float64 `json:"kills"`
					ObjectiveTeamWiped             float64 `json:"objectiveTeamWiped"`
					ObjectiveLastStandKill         float64 `json:"objectiveLastStandKill"`
					WallBangs                      float64 `json:"wallBangs"`
					AvgLifeTime                    float64 `json:"avgLifeTime"`
					Score                          float64 `json:"score"`
					Headshots                      float64 `json:"headshots"`
					Assists                        float64 `json:"assists"`
					ObjectiveDestroyedVehicleLight float64 `json:"objectiveDestroyedVehicleLight"`
					KillsPerGame                   float64 `json:"killsPerGame"`
					ScorePerMinute                 float64 `json:"scorePerMinute"`
					DistanceTraveled               float64 `json:"distanceTraveled"`
					Deaths                         float64 `json:"deaths"`
					ObjectiveDestroyedEquipment    float64 `json:"objectiveDestroyedEquipment"`
					ObjectiveBrDownEnemyCircle4    float64 `json:"objectiveBrDownEnemyCircle4"`
					ObjectiveBrDownEnemyCircle3    float64 `json:"objectiveBrDownEnemyCircle3"`
					ObjectiveBrDownEnemyCircle2    float64 `json:"objectiveBrDownEnemyCircle2"`
					KdRatio                        float64 `json:"kdRatio"`
					ObjectiveBrDownEnemyCircle1    float64 `json:"objectiveBrDownEnemyCircle1"`
					ObjectiveBrMissionPickupTablet float64 `json:"objectiveBrMissionPickupTablet"`
					ObjectiveReviver               float64 `json:"objectiveReviver"`
					ObjectiveBrKioskBuy            float64 `json:"objectiveBrKioskBuy"`
					GulagDeaths                    float64 `json:"gulagDeaths"`
					ObjectiveBrDownEnemyCircle6    float64 `json:"objectiveBrDownEnemyCircle6"`
					ObjectiveBrDownEnemyCircle5    float64 `json:"objectiveBrDownEnemyCircle5"`
					TimePlayed                     float64 `json:"timePlayed"`
					HeadshotPercentage             float64 `json:"headshotPercentage"`
					Executions                     float64 `json:"executions"`
					MatchesPlayed                  float64 `json:"matchesPlayed"`
					GulagKills                     float64 `json:"gulagKills"`
					Nearmisses                     float64 `json:"nearmisses"`
					ObjectiveBrCacheOpen           float64 `json:"objectiveBrCacheOpen"`
					DamageDone                     float64 `json:"damageDone"`
					DamageTaken                    float64 `json:"damageTaken"`
				} `json:"properties"`
			} `json:"all"`
			Mode struct {
				BrBrsolo struct {
					Properties struct {
						Kills                          float64 `json:"kills"`
						ObjectiveLastStandKill         float64 `json:"objectiveLastStandKill"`
						WallBangs                      float64 `json:"wallBangs"`
						AvgLifeTime                    float64 `json:"avgLifeTime"`
						Score                          float64 `json:"score"`
						Headshots                      float64 `json:"headshots"`
						Assists                        float64 `json:"assists"`
						KillsPerGame                   float64 `json:"killsPerGame"`
						ScorePerMinute                 float64 `json:"scorePerMinute"`
						DistanceTraveled               float64 `json:"distanceTraveled"`
						Deaths                         float64 `json:"deaths"`
						ObjectiveBrDownEnemyCircle3    float64 `json:"objectiveBrDownEnemyCircle3"`
						ObjectiveBrDownEnemyCircle2    float64 `json:"objectiveBrDownEnemyCircle2"`
						KdRatio                        float64 `json:"kdRatio"`
						ObjectiveBrDownEnemyCircle1    float64 `json:"objectiveBrDownEnemyCircle1"`
						ObjectiveBrMissionPickupTablet float64 `json:"objectiveBrMissionPickupTablet"`
						ScorePerGame                   float64 `json:"scorePerGame"`
						ObjectiveBrKioskBuy            float64 `json:"objectiveBrKioskBuy"`
						GulagDeaths                    float64 `json:"gulagDeaths"`
						TimePlayed                     float64 `json:"timePlayed"`
						HeadshotPercentage             float64 `json:"headshotPercentage"`
						Executions                     float64 `json:"executions"`
						MatchesPlayed                  float64 `json:"matchesPlayed"`
						GulagKills                     float64 `json:"gulagKills"`
						Nearmisses                     float64 `json:"nearmisses"`
						ObjectiveBrCacheOpen           float64 `json:"objectiveBrCacheOpen"`
						DamageDone                     float64 `json:"damageDone"`
						DamageTaken                    float64 `json:"damageTaken"`
					} `json:"properties"`
				} `json:"br_brsolo"`
				BrBrtrios struct {
					Properties struct {
						Kills                          float64 `json:"kills"`
						ObjectiveTeamWiped             float64 `json:"objectiveTeamWiped"`
						ObjectiveLastStandKill         float64 `json:"objectiveLastStandKill"`
						WallBangs                      float64 `json:"wallBangs"`
						AvgLifeTime                    float64 `json:"avgLifeTime"`
						Score                          float64 `json:"score"`
						Headshots                      float64 `json:"headshots"`
						Assists                        float64 `json:"assists"`
						KillsPerGame                   float64 `json:"killsPerGame"`
						ScorePerMinute                 float64 `json:"scorePerMinute"`
						DistanceTraveled               float64 `json:"distanceTraveled"`
						Deaths                         float64 `json:"deaths"`
						ObjectiveBrDownEnemyCircle4    float64 `json:"objectiveBrDownEnemyCircle4"`
						ObjectiveBrDownEnemyCircle3    float64 `json:"objectiveBrDownEnemyCircle3"`
						KdRatio                        float64 `json:"kdRatio"`
						ObjectiveBrDownEnemyCircle1    float64 `json:"objectiveBrDownEnemyCircle1"`
						ObjectiveBrMissionPickupTablet float64 `json:"objectiveBrMissionPickupTablet"`
						ScorePerGame                   float64 `json:"scorePerGame"`
						ObjectiveBrKioskBuy            float64 `json:"objectiveBrKioskBuy"`
						GulagDeaths                    float64 `json:"gulagDeaths"`
						ObjectiveBrDownEnemyCircle6    float64 `json:"objectiveBrDownEnemyCircle6"`
						ObjectiveBrDownEnemyCircle5    float64 `json:"objectiveBrDownEnemyCircle5"`
						TimePlayed                     float64 `json:"timePlayed"`
						HeadshotPercentage             float64 `json:"headshotPercentage"`
						Executions                     float64 `json:"executions"`
						MatchesPlayed                  float64 `json:"matchesPlayed"`
						GulagKills                     float64 `json:"gulagKills"`
						Nearmisses                     float64 `json:"nearmisses"`
						ObjectiveBrCacheOpen           float64 `json:"objectiveBrCacheOpen"`
						DamageDone                     float64 `json:"damageDone"`
						DamageTaken                    float64 `json:"damageTaken"`
					} `json:"properties"`
				} `json:"br_brtrios"`
				BrtdmRmbl struct {
					Properties struct {
						Kills                          float64 `json:"kills"`
						KdRatio                        float64 `json:"kdRatio"`
						ScorePerGame                   float64 `json:"scorePerGame"`
						WallBangs                      float64 `json:"wallBangs"`
						AvgLifeTime                    float64 `json:"avgLifeTime"`
						Score                          float64 `json:"score"`
						TimePlayed                     float64 `json:"timePlayed"`
						HeadshotPercentage             float64 `json:"headshotPercentage"`
						Headshots                      float64 `json:"headshots"`
						Executions                     float64 `json:"executions"`
						MatchesPlayed                  float64 `json:"matchesPlayed"`
						Assists                        float64 `json:"assists"`
						Nearmisses                     float64 `json:"nearmisses"`
						ObjectiveDestroyedVehicleLight float64 `json:"objectiveDestroyedVehicleLight"`
						KillsPerGame                   float64 `json:"killsPerGame"`
						ScorePerMinute                 float64 `json:"scorePerMinute"`
						DistanceTraveled               float64 `json:"distanceTraveled"`
						DamageDone                     float64 `json:"damageDone"`
						Deaths                         float64 `json:"deaths"`
						DamageTaken                    float64 `json:"damageTaken"`
						ObjectiveDestroyedEquipment    float64 `json:"objectiveDestroyedEquipment"`
					} `json:"properties"`
				} `json:"brtdm_rmbl"`
				BrAll struct {
					Properties struct {
						Kills                          float64 `json:"kills"`
						ObjectiveTeamWiped             float64 `json:"objectiveTeamWiped"`
						ObjectiveLastStandKill         float64 `json:"objectiveLastStandKill"`
						WallBangs                      float64 `json:"wallBangs"`
						AvgLifeTime                    float64 `json:"avgLifeTime"`
						Score                          float64 `json:"score"`
						Headshots                      float64 `json:"headshots"`
						Assists                        float64 `json:"assists"`
						ObjectiveDestroyedVehicleLight float64 `json:"objectiveDestroyedVehicleLight"`
						KillsPerGame                   float64 `json:"killsPerGame"`
						ScorePerMinute                 float64 `json:"scorePerMinute"`
						DistanceTraveled               float64 `json:"distanceTraveled"`
						Deaths                         float64 `json:"deaths"`
						ObjectiveDestroyedEquipment    float64 `json:"objectiveDestroyedEquipment"`
						ObjectiveBrDownEnemyCircle4    float64 `json:"objectiveBrDownEnemyCircle4"`
						ObjectiveBrDownEnemyCircle3    float64 `json:"objectiveBrDownEnemyCircle3"`
						ObjectiveBrDownEnemyCircle2    float64 `json:"objectiveBrDownEnemyCircle2"`
						KdRatio                        float64 `json:"kdRatio"`
						ObjectiveBrDownEnemyCircle1    float64 `json:"objectiveBrDownEnemyCircle1"`
						ObjectiveBrMissionPickupTablet float64 `json:"objectiveBrMissionPickupTablet"`
						ObjectiveReviver               float64 `json:"objectiveReviver"`
						ObjectiveBrKioskBuy            float64 `json:"objectiveBrKioskBuy"`
						GulagDeaths                    float64 `json:"gulagDeaths"`
						ObjectiveBrDownEnemyCircle6    float64 `json:"objectiveBrDownEnemyCircle6"`
						ObjectiveBrDownEnemyCircle5    float64 `json:"objectiveBrDownEnemyCircle5"`
						TimePlayed                     float64 `json:"timePlayed"`
						HeadshotPercentage             float64 `json:"headshotPercentage"`
						Executions                     float64 `json:"executions"`
						MatchesPlayed                  float64 `json:"matchesPlayed"`
						GulagKills                     float64 `json:"gulagKills"`
						Nearmisses                     float64 `json:"nearmisses"`
						ObjectiveBrCacheOpen           float64 `json:"objectiveBrCacheOpen"`
						DamageDone                     float64 `json:"damageDone"`
						DamageTaken                    float64 `json:"damageTaken"`
					} `json:"properties"`
				} `json:"br_all"`
				BrBrquads struct {
					Properties struct {
						Kills                          float64 `json:"kills"`
						ObjectiveLastStandKill         float64 `json:"objectiveLastStandKill"`
						WallBangs                      float64 `json:"wallBangs"`
						AvgLifeTime                    float64 `json:"avgLifeTime"`
						Score                          float64 `json:"score"`
						Headshots                      float64 `json:"headshots"`
						Assists                        float64 `json:"assists"`
						KillsPerGame                   float64 `json:"killsPerGame"`
						ScorePerMinute                 float64 `json:"scorePerMinute"`
						DistanceTraveled               float64 `json:"distanceTraveled"`
						Deaths                         float64 `json:"deaths"`
						KdRatio                        float64 `json:"kdRatio"`
						ObjectiveBrDownEnemyCircle1    float64 `json:"objectiveBrDownEnemyCircle1"`
						ObjectiveBrMissionPickupTablet float64 `json:"objectiveBrMissionPickupTablet"`
						ObjectiveReviver               float64 `json:"objectiveReviver"`
						ScorePerGame                   float64 `json:"scorePerGame"`
						ObjectiveBrKioskBuy            float64 `json:"objectiveBrKioskBuy"`
						GulagDeaths                    float64 `json:"gulagDeaths"`
						TimePlayed                     float64 `json:"timePlayed"`
						HeadshotPercentage             float64 `json:"headshotPercentage"`
						Executions                     float64 `json:"executions"`
						MatchesPlayed                  float64 `json:"matchesPlayed"`
						GulagKills                     float64 `json:"gulagKills"`
						Nearmisses                     float64 `json:"nearmisses"`
						ObjectiveBrCacheOpen           float64 `json:"objectiveBrCacheOpen"`
						DamageDone                     float64 `json:"damageDone"`
						DamageTaken                    float64 `json:"damageTaken"`
					} `json:"properties"`
				} `json:"br_brquads"`
			} `json:"mode"`
			Map struct {
			} `json:"map"`
		} `json:"weekly"`
		Engagement interface{} `json:"engagement"`
	} `json:"data"`
}
