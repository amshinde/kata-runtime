// Copyright (c) 2016 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0
//

package virtcontainers

type mockHypervisor struct {
}

func (m *mockHypervisor) init(sandbox *Sandbox) error {
	valid, err := sandbox.config.HypervisorConfig.valid()
	if valid == false || err != nil {
		return err
	}

	return nil
}

func (m *mockHypervisor) capabilities() capabilities {
	return capabilities{}
}

func (m *mockHypervisor) createSandbox(sandboxConfig SandboxConfig) error {
	return nil
}

func (m *mockHypervisor) startSandbox() error {
	return nil
}

func (m *mockHypervisor) waitSandbox(timeout int) error {
	return nil
}

func (m *mockHypervisor) stopSandbox() error {
	return nil
}

func (m *mockHypervisor) pauseSandbox() error {
	return nil
}

func (m *mockHypervisor) resumeSandbox() error {
	return nil
}

func (m *mockHypervisor) addDevice(devInfo interface{}, devType deviceType) error {
	return nil
}

func (m *mockHypervisor) hotplugAddDevice(devInfo interface{}, devType deviceType) error {
	return nil
}

func (m *mockHypervisor) hotplugRemoveDevice(devInfo interface{}, devType deviceType) error {
	return nil
}

func (m *mockHypervisor) getSandboxConsole(sandboxID string) string {
	return ""
}
