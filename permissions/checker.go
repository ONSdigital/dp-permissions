package permissions

import (
	"context"

	"github.com/ONSdigital/log.go/log"
)

func NewChecker(host string, httpClient HTTPClient) *Checker {
	return &Checker{
		host: host,
		cli:  httpClient,
	}
}

func (c *Checker) Check(ctx context.Context, required CRUD, serviceToken string, userToken string, collectionID string, datasetID string) (int, error) {
	data := log.Data{
		"collection_id": collectionID,
		"dataset_id":    datasetID,
		"user_token":    userToken != "",
		"service_token": serviceToken != "",
	}

	r, err := c.getPermissionsRequest(serviceToken, userToken, collectionID, datasetID)
	if err != nil {
		return 0, err
	}

	resp, err := c.cli.Do(ctx, r)
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		//
		return handleErrorResponse(ctx, resp, data), nil
	}

	return handleSuccessfulResponse(ctx, resp.Body, &required, data)
}
