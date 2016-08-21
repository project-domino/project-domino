package db

// UpdateNoteSearchText sets the searchtext field for a note with the given id
func UpdateNoteSearchText(id uint) error {
	if err := DB.Exec(`UPDATE notes n SET searchtext = result_vector
		FROM (select note_id, to_tsvector('english',note_title) ||
		to_tsvector('english', note_description) ||
		to_tsvector('english', tag_name) ||
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

// UpdateCollectionSearchText sets the searchtext field for a collection with the given id
func UpdateCollectionSearchText(id uint) error {
	if err := DB.Exec(`UPDATE collections c SET searchtext = result_vector
		FROM (select collection_id, to_tsvector('english',collection_title) ||
		to_tsvector('english', collection_description) ||
		to_tsvector('english', tag_name) ||
		to_tsvector('english', tag_description) AS
		result_vector FROM
			(SELECT
				c.id AS collection_id,
				c.title AS collection_title,
				c.description AS collection_description,
				string_agg(t.name,' ') AS tag_name,
				string_agg(t.description, ' ') AS tag_description
				FROM
					tags t JOIN
					collection_tag ct ON
					t.id = ct.tag_id JOIN
					collections c on c.id=ct.collection_id
				WHERE c.id = ?
				GROUP BY c.id, c.title, c.description
			) sub_q
		)
	r WHERE c.id = r.collection_id;`, id).Error; err != nil {
		return err
	}
	return nil
}
