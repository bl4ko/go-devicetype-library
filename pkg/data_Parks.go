// Code generated by go generate; DO NOT EDIT.
package devices

var DeviceTypesMapParks = map[string]*DeviceData{
    "Fiberlink 30028": {
        Manufacturer: "Parks",
        Model: "Fiberlink 30028",
        Slug: "parks-fiberlink-30028",
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
            { Name: "console", Type: "rj-45", Label: "", Poe: false },
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
            { Name: "PSU1", Label: "", Type: "dc-terminal", MaximumDraw: 50, AllocatedDraw: 30 },
            { Name: "PSU2", Label: "", Type: "dc-terminal", MaximumDraw: 50, AllocatedDraw: 30 },
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
            { Name: "giga-ethernet0/0", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "giga-ethernet0/1", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "giga-ethernet0/2", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "giga-ethernet0/3", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "giga-ethernet0/4", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "giga-ethernet0/5", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "giga-ethernet0/6", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "giga-ethernet0/7", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "gpon1/1", Label: "", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "gpon1/2", Label: "", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "gpon1/3", Label: "", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "gpon1/4", Label: "", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "gpon1/5", Label: "", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "gpon1/6", Label: "", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "gpon1/7", Label: "", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "gpon1/8", Label: "", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "10giga-ethernet0/0", Label: "", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "10giga-ethernet0/1", Label: "", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "mgmt", Label: "", Type: "1000base-t", MgmtOnly: true },
        },
    },
}