<script lang="ts">
	import { fade } from 'svelte/transition';
	// initialize stateful variable isHoveredOnce

	let isHoveredOnce = false;

	type MemberType = 'gold' | 'blue' | 'red';

	export let role: string;
	export let type: MemberType = 'gold';

	function getCardBacksideUrl(type: MemberType) {
		switch (type) {
			case 'gold':
				return '/graphics/gold-member-card-backside.png';
			case 'blue':
				return '/graphics/blue-member-card-backside.png';
			case 'red':
				return '/graphics/red-member-card-backside.png';
			default:
				return '/graphics/red-member-card-backside.png';
		}
	}

	function getColor(type: MemberType) {
		switch (type) {
			case 'gold':
				return '#EBE764';
			case 'blue':
				return '#5AC4FF';
			case 'red':
				return '#FF5A5F';
			default:
				return '#FF5A5F';
		}
	}

	$: backsideUrl = getCardBacksideUrl(type);
	$: color = getColor(type);

	export let githubUrl: string = '';
	export let linkedinUrl: string = '';
	export let telegramUrl: string = '';
	export let frontsideUrl: string;
</script>

<div class="space-y-8 relative">
	<div
		class="h-96 w-72 flip-container"
		role="button"
		tabindex="0"
		on:mouseenter={() => (isHoveredOnce = true)}
	>
		<div style="transform: rotateY({isHoveredOnce ? 180 : 0}deg);" class="frontside">
			<img src={backsideUrl} alt="profile" class="w-full h-full object-contain" />
		</div>
		<!--  -->
		<div style="transform: rotateY({isHoveredOnce ? 0 : 180}deg);" class="backside">
			<img src={frontsideUrl} alt="profile" class="w-full h-full object-contain" />
		</div>
	</div>

	{#if isHoveredOnce}
		<div
			transition:fade={{ delay: 250, duration: 600 }}
			class="absolute top-full left-0 w-full flex flex-col items-center gap-2"
		>
			<h2 style="color: {color};" class="h2 tracking-normal">{role}</h2>
			<!-- Social icons -->
			<div class="w-full flex flex-row items-center justify-between">
				<!-- Github Icon -->
				{#if githubUrl}
					<a
						href={githubUrl}
						style="background-color: {color};"
						class={`p-3 flex items-center justify-center rounded-md`}
					>
						<img
							src="https://api.iconify.design/mdi:github-face.svg?color=%23000000"
							alt="github icon"
							class="h-8 w-8"
						/>
					</a>
				{/if}
				<!-- Linkedin Icon -->
				{#if linkedinUrl}
					<a
						href={linkedinUrl}
						style="background-color: {color};"
						class="p-3 flex items-center justify-center rounded-md"
					>
						<img
							src="https://api.iconify.design/ion:logo-linkedin.svg?color=%23000000"
							alt="linkedin icon"
							class="h-8 w-8"
						/>
					</a>
				{/if}

				<!-- Telegram Icon -->
				{#if telegramUrl}
					<a
						href={telegramUrl}
						style="background-color: {color};"
						class="p-3 flex items-center justify-center rounded-md"
					>
						<img
							src="https://api.iconify.design/mingcute:telegram-fill.svg?color=%23000000"
							alt="telegram icon"
							class="h-8 w-8"
						/>
					</a>
				{/if}
			</div>
		</div>
	{/if}
</div>

<style>
	.backside {
		transform: rotateY(180deg);
		transform-style: preserve-3d;
		backface-visibility: hidden;
	}

	.frontside {
		transform: rotateY(0deg);
		transform-style: preserve-3d;
		backface-visibility: hidden;
	}

	.frontside,
	.backside {
		position: absolute;
		top: 0;
		left: 0;
		width: 100%;
		height: 100%;
		backface-visibility: hidden;
		transform-style: preserve-3d;
		transition: all 1.5s;
	}

	.flip-container {
		perspective: 2000px;
	}
</style>
