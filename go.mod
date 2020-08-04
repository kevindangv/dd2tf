module github.com/devops-works/dd2tf

require (
	github.com/DataDog/datadog-api-client-go v1.0.0-beta.6
	github.com/cenkalti/backoff v2.2.1+incompatible // indirect
	github.com/containous/go-bindata v1.0.0 // indirect
	github.com/davecgh/go-spew v1.1.1
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/pflag v1.0.5
	github.com/zorkian/go-datadog-api v2.27.0+incompatible
	golang.org/dl v0.0.0-20200611200201-72429b14455f // indirect
)

replace github.com/DataDog/datadog-api-client-go => /home/kevin/go/src/github.com/DataDog/datadog-api-client-go

go 1.13
