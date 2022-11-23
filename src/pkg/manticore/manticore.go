package manticore

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type ManticoreClient struct {
	*sqlx.DB
}

func NewClient(connectString string) (*ManticoreClient, error) {
	cfg, err := mysql.ParseDSN(connectString)
	if err != nil {
		return nil, err
	}
	cfg.InterpolateParams = true
	db, err := sqlx.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}
	return &ManticoreClient{DB: db}, nil
}

type Suggestion struct {
	Suggest  string `db:"suggest"`
	Distance int    `db:"distance"`
	Docs     int    `db:"docs"`
}

func (s *ManticoreClient) Suggest(word, index string, opts ...procedureOption) ([]Suggestion, error) {
	var suggestions []Suggestion
	var optsStr string
	if len(opts) > 0 {
		optsStr = optsToString(opts)
	}
	query := `CALL SUGGEST(?,?` + optsStr + `)`
	rows, err := s.Queryx(query, word, index)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var suggestion Suggestion
		err := rows.StructScan(&suggestion)
		if err != nil {
			return nil, err
		}
		suggestions = append(suggestions, suggestion)
	}
	return suggestions, nil
}

type KeywordData struct {
	Qpos       string `db:"qpos"`
	Tokenized  string `db:"tokenized"`
	Normalized string `db:"normalized"`
}

func (s *ManticoreClient) Keywords(word, index string) ([]KeywordData, error) {
	var keywords []KeywordData
	rows, err := s.Queryx(`CALL KEYWORDS(?,?)`, word, index)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var data KeywordData
	for rows.Next() {
		err := rows.StructScan(&data)
		if err != nil {
			return nil, err
		}
		keywords = append(keywords, data)
	}
	return keywords, nil
}
