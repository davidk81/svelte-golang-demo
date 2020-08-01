package patient

import (
	"encoding/json"

	"github.com/davidk81/svelte-golang-demo/backend/session"
	"github.com/valyala/fasthttp"
)

//Patient struct for (de)marshalling
type Patient struct {
	PatientID string `json:"patientId"`
	Name      string `json:"name"`
}

// HandlePatient entrypoint http request handler
func HandlePatient(ctx *fasthttp.RequestCtx) error {
	session.VerifySession(ctx, "nurse")
	switch string(ctx.Request.Header.Method()) {
	case "POST":
		return handleMethodPost(ctx)
	case "DELETE":
		return handleMethodDelete(ctx)
	default:
		ctx.NotFound()
		return nil
	}
}

// HandlePatientList entrypoint http request handler
func HandlePatientList(ctx *fasthttp.RequestCtx) error {
	session.VerifySession(ctx, "nurse")
	switch string(ctx.Request.Header.Method()) {
	case "GET":
		return handleMethodGetList(ctx)
	default:
		ctx.NotFound()
		return nil
	}
}

func handleMethodDelete(ctx *fasthttp.RequestCtx) error {
	// TODO
	ctx.SetStatusCode(fasthttp.StatusOK)
	return nil
}

func handleMethodPost(ctx *fasthttp.RequestCtx) error {
	// decode post body
	var patient Patient
	err := json.Unmarshal(ctx.Request.Body(), &patient)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return nil
	}

	// TODO: validate data
	// TODO: insert to db

	// return patient info in response
	b, err := json.Marshal(patient)
	if err != nil {
		return err
	}
	ctx.SetBody([]byte(b))
	ctx.SetStatusCode(fasthttp.StatusCreated)
	return nil
}

func handleMethodGetList(ctx *fasthttp.RequestCtx) error {
	// return patient info in response
	p, err := GetPatients(ctx)
	if err != nil {
		return err
	}
	b, err := json.Marshal(p)
	if err != nil {
		return err
	}
	ctx.SetBody([]byte(b))
	ctx.SetStatusCode(fasthttp.StatusCreated)
	return nil
}
