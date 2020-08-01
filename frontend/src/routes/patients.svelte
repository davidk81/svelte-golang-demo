<script>
  import { onMount } from 'svelte';
  import { stores } from '@sapper/app'
  import Patient from '../components/Patient.svelte';

  const { session } = stores()
  let patients
  let error

	onMount(async () => {
    const response = await fetch('http://localhost:8000/api/v1/patients', {
      method: 'GET',
      mode: 'cors',
      credentials: 'include',
      headers: {
        'Content-Type': 'application/json'
      }
    })
    if (response.ok) {
      error = null
      patients = await response.json();
    }
    else {
      error = await response.text()
    }
	});
</script>

<h1>Patients</h1>
{#if $session && $session.authenticated}
  {#if patients}
    {#each patients as patient}
      <Patient patient={patient}/>
    {/each}
  {/if}
{/if}
{#if error}
  {error}
{/if}
