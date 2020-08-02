export function DoLogin(username, password) {
  return fetch('http://localhost:8000/api/v1/session', {
    method: 'POST',
    mode: 'cors',
    credentials: 'include',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({
      username,
      password
    })
  })
}

export function DoRegister(user) {
  return fetch('http://localhost:8000/api/v1/register', {
    method: 'POST',
    mode: 'cors',
    credentials: 'include',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(user)
  })
}

