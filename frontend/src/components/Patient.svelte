<script>
  export let patient

  function onClick() {
		window.location.href= 'patient?patientId=' + patient.patientId
	}
</script>

<style>
  div {
    background-color: aquamarine;
    padding: 5px;
    margin: 5px;
    cursor: pointer;
  }
</style>

<div on:click={onClick}>
{patient.name} ({patient.patientId})
</div>
