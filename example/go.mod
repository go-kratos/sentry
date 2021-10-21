module example

go 1.15

require (
	github.com/davecgh/go-spew v1.1.1
	github.com/fatih/color v1.11.0 // indirect
	github.com/getsentry/sentry-go v0.11.0
	github.com/go-kratos/kratos/cmd/kratos/v2 v2.0.0-20210515081852-24b1ca6bc3dc // indirect
	github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2 v2.0.0-20210515081852-24b1ca6bc3dc // indirect
	github.com/go-kratos/kratos/v2 v2.1.1
	github.com/go-kratos/sentry v0.0.0
	github.com/golang/glog v0.0.0-20210429001901-424d2337a529 // indirect
	github.com/golang/protobuf v1.5.2
	github.com/google/subcommands v1.2.0 // indirect
	github.com/google/wire v0.5.0
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.4.0 // indirect
	github.com/mgutz/ansi v0.0.0-20200706080929-d51e80ef957d // indirect
	github.com/spf13/cobra v1.1.3 // indirect
	go.opentelemetry.io/otel/metric v0.20.0 // indirect
	golang.org/x/crypto v0.0.0-20210513164829-c07d793c2f9a // indirect
	golang.org/x/mod v0.4.2 // indirect
	golang.org/x/net v0.0.0-20210525063256-abc453219eb5 // indirect
	golang.org/x/term v0.0.0-20210503060354-a79de5458b56 // indirect
	golang.org/x/text v0.3.6 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/genproto v0.0.0-20210805201207-89edb61ffb67
	google.golang.org/grpc v1.39.1
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.1.0 // indirect
	google.golang.org/protobuf v1.27.1
	gopkg.in/yaml.v2 v2.4.0
)

replace github.com/go-kratos/sentry => ./..
