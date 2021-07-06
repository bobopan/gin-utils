module github.com/bobopan/gin-utils

go 1.13

require (
	github.com/onsi/ginkgo v1.16.3 // indirect
	github.com/onsi/gomega v1.13.0 // indirect
	github.com/op/go-logging v0.0.0-20160315200505-970db520ece7
	github.com/prometheus/client_golang v1.10.0
	github.com/rs/xid v1.3.0
	gopkg.in/redis.v5 v5.2.9
	github.com/natefinch/lumberjack v2.0.0+incompatible
	go.uber.org/zap v1.18.1
)

//replace go.uber.org/zap => github.com/uber-go/zap v1.18.1
