// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type MovieWriter struct {
  obj gdc.ObjectPtr
}

func createMovieWriter(obj gdc.ObjectPtr) *MovieWriter {
  return &MovieWriter{
    obj: obj,
  }
}

func (me *MovieWriter) BaseClass() string {
  return "MovieWriter"
}
