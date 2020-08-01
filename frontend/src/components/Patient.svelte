<script>
  import { goto, stores } from '@sapper/app'

  export let patient

  function onClick() {
		goto('patient?patientid=' + patient.patientid)
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
<span>{patient.patientid},</span>
<span>{patient.name},</span>
<span>{patient.location},</span>
<span>{patient.created}</span>
</div>
