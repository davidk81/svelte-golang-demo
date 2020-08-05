package patient

// handles http requests for /patient/note and /patient/notes

import (
	"encoding/json"

	"github.com/davidk81/svelte-golang-demo/backend/patientdb/models"
	"github.com/valyala/fasthttp"
)

// HandlePatientNote entrypoint http request handler /patient/note
func HandlePatientNote(ctx *fasthttp.RequestCtx) error {
	switch string(ctx.Request.Header.Method()) {
	case "POST":
		return handleMethodNotePost(ctx)
	default:
		ctx.NotFound()
		return nil
	}
}

// HandlePatientNoteList entrypoint http request handler /patient/notes
func HandlePatientNoteList(ctx *fasthttp.RequestCtx) error {
	switch string(ctx.Request.Header.Method()) {
	case "GET":
		return handleMethodNoteGetList(ctx)
	default:
		ctx.NotFound()
		return nil
	}
}

func handleMethodNotePost(ctx *fasthttp.RequestCtx) error {
	// decode post body
	var note models.PatientNote
	err := json.Unmarshal(ctx.Request.Body(), &note)
	if err != nil {
		return err
	}

	err = addPatientNote(ctx, &note)
	if err != nil {
		return err
	}

	// return patient info in response
	b, err := json.Marshal(note)
	if err != nil {
		return err
	}
	ctx.SetBody([]byte(b))
	ctx.SetStatusCode(fasthttp.StatusCreated)
	return nil
}

func handleMethodNoteGetList(ctx *fasthttp.RequestCtx) error {
	patientID := string(ctx.QueryArgs().Peek("patientid"))

	// return patient info in response
	p, err := getPatientNotes(ctx, patientID)
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
