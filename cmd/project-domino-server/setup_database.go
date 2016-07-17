package main

import (
	"github.com/jinzhu/gorm"
	"github.com/project-domino/project-domino/models"
)

// SetupDatabase initializes the database with empty tables of all the needed
// types.
func SetupDatabase(db *gorm.DB) error {
	if !db.HasTable(&models.User{}) {
		db.CreateTable(&models.User{})
		db.Exec("ALTER TABLE users ADD COLUMN searchtext TSVECTOR")
		db.Exec("CREATE INDEX searchtext_user_gin ON users USING GIN(searchtext)")
		db.Exec(`CREATE TRIGGER ts_searchtext_user
			BEFORE INSERT OR UPDATE ON users
			FOR EACH ROW EXECUTE PROCEDURE
			tsvector_update_trigger('searchtext', 'pg_catalog.english', 'user_name')`)
	}
	if !db.HasTable(&models.Note{}) {
		db.CreateTable(&models.Note{})
		db.Exec("ALTER TABLE notes ADD COLUMN searchtext TSVECTOR")
		db.Exec("CREATE INDEX searchtext_note_gin ON notes USING GIN(searchtext)")
		db.Exec(`CREATE TRIGGER ts_searchtext_note
			BEFORE INSERT OR UPDATE ON notes
			FOR EACH ROW EXECUTE PROCEDURE
			tsvector_update_trigger('searchtext', 'pg_catalog.english', 'title', 'description')`)
	}
	if !db.HasTable(&models.Collection{}) {
		setupTable(db, &models.Collection{})
		db.Exec("ALTER TABLE collections ADD COLUMN searchtext TSVECTOR")
		db.Exec("CREATE INDEX searchtext_collection_gin ON collections USING GIN(searchtext)")
		db.Exec(`CREATE TRIGGER ts_searchtext_collection
			BEFORE INSERT OR UPDATE ON collections
			FOR EACH ROW EXECUTE PROCEDURE
			tsvector_update_trigger('searchtext', 'pg_catalog.english', 'title', 'description')`)
	}
	if !db.HasTable(&models.University{}) {
		setupTable(db, &models.University{})
		db.Exec("ALTER TABLE universities ADD COLUMN searchtext TSVECTOR")
		db.Exec("CREATE INDEX searchtext_university_gin ON universities USING GIN(searchtext)")
		db.Exec(`CREATE TRIGGER ts_searchtext_university
			BEFORE INSERT OR UPDATE ON universities
			FOR EACH ROW EXECUTE PROCEDURE
			tsvector_update_trigger('searchtext', 'pg_catalog.english', 'name', 'short_name')`)
	}
	if !db.HasTable(&models.Tag{}) {
		setupTable(db, &models.Tag{})
		db.Exec("ALTER TABLE tags ADD COLUMN searchtext TSVECTOR")
		db.Exec("CREATE INDEX searchtext_tag_gin ON tags USING GIN(searchtext)")
		db.Exec(`CREATE TRIGGER ts_searchtext_tag
			BEFORE INSERT OR UPDATE ON tags
			FOR EACH ROW EXECUTE PROCEDURE
			tsvector_update_trigger('searchtext', 'pg_catalog.english', 'name', 'description')`)
	}
	setupTable(db, &models.AuthToken{})
	setupTable(db, &models.Comment{})
	setupTable(db, &models.CollectionNote{})

	return db.Error
}

// Creates a table for a specified struct if one doesn't exist
func setupTable(db *gorm.DB, val interface{}) {
	if !db.HasTable(val) {
		db.CreateTable(val)
	}
}
