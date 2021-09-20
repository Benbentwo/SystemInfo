package system_profiler

import "time"

type SystemProfiler struct {
	SPAirPortDataType []struct {
		SpairportAirportInterfaces []struct {
			Name                             string `json:"_name"`
			SpairportAirdropChannel          int    `json:"spairport_airdrop_channel"`
			SpairportCapsAirdrop             string `json:"spairport_caps_airdrop"`
			SpairportCapsAutounlock          string `json:"spairport_caps_autounlock"`
			SpairportCapsWow                 string `json:"spairport_caps_wow"`
			SpairportStatusInformation       string `json:"spairport_status_information"`
			SpairportSupportedChannels       []int  `json:"spairport_supported_channels"`
			SpairportSupportedPhymodes       string `json:"spairport_supported_phymodes"`
			SpairportWirelessCardType        string `json:"spairport_wireless_card_type"`
			SpairportWirelessCountryCode     string `json:"spairport_wireless_country_code"`
			SpairportWirelessFirmwareVersion string `json:"spairport_wireless_firmware_version"`
			SpairportWirelessLocale          string `json:"spairport_wireless_locale"`
			SpairportWirelessMacAddress      string `json:"spairport_wireless_mac_address"`
		} `json:"spairport_airport_interfaces"`
		SpairportSoftwareInformation struct {
			SpairportCorewlanVersion    string `json:"spairport_corewlan_version"`
			SpairportCorewlankitVersion string `json:"spairport_corewlankit_version"`
			SpairportDiagnosticsVersion string `json:"spairport_diagnostics_version"`
			SpairportExtraVersion       string `json:"spairport_extra_version"`
			SpairportFamilyVersion      string `json:"spairport_family_version"`
			SpairportProfilerVersion    string `json:"spairport_profiler_version"`
			SpairportUtilityVersion     string `json:"spairport_utility_version"`
		} `json:"spairport_software_information"`
	} `json:"SPAirPortDataType"`
	SPAudioDataType []struct {
		Items []struct {
			Name                              string `json:"_name"`
			CoreaudioDefaultAudioInputDevice  string `json:"coreaudio_default_audio_input_device,omitempty"`
			CoreaudioDeviceInput              int    `json:"coreaudio_device_input,omitempty"`
			CoreaudioDeviceManufacturer       string `json:"coreaudio_device_manufacturer"`
			CoreaudioDeviceSrate              int    `json:"coreaudio_device_srate"`
			CoreaudioDeviceTransport          string `json:"coreaudio_device_transport"`
			CoreaudioInputSource              string `json:"coreaudio_input_source,omitempty"`
			Properties                        string `json:"_properties,omitempty"`
			CoreaudioDefaultAudioOutputDevice string `json:"coreaudio_default_audio_output_device,omitempty"`
			CoreaudioDefaultAudioSystemDevice string `json:"coreaudio_default_audio_system_device,omitempty"`
			CoreaudioDeviceOutput             int    `json:"coreaudio_device_output,omitempty"`
			CoreaudioOutputSource             string `json:"coreaudio_output_source,omitempty"`
		} `json:"_items"`
		Name string `json:"_name"`
	} `json:"SPAudioDataType"`
	SPBluetoothDataType []struct {
		AppleBluetoothVersion string `json:"apple_bluetooth_version"`
		DeviceTitle           []struct {
			Device0 struct {
				DeviceCoreSpec                 string `json:"device_core_spec"`
				DeviceFwVersion                string `json:"device_fw_version"`
				DeviceMajorClassOfDeviceString string `json:"device_majorClassOfDevice_string"`
				DeviceManufacturer             string `json:"device_manufacturer"`
				DeviceMinorClassOfDeviceString string `json:"device_minorClassOfDevice_string"`
				DeviceProductID                string `json:"device_productID"`
				DeviceServices                 string `json:"device_services"`
				DeviceSupportsEDR              string `json:"device_supportsEDR"`
				DeviceSupportsESCO             string `json:"device_supportsESCO"`
				DeviceSupportsSSP              string `json:"device_supportsSSP"`
				DeviceVendorID                 string `json:"device_vendorID"`
			} `json:"Device 0,omitempty"`
			Device1 struct {
				DeviceCoreSpec                 string `json:"device_core_spec"`
				DeviceFwVersion                string `json:"device_fw_version"`
				DeviceMajorClassOfDeviceString string `json:"device_majorClassOfDevice_string"`
				DeviceManufacturer             string `json:"device_manufacturer"`
				DeviceMinorClassOfDeviceString string `json:"device_minorClassOfDevice_string"`
				DeviceProductID                string `json:"device_productID"`
				DeviceServices                 string `json:"device_services"`
				DeviceSupportsEDR              string `json:"device_supportsEDR"`
				DeviceSupportsESCO             string `json:"device_supportsESCO"`
				DeviceSupportsSSP              string `json:"device_supportsSSP"`
				DeviceVendorID                 string `json:"device_vendorID"`
			} `json:"Device 1,omitempty"`
			Device2 struct {
				DeviceCoreSpec                 string `json:"device_core_spec"`
				DeviceFwVersion                string `json:"device_fw_version"`
				DeviceMajorClassOfDeviceString string `json:"device_majorClassOfDevice_string"`
				DeviceManufacturer             string `json:"device_manufacturer"`
				DeviceMinorClassOfDeviceString string `json:"device_minorClassOfDevice_string"`
				DeviceProductID                string `json:"device_productID"`
				DeviceServices                 string `json:"device_services"`
				DeviceSupportsEDR              string `json:"device_supportsEDR"`
				DeviceSupportsESCO             string `json:"device_supportsESCO"`
				DeviceSupportsSSP              string `json:"device_supportsSSP"`
				DeviceVendorID                 string `json:"device_vendorID"`
			} `json:"Device 2,omitempty"`
		} `json:"device_title"`
		LocalDeviceTitle struct {
			GeneralAutoseekKeyboard       string `json:"general_autoseek_keyboard"`
			GeneralAutoseekPointing       string `json:"general_autoseek_pointing"`
			GeneralChipset                string `json:"general_chipset"`
			GeneralConnectable            string `json:"general_connectable"`
			GeneralFwVersion              string `json:"general_fw_version"`
			GeneralHardwareTransport      string `json:"general_hardware_transport"`
			GeneralMfg                    string `json:"general_mfg"`
			GeneralProductID              string `json:"general_productID"`
			GeneralRemoteWake             string `json:"general_remoteWake"`
			GeneralSupportsHandoff        string `json:"general_supports_handoff"`
			GeneralSupportsInstantHotspot string `json:"general_supports_instantHotspot"`
			GeneralSupportsLowEnergy      string `json:"general_supports_lowEnergy"`
			GeneralVendorID               string `json:"general_vendorID"`
		} `json:"local_device_title"`
	} `json:"SPBluetoothDataType"`
	SPCameraDataType []struct {
		Name             string `json:"_name"`
		SpcameraModelId  string `json:"spcamera_model-id"`
		SpcameraUniqueId string `json:"spcamera_unique-id"`
	} `json:"SPCameraDataType"`
	SPCardReaderDataType  []interface{} `json:"SPCardReaderDataType"`
	SPDiagnosticsDataType []struct {
		Name              string    `json:"_name"`
		SpdiagsLastRunKey time.Time `json:"spdiags_last_run_key"`
		SpdiagsResultKey  string    `json:"spdiags_result_key"`
	} `json:"SPDiagnosticsDataType"`
	SPDiscBurningDataType []interface{} `json:"SPDiscBurningDataType"`
	SPDisplaysDataType    []struct {
		Name                                 string `json:"_name"`
		SpdisplaysVram                       string `json:"_spdisplays_vram,omitempty"`
		SpdisplaysAutomaticGraphicsSwitching string `json:"spdisplays_automatic_graphics_switching"`
		SpdisplaysDeviceId                   string `json:"spdisplays_device-id"`
		SpdisplaysGmuxVersion                string `json:"spdisplays_gmux-version"`
		SpdisplaysMetalfamily                string `json:"spdisplays_metalfamily"`
		SpdisplaysNdrvs                      []struct {
			IODisplayEDID                  string `json:"_IODisplayEDID"`
			Name                           string `json:"_name"`
			SpdisplaysDisplayProductId     string `json:"_spdisplays_display-product-id"`
			SpdisplaysDisplaySerialNumber2 string `json:"_spdisplays_display-serial-number2"`
			SpdisplaysDisplayVendorId      string `json:"_spdisplays_display-vendor-id"`
			SpdisplaysDisplayWeek          string `json:"_spdisplays_display-week"`
			SpdisplaysDisplayYear          string `json:"_spdisplays_display-year"`
			SpdisplaysDisplayID            string `json:"_spdisplays_displayID"`
			SpdisplaysDisplayPath          string `json:"_spdisplays_displayPath"`
			SpdisplaysDisplayportDevice    struct {
				Name                                  string `json:"_name"`
				SpdisplaysDisplayportCurrentBandwidth string `json:"spdisplays_displayport_current_bandwidth"`
				SpdisplaysDisplayportCurrentLanes     string `json:"spdisplays_displayport_current_lanes"`
				SpdisplaysDisplayportCurrentSpread    string `json:"spdisplays_displayport_current_spread"`
				SpdisplaysDisplayportDPCDVersion      string `json:"spdisplays_displayport_DPCD_version"`
				SpdisplaysDisplayportErrorsLane0      string `json:"spdisplays_displayport_errors_lane0"`
				SpdisplaysDisplayportErrorsLane1      string `json:"spdisplays_displayport_errors_lane1"`
				SpdisplaysDisplayportErrorsLane2      string `json:"spdisplays_displayport_errors_lane2"`
				SpdisplaysDisplayportErrorsLane3      string `json:"spdisplays_displayport_errors_lane3"`
				SpdisplaysDisplayportHdcpCapability   string `json:"spdisplays_displayport_hdcp_capability"`
				SpdisplaysDisplayportMaxBandwidth     string `json:"spdisplays_displayport_max_bandwidth"`
				SpdisplaysDisplayportMaxLanes         string `json:"spdisplays_displayport_max_lanes"`
				SpdisplaysDisplayportMaxSpread        string `json:"spdisplays_displayport_max_spread"`
				SpdisplaysDisplayportSinkAsciiName    string `json:"spdisplays_displayport_sink_ascii_name"`
				SpdisplaysDisplayportSinkChipVersion  string `json:"spdisplays_displayport_sink_chip_version"`
				SpdisplaysDisplayportSinkCount        string `json:"spdisplays_displayport_sink_count"`
				SpdisplaysDisplayportSinkSwVersion    string `json:"spdisplays_displayport_sink_sw_version"`
				SpdisplaysDisplayportSinkVendor       string `json:"spdisplays_displayport_sink_vendor"`
				SpdisplaysDisplayportValidErrorLane0  string `json:"spdisplays_displayport_valid_error_lane0"`
				SpdisplaysDisplayportValidErrorLane1  string `json:"spdisplays_displayport_valid_error_lane1"`
				SpdisplaysDisplayportValidErrorLane2  string `json:"spdisplays_displayport_valid_error_lane2"`
				SpdisplaysDisplayportValidErrorLane3  string `json:"spdisplays_displayport_valid_error_lane3"`
			} `json:"_spdisplays_displayport_device"`
			SpdisplaysDisplayRegID      string `json:"_spdisplays_displayRegID"`
			SpdisplaysEdid              string `json:"_spdisplays_edid"`
			SpdisplaysPixels            string `json:"_spdisplays_pixels"`
			SpdisplaysResolution        string `json:"_spdisplays_resolution"`
			SpdisplaysAmbientBrightness string `json:"spdisplays_ambient_brightness"`
			SpdisplaysConnectionType    string `json:"spdisplays_connection_type"`
			SpdisplaysDepth             string `json:"spdisplays_depth"`
			SpdisplaysDisplayType       string `json:"spdisplays_display_type"`
			SpdisplaysMain              string `json:"spdisplays_main"`
			SpdisplaysMirror            string `json:"spdisplays_mirror"`
			SpdisplaysOnline            string `json:"spdisplays_online"`
			SpdisplaysPixelresolution   string `json:"spdisplays_pixelresolution"`
		} `json:"spdisplays_ndrvs,omitempty"`
		SpdisplaysRevisionId       string `json:"spdisplays_revision-id"`
		SpdisplaysVendor           string `json:"spdisplays_vendor"`
		SpdisplaysVramShared       string `json:"spdisplays_vram_shared,omitempty"`
		SppciBus                   string `json:"sppci_bus"`
		SppciDeviceType            string `json:"sppci_device_type"`
		SppciModel                 string `json:"sppci_model"`
		SpdisplaysEfiVersion       string `json:"spdisplays_efi-version,omitempty"`
		SpdisplaysOptionromVersion string `json:"spdisplays_optionrom-version,omitempty"`
		SpdisplaysPcieWidth        string `json:"spdisplays_pcie_width,omitempty"`
		SpdisplaysRomRevision      string `json:"spdisplays_rom-revision,omitempty"`
		SpdisplaysVbiosVersion     string `json:"spdisplays_vbios-version,omitempty"`
		SpdisplaysVram1            string `json:"spdisplays_vram,omitempty"`
	} `json:"SPDisplaysDataType"`
	SPEthernetDataType []struct {
		Name                       string `json:"_name"`
		SpethernetBSDName          string `json:"spethernet_BSD_Name"`
		SpethernetBUNDLEIDentifier string `json:"spethernet_BUNDLE_IDentifier"`
		SpethernetDeviceType       string `json:"spethernet_device_type"`
		SpethernetKextPath         string `json:"spethernet_kext_path"`
		SpethernetVersion          string `json:"spethernet_version"`
		SpusbethernetMacAddress    string `json:"spusbethernet_mac_address"`
		SpusbethernetPID           string `json:"spusbethernet_PID"`
		SpusbethernetVID           string `json:"spusbethernet_VID"`
	} `json:"SPEthernetDataType"`
	SPFibreChannelDataType []interface{} `json:"SPFibreChannelDataType"`
	SPFireWireDataType     []interface{} `json:"SPFireWireDataType"`
	SPHardwareDataType     []struct {
		Name                  string `json:"_name"`
		BootRomVersion        string `json:"boot_rom_version"`
		CpuType               string `json:"cpu_type"`
		CurrentProcessorSpeed string `json:"current_processor_speed"`
		L2CacheCore           string `json:"l2_cache_core"`
		L3Cache               string `json:"l3_cache"`
		MachineModel          string `json:"machine_model"`
		MachineName           string `json:"machine_name"`
		NumberProcessors      int    `json:"number_processors"`
		Packages              int    `json:"packages"`
		PhysicalMemory        string `json:"physical_memory"`
		PlatformCpuHtt        string `json:"platform_cpu_htt"`
	} `json:"SPHardwareDataType"`
	SPiBridgeDataType []struct {
		Name             string `json:"_name"`
		IbridgeBootUuid  string `json:"ibridge_boot_uuid"`
		IbridgeBuild     string `json:"ibridge_build"`
		IbridgeModelName string `json:"ibridge_model_name"`
	} `json:"SPiBridgeDataType"`
	SPMemoryDataType []struct {
		Items []struct {
			Name             string `json:"_name"`
			DimmManufacturer string `json:"dimm_manufacturer"`
			DimmPartNumber   string `json:"dimm_part_number"`
			DimmSize         string `json:"dimm_size"`
			DimmSpeed        string `json:"dimm_speed"`
			DimmStatus       string `json:"dimm_status"`
			DimmType         string `json:"dimm_type"`
		} `json:"_items"`
		Name                string `json:"_name"`
		GlobalEccState      string `json:"global_ecc_state"`
		IsMemoryUpgradeable string `json:"is_memory_upgradeable"`
	} `json:"SPMemoryDataType"`
	SPNetworkDataType []struct {
		Name                  string `json:"_name"`
		Hardware              string `json:"hardware"`
		Interface             string `json:"interface"`
		SpnetworkServiceOrder int    `json:"spnetwork_service_order"`
		Type                  string `json:"type"`
	} `json:"SPNetworkDataType"`
	SPNVMeDataType []struct {
		Items []struct {
			Name              string `json:"_name"`
			BsdName           string `json:"bsd_name"`
			DetachableDrive   string `json:"detachable_drive"`
			DeviceModel       string `json:"device_model"`
			DeviceRevision    string `json:"device_revision"`
			PartitionMapType  string `json:"partition_map_type"`
			RemovableMedia    string `json:"removable_media"`
			Size              string `json:"size"`
			SizeInBytes       int64  `json:"size_in_bytes"`
			SmartStatus       string `json:"smart_status"`
			SpnvmeLinkspeed   string `json:"spnvme_linkspeed"`
			SpnvmeLinkwidth   string `json:"spnvme_linkwidth"`
			SpnvmeTrimSupport string `json:"spnvme_trim_support"`
		} `json:"_items"`
		Name string `json:"_name"`
	} `json:"SPNVMeDataType"`
	SPParallelATADataType  []interface{} `json:"SPParallelATADataType"`
	SPParallelSCSIDataType []interface{} `json:"SPParallelSCSIDataType"`
	SPPCIDataType          []interface{} `json:"SPPCIDataType"`
	SPPrintersDataType     []struct {
		Name            string `json:"_name"`
		Airprintversion string `json:"airprintversion"`
		CreationDate    string `json:"creationDate"`
		CupsFilters     []struct {
			ApplicationVndCupsPdf string `json:"application/vnd.cups-pdf,omitempty"`
			ImageJpeg             string `json:"image/jpeg,omitempty"`
			ImageUrf              string `json:"image/urf,omitempty"`
		} `json:"cups filters"`
		Cupsversion            string        `json:"cupsversion"`
		Default                string        `json:"default"`
		Driverversion          string        `json:"driverversion"`
		FaxSupport             string        `json:"Fax Support"`
		Ppd                    string        `json:"ppd"`
		Ppdfileversion         string        `json:"ppdfileversion"`
		Printercommands        string        `json:"printercommands"`
		Printerfirmwareversion string        `json:"printerfirmwareversion"`
		Printerpdes            []interface{} `json:"printerpdes"`
		Printersharing         string        `json:"printersharing"`
		Printserver            string        `json:"printserver"`
		Psversion              string        `json:"psversion"`
		Scanner                string        `json:"scanner"`
		Scannerappbundlepath   string        `json:"scannerappbundlepath,omitempty"`
		Scannerapppath         string        `json:"scannerapppath,omitempty"`
		Scannerappversion      string        `json:"scannerappversion,omitempty"`
		ScannerUUID            string        `json:"scannerUUID,omitempty"`
		Shared                 string        `json:"shared"`
		Status                 string        `json:"status"`
		Urfversion             string        `json:"urfversion"`
		Uri                    string        `json:"uri"`
	} `json:"SPPrintersDataType"`
	SPPrintersSoftwareDataType []struct {
		Name string `json:"_name"`
		Ppds []struct {
			InfoPath    string `json:"info path"`
			InfoVersion string `json:"info version"`
		} `json:"ppds,omitempty"`
		Printers []struct {
			InfoPath    string `json:"info path"`
			InfoVersion string `json:"info version"`
		} `json:"printers,omitempty"`
		ImageCaptureDevices []struct {
			InfoPath    string `json:"info path"`
			InfoVersion string `json:"info version"`
		} `json:"image capture devices,omitempty"`
		ImageCaptureSupport []struct {
			InfoPath    string `json:"info path"`
			InfoVersion string `json:"info version"`
		} `json:"image capture support,omitempty"`
		Extensions []struct {
			InfoPath    string `json:"info path"`
			InfoVersion string `json:"info version"`
		} `json:"extensions,omitempty"`
		LibExtensions []interface{} `json:"libExtensions,omitempty"`
	} `json:"SPPrintersSoftwareDataType"`
	SPSASDataType           []interface{} `json:"SPSASDataType"`
	SPSecureElementDataType []struct {
		CtlFw              string `json:"ctl_fw"`
		CtlHw              string `json:"ctl_hw"`
		CtlInfo            string `json:"ctl_info"`
		CtlMw              string `json:"ctl_mw"`
		SeDevice           string `json:"se_device"`
		SeFw               string `json:"se_fw"`
		SeHw               string `json:"se_hw"`
		SeId               string `json:"se_id"`
		SeInRestrictedMode string `json:"se_in_restricted_mode"`
		SeInfo             string `json:"se_info"`
		SeOsVersion        string `json:"se_os_version"`
		SePlt              string `json:"se_plt"`
		SeProdSigned       string `json:"se_prod_signed"`
	} `json:"SPSecureElementDataType"`
	SPSerialATADataType []interface{} `json:"SPSerialATADataType"`
	SPSoftwareDataType  []struct {
		Name          string `json:"_name"`
		KernelVersion string `json:"kernel_version"`
		OsVersion     string `json:"os_version"`
		Uptime        string `json:"uptime"`
	} `json:"SPSoftwareDataType"`
	SPSPIDataType         []interface{} `json:"SPSPIDataType"`
	SPThunderboltDataType []struct {
		Name           string `json:"_name"`
		DeviceNameKey  string `json:"device_name_key"`
		DomainUuidKey  string `json:"domain_uuid_key"`
		Receptacle1Tag struct {
			CurrentLinkWidthKey string `json:"current_link_width_key"`
			CurrentSpeedKey     string `json:"current_speed_key"`
			LcVersionKey        string `json:"lc_version_key"`
			LinkStatusKey       string `json:"link_status_key"`
			ReceptacleIdKey     string `json:"receptacle_id_key"`
			ReceptacleStatusKey string `json:"receptacle_status_key"`
		} `json:"receptacle_1_tag"`
		Receptacle2Tag struct {
			CurrentLinkWidthKey string `json:"current_link_width_key"`
			CurrentSpeedKey     string `json:"current_speed_key"`
			LcVersionKey        string `json:"lc_version_key"`
			LinkStatusKey       string `json:"link_status_key"`
			ReceptacleIdKey     string `json:"receptacle_id_key"`
			ReceptacleStatusKey string `json:"receptacle_status_key"`
		} `json:"receptacle_2_tag"`
		RouteStringKey   string `json:"route_string_key"`
		SwitchUidKey     string `json:"switch_uid_key"`
		SwitchVersionKey string `json:"switch_version_key"`
		VendorNameKey    string `json:"vendor_name_key"`
	} `json:"SPThunderboltDataType"`
	SPUSBDataType []struct {
		Name           string `json:"_name"`
		HostController string `json:"host_controller"`
		PciDevice      string `json:"pci_device,omitempty"`
		PciRevision    string `json:"pci_revision,omitempty"`
		PciVendor      string `json:"pci_vendor,omitempty"`
		UsbBusNumber   string `json:"usb_bus_number,omitempty"`
		Items          []struct {
			Name             string `json:"_name"`
			BcdDevice        string `json:"bcd_device"`
			BuiltInDevice    string `json:"Built-in_Device,omitempty"`
			BusPower         string `json:"bus_power,omitempty"`
			BusPowerUsed     string `json:"bus_power_used,omitempty"`
			DeviceSpeed      string `json:"device_speed,omitempty"`
			ExtraCurrentUsed string `json:"extra_current_used,omitempty"`
			LocationId       string `json:"location_id"`
			Manufacturer     string `json:"manufacturer"`
			ProductId        string `json:"product_id"`
			SerialNum        string `json:"serial_num"`
			VendorId         string `json:"vendor_id"`
		} `json:"_items,omitempty"`
	} `json:"SPUSBDataType"`
	SPWWANDataType []interface{} `json:"SPWWANDataType"`
}
