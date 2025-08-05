package data

//
//type Access struct {
//	DSN string
//	DB  *gorm.DB
//}
//
//func NewDataAccess(dsn string) *Access {
//	da := &Access{
//		DSN: dsn,
//	}
//	err := da.Connect()
//	if err != nil {
//		log.Fatalf("unable to connect to database: %v", err)
//	}
//	return da
//}
//
//func (da *Access) Execute(sql string, values ...interface{}) error {
//	var tx *gorm.DB
//	if len(values) > 0 {
//		tx = da.DB.Exec(sql, values)
//	} else {
//		tx = da.DB.Exec(sql)
//	}
//
//	if err := tx.Error; err != nil {
//		return fmt.Errorf("unable to execute sql %s, err: %v", sql, err)
//	}
//	return nil
//}
//
//func (da *Access) Connect() error {
//	var err error
//	da.DB, err = gorm.Open(postgres.Open(da.DSN), &gorm.Config{
//		Logger: logger.Default.LogMode(logger.Info),
//	})
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//func (da *Access) AutoMigrate() error {
//	err := da.DB.AutoMigrate(
//	//entity.Region{},
//
//	// workload specific entities (the profile of a workload)
//	//entity.WorkloadProfile{},
//	//entity.Workload{},
//
//	// compute specific entities (the target compute resource for the profile for a given compute provider)
//	//entity.ComputeProfile{},
//	)
//
//	return err
//}
//
//// Query executes a state query and returns the results.
//func (da *Access) Query(stateID string, query dsl.StateQuery) ([]dsl.StateQueryResult, error) {
//	// Validate UUID
//	if err := utils.ValidateUUID(stateID); err != nil {
//		return nil, fmt.Errorf("invalid UUID: %v", err)
//	}
//
//	// Build the final SQL query and arguments
//	dataSQL, dataArgs, err := query.BuildFinalQuery(stateID)
//	if err != nil {
//		return nil, fmt.Errorf("failed to build final query: %v", err)
//	}
//
//	// Execute the final query to get the results
//	var results []dsl.StateQueryResult
//	err = da.DB.Raw(dataSQL, dataArgs...).Scan(&results).Error
//	if err != nil {
//		return nil, fmt.Errorf("failed to fetch data values: %v", err)
//	}
//
//	return results, nil
//}
//
//// InitializeNewDataAccessFromEnvDSN initializes a new data access object from the DSN environment variable.
//func InitializeNewDataAccessFromEnvDSN() *Access {
//	dsn, ok := os.LookupEnv("DSN")
//	if !ok {
//		dsn = "host=localhost port=5432 user=postgres password=postgres1 dbname=postgres sslmode=disable"
//	}
//
//	dataAccess := NewDataAccess(dsn)
//	if dataAccess == nil {
//		panic(fmt.Errorf("unable to connect to database, check database is accessible and dsn: %s", dsn))
//	}
//
//	return dataAccess
//}
