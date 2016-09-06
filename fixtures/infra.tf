"output" = {
  "arn" = {
    "value" = "${aws_dynamodb_table.basic-dynamodb-table.arn}"
  }
}

"output" = {
  "endpoint" = {
    "value" = "http://dynamodb.com:8080/endpoint/${aws_dynamodb_table.basic-dynamodb-table.arn}"
  }
}

"provider" = {
  "aws" = {
    "access_key" = "${var.FREY_AWS_ACCESS_KEY}"

    "region" = "us-east-1"

    "secret_key" = "${var.FREY_AWS_SECRET_KEY}"
  }
}

"resource" = {
  "aws_dynamodb_table" = {
    "basic-dynamodb-table" = {
      "attribute" = {
        "name" = "TopScore"

        "type" = "N"
      }

      "attribute" = {
        "name" = "UserId"

        "type" = "N"
      }

      "attribute" = {
        "name" = "GameTitle"

        "type" = "S"
      }

      "global_secondary_index" = {
        "hash_key" = "GameTitle"

        "name" = "GameTitleIndex"

        "non_key_attributes" = ["UserId"]

        "projection_type" = "INCLUDE"

        "range_key" = "TopScore"

        "read_capacity" = 10

        "write_capacity" = 10
      }

      "hash_key" = "UserId"

      "name" = "GameScores"

      "range_key" = "GameTitle"

      "read_capacity" = 20

      "write_capacity" = 20
    }
  }
}

"variable" = {
  "FREY_OPENSTACK_TENANT_NAME" = {
    "type" = "string"
  }
}

"variable" = {
  "FREY_AWS_ACCESS_KEY" = {
    "type" = "string"
  }
}

"variable" = {
  "FREY_AWS_SECRET_KEY" = {
    "type" = "string"
  }
}

"variable" = {
  "FREY_OPENSTACK_PROJECT_NAME" = {
    "type" = "string"
  }
}

"variable" = {
  "FREY_OPENSTACK_PASSWORD" = {
    "type" = "string"
  }
}

"variable" = {
  "FREY_OPENSTACK_TENANT_ID" = {
    "type" = "string"
  }
}

"variable" = {
  "FREY_OPENSTACK_USERNAME" = {
    "type" = "string"
  }
}

"variable" = {
  "FREY_DO_TOKEN" = {
    "type" = "string"
  }
}

"variable" = {
  "FREY_ENCRYPTION_SECRET" = {
    "type" = "string"
  }
}

"variable" = {
  "FREY_OPENSTACK_AUTH_URL" = {
    "type" = "string"
  }
}

"variable" = {
  "FREY_OPENSTACK_EXTERNAL_GATEWAY" = {
    "type" = "string"
  }
}