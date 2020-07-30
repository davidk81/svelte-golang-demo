<script>
  import { goto, stores } from '@sapper/app'
  
  const { session } = stores()
  let username = null
  let password = 'password'

  async function login () {
    const response = await fetch('http://localhost:8000/session', {
      method: 'POST',
      mode: 'cors',
      credentials: 'include',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        username,
        password
      })
    })
    if (response.ok) {
      const profile = await response.json();
      session.update(() => {
                    return {
                        authenticated: !!profile,
                        profile
                    }
                });
    } else {
      session.update(() => {
                    return {
                        authenticated: false,
                        profile : null
                    }
                });
      console.log(response);
    }
    goto('patients')
  }
</script>

<svelte:head>
	<title>Patient Management System</title>
</svelte:head>

{#if $session && $session.authenticated}
  <p>Welcome, {$session.profile.name}</p>
{:else}
  <form>
  <p>Nurse Login</p>
  <input type="txt" bind:value={username} />
  <input type="password" bind:value={password} />
  <button type="button" disabled={!username} on:click={login}>login</button>
  </form>
{/if}
