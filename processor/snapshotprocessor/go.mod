module github.com/observiq/bindplane-otel-collector/processor/snapshotprocessor

go 1.22.7

require (
	github.com/observiq/bindplane-otel-collector/internal/report v1.68.0
	github.com/open-telemetry/opamp-go v0.17.0
	github.com/open-telemetry/opentelemetry-collector-contrib/extension/opampcustommessages v0.116.0
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/golden v0.116.0
	github.com/stretchr/testify v1.10.0
	go.opentelemetry.io/collector/component v0.116.0
	go.opentelemetry.io/collector/consumer v1.22.0
	go.opentelemetry.io/collector/consumer/consumertest v0.116.0
	go.opentelemetry.io/collector/pdata v1.22.0
	go.opentelemetry.io/collector/processor v0.116.0
	go.opentelemetry.io/collector/processor/processortest v0.116.0
	go.uber.org/zap v1.27.0
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil v0.116.0 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/rogpeppe/go-internal v1.11.0 // indirect
	go.opentelemetry.io/collector/component/componentstatus v0.116.0 // indirect
	go.opentelemetry.io/collector/component/componenttest v0.116.0 // indirect
	go.opentelemetry.io/collector/config/configtelemetry v0.116.0 // indirect
	go.opentelemetry.io/collector/consumer/xconsumer v0.116.0 // indirect
	go.opentelemetry.io/collector/pdata/pprofile v0.116.0 // indirect
	go.opentelemetry.io/collector/pdata/testdata v0.116.0 // indirect
	go.opentelemetry.io/collector/pipeline v0.116.0 // indirect
	go.opentelemetry.io/collector/processor/xprocessor v0.116.0 // indirect
	go.opentelemetry.io/otel v1.32.0 // indirect
	go.opentelemetry.io/otel/metric v1.32.0 // indirect
	go.opentelemetry.io/otel/sdk v1.32.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.32.0 // indirect
	go.opentelemetry.io/otel/trace v1.32.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/net v0.29.0 // indirect
	golang.org/x/sys v0.27.0 // indirect
	golang.org/x/text v0.18.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240903143218-8af14fe29dc1 // indirect
	google.golang.org/grpc v1.68.1 // indirect
	google.golang.org/protobuf v1.35.2 // indirect
)

replace github.com/observiq/bindplane-otel-collector/internal/report => ../../internal/report