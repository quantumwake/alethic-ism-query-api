package data_test

//
//func TestSomething(t *testing.T) {
//	// This is a placeholder test function.
//	// You can add your test cases here.
//
//	query := dsl.StateQuery{
//		FilterGroups: []dsl.FilterGroup{
//			{
//				Filters: []dsl.Filter{
//					{Column: "age_stage", Value: "Infancy & Toddlerhood", Operator: dsl.Equal},
//					{Column: "ethical_framework_code", Value: "VIR", Operator: dsl.Equal},
//				},
//				GroupLogic: "AND",
//			},
//		},
//	}
//
//	dataAccess := data.NewDataAccess(test.DSN)
//	results, err := dataAccess.Query("dd0ce044-53eb-4e88-92e8-bac47cb20d97", query)
//	require.NoError(t, err)
//	require.NotNil(t, results)
//
//	rows := Pivot(results)
//	require.Greater(t, len(rows), 0)
//
//	for _, row := range rows {
//		for _, value := range row {
//			fmt.Printf("%s\t", value)
//		}
//		fmt.Println()
//	}
//}
//
//func Pivot(data []dsl.StateQueryResult) []map[string]any {
//	curIdx := 1
//	var current map[string]any = nil
//	var results = make([]map[string]any, 0)
//
//	for _, cell := range data {
//		if cell.DataIndex != curIdx {
//			if current != nil {
//				results = append(results, current)
//			}
//			current = map[string]any{}
//			curIdx = cell.DataIndex
//		}
//		current[cell.ColumnName] = cell.DataValue
//	}
//
//	return results
//}
//
///*
//
//   def pivot_list_of_dicts(self, data):
//       result = []  # List to hold the pivoted rows
//       current_dict = {}  # Dictionary to hold the current row
//       current_index = 1  # Track the current index
//
//       if not data:
//           return result
//
//       logger.info("pivoting data to table")
//       # Iterate over the data
//       for row in data:
//           # Check if we're starting a new index
//           if row['data_index'] != current_index:
//               # If current_dict is not empty, add it to the result list
//               if current_dict:
//                   result.append(current_dict)
//
//               # Start a new dictionary for the new index
//               current_dict = {}
//               current_index = row['data_index']
//
//           current_dict = {**current_dict, row['column_name']: row['data_value']}
//
//       logger.info("appending final dictionary to list")
//       # Append the final dictionary to the result list
//       if current_dict:
//           result.append(current_dict)
//
//       return result
//
//*/
