// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type TranslationServer struct {
  obj gdc.ObjectPtr
}

func createTranslationServer(obj gdc.ObjectPtr) *TranslationServer {
  return &TranslationServer{
    obj: obj,
  }
}

func (me *TranslationServer) BaseClass() string {
  return "TranslationServer"
}
