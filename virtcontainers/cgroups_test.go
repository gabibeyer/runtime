// Copyright (c) 2018 Huawei Corporation
//
// SPDX-License-Identifier: Apache-2.0
//

package virtcontainers

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/containerd/cgroups"
	. "github.com/kata-containers/runtime/virtcontainers/pkg/types"
	"github.com/kata-containers/runtime/virtcontainers/types"
	specs "github.com/opencontainers/runtime-spec/specs-go"
	"github.com/stretchr/testify/assert"
)

type mockCgroup struct {
}

func (m *mockCgroup) New(string, *specs.LinuxResources) (cgroups.Cgroup, error) {
	return &mockCgroup{}, nil
}
func (m *mockCgroup) Add(cgroups.Process) error {
	return nil
}

func (m *mockCgroup) AddTask(cgroups.Process) error {
	return nil
}

func (m *mockCgroup) Delete() error {
	return nil
}

func (m *mockCgroup) MoveTo(cgroups.Cgroup) error {
	return nil
}

func (m *mockCgroup) Stat(...cgroups.ErrorHandler) (*cgroups.Metrics, error) {
	return &cgroups.Metrics{}, nil
}

func (m *mockCgroup) Update(resources *specs.LinuxResources) error {
	return nil
}

func (m *mockCgroup) Processes(cgroups.Name, bool) ([]cgroups.Process, error) {
	return nil, nil
}

func (m *mockCgroup) Freeze() error {
	return nil
}

func (m *mockCgroup) Thaw() error {
	return nil
}

func (m *mockCgroup) OOMEventFD() (uintptr, error) {
	return 0, nil
}

func (m *mockCgroup) State() cgroups.State {
	return ""
}

func (m *mockCgroup) Subsystems() []cgroups.Subsystem {
	return nil
}

func mockCgroupNew(hierarchy cgroups.Hierarchy, path cgroups.Path, resources *specs.LinuxResources) (cgroups.Cgroup, error) {
	return &mockCgroup{}, nil
}

func mockCgroupLoad(hierarchy cgroups.Hierarchy, path cgroups.Path) (cgroups.Cgroup, error) {
	return &mockCgroup{}, nil
}

func init() {
	cgroupsNewFunc = mockCgroupNew
	cgroupsLoadFunc = mockCgroupLoad
}

func TestV1Constraints(t *testing.T) {
	assert := assert.New(t)

	systems, err := V1Constraints()
	assert.NoError(err)
	assert.NotEmpty(systems)
}

func TestV1NoConstraints(t *testing.T) {
	assert := assert.New(t)

	systems, err := V1NoConstraints()
	assert.NoError(err)
	assert.NotEmpty(systems)
}

func TestCgroupNoConstraintsPath(t *testing.T) {
	assert := assert.New(t)

	cgrouPath := "abc"
	expectedPath := filepath.Join(cgroupKataPath, cgrouPath)
	path := cgroupNoConstraintsPath(cgrouPath)
	assert.Equal(expectedPath, path)
}

func TestUpdateCgroups(t *testing.T) {
	assert := assert.New(t)

	oldCgroupsNew := cgroupsNewFunc
	oldCgroupsLoad := cgroupsLoadFunc
	cgroupsNewFunc = cgroups.New
	cgroupsLoadFunc = cgroups.Load
	defer func() {
		cgroupsNewFunc = oldCgroupsNew
		cgroupsLoadFunc = oldCgroupsLoad
	}()

	s := &Sandbox{
		state: types.State{
			CgroupPath: "",
		},
	}

	// empty path
	err := s.updateCgroups()
	assert.NoError(err)

	// path doesn't exist
	s.state.CgroupPath = "/abc/123/rgb"
	err = s.updateCgroups()
	assert.Error(err)

	if os.Getuid() != 0 {
		return
	}

	s.state.CgroupPath = fmt.Sprintf("/kata-tests-%d", os.Getpid())
	testCgroup, err := cgroups.New(cgroups.V1, cgroups.StaticPath(s.state.CgroupPath), &specs.LinuxResources{})
	assert.NoError(err)
	defer testCgroup.Delete()
	s.hypervisor = &mockHypervisor{mockPid: 0}

	// bad pid
	err = s.updateCgroups()
	assert.Error(err)

	// fake workload
	cmd := exec.Command("tail", "-f", "/dev/null")
	assert.NoError(cmd.Start())
	s.state.Pid = cmd.Process.Pid
	s.hypervisor = &mockHypervisor{mockPid: s.state.Pid}

	// no containers
	err = s.updateCgroups()
	assert.NoError(err)

	s.config = &SandboxConfig{}
	s.config.HypervisorConfig.NumVCPUs = 1

	s.containers = map[ContainerID]*Container{
		"abc": {
			process: Process{
				Pid: s.state.Pid,
			},
			config: &ContainerConfig{
				Annotations: containerAnnotations,
			},
		},
		"xyz": {
			process: Process{
				Pid: s.state.Pid,
			},
			config: &ContainerConfig{
				Annotations: containerAnnotations,
			},
		},
	}

	err = s.updateCgroups()
	assert.NoError(err)

	// cleanup
	assert.NoError(cmd.Process.Kill())
	err = s.deleteCgroups()
	assert.NoError(err)
}
