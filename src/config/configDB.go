package config

type ConfigDB struct {
	DriverName string `json:"drivername"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	DBName     string `json:"dbname"`
}
