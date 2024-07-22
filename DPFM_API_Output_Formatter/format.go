package dpfm_api_output_formatter

import (
	"data-platform-api-article-type-reads-rmq-kube/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)

func ConvertToArticleType(rows *sql.Rows) (*[]ArticleType, error) {
	defer rows.Close()
	articleType := make([]ArticleType, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.ArticleType{}

		err := rows.Scan(
			&pm.ArticleType,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &articleType, nil
		}

		data := pm
		articleType = append(articleType, ArticleType{
			ArticleType:			data.ArticleType,
			CreationDate:			data.CreationDate,
			LastChangeDate:			data.LastChangeDate,
			IsMarkedForDeletion:	data.IsMarkedForDeletion,
		})
	}

	return &articleType, nil
}

func ConvertToText(rows *sql.Rows) (*[]Text, error) {
	defer rows.Close()
	text := make([]Text, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Text{}

		err := rows.Scan(
			&pm.ArticleType,
			&pm.Language,
			&pm.ArticleTypeName,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &text, err
		}

		data := pm
		text = append(text, Text{
			ArticleType:     		data.ArticleType,
			Language:          		data.Language,
			ArticleTypeName:		data.ArticleTypeName,
			CreationDate:			data.CreationDate,
			LastChangeDate:			data.LastChangeDate,
			IsMarkedForDeletion:	data.IsMarkedForDeletion,
		})
	}

	return &text, nil
}
