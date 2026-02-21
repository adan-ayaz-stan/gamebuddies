<script>
	import { page } from '$app/state';
	import HighlightCard from './highlight-card.svelte';

	const navItems = [
		{
			header: 'Home',
			copy: 'Your personal command center and gaming hub.',
			href: '/dashboard'
		},
		{
			header: 'Games',
			copy: 'Browse, discover and launch your favorite titles.',
			href: '/dashboard/games'
		},
		{
			header: 'Matchmaking',
			copy: 'Find the perfect teammates for your next adventure.',
			href: '/dashboard/matchmaking'
		},
		{
			header: 'Friends',
			copy: 'Connect with your gaming crew and make new allies.',
			href: '/dashboard/friends'
		}
	];
</script>

<div class="-mt-1/2 relative flex scale-75 items-stretch gap-4 lg:gap-8">
	{#each navItems as item}
		<a href={item.href} class="block">
			<HighlightCard
				header={item.header}
				cardClassname={page.url.pathname === item.href ? '' : 'border-white/13'}
				headerClassname={page.url.pathname === item.href
					? ''
					: 'bg-white/13 group-hover:bg-primary'}
			>
				<p>{item.copy}</p>
			</HighlightCard>
		</a>
	{/each}
</div>
