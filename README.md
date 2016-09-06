# json2hcl

Convert JSON to HCL via STDIN / STDOUT.

This tool is in BETA. We'll update the readme from `go run main.go` -> `json2hcl` once we're
shipping binaries for all platforms.

## Use

Here's an example [`fixtures/infra.tf.json`](fixtures/infra.tf.json)

```bash
$ cat fixtures/infra.tf.json
{
  "output": {
    "arn": {
      "value": "${aws_dynamodb_table.basic-dynamodb-table.arn}"
    },
    "endpoint": {
      "value": "http://dynamodb.com:8080/endpoint/${aws_dynamodb_table.basic-dynamodb-table.arn}"
    }
  }
  ... rest of JSON truncated
}
... 
```

Now run it through json2hcl, and it will be converted to HCL

```bash
$ cat fixtures/infra.tf.json | go run main.go
"output" "arn" {
  "value" "${aws_dynamodb_table.basic-dynamodb-table.arn}"
}

"output" "endpoint" {
  "value" "http://dynamodb.com:8080/endpoint/${aws_dynamodb_table.basic-dynamodb-table.arn}"
}
... rest of HCL truncated
```

Typical use would be

```bash
$ cat fixtures/infra.tf.json | go run main.go > infra.tf
```

## hcl2json

As a bonus, the conversation the other way around is also supported via the `-reverse` flag:

```bash
cat fixtures/infra.tf | go run main.go -reverse
# Writes JSON to the STDOUT
```


## Development

```bash
mkdir -p ~/go/src/github.com/kvz
cd ~/go/src/github.com/kvz
git clone git@github.com:kvz/json2hcl.git
cd json2hcl
go get
```

## Why?

We're building a tool called Frey, and we like HCL. We'd like configuration previously 
built in YAML or TOML to be in HCL now as well. It's easy to convert those to JSON,
and HCL is able to read JSON natively, but for readability and maintainability, we wanted to save 
our infra recipes as HCL directly in our repos, instead of only having machine readable intermediate 
JSON that we'd need to hack on.

In the off-chance you too have machine-readable JSON but are interested in converting that
to the more human-being friendly HCL format, we thought we'd share.

This tool uses already available HashiCorp libraries to support the conversion.

## Todo

- [ ] Should we perhaps also support the conversion the other way? Just for fun?
- [ ] Tests
- [ ] Error handling
- [ ] Tasks for cross-compiling and shipping releases
- [ ] Give the README.md some love

## Contributors

- [Marius Kleidl](https://github.com/Acconut)
- [Kevin van Zonneveld](https://github.com/kvz)
