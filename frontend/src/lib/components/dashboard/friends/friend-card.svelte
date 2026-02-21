<script lang="ts">
	import type { ColorVariant } from '$lib/types';

	/**
	 * @type {{ friendCode?: string, friendName?: string, isOnline?: boolean, color?: ColorVariant }}
	 */
	let { friendCode = '', friendName = '', isOnline = false, color = 'red' } = $props();
</script>

<div
	class="friend-card flex w-20 flex-col items-center transition-opacity duration-300"
	class:opacity-40={!isOnline}
	data-color={color}
>
	<img
		class="friend-card-image relative aspect-square w-20 border"
		src="https://www.shutterstock.com/image-vector/game-controller-on-fire-vector-260nw-2169427787.jpg"
		alt="gamer-profile-show"
	/>
	<div class="friend-card-code w-full p-1 text-center text-black shadow-2xl">
		{friendCode}
	</div>

	<h4 class="friend-card-name mt-2">
		{friendName}
	</h4>
</div>

<style>
	.friend-card[data-color='red'] {
		--color-main: #f87171; /* red-400 */
		--color-border: rgba(248, 113, 113, 0.7); /* red-400/70 */
		--color-shadow: #f87171; /* red-400 */
	}

	.friend-card[data-color='blue'] {
		--color-main: #60a5fa; /* blue-400 */
		--color-border: rgba(96, 165, 250, 0.7); /* blue-400/70 */
		--color-shadow: #60a5fa; /* blue-400 */
	}

	.friend-card[data-color='green'] {
		--color-main: #4ade80; /* green-400 */
		--color-border: rgba(74, 222, 128, 0.7); /* green-400/70 */
		--color-shadow: #4ade80; /* green-400 */
	}

	.friend-card[data-color='yellow'] {
		--color-main: #facc15; /* yellow-400 */
		--color-border: rgba(250, 204, 21, 0.7); /* yellow-400/70 */
		--color-shadow: #facc15; /* yellow-400 */
	}

	.friend-card[data-color='purple'] {
		--color-main: #a78bfa; /* purple-400 */
		--color-border: rgba(167, 139, 250, 0.7); /* purple-400/70 */
		--color-shadow: #a78bfa; /* purple-400 */
	}

	.friend-card[data-color='pink'] {
		--color-main: #f472b6; /* pink-400 */
		--color-border: rgba(244, 114, 182, 0.7); /* pink-400/70 */
		--color-shadow: #f472b6; /* pink-400 */
	}

	.friend-card-image {
		border-color: var(--color-border);
	}

	.friend-card-code {
		background-color: var(--color-main);
		box-shadow: 0 25px 50px -12px var(--color-shadow);
	}

	.friend-card-name {
		color: var(--color-main);
	}
</style>
