<script>
	import { goto } from '$app/navigation';
	import { Headphones, MicOff, Phone } from '@lucide/svelte';
	import HexagonButton from '../hexagon-button.svelte';

	function endCall() {
		// End the call and redirect to matchmaking
		goto('/dashboard/matchmaking');
	}
</script>

<div class="flex items-center justify-center gap-4 pt-2 pb-4">
	<!--  -->
	<HexagonButton>
		<MicOff size={24} class="text-inherit" />
	</HexagonButton>
	<HexagonButton>
		<Headphones size={24} class="text-inherit" />
	</HexagonButton>
	<HexagonButton onclick={endCall}>
		<Phone size={24} class={'rotate-135 text-inherit'} />
	</HexagonButton>
</div>
