<script>
  import { onMount } from 'svelte';
  import Patient from '../components/Patient.svelte';

  let patients

	onMount(async () => {
    const response = await fetch('mock-patients.json', {
      method: 'GET',
      mode: 'cors',
      credentials: 'include',
      headers: {
        'Content-Type': 'application/json'
      }
    })
    patients = await response.json();
	});
</script>

<h1>Patients</h1>
{#if patients}
  {#each patients as patient}
    <Patient patient={patient}/>
  {/each}
{/if}
