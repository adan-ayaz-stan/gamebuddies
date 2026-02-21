<script lang="ts">
	import { enhance } from '$app/forms';
	import BorderBitsCard from '$lib/components/border-bits-card.svelte';

	let { data, form } = $props();

	const user = $derived(data.user);

	let displayName = $state(data.user?.display_name ?? '');
	let avatarPreview = $state<string | null>(null);
	let avatarFile = $state<File | null>(null);
	let uploading = $state(false);

	function handleFileSelect(e: Event) {
		const input = e.target as HTMLInputElement;
		const file = input.files?.[0] ?? null;
		avatarFile = file;
		if (file) {
			const reader = new FileReader();
			reader.onload = () => (avatarPreview = reader.result as string);
			reader.readAsDataURL(file);
		} else {
			avatarPreview = null;
		}
	}

	// Keep displayName in sync if user data changes
	$effect(() => {
		if (user?.display_name) displayName = user.display_name;
	});
</script>

<div class="min-h-screen bg-black px-6 py-8 text-white">
	<!-- Back link -->
	<a
		href="/dashboard"
		class="mb-8 inline-flex items-center gap-2 text-sm text-white/40 transition-colors hover:text-white"
	>
		← Back to Dashboard
	</a>

	<div class="mx-auto max-w-2xl">
		<h1
			class="mb-8 text-3xl font-bold tracking-widest text-white uppercase"
			style="font-family: 'Iceland', sans-serif"
		>
			Profile & Settings
		</h1>

		<!-- Avatar Section -->
		<BorderBitsCard class="mb-6">
			<form
				method="POST"
				action="?/uploadAvatar"
				enctype="multipart/form-data"
				use:enhance={() => {
					uploading = true;
					return ({ update }) => {
						uploading = false;
						return update();
					};
				}}
			>
				<div class="flex flex-col items-center gap-5 py-4 sm:flex-row sm:items-start">
					<!-- Avatar preview -->
					<div class="relative shrink-0">
						{#if avatarPreview || user?.avatar_url}
							<img
								src={avatarPreview ?? user?.avatar_url}
								alt="avatar"
								class="h-24 w-24 rounded-full object-cover ring-2 ring-red-500"
							/>
						{:else}
							<div
								class="flex h-24 w-24 items-center justify-center rounded-full bg-red-500/20 text-4xl font-bold text-red-400 ring-2 ring-red-500"
								style="font-family: 'Iceland', sans-serif"
							>
								{(user?.display_name ?? user?.username ?? '?')[0].toUpperCase()}
							</div>
						{/if}
						<label
							for="avatar-input"
							class="absolute -right-1 -bottom-1 flex h-7 w-7 cursor-pointer items-center justify-center rounded-full bg-red-500 text-xs text-white shadow-lg transition-transform hover:scale-110"
							title="Change avatar"
						>
							✎
						</label>
						<input
							id="avatar-input"
							name="avatar"
							type="file"
							accept="image/jpeg,image/png,image/webp,image/gif"
							class="hidden"
							onchange={handleFileSelect}
						/>
					</div>

					<!-- User info + upload action -->
					<div class="flex-1 text-center sm:text-left">
						<p class="text-lg font-semibold text-white" style="font-family: 'Iceland', sans-serif">
							{user?.display_name}
						</p>
						<p class="text-sm text-white/40">@{user?.username}</p>
						<p class="mt-3 text-xs text-white/30">JPG, PNG, WEBP or GIF · max 5 MB</p>

						{#if avatarFile}
							<p class="mt-1 max-w-xs truncate text-xs text-white/50">{avatarFile.name}</p>
							<button
								type="submit"
								disabled={uploading}
								class="mt-3 rounded border border-red-500/50 bg-red-500/15 px-4 py-1.5 text-sm text-red-400 transition-colors hover:bg-red-500/25 disabled:opacity-50"
								style="font-family: 'Iceland', sans-serif"
							>
								{uploading ? 'Uploading…' : 'Upload Avatar'}
							</button>
						{/if}
					</div>
				</div>

				{#if form?.avatarError}
					<p class="mt-2 text-xs text-red-400">{form.avatarError}</p>
				{/if}
				{#if form?.avatarOk}
					<p class="mt-2 text-xs text-green-400">Avatar updated successfully! Refresh to see it.</p>
				{/if}
			</form>
		</BorderBitsCard>

		<!-- Display Name Section -->
		<BorderBitsCard class="mb-6">
			<form
				method="POST"
				action="?/updateDisplayName"
				use:enhance={() =>
					({ update }) =>
						update({ reset: false })}
				class="py-2"
			>
				<p
					class="mb-4 text-sm tracking-widest text-white/60 uppercase"
					style="font-family: 'Iceland', sans-serif"
				>
					Display Name
				</p>

				<div class="flex gap-2">
					<input
						type="text"
						name="display_name"
						bind:value={displayName}
						minlength="2"
						maxlength="50"
						required
						class="flex-1 rounded border border-white/15 bg-white/5 px-3 py-2 text-sm text-white placeholder-white/25 transition-colors outline-none focus:border-red-500/60 focus:ring-1 focus:ring-red-500/30"
					/>
					<button
						type="submit"
						class="rounded border border-red-500/40 bg-red-500/10 px-4 py-2 text-sm text-red-400 transition-colors hover:bg-red-500/20 hover:text-red-300"
						style="font-family: 'Iceland', sans-serif"
					>
						Save
					</button>
				</div>

				{#if form?.updateError}
					<p class="mt-2 text-xs text-red-400">{form.updateError}</p>
				{/if}
				{#if form?.updateOk}
					<p class="mt-2 text-xs text-green-400">Display name updated!</p>
				{/if}
			</form>
		</BorderBitsCard>

		<!-- Account Info (read-only) -->
		<BorderBitsCard>
			<div class="py-2">
				<p
					class="mb-4 text-sm tracking-widest text-white/60 uppercase"
					style="font-family: 'Iceland', sans-serif"
				>
					Account Info
				</p>
				<div class="space-y-3">
					<div>
						<p class="mb-0.5 text-xs tracking-wide text-white/30 uppercase">Username</p>
						<p class="text-sm text-white/70">@{user?.username}</p>
					</div>
					<div>
						<p class="mb-0.5 text-xs tracking-wide text-white/30 uppercase">User ID</p>
						<p class="font-mono text-sm text-white/40">{user?.id}</p>
					</div>
				</div>
			</div>
		</BorderBitsCard>
	</div>
</div>
