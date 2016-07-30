package db

import "github.com/project-domino/project-domino/models"

// Setup initializes the database with empty tables of all the needed
// types.
// TODO weights on searchtext
func Setup() error {
	if !DB.HasTable(&models.User{}) {
		DB.CreateTable(&models.User{})
		DB.Exec("ALTER TABLE users ADD COLUMN searchtext TSVECTOR")
		DB.Exec("CREATE INDEX searchtext_user_gin ON users USING GIN(searchtext)")
		DB.Exec(`CREATE TRIGGER ts_searchtext_user
			BEFORE INSERT OR UPDATE ON users
			FOR EACH ROW EXECUTE PROCEDURE
			tsvector_update_trigger('searchtext', 'pg_catalog.english', 'user_name')`)
	}
	if !DB.HasTable(&models.Note{}) {
		DB.CreateTable(&models.Note{})
		DB.Exec("ALTER TABLE notes ADD COLUMN searchtext TSVECTOR")
		DB.Exec("CREATE INDEX searchtext_note_gin ON notes USING GIN(searchtext)")
	}
	if !DB.HasTable(&models.Collection{}) {
		DB.CreateTable(&models.Collection{})
		DB.Exec("ALTER TABLE collections ADD COLUMN searchtext TSVECTOR")
		DB.Exec("CREATE INDEX searchtext_collection_gin ON collections USING GIN(searchtext)")
	}
	if !DB.HasTable(&models.Tag{}) {
		DB.CreateTable(&models.Tag{})
		DB.Exec("ALTER TABLE tags ADD COLUMN searchtext TSVECTOR")
		DB.Exec("CREATE INDEX searchtext_tag_gin ON tags USING GIN(searchtext)")
		DB.Exec(`CREATE TRIGGER ts_searchtext_tag
			BEFORE INSERT OR UPDATE ON tags
			FOR EACH ROW EXECUTE PROCEDURE
			tsvector_update_trigger('searchtext', 'pg_catalog.english', 'name', 'description')`)
	}
	setupTable(&models.AuthToken{})
	setupTable(&models.Comment{})
	setupTable(&models.CollectionNote{})
	setupTable(&models.Email{})

	return DB.Error
}

// Creates a table for a specified struct if one doesn't exist
func setupTable(val interface{}) {
	if !DB.HasTable(val) {
		DB.CreateTable(val)
	}
}
