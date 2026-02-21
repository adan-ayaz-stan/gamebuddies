<script lang="ts">
	import { page } from '$app/state';
	import BorderBitsCard from '$lib/components/border-bits-card.svelte';
	import type { Friend } from '$lib/types/auth';

	let { data, children } = $props();

	const user = $derived(data.user);
	const friends = $derived((data.friends ?? []) as Friend[]);
	const pendingRequestCount = $derived((data.pendingRequestCount ?? 0) as number);

	const onlineFriends = $derived(friends.filter((f) => f.is_online));
	const offlineFriends = $derived(friends.filter((f) => !f.is_online));

	const navLinks = [
		{ href: '/dashboard', label: 'Dashboard', icon: '⬡' },
		{ href: '/dashboard/matchmaking', label: 'Matchmaking', icon: '◈' },
		{ href: '/dashboard/friends', label: 'Friends', icon: '◌' },
		{ href: '/dashboard/messages', label: 'Messages', icon: '◉' }
	];

	function isActive(href: string) {
		if (href === '/dashboard') return page.url.pathname === '/dashboard';
		return page.url.pathname.startsWith(href);
	}
</script>

<div class="flex h-screen overflow-hidden bg-black text-white">
	<!-- ── Left Sidebar ─────────────────────────────────────────────────── -->
	<aside
		class="flex h-full w-64 shrink-0 flex-col gap-4 overflow-y-auto border-r border-white/10 bg-black/90 px-4 py-5"
	>
		<!-- Profile Card -->
		<a href="/profile" class="block transition-opacity hover:opacity-80">
			<BorderBitsCard class="w-full">
				<div class="flex flex-col items-center gap-3 py-2">
					{#if user?.avatar_url}
						<img
							src={user.avatar_url}
							alt="avatar"
							class="h-16 w-16 rounded-full object-cover ring-2 ring-red-500"
						/>
					{:else}
						<div
							class="flex h-16 w-16 items-center justify-center rounded-full bg-red-500/20 text-2xl font-bold text-red-400 ring-2 ring-red-500"
						>
							{(user?.display_name ?? user?.username ?? '?')[0].toUpperCase()}
						</div>
					{/if}
					<div class="text-center">
						<p
							class="font-semibold tracking-wide text-white"
							style="font-family: 'Iceland', sans-serif"
						>
							{user?.display_name ?? 'Player'}
						</p>
						<p class="text-xs text-white/50">@{user?.username ?? ''}</p>
					</div>
				</div>
			</BorderBitsCard>
		</a>

		<!-- Nav Links -->
		<nav class="flex flex-col gap-1">
			{#each navLinks as link}
				<a
					href={link.href}
					class="flex items-center gap-3 rounded px-3 py-2.5 text-sm tracking-wide transition-colors
						{isActive(link.href)
						? 'border-l-2 border-red-500 bg-red-500/15 text-red-400'
						: 'border-l-2 border-transparent text-white/60 hover:bg-white/5 hover:text-white'}"
					style="font-family: 'Iceland', sans-serif; font-size: 1rem"
				>
					<span class="text-lg leading-none">{link.icon}</span>
					{link.label}
					{#if link.href === '/dashboard/friends' && pendingRequestCount > 0}
						<span class="ml-auto rounded-full bg-red-500/30 px-1.5 py-0.5 text-[10px] text-red-400">
							{pendingRequestCount}
						</span>
					{/if}
				</a>
			{/each}
		</nav>

		<!-- Friends Section -->
		<div class="flex-1">
			<p class="mb-2 px-1 text-xs tracking-widest text-white/30 uppercase">Friends</p>

			{#if friends.length === 0}
				<p class="px-1 text-xs text-white/30 italic">No friends yet</p>
			{:else}
				{#if onlineFriends.length > 0}
					<div class="mb-3">
						<p class="mb-1 px-1 text-[10px] tracking-widest text-green-500/60 uppercase">
							Online — {onlineFriends.length}
						</p>
						{#each onlineFriends as friend}
							<div class="flex items-center gap-2 rounded px-2 py-1.5 hover:bg-white/5">
								<div class="relative">
									{#if friend.avatar_url}
										<img
											src={friend.avatar_url}
											alt={friend.username}
											class="h-7 w-7 rounded-full object-cover"
										/>
									{:else}
										<div
											class="flex h-7 w-7 items-center justify-center rounded-full bg-white/10 text-xs font-bold"
										>
											{(friend.display_name || friend.username || '?')[0].toUpperCase()}
										</div>
									{/if}
									<span
										class="absolute -right-0.5 -bottom-0.5 h-2.5 w-2.5 rounded-full bg-green-500 ring-1 ring-black"
									></span>
								</div>
								<div class="min-w-0">
									<p class="truncate text-xs text-white/80">{friend.display_name}</p>
									<p class="truncate text-[10px] text-white/30">@{friend.username}</p>
								</div>
							</div>
						{/each}
					</div>
				{/if}

				{#if offlineFriends.length > 0}
					<div>
						<p class="mb-1 px-1 text-[10px] tracking-widest text-white/25 uppercase">
							Offline — {offlineFriends.length}
						</p>
						{#each offlineFriends as friend}
							<div class="flex items-center gap-2 rounded px-2 py-1.5 opacity-50 hover:opacity-70">
								<div class="relative">
									{#if friend.avatar_url}
										<img
											src={friend.avatar_url}
											alt={friend.username}
											class="h-7 w-7 rounded-full object-cover grayscale"
										/>
									{:else}
										<div
											class="flex h-7 w-7 items-center justify-center rounded-full bg-white/10 text-xs font-bold"
										>
											{(friend.display_name || friend.username || '?')[0].toUpperCase()}
										</div>
									{/if}
									<span
										class="absolute -right-0.5 -bottom-0.5 h-2.5 w-2.5 rounded-full bg-white/20 ring-1 ring-black"
									></span>
								</div>
								<div class="min-w-0">
									<p class="truncate text-xs text-white/60">{friend.display_name}</p>
									<p class="truncate text-[10px] text-white/25">@{friend.username}</p>
								</div>
							</div>
						{/each}
					</div>
				{/if}
			{/if}
		</div>

		<!-- Bottom: Profile + Logout -->
		<div class="mt-auto flex flex-col gap-2 border-t border-white/10 pt-4">
			<a
				href="/profile"
				class="flex items-center gap-2 rounded px-3 py-2 text-sm text-white/50 transition-colors hover:bg-white/5 hover:text-white"
			>
				<span class="text-base">⚙</span>
				<span style="font-family: 'Iceland', sans-serif; font-size: 0.95rem"
					>Profile & Settings</span
				>
			</a>
			<form method="POST" action="/logout">
				<button
					type="submit"
					class="flex w-full items-center gap-2 rounded px-3 py-2 text-sm text-red-500/70 transition-colors hover:bg-red-500/10 hover:text-red-400"
				>
					<span class="text-base">⏻</span>
					<span style="font-family: 'Iceland', sans-serif; font-size: 0.95rem">Logout</span>
				</button>
			</form>
		</div>
	</aside>

	<!-- ── Main Content ──────────────────────────────────────────────────── -->
	<main class="flex-1 overflow-y-auto">
		{@render children?.()}
	</main>
</div>
