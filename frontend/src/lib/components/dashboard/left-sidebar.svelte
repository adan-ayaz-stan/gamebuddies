<script>
	import BorderBitsCard from '../border-bits-card.svelte';
	import { Popover } from '../ui/popover';
	import PopoverContent from '../ui/popover/popover-content.svelte';
	import PopoverTrigger from '../ui/popover/popover-trigger.svelte';
</script>

<div class="flex h-full flex-col gap-4">
	<BorderBitsCard class="mx-auto w-fit max-w-sm">
		<img
			src="https://c4.wallpaperflare.com/wallpaper/270/674/468/the-mandalorian-the-mandalorian-character-minimalism-red-black-hd-wallpaper-preview.jpg"
			alt="profile"
			width="194"
			height="175"
			class="aspect-square object-cover"
		/>
	</BorderBitsCard>

	<div class="space-y-2">
		<p>Name</p>
		<h3>Adan Ayaz</h3>
	</div>
	<div class="space-y-2">
		<p>Games Played</p>
		<h3>127</h3>
	</div>

	<div class="space-y-2">
		<p>Buddy Rating</p>
		<h3>4.8/5.0</h3>
	</div>

	<div class="space-y-2">
		<p>Favorite Game</p>
		<Popover>
			<PopoverTrigger>
				<h3>Valorant</h3>
			</PopoverTrigger>
			<PopoverContent side="right">
				<div class="max-w-xs p-4">
					<h4 class="mb-2 font-semibold">Valorant</h4>
					<p class="mb-3 leading-tight tracking-tight text-muted-foreground">
						A tactical first-person shooter developed by Riot Games. Known for its precise gunplay,
						unique agent abilities, and competitive ranked system.
					</p>
					<div class="space-y-1">
						<p><span class="font-medium">Genre:</span> Tactical FPS</p>
						<p><span class="font-medium">Developer:</span> Riot Games</p>
						<p><span class="font-medium">Release:</span> 2020</p>
					</div>
				</div>
			</PopoverContent>
		</Popover>
	</div>

	<div class="mt-auto uppercase">
		<p class="text-xl text-primary">Motto:</p>
		<p class="text-white/73">
			You did not break me <br />
			I'm still fighting for peace
		</p>
	</div>
</div>
