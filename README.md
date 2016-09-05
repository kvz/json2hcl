# json2hcl

Convert JSON to HCL via STDIN / STDOUT

## Use

```bash
cat infra.tf.json | go run main.go
```

## Development

```bash
mkdir -p ~/go/src/github.com/kvz
cd ~/go/src/github.com/kvz
git clone git@github.com:kvz/json2hcl.git
cd json2hcl
go get
```

## Contributors

[Marius Kleidl](https://github.com/Acconut)
[Kevin van Zonneveld](https://github.com/kvz)
