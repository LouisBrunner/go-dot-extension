// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type HTTPClient struct {
  obj gdc.ObjectPtr
}

func createHTTPClient(obj gdc.ObjectPtr) *HTTPClient {
  return &HTTPClient{
    obj: obj,
  }
}

func (me *HTTPClient) BaseClass() string {
  return "HTTPClient"
}
