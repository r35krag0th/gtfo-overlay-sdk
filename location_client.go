package gtfo_overlay_sdk

import "fmt"

type LocationClient struct {
	baseClient
}

func (c LocationClient) Update(d int, z int, a string) error {
	r, err := c.makePutRequest(fmt.Sprintf("/api/v1/location/update?dimension=%d&zone=%d&area=%s", d, z, a))
	if err != nil {
		return err
	}
	if r.Response().StatusCode >= 200 && r.Response().StatusCode < 300 {
		return nil
	}
	return fmt.Errorf("updating location to dimension %d of %d%s failed with unexpected status: %d - %s", d, z, a, r.Response().StatusCode, r.Response().Status)
}
