package main

import (
	"database/sql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"

	_ "modernc.org/sqlite"
	//_ "github.com/mattn/go-sqlite3"
)

func Test_SelectClient_WhenOk(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "/Users/yong3rz/go/projects/ya_practicum/M3_SQL_DB/go-db-sql-query-testify/demo.db")
	require.NoError(t, err)
	defer db.Close()

	clientID := 1

	// напиши тест здесь
	client, err := selectClient(db, clientID)
	require.NoError(t, err)

	assert.Equal(t, clientID, client.ID)
	assert.NotEmpty(t, client.FIO)
	assert.NotEmpty(t, client.Login)
	assert.NotEmpty(t, client.Birthday)
	assert.NotEmpty(t, client.Email)
}

func Test_SelectClient_WhenNoClient(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "/Users/yong3rz/go/projects/ya_practicum/M3_SQL_DB/go-db-sql-query-testify/demo.db")
	require.NoError(t, err)
	defer db.Close()

	clientID := -1

	// напиши тест здесь
	client, err := selectClient(db, clientID)
	require.Equal(t, sql.ErrNoRows, err)

	assert.Empty(t, client.ID)
	assert.Empty(t, client.FIO)
	assert.Empty(t, client.Login)
	assert.Empty(t, client.Birthday)
	assert.Empty(t, client.Email)
}

func Test_InsertClient_ThenSelectAndCheck(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "/Users/yong3rz/go/projects/ya_practicum/M3_SQL_DB/go-db-sql-query-testify/demo.db")
	if err != nil {
		require.NoError(t, err)
	}
	defer db.Close()

	cl := Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}

	// напиши тест здесь
	clientId, err := insertClient(db, cl)

	cl.ID = clientId

	require.NotEmpty(t, clientId)

	client, err := selectClient(db, clientId)
	require.NoError(t, err)

	assert.Equal(t, cl, client)
}

func Test_InsertClient_DeleteClient_ThenCheck(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "/Users/yong3rz/go/projects/ya_practicum/M3_SQL_DB/go-db-sql-query-testify/demo.db")
	if err != nil {
		require.NoError(t, err)
	}
	defer db.Close()

	cl := Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}

	// напиши тест здесь
	clientId, err := insertClient(db, cl)
	require.NoError(t, err)
	require.NotEmpty(t, clientId)

	_, err = selectClient(db, clientId)
	require.NoError(t, err)

	err = deleteClient(db, clientId)
	require.NoError(t, err)

	_, err = selectClient(db, clientId)
	require.Equal(t, sql.ErrNoRows, err)
}
