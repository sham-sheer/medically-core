package nlp_processor

import (
        "context"
		"log"

        language "cloud.google.com/go/language/apiv1"
        languagepb "google.golang.org/genproto/googleapis/cloud/language/v1"
		"google.golang.org/api/option"
)

type MCGCL struct {
	client *language.Client
}

var (
	credentialsFile = "credentials.json"
)

func (gcl *MCGCL) analyzeEntities(ctx context.Context) error {
	req := &languagepb.AnalyzeEntitiesRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/language/v1#AnalyzeEntitiesRequest.
	}
	resp, err := gcl.client.AnalyzeEntities(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
	return nil
}


func New(ctx context.Context) (MCGCL, error)  {
	c, err := language.NewClient(ctx, option.WithCredentialsFile(credentialsFile))
	if err != nil {
		log.Panicf("error in setting up GCS client with credentials %v", credentialsFile)
		return MCGCL{}, err
	}
	defer c.Close()
	gcl := MCGCL{}
	gcl.client = c

	return gcl, nil
}