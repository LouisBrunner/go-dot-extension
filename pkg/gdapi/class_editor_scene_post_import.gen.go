// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type EditorScenePostImport struct {
  obj gdc.ObjectPtr
}

func (me *EditorScenePostImport) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *EditorScenePostImport) BaseClass() string {
  return "EditorScenePostImport"
}



// Enums

func (me *EditorScenePostImport) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EditorScenePostImport) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *EditorScenePostImport) XPostImport(scene Node, )  {
  panic("TODO: implement")
}

func  (me *EditorScenePostImport) GetSourceFile()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
