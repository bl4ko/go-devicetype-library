// Code generated by go generate; DO NOT EDIT.
package devices

var DeviceTypesMapschweitzer_engineering_laboratories = map[string]*DeviceData{
    "SEL-2242 10-bay Chassis/Backplane": {
        Manufacturer: "Schweitzer Engineering Laboratories",
        Model: "SEL-2242 10-bay Chassis/Backplane",
        Slug: "schweitzer-engineering-laboratories-sel-2242-10-bay-chassis-backplane",
        UHeight: 2,
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
            { Name: "Slot A", Label: "", Position: "" },
            { Name: "Slot B", Label: "", Position: "" },
            { Name: "Slot C", Label: "", Position: "" },
            { Name: "Slot D", Label: "", Position: "" },
            { Name: "Slot E", Label: "", Position: "" },
            { Name: "Slot F", Label: "", Position: "" },
            { Name: "Slot G", Label: "", Position: "" },
            { Name: "Slot H", Label: "", Position: "" },
            { Name: "Slot I", Label: "", Position: "" },
            { Name: "Slot J", Label: "", Position: "" },
        },
			  DeviceBays: []DeviceBay{
        },
        InventoryItems: []InventoryItem{
        },
        Interfaces: []Interface{
        },
    },
    "SEL-2242 4-bay Chassis/Backplane": {
        Manufacturer: "Schweitzer Engineering Laboratories",
        Model: "SEL-2242 4-bay Chassis/Backplane",
        Slug: "schweitzer-engineering-laboratories-sel-2242-4-bay-chassis-backplane",
        UHeight: 2,
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
            { Name: "Slot A", Label: "", Position: "" },
            { Name: "Slot B", Label: "", Position: "" },
            { Name: "Slot C", Label: "", Position: "" },
            { Name: "Slot D", Label: "", Position: "" },
        },
			  DeviceBays: []DeviceBay{
        },
        InventoryItems: []InventoryItem{
        },
        Interfaces: []Interface{
        },
    },
    "SEL-351 Protection System": {
        Manufacturer: "Schweitzer Engineering Laboratories",
        Model: "SEL-351 Protection System",
        Slug: "schweitzer-engineering-laboratories-sel-351-protection-system",
        UHeight: 2,
        PartNumber: "SEL-351",
        IsFullDepth: true,
        Airflow: "front-to-rear",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 11,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
            { Name: "Port F", Type: "other", Label: "Console Port", Poe: false },
            { Name: "USB 1", Type: "usb-b", Label: "USB Console Port", Poe: false },
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
            { Name: "Power", Label: "", Type: "dc-terminal", MaximumDraw: 25, AllocatedDraw: 0 },
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
            { Name: "IRIG-B", Label: "NTP, SNTP Time Synchronation", Type: "other", MgmtOnly: false },
            { Name: "5A", Label: "", Type: "100base-tx", MgmtOnly: false },
            { Name: "5B", Label: "", Type: "100base-tx", MgmtOnly: false },
            { Name: "Serial Port 1", Label: "EIA-485 Engineering Compression Terminal", Type: "other", MgmtOnly: false },
            { Name: "Serial Port 2", Label: "ASCII, TCP, Modbus RTU, DNP3, Telnet, Webserver, and IEC 61850", Type: "other", MgmtOnly: false },
            { Name: "Serial Port 3", Label: "SEL Dual-channel MIRRORED BITS", Type: "other", MgmtOnly: false },
        },
    },
    "SEL-451-5 Relay": {
        Manufacturer: "Schweitzer Engineering Laboratories",
        Model: "SEL-451-5 Relay",
        Slug: "schweitzer-engineering-laboratories-sel-451-5-relay",
        UHeight: 2,
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
            { Name: "EIA-232-1", Type: "de-9", Label: "", Poe: false },
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
            { Name: "power", Label: "", Type: "hardwired", MaximumDraw: 0, AllocatedDraw: 0 },
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
            { Name: "EIA-232-2", Label: "", Type: "other", MgmtOnly: false },
            { Name: "EIA-232-3", Label: "", Type: "other", MgmtOnly: false },
            { Name: "EIA-232-4", Label: "", Type: "other", MgmtOnly: false },
            { Name: "Ethernet0/0", Label: "", Type: "100base-fx", MgmtOnly: false },
            { Name: "Ethernet0/1", Label: "", Type: "100base-fx", MgmtOnly: false },
        },
    },
    "SEL-451-6 SV": {
        Manufacturer: "Schweitzer Engineering Laboratories",
        Model: "SEL-451-6 SV",
        Slug: "schweitzer-engineering-laboratories-sel-451-6-sv",
        UHeight: 4,
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
            { Name: "EIA-232-1", Type: "de-9", Label: "", Poe: false },
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
            { Name: "power", Label: "", Type: "hardwired", MaximumDraw: 0, AllocatedDraw: 0 },
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
            { Name: "EIA-232-2", Label: "", Type: "other", MgmtOnly: false },
            { Name: "EIA-232-3", Label: "", Type: "other", MgmtOnly: false },
            { Name: "EIA-232-4", Label: "", Type: "other", MgmtOnly: false },
            { Name: "Ethernet0/0", Label: "", Type: "1000base-x-sfp", MgmtOnly: false },
            { Name: "Ethernet0/1", Label: "", Type: "1000base-x-sfp", MgmtOnly: false },
            { Name: "Ethernet1/0", Label: "", Type: "100base-fx", MgmtOnly: false },
            { Name: "Ethernet1/1", Label: "", Type: "100base-fx", MgmtOnly: false },
            { Name: "Ethernet1/2", Label: "", Type: "100base-fx", MgmtOnly: false },
        },
    },
    "SEL-451-6 TiDL": {
        Manufacturer: "Schweitzer Engineering Laboratories",
        Model: "SEL-451-6 TiDL",
        Slug: "schweitzer-engineering-laboratories-sel-451-6-tidl",
        UHeight: 4,
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
            { Name: "EIA-232-1", Type: "de-9", Label: "", Poe: false },
        },
        ConsoleServerPorts: []ConsoleServerPort{
        },
        PowerPorts: []PowerPort{
            { Name: "power", Label: "", Type: "hardwired", MaximumDraw: 0, AllocatedDraw: 0 },
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
            { Name: "EIA-232-2", Label: "", Type: "other", MgmtOnly: false },
            { Name: "EIA-232-3", Label: "", Type: "other", MgmtOnly: false },
            { Name: "EIA-232-4", Label: "", Type: "other", MgmtOnly: false },
            { Name: "Ethernet0/0", Label: "", Type: "1000base-x-sfp", MgmtOnly: false },
            { Name: "Ethernet0/1", Label: "", Type: "1000base-x-sfp", MgmtOnly: false },
            { Name: "Ethernet1/0", Label: "", Type: "100base-fx", MgmtOnly: false },
            { Name: "Ethernet1/1", Label: "", Type: "100base-fx", MgmtOnly: false },
            { Name: "Ethernet1/2", Label: "", Type: "100base-fx", MgmtOnly: false },
            { Name: "TiDL0/0", Label: "", Type: "other", MgmtOnly: false },
            { Name: "TiDL0/1", Label: "", Type: "other", MgmtOnly: false },
            { Name: "TiDL0/2", Label: "", Type: "other", MgmtOnly: false },
            { Name: "TiDL0/3", Label: "", Type: "other", MgmtOnly: false },
            { Name: "TiDL0/4", Label: "", Type: "other", MgmtOnly: false },
            { Name: "TiDL0/5", Label: "", Type: "other", MgmtOnly: false },
            { Name: "TiDL0/6", Label: "", Type: "other", MgmtOnly: false },
            { Name: "TiDL0/7", Label: "", Type: "other", MgmtOnly: false },
        },
    },
    "SEL-587 Current Differential Relay": {
        Manufacturer: "Schweitzer Engineering Laboratories",
        Model: "SEL-587 Current Differential Relay",
        Slug: "schweitzer-engineering-laboratories-sel-587-current-differential-relay",
        UHeight: 2,
        PartNumber: "SEL-587",
        IsFullDepth: true,
        Airflow: "front-to-rear",
        FrontImage: false,
        RearImage: false,
        SubdeviceRole: "",
        Weight: 5,
        WeightUnit: "",
        IsPowered: false,
        ConsolePorts: []ConsolePort{
        },
        ConsoleServerPorts: []ConsoleServerPort{
            { Name: "Port F", Type: "de-9", Label: "Modbus RTU-Slave, ASCII terminals, HMI, SCADA, or DCS systems" },
        },
        PowerPorts: []PowerPort{
            { Name: "Power", Label: "", Type: "dc-terminal", MaximumDraw: 6, AllocatedDraw: 0 },
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
        },
    },
}
