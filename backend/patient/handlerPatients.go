package patient

// handles http requests for /patient and /patients

import (
	"encoding/json"

	"github.com/davidk81/svelte-golang-demo/backend/patientdb/models"
	"github.com/valyala/fasthttp"
)

// HandlePatient entrypoint http request handler for /patient
func HandlePatient(ctx *fasthttp.RequestCtx) error {
	switch string(ctx.Request.Header.Method()) {
	case "GET":
		return handleMethodGet(ctx)
	case "POST":
		return handleMethodPost(ctx) // not implemented
	case "DELETE":
		return handleMethodDelete(ctx) // not implemented
	default:
		ctx.NotFound()
		return nil
	}
}

// HandlePatientList entrypoint http request handler for /patients
func HandlePatientList(ctx *fasthttp.RequestCtx) error {
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
	ctx.SetStatusCode(fasthttp.StatusNotImplemented)
	return nil
}

func handleMethodPost(ctx *fasthttp.RequestCtx) error {
	// decode post body
	var patient models.Patient
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
	// ctx.SetStatusCode(fasthttp.StatusCreated)
	ctx.SetStatusCode(fasthttp.StatusNotImplemented)
	return nil
}

func handleMethodGetList(ctx *fasthttp.RequestCtx) error {
	p, err := getPatients(ctx)
	if err != nil {
		return err
	}
	b, err := json.Marshal(p)
	if err != nil {
		return err
	}
	ctx.SetBody([]byte(b))
	ctx.SetStatusCode(fasthttp.StatusOK)
	return nil
}

func handleMethodGet(ctx *fasthttp.RequestCtx) error {
	patientID := string(ctx.QueryArgs().Peek("patientid"))

	// return patient info in response
	p, err := getPatient(ctx, patientID)
	if err != nil {
		return err
	}
	b, err := json.Marshal(p)
	if err != nil {
		return err
	}
	ctx.SetBody([]byte(b))
	ctx.SetStatusCode(fasthttp.StatusOK)
	return nil
}
