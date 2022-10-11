package documentsdb

import (
	"context"
	"fmt"
	"github.com/lib/pq"
	"time"
	"web-server/pkg/db"
	"web-server/pkg/models/documentsModels"
)

func InsertDocuments(documents *documentsModels.Document, json string, pathDirectory string) error {
	db, err := db.InitDB()
	if err != nil {
		return fmt.Errorf("InitDB failed: %v", err)
	}
	_, err = db.Query(`INSERT INTO "documents"(
						"name",
                      	"file",
                      	"public",
                      	"token",
                      	"mime",
                      	"grant",
                      	"json",
                      	"directory",
                      	"created"
						 )
						 VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)`, documents.Name, documents.File, documents.Public,
		documents.Token, documents.Mime, pq.Array(documents.Grant), json, pathDirectory, time.Now())
	if err != nil {
		return err
	}
	return nil
}

func GetDocument(id int) (*documentsModels.Document, error) {
	conn, err := db.InitDB()
	if err != nil {
		return nil, fmt.Errorf("InitDB failed: %v", err)
	}

	stmt := `SELECT * FROM documents WHERE "id" = $1`

	row := conn.QueryRow(stmt, id)

	doc := &documentsModels.Document{}

	err = row.Scan(&doc.Id, &doc.Name, &doc.File, &doc.Public,
		&doc.Token, &doc.Mime, pq.Array(&doc.Grant), &doc.Json, &doc.Directory, &doc.Created)
	if err != nil {
		return nil, fmt.Errorf("Scan failed: %v", err)
	}

	return doc, nil
}

func GetDocumentsList(ctx context.Context, info documentsModels.Pass) ([]*documentsModels.Document, error) {
	conn, err := db.InitDB()
	if err != nil {
		return nil, err
	}

	stmt := fmt.Sprintf(`SELECT * FROM documents WHERE %v = '%v' ORDER BY name ASC, created ASC LIMIT $1`, info.Key, info.Value)

	rows, err := conn.QueryContext(ctx, stmt, info.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	list := []*documentsModels.Document{}
	for rows.Next() {
		doc := documentsModels.Document{}
		if err := rows.Scan(&doc.Id, &doc.Name, &doc.File, &doc.Public,
			&doc.Token, &doc.Mime, pq.Array(&doc.Grant), &doc.Json, &doc.Directory, &doc.Created); err != nil {
			return nil, err
		}

		list = append(list, &doc)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return list, err
}

func DeleteDocument(id int) (string, error) {
	conn, err := db.InitDB()
	if err != nil {
		return "", fmt.Errorf("InitDB failed: %v", err)
	}

	var directory string

	err = conn.QueryRow(`DELETE FROM documents WHERE "id" = $1 RETURNING "directory" `, id).Scan(&directory)
	if err != nil {
		return "", fmt.Errorf("DELETE failed: %w", err)
	}

	return directory, nil
}
