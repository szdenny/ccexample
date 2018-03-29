package bean

// Define the car structure, with 4 properties.  Structure tags are used by encoding/json library
type Car struct {
	ID string `json:"id"`
	Make   string `json:"make"`
	Model  string `json:"model"`
	Colour string `json:"colour"`
	Owner  string `json:"owner"`
}
