<script>
  import { goto, stores } from '@sapper/app'
  import { DoLogin, DoRegister } from '../session.js'
  
  const { session } = stores()
  let name = ''
  let username = 'nurse1'
  let password = 'password'
  let verify = ''
  let register = false
  let error

  async function onClick () {
    if (register) {
      btnRegister();
    } else {
      btnLogin();
    }
  }

  async function btnLogin() {
    const response = await DoLogin(username, password);
    processResponse(response);
  }

  async function btnRegister() {
    if (!verifyPassword()) return;
    const response = await DoRegister({
      name,
      username,
      password,
      roles: ["nurse"]
    });
    processResponse(response);
  }

  async function processResponse(response) {
    if (response.ok) {
      const profile = await response.json();
      session.update(() => {
                    return {
                        authenticated: !!profile,
                        profile
                    }
                });
      if (!profile)
        error = 'login error'
      if (profile.roles.includes('nurse'))
        goto('patients')
    } else {
      session.update(() => {
                    return {
                        authenticated: false,
                        profile : null
                    }
                });
      if (response.text)
        error = await response.text()
      else
        error = "server error"
    }
  }

  function verifyPassword() {
    if (name.length == 0) {
      error = "name cannot be empty"
      return false
    }
    if (password.length < 8) {
      error = "password must be atleast 8 characters"
      return false
    }
    if (verify != password) {
      error = "passwords don't match"
      return false
    }
    error = null
    return true
  }
</script>

<style>
  .center {
    margin: auto;
    width: 100%;
    display: flex;
    justify-content: center;
  }

  .table tr td {
    text-align: right;
  }  
</style>

<div class=center>
<form>
{#if $session && !$session.loading}
  <p>Nurse Login</p>
  <table>
    <tbody>
      <tr>
        <td>Register new user ?</td>
        <td><input type=checkbox bind:checked={register}></td>
      </tr>
      {#if register}
      <tr>
        <td>Full Name:</td>
        <td><input bind:value={name}/></td>
      </tr>
      {/if}
      <tr>
        <td>Login ID:</td>
        <td><input bind:value={username}/></td>
      </tr>
      <tr>
        <td>Password:</td>
        <td><input type="password" bind:value={password} autocomplete="off" readonly onfocus="this.removeAttribute('readonly');" style="text-security:disc; -webkit-text-security:disc;"/></td>
      </tr>
      {#if register}
      <tr>
        <td>Verify Password:</td>
        <td><input type="password" bind:value={verify} on:change={verifyPassword} autocomplete="off" readonly onfocus="this.removeAttribute('readonly');" style="text-security:disc; -webkit-text-security:disc;"/></td>
      </tr>
      {/if}
      {#if error}
      <tr>
        <td></td>
        <td class="error">{error}</td>
      </tr>
      {/if}
      <tr>
        <td></td>
        <td><button type="button" disabled={!username} on:click={onClick}>{register? "register" : "login"} </button></td>
      </tr>
    </tbody>
  </table>  
{/if}
</form>
</div>
