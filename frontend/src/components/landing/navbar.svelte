<script lang="ts">
	import { Hamburger } from 'svelte-hamburgers';
	import { blur } from 'svelte/transition';
	let open = false;
</script>

<div
	class="fixed top-0 left-0 px-8 w-full flex items-center justify-between gap-2 z-50 bg-gradient-to-b from-black to-transparent backdrop-blur-sm"
>
	<!-- Logo -->
	<img src="/svgs/logo.svg" alt="Logo" class="h-20 w-20" />

	<!-- CTA Buttons -->
	<div class="hidden md:flex items-center gap-12 mx-auto uppercase">
		<a href="/">Home</a>
		<button disabled class="btn variant-filled-primary font-bold"
			><svg
				xmlns="http://www.w3.org/2000/svg"
				width="24"
				height="24"
				viewBox="0 0 24 24"
				class="mr-2"
				><path
					fill="currentColor"
					d="M6 22q-.825 0-1.412-.587T4 20V10q0-.825.588-1.412T6 8h1V6q0-2.075 1.463-3.537T12 1t3.538 1.463T17 6v2h1q.825 0 1.413.588T20 10v10q0 .825-.587 1.413T18 22zm6-5q.825 0 1.413-.587T14 15t-.587-1.412T12 13t-1.412.588T10 15t.588 1.413T12 17M9 8h6V6q0-1.25-.875-2.125T12 3t-2.125.875T9 6z"
				/></svg
			> Sign up</button
		>
	</div>

	<!-- Github Icon -->
	<a
		href="https://github.com/adan-ayaz-stan/gamebuddies"
		target="_blank"
		rel="noopener noreferrer"
		class="hidden md:block"
	>
		<button class="btn variant-ghost-surface text-glowRed"
			><img
				src="data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxZW0iIGhlaWdodD0iMWVtIiB2aWV3Qm94PSIwIDAgMjQgMjQiPjxwYXRoIGQ9Ik0yMC4zOCA4LjUzYy4xNi0uNC42OC0xLjk5LS4xNy00LjE0YzAgMC0xLjMxLS4zOS00LjMgMS42MWMtMS4yNS0uMzMtMi41OC0uMzgtMy45MS0uMzhjLTEuMzIgMC0yLjY2LjA1LTMuOTEuMzhjLTIuOTktMi4wMy00LjMtMS42MS00LjMtMS42MWMtLjg1IDIuMTUtLjMzIDMuNzQtLjE2IDQuMTRDMi42MSA5LjYyIDIgMTEgMiAxMi43MmMwIDYuNDQgNC4xNiA3Ljg5IDEwIDcuODljNS43OSAwIDEwLTEuNDUgMTAtNy44OWMwLTEuNzItLjYxLTMuMS0xLjYyLTQuMTlNMTIgMTkuMzhjLTQuMTIgMC03LjQ3LS4xOS03LjQ3LTQuMTljMC0uOTUuNDctMS44NSAxLjI3LTIuNThjMS4zNC0xLjIzIDMuNjMtLjU4IDYuMi0uNThjMi41OSAwIDQuODUtLjY1IDYuMi41OGMuOC43MyAxLjMgMS42MiAxLjMgMi41OGMwIDMuOTktMy4zNyA0LjE5LTcuNSA0LjE5bS0zLjE0LTYuMjZjLS44MiAwLTEuNSAxLTEuNSAyLjIyYzAgMS4yMy42OCAyLjI0IDEuNSAyLjI0Yy44MyAwIDEuNS0xIDEuNS0yLjI0YzAtMS4yMy0uNjctMi4yMi0xLjUtMi4yMm02LjI4IDBjLS44MyAwLTEuNS45OS0xLjUgMi4yMmMwIDEuMjQuNjcgMi4yNCAxLjUgMi4yNGMuODIgMCAxLjUtMSAxLjUtMi4yNGMwLTEuMjMtLjY0LTIuMjItMS41LTIuMjJ6IiBmaWxsPSIjZmYwMDAwIi8+PC9zdmc+"
				alt="github icon"
				class="h-6 w-6 mr-2"
			/> Github</button
		></a
	>

	<div class="md:hidden text-white">
		<Hamburger bind:open --color="white" --border-radius={2} />
	</div>
</div>
<!-- Menu -->
{#if open}
	<div
		class="min-h-screen bg-black/50 backdrop-blur-md fixed top-0 left-0 w-screen h-screen z-20 py-20 px-8 flex flex-col justify-center transition-all duration-300 gap-8"
		transition:blur={{ amount: 10 }}
	>
		<a href="/">Home</a>
		<a href="/">Features</a>
		<a href="/">Team</a>
		<a href="/" class="btn variant-filled-primary font-light mt-16"
			>Sign up <img
				src="https://api.iconify.design/material-symbols:lock-clock-rounded.svg?color=%23ffffff"
				alt="lock clock icon"
				class="h-6 w-6 ml-2"
			/></a
		>
	</div>
{/if}
