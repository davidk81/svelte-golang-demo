<script>
  import { stores } from '@sapper/app'
  
  const { session } = stores()
  export let patient
  export let onUpdated
  let note = ""
  let error
  let loading = false

  async function onSave() {
    loading = true
    if (note.length == 0) {
      error = "health note must not be empty"
      return
    }
    var newNote = {
      userid: $session.profile.username,
      patientid: patient.patientid,
      note: note
    }
    const response = await fetch('http://localhost:8000/api/v1/patient/note', {
      method: 'POST',
      mode: 'cors',
      credentials: 'include',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(newNote)
    })
    if (response.ok) {
      error = null
      note = ""
      if (onUpdated != null) onUpdated()
    }
    else {
      error = await response.text()
    }
    loading = false
  }
</script>

<style>
  div {
    background-color: rgb(241, 241, 240);
    padding: 5px;
    margin: 5px;
  }
	textarea { width: 95%; height: 200px; }
</style>

{#if $session && $session.authenticated}
  <div>
  Add new note:<br/>
  patient : {patient.name} ({patient.patientid})<br/>
  <textarea bind:value={note}></textarea><br/>
  by : {$session.profile.name} ({$session.profile.username})<br/>
  <button disabled={loading} on:click={onSave}>save</button>
  </div>
{/if}
{#if error}
  {error}
{/if}
