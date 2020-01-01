[![Build Status](https://travis-ci.org/kvz/json2hcl.svg?branch=master)](https://travis-ci.org/kvz/json2hcl)

# json2hcl (and hcl2json)

Convert JSON to HCL and HCL to JSON via STDIN / STDOUT.

## Warning

We don't use json2hcl anymore ourselves, so we can't invest time into it. However, we're still welcoming PRs.

## Install

Check the [releases](https://github.com/kvz/json2hcl/releases) for the latest version.
Then it's just a matter of downloading the right one for you platform, and making the binary
executable. 

### Linux

Here's how it could look for 64 bits Linux, if you wanted `json2hcl` available globally inside
`/usr/local/bin`:

```bash
curl -SsL https://github.com/kvz/json2hcl/releases/download/v0.0.6/json2hcl_v0.0.6_linux_amd64 \
  | sudo tee /usr/local/bin/json2hcl > /dev/null && sudo chmod 755 /usr/local/bin/json2hcl && json2hcl -version
```

### OSX

Here's how it could look for 64 bits Darwin, if you wanted `json2hcl` available globally inside
`/usr/local/bin`:

```bash
curl -SsL https://github.com/kvz/json2hcl/releases/download/v0.0.6/json2hcl_v0.0.6_darwin_amd64 \
  | sudo tee /usr/local/bin/json2hcl > /dev/null && sudo chmod 755 /usr/local/bin/json2hcl && json2hcl -version
```

## Use

Here's an example [`fixtures/infra.tf.json`](fixtures/infra.tf.json) being
converted to HCL:

```bash
$ json2hcl < fixtures/infra.tf.json
"output" "arn" {
  "value" = "${aws_dynamodb_table.basic-dynamodb-table.arn}"
}
... rest of HCL truncated
```

Typical use would be

```bash
$ json2hcl < fixtures/infra.tf.json > fixtures/infra.tf
```

## hcl2json

As a bonus, the conversation the other way around is also supported via the `-reverse` flag:

```bash
$ json2hcl -reverse < fixtures/infra.tf
{
  "output": [
    {
      "arn": [
        {
          "value": "${aws_dynamodb_table.basic-dynamodb-table.arn}"
        }
      ]
    }, 
  ... rest of JSON truncated
  ]
}
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

If you don't know HCL, read [Why HCL](https://github.com/hashicorp/hcl#why).

As for why json2hcl and hcl2json, we're building a tool called Frey that marries multiple underlying
tools. We'd like configuration previously written in YAML or TOML to now be in HCL now as well. 
It's easy enough to convert the mentioned formats to JSON, and strictly speaking HCL is already 
able to read JSON natively, so why the extra step?

We're doing this for readability and maintainability, we wanted to save 
our infra recipes as HCL directly in our repos, instead of only having machine readable intermediate 
JSON that we'd need to hack on. This saves time spotting problems, and makes the experience somewhat 
enjoyable even.

In the off-chance you too have machine-readable JSON and are interested in converting that
to the more human-being friendly HCL format, we thought we'd share this.

It's no rocket science, we're using already available HashiCorp libraries to support the conversion,
HashiCorp could have easily released their own tools around this, and perhaps they will, but 
so far, they haven't.

## Changelog

### Ideabox (Unplanned)

- [ ] Give the README.md some love

### v0.0.7 (Unreleased)

- [ ] Tests

### v0.0.6 (2016-09-06)

- [x] Deprecate goxc in favor of native builds

### v0.0.5 (2016-09-06)

- [x] Add hcl2json via the `-reverse` flag 

### v0.0.4 (2016-09-05)

- [x] Error handling
- [x] Cross-compiling and shipping releases

## Contributors

- [Marius Kleidl](https://github.com/Acconut)
- [Kevin van Zonneveld](https://github.com/kvz)
