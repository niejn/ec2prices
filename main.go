package ec2prices

import (
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/recursionpharma/ec2prices/files"
)

const (
	Spot     = "Spot"
	OnDemand = "On-Demand"
	Reserved = "Reserved"

	Linux = "Linux"
)

var (
	filesByResource = map[Resource]string{
		Resource{Platform: Linux, PurchaseModel: OnDemand}: files.LinuxOnDemand,
	}

	innerJSON = regexp.MustCompile(`callback\((.*)\);$`) // Ignore comments and callback
)

// Resource represents a kind of thing you can buy on EC2
type Resource struct {
	Platform      string // RHEL, Linux, etc
	PurchaseModel string // Spot, Reserved, On-Demand
}

// PriceList represents the JSON file
type PriceList struct {
	Version float64 `json:"float,omitempty"`
	Config  Config  `json:"config,omitempty"`
}

type Config struct {
	Rate         string   `json:"rate,omitempty"`
	ValueColumns []string `json:"valueColumns,omitempty"`
	Currencies   []string `json:"currencies,omitempty"`
	Regions      []Region `json:"regions,omitempty"`
}

type Region struct {
	Region        string         `json:"region,omitempty"`
	InstanceTypes []InstanceType `json:"instanceTypes,omitempty"`
}

type InstanceType struct {
	Type  string `json:"type,omitempty"`
	Sizes []Size `json:"sizes,omitempty"`
}

type Size struct {
	Size         string        `json:"size,omitempty"`
	VCPU         int           `json:"vCPU,omitempty"`
	ECU          string        `json:"ECU,omitempty"`
	MemoryGiB    string        `json:"memoryGiB,omitempty"`
	StorageGB    string        `json:"storageGB,omitempty"`
	ValueColumns []ValueColumn `json:"valueColumns,omitempty"`
}

type ValueColumn struct {
	Name   string             `json:"name,omitempty"`
	Prices map[string]float64 `json:"prices,omitempty"`
}

// GetPrices parses the JSON file for the given resource, and unmarshals it into a struct
func GetPriceList(r Resource) (*PriceList, error) {
	var f string
	var ok bool
	if f, ok = filesByResource[r]; !ok {
		return nil, fmt.Errorf("Couldn't find a JSON price file for the resource: %+v", r)
	}
	data := []byte(f)
	p := &PriceList{}
	err := json.Unmarshal(data, p)
	if err != nil {
		return nil, err
	}
	return p, nil
}
