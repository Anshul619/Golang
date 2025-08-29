package user

import "github.com/jarcoal/httpmock"

func TestHTTP() {
    httpmock.Activate()
    defer httpmock.DeactivateAndReset()

    httpmock.RegisterResponder("GET", "URL",
        func(req *http.Request) (*http.Response, error) {
            return httpmock.NewBytesResponse(200, textBytes), nil
        })
}