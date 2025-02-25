#!/bin/bash

curl -vvv -H -X 'POST' \
  'https://api.ism.quantumwake.io/api/v1/query/state/4910bc34-41ce-4425-82e1-a7f2cc20be61' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "filter_groups": [
    {
      "filters": [
        {
          "column": "instruction",
          "operator": "=",
          "value": "Is it problematic to eat cats?"
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
  "state_id": "4910bc34-41ce-4425-82e1-a7f2cc20be61",
  "user_id": "77c17315-3013-5bb8-8c42-32c28618101f"
}'
