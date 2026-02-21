<script lang="ts">
	import BorderBitsCard from '$lib/components/border-bits-card.svelte';
	import { matchmakingStore } from '$lib/stores/matchmaking.svelte';
	import { onDestroy, onMount } from 'svelte';

	let { data } = $props();

	interface Game {
		id: number;
		title: string;
		subtitle: string;
		img_url: string;
	}

	const games = $derived((data.games ?? []) as Game[]);
	let selectedGameIds = $state<number[]>([]);

	// ── Connect on mount ───────────────────────────────────────────────────
	onMount(() => {
		matchmakingStore.connect();
	});

	onDestroy(() => {
		matchmakingStore.disconnect();
	});

	// ── Game selection ─────────────────────────────────────────────────────
	function toggleGame(id: number) {
		if (selectedGameIds.includes(id)) {
			selectedGameIds = selectedGameIds.filter((g) => g !== id);
		} else {
			selectedGameIds = [...selectedGameIds, id];
		}
	}

	// ── Queue actions ──────────────────────────────────────────────────────
	function handleJoin() {
		if (selectedGameIds.length === 0) return;
		matchmakingStore.joinQueue(selectedGameIds);
	}

	function handleLeave() {
		matchmakingStore.leaveQueue();
	}

	function handleClearMatch() {
		matchmakingStore.clearMatch();
	}

	// ── Pulse animation for queue indicator ────────────────────────────────
	let pulseCount = $state(0);
	let pulseInterval: ReturnType<typeof setInterval>;

	$effect(() => {
		if (matchmakingStore.inQueue) {
			pulseInterval = setInterval(() => pulseCount++, 1000);
		} else {
			clearInterval(pulseInterval);
		}
		return () => clearInterval(pulseInterval);
	});
</script>

<div class="min-h-full px-8 py-8">
	<!-- Header -->
	<div class="mb-8">
		<h1
			class="text-3xl font-bold tracking-widest text-white uppercase"
			style="font-family: 'Iceland', sans-serif"
		>
			Matchmaking
		</h1>
		<p class="mt-1 text-sm tracking-wide text-white/40">Find your next game buddy</p>
	</div>

	<!-- Match Found Modal -->
	{#if matchmakingStore.matchFound}
		<div class="fixed inset-0 z-50 flex items-center justify-center bg-black/80">
			<BorderBitsCard class="w-full max-w-sm">
				<div class="py-4 text-center">
					<div class="mb-4 text-5xl">⬡</div>
					<h2
						class="mb-2 text-2xl font-bold tracking-widest text-red-400 uppercase"
						style="font-family: 'Iceland', sans-serif"
					>
						Match Found!
					</h2>
					<p class="mb-6 text-sm text-white/60">Your game buddy is waiting…</p>
					{#if matchmakingStore.roomURL}
						<a
							href={matchmakingStore.roomURL}
							class="block w-full rounded border border-red-500 bg-red-500/20 py-2.5 text-sm font-semibold text-red-300 transition-colors hover:bg-red-500/30"
							style="font-family: 'Iceland', sans-serif; font-size: 1rem"
						>
							Enter Room →
						</a>
					{/if}
					<button
						onclick={handleClearMatch}
						class="mt-3 w-full rounded border border-white/15 py-2 text-xs text-white/40 transition-colors hover:text-white/70"
					>
						Dismiss
					</button>
				</div>
			</BorderBitsCard>
		</div>
	{/if}

	<div class="grid grid-cols-1 gap-6 lg:grid-cols-3">
		<!-- Game Selection -->
		<div class="lg:col-span-2">
			<p
				class="mb-4 text-xs tracking-widest text-white/40 uppercase"
				style="font-family: 'Iceland', sans-serif"
			>
				Select Games
			</p>
			{#if games.length === 0}
				<p class="text-sm text-white/30 italic">No games available.</p>
			{:else}
				<div class="grid grid-cols-1 gap-3 sm:grid-cols-2">
					{#each games as game}
						{@const selected = selectedGameIds.includes(game.id)}
						<button
							onclick={() => toggleGame(game.id)}
							disabled={matchmakingStore.inQueue}
							class="relative flex items-center gap-3 rounded border p-3 text-left transition-all
								{selected
								? 'border-red-500/60 bg-red-500/10'
								: 'border-white/10 bg-white/3 hover:border-white/25 hover:bg-white/5'}
								{matchmakingStore.inQueue ? 'cursor-not-allowed opacity-50' : 'cursor-pointer'}"
						>
							{#if game.img_url}
								<img
									src={game.img_url}
									alt={game.title}
									class="h-12 w-12 shrink-0 rounded object-cover"
								/>
							{:else}
								<div
									class="flex h-12 w-12 shrink-0 items-center justify-center rounded bg-white/10 text-xl"
								>
									◈
								</div>
							{/if}
							<div class="min-w-0 flex-1">
								<p
									class="truncate font-semibold text-white"
									style="font-family: 'Iceland', sans-serif"
								>
									{game.title}
								</p>
								<p class="truncate text-xs text-white/40">{game.subtitle}</p>
							</div>
							{#if selected}
								<span class="h-2.5 w-2.5 shrink-0 rounded-full bg-red-500"></span>
							{/if}
						</button>
					{/each}
				</div>
			{/if}
		</div>

		<!-- Queue Panel -->
		<div class="space-y-4">
			<!-- Connection Status -->
			<BorderBitsCard>
				<div class="flex items-center gap-2 py-1">
					<span
						class="h-2 w-2 rounded-full {matchmakingStore.connected
							? 'bg-green-500'
							: 'bg-red-500/60'}"
					></span>
					<p class="text-xs text-white/50">
						{matchmakingStore.connected ? 'Connected' : 'Connecting…'}
					</p>
				</div>
			</BorderBitsCard>

			<!-- Queue Status -->
			<BorderBitsCard>
				<div class="py-2">
					<p
						class="mb-3 text-xs tracking-widest text-white/40 uppercase"
						style="font-family: 'Iceland', sans-serif"
					>
						Queue
					</p>

					<div class="mb-4 flex items-end gap-2">
						<span
							class="leading-none text-red-400 tabular-nums"
							style="font-family: 'Iceland', sans-serif; font-size: 2.5rem"
						>
							{matchmakingStore.queueSize}
						</span>
						<span class="mb-1 text-xs text-white/40">players waiting</span>
					</div>

					{#if matchmakingStore.inQueue}
						<div class="mb-4 flex items-center gap-2">
							<span
								class="h-2 w-2 rounded-full bg-red-500 {pulseCount % 2 === 0
									? 'opacity-100'
									: 'opacity-30'} transition-opacity"
							></span>
							<p
								class="text-xs tracking-wider text-red-400/80 uppercase"
								style="font-family: 'Iceland', sans-serif"
							>
								In Queue
							</p>
						</div>
					{/if}

					{#if matchmakingStore.error}
						<p class="mb-3 text-xs text-red-400">{matchmakingStore.error}</p>
					{/if}

					{#if matchmakingStore.inQueue}
						<button
							onclick={handleLeave}
							class="w-full rounded border border-white/20 py-2.5 text-sm text-white/60 transition-colors hover:border-white/40 hover:text-white"
							style="font-family: 'Iceland', sans-serif; font-size: 0.95rem"
						>
							Leave Queue
						</button>
					{:else}
						<button
							onclick={handleJoin}
							disabled={selectedGameIds.length === 0 || !matchmakingStore.connected}
							class="w-full rounded border py-2.5 text-sm transition-colors
								{selectedGameIds.length > 0 && matchmakingStore.connected
								? 'border-red-500/60 bg-red-500/15 text-red-400 hover:bg-red-500/25 hover:text-red-300'
								: 'cursor-not-allowed border-white/10 bg-white/5 text-white/25'}"
							style="font-family: 'Iceland', sans-serif; font-size: 0.95rem"
						>
							{selectedGameIds.length === 0
								? 'Select at least 1 game'
								: `Join Queue (${selectedGameIds.length} game${selectedGameIds.length > 1 ? 's' : ''})`}
						</button>
					{/if}
				</div>
			</BorderBitsCard>

			<p class="text-center text-[10px] tracking-wide text-white/20">Matches require 5 players</p>
		</div>
	</div>
</div>
