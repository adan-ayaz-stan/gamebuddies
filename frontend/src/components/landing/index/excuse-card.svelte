<script lang="ts">
	type ExcuseCardType = 'yellow' | 'blue' | 'gray';

	export let content: string;
	export let type: ExcuseCardType = 'gray';
</script>

<div class="relative">
	<img src="/graphics/ex-card-{type}.png" alt="Excuse Card" class="h-96 w-96 object-contain" />

	<div class="absolute top-0 left-0 w-full h-full uppercase p-10">
		<h2 class="h2">{content}</h2>
	</div>
</div>
