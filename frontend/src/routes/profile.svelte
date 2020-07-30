<script>
  import { onMount } from 'svelte';

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
{#if profile}
<h2>Name</h2>
<ul>{profile.name}</ul>

<h2>Login ID</h2>
<ul>{profile.username}</ul>

<h2>Roles</h2>
<ul>
  {#each profile.roles as role}
  <li>{role}</li>
  {/each}
</ul>
{/if}
