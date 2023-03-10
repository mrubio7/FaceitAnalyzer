const getPredictedData = (matchId) => fetch(`/analyze?q=${matchId}`)

const fakeData = {
	"Data": [
	  1.0287201404571533,
	  0.701583981513977,
	  3.7248001098632812,
	  1.051740050315857,
	  0.7106479406356812,
	  3.895599842071533,
	  1.2260000705718994,
	  0.8402000665664673,
	  4.659999847412109,
	  0.9898281097412109,
	  0.6740398406982422,
	  3.693999767303467,
	  1.0472561120986938,
	  0.6947080492973328,
	  3.810399293899536,
	  1.073799967765808,
	  0.7401999831199646,
	  3.9800000190734863
	],
	"Label": 0
};

const fakeResult = {
	Result: 0.235673,
	TeamA: {
	  teamName: "team_Kara0ne",
	  averageElo: "1206",
	  players: [
		{
		  playerId: "f688248f-c5e5-4ea0-8c49-83101c5ad19a",
		  nickname: "Kara0ne",
		  lvl: 3,
		  avatar: "https://distribution.faceit-cdn.net/images/2b457438-f692-476c-862b-bc24eb62ec33.jpeg",
		  stats: {
			c3: 0.75299996,
			c2: 1.1069999,
			c4: 3.9
		  },
		  teamStats: {
			c3: 0.71798,
			c2: 1.0593407,
			c4: 3.7220008
		  },
		  enemyStats: {
			c3: 0.70711976,
			c2: 1.0388601,
			c4: 3.468001
		  }
		},
		{
		  playerId: "d8f7f0c9-5771-43ca-8a96-f219eb28e7ab",
		  nickname: "iksagi",
		  lvl: 4,
		  avatar: "https://assets.faceit-cdn.net/avatars/d8f7f0c9-5771-43ca-8a96-f219eb28e7ab_1601909488716.jpg",
		  stats: {
			c3: 0.722,
			c2: 0.984,
			c4: 4.9
		  },
		  teamStats: {
			c3: 0.6936199,
			c2: 1.0265998,
			c4: 4.0479994
		  },
		  enemyStats: {
			c3: 0.70017993,
			c2: 1.0328001,
			c4: 3.9920006
		  }
		},
		{
		  playerId: "cd0770ab-a2a6-46dd-95bf-e01b1874de31",
		  nickname: "LaviJr",
		  lvl: 5,
		  avatar: "https://assets.faceit-cdn.net/avatars/cd0770ab-a2a6-46dd-95bf-e01b1874de31_1550757598307.jpg",
		  stats: {
			c3: 1.17,
			c2: 1.8630002,
			c4: 2.8
		  },
		  teamStats: {
			c3: 0.7148,
			c2: 1.0527804,
			c4: 3.692
		  },
		  enemyStats: {
			c3: 0.7067996,
			c2: 1.0388198,
			c4: 3.558001
		  }
		},
		{
		  playerId: "6619ce5b-6c97-45da-916e-ae344f485d30",
		  nickname: "Crassuss",
		  lvl: 4,
		  avatar: "https://assets.faceit-cdn.net/avatars/6619ce5b-6c97-45da-916e-ae344f485d30_1585063481292.jpg",
		  stats: {
			c3: 0.70900005,
			c2: 0.94500005,
			c4: 5.4
		  },
		  teamStats: {
			c3: 0.71044,
			c2: 1.0681198,
			c4: 4.126
		  },
		  enemyStats: {
			c3: 0.68622005,
			c2: 1.0218003,
			c4: 3.7539978
		  }
		},
		{
		  playerId: "bf1a4910-2ea9-42f1-8048-a762a56820d6",
		  nickname: "W3Zy7",
		  lvl: 4,
		  avatar: "https://assets.faceit-cdn.net/avatars/bf1a4910-2ea9-42f1-8048-a762a56820d6_1551035571060.jpg",
		  stats: {
			c3: 0.8470001,
			c2: 1.2310002,
			c4: 6.3
		  },
		  teamStats: {
			c3: 0.7164,
			c2: 1.0518599,
			c4: 3.8899994
		  },
		  enemyStats: {
			c3: 0.7076001,
			c2: 1.0113199,
			c4: 3.8519998
		  }
		}
	  ],
	  win: ""
	},
	TeamB: {
	  teamName: "team_greensk1y",
	  averageElo: "1152",
	  players: [
		{
		  playerId: "03f08112-d3df-478e-a138-37232624555a",
		  nickname: "greensk1y",
		  lvl: 4,
		  avatar: "5_1596115963730.jpg",
		  stats: {
			c3: 0.94300014,
			c2: 1.434,
			c4: 4.7
		  },
		  teamStats: {
			c3: 0.66494,
			c2: 1.0013002,
			c4: 3.6179976
		  },
		  enemyStats: {
			c3: 0.66424,
			c2: 0.9966003,
			c4: 3.48
		  }
		},
		{
		  playerId: "76d91c84-0d3f-4e8e-9ac9-9f6428940d64",
		  nickname: "damascx",
		  lvl: 4,
		  avatar: "https://distribution.faceit-cdn.net/images/275a340c-dced-4dcb-8e4f-6d55e8792718.jpeg",
		  stats: {
			c3: 0.6830001,
			c2: 0.93799984,
			c4: 3.1
		  },
		  teamStats: {
			c3: 0.73810005,
			c2: 1.122,
			c4: 4.0360007
		  },
		  enemyStats: {
			c3: 0.6802999,
			c2: 0.95419943,
			c4: 3.9440007
		  }
		},
		{
			playerId: "76d91c84-0d3f-4e8e-9ac9-9f6428940d64",
			nickname: "dafdff",
			lvl: 4,
			avatar: "https://distribution.faceit-cdn.net/images/275a340c-dced-4dcb-8e4f-6d55e8792718.jpeg",
			stats: {
			  c3: 0.6830001,
			  c2: 0.93799984,
			  c4: 3.1
			},
			teamStats: {
			  c3: 0.73810005,
			  c2: 1.122,
			  c4: 4.0360007
			},
			enemyStats: {
			  c3: 0.6802999,
			  c2: 0.95419943,
			  c4: 3.9440007
			}
		},
		{
		playerId: "76d91c84-0d3f-4e8e-9ac9-9f6428940d64",
		nickname: "sddfsdfsdfsd",
		lvl: 4,
		avatar: "https://distribution.faceit-cdn.net/images/275a340c-dced-4dcb-8e4f-6d55e8792718.jpeg",
		stats: {
			c3: 0.6830001,
			c2: 0.93799984,
			c4: 3.1
		},
		teamStats: {
			c3: 0.73810005,
			c2: 1.122,
			c4: 4.0360007
		},
		enemyStats: {
			c3: 0.6802999,
			c2: 0.95419943,
			c4: 3.9440007
		}
		},
		{
		playerId: "76d91c84-0d3f-4e8e-9ac9-9f6428940d64",
		nickname: "jhhhhh",
		lvl: 4,
		avatar: "https://distribution.faceit-cdn.net/images/275a340c-dced-4dcb-8e4f-6d55e8792718.jpeg",
		stats: {
			c3: 0.6830001,
			c2: 0.93799984,
			c4: 3.1
		},
		teamStats: {
			c3: 0.73810005,
			c2: 1.122,
			c4: 4.0360007
		},
		enemyStats: {
			c3: 0.6802999,
			c2: 0.95419943,
			c4: 3.9440007
		}
		}
	  ],
	  win: ""
	}
}
