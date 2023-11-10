// models_gen.go

package model

// Character represents a character in the Rick and Morty universe.
type Character struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Status   string    `json:"status,omitempty"`
	Species  string    `json:"species,omitempty"`
	Type     string    `json:"type,omitempty"`
	Gender   string    `json:"gender,omitempty"`
	Origin   *Location `json:"origin,omitempty"`
	Location *Location `json:"location,omitempty"`
	Image    string    `json:"image,omitempty"`
}

// Episode represents an episode in the Rick and Morty universe.
type Episode struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Location represents a location in the Rick and Morty universe.
type Location struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// RickAndMortyAssociation represents a link between a Rick and his Morties.
type RickAndMortyAssociation struct {
	Rick    *Character   `json:"rick"`
	Morties []*Character `json:"morties"`
}

// Add any additional types and enums as defined in your GraphQL schema.
