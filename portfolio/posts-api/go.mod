module github.com/kytra-app/posts-api

go 1.12

require (
	github.com/golang/protobuf v1.3.2
	github.com/kytra-app/bullbear-srv v1.0.0
	github.com/kytra-app/comments-srv v1.0.0
	github.com/kytra-app/helpers/authentication v1.0.0
	github.com/kytra-app/helpers/photos v0.0.0-00010101000000-000000000000
	github.com/kytra-app/helpers/textenhancer v0.0.0-00010101000000-000000000000
	github.com/kytra-app/posts-srv v1.0.0
	github.com/kytra-app/stocks-srv v0.0.0-00010101000000-000000000000
	github.com/kytra-app/users-srv v1.0.0
	github.com/micro/go-micro v1.8.1
	github.com/micro/go-plugins v1.2.0
	github.com/nats-io/nats.go v1.8.2-0.20190607221125-9f4d16fe7c2d // indirect
)

replace github.com/kytra-app/comments-srv => ../comments-srv

replace github.com/kytra-app/bullbear-srv => ../bullbear-srv

replace github.com/kytra-app/helpers/textenhancer => ../helpers/textenhancer

replace github.com/kytra-app/users-srv => ../users-srv

replace github.com/kytra-app/posts-srv => ../posts-srv

replace github.com/kytra-app/stocks-srv => ../stocks-srv

replace github.com/kytra-app/helpers/authentication => ../helpers/authentication

replace github.com/kytra-app/helpers/photos => ../helpers/photos

replace github.com/kytra-app/helpers/microgorm => ../helpers/microgorm

replace github.com/kytra-app/helpers/passwordhasher => ../helpers/passwordhasher

replace github.com/hashicorp/consul => github.com/hashicorp/consul v1.5.1