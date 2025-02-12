// Code generated by go generate; DO NOT EDIT.
package devices

var DeviceTypesMapdepo = map[string]*DeviceData{
    "DEPO-Storage-1304": {
        Manufacturer: "DEPO",
        Model: "DEPO-Storage-1304",
        Slug: "depo-storage-1304",
        UHeight: 1,
        PartNumber: "",
        IsFullDepth: true,
        Airflow: "front-to-rear",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 10.9,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
            { Name: "PSU1", Label: "", Type: "iec-60320-c14", MaximumDraw: 250, AllocatedDraw: 100 },
            { Name: "PSU2", Label: "", Type: "iec-60320-c14", MaximumDraw: 250, AllocatedDraw: 100 },
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
            { Name: "Eth1", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "Eth2", Label: "", Type: "1000base-t", MgmtOnly: false },
        },
    },
}
