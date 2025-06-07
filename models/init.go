package models

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db *gorm.DB
)

func Connect() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("can not find the .env file:", err)
		return
	}

	url := os.Getenv("POSTGRES_URL")
	username := os.Getenv("POSTGRES_USERNAME")
	password := os.Getenv("POSTGRES_PASSWORD")
	port := os.Getenv("POSTGRES_PORT")
	DBName := os.Getenv("POSTGRES_DATABASE_NAME")
	pgSsl := os.Getenv("POSTGRES_SSL")
	pgTimeZone := os.Getenv("POSTGRES_TIMEZONE")

	dsn := "host=" + url + " user=" + username + " password=" + password + " dbname=" + DBName + " port=" + port + " sslmode=" + pgSsl + " TimeZone=" + pgTimeZone

	cert, err := tls.LoadX509KeyPair("cert/crt.pem", "cert/key.pem")
	if err != nil {
		log.Fatalln("failed to load client certificate: " + err.Error())
	}

	caCert, err := os.ReadFile("cert/root.pem")
	if err != nil {
		log.Fatalln("can not load the root CA")
		return
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(caCert) {
		log.Fatalln("Failed to add CA cert to pool")
		return
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      certPool,
		MinVersion:   tls.VersionTLS12,
		ServerName:   "postgres.backend.ctf.me",
	}

	connConfig, err := pgx.ParseConfig(dsn)
	if err != nil {
		log.Fatalln("faild to parse dsn")
		return
	}
	connConfig.TLSConfig = tlsConfig

	// Create a custom logger
	file, err := os.OpenFile("gorm.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("failed to open log file:", err)
		return
	}
	newLogger := logger.New(
		log.New(file, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Enabled color
		},
	)

	// d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
	// 	Logger: newLogger,
	// })

	d, err := gorm.Open(postgres.New(postgres.Config{
		DriverName: "pgx",
		DSN:        stdlib.RegisterConnConfig(connConfig),
	}), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		log.Fatalln("failed to connect to the database:", err)
		return
	}
	db = d
}

// func ConnectDB() {
// 	err := godotenv.Load(".env")
// 	if err != nil {
// 		log.Fatalln("can not find the .env file:", err)
// 		return
// 	}

// 	url := os.Getenv("POSTGRES_URL")
// 	username := os.Getenv("POSTGRES_USERNAME")
// 	password := os.Getenv("POSTGRES_PASSWORD")
// 	port := os.Getenv("POSTGRES_PORT")
// 	DBName := os.Getenv("POSTGRES_DATABASE_NAME")
// 	pgSsl := os.Getenv("POSTGRES_SSL")
// 	pgTimeZone := os.Getenv("POSTGRES_TIMEZONE")

// 	dsn := "host=" + url + " user=" + username + " password=" + password + " dbname=" + DBName + " port=" + port + " sslmode=" + pgSsl + " TimeZone=" + pgTimeZone

// 	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
// 		Logger: logger.Default.LogMode(logger.Silent),
// 	})

// 	if err != nil {
// 		log.Fatalln("failed to connect to the database:", err)
// 		return
// 	}
// 	db = d
// }

func Init() {
	Connect()
	models := []interface{}{
		&DifficultyLevel{}, &OperatingSystemType{},
		&User{}, &Student{}, &Admin{}, &Instructor{},
		&RevokedToken{}, &RevokedRefreshToken{}, &RevokedRefreshTokenAdmin{},
		&RevokedRefreshTokenInstructor{}, &RevokedRefreshTokenStudent{},
		&RevokedTokenAdmin{}, &RevokedTokenInstructor{}, &RevokedTokenStudent{},
		&Room{}, &RoomStudent{}, &CourseMachine{}, &Machine{}, &Course{},
		&CourseStudent{}, &MachineStudent{},
		&Question{}, &QuestionStudent{},
		&Flag{}, &Notification{}, &NotificationAdmin{},
		&NotificationInstructor{}, &NotificationStudent{}, HackingMachine{}, &HackingMachineStudent{},
	}

	// added for testing purposes
	for _, model := range models {
		err := db.Migrator().DropTable(model)
		if err != nil {
			log.Printf("Failed to drop table for model %T: %v \n", model, err)
		}
	}

	for _, model := range models {
		err := db.AutoMigrate(model)
		if err != nil {
			log.Fatalf("Failed to migrate model %T: %v \n", model, err)
		}
	}
	CreateDifficultyLevelAndOperatingSystemTypes()

	//added for test purposes
	test()

	hm := &HackingMachine{
		ImageNameOrId: "mintesnotafework/browser-attackbox:v1.0",
	}

	hm.CreateHackingMachine()
}

func CreateDifficultyLevelAndOperatingSystemTypes() {
	var difficultyLevels = []DifficultyLevel{
		{Level: "very easy"},
		{Level: "easy"},
		{Level: "medium"},
		{Level: "hard"},
		{Level: "insane"},
	}

	var operatingSystemTypes = []OperatingSystemType{
		{Type: "windows"},
		{Type: "linux"},
		{Type: "macos"},
		{Type: "android"},
		{Type: "ios"},
		{Type: "other"},
	}

	for _, difficultyLevel := range difficultyLevels {
		if err := db.Create(&difficultyLevel).Error; err != nil {
			log.Printf("Failed to create difficulty level: %v\n", err)
		}
	}

	for _, operatingSystemType := range operatingSystemTypes {
		if err := db.Create(&operatingSystemType).Error; err != nil {
			log.Printf("Failed to create operating system type: %v \n", err)
		}
	}
}
