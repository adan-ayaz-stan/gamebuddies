<script lang="ts">
	type CardItem = {
		title: string;
		description: string;
		icon_url: string;
	};

	const data: CardItem[] = [
		{
			title: 'Voice Rooms',
			description: 'Public and private voice rooms to fit your needs.',
			icon_url:
				'https://api.iconify.design/streamline:interface-id-voice-1-identification-secure-id-soundwave-sound-voice-security.svg?color=%23f92424'
		},
		{
			title: 'Matchmaking',
			description: 'The right people at your right time, in your right place.',
			icon_url: 'https://api.iconify.design/iconoir:archery-match.svg?color=%23f92424'
		},
		{
			title: 'Karma Leaderboards',
			description: 'Making sure you fit in with your type of crowd.',
			icon_url: 'https://api.iconify.design/mdi:trophy.svg?color=%23f92424'
		},
		{
			title: 'Integrations',
			description:
				'Originating as a lightweight client, allows addition of integrations on top of it.',
			icon_url: 'https://api.iconify.design/carbon:continuous-integration.svg?color=%23f92424'
		},
		{
			title: 'Profile Skins',
			description: 'Build your own profile skin or distribute to community.',
			icon_url: 'https://api.iconify.design/ant-design:skin-filled.svg?color=%23f92424'
		},
		{
			title: 'Lightweight Client',
			description: 'Aims to be as lightweight as no weight.',
			icon_url:
				'https://api.iconify.design/material-symbols:laptop-windows-sharp.svg?color=%23f92424'
		}
	];
</script>

<div class="min-h-screen grid relative bg-gradient-to-b from-black to-darkRed">
	<!-- Background Image -->
	<img src="/svgs/logo.svg" alt="Logo" class="h-full w-full absolute top-0 left-0 object-contain" />

	<!-- Content -->
	<div
		class="w-full h-full flex items-center justify-center z-10 relative backdrop-blur-md bg-gradient-to-b from-black via-darkRed/70 to-black"
	>
		<!-- Main Card -->
		<div class="grid lg:grid-cols-2 xl:grid-cols-3 auto-rows-fr gap-8 p-8 px-16">
			{#each data as item}
				<div
					class="border-2 rounded-xl border-lightRed h-full text-lightRed p-6 text-center bg-gradient-to-tr flex flex-row items-center gap-4 from-black via-darkRed to-black"
				>
					<img src={item.icon_url} alt="icon" class="h-12 w-12" />
					<div>
						<h3 class="h3 tracking-normal">{item.title}</h3>
						<p class="text-white">{item.description}</p>
					</div>
				</div>
			{/each}
		</div>
	</div>
</div>

<!-- Features: Matchmaking, Rooms, Leaderboards,  -->
