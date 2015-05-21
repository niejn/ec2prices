# ec2prices [![Doc Status](https://godoc.org/github.com/recursionpharma/ec2prices?status.png)](https://godoc.org/github.com/recursionpharma/ec2prices)
Parses the unofficial EC2 price lists into a struct since they are not available via an API call.

The data is culled from JSON files available from Amazon, but they state:
> This file is intended for use only on aws.amazon.com. We do not guarantee its availability or accuracy.

Right now supports Linux On-Demand instances, but is easily expanded.

## Usage

```go
package main

import (
	"github.com/kr/pretty"
	ec2p "github.com/recursionpharma/ec2prices"
)

func main() {
	// Pricing for Linux On-Demand instances
	r := ec2p.Resource{Platform: ec2p.Linux, PurchaseModel: ec2p.OnDemand}
	prices, _ := ec2p.GetPriceList(r)
	pretty.Print(prices)
}
```

Output:

```
&ec2prices.PriceList{
    Version: 0.01,
    Config:  ec2prices.Config{
        Rate:         "perhr",
        ValueColumns: {"vCPU", "ECU", "memoryGiB", "storageGB", "linux"},
        Currencies:   {"USD"},
        Regions:      {
            {
                Region:        "us-east-1",
                InstanceTypes: {
                    {
                        Type:  "generalCurrentGen",
                        Sizes: {
                            {
                                Size:         "t2.micro",
                                VCPU:         1,
                                ECU:          "variable",
                                MemoryGiB:    1,
                                StorageGB:    "ebsonly",
                                ValueColumns: {
                                    {
                                        Name:   "linux",
                                        Prices: {"USD":0.013},
                                    },
                                },
                            },
        [...]

```

## Contributing

Enjoy and do please contribute pricing for other resource types if you use
them. For most cases this involves downloading the right JSON file and adding
a few constants.

1. Download the JSON file you want from: http://stackoverflow.com/questions/7334035/get-ec2-pricing-programmatically

2. Put the inner JSON (from between `callback( ... );` into http://jsonformat.com/ so that the keys are double-quoted (Go's JSON unmarshaller will blow up otherwise).

3. Re-compress the resulting properly formatted JSON using something like
   http://www.httputility.net/json-minifier.aspx.

4. Add this JSON data as a string variable in files/files.go

5. Add a mapping from a resource to this JSON variable in main.go.
