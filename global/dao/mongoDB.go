package dao

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"mall/global/config"
	"mall/global/log"
	"time"

	options "go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

func init() {
	dbConfig := config.MongoDBConfig{}
	option := options.Client().SetAuth(options.Credential{
		AuthMechanism: dbConfig.Credential.AuthMechanism,
		AuthSource:    dbConfig.Credential.AuthSource,
		Username:      dbConfig.Credential.Username,
		Password:      dbConfig.Credential.Password,
	})
	if dbConfig.Timeout == 0 {
		option.SetConnectTimeout(10 * time.Second)
	} else {
		option.SetConnectTimeout(time.Duration(dbConfig.Timeout) * time.Second)
	}

	if dbConfig.MinPoolSize == 0 {
		option.SetMinPoolSize(2)
	} else {
		option.SetMinPoolSize(dbConfig.MinPoolSize)
	}

	if dbConfig.MaxPoolSize == 0 {
		option.SetMaxPoolSize(2)
	} else {
		option.SetMaxPoolSize(dbConfig.MaxPoolSize)
	}

	option.SetReadPreference(readpref.Primary())
	option.SetHosts(dbConfig.Host)
	c, err := mongo.Connect(context.TODO(), option)

	if err != nil {
		log.Logger.Error(err.Error())
		panic(err.Error())
		return
	}

	db = c.Database(dbConfig.Database)

}

func GetMongoDB() *mongo.Database {
	return db
}
