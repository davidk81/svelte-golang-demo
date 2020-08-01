<script>
  import { onMount } from 'svelte';
  import { stores } from '@sapper/app'
  import HealthNoteView from '../components/HealthNoteView.svelte';
  import HealthNoteEdit from '../components/HealthNoteEdit.svelte';
  
  const { session } = stores()
  let patient
  let healthNotes;
  let error

  let tokens = window.location.search.split("patientid=")
  let patientid = tokens[tokens.length-1]
  let newNote = { note: "",  patientid}

	onMount(async () => {
    const response = await fetch('http://localhost:8000/api/v1/patient?patientid=' + patientid, {
      method: 'GET',
      mode: 'cors',
      credentials: 'include',
      headers: {
        'Content-Type': 'application/json'
      }
    })
    if (response.ok) {
      error = null
      patient = await response.json();
    }
    else {
      error = await response.text()
    }
	});    

  var refreshNotes = async () => {
    const response = await fetch('http://localhost:8000/api/v1/patient/notes?patientid=' + patientid, {
      method: 'GET',
      mode: 'cors',
      credentials: 'include',
      headers: {
        'Content-Type': 'application/json'
      }
    })
    if (response.ok) {
      error = null
      healthNotes = await response.json();
    }
    else {
      error = await response.text()
    }
  }

	onMount(refreshNotes);  
</script>

<h1>Patient</h1>
{#if $session && $session.authenticated}
  {#if patient}
    <div>
    {patient.name} ({patient.patientid})
    </div>
    {#if healthNotes}
      {#each healthNotes as healthNote}
        <HealthNoteView healthNote={healthNote}/>
      {/each}
    {/if}
    {#if patient}
      <HealthNoteEdit onUpdated={refreshNotes} healthNote={newNote} patient={patient}/>
    {/if}
  {/if}
{/if}
{#if error}
  {error}
{/if}
