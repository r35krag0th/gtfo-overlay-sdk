package gtfo_overlay_sdk

import "fmt"

type SectorClient struct {
	baseClient
}

func (c SectorClient) SetMain() error {
	return c.Set("Main")
}

func (c SectorClient) SetSecondary() error {
	return c.Set("Secondary")
}

func (c SectorClient) SetOverload() error {
	return c.Set("Overload")
}

func (c SectorClient) SetPE() error {
	return c.Set("Prisoner Efficiency")
}

func (c SectorClient) Get() error {
	r, err := c.makeGetRequest("/api/v1/overlay/sector")
	if err != nil {
		return err
	}

	// Healthy response
	if r.Response().StatusCode >= 200 && r.Response().StatusCode < 300 {
		return nil
	}

	// Treat everything else as an error
	return fmt.Errorf("getting current sector failed with unexpected status: %d - %s", r.Response().StatusCode, r.Response().Status)
}

func (c SectorClient) Set(sectorName string) error {
	r, err := c.makePutRequest(fmt.Sprintf("/api/v1/overlay/sector?value=%s", sectorName))
	// A problem with the request itself
	if err != nil {
		return err
	}

	// Healthy response
	if r.Response().StatusCode >= 200 && r.Response().StatusCode < 300 {
		return nil
	}

	// Treat everything else as an error
	return fmt.Errorf("setting sector to %s failed with unexpected status: %d - %s", sectorName, r.Response().StatusCode, r.Response().Status)
}
