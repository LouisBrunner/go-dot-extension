// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type HTTPRequest struct {
  obj gdc.ObjectPtr
}

func createHTTPRequest(obj gdc.ObjectPtr) *HTTPRequest {
  return &HTTPRequest{
    obj: obj,
  }
}

func (me *HTTPRequest) BaseClass() string {
  return "HTTPRequest"
}
