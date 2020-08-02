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
  td {
    margin: 5px;
    padding: 5px;
  }
  .active {
		background-color: #ff3e00;
		color: white;
	}
</style>

<div on:click={onClick} let:hovering={active}>
  <table>
    <tr>
      <td>Name:</td>
      <td>{patient.name}</td>
    </tr>
    <tr>
      <td>Location:</td>
      <td>{patient.location}</td>
    </tr>
  </table>
</div>
