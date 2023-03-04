package dto

import "time"

type LiveMatchDto struct {
	Code    string `json:"code"`
	Env     string `json:"env"`
	Message string `json:"message"`
	Payload struct {
		ID          string `json:"id"`
		Type        string `json:"type"`
		Game        string `json:"game"`
		Region      string `json:"region"`
		OrganizerID string `json:"organizerId"`
		Entity      struct {
			Type string `json:"type"`
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"entity"`
		EntityCustom struct {
			EffectiveRanking float64 `json:"effectiveRanking"`
			MatcherMatchID   string  `json:"matcherMatchId"`
			Parties          struct {
				NAMING_FAILED  []string `json:""`
				NAMING_FAILED0 []string `json:""`
				NAMING_FAILED1 []string `json:""`
				NAMING_FAILED2 []string `json:""`
				NAMING_FAILED3 []string `json:""`
				NAMING_FAILED4 []string `json:""`
				NAMING_FAILED5 []string `json:""`
				NAMING_FAILED6 []string `json:""`
			} `json:"parties"`
			PartyQueueDurations struct {
				NAMING_FAILED  float64 `json:""`
				NAMING_FAILED0 float64 `json:""`
				NAMING_FAILED1 float64 `json:""`
				NAMING_FAILED2 float64 `json:""`
				NAMING_FAILED3 float64 `json:""`
				NAMING_FAILED4 float64 `json:""`
				NAMING_FAILED5 float64 `json:""`
				NAMING_FAILED6 float64 `json:""`
			} `json:"partyQueueDurations"`
			QueueID string `json:"queueId"`
		} `json:"entityCustom"`
		AllowOngoingJoin  bool   `json:"allowOngoingJoin"`
		AnticheatRequired bool   `json:"anticheatRequired"`
		AnticheatMode     string `json:"anticheatMode"`
		CalculateElo      bool   `json:"calculateElo"`
		SkillFeedback     string `json:"skillFeedback"`
		AfkAction         string `json:"afkAction"`
		FbiManagement     bool   `json:"fbiManagement"`
		AdminTool         bool   `json:"adminTool"`
		CheckIn           struct {
			Time           int       `json:"time"`
			TotalCheckedIn int       `json:"totalCheckedIn"`
			TotalPlayers   int       `json:"totalPlayers"`
			EndTime        time.Time `json:"endTime"`
			CheckedIn      bool      `json:"checkedIn"`
		} `json:"checkIn"`
		State  string   `json:"state"`
		Status string   `json:"status"`
		Round  int      `json:"round"`
		States []string `json:"states"`
		Teams  struct {
			Faction1 Faction `json:"faction1"`
			Faction2 Faction `json:"faction2"`
		} `json:"teams"`
		Spectators  []any `json:"spectators"`
		MatchCustom struct {
			ID       string `json:"id"`
			Overview struct {
				Description struct {
					En string `json:"en"`
				} `json:"description"`
				Game  string `json:"game"`
				Label struct {
					En string `json:"en"`
				} `json:"label"`
				Name   string `json:"name"`
				Region string `json:"region"`
				Round  struct {
					Label struct {
						En string `json:"en"`
					} `json:"label"`
					ID     string `json:"id"`
					Type   string `json:"type"`
					ToPlay int    `json:"to_play"`
					ToWin  int    `json:"to_win"`
				} `json:"round"`
				Detections struct {
					Afk     bool `json:"afk"`
					Leavers bool `json:"leavers"`
				} `json:"detections"`
				Spectators             bool   `json:"spectators"`
				EloMode                string `json:"elo_mode"`
				ExpireSeconds          int    `json:"expire_seconds"`
				GroupingStats          string `json:"grouping_stats"`
				MaxPlayers             int    `json:"max_players"`
				MinPlayers             int    `json:"min_players"`
				TeamSize               int    `json:"team_size"`
				TimeToConnect          int    `json:"time_to_connect"`
				TimeOutSelectRandom    bool   `json:"time_out_select_random"`
				OrganizerID            string `json:"organizer_id"`
				EloType                string `json:"elo_type"`
				MatchConfigurationType struct {
					Value string `json:"value"`
					Label struct {
						En string `json:"en"`
					} `json:"label"`
				} `json:"match_configuration_type"`
				MatchFinishedType struct {
					Value string `json:"value"`
					Label struct {
						En string `json:"en"`
					} `json:"label"`
				} `json:"match_finished_type"`
				GameType string `json:"game_type"`
			} `json:"overview"`
			Tree struct {
				GameConfig struct {
					DataType string `json:"data_type"`
					Flags    struct {
					} `json:"flags"`
					ID       string `json:"id"`
					LeafNode bool   `json:"leaf_node"`
					Values   struct {
						Value string `json:"value"`
					} `json:"values"`
				} `json:"game_config"`
				Location struct {
					DataType string `json:"data_type"`
					Display  struct {
						Priority int `json:"priority"`
					} `json:"display"`
					Flags struct {
						Votable bool `json:"votable"`
					} `json:"flags"`
					ID    string `json:"id"`
					Label struct {
						En string `json:"en"`
					} `json:"label"`
					LeafNode bool   `json:"leaf_node"`
					Name     string `json:"name"`
					Values   struct {
						Value []struct {
							ClassName      string `json:"class_name"`
							GameLocationID string `json:"game_location_id"`
							GUID           string `json:"guid"`
							ImageLg        string `json:"image_lg"`
							ImageSm        string `json:"image_sm"`
							Name           string `json:"name"`
						} `json:"value"`
						VotingSteps []string `json:"voting_steps"`
					} `json:"values"`
				} `json:"location"`
				Map struct {
					DataType string `json:"data_type"`
					Display  struct {
						Priority int `json:"priority"`
					} `json:"display"`
					Flags struct {
						Votable bool `json:"votable"`
					} `json:"flags"`
					ID    string `json:"id"`
					Label struct {
						En string `json:"en"`
					} `json:"label"`
					LeafNode bool   `json:"leaf_node"`
					Name     string `json:"name"`
					Values   struct {
						MultiSelect struct {
							Memberships []string `json:"memberships"`
							Minimum     int      `json:"minimum"`
						} `json:"multi_select"`
						Value []struct {
							ClassName string `json:"class_name"`
							GameMapID string `json:"game_map_id"`
							GUID      string `json:"guid"`
							ImageLg   string `json:"image_lg"`
							ImageSm   string `json:"image_sm"`
							Name      string `json:"name"`
						} `json:"value"`
						VotingSteps []string `json:"voting_steps"`
					} `json:"values"`
				} `json:"map"`
				ServerConfig struct {
					Children struct {
						Bots struct {
							DataType string `json:"data_type"`
							Flags    struct {
							} `json:"flags"`
							ID       string `json:"id"`
							LeafNode bool   `json:"leaf_node"`
							Values   struct {
								Value bool `json:"value"`
							} `json:"values"`
						} `json:"bots"`
						BotsDifficulty struct {
							DataType string `json:"data_type"`
							Flags    struct {
							} `json:"flags"`
							ID       string `json:"id"`
							LeafNode bool   `json:"leaf_node"`
							Values   struct {
								MaxValue int `json:"max_value"`
								MinValue int `json:"min_value"`
								Value    int `json:"value"`
							} `json:"values"`
						} `json:"botsDifficulty"`
						DeadTalk struct {
							DataType string `json:"data_type"`
							Display  struct {
								Priority int `json:"priority"`
							} `json:"display"`
							Flags struct {
							} `json:"flags"`
							ID    string `json:"id"`
							Label struct {
								En string `json:"en"`
							} `json:"label"`
							LeafNode bool   `json:"leaf_node"`
							Name     string `json:"name"`
							Values   struct {
								Value bool `json:"value"`
							} `json:"values"`
						} `json:"deadTalk"`
						FreezeTime struct {
							DataType string `json:"data_type"`
							Flags    struct {
							} `json:"flags"`
							ID       string `json:"id"`
							LeafNode bool   `json:"leaf_node"`
							Values   struct {
								Value int `json:"value"`
							} `json:"values"`
						} `json:"freezeTime"`
						FriendlyFire struct {
							DataType string `json:"data_type"`
							Flags    struct {
							} `json:"flags"`
							ID       string `json:"id"`
							LeafNode bool   `json:"leaf_node"`
							Values   struct {
								Value bool `json:"value"`
							} `json:"values"`
						} `json:"friendlyFire"`
						GameMode struct {
							DataType string `json:"data_type"`
							Flags    struct {
							} `json:"flags"`
							ID       string `json:"id"`
							LeafNode bool   `json:"leaf_node"`
							Values   struct {
								Value int `json:"value"`
							} `json:"values"`
						} `json:"gameMode"`
						GameType struct {
							DataType string `json:"data_type"`
							Flags    struct {
							} `json:"flags"`
							ID       string `json:"id"`
							LeafNode bool   `json:"leaf_node"`
							Values   struct {
								MaxValue int `json:"max_value"`
								MinValue int `json:"min_value"`
								Value    int `json:"value"`
							} `json:"values"`
						} `json:"gameType"`
						KnifeRound struct {
							DataType string `json:"data_type"`
							Flags    struct {
							} `json:"flags"`
							ID       string `json:"id"`
							LeafNode bool   `json:"leaf_node"`
							Values   struct {
								Value bool `json:"value"`
							} `json:"values"`
						} `json:"knifeRound"`
						MaxRounds struct {
							DataType string `json:"data_type"`
							Flags    struct {
							} `json:"flags"`
							ID       string `json:"id"`
							LeafNode bool   `json:"leaf_node"`
							Values   struct {
								MaxValue int `json:"max_value"`
								MinValue int `json:"min_value"`
								Value    int `json:"value"`
							} `json:"values"`
						} `json:"maxRounds"`
						OvertimeHalftimePausetimer struct {
							DataType string `json:"data_type"`
							Flags    struct {
							} `json:"flags"`
							ID       string `json:"id"`
							LeafNode bool   `json:"leaf_node"`
							Values   struct {
								Value bool `json:"value"`
							} `json:"values"`
						} `json:"overtimeHalftimePausetimer"`
						OvertimeMaxRounds struct {
							DataType string `json:"data_type"`
							Flags    struct {
							} `json:"flags"`
							ID       string `json:"id"`
							LeafNode bool   `json:"leaf_node"`
							Values   struct {
								MaxValue int `json:"max_value"`
								MinValue int `json:"min_value"`
								Value    int `json:"value"`
							} `json:"values"`
						} `json:"overtimeMaxRounds"`
						OvertimeStartMoney struct {
							DataType string `json:"data_type"`
							Flags    struct {
							} `json:"flags"`
							ID       string `json:"id"`
							LeafNode bool   `json:"leaf_node"`
							Values   struct {
								Value int `json:"value"`
							} `json:"values"`
						} `json:"overtimeStartMoney"`
						Pause struct {
							DataType string `json:"data_type"`
							Flags    struct {
							} `json:"flags"`
							ID       string `json:"id"`
							LeafNode bool   `json:"leaf_node"`
							Values   struct {
								Value bool `json:"value"`
							} `json:"values"`
						} `json:"pause"`
						PauseTime struct {
							DataType string `json:"data_type"`
							Flags    struct {
							} `json:"flags"`
							ID       string `json:"id"`
							LeafNode bool   `json:"leaf_node"`
							Values   struct {
								MaxValue int `json:"max_value"`
								MinValue int `json:"min_value"`
								Value    int `json:"value"`
							} `json:"values"`
						} `json:"pauseTime"`
						RespawnImmunityTime struct {
							DataType string `json:"data_type"`
							Flags    struct {
							} `json:"flags"`
							ID       string `json:"id"`
							LeafNode bool   `json:"leaf_node"`
							Optional bool   `json:"optional"`
							Values   struct {
								Value int `json:"value"`
							} `json:"values"`
						} `json:"respawnImmunityTime"`
						StartMoney struct {
							DataType string `json:"data_type"`
							Flags    struct {
							} `json:"flags"`
							ID       string `json:"id"`
							LeafNode bool   `json:"leaf_node"`
							Values   struct {
								MaxValue int `json:"max_value"`
								MinValue int `json:"min_value"`
								Value    int `json:"value"`
							} `json:"values"`
						} `json:"startMoney"`
						StartOnReady struct {
							DataType string `json:"data_type"`
							Flags    struct {
							} `json:"flags"`
							ID       string `json:"id"`
							LeafNode bool   `json:"leaf_node"`
							Values   struct {
								Value bool `json:"value"`
							} `json:"values"`
						} `json:"startOnReady"`
						TimeToConnect struct {
							DataType string `json:"data_type"`
							Flags    struct {
							} `json:"flags"`
							ID       string `json:"id"`
							LeafNode bool   `json:"leaf_node"`
							Link     string `json:"link"`
							Optional bool   `json:"optional"`
							Values   struct {
								MaxValue int `json:"max_value"`
								MinValue int `json:"min_value"`
								Value    int `json:"value"`
							} `json:"values"`
						} `json:"timeToConnect"`
						TimeoutMax struct {
							DataType string `json:"data_type"`
							Display  struct {
								Priority int `json:"priority"`
							} `json:"display"`
							Flags struct {
							} `json:"flags"`
							ID    string `json:"id"`
							Label struct {
								En string `json:"en"`
							} `json:"label"`
							LeafNode bool   `json:"leaf_node"`
							Name     string `json:"name"`
							Values   struct {
								MaxValue int `json:"max_value"`
								MinValue int `json:"min_value"`
								Value    int `json:"value"`
							} `json:"values"`
						} `json:"timeoutMax"`
						TimeoutTime struct {
							DataType string `json:"data_type"`
							Display  struct {
								Priority int `json:"priority"`
							} `json:"display"`
							Flags struct {
							} `json:"flags"`
							ID    string `json:"id"`
							Label struct {
								En string `json:"en"`
							} `json:"label"`
							LeafNode bool   `json:"leaf_node"`
							Name     string `json:"name"`
							Values   struct {
								MaxValue int `json:"max_value"`
								MinValue int `json:"min_value"`
								Value    int `json:"value"`
							} `json:"values"`
						} `json:"timeoutTime"`
						TvDelay struct {
							DataType string `json:"data_type"`
							Flags    struct {
							} `json:"flags"`
							ID       string `json:"id"`
							LeafNode bool   `json:"leaf_node"`
							Values   struct {
								MaxValue int `json:"max_value"`
								MinValue int `json:"min_value"`
								Value    int `json:"value"`
							} `json:"values"`
						} `json:"tvDelay"`
						VoteKick struct {
							DataType string `json:"data_type"`
							Flags    struct {
							} `json:"flags"`
							ID       string `json:"id"`
							LeafNode bool   `json:"leaf_node"`
							Values   struct {
								Value bool `json:"value"`
							} `json:"values"`
						} `json:"voteKick"`
					} `json:"children"`
					ID string `json:"id"`
				} `json:"server_config"`
				Stream struct {
					DataType string `json:"data_type"`
					Flags    struct {
					} `json:"flags"`
					ID       string `json:"id"`
					LeafNode bool   `json:"leaf_node"`
					Values   struct {
						Value bool `json:"value"`
					} `json:"values"`
				} `json:"stream"`
			} `json:"tree"`
		} `json:"matchCustom"`
		Voting struct {
			VotedEntityTypes []string `json:"voted_entity_types"`
			Location         struct {
				Entities []struct {
					ClassName      string `json:"class_name"`
					GameLocationID string `json:"game_location_id"`
					GUID           string `json:"guid"`
					ImageLg        string `json:"image_lg"`
					ImageSm        string `json:"image_sm"`
					Name           string `json:"name"`
				} `json:"entities"`
				Pick []string `json:"pick"`
			} `json:"location"`
			Map struct {
				Entities []struct {
					ClassName string `json:"class_name"`
					GameMapID string `json:"game_map_id"`
					GUID      string `json:"guid"`
					ImageLg   string `json:"image_lg"`
					ImageSm   string `json:"image_sm"`
					Name      string `json:"name"`
				} `json:"entities"`
				Pick []string `json:"pick"`
			} `json:"map"`
		} `json:"voting"`
		Maps []struct {
			ClassName string `json:"class_name"`
			GameMapID string `json:"game_map_id"`
			GUID      string `json:"guid"`
			ImageLg   string `json:"image_lg"`
			ImageSm   string `json:"image_sm"`
			Name      string `json:"name"`
		} `json:"maps"`
		Locations []struct {
			ClassName      string `json:"class_name"`
			GameLocationID string `json:"game_location_id"`
			GUID           string `json:"guid"`
			ImageLg        string `json:"image_lg"`
			ImageSm        string `json:"image_sm"`
			Name           string `json:"name"`
		} `json:"locations"`
		SummaryResults struct {
			AscScore bool `json:"ascScore"`
			Factions struct {
				Faction2 struct {
					Score int `json:"score"`
				} `json:"faction2"`
				Faction1 struct {
					Score int `json:"score"`
				} `json:"faction1"`
			} `json:"factions"`
			Leavers []any `json:"leavers"`
			Afk     []any `json:"afk"`
		} `json:"summaryResults"`
		Results []struct {
			AscScore bool `json:"ascScore"`
			Factions struct {
				Faction2 struct {
					Score int `json:"score"`
				} `json:"faction2"`
				Faction1 struct {
					Score int `json:"score"`
				} `json:"faction1"`
			} `json:"factions"`
			Leavers []any `json:"leavers"`
			Afk     []any `json:"afk"`
		} `json:"results"`
		StartedAt     time.Time `json:"startedAt"`
		ConfiguredAt  time.Time `json:"configuredAt"`
		TimeToConnect int       `json:"timeToConnect"`
		Version       int       `json:"version"`
		CreatedAt     time.Time `json:"createdAt"`
		LastModified  time.Time `json:"lastModified"`
		Parties       []struct {
			PartyID string   `json:"partyId"`
			Users   []string `json:"users"`
		} `json:"parties"`
	} `json:"payload"`
	Time    int64  `json:"time"`
	Version string `json:"version"`
}

type Faction struct {
	ID     string   `json:"id"`
	Name   string   `json:"name"`
	Avatar string   `json:"avatar"`
	Leader string   `json:"leader"`
	Roster []Roster `json:"roster"`
	Stats  struct {
		WinProbability float64 `json:"winProbability"`
		SkillLevel     struct {
			Average int `json:"average"`
			Range   struct {
				Min int `json:"min"`
				Max int `json:"max"`
			} `json:"range"`
		} `json:"skillLevel"`
		Rating int `json:"rating"`
	} `json:"stats"`
	Substituted bool `json:"substituted"`
}

type Roster struct {
	ID             string   `json:"id"`
	Nickname       string   `json:"nickname"`
	Avatar         string   `json:"avatar,omitempty"`
	GameID         string   `json:"gameId"`
	GameName       string   `json:"gameName"`
	Memberships    []string `json:"memberships"`
	Elo            int      `json:"elo"`
	GameSkillLevel int      `json:"gameSkillLevel"`
	AcReq          bool     `json:"acReq"`
	PartyID        string   `json:"partyId"`
}
