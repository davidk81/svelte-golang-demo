package patient

import (
	"encoding/json"
	"log"

	"github.com/davidk81/svelte-golang-demo/backend/patientdb/models"
	"github.com/davidk81/svelte-golang-demo/backend/session"
	"github.com/valyala/fasthttp"
)

// HandlePatientNote entrypoint http request handler /patient/note
func HandlePatientNote(ctx *fasthttp.RequestCtx) error {
	session.VerifySession(ctx, "nurse")
	switch string(ctx.Request.Header.Method()) {
	case "POST":
		return handleMethodNotePost(ctx)
	default:
		ctx.Error("Unsupported path", fasthttp.StatusNotFound)
	}
	return nil
}

// HandlePatientNoteList entrypoint http request handler /patient/notes
func HandlePatientNoteList(ctx *fasthttp.RequestCtx) error {
	session.VerifySession(ctx, "nurse")
	switch string(ctx.Request.Header.Method()) {
	case "GET":
		handleMethodNoteGetList(ctx)
	default:
		ctx.Error("Unsupported path", fasthttp.StatusNotFound)
	}
	return nil
}

func handleMethodNotePost(ctx *fasthttp.RequestCtx) error {
	// decode post body
	var note models.PatientNote
	err := json.Unmarshal(ctx.Request.Body(), &note)
	if err != nil {
		return err
	}
	err = AddPatientNote(&note, ctx)
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
	// return patient info in response
	p, err := GetPatients(ctx)
	if err != nil {
		return err
	}
	b, err := json.Marshal(p)
	log.Println(string(b))
	if err != nil {
		return err
	}
	ctx.SetBody([]byte(b))
	ctx.SetStatusCode(fasthttp.StatusCreated)
	return nil
}