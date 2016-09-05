# json2hcl

Convert JSON to HCL via STDIN / STDOUT.

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
and HCL is able to read JSON natively, but for readability, we wanted to have 
HCL directly in our repos, and not the machine readable, intermediate JSON configs.

This tool uses already available Hashicorp libraries to support the conversion.

## Contributors

- [Marius Kleidl](https://github.com/Acconut)
- [Kevin van Zonneveld](https://github.com/kvz)
