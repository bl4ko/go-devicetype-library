// Code generated by go generate; DO NOT EDIT.
package devices

var DeviceTypesMapsnr = map[string]*DeviceData{
    "S2995G-24FX": {
        Manufacturer: "SNR",
        Model: "S2995G-24FX",
        Slug: "snr-s2995g-24fx",
        UHeight: 1,
        PartNumber: "S2995G-24FX-UPS",
        IsFullDepth: false,
        Airflow: "right-to-left",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 4.1,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
            { Name: "con0", Type: "rj-45", Label: "Console", Poe: false },
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
            { Name: "PSU1", Label: "", Type: "iec-60320-c14", MaximumDraw: 80, AllocatedDraw: 0 },
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
            { Name: "Ethernet1/0/1", Label: "1", Type: "1000base-tx", MgmtOnly: false },
            { Name: "Ethernet1/0/2", Label: "2", Type: "1000base-tx", MgmtOnly: false },
            { Name: "Ethernet1/0/3", Label: "3", Type: "1000base-tx", MgmtOnly: false },
            { Name: "Ethernet1/0/4", Label: "4", Type: "1000base-tx", MgmtOnly: false },
            { Name: "Ethernet1/0/5", Label: "5", Type: "1000base-tx", MgmtOnly: false },
            { Name: "Ethernet1/0/6", Label: "6", Type: "1000base-tx", MgmtOnly: false },
            { Name: "Ethernet1/0/7", Label: "7", Type: "1000base-tx", MgmtOnly: false },
            { Name: "Ethernet1/0/8", Label: "8", Type: "1000base-tx", MgmtOnly: false },
            { Name: "Ethernet1/0/9", Label: "9", Type: "1000base-x-sfp", MgmtOnly: false },
            { Name: "Ethernet1/0/10", Label: "10", Type: "1000base-x-sfp", MgmtOnly: false },
            { Name: "Ethernet1/0/11", Label: "11", Type: "1000base-x-sfp", MgmtOnly: false },
            { Name: "Ethernet1/0/12", Label: "12", Type: "1000base-x-sfp", MgmtOnly: false },
            { Name: "Ethernet1/0/13", Label: "13", Type: "1000base-x-sfp", MgmtOnly: false },
            { Name: "Ethernet1/0/14", Label: "14", Type: "1000base-x-sfp", MgmtOnly: false },
            { Name: "Ethernet1/0/15", Label: "15", Type: "1000base-x-sfp", MgmtOnly: false },
            { Name: "Ethernet1/0/16", Label: "16", Type: "1000base-x-sfp", MgmtOnly: false },
            { Name: "Ethernet1/0/17", Label: "17", Type: "1000base-x-sfp", MgmtOnly: false },
            { Name: "Ethernet1/0/18", Label: "18", Type: "1000base-x-sfp", MgmtOnly: false },
            { Name: "Ethernet1/0/19", Label: "19", Type: "1000base-x-sfp", MgmtOnly: false },
            { Name: "Ethernet1/0/20", Label: "20", Type: "1000base-x-sfp", MgmtOnly: false },
            { Name: "Ethernet1/0/21", Label: "21", Type: "1000base-x-sfp", MgmtOnly: false },
            { Name: "Ethernet1/0/22", Label: "22", Type: "1000base-x-sfp", MgmtOnly: false },
            { Name: "Ethernet1/0/23", Label: "23", Type: "1000base-x-sfp", MgmtOnly: false },
            { Name: "Ethernet1/0/24", Label: "24", Type: "1000base-x-sfp", MgmtOnly: false },
            { Name: "Ethernet1/0/25", Label: "25", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "Ethernet1/0/26", Label: "26", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "Ethernet1/0/27", Label: "27", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "Ethernet1/0/28", Label: "28", Type: "10gbase-x-sfpp", MgmtOnly: false },
            { Name: "mgmt0", Label: "MGMT", Type: "1000base-tx", MgmtOnly: true },
        },
    },
    "SNR-PDU-08A-W2": {
        Manufacturer: "SNR",
        Model: "SNR-PDU-08A-W2",
        Slug: "snr-pdu-08a-w2",
        UHeight: 1,
        PartNumber: "SNR-PDU-08A-W2",
        IsFullDepth: false,
        Airflow: "",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 0,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
            { Name: "Power Port 1", Label: "", Type: "iec-60320-c14", MaximumDraw: 0, AllocatedDraw: 0 },
        },
        PowerOutlets: []PowerOutlet{
            { Name: "Outlet 1", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 2", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 3", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 4", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 5", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 6", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 7", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 8", Type: "iec-60320-c13", Label: "", PowerPort: "Power Port 1", FeedLeg: "", MaximumDraw: 0, AllocatedDraw: 0 },
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
        },
    },
    "SNR-PDU-08S-W2": {
        Manufacturer: "SNR",
        Model: "SNR-PDU-08S-W2",
        Slug: "snr-pdu-08s-w2",
        UHeight: 1,
        PartNumber: "SNR-PDU-08S-W2",
        IsFullDepth: false,
        Airflow: "",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 0,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
            { Name: "Power Port 1", Label: "", Type: "iec-60320-c14", MaximumDraw: 0, AllocatedDraw: 0 },
        },
        PowerOutlets: []PowerOutlet{
            { Name: "Outlet 1", Type: "ita-f", Label: "", PowerPort: "Power Port 1", FeedLeg: "", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 2", Type: "ita-f", Label: "", PowerPort: "Power Port 1", FeedLeg: "", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 3", Type: "ita-f", Label: "", PowerPort: "Power Port 1", FeedLeg: "", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 4", Type: "ita-f", Label: "", PowerPort: "Power Port 1", FeedLeg: "", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 5", Type: "ita-f", Label: "", PowerPort: "Power Port 1", FeedLeg: "", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 6", Type: "ita-f", Label: "", PowerPort: "Power Port 1", FeedLeg: "", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 7", Type: "ita-f", Label: "", PowerPort: "Power Port 1", FeedLeg: "", MaximumDraw: 0, AllocatedDraw: 0 },
            { Name: "Outlet 8", Type: "ita-f", Label: "", PowerPort: "Power Port 1", FeedLeg: "", MaximumDraw: 0, AllocatedDraw: 0 },
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
        },
    },
}
