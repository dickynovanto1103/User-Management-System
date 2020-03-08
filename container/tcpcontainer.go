package container

import (
	"github.com/dickynovanto1103/User-Management-System/internal/repository/dbsql"
	"github.com/dickynovanto1103/User-Management-System/internal/repository/redis"
	"github.com/dickynovanto1103/User-Management-System/internal/service/config"
)

type Repository struct {
	DBImpl        *dbsql.DBImpl
	RedisImpl     *redis.RedisImpl
	SQLStatements *dbsql.SQLStatements
}

func NewRepository(dbImpl *dbsql.DBImpl, redisImpl *redis.RedisImpl, statements *dbsql.SQLStatements) *Repository {
	return &Repository{
		DBImpl:        dbImpl,
		RedisImpl:     redisImpl,
		SQLStatements: statements,
	}
}

func BuildTCPServerDep() *Repository {
	configDB := config.LoadConfigDB("config/configDB.json")
	configRedis := config.LoadConfigRedis("config/configRedis.json")

	redisClient := redis.CreateRedisClient(configRedis)
	redisImpl := redis.CreateRedisWrapper(redisClient)

	dbClient, err := dbsql.PrepareDB(configDB)
	if err != nil {
		panic("error creating db client, err:" + err.Error())
	}

	dbImpl := dbsql.NewDBImpl(dbClient)

	statements := prepareStatements(dbImpl)
	return NewRepository(dbImpl, redisImpl, statements)
}

func prepareStatements(dbImpl *dbsql.DBImpl) *dbsql.SQLStatements {
	statementQueryUser, err := dbImpl.PrepareQueryUser()
	if err != nil {
		panic("error preparing statement query user, err:" + err.Error())
	}

	statementQueryPassword, err := dbImpl.PrepareQueryPassword()
	if err != nil {
		panic("error preparing statement query password, err:" + err.Error())
	}

	statementUpdateNickname, err := dbImpl.PrepareUpdateNickname()
	if err != nil {
		panic("error preparing statement update nickname, err:" + err.Error())
	}

	statementUpdateProfile, err := dbImpl.PrepareUpdateProfile()
	if err != nil {
		panic("error preparing statement update profile, err:" + err.Error())
	}

	sqlStatements := dbsql.NewSQLStatements(statementQueryUser, statementQueryPassword, statementUpdateNickname, statementUpdateProfile)
	return sqlStatements
}
