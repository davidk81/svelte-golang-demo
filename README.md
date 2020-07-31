# svelte-golang-demo

A demo app using Svelte as frontend and GoLang as backend. This app provides an web-based GUI for nurses to manage patient's health notes. 

## site map

- home (login, logout)
- patients list view (accessible by nurses & admin, can add/edit patient)
	- patient view (list of notes, nurses can add new notes and edit existing)
- admin (add/edit users: admin & nurses)
- profile (view/modify personal details for logged-in user)

## application design

#### frontend

Svelte will be used to compile frontend into static html/css/js pages that can be served over HTTP(S). JWT is used to authenticate client (web browser), as well as validate client session using stateless servers. HTTPS should be employed when deployed in production environment.

One advantage of using svelte is its extremely small and fast compiled size. site built with svelt can fun very fast and require little resources, this can be advantageous especially in context of mobile devices. browser compatibility of the target mobile device should be investigated.

internationalization possibiles will not be address; if needed, it should be possible to achieve localization using plugins such as [svelte-i18n](https://github.com/kaisermann/svelte-i18n).

#### backend

for the sake of keeping this demo simple during developemnt, a single goLang backend service is used to authenticate and serve all api requests. the backend will be used to authenticate client session by providing JWT tokens; it should also be easy to use a 3rd party authentication provider instead. 

if/when the need arises to separate the backend into individual microserves, the api structure is designed to make this process relatively simple. some code refactoring may need to be done, for example, the client JWT token validation may be refactored to make a

data is persisted by backend using sql interface. any popular sql backend such my mysql or postgres may be used.

#### services & data modelling

- auth service
	- /api/v1/session
		- POST/DELETE (login, logout)

- user service
	- /api/v1/users
		- GET
	- /api/v1/user?user-id=
		- POST/GET/DELETE

	- dao nurse
		- username (login id)
		- name (full name)
		- enabled
		- last-login
		- created-time
		- last-modified

- patient service		
	- /api/v1/patients
		- GET
	- /api/v1/patient?patient-id=
		- POST/GET/DELETE
	- /api/v1/patient-notes?patient-id=
		- GET
	- /api/v1/patient-note?note-id=
		- POST/GET/DELETE

	- dao patient
		- name (full name)
		- location (facility-id, bed-id, room-id, etc)
		- created-time
		- last-modified
	- dao patient-note
		- note-id
		- nurse-id
		- patient-id
		- note (text)
		- created-time
		- last-modified

## repository folder structure

- project root
	- api (folder)
		- api definition
		- stub generation scripts
	- frontend (folder)
		- src code, svelte
		- Dockerfiles for dev & testing
	- backend (folder)
		- src code
		- Dockerfiles for dev & testing
	- db schema (folder)
		- schema definition
		- pre-seeded data for dev/testing
		- db schema migration scripts
	- integration tests (folder)
		- synthetic testing (eg. selenium)
		- Docker-compose files for dev & testing
	- cicd (folder)
		- Dockerfiles for ci/cd & production services
		- cloud deployment manifests (k8s, terraform, etc)
		- ci/cd bash scripts

## ci/cd considerations

code lint, unit testing, gui testing in ci/cd. Dockerfiles provided to facilitate build, test, and service execution environments.

## sample deployment architecture

frontend website can deployed as static website; CDN may be utilized. Backend services can be deployed as stateless containerized (K8S) or serverless cloud architecture, using L7 cloud load balancers with TLS termination. A reverse proxy (or load balancer with path-based-routing support) may be used to scale different APIs individually, even if the codebase for different APIs is not separated.

## service monitoring consideration

frontend can be instrucmented with services such as Google Analytics. backend goLang services can be instrumented for Prometheus using official [go client](https://github.com/prometheus/client_golang), and for OpenTracing using [OpenTracing API for Go](https://github.com/opentracing/opentracing-go)
