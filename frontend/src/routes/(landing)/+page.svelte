<script>
	import Footer from '../../components/landing/footer.svelte';
	import FifthSection from '../../components/landing/index/fifth-section.svelte';
	import FirstSection from '../../components/landing/index/first-section.svelte';
	import FourthSection from '../../components/landing/index/fourth-section.svelte';
	import SecondSection from '../../components/landing/index/second-section.svelte';
	import TeamSection from '../../components/landing/index/team-section.svelte';
	import ThirdSection from '../../components/landing/index/third-section.svelte';
	import WaitlistRegisterSection from '../../components/landing/index/waitlist-register-section.svelte';
	import { MetaTags } from 'svelte-meta-tags';
</script>

<MetaTags
	title="Gamebuddies | By Luceforge Studio"
	titleTemplate="%s - GB Spitfire"
	description="Connect with other gamebuddies, team up on the spot and start dealing with business! Don't wait for anyone. Matchmake, Converse & Play!"
	canonical="https://gamebuddies.vercel.app"
	openGraph={{
		url: 'https://gamebuddies.vercel.app',
		title: 'Gamebuddies | By Luceforge Studio',
		description:
			"Connect with other gamebuddies, team up on the spot and start dealing with business! Don't wait for anyone. Matchmake, Converse & Play!",
		images: [
			{
				url: 'https://gamebuddies.vercel.app/graphics/meta.png',
				width: 800,
				height: 600,
				alt: 'Gamebuddies | By Luceforge Studio'
			}
		],
		siteName: 'Your Website'
	}}
	twitter={{
		handle: '@AdanAyaz1',
		site: '@adanayaz',
		cardType: 'summary_large_image',
		title: 'Gamebuddies | By Luceforge Studio',
		description:
			"Connect with other gamebuddies, team up on the spot and start dealing with business! Don't wait for anyone. Matchmake, Converse & Play!",
		image: 'https://gamebuddies.vercel.app/graphics/meta.png',
		imageAlt: 'Gamebuddies | By Luceforge Studio'
	}}
/>
<div class="min-h-screen">
	<FirstSection />
	<SecondSection />
	<ThirdSection />
	<FourthSection />
	<FifthSection />
	<WaitlistRegisterSection />
	<TeamSection />
	<Footer />
</div>
