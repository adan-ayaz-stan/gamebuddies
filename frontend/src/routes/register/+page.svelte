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
	const footerOpacity = new Tween(0, { duration: 400, easing: cubicOut, delay: 600 });

	let submitting = $state(false);
	let showPassword = $state(false);

	// Live password strength
	let password = $state('');
	const strength = $derived(getStrength(password));

	function getStrength(p: string): { score: number; label: string; color: string } {
		if (!p) return { score: 0, label: '', color: '' };
		let score = 0;
		if (p.length >= 8) score++;
		if (p.length >= 12) score++;
		if (/[A-Z]/.test(p)) score++;
		if (/[0-9]/.test(p)) score++;
		if (/[^A-Za-z0-9]/.test(p)) score++;
		if (score <= 1) return { score, label: 'Weak', color: 'bg-red-600' };
		if (score <= 2) return { score, label: 'Fair', color: 'bg-orange-500' };
		if (score <= 3) return { score, label: 'Good', color: 'bg-yellow-400' };
		return { score, label: 'Strong', color: 'bg-green-500' };
	}

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
	<title>Register – GameBuddies</title>
</svelte:head>

<!-- Ambient glows -->
<div class="pointer-events-none fixed inset-0 overflow-hidden">
	<div class="absolute -top-32 -right-32 h-96 w-96 rounded-full bg-red-600/20 blur-[120px]"></div>
	<div class="absolute bottom-0 -left-32 h-80 w-80 rounded-full bg-red-800/15 blur-[100px]"></div>
</div>

<div
	class="relative flex min-h-screen items-center justify-center px-4 py-12"
	style="opacity: {wrapOpacity.current}; transform: translateY({wrapY.current}px)"
>
	<div class="w-full max-w-sm">
		<!-- Badge -->
		<div class="mb-6 flex justify-center" style="opacity: {badgeOpacity.current}">
			<span
				class="border border-red-500/40 bg-red-500/10 px-3 py-1 font-mono text-xs tracking-[0.2em] text-red-400 uppercase"
			>
				New Operative Enrollment
			</span>
		</div>

		<!-- Title -->
		<div class="mb-8 text-center" style="opacity: {titleOpacity.current}">
			<h1 class="font-['Iceland'] text-5xl font-normal tracking-widest text-white uppercase">
				Join the System
			</h1>
			<p
				class="mt-2 font-mono text-xs tracking-widest text-white/30 uppercase"
				style="opacity: {subtitleOpacity.current}"
			>
				Create your operative profile
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
					<!-- Error message -->
					{#if form?.error}
						<div
							class="border border-red-500/50 bg-red-500/10 px-3 py-2 font-mono text-xs text-red-400"
						>
							▸ {form.error}
						</div>
					{/if}

					<!-- Display name -->
					<div class="flex flex-col gap-1.5">
						<label
							for="display_name"
							class="font-mono text-xs tracking-widest text-white/50 uppercase"
						>
							Display Name
						</label>
						<input
							id="display_name"
							name="display_name"
							type="text"
							autocomplete="name"
							value={form?.display_name ?? ''}
							required
							class="border border-white/15 bg-white/5 px-3 py-2.5 font-mono text-sm text-white placeholder-white/20
								   transition-all duration-200 outline-none
								   focus:border-red-500/60 focus:bg-red-950/20 focus:ring-1 focus:ring-red-500/30"
							placeholder="Ghost Reaper"
						/>
					</div>

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
							placeholder="ghost_reaper"
						/>
						<p class="font-mono text-[10px] text-white/25">Letters, numbers, underscores only.</p>
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
								autocomplete="new-password"
								bind:value={password}
								required
								class="w-full border border-white/15 bg-white/5 px-3 py-2.5 pr-10 font-mono text-sm text-white placeholder-white/20
									   transition-all duration-200 outline-none
									   focus:border-red-500/60 focus:bg-red-950/20 focus:ring-1 focus:ring-red-500/30"
								placeholder="Min. 8 characters"
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

						<!-- Strength bar -->
						{#if password.length > 0}
							<div class="flex items-center gap-2">
								<div class="flex flex-1 gap-0.5">
									{#each Array(5) as _, i}
										<div
											class="h-0.5 flex-1 transition-all duration-300 {i < strength.score
												? strength.color
												: 'bg-white/10'}"
										></div>
									{/each}
								</div>
								<span class="font-mono text-[10px] text-white/40">{strength.label}</span>
							</div>
						{/if}
					</div>

					<!-- Confirm password -->
					<div class="flex flex-col gap-1.5">
						<label
							for="confirm_password"
							class="font-mono text-xs tracking-widest text-white/50 uppercase"
						>
							Confirm Password
						</label>
						<input
							id="confirm_password"
							name="confirm_password"
							type={showPassword ? 'text' : 'password'}
							autocomplete="new-password"
							required
							class="border border-white/15 bg-white/5 px-3 py-2.5 font-mono text-sm text-white placeholder-white/20
								   transition-all duration-200 outline-none
								   focus:border-red-500/60 focus:bg-red-950/20 focus:ring-1 focus:ring-red-500/30"
							placeholder="Repeat password"
						/>
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
							<span class="animate-pulse">Enrolling...</span>
						{:else}
							Enroll Operative
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
			Already enrolled?
			<a
				href="/login"
				class="text-red-400/80 underline-offset-4 transition-colors hover:text-red-300 hover:underline"
			>
				Access system
			</a>
		</p>
	</div>
</div>
