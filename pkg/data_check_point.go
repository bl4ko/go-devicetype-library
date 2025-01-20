// Code generated by go generate; DO NOT EDIT.
package devices

var DeviceTypesMapcheck_point = map[string]*DeviceData{
    "CPAC-1500/3600/3800-RM-DUAL": {
        Manufacturer: "Check Point",
        Model: "CPAC-1500/3600/3800-RM-DUAL",
        Slug: "check-point-cpac-1500-3600-3800-rm-dual",
        UHeight: 1,
        PartNumber: "CPAC-1500/3600/3800-RM-DUAL",
        IsFullDepth: false,
        Airflow: "",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "parent",
        Weight: 0,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
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
        },
			  DeviceBays: []DeviceBay{
            { Name: "Left", Label: "" },
            { Name: "Right", Label: "" },
        },
        InventoryItems: []InventoryItem{
        },
        Interfaces: []Interface{
        },
    },
    "MHO-140": {
        Manufacturer: "Check Point",
        Model: "MHO-140",
        Slug: "check-point-mho-140",
        UHeight: 1,
        PartNumber: "CPAP-MHO-140",
        IsFullDepth: true,
        Airflow: "",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 8.52,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
            { Name: "CONSOLE", Type: "rj-45", Label: "", Poe: false },
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
            { Name: "Power1", Label: "", Position: "1" },
            { Name: "Power2", Label: "", Position: "2" },
        },
			  DeviceBays: []DeviceBay{
        },
        InventoryItems: []InventoryItem{
        },
        Interfaces: []Interface{
            { Name: "0", Label: "", Type: "1000base-t", MgmtOnly: true },
            { Name: "1", Label: "", Type: "1000base-t", MgmtOnly: true },
            { Name: "sfp1", Label: "", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "sfp2", Label: "", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "sfp3", Label: "", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "sfp4", Label: "", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "sfp5", Label: "", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "sfp6", Label: "", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "sfp7", Label: "", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "sfp8", Label: "", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "sfp9", Label: "", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "sfp10", Label: "", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "sfp11", Label: "", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "sfp12", Label: "", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "sfp13", Label: "", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "sfp14", Label: "", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "sfp15", Label: "", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "sfp16", Label: "", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "sfp17", Label: "", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "sfp18", Label: "", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "sfp19", Label: "", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "sfp20", Label: "", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "sfp21", Label: "", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "sfp22", Label: "", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "sfp23", Label: "", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "sfp24", Label: "", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "sfp25", Label: "", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "sfp26", Label: "", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "sfp27", Label: "", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "sfp28", Label: "", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "sfp29", Label: "", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "sfp30", Label: "", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "sfp31", Label: "", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "sfp32", Label: "", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "sfp33", Label: "", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "sfp34", Label: "", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "sfp35", Label: "", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "sfp36", Label: "", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "sfp37", Label: "", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "sfp38", Label: "", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "sfp39", Label: "", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "sfp40", Label: "", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "sfp41", Label: "", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "sfp42", Label: "", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "sfp43", Label: "", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "sfp44", Label: "", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "sfp45", Label: "", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "sfp46", Label: "", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "sfp47", Label: "", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "sfp48", Label: "", Type: "25gbase-x-sfp28", MgmtOnly: false },
            { Name: "qsfp49", Label: "", Type: "100gbase-x-qsfp28", MgmtOnly: false },
            { Name: "qsfp50", Label: "", Type: "100gbase-x-qsfp28", MgmtOnly: false },
            { Name: "qsfp51", Label: "", Type: "100gbase-x-qsfp28", MgmtOnly: false },
            { Name: "qsfp52", Label: "", Type: "100gbase-x-qsfp28", MgmtOnly: false },
            { Name: "qsfp53", Label: "", Type: "100gbase-x-qsfp28", MgmtOnly: false },
            { Name: "qsfp54", Label: "", Type: "100gbase-x-qsfp28", MgmtOnly: false },
            { Name: "qsfp55", Label: "", Type: "100gbase-x-qsfp28", MgmtOnly: false },
            { Name: "qsfp56", Label: "", Type: "100gbase-x-qsfp28", MgmtOnly: false },
        },
    },
    "SG1590": {
        Manufacturer: "Check Point",
        Model: "SG1590",
        Slug: "check-point-sg1590",
        UHeight: 1,
        PartNumber: "CPAP-SG1590",
        IsFullDepth: false,
        Airflow: "passive",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 0.87,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
            { Name: "CONSOLE", Type: "usb-c", Label: "", Poe: false },
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
            { Name: "DC", Label: "", Type: "dc-terminal", MaximumDraw: 0, AllocatedDraw: 0 },
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
            { Name: "DMZ-RJ-45", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "DMZ-SFP", Label: "", Type: "1000base-x-sfp", MgmtOnly: false },
            { Name: "1", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "2", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "3", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "4", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "5", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "6", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "7", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "8", Label: "", Type: "1000base-t", MgmtOnly: false },
            { Name: "WAN", Label: "", Type: "1000base-t", MgmtOnly: false },
        },
    },
    "SG23800": {
        Manufacturer: "Check Point",
        Model: "SG23800",
        Slug: "check-point-sg23800",
        UHeight: 2,
        PartNumber: "CPAP-SG23800",
        IsFullDepth: true,
        Airflow: "",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 0,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
            { Name: "Console", Type: "rj-45", Label: "", Poe: false },
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
            { Name: "PSU0", Label: "", Type: "iec-60320-c14", MaximumDraw: 394, AllocatedDraw: 0 },
            { Name: "PSU1", Label: "", Type: "iec-60320-c14", MaximumDraw: 394, AllocatedDraw: 0 },
        },
        PowerOutlets: []PowerOutlet{
        },
        FrontPorts: []FrontPort{
        },
        RearPorts: []RearPort{
        },
        ModuleBays: []ModuleBay{
            { Name: "1", Label: "", Position: "1" },
            { Name: "2", Label: "", Position: "2" },
            { Name: "3", Label: "", Position: "3" },
            { Name: "4", Label: "", Position: "4" },
            { Name: "5", Label: "", Position: "5" },
        },
			  DeviceBays: []DeviceBay{
        },
        InventoryItems: []InventoryItem{
        },
        Interfaces: []Interface{
            { Name: "LOM", Label: "", Type: "1000base-t", MgmtOnly: true },
            { Name: "MGMT", Label: "", Type: "1000base-t", MgmtOnly: true },
            { Name: "SYNC", Label: "", Type: "1000base-t", MgmtOnly: false },
        },
    },
    "SG3600": {
        Manufacturer: "Check Point",
        Model: "SG3600",
        Slug: "check-point-sg3600",
        UHeight: 0,
        PartNumber: "CPAP-SG3600-SNBT",
        IsFullDepth: false,
        Airflow: "",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "child",
        Weight: 0,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
            { Name: "ttyS0", Type: "rj-45", Label: "CONSOLE", Poe: false },
            { Name: "ttyS1", Type: "usb-c", Label: "CONSOLE", Poe: false },
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
            { Name: "PS1", Label: "", Type: "dc-terminal", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "PS2", Label: "", Type: "dc-terminal", MaximumDraw: 0, AllocatedDraw: 0 },
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
            { Name: "eth1", Label: "1", Type: "1000base-t", MgmtOnly: false },
            { Name: "eth2", Label: "2", Type: "1000base-t", MgmtOnly: false },
            { Name: "eth3", Label: "3", Type: "1000base-t", MgmtOnly: false },
            { Name: "eth4", Label: "4", Type: "1000base-t", MgmtOnly: false },
            { Name: "eth5", Label: "5", Type: "1000base-t", MgmtOnly: false },
            { Name: "Mgmt", Label: "MGMT", Type: "1000base-t", MgmtOnly: false },
            { Name: "lom", Label: "LOM", Type: "1000base-t", MgmtOnly: true },
        },
    },
    "SG6200": {
        Manufacturer: "Check Point",
        Model: "SG6200",
        Slug: "check-point-sg6200",
        UHeight: 1,
        PartNumber: "CPAP-SG6200-SNBT",
        IsFullDepth: true,
        Airflow: "",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 6.75,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
            { Name: "ttyS0", Type: "rj-45", Label: "CONSOLE", Poe: false },
            { Name: "ttyS1", Type: "usb-c", Label: "CONSOLE", Poe: false },
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
            { Name: "PS1", Label: "", Type: "iec-60320-c14", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "PS2", Label: "", Type: "iec-60320-c14", MaximumDraw: 0, AllocatedDraw: 0 },
        },
        PowerOutlets: []PowerOutlet{
        },
        FrontPorts: []FrontPort{
        },
        RearPorts: []RearPort{
        },
        ModuleBays: []ModuleBay{
            { Name: "Line card 1", Label: "", Position: "" },
        },
			  DeviceBays: []DeviceBay{
        },
        InventoryItems: []InventoryItem{
        },
        Interfaces: []Interface{
            { Name: "eth1", Label: "1", Type: "1000base-t", MgmtOnly: false },
            { Name: "eth2", Label: "2", Type: "1000base-t", MgmtOnly: false },
            { Name: "eth3", Label: "3", Type: "1000base-t", MgmtOnly: false },
            { Name: "eth4", Label: "4", Type: "1000base-t", MgmtOnly: false },
            { Name: "eth5", Label: "5", Type: "1000base-t", MgmtOnly: false },
            { Name: "eth6", Label: "6", Type: "1000base-t", MgmtOnly: false },
            { Name: "eth7", Label: "7", Type: "1000base-t", MgmtOnly: false },
            { Name: "eth8", Label: "8", Type: "1000base-t", MgmtOnly: false },
            { Name: "Sync", Label: "SYNC", Type: "1000base-t", MgmtOnly: false },
            { Name: "Mgmt", Label: "MGMT", Type: "1000base-t", MgmtOnly: false },
            { Name: "lom", Label: "LOM", Type: "1000base-t", MgmtOnly: true },
        },
    },
    "SG6400": {
        Manufacturer: "Check Point",
        Model: "SG6400",
        Slug: "check-point-sg6400",
        UHeight: 1,
        PartNumber: "CPAP-SG6400-SNBT",
        IsFullDepth: true,
        Airflow: "",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "parent",
        Weight: 0,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
            { Name: "ttyS0", Type: "rj-45", Label: "CONSOLE", Poe: false },
            { Name: "ttyS1", Type: "usb-c", Label: "CONSOLE", Poe: false },
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
            { Name: "PS1", Label: "", Type: "iec-60320-c14", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "PS2", Label: "", Type: "iec-60320-c14", MaximumDraw: 0, AllocatedDraw: 0 },
        },
        PowerOutlets: []PowerOutlet{
        },
        FrontPorts: []FrontPort{
        },
        RearPorts: []RearPort{
        },
        ModuleBays: []ModuleBay{
            { Name: "Line card 1", Label: "", Position: "" },
        },
			  DeviceBays: []DeviceBay{
        },
        InventoryItems: []InventoryItem{
        },
        Interfaces: []Interface{
            { Name: "eth1", Label: "1", Type: "1000base-t", MgmtOnly: false },
            { Name: "eth2", Label: "2", Type: "1000base-t", MgmtOnly: false },
            { Name: "eth3", Label: "3", Type: "1000base-t", MgmtOnly: false },
            { Name: "eth4", Label: "4", Type: "1000base-t", MgmtOnly: false },
            { Name: "eth5", Label: "5", Type: "1000base-t", MgmtOnly: false },
            { Name: "eth6", Label: "6", Type: "1000base-t", MgmtOnly: false },
            { Name: "eth7", Label: "7", Type: "1000base-t", MgmtOnly: false },
            { Name: "eth8", Label: "8", Type: "1000base-t", MgmtOnly: false },
            { Name: "Sync", Label: "SYNC", Type: "1000base-t", MgmtOnly: false },
            { Name: "Mgmt", Label: "MGMT", Type: "1000base-t", MgmtOnly: false },
            { Name: "lom", Label: "LOM", Type: "1000base-t", MgmtOnly: true },
        },
    },
}
