package main


// RequestParams - parameters for the DoRequest function
type RequestParams struct {
	Body         []byte
	URL          string
	Method       string

	Context context.Context // Context for request timeout
}

func requestWithContext(params RequestParams) ([]byte, error) {
	// request all the things
	b := make([]byte, 0)

	if params.Context == nil {
		params.Context = context.Background()
	}
	// Don't create a new timeout context - use the one passed in
	// This allows the caller to control the overall timeout
	ctx := params.Context

	// Only add timeout if the context doesn't already have one
	if _, ok := ctx.Deadline(); !ok {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(params.Context, 25*time.Second)
		defer cancel()
	}

	client := &http.Client{}

	var reader io.Reader
	if len(params.Body) > 0 {
		reader = bytes.NewBuffer(params.Body)
	}

	// Normal HTTP Request
	// req, err := http.NewRequest(params.Method, params.URL, reader)

	// Main HTTP method to call API with context
	req, err := http.NewRequestWithContext(ctx, params.Method, params.URL, reader)
	if err != nil {
		return b, RequestFailedError
	}

	req.Header.Add("Content-Type", "application/json")

	for k, v := range params.Headers {
		req.Header.Add(k, v)
	}

	// send the request
	resp, err := client.Do(req)

	// if there was a request error, fail
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("Request timed out after 25 seconds")
		} else if ctx.Err() == context.Canceled {
			log.Println("Request was cancelled")
		}
		log.Printf("Request error %s %s with body %s - %v", params.Method, params.URL, string(params.Body), err)
		return b, RequestFailedError
	}

	defer resp.Body.Close()
	b, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Invalid response error when requesting %s %s with body %s - %v", params.Method, params.URL, string(params.Body), err)
		return b, InvalidResponseError
	}

	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		// error handling
		return nil, err
	}

	return b, nil
}