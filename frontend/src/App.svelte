<script>
	import Galaxy from './lib/Galaxy.svelte';
	import Popup from './lib/Popup.svelte';
	// var
	let planet = '';
	let speed = '';

	let showPopup = false;
	//
	// planetSpeed
	function planetSpeed(name) {
		window.backend.main.Galaxy.PlanetSpeed(name)
			.then((r) => {
				planet = name;
				speed = r;
			})
			.catch((e) => {
				console.log(e);
			});
	}
	//
	function galaxySpeed() {
		window.backend.main.Galaxy.Speed()
			.then((r) => {
				planet = 'galaxy';
				speed = r;
			})
			.catch((e) => {
				console.log(e);
			});
	}
</script>

<main>
	{#if showPopup}
		<Popup
			planet="{planet}"
			speed="{speed}"
			on:close="{() => {
				showPopup = !showPopup;
			}}"
		/>
	{/if}
	<!--  -->
	<Galaxy
		on:select="{(e) => {
			planetSpeed(e.detail);
			showPopup = !showPopup;
		}}"
		on:check="{() => {
			galaxySpeed();
			showPopup = !showPopup;
		}}"
	/>
</main>

<style global>
	@tailwind base;
	@tailwind components;
	@tailwind utilities;
</style>
