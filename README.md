# ec2prices
Parses the unofficial EC2 price lists into a struct since they are not available via API.

The data is culled from json files that have the disclaimer:
"This file is intended for use only on aws.amazon.com. We do not guarantee its availability or accuracy."

Currently only supports Linux on-demand instances, but could easily be expanded. The full list of JSON price files is here: http://stackoverflow.com/questions/7334035/get-ec2-pricing-programmatically
