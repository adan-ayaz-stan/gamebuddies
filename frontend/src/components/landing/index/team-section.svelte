<script>
	import MemberCard from './_team-section-components/member-card.svelte';
</script>

<div class="min-h-screen bg-black py-16 px-8 flex flex-col items-center justify-center gap-16">
	<h1 class="h1 text-lightRed">Meet the team</h1>
	<!--  -->
	<div
		class="flex flex-col gap-52 md:gap-0 md:flex-row items-center justify-between w-full max-w-7xl"
	>
		<MemberCard
			type="blue"
			role="First Hand"
			frontsideUrl="/graphics/openrole-card-frontside.png"
		/>

		<MemberCard
			type="gold"
			role="Founder"
			frontsideUrl="/graphics/founder-card-frontside.png"
			githubUrl="https://github.com/adan-ayaz-stan"
			linkedinUrl="https://www.linkedin.com/in/adan-ayaz/"
			telegramUrl="https://t.me/AdanAyaz1"
		/>

		<MemberCard
			type="blue"
			role="Second Hand"
			frontsideUrl="/graphics/openrole-card-r2-frontside.png"
		/>
	</div>
</div>
