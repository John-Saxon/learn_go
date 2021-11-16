package main

import (
	"fmt"

	pg_query "github.com/pganalyze/pg_query_go"
)

func main() {
	// sql := `CREATE OR REPLACE FUNCTION func_name(res_table_name text) RETURNS void AS $$ DECLARE tag text = '青少年';  BEGIN   EXECUTE 'CREATE TABLE ' || res_table_name || '(id serial, value text);'; EXECUTE 'INSERT INTO ' || res_table_name || '(value) VALUES('|| quote_literal(tag) ||');';  END  $$  LANGUAGE PLPGSQL   RETURNS NULL ON NULL INPUT;`

	// sql := `CREATE OR REPLACE FUNCTION func_name(res_table_name text) RETURNS TABLE(id int, value text) AS $$ DECLARE tag text = '青少年';  BEGIN   EXECUTE 'CREATE TABLE ' || res_table_name || '(id serial, value text);'; EXECUTE 'INSERT INTO ' || res_table_name || '(value) VALUES('|| quote_literal(tag) ||');';  END  $$  LANGUAGE PLPGSQL   RETURNS NULL ON NULL INPUT;`
	// sql := `CREATE OR REPLACE FUNCTION func_name(res_table_name text) RETURNS SETOF RECORD AS $$ DECLARE tag text = '青少年';  BEGIN   EXECUTE 'CREATE TABLE ' || res_table_name || '(id serial, value text);'; EXECUTE 'INSERT INTO ' || res_table_name || '(value) VALUES('|| quote_literal(tag) ||');';  END  $$  LANGUAGE PLPGSQL   RETURNS NULL ON NULL INPUT;`
	sql := `CREATE OR REPLACE FUNCTION func_name(res_table_name text, b OUT text) RETURNS SETOF RECORD AS $$ DECLARE tag text = '青少年';  BEGIN   EXECUTE 'CREATE TABLE ' || res_table_name || '(id serial, value text);'; EXECUTE 'INSERT INTO ' || res_table_name || '(value) VALUES('|| quote_literal(tag) ||');';  END  $$  LANGUAGE PLPGSQL   RETURNS NULL ON NULL INPUT;`
	tree, err := pg_query.ParseToJSON(sql)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", tree)
}
