// Copyright (c) 2017 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0
//

package vcmock

import (
	"context"
	"syscall"

	vc "github.com/kata-containers/runtime/virtcontainers"
	"github.com/kata-containers/runtime/virtcontainers/device/api"
	"github.com/kata-containers/runtime/virtcontainers/device/config"
	. "github.com/kata-containers/runtime/virtcontainers/pkg/types"
	"github.com/kata-containers/runtime/virtcontainers/types"
	specs "github.com/opencontainers/runtime-spec/specs-go"
	"github.com/sirupsen/logrus"
)

// Sandbox is a fake Sandbox type used for testing
type Sandbox struct {
	MockID          SandboxID
	MockURL         string
	MockAnnotations map[string]string
	MockContainers  []*Container
	MockNetNs       string
}

// Container is a fake Container type used for testing
type Container struct {
	MockID          ContainerID
	MockURL         string
	MockToken       string
	MockProcess     vc.Process
	MockPid         int
	MockSandbox     *Sandbox
	MockAnnotations map[string]string
}

// VCMock is a type that provides an implementation of the VC interface.
// It is used for testing.
type VCMock struct {
	SetLoggerFunc  func(ctx context.Context, logger *logrus.Entry)
	SetFactoryFunc func(ctx context.Context, factory vc.Factory)

	CreateSandboxFunc  func(ctx context.Context, sandboxConfig vc.SandboxConfig) (vc.VCSandbox, error)
	DeleteSandboxFunc  func(ctx context.Context, sandboxID SandboxID) (vc.VCSandbox, error)
	ListSandboxFunc    func(ctx context.Context) ([]vc.SandboxStatus, error)
	FetchSandboxFunc   func(ctx context.Context, sandboxID SandboxID) (vc.VCSandbox, error)
	PauseSandboxFunc   func(ctx context.Context, sandboxID SandboxID) (vc.VCSandbox, error)
	ResumeSandboxFunc  func(ctx context.Context, sandboxID SandboxID) (vc.VCSandbox, error)
	RunSandboxFunc     func(ctx context.Context, sandboxConfig vc.SandboxConfig) (vc.VCSandbox, error)
	StartSandboxFunc   func(ctx context.Context, sandboxID SandboxID) (vc.VCSandbox, error)
	StatusSandboxFunc  func(ctx context.Context, sandboxID SandboxID) (vc.SandboxStatus, error)
	StatsContainerFunc func(ctx context.Context, sandboxID SandboxID, containerID ContainerID) (vc.ContainerStats, error)
	StopSandboxFunc    func(ctx context.Context, sandboxID SandboxID) (vc.VCSandbox, error)

	CreateContainerFunc      func(ctx context.Context, sandboxID SandboxID, containerConfig vc.ContainerConfig) (vc.VCSandbox, vc.VCContainer, error)
	DeleteContainerFunc      func(ctx context.Context, sandboxID SandboxID, containerID ContainerID) (vc.VCContainer, error)
	EnterContainerFunc       func(ctx context.Context, sandboxID SandboxID, containerID ContainerID, cmd types.Cmd) (vc.VCSandbox, vc.VCContainer, *vc.Process, error)
	KillContainerFunc        func(ctx context.Context, sandboxID SandboxID, containerID ContainerID, signal syscall.Signal, all bool) error
	StartContainerFunc       func(ctx context.Context, sandboxID SandboxID, containerID ContainerID) (vc.VCContainer, error)
	StatusContainerFunc      func(ctx context.Context, sandboxID SandboxID, containerID ContainerID) (vc.ContainerStatus, error)
	StopContainerFunc        func(ctx context.Context, sandboxID SandboxID, containerID ContainerID) (vc.VCContainer, error)
	ProcessListContainerFunc func(ctx context.Context, sandboxID SandboxID, containerID ContainerID, options vc.ProcessListOptions) (vc.ProcessList, error)
	UpdateContainerFunc      func(ctx context.Context, sandboxID SandboxID, containerID ContainerID, resources specs.LinuxResources) error
	PauseContainerFunc       func(ctx context.Context, sandboxID SandboxID, containerID ContainerID) error
	ResumeContainerFunc      func(ctx context.Context, sandboxID SandboxID, containerID ContainerID) error

	AddDeviceFunc func(ctx context.Context, sandboxID SandboxID, info config.DeviceInfo) (api.Device, error)

	AddInterfaceFunc    func(ctx context.Context, sandboxID SandboxID, inf *Interface) (*Interface, error)
	RemoveInterfaceFunc func(ctx context.Context, sandboxID SandboxID, inf *Interface) (*Interface, error)
	ListInterfacesFunc  func(ctx context.Context, sandboxID SandboxID) ([]*Interface, error)
	UpdateRoutesFunc    func(ctx context.Context, sandboxID SandboxID, routes []*Route) ([]*Route, error)
	ListRoutesFunc      func(ctx context.Context, sandboxID SandboxID) ([]*Route, error)
}
