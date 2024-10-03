# PinGen Usage

Go implementation of https://tools.ietf.org/html/rfc4226#section-5

However this is purely to create 1-8 random digit integers to use for
user specified PIN number or similar. More of an exercise in Go than
a widely useful tool ;)

## Usage

```sh
go build

./pingen   # generates 6 digit integer
./pingen 8 # generates 8 digit integer
./pingen 3 # generates 3 digit integer
```
