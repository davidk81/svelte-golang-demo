<script>
  export let healthNote
</script>

<style>
  div {
    background-color: rgb(251, 252, 206);
    padding: 5px;
    margin: 5px;
  }
  table {
    width: 100%;
  }
  td {
    white-space: pre-wrap;
    vertical-align: text-top;
  }
</style>

<div>
  <table>
    <tr>
      <td>Health Note:</td>
      <td>{healthNote.note}</td>
    </tr>
    <tr>
      <td>Nurse ID:</td>
      <td>{healthNote.userid}</td>
    </tr>
    <tr>
      <td>Date:</td>
      <td>{new Date(healthNote.created)}</td>
    </tr>
  </table>
</div>
