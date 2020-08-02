<script>
  import { goto, stores } from '@sapper/app'
  import Login from '../components/Login.svelte'
  
  const { session } = stores()
  let error

</script>

<style>
  .center {
    margin: auto;
    width: 100%;
    display: flex;
    justify-content: center;
  }
</style>

<svelte:head>
	<title>Patient Management System</title>
</svelte:head>

<div class=center>
{#if $session && !$session.loading}
  {#if $session.authenticated}
    <p>Welcome, {$session.profile.name}</p>
  {:else}
    <Login/>
  {/if}
{/if}
{#if error}
  {error}
{/if}
</div>
