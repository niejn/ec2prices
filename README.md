# ec2prices
Parses the unofficial EC2 price lists into a struct since they are not available via an API call.

The data is culled from JSON files available from Amazon, but they state:
> This file is intended for use only on aws.amazon.com. We do not guarantee its availability or accuracy.

Right now supports Linux On-Demand instances, but is easily expanded.

## Usage

```
import ec2p "github.com/recursionpharma/ec2prices"

// I want pricing for Linux On-Demand instances

r := ec2p.Resource{Platform: ec2p.Linux, PurchaseModel: ec2p.OnDemand}
prices, err := ec2p.GetPriceList(r)
```

## Contributing

Enjoy and do please contribute pricing for other resource types if you use
them. For most cases this involves downloading the right JSON file and adding
a few constants.

The full list of JSON price files is here: http://stackoverflow.com/questions/7334035/get-ec2-pricing-programmatically
