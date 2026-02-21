<script lang="ts">
	import BorderBitsCard from '$lib/components/border-bits-card.svelte';
	import type { Stats } from '$lib/types/auth';
	import { onDestroy, onMount } from 'svelte';
	import { Tween } from 'svelte/motion';

	let { data, form } = $props();

	const stats = $derived(data.stats as Stats | null);

	// ── Live session timer ─────────────────────────────────────────────────
	let sessionIncrement = $state(0);
	let timerInterval: ReturnType<typeof setInterval>;

	onMount(() => {
		timerInterval = setInterval(() => {
			sessionIncrement++;
		}, 1000);
	});

	onDestroy(() => {
		clearInterval(timerInterval);
	});

	const totalSeconds = $derived((stats?.total_seconds ?? 0) + sessionIncrement);

	function formatTime(totalSecs: number) {
		const h = Math.floor(totalSecs / 3600);
		const m = Math.floor((totalSecs % 3600) / 60);
		const s = totalSecs % 60;
		return `${String(h).padStart(2, '0')}:${String(m).padStart(2, '0')}:${String(s).padStart(2, '0')}`;
	}

	// ── Animated counters ──────────────────────────────────────────────────
	const tweenMatches = new Tween(0, { duration: 1200 });
	const tweenSessions = new Tween(0, { duration: 1200 });

	$effect(() => {
		tweenMatches.set(stats?.total_matches ?? 0);
		tweenSessions.set(stats?.total_sessions ?? 0);
	});

	// ── Add friend form ────────────────────────────────────────────────────
	let friendUsername = $state('');
</script>

<div class="min-h-full px-8 py-8">
	<!-- Header -->
	<div class="mb-8">
		<h1
			class="text-3xl font-bold tracking-widest text-white uppercase"
			style="font-family: 'Iceland', sans-serif"
		>
			Dashboard
		</h1>
		<p class="mt-1 text-sm tracking-wide text-white/40">
			Welcome back, {data.user?.display_name ?? 'Player'}
		</p>
	</div>

	<!-- Stats Grid -->
	<div class="mb-10 grid grid-cols-1 gap-5 sm:grid-cols-3">
		<!-- Platform Time -->
		<BorderBitsCard>
			<div class="py-2">
				<p class="mb-3 text-xs tracking-widest text-white/40 uppercase">Platform Time</p>
				<p
					class="leading-none text-red-400 tabular-nums"
					style="font-family: 'Iceland', sans-serif; font-size: 2.2rem"
				>
					{formatTime(totalSeconds)}
				</p>
				<p class="mt-2 text-xs text-white/30">
					{Math.floor(totalSeconds / 3600)}h total accumulated
				</p>
			</div>
		</BorderBitsCard>

		<!-- Games Matched -->
		<BorderBitsCard>
			<div class="py-2">
				<p class="mb-3 text-xs tracking-widest text-white/40 uppercase">Games Matched</p>
				<p
					class="leading-none text-red-400 tabular-nums"
					style="font-family: 'Iceland', sans-serif; font-size: 2.8rem"
				>
					{Math.round(tweenMatches.current)}
				</p>
				<p class="mt-2 text-xs text-white/30">total matches found</p>
			</div>
		</BorderBitsCard>

		<!-- Sessions -->
		<BorderBitsCard>
			<div class="py-2">
				<p class="mb-3 text-xs tracking-widest text-white/40 uppercase">Sessions</p>
				<p
					class="leading-none text-red-400 tabular-nums"
					style="font-family: 'Iceland', sans-serif; font-size: 2.8rem"
				>
					{Math.round(tweenSessions.current)}
				</p>
				<p class="mt-2 text-xs text-white/30">login sessions</p>
			</div>
		</BorderBitsCard>
	</div>
</div>
