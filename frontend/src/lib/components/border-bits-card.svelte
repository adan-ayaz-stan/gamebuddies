<script>
	let { class: className = '', children } = $props();
</script>

<div class="relative {className}">
	<!-- Corner caps -->
	<div class="absolute top-0 left-0 h-4 w-4 border-t-2 border-l-2 border-red-500"></div>
	<div class="absolute top-0 right-0 h-4 w-4 border-t-2 border-r-2 border-red-500"></div>
	<div class="absolute bottom-0 left-0 h-4 w-4 border-b-2 border-l-2 border-red-500"></div>
	<div class="absolute right-0 bottom-0 h-4 w-4 border-r-2 border-b-2 border-red-500"></div>

	<!-- Main border -->
	<div class="h-full border border-white/13 p-4">
		{@render children?.()}
	</div>
</div>
