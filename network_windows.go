//go:build windows

package main

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

type WindowsXion struct {
	devices []Device
}

func NewXion() Xion {
	return &WindowsXion{
		devices: getDevices(),
	}
}

const mibFilterInterface = 0x02

type Device struct {
	LUID uint64
	Name string
}

func getDevices() []Device {
	var devices []Device

	var table *windows.MibIfTable2
	err := windows.GetIfTable2Ex(windows.MibIfTableNormal, &table)
	if err != nil {
		panic("Failed to get the table")
	}
	defer windows.FreeMibTable(unsafe.Pointer(table))

	rows := unsafe.Slice(&table.Table[0], table.NumEntries)

	for _, row := range rows {

		if row.InterfaceAndOperStatusFlags&mibFilterInterface != 0 {
			continue
		}

		if row.InOctets == 0 && row.OutOctets == 0 {
			continue
		}

		if row.OperStatus != windows.IfOperStatusUp {
			continue
		}

		if row.Type == windows.IF_TYPE_ETHERNET_CSMACD || row.Type == windows.IF_TYPE_IEEE80211 {
			devices = append(devices, Device{
				LUID: row.InterfaceLuid,
				Name: windows.UTF16ToString(row.Description[:]),
			})
		}
	}

	return devices
}

func (w *WindowsXion) GetData() []uint64 {
	row := windows.MibIfRow2{
		InterfaceLuid: w.devices[0].LUID,
	}

	err := windows.GetIfEntry2Ex(windows.MibIfEntryNormal, &row)
	if err != nil {
		panic("There is a issue")
	}

	return []uint64{row.OutOctets, row.InOctets}
}
