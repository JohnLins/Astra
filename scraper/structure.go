package scraper

import (
	"time"
)

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

		// AlternateName .
		// AlternateName []string

		// ReleaseDate .
		ReleaseDate time.Time

		// Description .
		Description string

		// Contact .
		Contact Contact

		// InformationURL .
		InformationURL []struct {
			// Name .
			Name string

			// URL .
			URL string

			// Description .
			Description string
		}

		// PriorID .
		PriorID string
	}

	// Location .
	Location Location
}

// Contact .
type Contact struct {
	// PersonID .
	PersonID string

	// Role .
	Role string
}

// Location .
type Location struct {
	// ObservatoryRegion .
	ObservatoryRegion []string
}
