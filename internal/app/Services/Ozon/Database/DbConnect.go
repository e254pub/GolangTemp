package Database

import (
	"context"
	"log"
	conf "main/configs/connects"
	"main/internal/app/Services/Ozon/Structures"
	"main/pkg/utils"
	"strconv"
)

var conn *pgx.Conn

func init() {
	conn = conf.PostgresConnect()
}

type Row struct {
	ExternalId string
}

func GetDbAttrs(Query string, ShopId int, Limit interface{}) []int {
	rows, err := conn.Query(context.Background(), Query, ShopId, Limit)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var rowSlice []int
	for rows.Next() {
		var r Row
		err := rows.Scan(&r.ExternalId)
		if err != nil {
			log.Fatal(err)
		}
		AttrId, _ := strconv.Atoi(utils.TrimLeftChars(r.ExternalId, 2))
		rowSlice = append(rowSlice, AttrId)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return rowSlice
}

// ProductRecords добавляет продукты
func ProductRecords(records []Structures.ProductStruct) (int64, error) {
	res, err := conn.CopyFrom(context.Background(), pgx.Identifier{"схема", "таблица"},
		[]string{"поля"},
		pgx.CopyFromSlice(len(records), func(i int) ([]interface{}, error) {
			return []interface{}{}, nil
		}),
	)
	return res, err
}

// AttributeRecords добавляет атрибуты
func AttributeRecords(records []Structures.AttributeStruct) (int64, error) {
	res, err := conn.CopyFrom(context.Background(), pgx.Identifier{"схема", "таблица"},
		[]string{"поля"},
		pgx.CopyFromSlice(len(records), func(i int) ([]interface{}, error) {
			return []interface{}{}, nil
		}),
	)
	return res, err
}

// UniqueAttrs возвращает уникальные атрибуты
func UniqueAttrs(attributes []Structures.AttributeStruct) []Structures.AttributeStruct {
	var unique []Structures.AttributeStruct
sampleLoop:
	for _, v := range attributes {
		for i, u := range unique {
			if //перебор значений

			{
				unique[i] = v
				continue sampleLoop
			}
		}
		unique = append(unique, v)
	}
	return unique
}
