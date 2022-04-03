package db

import (
	"fmt"
	"CampusRecruitment/pkg/consts"
	"CampusRecruitment/pkg/infra/log"
	"strings"
	"time"

	"github.com/pkg/errors"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	defaultDB *gorm.DB
	drivers   = make(map[string]func(dsn string) gorm.Dialector)
)

func openDB(dsn string) (*gorm.DB, error) {
	driverNameIdx := strings.Index(dsn, "://")
	if driverNameIdx < 0 {
		return nil, fmt.Errorf("invalid dsn")
	}
	var driverName string
	driverName, dsn = dsn[0:driverNameIdx], dsn[driverNameIdx+3:]

	if driverName == "mysql" && !strings.Contains(dsn, "parseTime") {
		if strings.Contains(dsn, "?") {
			dsn += "&parseTime=true"
		} else {
			dsn += "?parseTime=true"
		}
	}

	var dialector gorm.Dialector
	if openFunc, ok := drivers[strings.ToLower(driverName)]; !ok {
		return nil, fmt.Errorf("unsupported db type '%s'", driverName)
	} else {
		dialector = openFunc(dsn)
	}

	return gorm.Open(dialector, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			TablePrefix:   consts.TablePrefix,
		},
		Logger: gormLogger.New(log.Get(), gormLogger.Config{
			SlowThreshold:             200 * time.Millisecond,
			LogLevel:                  gormLogger.Warn,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		}),
	})
}

func Init(dsn string) error {
	var err error
	defaultDB, err = openDB(dsn)
	if err != nil {
		return errors.Wrap(err, "open database")
	}
	return nil
}

func Get() *gorm.DB {
	if defaultDB == nil {
		log.Get().Panic("db connection is not initialized")
	}
	return defaultDB
}
