# Windisk

Windisk is a tiny library to retrieve disk info such as serial number, product id, vendor id, etc... It uses winapi so no need to install external dependencies. 

# Usage

windisk.GetDiskInfo() returns DeviceInfo struct which is;

```go
type DeviceInfo struct {
	SerialNumber     string
	IsRemovableMedia bool
	VendorID         string
	ProductID        string
	ProductRevision  string
	BusType          bustype.BusType
}
```

```go
package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"

	"github.com/sigkilled/windisk"
)

func main() {
	volumeName := filepath.VolumeName(os.Args[0])
	di, err := windisk.GetDiskInfo("\\\\.\\" + volumeName)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%+v\n", di)
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	<-c
}
```