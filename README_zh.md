[English Version](README.md)

## 说明

将一个创建表的SQL语句转换成Golang的ORM结构体的go函数。 例子：下面是一个创建`user`表的sql语句

```sql
CREATE TABLE `user`
(
    `id`            INT UNSIGNED PRIMARY KEY NOT NULL AUTO_INCREMENT COMMENT 'primary key',
    `ip_address`    INT                      NOT NULL DEFAULT 0 COMMENT 'ip_address',
    `nickname`      VARCHAR(128)             NOT NULL DEFAULT '' COMMENT 'user note',
    `description`   VARCHAR(256)             NOT NULL DEFAULT '' COMMENT 'user description',
    `creator_email` VARCHAR(64)              NOT NULL DEFAULT '' COMMENT 'creator email',
    `created_at`    TIMESTAMP                NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'create time',
    `deleted_at`    TIMESTAMP                NULL     DEFAULT NULL COMMENT 'delete time',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4 COMMENT ='user table';
```

函数`Parse`将其转化成下面的代码。其中，包名是可以选择的。

```go
package main

import (
	"time"
)

// User ENGINE=InnoDB auto_increment=1 default charset=utf8mb4 comment='user table'
type User struct {
	Id           uint      `gorm:"column:id;type:int;not null;autoIncrement;primaryKey;comment:'primary key'"`
	IpAddress    int       `gorm:"column:ip_address;type:int;not null;default:0;comment:'ip_address'"`
	Nickname     string    `gorm:"column:nickname;type:varchar(128);not null;default:'';comment:'user note'"`
	Description  string    `gorm:"column:description;type:varchar(256);not null;default:'';comment:'user description'"`
	CreatorEmail string    `gorm:"column:creator_email;type:varchar(64);not null;default:'';comment:'creator email'"`
	CreatedAt    time.Time `gorm:"column:created_at;type:timestamp;not null;default:current_timestamp;comment:'create time'"`
	DeletedAt    time.Time `gorm:"column:deleted_at;type:timestamp;default:null;comment:'delete time'"`
}
```

## 下划线命名改为驼峰式命名

在SQL的命名规范中，字段的命名一般都是下划线分隔的,例如`ip_address`。而Golang的`struct`的字段的命名是驼峰式的。
`Parse`会将其字段命名转化为驼峰式的。
