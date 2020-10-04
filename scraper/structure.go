package scraper

// Satellite .
type Satellite struct {
	// Spase .
	Spase struct {
		// Version .
		Version string

		// Observatory .
		Observatory Observatory
	}
}

// Observatory .
type Observatory struct {
	// ResourceID .
	ResourceID string

	// ResourceHeader .
	ResourceHeader struct {
		// ResourceName .
		ResourceName string

		// Description .
		Description string

		// PriorID .
		PriorID string
	}
}
