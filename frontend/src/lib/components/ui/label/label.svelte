<script lang="ts">
	import { cn } from '$lib/utils';

	let {
		class: className,
		for: htmlFor,
		children,
		...restProps
	}: {
		class?: string;
		for?: string;
		children?: any;
		[key: string]: any;
	} = $props();
</script>

<label
	for={htmlFor}
	class={cn(
		'text-sm leading-none font-medium peer-disabled:cursor-not-allowed peer-disabled:opacity-70',
		className
	)}
	{...restProps}
>
	{@render children?.()}
</label>
