<template>
	<div>
		<span style="color: white; padding-left: 5px;">Buscar partida:</span>
		<div class="w1000">
			<form role="search">
				<input type="search" v-model="matchIdInput" placeholder="https://www.faceit.com/es/csgo/room/1-4d22f711-afae-474d-8443-9ca03943f843" required />
				<button type="button" @click="forceRerender">Go</button>
			</form>
		</div>

		<div v-if="loading == false">

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
		<div v-else>

			<div style="margin-top: 200px; color: #232323;" class="flex-column center">
				<main>
					<div class="dank-ass-loader">
						<div class="row">
							<div class="arrow up outer outer-18"></div>
							<div class="arrow down outer outer-17"></div>
							<div class="arrow up outer outer-16"></div>
							<div class="arrow down outer outer-15"></div>
							<div class="arrow up outer outer-14"></div>
						</div>
						<div class="row">
							<div class="arrow up outer outer-1"></div>
							<div class="arrow down outer outer-2"></div>
							<div class="arrow up inner inner-6"></div>
							<div class="arrow down inner inner-5"></div>
							<div class="arrow up inner inner-4"></div>
							<div class="arrow down outer outer-13"></div>
							<div class="arrow up outer outer-12"></div>
						</div>
						<div class="row">
							<div class="arrow down outer outer-3"></div>
							<div class="arrow up outer outer-4"></div>
							<div class="arrow down inner inner-1"></div>
							<div class="arrow up inner inner-2"></div>
							<div class="arrow down inner inner-3"></div>
							<div class="arrow up outer outer-11"></div>
							<div class="arrow down outer outer-10"></div>
						</div>
						<div class="row">
							<div class="arrow down outer outer-5"></div>
							<div class="arrow up outer outer-6"></div>
							<div class="arrow down outer outer-7"></div>
							<div class="arrow up outer outer-8"></div>
							<div class="arrow down outer outer-9"></div>
						</div>
					</div>
					</main>
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
				loading: false,
			};
		},
		methods: {
			async forceRerender() {
				this.loading = true;
				await this.$nextTick();

				const re = new RegExp("[a-zA-Z0-9-]*$")
				const matchId = re.exec(this.matchIdInput)[0]
				
				
				const res = await fetch(`/analyze?q=${matchId}`);
				const finalRes = await res.json();
				this.match = Object.assign({}, finalRes);
				this.loading = false;
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




	main {
		display: flex;
		justify-content: center;
		align-items: center;
	}
	.dank-ass-loader {
		display: flex;
		flex-direction: column;
		align-items: center;
	}
	.dank-ass-loader .row {
		display: flex;
	}
	.arrow {
		width: 0;
		height: 0;
		margin: 0 -6px;
		border-left: 12px solid transparent;
		border-right: 12px solid transparent;
		border-bottom: 21.6px solid #fd7000;
		animation: blink 1s infinite;
		filter: drop-shadow(0 0 18px #fd7000);
	}
	.arrow.down {
		transform: rotate(180deg);
	}
	.arrow.outer-1 {
		animation-delay: -0.0555555556s;
	}
	.arrow.outer-2 {
		animation-delay: -0.1111111111s;
	}
	.arrow.outer-3 {
		animation-delay: -0.1666666667s;
	}
	.arrow.outer-4 {
		animation-delay: -0.2222222222s;
	}
	.arrow.outer-5 {
		animation-delay: -0.2777777778s;
	}
	.arrow.outer-6 {
		animation-delay: -0.3333333333s;
	}
	.arrow.outer-7 {
		animation-delay: -0.3888888889s;
	}
	.arrow.outer-8 {
		animation-delay: -0.4444444444s;
	}
	.arrow.outer-9 {
		animation-delay: -0.5s;
	}
	.arrow.outer-10 {
		animation-delay: -0.5555555556s;
	}
	.arrow.outer-11 {
		animation-delay: -0.6111111111s;
	}
	.arrow.outer-12 {
		animation-delay: -0.6666666667s;
	}
	.arrow.outer-13 {
		animation-delay: -0.7222222222s;
	}
	.arrow.outer-14 {
		animation-delay: -0.7777777778s;
	}
	.arrow.outer-15 {
		animation-delay: -0.8333333333s;
	}
	.arrow.outer-16 {
		animation-delay: -0.8888888889s;
	}
	.arrow.outer-17 {
		animation-delay: -0.9444444444s;
	}
	.arrow.outer-18 {
		animation-delay: -1s;
	}
	.arrow.inner-1 {
		animation-delay: -0.1666666667s;
	}
	.arrow.inner-2 {
		animation-delay: -0.3333333333s;
	}
	.arrow.inner-3 {
		animation-delay: -0.5s;
	}
	.arrow.inner-4 {
		animation-delay: -0.6666666667s;
	}
	.arrow.inner-5 {
		animation-delay: -0.8333333333s;
	}
	.arrow.inner-6 {
		animation-delay: -1s;
	}
	@keyframes blink {
		0% {
			opacity: 0.1;
		}
		30% {
			opacity: 1;
		}
		100% {
			opacity: 0.1;
		}
	}

</style>