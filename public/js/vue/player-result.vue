<template>
	<div class="player-card">
		<div>
			<span class="nickname">
				{{ player?.nickname }}
			</span>
			<img class="avatar" :src="player?.avatar" width="40">
			<img class="lvl" :src="'https://beta.leetify.com/assets/images/rank-icons/faceit' + player?.lvl + '.svg'" width="30">

			<div class="flex-justify-between stats">
				<div class="flex-column center">
					<span>Avg. K/D</span>
					<span>{{ player?.stats?.c2.toFixed(2) }}</span>
				</div>
				<div class="flex-column center">
					<span>Avg. K/R</span>
					<span>{{ player?.stats?.c3.toFixed(2) }}</span>
				</div>
				<div class="flex-column center">
					<span>Avg. MVP</span>
					<span>{{ player?.stats?.c4 }}</span>
				</div>
			</div>
		</div>
	</div>
</template>

<script>
	export default {
		props: ['player']
	}
	
</script>


<style scoped>
	.player-card {
		background-color: rgb(26, 26, 26);
		width: 300px;
		height: 60px;
		padding: 10px;
		border-radius: 10px;
		margin-top: 20px;
	}

	.nickname {
		font-size: 14pt;
		color: #ff5500;
	}

	.avatar {
		border-radius: 5000px;
		box-shadow: rgba(0, 0, 0, 0.16) 0px 1px 4px, rgb(51, 51, 51) 0px 0px 0px 3px;
		float:right;
		margin-top: 5px;
	}
	.lvl {
		border-radius: 5000px;
		float:right;
		margin-right: -10px;
		margin-top: -5px;
	}

	.stats {
		font-size: 10pt;
		margin-right: 80px;
		padding-top: 10px;
		color:rgb(90, 90, 90);
	}

	.center {
		align-items: center
	}
</style>