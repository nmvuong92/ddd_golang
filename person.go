// entity package entities holds all the entities that are shared across subdomains
package tavern

import "github.com/google/uuid"

// Person is an entity that represents a person in all domains
type Person struct {
	// ID is an identifier of the entity
	ID   uuid.UUID
	Name string
	Age  int
}
