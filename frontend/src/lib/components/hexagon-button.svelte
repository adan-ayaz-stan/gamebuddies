<script>
	let { children, ...props } = $props();
</script>

<div
	{...props}
	class="group relative h-[4em] w-[3.464em] cursor-pointer items-center justify-center bg-primary"
	style="clip-path: polygon(
  50% 0%,
  100% 25%,
  100% 75%,
  50% 100%,
  0% 75%,
  0% 25%
);"
>
	<div
		class="absolute top-1/2 left-1/2 flex h-[3.8em] w-[3.264em] -translate-x-1/2 -translate-y-1/2 items-center justify-center bg-background text-primary group-hover:bg-primary group-hover:text-white"
		style="clip-path: polygon(
  50% 0%,
  100% 25%,
  100% 75%,
  50% 100%,
  0% 75%,
  0% 25%
);"
	>
		{@render children()}
	</div>
</div>
