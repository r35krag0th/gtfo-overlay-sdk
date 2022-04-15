package gtfo_overlay_sdk

// Client is the SDK Client for the GTFO Overlay SDK Client.
// It has hard-coded endpoints based, and conditionally can use the development environment based by setting
// the useDev flag.
type Client struct {
	useDev     bool
	Expedition *ExpeditionClient
	Sector     *SectorClient
	Location   *LocationClient
}

// NewClient creates a GTFO Overlay SDK Client
func NewClient(useDevEndpoint bool) *Client {
	return &Client{
		useDev: useDevEndpoint,
		Expedition: &ExpeditionClient{
			baseClient{useDev: useDevEndpoint},
		},
		Sector: &SectorClient{
			baseClient{useDev: useDevEndpoint},
		},
		Location: &LocationClient{
			baseClient{useDev: useDevEndpoint},
		},
	}
}
