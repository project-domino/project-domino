package db

import (
	"github.com/jinzhu/gorm"
	"github.com/project-domino/project-domino/errors"
	"github.com/project-domino/project-domino/models"
)

// GetTags gets tags from the db whose ids are in a specified slice
func GetTags(ids []uint) ([]models.Tag, error) {
	var tags []models.Tag
	if err := DB.Where("id in (?)", ids).
		Find(&tags).Error; err != nil && err != gorm.ErrRecordNotFound {
		return tags, err
	}
	return tags, nil
}

// VerifyNotes checks if notes with given ids exist
func VerifyNotes(ids []uint) error {
	for _, id := range ids {
		var note models.Note
		DB.Where("id = ?", id).Find(&note)
		if note.ID == 0 {
			return errors.NoteNotFound
		}
	}
	return nil
}

// LoadCollectionNotes loads the notes into a collection
func LoadCollectionNotes(c *models.Collection) error {
	// Find collection note relationships
	var collectioNotes []models.CollectionNote
	if err := DB.Preload("Note").
		Where("collection_id = ?", c.ID).
		Order("order").
		Find(&collectioNotes).Error; err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	// Load notes into collection
	for _, collectionNote := range collectioNotes {
		c.Notes = append(c.Notes, collectionNote.Note)
	}

	return nil
}

// UpdateNoteSearchText sets the searchtext field for a note with the given id
func UpdateNoteSearchText(id uint) error {
	if err := DB.Exec(`UPDATE notes n SET searchtext = result_vector
		FROM (select note_id, to_tsvector('english',note_title) ||
		to_tsvector('english',note_description) ||
		to_tsvector('english',tag_name) ||
		to_tsvector('english', tag_description) AS
		result_vector FROM
			(SELECT
				n.id AS note_id,
				n.title AS note_title,
				n.description AS note_description,
				string_agg(t.name,' ') AS tag_name,
				string_agg(t.description, ' ') AS tag_description
				FROM
					tags t JOIN
					note_tag nt ON
					t.id = nt.tag_id JOIN
					notes n on n.id=nt.note_id
				WHERE n.id = ?
				GROUP BY n.id, n.title, n.description
			) sub_q
		)
	r WHERE n.id = r.note_id;`, id).Error; err != nil {
		return err
	}
	return nil
}
