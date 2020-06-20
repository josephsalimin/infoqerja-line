package main

import (
	iqc "infoqerja-line/app/config"
	constant "infoqerja-line/app/utils/constant"
	"log"

	"github.com/Kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// InitiateDatabaseConnection : using configuration files to connect to the mongo db database
func InitiateDatabaseConnection(config iqc.Config) error {
	if err := mgm.SetDefaultConfig(nil, constant.DatabaseName, options.Client().ApplyURI(config.DatabaseURI)); err != nil {
		log.Print(err)
		return err
	}
	return nil
}
