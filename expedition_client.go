package gtfo_overlay_sdk

import "fmt"

type ExpeditionClient struct {
	baseClient
}

func (c ExpeditionClient) Get() error {
	r, err := c.makeGetRequest("/api/v1/overlay/expedition")
	if err != nil {
		return err
	}

	// Healthy response
	if r.Response().StatusCode >= 200 && r.Response().StatusCode < 300 {
		return nil
	}

	// Treat everything else as an error
	return fmt.Errorf("getting current expedition failed with unexpected status: %d - %s", r.Response().StatusCode, r.Response().Status)
}

func (c ExpeditionClient) Set(expeditionId string) error {
	r, err := c.makePutRequest(fmt.Sprintf("/api/v1/overlay/expedition?value=%s", expeditionId))
	// A problem with the request itself
	if err != nil {
		return err
	}

	// Healthy response
	if r.Response().StatusCode >= 200 && r.Response().StatusCode < 300 {
		return nil
	}

	// Treat everything else as an error
	return fmt.Errorf("setting expedition to %s failed with unexpected status: %d - %s", expeditionId, r.Response().StatusCode, r.Response().Status)
}
