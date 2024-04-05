// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type singletons struct {
  Performance Performance
  TextServerManager TextServerManager
  PhysicsServer2DManager PhysicsServer2DManager
  PhysicsServer3DManager PhysicsServer3DManager
  NavigationMeshGenerator NavigationMeshGenerator
  ProjectSettings ProjectSettings
  IP IP
  Geometry2D Geometry2D
  Geometry3D Geometry3D
  ResourceLoader ResourceLoader
  ResourceSaver ResourceSaver
  OS OS
  Engine Engine
  ClassDB ClassDB
  Marshalls Marshalls
  TranslationServer TranslationServer
  Input Input
  InputMap InputMap
  EngineDebugger EngineDebugger
  Time Time
  GDExtensionManager GDExtensionManager
  ResourceUID ResourceUID
  WorkerThreadPool WorkerThreadPool
  ThemeDB ThemeDB
}

func newSingletons(iface gdc.Interface) *singletons {
  strPerformance := StringNameFromStr("Performance")
  defer strPerformance.Destroy()
  strTextServerManager := StringNameFromStr("TextServerManager")
  defer strTextServerManager.Destroy()
  strPhysicsServer2DManager := StringNameFromStr("PhysicsServer2DManager")
  defer strPhysicsServer2DManager.Destroy()
  strPhysicsServer3DManager := StringNameFromStr("PhysicsServer3DManager")
  defer strPhysicsServer3DManager.Destroy()
  strNavigationMeshGenerator := StringNameFromStr("NavigationMeshGenerator")
  defer strNavigationMeshGenerator.Destroy()
  strProjectSettings := StringNameFromStr("ProjectSettings")
  defer strProjectSettings.Destroy()
  strIP := StringNameFromStr("IP")
  defer strIP.Destroy()
  strGeometry2D := StringNameFromStr("Geometry2D")
  defer strGeometry2D.Destroy()
  strGeometry3D := StringNameFromStr("Geometry3D")
  defer strGeometry3D.Destroy()
  strResourceLoader := StringNameFromStr("ResourceLoader")
  defer strResourceLoader.Destroy()
  strResourceSaver := StringNameFromStr("ResourceSaver")
  defer strResourceSaver.Destroy()
  strOS := StringNameFromStr("OS")
  defer strOS.Destroy()
  strEngine := StringNameFromStr("Engine")
  defer strEngine.Destroy()
  strClassDB := StringNameFromStr("ClassDB")
  defer strClassDB.Destroy()
  strMarshalls := StringNameFromStr("Marshalls")
  defer strMarshalls.Destroy()
  strTranslationServer := StringNameFromStr("TranslationServer")
  defer strTranslationServer.Destroy()
  strInput := StringNameFromStr("Input")
  defer strInput.Destroy()
  strInputMap := StringNameFromStr("InputMap")
  defer strInputMap.Destroy()
  strEngineDebugger := StringNameFromStr("EngineDebugger")
  defer strEngineDebugger.Destroy()
  strTime := StringNameFromStr("Time")
  defer strTime.Destroy()
  strGDExtensionManager := StringNameFromStr("GDExtensionManager")
  defer strGDExtensionManager.Destroy()
  strResourceUID := StringNameFromStr("ResourceUID")
  defer strResourceUID.Destroy()
  strWorkerThreadPool := StringNameFromStr("WorkerThreadPool")
  defer strWorkerThreadPool.Destroy()
  strThemeDB := StringNameFromStr("ThemeDB")
  defer strThemeDB.Destroy()
  return &singletons{
    Performance: Performance{
      Object: Object{
        obj: ensurePtr(iface.GlobalGetSingleton(strPerformance.AsCPtr())),
      },
    },
    TextServerManager: TextServerManager{
      Object: Object{
        obj: ensurePtr(iface.GlobalGetSingleton(strTextServerManager.AsCPtr())),
      },
    },
    PhysicsServer2DManager: PhysicsServer2DManager{
      Object: Object{
        obj: ensurePtr(iface.GlobalGetSingleton(strPhysicsServer2DManager.AsCPtr())),
      },
    },
    PhysicsServer3DManager: PhysicsServer3DManager{
      Object: Object{
        obj: ensurePtr(iface.GlobalGetSingleton(strPhysicsServer3DManager.AsCPtr())),
      },
    },
    NavigationMeshGenerator: NavigationMeshGenerator{
      Object: Object{
        obj: ensurePtr(iface.GlobalGetSingleton(strNavigationMeshGenerator.AsCPtr())),
      },
    },
    ProjectSettings: ProjectSettings{
      Object: Object{
        obj: ensurePtr(iface.GlobalGetSingleton(strProjectSettings.AsCPtr())),
      },
    },
    IP: IP{
      Object: Object{
        obj: ensurePtr(iface.GlobalGetSingleton(strIP.AsCPtr())),
      },
    },
    Geometry2D: Geometry2D{
      Object: Object{
        obj: ensurePtr(iface.GlobalGetSingleton(strGeometry2D.AsCPtr())),
      },
    },
    Geometry3D: Geometry3D{
      Object: Object{
        obj: ensurePtr(iface.GlobalGetSingleton(strGeometry3D.AsCPtr())),
      },
    },
    ResourceLoader: ResourceLoader{
      Object: Object{
        obj: ensurePtr(iface.GlobalGetSingleton(strResourceLoader.AsCPtr())),
      },
    },
    ResourceSaver: ResourceSaver{
      Object: Object{
        obj: ensurePtr(iface.GlobalGetSingleton(strResourceSaver.AsCPtr())),
      },
    },
    OS: OS{
      Object: Object{
        obj: ensurePtr(iface.GlobalGetSingleton(strOS.AsCPtr())),
      },
    },
    Engine: Engine{
      Object: Object{
        obj: ensurePtr(iface.GlobalGetSingleton(strEngine.AsCPtr())),
      },
    },
    ClassDB: ClassDB{
      Object: Object{
        obj: ensurePtr(iface.GlobalGetSingleton(strClassDB.AsCPtr())),
      },
    },
    Marshalls: Marshalls{
      Object: Object{
        obj: ensurePtr(iface.GlobalGetSingleton(strMarshalls.AsCPtr())),
      },
    },
    TranslationServer: TranslationServer{
      Object: Object{
        obj: ensurePtr(iface.GlobalGetSingleton(strTranslationServer.AsCPtr())),
      },
    },
    Input: Input{
      Object: Object{
        obj: ensurePtr(iface.GlobalGetSingleton(strInput.AsCPtr())),
      },
    },
    InputMap: InputMap{
      Object: Object{
        obj: ensurePtr(iface.GlobalGetSingleton(strInputMap.AsCPtr())),
      },
    },
    EngineDebugger: EngineDebugger{
      Object: Object{
        obj: ensurePtr(iface.GlobalGetSingleton(strEngineDebugger.AsCPtr())),
      },
    },
    Time: Time{
      Object: Object{
        obj: ensurePtr(iface.GlobalGetSingleton(strTime.AsCPtr())),
      },
    },
    GDExtensionManager: GDExtensionManager{
      Object: Object{
        obj: ensurePtr(iface.GlobalGetSingleton(strGDExtensionManager.AsCPtr())),
      },
    },
    ResourceUID: ResourceUID{
      Object: Object{
        obj: ensurePtr(iface.GlobalGetSingleton(strResourceUID.AsCPtr())),
      },
    },
    WorkerThreadPool: WorkerThreadPool{
      Object: Object{
        obj: ensurePtr(iface.GlobalGetSingleton(strWorkerThreadPool.AsCPtr())),
      },
    },
    ThemeDB: ThemeDB{
      Object: Object{
        obj: ensurePtr(iface.GlobalGetSingleton(strThemeDB.AsCPtr())),
      },
    },
}
}
