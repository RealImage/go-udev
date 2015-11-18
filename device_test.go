// +build linux

package udev

import (
	"fmt"
	"runtime"
	"testing"
)

func ExampleDevice() {

	// Create Udev
	u := Udev{}

	// Create new Device based on subsystem and sysname
	d := u.NewDeviceFromSubsystemSysname("mem", "zero")

	// Extract information
	fmt.Printf("Syspath:%v\n", d.Syspath())
	fmt.Printf("Devpath:%v\n", d.Devpath())
	fmt.Printf("Devnode:%v\n", d.Devnode())
	fmt.Printf("Subsystem:%v\n", d.Subsystem())
	fmt.Printf("Devtype:%v\n", d.Devtype())
	fmt.Printf("Sysnum:%v\n", d.Sysnum())
	fmt.Printf("IsInitialized:%v\n", d.IsInitialized())
	if s, e := d.Driver(); e != nil {
		fmt.Printf("Driver:%v\n", s)
	}

	// Output:
	// Syspath:/sys/devices/virtual/mem/zero
	// Devpath:/devices/virtual/mem/zero
	// Devnode:/dev/zero
	// Subsystem:mem
	// Devtype:
	// Sysnum:
	// IsInitialized:true
	// Driver:
}

func TestDeviceZero(t *testing.T) {
	u := Udev{}
	d := u.NewDeviceFromDeviceID("c1:5")
	if d.Subsystem() != "mem" {
		t.Fail()
	}
	if d.Syspath() != "/sys/devices/virtual/mem/zero" {
		t.Fail()
	}
	if d.Devnode() != "/dev/zero" {
		t.Fail()
	}
	p, e := d.PropertyValue("SUBSYSTEM")
	if e != nil || p != "mem" {
		t.Fail()
	}
	if !d.IsInitialized() {
		t.Fail()
	}
	s, e := d.SysattrValue("subsystem")
	if e != nil || s != "mem" {
		t.Fail()
	}
	// Device should have Properties
	properties := d.Properties()
	if len(properties) == 0 {
		t.Fail()
	}
	// Device should have Sysattrs
	sysattrs := d.Sysattrs()
	if len(sysattrs) == 0 {
		t.Fail()
	}
}

func TestDeviceRandom(t *testing.T) {
	u := Udev{}
	d := u.NewDeviceFromDeviceID("c1:8")
	if d.Subsystem() != "mem" {
		t.Fail()
	}
	if d.Syspath() != "/sys/devices/virtual/mem/random" {
		t.Fail()
	}
	if d.Devnode() != "/dev/random" {
		t.Fail()
	}
	p, e := d.PropertyValue("SUBSYSTEM")
	if e != nil || p != "mem" {
		t.Fail()
	}
	if !d.IsInitialized() {
		t.Fail()
	}
	s, e := d.SysattrValue("subsystem")
	if e != nil || s != "mem" {
		t.Fail()
	}
	// Device should have Properties
	properties := d.Properties()
	if len(properties) == 0 {
		t.Fail()
	}
	// Device should have Sysattrs
	sysattrs := d.Sysattrs()
	if len(sysattrs) == 0 {
		t.Fail()
	}
}

func TestDeviceGC(t *testing.T) {
	runtime.GC()
}
