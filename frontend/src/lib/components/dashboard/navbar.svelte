<script>
	import { onMount } from 'svelte';
	import { Button } from '../ui/button';

	let serverTime = '8:42';
	let localTime = '15:42';

	// Update times every second
	onMount(() => {
		const interval = setInterval(() => {
			const now = new Date();
			localTime = now.toLocaleTimeString('en-US', {
				hour12: false,
				hour: '2-digit',
				minute: '2-digit'
			});
			// You can adjust server time logic as needed
			serverTime = new Date(now.getTime() - 7 * 60 * 60 * 1000).toLocaleTimeString('en-US', {
				hour12: false,
				hour: '2-digit',
				minute: '2-digit'
			});
		}, 1000);

		return () => clearInterval(interval);
	});
</script>

<nav
	class="flex items-center justify-between border-b border-white/13 bg-black p-4 px-8 text-xs text-white"
>
	<div class="flex items-center gap-6">
		<div class="flex items-center gap-2">
			<h3 class="text-2xl font-bold text-success">48</h3>
			<span class="text-[10px] tracking-wider text-gray-500">LEVEL</span>
		</div>

		<div class="flex items-center gap-2">
			<Button size="icon" variant="outline">+</Button>
			<h3 class="text-2xl font-bold text-success">1,248</h3>
			<span class="text-[10px] tracking-wider text-gray-500">COINS AWARDED</span>
		</div>
	</div>

	<div class="flex items-center gap-4 text-base">
		<span class="tracking-wider">CREDITS</span>
		<div class="flex gap-2">
			<span class="tracking-wider text-gray-500">SERVER TIME</span>
			<span class="font-bold text-white">{serverTime}</span>
		</div>
		<div class="flex gap-2">
			<span class="tracking-wider text-gray-500">LOCAL TIME</span>
			<span class="font-bold text-white">{localTime}</span>
		</div>
	</div>
</nav>
