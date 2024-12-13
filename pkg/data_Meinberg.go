// Code generated by go generate; DO NOT EDIT.
package devices

var DeviceTypesMapMeinberg = map[string]*DeviceData{
    "LANTIME M1000": {
        Manufacturer: "Meinberg",
        Model: "LANTIME M1000",
        Slug: "meinberg-lantime-m1000",
        UHeight: 1,
        PartNumber: "",
        IsFullDepth: true,
        Airflow: "passive",
        FrontImage: true,
        RearImage: true,
        SubdeviceRole: "",
        Weight: 0,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
            { Name: "Console", Type: "de-9", Label: "", Poe: false },
            { Name: "USB", Type: "usb-a", Label: "", Poe: false },
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
        },
        PowerOutlets: []PowerOutlet{
        },
        FrontPorts: []FrontPort{
        },
        RearPorts: []RearPort{
        },
        ModuleBays: []ModuleBay{
            { Name: "IO1", Label: "", Position: "1" },
            { Name: "CPU", Label: "", Position: "2" },
            { Name: "IO2", Label: "", Position: "3" },
            { Name: "CLK", Label: "", Position: "4" },
            { Name: "ESI/IO", Label: "", Position: "5" },
            { Name: "MRI/ESI/IO", Label: "", Position: "6" },
            { Name: "PWR1", Label: "", Position: "7" },
            { Name: "PWR2", Label: "", Position: "8" },
        },
			  DeviceBays: []DeviceBay{
        },
        InventoryItems: []InventoryItem{
        },
        Interfaces: []Interface{
        },
    },
    "Lantime M200": {
        Manufacturer: "Meinberg",
        Model: "Lantime M200",
        Slug: "meinberg-lantime-m200",
        UHeight: 1,
        PartNumber: "",
        IsFullDepth: false,
        Airflow: "passive",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 0,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
            { Name: "usb0", Type: "usb-b", Label: "USB0", Poe: false },
            { Name: "terminal", Type: "de-9", Label: "Termnal", Poe: false },
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
            { Name: "PWR1", Label: "", Type: "iec-60320-c14", MaximumDraw: 20, AllocatedDraw: 0 },
        },
        PowerOutlets: []PowerOutlet{
        },
        FrontPorts: []FrontPort{
        },
        RearPorts: []RearPort{
        },
        ModuleBays: []ModuleBay{
        },
			  DeviceBays: []DeviceBay{
        },
        InventoryItems: []InventoryItem{
        },
        Interfaces: []Interface{
            { Name: "eth0", Label: "ETH0", Type: "100base-tx", MgmtOnly: false },
        },
    },
    "Lantime M300": {
        Manufacturer: "Meinberg",
        Model: "Lantime M300",
        Slug: "meinberg-lantime-m300",
        UHeight: 1,
        PartNumber: "",
        IsFullDepth: false,
        Airflow: "passive",
        FrontImage: true,
        RearImage: true,
        SubdeviceRole: "",
        Weight: 3.5,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
            { Name: "serial", Type: "de-9", Label: "Terminal", Poe: false },
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
            { Name: "PWR1", Label: "", Type: "iec-60320-c14", MaximumDraw: 20, AllocatedDraw: 0 },
        },
        PowerOutlets: []PowerOutlet{
        },
        FrontPorts: []FrontPort{
        },
        RearPorts: []RearPort{
        },
        ModuleBays: []ModuleBay{
        },
			  DeviceBays: []DeviceBay{
        },
        InventoryItems: []InventoryItem{
        },
        Interfaces: []Interface{
            { Name: "lan0", Label: "LAN0", Type: "100base-tx", MgmtOnly: false },
            { Name: "lan1", Label: "LAN1", Type: "1000base-tx", MgmtOnly: false },
        },
    },
}