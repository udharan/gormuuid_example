package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/ubgo/gormuuid"
	"github.com/ubgo/goutil"
	"github.com/ubgo/gouuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	ID            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name          string
	TagIDs        gormuuid.UUIDArray  `gorm:"type:uuid[]"`
	TagIDsPointer *gormuuid.UUIDArray `gorm:"type:uuid[]"`
}

func main() {

	dsn := "host=localhost user=postgres password=postgres dbname=testdb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	// fmt.Println(db)

	db.Migrator().DropTable(&User{})
	db.AutoMigrate(&User{})

	// imageId := uuid.New()

	var tagIds gormuuid.UUIDArray = []uuid.UUID{gouuid.ParseToDefault("c4ba81a1-9e57-4e60-b811-2860136ab803"), gouuid.ParseToDefault("0f49cc97-fba6-41c7-9063-834ecccb681e")}

	db.Create(&User{
		TagIDs:        tagIds,
		TagIDsPointer: &tagIds,
	})

	var users []User
	tagId, _ := uuid.Parse("c4ba81a1-9e57-4e60-b811-2860136ab803")
	db.Model(&User{}).Where("tag_ids && ?", pq.Array([]uuid.UUID{tagId})).First(&users)
	goutil.PrintToJSON(&users)

	// Output
	/*
		[
			{
				"ID": "8f050baa-5a3b-4635-93a8-184c48ac4f3d",
				"Name": "",
				"TagIDs": [
					"c4ba81a1-9e57-4e60-b811-2860136ab803",
					"0f49cc97-fba6-41c7-9063-834ecccb681e"
				],
				"TagIDsPointer": [
					"c4ba81a1-9e57-4e60-b811-2860136ab803",
					"0f49cc97-fba6-41c7-9063-834ecccb681e"
				]
			}
		]
	*/
}
