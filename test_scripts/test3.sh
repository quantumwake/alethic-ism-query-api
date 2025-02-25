#!/bin/bash

curl -vvv -H -X 'POST' \
  'https://api.ism.quantumwake.io/api/v1/query/state/97a15dcf-a541-431e-a449-42d740bb9e5e' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
    filter_groups": [
    {
      "filters": [
        {
          "column": "instruction",
          "operator": "like",
          "value": "%problematic%"
        },
        {
          "column": "animal",
          "operator": "=",
          "value": "cat"
        }
      ],
      "group_logic": "AND"
    }
  ],
  "state_id": "97a15dcf-a541-431e-a449-42d740bb9e5e",
  "user_id": "77c17315-3013-5bb8-8c42-32c28618101f"
}'
