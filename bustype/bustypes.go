package bustype

const (
	BusTypeUnknown = iota
	BusTypeScsi
	BusTypeAtapi
	BusTypeAta
	BusType1394
	BusTypeSsa
	BusTypeFibre
	BusTypeUsb
	BusTypeRAID
	BusTypeiScsi
	BusTypeSas
	BusTypeSata
	BusTypeSd
	BusTypeMmc
	BusTypeVirtual
	BusTypeFileBackedVirtual
	BusTypeSpaces
	BusTypeNvme
	BusTypeSCM
	BusTypeUfs
	BusTypeMax
	BusTypeMaxReserved
)

type BusType byte
