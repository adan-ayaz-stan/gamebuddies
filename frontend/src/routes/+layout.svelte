<script lang="ts">
	import favicon from '$lib/assets/favicon.svg';
	import { authStore } from '$lib/stores/auth.svelte';
	import '../app.css';

	let { data, children } = $props();

	// Keep the client-side store in sync with data from server load functions
	$effect(() => {
		authStore.setUser(data.user ?? null);
	});
</script>

<svelte:head>
	<link rel="icon" href={favicon} />
</svelte:head>

<div class="min-h-screen bg-black text-white">
	{@render children?.()}
</div>
