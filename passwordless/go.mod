module gitee.com/we-mid/go/passwordless

go 1.21.1

replace (
	gitee.com/we-mid/go/cookiesession => ../cookiesession
	gitee.com/we-mid/go/session/v2 => ../session/v2
	gitee.com/we-mid/go/session_memory/v2 => ../session_memory/v2
)

require (
	gitee.com/we-mid/go/bec_http v0.0.0-20240924065140-a2727aea70d9
	gitee.com/we-mid/go/cookiesession v0.0.0-00010101000000-000000000000
	gitee.com/we-mid/go/mailer v0.0.0-20240924065140-a2727aea70d9
	gitee.com/we-mid/go/util v0.0.0-20240924065140-a2727aea70d9
)

require (
	gitee.com/we-mid/go/ratelimit v0.0.0-20240814094913-e4842229d27c // indirect
	gitee.com/we-mid/go/session/v2 v2.0.0-00010101000000-000000000000 // indirect
	gitee.com/we-mid/go/session_memory/v2 v2.0.0-00010101000000-000000000000 // indirect
	golang.org/x/time v0.5.0 // indirect
)
