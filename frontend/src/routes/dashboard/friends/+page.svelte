<script lang="ts">
	import { enhance } from '$app/forms';
	import BorderBitsCard from '$lib/components/border-bits-card.svelte';
	import type { Friend, FriendRequest } from '$lib/types/auth';

	let { data, form } = $props();

	const friends = $derived((data.friends ?? []) as Friend[]);
	const requests = $derived((data.requests ?? []) as FriendRequest[]);

	const onlineFriends = $derived(friends.filter((f) => f.is_online));
	const offlineFriends = $derived(friends.filter((f) => !f.is_online));

	let addUsername = $state('');
</script>

<div class="min-h-full px-8 py-8">
	<!-- Header -->
	<div class="mb-8">
		<h1
			class="text-3xl font-bold tracking-widest text-white uppercase"
			style="font-family: 'Iceland', sans-serif"
		>
			Friends
		</h1>
		<p class="mt-1 text-sm tracking-wide text-white/40">Manage your gaming crew</p>
	</div>

	<div class="grid grid-cols-1 gap-8 lg:grid-cols-2">
		<!-- ── Left column ─────────────────────────────── -->
		<div class="flex flex-col gap-6">
			<!-- Send Friend Request -->
			<BorderBitsCard>
				<div class="py-2">
					<p
						class="mb-4 text-sm tracking-widest text-white/60 uppercase"
						style="font-family: 'Iceland', sans-serif"
					>
						Add a Friend
					</p>
					<form
						method="POST"
						action="?/sendRequest"
						use:enhance={() => {
							return ({ update }) => update({ reset: false });
						}}
						class="flex gap-2"
					>
						<input
							type="text"
							name="username"
							bind:value={addUsername}
							placeholder="Enter username…"
							class="flex-1 rounded border border-white/15 bg-white/5 px-3 py-2 text-sm text-white placeholder-white/25 transition-colors outline-none focus:border-red-500/60 focus:ring-1 focus:ring-red-500/30"
						/>
						<button
							type="submit"
							class="rounded border border-red-500/40 bg-red-500/10 px-4 py-2 text-sm text-red-400 transition-colors hover:bg-red-500/20 hover:text-red-300"
							style="font-family: 'Iceland', sans-serif"
						>
							Send Request
						</button>
					</form>
					{#if form?.sendError}
						<p class="mt-2 text-xs text-red-400">{form.sendError}</p>
					{/if}
					{#if form?.sendOk}
						<p class="mt-2 text-xs text-green-400">Friend request sent!</p>
					{/if}
				</div>
			</BorderBitsCard>

			<!-- Incoming Friend Requests -->
			<BorderBitsCard>
				<div class="py-2">
					<div class="mb-4 flex items-center justify-between">
						<p
							class="text-sm tracking-widest text-white/60 uppercase"
							style="font-family: 'Iceland', sans-serif"
						>
							Incoming Requests
						</p>
						{#if requests.length > 0}
							<span class="rounded-full bg-red-500/20 px-2 py-0.5 text-xs text-red-400">
								{requests.length}
							</span>
						{/if}
					</div>

					{#if requests.length === 0}
						<p class="text-xs text-white/30 italic">No pending requests</p>
					{:else}
						<div class="flex flex-col gap-2">
							{#each requests as req}
								<div
									class="flex items-center justify-between rounded border border-white/8 bg-white/3 px-3 py-2.5"
								>
									<!-- Avatar + name -->
									<div class="flex items-center gap-3">
										{#if req.avatar_url}
											<img
												src={req.avatar_url}
												alt={req.username}
												class="h-9 w-9 rounded-full object-cover ring-1 ring-white/20"
											/>
										{:else}
											<div
												class="flex h-9 w-9 items-center justify-center rounded-full bg-red-500/20 text-sm font-bold text-red-400 ring-1 ring-red-500/30"
											>
												{(req.display_name || req.username)[0].toUpperCase()}
											</div>
										{/if}
										<div>
											<p class="text-sm text-white">
												{req.display_name || req.username}
											</p>
											<p class="text-xs text-white/40">@{req.username}</p>
										</div>
									</div>

									<!-- Accept / Decline -->
									<div class="flex gap-2">
										<form method="POST" action="?/accept" use:enhance>
											<input type="hidden" name="id" value={req.id} />
											<button
												type="submit"
												class="rounded border border-green-500/40 bg-green-500/10 px-3 py-1 text-xs text-green-400 transition-colors hover:bg-green-500/20"
											>
												Accept
											</button>
										</form>
										<form method="POST" action="?/decline" use:enhance>
											<input type="hidden" name="id" value={req.id} />
											<button
												type="submit"
												class="rounded border border-white/15 bg-white/5 px-3 py-1 text-xs text-white/50 transition-colors hover:bg-white/10 hover:text-white/80"
											>
												Decline
											</button>
										</form>
									</div>
								</div>
							{/each}
						</div>
					{/if}
					{#if form?.actionError}
						<p class="mt-2 text-xs text-red-400">{form.actionError}</p>
					{/if}
				</div>
			</BorderBitsCard>
		</div>

		<!-- ── Right column — Friends list ────────────── -->
		<div>
			<BorderBitsCard>
				<div class="py-2">
					<div class="mb-4 flex items-center justify-between">
						<p
							class="text-sm tracking-widest text-white/60 uppercase"
							style="font-family: 'Iceland', sans-serif"
						>
							Your Friends
						</p>
						<span class="text-xs text-white/30">{friends.length} total</span>
					</div>

					{#if friends.length === 0}
						<p class="text-xs text-white/30 italic">
							No friends yet — send a request to get started!
						</p>
					{:else}
						<div class="flex flex-col gap-1">
							<!-- Online -->
							{#if onlineFriends.length > 0}
								<p class="mb-1 text-[10px] tracking-widest text-green-500/60 uppercase">
									Online — {onlineFriends.length}
								</p>
								{#each onlineFriends as friend}
									<div class="flex items-center justify-between rounded px-2 py-2 hover:bg-white/5">
										<div class="flex items-center gap-3">
											<div class="relative shrink-0">
												{#if friend.avatar_url}
													<img
														src={friend.avatar_url}
														alt={friend.username}
														class="h-8 w-8 rounded-full object-cover"
													/>
												{:else}
													<div
														class="flex h-8 w-8 items-center justify-center rounded-full bg-white/10 text-xs font-bold"
													>
														{(friend.display_name || friend.username)[0].toUpperCase()}
													</div>
												{/if}
												<span
													class="absolute -right-0.5 -bottom-0.5 h-2.5 w-2.5 rounded-full bg-green-500 ring-1 ring-black"
												></span>
											</div>
											<div>
												<p class="text-xs text-white/80">
													{friend.display_name || friend.username}
												</p>
												<p class="text-[10px] text-white/30">@{friend.username}</p>
											</div>
										</div>
										<form method="POST" action="?/remove" use:enhance>
											<input type="hidden" name="id" value={friend.id} />
											<button
												type="submit"
												class="rounded px-2 py-1 text-[10px] text-white/25 transition-colors hover:bg-red-500/10 hover:text-red-400"
											>
												Remove
											</button>
										</form>
									</div>
								{/each}
							{/if}

							<!-- Offline -->
							{#if offlineFriends.length > 0}
								{#if onlineFriends.length > 0}
									<div class="my-2 border-t border-white/8"></div>
								{/if}
								<p class="mb-1 text-[10px] tracking-widest text-white/25 uppercase">
									Offline — {offlineFriends.length}
								</p>
								{#each offlineFriends as friend}
									<div
										class="flex items-center justify-between rounded px-2 py-2 opacity-50 hover:opacity-70"
									>
										<div class="flex items-center gap-3">
											<div class="relative shrink-0">
												{#if friend.avatar_url}
													<img
														src={friend.avatar_url}
														alt={friend.username}
														class="h-8 w-8 rounded-full object-cover grayscale"
													/>
												{:else}
													<div
														class="flex h-8 w-8 items-center justify-center rounded-full bg-white/10 text-xs font-bold"
													>
														{(friend.display_name || friend.username)[0].toUpperCase()}
													</div>
												{/if}
												<span
													class="absolute -right-0.5 -bottom-0.5 h-2.5 w-2.5 rounded-full bg-white/20 ring-1 ring-black"
												></span>
											</div>
											<div>
												<p class="text-xs text-white/60">
													{friend.display_name || friend.username}
												</p>
												<p class="text-[10px] text-white/25">@{friend.username}</p>
											</div>
										</div>
										<form method="POST" action="?/remove" use:enhance>
											<input type="hidden" name="id" value={friend.id} />
											<button
												type="submit"
												class="rounded px-2 py-1 text-[10px] text-white/25 transition-colors hover:bg-red-500/10 hover:text-red-400"
											>
												Remove
											</button>
										</form>
									</div>
								{/each}
							{/if}
						</div>
					{/if}
				</div>
			</BorderBitsCard>
		</div>
	</div>
</div>
