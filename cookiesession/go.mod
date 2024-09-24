module gitee.com/we-mid/go/cookiesession

go 1.21.1

replace (
	gitee.com/we-mid/go/session/v2 => ../session/v2
	gitee.com/we-mid/go/session_memory/v2 => ../session_memory/v2
)

require gitee.com/we-mid/go/session_memory/v2 v2.0.0-00010101000000-000000000000

require (
	gitee.com/we-mid/go/session/v2 v2.0.0-00010101000000-000000000000 // indirect
	gitee.com/we-mid/go/util v0.0.0-20240918131137-e86e22a73e87 // indirect
)
