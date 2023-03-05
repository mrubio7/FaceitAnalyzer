<template>
	<section style="margin-top: 100px; color: white;">
		<div>
			<section :key="componentKey">

				<div>
					<input type="range" min="0" max="100" :value="match?.Result*100" disabled>
				</div>
				<div class="flex-justify-between">
					<div class="team">
						<div class="flex-justify-between">
							<span style="font-variant: all-small-caps;">{{ match?.TeamA?.teamName }}</span>
							<span class="elo">{{ match?.TeamA?.averageElo }}</span>
						</div>
						<div>
							<Player :player="match?.TeamA?.players[0]"/>
							<Player :player="match?.TeamA?.players[1]"/>
							<Player :player="match?.TeamA?.players[2]"/>
							<Player :player="match?.TeamA?.players[3]"/>
							<Player :player="match?.TeamA?.players[4]"/>
						</div>
					</div>
					<div class="team">
						<div class="flex-justify-between">
							<span style="font-variant: all-small-caps;">{{ match?.TeamB?.teamName }}</span>
							<span class="elo">{{ match?.TeamB?.averageElo }}</span>
						</div>
						<div>
							<Player :player="match?.TeamB?.players[0]"/>
							<Player :player="match?.TeamB?.players[1]"/>
							<Player :player="match?.TeamB?.players[2]"/>
							<Player :player="match?.TeamB?.players[3]"/>
							<Player :player="match?.TeamB?.players[4]"/>
						</div>
					</div>
				</div>

			</section>
		</div>
	</section>
</template>

<script>
	import Player from './player-result.vue'

	export default {
		components: {
			Player
		},
		props: ["matchcode"],
		data() {
			return {
				componentKey: 0,
				match: {}
			}
		},
		watch: {
			async matchcode() {
				const re = new RegExp("[a-zA-Z0-9-]*$")
				const matchId = re.exec(this.matchcode)[0]

				const res = await fetch(`/analyze?q=${matchId}`);
				const finalRes = await res.json();
				console.log(Object.assign({}, this.match));
				this.match = Object.assign({}, finalRes);
				console.log(Object.assign({}, this.match));
				console.log(this.componentKey);
			},
			match: {
				async handler() {
					console.log("handler " + this.componentKey);
					this.componentKey += 1;
				},
			},
		}
	}
</script>

<style scoped>
	.team {
		border: 1px solid #ff550021;
		border-radius: 10px;
		margin-top: 20px;
		padding: 15px;
		color: #5f5f5f;
	}
	.elo {
		color: #ff5500;
		font-weight: 900;
		margin-top: -5px;
	}

	input[type=range] {
		-webkit-appearance: none;
		margin: 10px 0;
		width: 100%;
	}
	input[type=range]:focus {
		outline: none;
	}
	input[type=range]::-webkit-slider-runnable-track {
		width: 100%;
		height: 12.8px;
		cursor: pointer;
		box-shadow: 0px 0px 0px #000000, 0px 0px 0px #0d0d0d;
		background: linear-gradient(90deg, #29292986,#ff550086,#29292986);
		border-radius: 25px;
		border: 1px solid #00010186;
	}
	input[type=range]::-webkit-slider-thumb {
		box-shadow: 0px 0px 0px #000000, 0px 0px 0px #0d0d0d;
		border: 0px solid #000000;
		height: 20px;
		width: 39px;
		border-radius: 7px;
		background: #ff5500;
		cursor: pointer;
		-webkit-appearance: none;
		margin-top: -3.6px;
	}
	input[type=range]:focus::-webkit-slider-runnable-track {
		background: #d4751c;
	}
	input[type=range]::-moz-range-track {
		width: 100%;
		height: 12.8px;
		cursor: pointer;
		animate: 0.2s;
		box-shadow: 0px 0px 0px #000000, 0px 0px 0px #0d0d0d;
		background: #f54d0b;
		border-radius: 25px;
		border: 0px solid #000101;
	}
	input[type=range]::-moz-range-thumb {
		box-shadow: 0px 0px 0px #000000, 0px 0px 0px #0d0d0d;
		border: 0px solid #000000;
		height: 20px;
		width: 39px;
		border-radius: 7px;
		background: #d64105;
		cursor: pointer;
	}
	input[type=range]::-ms-track {
		width: 100%;
		height: 12.8px;
		cursor: pointer;
		animate: 0.2s;
		background: transparent;
		border-color: transparent;
		border-width: 39px 0;
		color: transparent;
	}
	input[type=range]::-ms-fill-lower {
		background: #ac51b5;
		border: 0px solid #000101;
		border-radius: 50px;
		box-shadow: 0px 0px 0px #000000, 0px 0px 0px #0d0d0d;
	}
	input[type=range]::-ms-fill-upper {
		background: #ac51b5;
		border: 0px solid #000101;
		border-radius: 50px;
		box-shadow: 0px 0px 0px #000000, 0px 0px 0px #0d0d0d;
	}
	input[type=range]::-ms-thumb {
		box-shadow: 0px 0px 0px #000000, 0px 0px 0px #0d0d0d;
		border: 0px solid #000000;
		height: 20px;
		width: 39px;
		border-radius: 7px;
		background: #df3906;
		cursor: pointer;
	}
	input[type=range]:focus::-ms-fill-lower {
		background: #ac51b5;
	}
	input[type=range]:focus::-ms-fill-upper {
		background: #ac51b5;
	}
</style>