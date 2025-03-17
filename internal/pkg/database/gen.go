package database

import (
	"gorm.io/gen"
	"gorm.io/gorm"
)

func NewGen(db *gorm.DB, path string, fc interface{}, models ...interface{}) {
	g := gen.NewGenerator(gen.Config{
		OutPath:      path,
		Mode:         gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
		WithUnitTest: true,
	})

	g.UseDB(db)
	g.ApplyBasic(models...)
	g.ApplyInterface(fc, models...)
	g.Execute()
}
