<script>
	import { Howl } from 'howler';

	const hoverSound = new Howl({
		src: ['/sounds/hover1.mp3'],
		volume: 2,
		rate: 2,
		sprite: {
			hover: [50, 150]
		}
	});

	const clickSound = new Howl({
		src: ['/sounds/click.mp3'],
		volume: 2,
		rate: 2
	});

	function playClickSound() {
		clickSound.play();
	}

	function playHoverSound() {
		hoverSound.play('hover');
	}

	let { header, subheader = '', cardClassname = '', headerClassname = '', children } = $props();
</script>

<div
	class="group border-l-2 {cardClassname}"
	onmouseenter={playHoverSound}
	onclick={playClickSound}
	role="button"
	onkeydown={null}
	tabindex="0"
>
	<div class="header-section">
		<div
			class="relative flex overflow-hidden bg-primary px-2 text-primary-foreground {headerClassname}"
		>
			<h2 class="text-primary-foreground">{header}</h2>
			<img
				src="/svgs/trihex.svg"
				alt="hexes"
				class="absolute top-1/2 right-0 h-16 w-16 -translate-y-1/2"
			/>
		</div>
		{#if subheader && subheader.length > 0}
			<p class="bg-primary/10 p-2 text-muted-foreground uppercase">{subheader}</p>
		{/if}
	</div>

	<div class="p-4">
		{@render children()}
	</div>
</div>
