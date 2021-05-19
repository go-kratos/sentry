# [WIP] Sentry middleware for Kratos
**This project is a Work In Progress**

## Quick Start
You could check the full demo in example folder.
```go
// Step 1: 
// init sentry in the entry of your application
import "github.com/getsentry/sentry-go"
sentry.Init(sentry.ClientOptions{
		Dsn: "<your dsn>",
		AttachStacktrace: true, // recommended
})


// Step 2: 
// set middleware
import 	"github.com/go-kratos/sentry"
// for http server
m := http.Middleware(
    middleware.Chain(
        recovery.Recovery(),
        sentrykratos.Server(), // must after Recovery middleware
        tracing.Server(),
        logging.Server(logger),
    ),
)
// for grpc server
var opts = []grpc.ServerOption{
     grpc.Middleware(
         middleware.Chain(
             recovery.Recovery(),
             sentrykratos.Server(), // must after Recovery middleware
             tracing.Server(),
             logging.Server(logger),
         ),
     ),
 }

// Then, the framework will report events to Sentry when your trigger panics.
// Or your can push events to Sentry manually
```

## Reference
* [https://docs.sentry.io/platforms/go/](https://docs.sentry.io/platforms/go/)