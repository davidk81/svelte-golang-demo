<script>
  import { onMount } from 'svelte';
  import HealthNoteView from '../components/HealthNoteView.svelte';
  import HealthNoteEdit from '../components/HealthNoteEdit.svelte';

  let patient
  let healthNotes;

	onMount(async () => {
    const response = await fetch('mock-patient.json', {
      method: 'GET',
      mode: 'cors',
      credentials: 'include',
      headers: {
        'Content-Type': 'application/json'
      }
    })
    patient = await response.json();
	});

  onMount(async () => {
    const response = await fetch('mock-health-notes.json', {
      method: 'GET',
      mode: 'cors',
      credentials: 'include',
      headers: {
        'Content-Type': 'application/json'
      }
    })
    healthNotes = await response.json();
	});
</script>

<h1>Patient</h1>
{#if patient}
  <div>
  {patient.name} ({patient.patientId})
  </div>
  {#if healthNotes}
    {#each healthNotes.records as healthNote}
      <HealthNoteView healthNote={healthNote}/>
    {/each}
  {/if}
  <HealthNoteEdit/>
{/if}
