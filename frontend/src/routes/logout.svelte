<script>
  import { onMount } from 'svelte';
  import { goto, stores } from '@sapper/app'

  const { session } = stores()

	onMount(async () => {
    session.set({ authenticated: false, profile: null })
    await fetch('http://localhost:8000/api/v1/session', {
      method: 'DELETE',
      mode: 'cors',
      credentials: 'include',
      headers: {
        'Content-Type': 'application/json'
      }
    })
    goto('/')
	});
</script>
