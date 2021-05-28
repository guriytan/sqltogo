[![Go Report Card](https://goreportcard.com/badge/github.com/liangyaopei/sqltogo)](https://goreportcard.com/report/github.com/liangyaopei/sqltogo)
[![GoDoc](https://godoc.org/github.com/liangyaopei/sqltogo?status.svg)](http://godoc.org/github.com/liangyaopei/sqltogo)
[中文版说明](./README_zh.md)

## Description

This repository provide a way to convert SQL create statement to Golang struct (ORM struct) by parsing the SQL
statement. For example, with the input

```mysql
CREATE TABLE `user`
(
    `id`            INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'primary key',
    `ip_address`    INT          NOT NULL DEFAULT 0 COMMENT 'ip_address',
    `nickname`      VARCHAR(128) NOT NULL DEFAULT '' COMMENT 'user note',
    `description`   VARCHAR(256) NOT NULL DEFAULT '' COMMENT 'user description',
    `creator_email` VARCHAR(64)  NOT NULL DEFAULT '' COMMENT 'creator email',
    `created_at`    TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'create time',
    `deleted_at`    TIMESTAMP    NULL     DEFAULT NULL COMMENT 'delete time',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4 COMMENT ='user table';
```

the function `Parse` will convert it as follows, with the `package` name is optional.

```go
package main

import (
	"time"
)

// User ENGINE=InnoDB auto_increment=1 default charset=utf8mb4 comment='user table'
type User struct {
	Id           uint      `gorm:"column:id;type:int;not null;autoIncrement;primaryKey;comment:primary key"`
	IpAddress    int       `gorm:"column:ip_address;type:int;not null;default:0;comment:ip_address"`
	Nickname     string    `gorm:"column:nickname;type:varchar(128);not null;default:'';comment:user note"`
	Description  string    `gorm:"column:description;type:varchar(256);not null;default:'';comment:user description"`
	CreatorEmail string    `gorm:"column:creator_email;type:varchar(64);not null;default:'';comment:creator email"`
	CreatedAt    time.Time `gorm:"column:created_at;type:timestamp;not null;default:current_timestamp;comment:create time"`
	DeletedAt    time.Time `gorm:"column:deleted_at;type:timestamp;default:null;comment:delete time"`
}
```

## Snake case To Camel case

In SQL, the naming convention of a filed is snake case, such as `ip_address`, while in Golang, the naming convention of
struct's field is Camel case. So the `Parse` function will convert snake case to Camel case.
