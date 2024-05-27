module github.com/we-task/Todo-as-a-Service/x/session_memory

go 1.21.1

replace (
	github.com/we-task/Todo-as-a-Service/x/session => ../session
	github.com/we-task/Todo-as-a-Service/x/util => ../util
)

require (
	github.com/we-task/Todo-as-a-Service/x/session v0.0.0-00010101000000-000000000000
	github.com/we-task/Todo-as-a-Service/x/util v0.0.0-00010101000000-000000000000
)
