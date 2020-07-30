<script>
  import { goto, stores } from '@sapper/app'

  export let patient

  function onClick() {
		goto('patient?patientId=' + patient.patientId)
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
