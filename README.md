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

1. Download the JSON file you want from: http://stackoverflow.com/questions/7334035/get-ec2-pricing-programmatically

2. Put the inner JSON (from between `callback( ... );` into http://jsonformat.com/ so that the keys are double-quoted (Go's JSON unmarshaller will blow up otherwise).

3. Re-compress the resulting properly formatted JSON using something like
   http://www.httputility.net/json-minifier.aspx.

4. Add this JSON data as a string variable in files/files.go

5. Add a mapping from a resource to this JSON variable in main.go.
