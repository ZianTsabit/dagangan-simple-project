package database

import "gorm.io/gorm"

type DbInstance struct {
	Db *gorm.DB
}

var Db DbInstance

func ConnectDb() {

	dsn := fnt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5434 sslmode=disable TimeZone=Asia/Jakarta", 
		os.Getenv("DB_USER"), 
		os.Getenv("DB_PASSWORD"), 
		os.Getenv("DB_NAME")
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		panic(err)
	}

	log.Println("Connected to database.")
	db.Logger = logger.Default.LogMode(logger.Info)


}

