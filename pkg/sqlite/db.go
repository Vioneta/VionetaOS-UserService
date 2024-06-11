/*
 * @Author: LinkLeong link@vioneta.com
 * @Date: 2022-05-13 18:15:46
 * @LastEditors: LinkLeong
 * @LastEditTime: 2022-07-11 18:10:53
 * @Description:
 * @Website: https://www.vionetaos.io
 * Copyright (c) 2022 by icewhale, All Rights Reserved.
 */
package sqlite

import (
	"time"

	"github.com/Vioneta/VionetaOS-Common/utils/logger"
	"github.com/Vioneta/VionetaOS-UserService/model"
	"github.com/Vioneta/VionetaOS-UserService/pkg/utils/file"
	model2 "github.com/Vioneta/VionetaOS-UserService/service/model"
	"github.com/glebarez/sqlite"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var gdb *gorm.DB

func GetDb(dbPath string) *gorm.DB {
	if gdb != nil {
		return gdb
	}

	file.IsNotExistMkDir(dbPath)
	db, err := gorm.Open(sqlite.Open(dbPath+"/user.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	c, _ := db.DB()
	c.SetMaxIdleConns(10)
	c.SetMaxOpenConns(1)
	c.SetConnMaxIdleTime(time.Second * 1000)

	gdb = db

	err = db.AutoMigrate(model2.UserDBModel{}, model.EventModel{})
	if err != nil {
		logger.Error("check or create db error", zap.Any("error", err))
	}
	return db
}
