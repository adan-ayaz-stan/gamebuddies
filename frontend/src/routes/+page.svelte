<script lang="ts">
	import { goto } from '$app/navigation';
	import { Button } from '$lib/components/ui/button';
	import { onMount } from 'svelte';
	import { cubicOut } from 'svelte/easing';
	import { Tween } from 'svelte/motion';
	import type { PageData } from './$types';

	let { data }: { data: PageData } = $props();

	// Animation stores
	const pageOpacity = new Tween(0, { duration: 800, easing: cubicOut });
	const titleOpacity = new Tween(0, { duration: 600, easing: cubicOut, delay: 200 });
	const descriptionOpacity = new Tween(0, { duration: 600, easing: cubicOut, delay: 400 });
	const buttonOpacity = new Tween(0, { duration: 600, easing: cubicOut, delay: 600 });

	// Entry animation on mount
	onMount(async () => {
		pageOpacity.set(1);
		titleOpacity.set(1);
		descriptionOpacity.set(1);
		buttonOpacity.set(1);
	});

	// Exit animation and navigation
	async function handleEnterSystem() {
		await buttonOpacity.set(0);
		await descriptionOpacity.set(0);
		await titleOpacity.set(0);

		await new Promise((resolve) => setTimeout(resolve, 300));

		// Authenticated users go straight to the dashboard
		goto(data.user ? '/dashboard' : '/login');
	}
</script>

<div
	class="relative flex min-h-screen items-center justify-center overflow-x-hidden bg-black"
	style="opacity: {pageOpacity.current}"
>
	<div
		class="absolute top-0 right-0 h-60 w-60 translate-x-1/2 -translate-y-1/2 bg-red-500 blur-[10em]"
	></div>
	<div
		class="absolute bottom-1/2 left-0 h-60 w-60 -translate-x-1/2 translate-y-1/2 bg-red-500 blur-[10em]"
	></div>

	<!--  -->
	<div class="mx-auto flex max-w-7xl items-center justify-center gap-12">
		<h1 class="text-[50vh]" style="opacity: {titleOpacity.current}">Hi!</h1>
		<div class="flex max-w-xl flex-col gap-4 text-white">
			<div style="opacity: {descriptionOpacity.current}">
				<p>
					Welcome to GameBuddies. The sci-fi hub to meet your daily gaming partners and have a
					banging time!
				</p>
				<p>
					You can join game rooms, chat globally, make friends and more. Join us right now to see
					what're all about!
				</p>
			</div>

			<div style="opacity: {buttonOpacity.current}">
				<Button variant="outline" class="w-fit" onclick={handleEnterSystem}>Enter the System</Button
				>
			</div>
		</div>
	</div>
</div>
