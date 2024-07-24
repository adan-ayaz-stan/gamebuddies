<div
	class="min-h-screen bg-gradient-to-b from-darkRed to-black from-0% to-50% flex flex-col items-center justify-center gap-4"
>
	<!-- Glow Logo -->
	<div
		class="w-32 h-32 shadow-lg shadow-glowRed rounded-xl border-2 border-lightRed p-4 bg-gradient-to-tr from-black to-darkRed"
	>
		<img src="/svgs/logo.svg" alt="logo" class="w-full h-full" />
	</div>
	<p class="text-gray-400 opacity-50 mt-4">GAMEBUDDIES</p>

	<!-- Heading -->
	<div class="text-center mt-7">
		<h1 class="h1 leading-0 tracking-normal">JOIN THE WAITLIST FOR</h1>
		<h1 class="h1 text-lightRed tracking-normal">THE SPOT IN THE 1000.</h1>
	</div>

	<!-- Sub heading -->
	<p class="text-xl md:text-2xl mt-4 uppercase">
		Get your spot in <a href="/" class="text-lightRed underline">the beta release</a>.
	</p>

	<!-- Form goes here -->
	<form method="POST" action="?/waitlist_register">
		<div class="input-group input-group-divider grid-cols-[auto_1fr_auto] my-2">
			<div class="input-group-shim">
				<img
					src="https://api.iconify.design/material-symbols:person-outline.svg?color=%23999999"
					alt="person icon"
					class="h-6 w-6"
				/>
			</div>
			<input name="name" type="text" placeholder="Spitfire" autocomplete="off" required />
		</div>
		<div class="input-group input-group-divider grid-cols-[auto_1fr_auto] my-2">
			<div class="input-group-shim">
				<img
					src="https://api.iconify.design/material-symbols:mail-rounded.svg?color=%23999999"
					alt="email icon"
					class="h-6 w-6"
				/>
			</div>
			<input
				type="email"
				name="email"
				autocomplete="off"
				placeholder="gamebuddies@luceforge.com"
				required
			/>
		</div>

		<button type="submit" class="btn variant-filled-primary w-full mt-4">Apply</button>
	</form>
</div>
