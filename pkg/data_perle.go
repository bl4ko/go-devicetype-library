// Code generated by go generate; DO NOT EDIT.
package devices

var DeviceTypesMapperle = map[string]*DeviceData{
    "IOLAN STS4-D": {
        Manufacturer: "Perle",
        Model: "IOLAN STS4-D",
        Slug: "perle-iolan-sts4-d",
        UHeight: 1,
        PartNumber: "STS4-D",
        IsFullDepth: false,
        Airflow: "passive",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 0.77,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
        },
        ConsoleServerPorts: []ConsoleServerPort{
            { Name: "RS232-1", Type: "rj-45", Label: "RS232-1" },
            { Name: "RS232-2", Type: "rj-45", Label: "RS232-2" },
            { Name: "RS232-3", Type: "rj-45", Label: "RS232-3" },
            { Name: "RS232-4", Type: "rj-45", Label: "RS232-4" },
        },
        PowerPorts: []PowerPort{
            { Name: "DC in", Label: "DC in", Type: "dc-terminal", MaximumDraw: 5, AllocatedDraw: 0 },
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
            { Name: "Interface 1", Label: "Interface 1", Type: "100base-tx", MgmtOnly: false },
        },
    },
}
