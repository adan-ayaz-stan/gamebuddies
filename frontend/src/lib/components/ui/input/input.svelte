<script lang="ts">
	import { cn } from '$lib/utils';

	let {
		class: className,
		value = $bindable(),
		placeholder,
		type = 'text',
		disabled = false,
		...restProps
	}: {
		class?: string;
		value?: string;
		placeholder?: string;
		type?: string;
		disabled?: boolean;
		[key: string]: any;
	} = $props();
</script>

<input
	bind:value
	{type}
	{placeholder}
	{disabled}
	class={cn(
		'flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 focus-visible:outline-none disabled:cursor-not-allowed disabled:opacity-50',
		className
	)}
	{...restProps}
/>
