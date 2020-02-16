package container

import (
	"database/sql"
	"github.com/dickynovanto1103/User-Management-System/internal/repository/dbsql"
	"github.com/dickynovanto1103/User-Management-System/internal/repository/redis"
	"github.com/dickynovanto1103/User-Management-System/internal/service/config"
	pb "github.com/dickynovanto1103/User-Management-System/proto"
	"google.golang.org/grpc"
)

var Client pb.UserDataServiceClient

func BuildHttpServerClient(conn *grpc.ClientConn) {
	Client = pb.NewUserDataServiceClient(conn)
}

var DBImpl *dbsql.DBImpl
var RedisImpl *redis.RedisImpl
var StatementQueryUser, StatementUpdateNickname, StatementUpdateProfile, StatementQueryPassword *sql.Stmt

func BuildTCPServerDep() {
	configDB := config.LoadConfigDB("config/configDB.json")
	configRedis := config.LoadConfigRedis("config/configRedis.json")

	redisClient := redis.CreateRedisClient(configRedis)
	RedisImpl = redis.CreateRedisWrapper(redisClient)

	dbClient, err := dbsql.PrepareDB(configDB)
	if err != nil {
		panic("error creating db client, err:" + err.Error())
	}

	DBImpl = dbsql.NewDBImpl(dbClient)

	prepareStatements()
}

func prepareStatements() {
	var err error
	StatementQueryUser, err = DBImpl.PrepareQueryUser()
	if err != nil {
		panic("error preparing statement query user, err:" + err.Error())
	}

	StatementQueryPassword, err = DBImpl.PrepareQueryPassword()
	if err != nil {
		panic("error preparing statement query password, err:" + err.Error())
	}

	StatementUpdateNickname, err = DBImpl.PrepareUpdateNickname()
	if err != nil {
		panic("error preparing statement update nickname, err:" + err.Error())
	}

	StatementUpdateProfile, err = DBImpl.PrepareUpdateProfile()
	if err != nil {
		panic("error preparing statement update profile, err:" + err.Error())
	}
}
