<script>
  export let healthNote = { note: "" }

  function onSave() {
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

<div>
Add new note:<br/>
<textarea bind:value={healthNote.note}></textarea><br/>
<button on:click={onSave}>save</button>
</div>
