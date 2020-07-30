<script>
  	import { stores } from '@sapper/app'
  
  	const { session } = stores()
	export let segment;
</script>

<style>
	nav {
		border-bottom: 1px solid rgba(255,62,0,0.1);
		font-weight: 300;
		padding: 0 1em;
    display: flex;
    justify-content: space-between;
	}

	ul {
		margin: 0;
		padding: 0;
	}

	/* clearfix */
	ul::after {
		content: '';
		display: block;
		clear: both;
	}

	li {
		display: block;
		float: left;
	}

	[aria-current] {
		position: relative;
		display: inline-block;
	}

	[aria-current]::after {
		position: absolute;
		content: '';
		width: calc(100% - 1em);
		height: 2px;
		background-color: rgb(255,62,0);
		display: block;
		bottom: -1px;
	}

	a {
		text-decoration: none;
		padding: 1em 0.5em;
		display: block;
	}
</style>

<nav>
	<ul>
		<li><a aria-current="{segment === undefined ? 'page' : undefined}" href=".">home</a></li>
		{#if $session && $session.authenticated}
			{#if $session.profile.roles.includes('nurse')}
				<li><a aria-current='{segment === "patients" ? "page" : undefined}' href='patients'>patients</a></li>
				{#if segment === "patient"}
					<li><a aria-current='{segment === "patient" ? "page" : undefined}' href='patient'>patient</a></li>
				{/if}
			{/if}
		{/if}
  </ul>
  <ul>
	{#if $session && $session.authenticated}
		{#if $session.profile.roles.includes('admin')}
			<li><a aria-current='{segment === "admin" ? "page" : undefined}' href='admin'>admin</a></li>
		{/if}
		<li><a aria-current='{segment === "profile" ? "page" : undefined}' href='profile'>profile</a></li>
		<li><a href='logout'>log out</a></li>
    {/if}
</ul>
</nav>
