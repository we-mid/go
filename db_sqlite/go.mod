module github.com/we-task/Todo-as-a-Service/x/db_sqlite

go 1.21.1

replace (
	github.com/we-task/Todo-as-a-Service/x/db => ../db
	github.com/we-task/Todo-as-a-Service/x/util => ../util
)

require (
	github.com/glebarez/go-sqlite v1.22.0
	github.com/we-task/Todo-as-a-Service/x/db v0.0.0-00010101000000-000000000000
)

require (
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/google/uuid v1.5.0 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20230129092748-24d4a6f8daec // indirect
	github.com/we-task/Todo-as-a-Service/x/util v0.0.0-00010101000000-000000000000 // indirect
	golang.org/x/sys v0.15.0 // indirect
	modernc.org/libc v1.37.6 // indirect
	modernc.org/mathutil v1.6.0 // indirect
	modernc.org/memory v1.7.2 // indirect
	modernc.org/sqlite v1.28.0 // indirect
)
