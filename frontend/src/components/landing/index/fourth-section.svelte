<script lang="ts">
	import Marquee from 'svelte-fast-marquee';
	import ExcuseCard from './excuse-card.svelte';
</script>

<div class="min-h-screen flex flex-col items-center justify-between bg-black py-16">
	<h1 class="h1 text-lightRed text-center tracking-normal">TIRED OF EXCUSES?</h1>

	<!-- Excuse Cards -->
	<Marquee speed={50} gap="11%" class="flex items-center justify-evenly">
		<ExcuseCard content="I have to go now." type="yellow" />
		<ExcuseCard content="Already teamed up. Some time later." type="gray" />
		<ExcuseCard content="Busy right now. Play later." type="blue" />
		<ExcuseCard content="Have to sleep. We'll see morrow." type="yellow" />
	</Marquee>

	<!-- CTA -->
	<div class="space-y-3 text-center">
		<a href="/#waitlist-form" class="text-lightRed underline hidden md:block"
			><h2 class="h2 tracking-normal">JOIN THE GAMEBUDDIES PLATFORM!</h2></a
		>
		<h2 class="h2 tracking-normal">GET ON-SPOT MATCHED WITH YOUR BUDDIES AND ROLL ON</h2>

		<a href="/" class="btn variant-filled-primary font-light uppercase"
			>More about buds like you <img
				src="https://api.iconify.design/material-symbols:arrow-outward.svg?color=%23ffffff"
				alt="arrow outward"
				class="h-6 w-6 ml-2"
			/></a
		>
	</div>
</div>
