package windisk

import (
	"syscall"
	"unsafe"

	"github.com/sigkilled/windisk/bustype"
)

type storagePropertyQuery struct {
	propertyId           int
	queryType            int
	additionalParameters byte
}

type storageDeviceDescriptor struct {
	version               uint32
	size                  uint32
	deviceType            byte
	deviceTypeModifier    byte
	removableMedia        bool
	commandQueueing       bool
	vendorIdOffset        uint32
	productIdOffset       uint32
	productRevisionOffset uint32
	serialNumberOffset    uint32
	busType               byte
	rawPropertiesLength   uint32
	rawDeviceProperties   [1]byte
}

type DeviceInfo struct {
	SerialNumber     string
	IsRemovableMedia bool
	VendorID         string
	ProductID        string
	ProductRevision  string
	BusType          bustype.BusType
}

const (
	ioctl_storage_query_property = uint32(0x002d1400)
	propertyStandardQuery        = 0
	storageDeviceProperty        = 0
)

// GetDiskInfo gathers informations about a disk pointed by the diskpath.
// diskpath must be provided in correct form.
// Ex: \\\\.\\F: or \\\\.\\PHYSICALDRIVE1
// PLease see: https://docs.microsoft.com/en-us/windows/win32/api/ioapiset/nf-ioapiset-deviceiocontrol

func GetDiskInfo(diskpath string) (*DeviceInfo, error) {
	fd, err := syscall.Open(diskpath, syscall.O_RDWR, 0)
	if err != nil {
		return nil, err
	}

	defer syscall.CloseHandle(fd)

	sQuery := storagePropertyQuery{
		propertyId: storageDeviceProperty,
		queryType:  propertyStandardQuery,
	}

	var recivedBytes uint32

	bb := new([1024]byte)
	buf := unsafe.Pointer(bb)

	err = syscall.DeviceIoControl(fd, ioctl_storage_query_property, (*byte)(unsafe.Pointer(&sQuery)), uint32(unsafe.Sizeof(sQuery)), (*byte)(buf), uint32(1024), &recivedBytes, nil)

	if err != nil {
		return nil, err
	}

	devInfo := &DeviceInfo{}
	devData := (*storageDeviceDescriptor)(buf)

	if devData.vendorIdOffset > 0 {
		devInfo.VendorID = string(getData(bb, devData.vendorIdOffset))
	}
	if devData.productIdOffset > 0 {
		devInfo.ProductID = string(getData(bb, devData.productIdOffset))
	}
	if devData.serialNumberOffset > 0 {
		devInfo.SerialNumber = string(getData(bb, devData.serialNumberOffset))
	}
	if devData.productRevisionOffset > 0 {
		devInfo.ProductRevision = string(getData(bb, devData.productRevisionOffset))
	}

	devInfo.IsRemovableMedia = devData.removableMedia
	devInfo.BusType = bustype.BusType(devData.busType)

	return devInfo, nil
}

func getData(buf *[1024]byte, offset uint32) []byte {
	var tmp []byte
	for _, v := range (*buf)[offset:] {
		if v == 0 {
			break
		}
		tmp = append(tmp, v)
	}
	return tmp
}
