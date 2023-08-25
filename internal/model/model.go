package model

// User is a model used by this application. For simplicity,
// this is used publicly by all layers. However, in a real world
// application each layer will have its own model and models are
// transofrmed from one layer to another, since models might be
// inconsitent across different layers (e.g. presentation model,
// business logic model, and storage model).
type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
