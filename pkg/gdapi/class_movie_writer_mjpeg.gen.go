// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type MovieWriterMJPEG struct {
  obj gdc.ObjectPtr
}

func createMovieWriterMJPEG(obj gdc.ObjectPtr) *MovieWriterMJPEG {
  return &MovieWriterMJPEG{
    obj: obj,
  }
}

func (me *MovieWriterMJPEG) BaseClass() string {
  return "MovieWriterMJPEG"
}
