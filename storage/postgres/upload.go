package postgres

import (
	"database/sql"
	"errors"
	"essy_travel/models"
	"essy_travel/pkg/helpers"
	"fmt"
	"strings"

	"github.com/spf13/cast"
)

type UploadRepo struct {
	db *sql.DB
}

func NewUploadRepo(db *sql.DB) *UploadRepo {
	return &UploadRepo{
		db: db,
	}
}

func (u *UploadRepo) Upload(req *models.UploadRequest) error {

	if _, ok := models.TableEntity[req.TableSlug]; !ok {
		return errors.New("not found table")
	}

	var (
		entity = models.TableEntity[req.TableSlug]
		query  = "INSERT INTO " + req.TableSlug + "("
		column = strings.Join(entity, ",")
	)
	query += column + ") VALUES"

	if helpers.Contains(entity, `"offset"`) {
		index := helpers.FindIndex(entity, `"offset"`)
		entity[index] = "offset"

	}

	for _, data := range req.Data {
		var obj = cast.ToStringMap(data)

		query += "("
		for ind, key := range entity {
			if _, ok := obj[key]; ok {
				if ind == len(entity)-1 {
					query += fmt.Sprintf(`'%s'`, cast.ToString(obj[key]))
					break
				}

				query += fmt.Sprintf(`'%s',`, cast.ToString(obj[key]))

				if obj["guuid"] == "US" {
					fmt.Println(obj)
					break

				}
			} else {
				if ind == len(entity)-1 {
					fmt.Println(":::", ind)
					query += `''`
					break
				}

				query += `'',`
			}
		}
		query += "),"
	}
	query = query[:len(query)-1]

	// fmt.Println("query:", query)

	_, err := u.db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
