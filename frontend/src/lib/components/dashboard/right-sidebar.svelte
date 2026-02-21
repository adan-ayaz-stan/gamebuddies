<script>
	import { Settings } from '@lucide/svelte';
	import { Button } from '../ui/button';
	import HighlightCard from './highlight-card.svelte';
</script>

<div class="flex h-full flex-col gap-4">
	<HighlightCard header="Actively Playing" subheader={'Valorant'}>Content goes here</HighlightCard>
	<div class="mt-auto w-full">
		<Button variant="outline" class="items-left w-full">Settings <Settings size={24} /></Button>
	</div>
</div>
