package main

import (
	"fmt"
	"log"

	"github.com/rbrick/mc-avatar-go/mojang"
	"github.com/rbrick/mc-avatar-go/uuidcache"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func main() {
	// TODO(ryan): load the database based on environment variables, move to init
	db, err := gorm.Open(sqlite.Open("avatar.db"), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(uuidcache.Name{}, uuidcache.CacheEntry{})

	uuidCache := uuidcache.NewSqlUuidCache(db)

	engine := gin.New()

	v1apiGroup := engine.Group("v1")
	{

		// GET /v1/name/:uuid
		v1apiGroup.GET("/name/:uuid", func(ctx *gin.Context) {
			uuid := mojang.Uuid(ctx.Param("uuid")).WithDashes()

			name, err := uuidCache.GetLatestName(uuid)

			if err != nil {

				if err == gorm.ErrRecordNotFound {
					// defer to mojang
					profile, err := mojang.LookupProfile(mojang.Uuid(uuid).WithoutDashes())

					if err != nil {

						log.Println("error:", err)
						ctx.JSON(500, gin.H{
							"message": "failed to lookup UUID from mojang",
						})
						return
					}

					uuidCache.Put(profile.Name, profile.ID.WithDashes())

					ctx.JSON(200, gin.H{
						"name": profile.Name,
					})
					return
				}

				// return early
				log.Println("error:", err)
				ctx.JSON(500, gin.H{
					"message": fmt.Sprintf("failed to lookup name for UUID: %s", uuid),
				})
				return
			}

			ctx.JSON(200, gin.H{
				"name": name,
			})

		})

		// GET /v1/uuid/:name
		v1apiGroup.GET("/uuid/:name", func(ctx *gin.Context) {
			name := ctx.Param("name")

			uuid, err := uuidCache.GetUUID(name)

			if err != nil {

				if err == gorm.ErrRecordNotFound {
					// defer to mojang
					profile, err := mojang.LookupUuid(name)

					if err != nil {

						log.Println("error:", err)
						ctx.JSON(500, gin.H{
							"message": "failed to lookup name from mojang",
						})
						return
					}

					uuidCache.Put(profile.Name, profile.ID.WithDashes())

					ctx.JSON(200, gin.H{
						"uuid": profile.ID.WithDashes(),
					})
					return
				}

				// return early
				log.Println("error:", err)
				ctx.JSON(500, gin.H{
					"message": fmt.Sprintf("failed to lookup UUID for name: %s", name),
				})
				return
			}

			ctx.JSON(200, gin.H{
				"uuid": uuid,
			})
		})

		engine.Run()

	}
}
