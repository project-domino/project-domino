package db

import (
	"github.com/jinzhu/gorm"
	"github.com/project-domino/project-domino/models"
)

// Setup initializes the database with empty tables of all the needed
// types.
// TODO weights on searchtext
func Setup() error {
	// Create transaction
	tx := DB.Begin()
	defer tx.Commit()
	tx.Exec("SET TRANSACTION ISOLATION LEVEL SERIALIZABLE")

	if !tx.HasTable(&models.User{}) {
		tx.CreateTable(&models.User{})
		tx.Exec("ALTER TABLE users ADD COLUMN searchtext TSVECTOR")
		tx.Exec("CREATE INDEX searchtext_user_gin ON users USING GIN(searchtext)")
		tx.Exec(`CREATE TRIGGER ts_searchtext_user
			BEFORE INSERT OR UPDATE ON users
			FOR EACH ROW EXECUTE PROCEDURE
			tsvector_update_trigger('searchtext', 'pg_catalog.english', 'user_name')`)
	}
	if !tx.HasTable(&models.Note{}) {
		tx.CreateTable(&models.Note{})
		tx.Exec("ALTER TABLE notes ADD COLUMN searchtext TSVECTOR")
		tx.Exec("CREATE INDEX searchtext_note_gin ON notes USING GIN(searchtext)")
	}
	if !tx.HasTable(&models.Collection{}) {
		tx.CreateTable(&models.Collection{})
		tx.Exec("ALTER TABLE collections ADD COLUMN searchtext TSVECTOR")
		tx.Exec("CREATE INDEX searchtext_collection_gin ON collections USING GIN(searchtext)")
	}
	if !tx.HasTable(&models.Tag{}) {
		tx.CreateTable(&models.Tag{})
		tx.Exec("ALTER TABLE tags ADD COLUMN searchtext TSVECTOR")
		tx.Exec("CREATE INDEX searchtext_tag_gin ON tags USING GIN(searchtext)")
		tx.Exec(`CREATE TRIGGER ts_searchtext_tag
			BEFORE INSERT OR UPDATE ON tags
			FOR EACH ROW EXECUTE PROCEDURE
			tsvector_update_trigger('searchtext', 'pg_catalog.english', 'name', 'description')`)
	}
	setupTable(&models.AuthToken{}, tx)
	setupTable(&models.Comment{}, tx)
	setupTable(&models.CollectionNote{}, tx)
	setupTable(&models.Email{}, tx)

	err := tx.Error

	if err != nil {
		tx.Rollback()
	}

	return err
}

// Creates a table for a specified struct if one doesn't exist
func setupTable(val interface{}, db *gorm.DB) {
	if !db.HasTable(val) {
		db.CreateTable(val)
	}
}
