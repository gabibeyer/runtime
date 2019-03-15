// Copyright (c) 2017 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0
//

package vcmock

import (
	"io"
	"syscall"

	vc "github.com/kata-containers/runtime/virtcontainers"
	"github.com/kata-containers/runtime/virtcontainers/device/api"
	"github.com/kata-containers/runtime/virtcontainers/device/config"
	. "github.com/kata-containers/runtime/virtcontainers/pkg/types"
	"github.com/kata-containers/runtime/virtcontainers/types"
	specs "github.com/opencontainers/runtime-spec/specs-go"
)

// ID implements the VCSandbox function of the same name.
func (s *Sandbox) ID() SandboxID {
	return s.MockID
}

// Annotations implements the VCSandbox function of the same name.
func (s *Sandbox) Annotations(key string) (string, error) {
	return s.MockAnnotations[key], nil
}

// SetAnnotations implements the VCSandbox function of the same name.
func (s *Sandbox) SetAnnotations(annotations map[string]string) error {
	return nil
}

// GetAnnotations implements the VCSandbox function of the same name.
func (s *Sandbox) GetAnnotations() map[string]string {
	return s.MockAnnotations
}

// GetNetNs returns the network namespace of the current sandbox.
func (s *Sandbox) GetNetNs() string {
	return s.MockNetNs
}

// GetAllContainers implements the VCSandbox function of the same name.
func (s *Sandbox) GetAllContainers() []vc.VCContainer {
	var ifa = make([]vc.VCContainer, len(s.MockContainers))

	for i, v := range s.MockContainers {
		ifa[i] = v
	}

	return ifa
}

// GetContainer implements the VCSandbox function of the same name.
func (s *Sandbox) GetContainer(containerID ContainerID) vc.VCContainer {
	for _, c := range s.MockContainers {
		if c.MockID == containerID {
			return c
		}
	}
	return &Container{}
}

// Release implements the VCSandbox function of the same name.
func (s *Sandbox) Release() error {
	return nil
}

// Start implements the VCSandbox function of the same name.
func (s *Sandbox) Start() error {
	return nil
}

// Stop implements the VCSandbox function of the same name.
func (s *Sandbox) Stop() error {
	return nil
}

// Pause implements the VCSandbox function of the same name.
func (s *Sandbox) Pause() error {
	return nil
}

// Resume implements the VCSandbox function of the same name.
func (s *Sandbox) Resume() error {
	return nil
}

// Delete implements the VCSandbox function of the same name.
func (s *Sandbox) Delete() error {
	return nil
}

// CreateContainer implements the VCSandbox function of the same name.
func (s *Sandbox) CreateContainer(conf vc.ContainerConfig) (vc.VCContainer, error) {
	return &Container{}, nil
}

// DeleteContainer implements the VCSandbox function of the same name.
func (s *Sandbox) DeleteContainer(contID ContainerID) (vc.VCContainer, error) {
	return &Container{}, nil
}

// StartContainer implements the VCSandbox function of the same name.
func (s *Sandbox) StartContainer(containerID ContainerID) (vc.VCContainer, error) {
	return &Container{}, nil
}

// StopContainer implements the VCSandbox function of the same name.
func (s *Sandbox) StopContainer(containerID ContainerID) (vc.VCContainer, error) {
	return &Container{}, nil
}

// KillContainer implements the VCSandbox function of the same name.
func (s *Sandbox) KillContainer(contID ContainerID, signal syscall.Signal, all bool) error {
	return nil
}

// StatusContainer implements the VCSandbox function of the same name.
func (s *Sandbox) StatusContainer(contID ContainerID) (vc.ContainerStatus, error) {
	return vc.ContainerStatus{}, nil
}

// StatsContainer implements the VCSandbox function of the same name.
func (s *Sandbox) StatsContainer(contID ContainerID) (vc.ContainerStats, error) {
	return vc.ContainerStats{}, nil
}

// PauseContainer implements the VCSandbox function of the same name.
func (s *Sandbox) PauseContainer(contID ContainerID) error {
	return nil
}

// ResumeContainer implements the VCSandbox function of the same name.
func (s *Sandbox) ResumeContainer(contID ContainerID) error {
	return nil
}

// Status implements the VCSandbox function of the same name.
func (s *Sandbox) Status() vc.SandboxStatus {
	return vc.SandboxStatus{}
}

// EnterContainer implements the VCSandbox function of the same name.
func (s *Sandbox) EnterContainer(containerID ContainerID, cmd types.Cmd) (vc.VCContainer, *vc.Process, error) {
	return &Container{}, &vc.Process{}, nil
}

// Monitor implements the VCSandbox function of the same name.
func (s *Sandbox) Monitor() (chan error, error) {
	return nil, nil
}

// UpdateContainer implements the VCSandbox function of the same name.
func (s *Sandbox) UpdateContainer(containerID ContainerID, resources specs.LinuxResources) error {
	return nil
}

// ProcessListContainer implements the VCSandbox function of the same name.
func (s *Sandbox) ProcessListContainer(containerID ContainerID, options vc.ProcessListOptions) (vc.ProcessList, error) {
	return nil, nil
}

// WaitProcess implements the VCSandbox function of the same name.
func (s *Sandbox) WaitProcess(containerID ContainerID, processID string) (int32, error) {
	return 0, nil
}

// SignalProcess implements the VCSandbox function of the same name.
func (s *Sandbox) SignalProcess(containerID ContainerID, processID string, signal syscall.Signal, all bool) error {
	return nil
}

// WinsizeProcess implements the VCSandbox function of the same name.
func (s *Sandbox) WinsizeProcess(containerID ContainerID, processID string, height, width uint32) error {
	return nil
}

// IOStream implements the VCSandbox function of the same name.
func (s *Sandbox) IOStream(containerID ContainerID, processID string) (io.WriteCloser, io.Reader, io.Reader, error) {
	return nil, nil, nil, nil
}

// AddDevice adds a device to sandbox
func (s *Sandbox) AddDevice(info config.DeviceInfo) (api.Device, error) {
	return nil, nil
}

// AddInterface implements the VCSandbox function of the same name.
func (s *Sandbox) AddInterface(inf *Interface) (*Interface, error) {
	return nil, nil
}

// RemoveInterface implements the VCSandbox function of the same name.
func (s *Sandbox) RemoveInterface(inf *Interface) (*Interface, error) {
	return nil, nil
}

// ListInterfaces implements the VCSandbox function of the same name.
func (s *Sandbox) ListInterfaces() ([]*Interface, error) {
	return nil, nil
}

// UpdateRoutes implements the VCSandbox function of the same name.
func (s *Sandbox) UpdateRoutes(routes []*Route) ([]*Route, error) {
	return nil, nil
}

// ListRoutes implements the VCSandbox function of the same name.
func (s *Sandbox) ListRoutes() ([]*Route, error) {
	return nil, nil
}
