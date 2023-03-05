<template>
	<div>
		<span style="color: white; padding-left: 5px;">Buscar partida:</span>
		<div class="w1000">
			<form role="search">
				<input type="search" v-model="matchIdInput" placeholder="https://www.faceit.com/es/csgo/room/1-4d22f711-afae-474d-8443-9ca03943f843" required />
				<button type="button" @click="forceRerender">Go</button>
			</form>
		</div>
		<div v-if="match != null">
			<MatchResult :match="match" :key="componentKey" />
		</div>
		<div v-else>
			<div style="margin-top: 100px; color: #232323;" class="flex-column center">
				<h1>Busca una partida</h1>
				<img src="/images/undraw_personal_goals_re_iow7.svg" width="500" style="mix-blend-mode: overlay;">
			</div>
		</div>
	</div>
</template>

<script>
	import MatchResult from './match-result.vue';
	
	export default {
		components: {
			MatchResult
		},
		data() {
			return {
				componentKey: 0,
				matchId: "",
				matchIdInput: "",
				match: null,
			};
		},
		methods: {
			async forceRerender() {
				const re = new RegExp("[a-zA-Z0-9-]*$")
				const matchId = re.exec(this.matchIdInput)[0]
				
				
				const res = await fetch(`/analyze?q=${matchId}`);
				const finalRes = await res.json();
				this.match = Object.assign({}, finalRes);
				await this.$nextTick(); // Espera a que Vue termine de actualizar la propiedad "matchId".
			},
		}
	}
</script>


<style scoped>
	form {
		position: relative;
		width: 100%;
		background: #ff5500;
		border-radius: 0.7rem;
	}
	input, button {
		height: 4rem;
		font-family: 'Lato', sans-serif;
		border: 0;
		color: #2f2f2f;
		font-size: 1.5rem;
	}
	input[type="search"]::-webkit-search-cancel-button {
		display: none;
	}
	input[type="search"] {
		outline: 0;
		width: 100%;
		background: #cecece;
		padding: 0 1.6rem;
		border-radius: 0.7rem;
		appearance: none;
		transition: all 0.3s cubic-bezier(0, 0, 0.43, 1.49);
		transition-property: width, border-radius;
		z-index: 1;
		position: relative;
	}
	button {
		display: none;
		position: absolute;
		top: 0;
		right: 0;
		width: 6rem;
		font-weight: bold;
		background: #ff5500;
		border-radius: 0 0.7rem 0.7rem 0;
		cursor: pointer;
	}
	button:hover {
		box-shadow: rgb(180 72 22) 0px 0px 70px -5px;
		transition: ease-in-out 0.3s;
	}
	input:not(:placeholder-shown) {
		border-radius: 0.7rem 0 0 0.7rem;
		width: calc(100% - 6rem);
	}
	input:not(:placeholder-shown) + button {
		display: block;
	}
	label {
		position: absolute;
		clip: rect(1px, 1px, 1px, 1px);
		padding: 0;
		border: 0;
		height: 1px;
		width: 1px;
		overflow: hidden;
	}
</style>