<script>
  import { stores } from '@sapper/app'
  
  const { session } = stores()
  export let patient = { name: "Patrick", patientId: "patient1"}
  export let healthNote
  let note = ""

  async function onSave() {
    var newNote = {
      user_id: $session.profile.username,
      patient_id: patient.patientId,
      note: note
    }
    console.log(newNote)
    await fetch('http://localhost:8000/api/v1/patient/note', {
      method: 'POST',
      mode: 'cors',
      credentials: 'include',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(newNote)
    })
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
  patient : {patient.name} ({patient.patientId})<br/>
  <textarea bind:value={note}></textarea><br/>
  by : {$session.profile.name} ({$session.profile.username})<br/>
  <button on:click={onSave}>save</button>
  </div>
{/if}
