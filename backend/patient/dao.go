package patient

// handles database operations for patient and patient_notes tables

import (
	"context"

	"github.com/davidk81/svelte-golang-demo/backend/patientdb"
	"github.com/davidk81/svelte-golang-demo/backend/patientdb/models"
	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

// getPatients (list)
func getPatients(ctx context.Context) (models.PatientSlice, error) {
	return models.Patients(qm.Limit(20)).All(ctx, patientdb.DB())
}

// getPatient returns patient with matching patientID
func getPatient(ctx context.Context, patientID string) (*models.Patient, error) {
	return models.Patients(models.PatientWhere.Patientid.EQ(patientID)).One(ctx, patientdb.DB())
}

// getPatientNotes (list)
func getPatientNotes(ctx context.Context, patientID string) (models.PatientNoteSlice, error) {
	// TODO: add orderby clause
	return models.PatientNotes(models.PatientNoteWhere.Patient_Id.EQ(patientID)).All(ctx, patientdb.DB())
}

// addPatientNote adds a new note to the patient
func addPatientNote(ctx context.Context, note *models.PatientNote) error {
	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	note.Noteid = id.String()
	return (*note).Insert(ctx, patientdb.DB(), boil.Infer())
}
