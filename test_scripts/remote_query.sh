curl -X 'POST' \
  'https://api.ism.quantumwake.io/api/v1/state/8a516a57-f473-4027-b01c-ba5268571831/query' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "filter_groups": [
    {
      "filters": [
        {
          "column": "response_id",
          "operator": "=",
          "value": "6711c9816bba2c051ed1b0c8"
        }
      ],
      "group_logic": "AND"
    }
  ]
}'
