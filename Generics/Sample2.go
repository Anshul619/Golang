package main

// UserRequest is a request for user information.
type UserRequest struct {
    Login string
}

// GroupRequest is a request for group information.
type GroupRequest struct {
    ID string
}

// Request is the set of all possible request types.
type Request interface {
    UserRequest | GroupRequest
}

// UnmarshalJSON safely decodes JSON into a pointer to a specific type.
func UnmarshalJSON[T Request](data []byte, d *T) error {
    return json.Unmarshal(data, d)
}