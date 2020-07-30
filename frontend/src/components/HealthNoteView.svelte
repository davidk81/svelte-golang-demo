<script>
  export let healthNote
</script>

<style>
  div {
    background-color: rgb(251, 252, 206);
    padding: 5px;
    margin: 5px;
  }
</style>

<div>
{healthNote.note}<br/>
by : {healthNote.nurseId}<br/>
date: {new Date(healthNote.timestamp)}<br/>
</div>
