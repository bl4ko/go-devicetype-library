// Code generated by go generate; DO NOT EDIT.
package devices

var DeviceTypesMapcryptopro = map[string]*DeviceData{
    "NGATE-320": {
        Manufacturer: "CryptoPro",
        Model: "NGATE-320",
        Slug: "cryptopro-ngate-320",
        UHeight: 1,
        PartNumber: "",
        IsFullDepth: false,
        Airflow: "",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 0,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
            { Name: "ttyS0", Type: "rj-45", Label: "CONSOLE", Poe: false },
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
            { Name: "PSU0", Label: "", Type: "dc-terminal", MaximumDraw: 36, AllocatedDraw: 0 },
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
            { Name: "enp1s0", Label: "1", Type: "1000base-t", MgmtOnly: false },
            { Name: "enp2s0", Label: "2", Type: "1000base-t", MgmtOnly: false },
            { Name: "enp3s0", Label: "3", Type: "1000base-t", MgmtOnly: false },
        },
    },
    "NGATE-600": {
        Manufacturer: "CryptoPro",
        Model: "NGATE-600",
        Slug: "cryptopro-ngate-600",
        UHeight: 1,
        PartNumber: "",
        IsFullDepth: true,
        Airflow: "",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 0,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
            { Name: "ttyS0", Type: "rj-45", Label: "CONSOLE", Poe: false },
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
            { Name: "PSU0", Label: "", Type: "iec-60320-c14", MaximumDraw: 150, AllocatedDraw: 0 },
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
            { Name: "enp0f20s0", Label: "1", Type: "1000base-t", MgmtOnly: false },
            { Name: "enp0f20s1", Label: "2", Type: "1000base-t", MgmtOnly: false },
            { Name: "enp0f20s2", Label: "3", Type: "1000base-t", MgmtOnly: false },
            { Name: "enp0f20s3", Label: "4", Type: "1000base-t", MgmtOnly: false },
            { Name: "enp1s0", Label: "5", Type: "1000base-t", MgmtOnly: false },
            { Name: "enp2s0", Label: "6", Type: "1000base-t", MgmtOnly: false },
        },
    },
}
