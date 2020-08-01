<script>
  import { onMount } from 'svelte';
  import { stores } from '@sapper/app'

  const { session } = stores()
  let profile
  
  onMount(async () => {
    const response = await fetch('mock-profile.json', {
      method: 'GET',
      mode: 'cors',
      credentials: 'include',
      headers: {
        'Content-Type': 'application/json'
      }
    })
    profile = await response.json();
	});
</script>

<h1>Profile</h1>
{#if $session && $session.authenticated}
  {#if profile}
  <h4>Name</h4>
  <ul>{profile.name}</ul>

  <h4>Login ID</h4>
  <ul>{profile.username}</ul>

  <h4>Roles</h4>
  <ul>
    {#each profile.roles as role}
    <li>{role}</li>
    {/each}
  </ul>
  {/if}
{/if}
