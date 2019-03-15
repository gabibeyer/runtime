// Copyright (c) 2017 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0
//

package virtcontainers

import (
	"context"
	"io"
	"syscall"

	"github.com/kata-containers/runtime/virtcontainers/device/api"
	"github.com/kata-containers/runtime/virtcontainers/device/config"
	. "github.com/kata-containers/runtime/virtcontainers/pkg/types"
	"github.com/kata-containers/runtime/virtcontainers/types"
	specs "github.com/opencontainers/runtime-spec/specs-go"
	"github.com/sirupsen/logrus"
)

// VC is the Virtcontainers interface
type VC interface {
	SetLogger(ctx context.Context, logger *logrus.Entry)
	SetFactory(ctx context.Context, factory Factory)

	CreateSandbox(ctx context.Context, sandboxConfig SandboxConfig) (VCSandbox, error)
	DeleteSandbox(ctx context.Context, sandboxID SandboxID) (VCSandbox, error)
	FetchSandbox(ctx context.Context, sandboxID SandboxID) (VCSandbox, error)
	ListSandbox(ctx context.Context) ([]SandboxStatus, error)
	PauseSandbox(ctx context.Context, sandboxID SandboxID) (VCSandbox, error)
	ResumeSandbox(ctx context.Context, sandboxID SandboxID) (VCSandbox, error)
	RunSandbox(ctx context.Context, sandboxConfig SandboxConfig) (VCSandbox, error)
	StartSandbox(ctx context.Context, sandboxID SandboxID) (VCSandbox, error)
	StatusSandbox(ctx context.Context, sandboxID SandboxID) (SandboxStatus, error)
	StopSandbox(ctx context.Context, sandboxID SandboxID) (VCSandbox, error)

	CreateContainer(ctx context.Context, sandboxID SandboxID, containerConfig ContainerConfig) (VCSandbox, VCContainer, error)
	DeleteContainer(ctx context.Context, sandboxID SandboxID, containerID ContainerID) (VCContainer, error)
	EnterContainer(ctx context.Context, sandboxID SandboxID, containerID ContainerID, cmd types.Cmd) (VCSandbox, VCContainer, *Process, error)
	KillContainer(ctx context.Context, sandboxID SandboxID, containerID ContainerID, signal syscall.Signal, all bool) error
	StartContainer(ctx context.Context, sandboxID SandboxID, containerID ContainerID) (VCContainer, error)
	StatusContainer(ctx context.Context, sandboxID SandboxID, containerID ContainerID) (ContainerStatus, error)
	StatsContainer(ctx context.Context, sandboxID SandboxID, containerID ContainerID) (ContainerStats, error)
	StopContainer(ctx context.Context, sandboxID SandboxID, containerID ContainerID) (VCContainer, error)
	ProcessListContainer(ctx context.Context, sandboxID SandboxID, containerID ContainerID, options ProcessListOptions) (ProcessList, error)
	UpdateContainer(ctx context.Context, sandboxID SandboxID, containerID ContainerID, resources specs.LinuxResources) error
	PauseContainer(ctx context.Context, sandboxID SandboxID, containerID ContainerID) error
	ResumeContainer(ctx context.Context, sandboxID SandboxID, containerID ContainerID) error

	AddDevice(ctx context.Context, sandboxID SandboxID, info config.DeviceInfo) (api.Device, error)

	AddInterface(ctx context.Context, sandboxID SandboxID, inf *Interface) (*Interface, error)
	RemoveInterface(ctx context.Context, sandboxID SandboxID, inf *Interface) (*Interface, error)
	ListInterfaces(ctx context.Context, sandboxID SandboxID) ([]*Interface, error)
	UpdateRoutes(ctx context.Context, sandboxID SandboxID, routes []*Route) ([]*Route, error)
	ListRoutes(ctx context.Context, sandboxID SandboxID) ([]*Route, error)
}

// VCSandbox is the Sandbox interface
// (required since virtcontainers.Sandbox only contains private fields)
type VCSandbox interface {
	Annotations(key string) (string, error)
	GetNetNs() string
	GetAllContainers() []VCContainer
	GetAnnotations() map[string]string
	GetContainer(containerID ContainerID) VCContainer
	ID() SandboxID
	SetAnnotations(annotations map[string]string) error

	Start() error
	Stop() error
	Pause() error
	Resume() error
	Release() error
	Monitor() (chan error, error)
	Delete() error
	Status() SandboxStatus
	CreateContainer(contConfig ContainerConfig) (VCContainer, error)
	DeleteContainer(contID ContainerID) (VCContainer, error)
	StartContainer(containerID ContainerID) (VCContainer, error)
	StopContainer(containerID ContainerID) (VCContainer, error)
	KillContainer(containerID ContainerID, signal syscall.Signal, all bool) error
	StatusContainer(containerID ContainerID) (ContainerStatus, error)
	StatsContainer(containerID ContainerID) (ContainerStats, error)
	PauseContainer(containerID ContainerID) error
	ResumeContainer(containerID ContainerID) error
	EnterContainer(containerID ContainerID, cmd types.Cmd) (VCContainer, *Process, error)
	UpdateContainer(containerID ContainerID, resources specs.LinuxResources) error
	ProcessListContainer(containerID ContainerID, options ProcessListOptions) (ProcessList, error)
	WaitProcess(containerID ContainerID, processID string) (int32, error)
	SignalProcess(containerID ContainerID, processID string, signal syscall.Signal, all bool) error
	WinsizeProcess(containerID ContainerID, processID string, height, width uint32) error
	IOStream(containerID ContainerID, processID string) (io.WriteCloser, io.Reader, io.Reader, error)

	AddDevice(info config.DeviceInfo) (api.Device, error)

	AddInterface(inf *Interface) (*Interface, error)
	RemoveInterface(inf *Interface) (*Interface, error)
	ListInterfaces() ([]*Interface, error)
	UpdateRoutes(routes []*Route) ([]*Route, error)
	ListRoutes() ([]*Route, error)
}

// VCContainer is the Container interface
// (required since virtcontainers.Container only contains private fields)
type VCContainer interface {
	GetAnnotations() map[string]string
	GetPid() int
	GetToken() string
	ID() ContainerID
	Sandbox() VCSandbox
	Process() Process
	SetPid(pid int) error
}
