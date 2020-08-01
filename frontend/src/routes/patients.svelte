<script>
  import { onMount } from 'svelte';
  import Patient from '../components/Patient.svelte';

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
  {#if patients}
    {#each patients as patient}
      <Patient patient={patient}/>
    {/each}
  {/if}
{#if error}
  {error}
{/if}
