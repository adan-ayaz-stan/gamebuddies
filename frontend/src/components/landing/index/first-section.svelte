<script>
	import Navbar from '../navbar.svelte';
</script>

<div class="min-h-screen bg-black flex flex-col justify-evenly items-center">
	<Navbar />

	<!-- DEI Hired Div -->
	<div />

	<!-- Center Content -->
	<div class="w-full relative flex flex-col items-center">
		<!-- Visual Graphics -->
		<!-- Top Left -->
		<div class="absolute bottom-full left-0 w-1/4 h-3/4">
			<div
				style="clip-path: polygon(0 98%, 75% 98%, 100% 0, 100% 5%, 75% 100%, 0 100%);"
				class="w-full h-full bg-gradient-to-tr from-glowRed to-transparent"
			/>

			<!-- Svg -->
			<img
				src="/svgs/meet.svg"
				alt="meet icon"
				class="absolute bottom-0 right-1/4 translate-x-1/2 translate-y-1/2 backdrop-blur-sm h-16 w-16 border rounded-full text-glowRed border-glowRed p-3"
			/>
			<p class="absolute bottom-0 right-1/4 -translate-x-1/2">Let's MEET</p>
		</div>

		<!-- Bottom Right -->
		<div class="absolute top-full right-0 w-1/4 h-3/4">
			<div
				style="clip-path: polygon(0 98%, 75% 98%, 100% 0, 100% 3%, 75% 100%, 0 100%);"
				class="w-full h-full bg-gradient-to-l from-glowRed to-transparent"
			/>

			<!-- Svg -->
			<img
				src="/svgs/dock.svg"
				alt="meet icon"
				class="absolute bottom-0 right-1/4 translate-x-1/2 translate-y-1/2 backdrop-blur-sm h-16 w-16 border rounded-full text-glowRed border-glowRed p-3"
			/>
			<p class="absolute bottom-0 right-1/4 -translate-x-2/4">Let's STAY</p>
		</div>

		<!-- Bottom Left -->

		<div class="absolute top-full left-0 w-1/4 h-3/4">
			<div
				style="clip-path: polygon(0 98%, 75% 98%, 100% 0, 100% 3%, 75% 100%, 0 100%);"
				class="w-full h-full rotate-180 bg-gradient-to-l from-glowRed to-transparent"
			/>

			<!-- Svg -->
			<img
				src="/svgs/music-cast.svg"
				alt="meet icon"
				class="absolute bottom-full left-1/4 -translate-x-1/2 translate-y-1/2 backdrop-blur-sm h-16 w-16 border rounded-full text-glowRed border-glowRed p-3"
			/>
			<p class="absolute bottom-full left-1/4 translate-x-2/4">Let's CHILL</p>
		</div>

		<!-- Top Right -->

		<div class="absolute bottom-full right-0 w-1/4 h-3/4">
			<div
				style="clip-path: polygon(100% 100%, 25% 100%, 0 3%, 0 0, 25% 98%, 100% 99%);"
				class="w-full h-full bg-gradient-to-l from-glowRed to-transparent"
			/>

			<!-- Svg -->
			<img
				src="/svgs/gamepad.svg"
				alt="meet icon"
				class="absolute bottom-0 left-1/4 -translate-x-1/2 translate-y-1/2 backdrop-blur-sm h-16 w-16 border rounded-full text-glowRed border-glowRed p-3"
			/>
			<p class="absolute bottom-0 right-1/4 -translate-x-3/4">Let's PLAY</p>
		</div>

		<!-- Content -->
		<span>Find your best people</span>
		<h1 class="text-8xl tracking-wider uppercase font-bold text-lightRed mt-4">Gamebuddies</h1>
		<h2 class="text-4xl tracking-wider uppercase">{'>> Exclusive Matchmaking <<'}</h2>
	</div>

	<!-- Waitlist Button -->
	<button class="btn variant-filled-secondary font-bold mt-8 uppercase"
		><img
			src="https://api.iconify.design/hugeicons:target-03.svg"
			alt="target icon"
			class="h-6 w-6 mr-2"
		/> Join The Waitlist</button
	>
</div>
