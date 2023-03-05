<template>
  <div class="header">
	<div class="header-items">
		<h3>faceit AI analyzer</h3>
		<a class="about" href="https://www.linkedin.com/in/mrubioa/">mrubio</a>
	</div>
  </div>
</template>

<style scoped>
	.header {
		top: 0;
		left: 0;
		width: 100vw;
		background-color: #1a1a1a;
		color:#ff5500;
		position: fixed;
		font-variant: petite-caps;
	}
	.header-items {
		display:flex;
		justify-content: space-between;
		padding-left: 20px;
		padding-right: 50px;
	}
	.about {
		margin-top: 22px;
		color: rgb(92, 92, 92);
		font-size: 10pt;
		text-decoration: none;
		font-variant: normal;
		font-weight: 600;
	}
	.about:hover {
		color: rgb(153, 153, 153);
		transition: 0.5s;
	}
</style>
