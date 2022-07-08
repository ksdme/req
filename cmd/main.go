package main

import (
	"fmt"

	"github.com/goccy/go-yaml"
	"github.com/ksdme/req/pkg/atoms"
)

var yml = `
version: 1

groups:
  example:
    requests:
      echo:
        method: get
        host: "{host}"
        endpoint: /

        query:
          value:
            if: "{host}"
            then: "{host}"
            else: null

        headers:
          Authorization: JWT {token}

        variables:
          host:
            source:
              - cli:
                  map:
                    local: http://localhost:8000
                  default: local
            required: true

          token:
            source:
              - storage:
                  key: key
                  scope: request
              - cli:
            required: false
`

func main() {
	var collection = atoms.Collection{}

	err := yaml.Unmarshal([]byte(yml), &collection)
	if err != nil {
		panic(err)
	}

	for group_name, group := range collection.Groups {
		for request_name := range group.Requests {
			fmt.Printf("%s:%s", group_name, request_name)
		}
	}
}
