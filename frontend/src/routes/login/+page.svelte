<script lang="ts">
	import { enhance } from '$app/forms';
	import BorderBitsCard from '$lib/components/border-bits-card.svelte';
	import { onMount } from 'svelte';
	import { cubicOut, quartOut } from 'svelte/easing';
	import { Tween } from 'svelte/motion';
	import type { ActionData, PageData } from './$types';

	let { data, form }: { data: PageData; form: ActionData } = $props();

	// ── Entry animations ────────────────────────────────────────────────────
	const wrapY = new Tween(24, { duration: 700, easing: quartOut });
	const wrapOpacity = new Tween(0, { duration: 600, easing: cubicOut });
	const badgeOpacity = new Tween(0, { duration: 400, easing: cubicOut, delay: 100 });
	const titleOpacity = new Tween(0, { duration: 500, easing: cubicOut, delay: 200 });
	const subtitleOpacity = new Tween(0, { duration: 400, easing: cubicOut, delay: 320 });
	const formOpacity = new Tween(0, { duration: 500, easing: cubicOut, delay: 420 });
	const footerOpacity = new Tween(0, { duration: 400, easing: cubicOut, delay: 550 });

	let submitting = $state(false);
	let showPassword = $state(false);

	onMount(() => {
		wrapY.set(0);
		wrapOpacity.set(1);
		badgeOpacity.set(1);
		titleOpacity.set(1);
		subtitleOpacity.set(1);
		formOpacity.set(1);
		footerOpacity.set(1);
	});
</script>

<svelte:head>
	<title>Sign In – GameBuddies</title>
</svelte:head>

<!-- Ambient glows -->
<div class="pointer-events-none fixed inset-0 overflow-hidden">
	<div class="absolute -top-32 -right-32 h-96 w-96 rounded-full bg-red-600/20 blur-[120px]"></div>
	<div class="absolute bottom-0 -left-32 h-80 w-80 rounded-full bg-red-800/15 blur-[100px]"></div>
</div>

<div
	class="relative flex min-h-screen items-center justify-center px-4"
	style="opacity: {wrapOpacity.current}; transform: translateY({wrapY.current}px)"
>
	<div class="w-full max-w-sm">
		<!-- Badge -->
		<div class="mb-6 flex justify-center" style="opacity: {badgeOpacity.current}">
			<span
				class="border border-red-500/40 bg-red-500/10 px-3 py-1 font-mono text-xs tracking-[0.2em] text-red-400 uppercase"
			>
				Authentication Protocol
			</span>
		</div>

		<!-- Title -->
		<div class="mb-8 text-center" style="opacity: {titleOpacity.current}">
			<h1 class="font-['Iceland'] text-5xl font-normal tracking-widest text-white uppercase">
				Access System
			</h1>
			<p
				class="mt-2 font-mono text-xs tracking-widest text-white/30 uppercase"
				style="opacity: {subtitleOpacity.current}"
			>
				Enter your operative credentials
			</p>
		</div>

		<!-- Form card -->
		<div style="opacity: {formOpacity.current}">
			<BorderBitsCard>
				<form
					method="POST"
					use:enhance={() => {
						submitting = true;
						return async ({ update }) => {
							submitting = false;
							update();
						};
					}}
					class="flex flex-col gap-5 p-2"
				>
					<!-- Success message after registration -->
					{#if data?.registered}
						<div
							class="border border-green-500/40 bg-green-500/10 px-3 py-2 font-mono text-xs text-green-400"
						>
							▸ Account created. You may now sign in.
						</div>
					{/if}

					<!-- Error message -->
					{#if form?.error}
						<div
							class="border border-red-500/50 bg-red-500/10 px-3 py-2 font-mono text-xs text-red-400"
						>
							▸ {form.error}
						</div>
					{/if}

					<!-- Username -->
					<div class="flex flex-col gap-1.5">
						<label for="username" class="font-mono text-xs tracking-widest text-white/50 uppercase">
							Username
						</label>
						<input
							id="username"
							name="username"
							type="text"
							autocomplete="username"
							value={form?.username ?? ''}
							required
							class="border border-white/15 bg-white/5 px-3 py-2.5 font-mono text-sm text-white placeholder-white/20
								   transition-all duration-200 outline-none
								   focus:border-red-500/60 focus:bg-red-950/20 focus:ring-1 focus:ring-red-500/30"
							placeholder="operative_handle"
						/>
					</div>

					<!-- Password -->
					<div class="flex flex-col gap-1.5">
						<label for="password" class="font-mono text-xs tracking-widest text-white/50 uppercase">
							Password
						</label>
						<div class="relative">
							<input
								id="password"
								name="password"
								type={showPassword ? 'text' : 'password'}
								autocomplete="current-password"
								required
								class="w-full border border-white/15 bg-white/5 px-3 py-2.5 pr-10 font-mono text-sm text-white placeholder-white/20
									   transition-all duration-200 outline-none
									   focus:border-red-500/60 focus:bg-red-950/20 focus:ring-1 focus:ring-red-500/30"
								placeholder="••••••••"
							/>
							<button
								type="button"
								onclick={() => (showPassword = !showPassword)}
								class="absolute top-1/2 right-3 -translate-y-1/2 font-mono text-xs text-white/30 transition-colors hover:text-white/60"
								aria-label="Toggle password visibility"
							>
								{showPassword ? 'hide' : 'show'}
							</button>
						</div>
					</div>

					<!-- Submit -->
					<button
						type="submit"
						disabled={submitting}
						class="relative mt-1 overflow-hidden border border-red-500/60 bg-red-500/10 py-3
							   font-mono text-xs tracking-[0.2em] text-red-400 uppercase transition-all duration-200
							   hover:border-red-400/80 hover:bg-red-500/20 hover:text-red-300
							   disabled:cursor-not-allowed disabled:opacity-40"
					>
						{#if submitting}
							<span class="animate-pulse">Authenticating...</span>
						{:else}
							Authenticate
						{/if}
					</button>
				</form>
			</BorderBitsCard>
		</div>

		<!-- Footer link -->
		<p
			class="mt-6 text-center font-mono text-xs text-white/30"
			style="opacity: {footerOpacity.current}"
		>
			New operative?
			<a
				href="/register"
				class="text-red-400/80 underline-offset-4 transition-colors hover:text-red-300 hover:underline"
			>
				Request access
			</a>
		</p>
	</div>
</div>
