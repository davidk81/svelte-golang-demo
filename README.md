# svelte-golang-demo

A demo app using Svelte as frontend and GoLang as backend. This app provides an web-based GUI for nurses to manage patient's health notes. 

## site map

- home (login, logout, register)
- patients list view (accessible by nurses & admin, can add/edit patient)
	- patient view (list of notes, nurses can add new notes and edit existing)
- ** admin (add/edit users: admin & nurses)
- ** profile (view/modify personal details for logged-in user)

** not implemented 

## application design

#### frontend

Svelte is used to compile frontend into static html/css/js pages that can be served over HTTP(S). JWT is used to authenticate client (web browser), as well as validate client session using stateless servers.

One advantage of using svelte is its extremely small and fast compiled size. Site built with svelt can fun very fast and require little resources, this can be advantageous especially in context of mobile devices. Browser compatibility of the target mobile device should be investigated.

Internationalization possibiles will not be address; if needed, it should be possible to achieve localization using plugins such as [svelte-i18n](https://github.com/kaisermann/svelte-i18n).

#### backend

A single goLang backend service is used to authenticate and serve all api requests. The backend will be used to authenticate client session by providing JWT tokens.

if/when the need arises to separate the backend into individual microserves, the backend folder structure is designed to make this process relatively simple. Some code refactoring may need to be done, for example, the session service may be updated to call user service endpoint instead of accessing it directly

Data is persisted by backend using postgres sql using (SQLBoiler)[https://github.com/volatiletech/sqlboiler]

#### services & data modelling

also refer to [api/README.md](api/README.md) for sample request/response

- auth service
	- /api/v1/session
		- POST/DELETE (login, logout)
	- /api/v1/regiser
		- POST (register new userz)

- user service
	- /api/v1/users
		- GET
	- /api/v1/user?user-id=
		- POST/GET/DELETE (not implemented)

	- dao user
		- userid (login id)
		- name (full name)
		- secret (salted hashed password)
		- last-login (not implemented)
		- created-time
		- last-modified (not implemented)

- patient service		
	- /api/v1/patients
		- GET (list)
	- /api/v1/patient?patient-id=
		- POST/GET/DELETE (single)
	- /api/v1/patient-notes?patient-id=
		- GET (list)
	- /api/v1/patient-note?note-id=
		- POST (new health note)
		- GET/DELETE (not implemented)

	- dao patient
		- patientid (uuid)
		- name (full name)
		- location (facility-id, bed-id, room-id, etc)
		- created-time
		- last-modified (not implemented)
	- dao patient-note
		- noteid (uuid)
		- userid (fk)
		- patientid (fk)
		- note (text)
		- created-time
		- last-modified (not implemented)

## repository folder structure

- project root
	- api (folder) -  (not implemented)
		- api definition
		- swagger/openapi stub generation scripts
	- frontend (folder)
		- src code, svelte
		- Dockerfiles for dev & testing
	- backend (folder)
		- src code
		- Dockerfiles for dev & testing
	- db-schema (folder)
		- schema definition
		- pre-seeded data for dev/testing
		- db schema migration scripts (https://github.com/golang-migrate/migrate)
		- db orm generator (sqlboiler, https://github.com/lqs/sqlingo)
	- integration-tests (folder) (not implemented)
		- synthetic testing (eg. selenium)
		- Docker-compose files for dev & testing
	- cicd (folder) (not implemented)
		- Dockerfiles for ci/cd & production services
		- cloud deployment manifests (k8s, terraform, etc)
		- ci/cd bash scripts

## ci/cd considerations

(not implemented) code lint, unit testing, gui testing in ci/cd. Dockerfiles provided to facilitate build, test, and service execution environments.

## sample deployment architecture

frontend website can deployed as static website; CDN may be utilized. Backend services can be deployed as stateless containerized (K8S) or serverless cloud architecture, using L7 cloud load balancers with TLS termination. A reverse proxy (or load balancer with path-based-routing support) may be used to scale different APIs individually, even if the codebase for different APIs is not separated.

## service monitoring consideration

frontend can be instrumented with services such as Google Analytics. backend goLang services can be instrumented for Prometheus using official [go client](https://github.com/prometheus/client_golang), and for OpenTracing using [OpenTracing API for Go](https://github.com/opentracing/opentracing-go)

## future work

- check licensing
- openapi/swagger definition & stub generation
- frontend unit tests
- frontend instrumentation (eg. Google Analytics)
- backend unit tests
- backend instrumentation (eg. prometheus, opentracing)
- integration tests
- code lint in ci/cd
- look into more secure authentication services/products & practices

## build & run compiled pages locally

to run frontend in developer mode, see [frontend/README.md](frontend/README.md)

```sh
# (re)build images & start
docker-compose up --build

# surf http://localhost:5000/

# stop
docker-compose down
```