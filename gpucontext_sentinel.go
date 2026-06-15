package wgpu

// gpucontext sentinel methods — satisfy gpucontext.Device/Queue/Adapter/Surface/Instance
// interfaces. These unexported methods prevent nil or arbitrary types from being
// passed where a real GPU object is expected (compile-time safety).
//
// Linter flags these as "unused" because it doesn't see cross-package interface
// satisfaction. The methods ARE used — gpucontext.Device requires gpuDevice(), etc.

func (*Device) gpuDevice()     {} //nolint:unused // implements gpucontext.Device
func (*Queue) gpuQueue()       {} //nolint:unused // implements gpucontext.Queue
func (*Adapter) gpuAdapter()   {} //nolint:unused // implements gpucontext.Adapter
func (*Surface) gpuSurface()   {} //nolint:unused // implements gpucontext.Surface
func (*Instance) gpuInstance() {} //nolint:unused // implements gpucontext.Instance
