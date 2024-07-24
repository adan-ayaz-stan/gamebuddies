<script>
	const row2Data = [
		{
			title: 'Team',
			url: '/team'
		},
		{
			title: 'Ecosystem',
			url: '/ecosystem'
		},
		{
			title: 'Features',
			url: '/features'
		},
		{
			title: 'Contact',
			url: '/contact'
		},
		{
			title: 'Support',
			url: '/support'
		}
	];
</script>

<div class="min-h-screen bg-black flex flex-col justify-end gap-16 pt-16">
	<!-- Qoute -->

	<div class="w-full space-y-4">
		<hr class="w-2/4 h-1 bg-lightRed mx-auto" />
		<h1 class="h2 w-fit mx-auto blockquote text-center text-lightRed">
			Connections in small moments of geniune interactions
		</h1>
	</div>

	<!--  -->
	<div class="w-full border-t-4 grid grid-cols-2 md:grid-cols-7 auto-rows-auto border-lightRed">
		<!-- Row 1 -->
		<div
			class="col-span-3 border-r-4 flex flex-row gap-4 items-center border-b-4 border-lightRed p-6"
		>
			<img src="/svgs/logo.svg" alt="Logo" class="h-20 w-20 inline" />
			<h2 class="h2 inline">Luceforge studio</h2>
		</div>

		<div
			class="hidden col-span-2 border-r-4 md:flex flex-row gap-4 items-center border-b-4 border-lightRed p-6"
		>
			<p class="text-2xl uppercase">About</p>
		</div>
		<div
			class="hidden col-span-2 border-r-4 md:border-r-0 md:flex flex-row gap-4 items-center border-b-4 border-lightRed p-6"
		>
			<p class="text-2xl uppercase">Documents</p>
		</div>

		<!-- Row 2 -->
		<div
			class="h-[25em] col-span-3 border-r-4 flex flex-col gap-4 justify-between border-b-4 border-lightRed p-6"
		>
			<div>
				<p>Get in touch!</p>
				<a href="mailto:luceforge@gmail.com" class="text-lightRed">luceforge@gmail.com</a>
			</div>

			<a href="/">
				<h1 class="h1 text-lightRed tracking-normal">
					Join Discord <img
						src="https://api.iconify.design/ic:baseline-arrow-outward.svg?color=%23FF5A5F"
						alt="outward arrow"
						class="h-16 w-16 inline -ml-2 -top-1 relative"
					/>
				</h1>
			</a>
		</div>

		<div class="col-span-2 border-r-4 flex flex-col gap-4 border-b-4 border-lightRed p-6">
			{#each row2Data as item}
				<a href={item.url} class="">{item.title}</a>
			{/each}
		</div>

		<div
			class="col-span-2 flex flex-col gap-4 border-b-4 border-r-4 md:border-r-0 border-lightRed p-6"
		>
			<a href="/"
				>Pitch Deck <img
					src="https://api.iconify.design/ic:baseline-arrow-outward.svg?color=%23FFFFFF"
					alt="outward arrow"
					class="h-8 w-8 inline relative -top-0.5"
				/></a
			>
			<a href="/"
				>Roadmap <img
					src="https://api.iconify.design/ic:baseline-arrow-outward.svg?color=%23FFFFFF"
					alt="outward arrow"
					class="h-8 w-8 inline relative -top-0.5"
				/></a
			>
			<a href="/">Legal</a>
		</div>

		<!-- Row 3 -->
		<div class="col-span-3 border-r-4 flex flex-row gap-4 items-center border-lightRed p-6">
			<p>© 2024 Luceforge Studios, Gamebuddies. Rights reserved.</p>
		</div>
		<div class="col-span-2 border-r-4 flex flex-row gap-4 items-center border-lightRed p-6">
			<p>TERMS OF SERVICE</p>
		</div>
		<div
			class="col-span-2 flex flex-row border-r-4 md:border-r-0 gap-4 items-center border-lightRed p-6"
		>
			<p>PRIVACY POLICY</p>
		</div>
	</div>
</div>
