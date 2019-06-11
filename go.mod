module github.com/hashgard/hashgard

go 1.12

require (
	github.com/btcsuite/btcutil v0.0.0-20190207003914-4c204d697803 // indirect
	github.com/cosmos/cosmos-sdk v0.34.7
	github.com/gogo/protobuf v1.2.1 // indirect
	github.com/gorilla/mux v1.7.0
	github.com/mattn/go-isatty v0.0.7 // indirect
	github.com/otiai10/copy v0.0.0-20180813032824-7e9a647135a1
	github.com/prometheus/procfs v0.0.0-20190306233201-d0f344d83b0c // indirect
	github.com/rakyll/statik v0.1.4
	github.com/rcrowley/go-metrics v0.0.0-20181016184325-3113b8401b8a // indirect
	github.com/spf13/cobra v0.0.3
	github.com/spf13/viper v1.0.3
	github.com/stretchr/testify v1.3.0
	github.com/tendermint/go-amino v0.14.1
	github.com/tendermint/tendermint v0.31.5
	golang.org/x/crypto v0.0.0 // indirect
	google.golang.org/grpc v1.19.1 // indirect

)

replace (
	github.com/cosmos/cosmos-sdk => github.com/hashgard/cosmos-sdk v0.34.4-hashgard.0.20190606102350-4249afd167f2
	golang.org/x/crypto => github.com/tendermint/crypto v0.0.0-20180820045704-3764759f34a5
)
