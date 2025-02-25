curl -X 'POST' \
  'http://localhost:8081/api/v1/query/state/9ab0f4ac-3d87-4437-9528-0163cc5367a8' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "filter_groups": [
    {
      "filters": [
        {
          "column": "question",
          "operator": "like",
          "value": "%animal%"
        }
      ],
      "group_logic": "AND"
    }
  ],
  "state_id": "9ab0f4ac-3d87-4437-9528-0163cc5367a8",
  "user_id": "77c17315-3013-5bb8-8c42-32c28618101f"
}'
