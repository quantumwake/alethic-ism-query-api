go install github.com/swaggo/swag/cmd/swag@latest
#swag init --parseDependency --parseInternal
swag init --parseDependency --parseInternal --dir ./.,./pkg/model,./pkg/dsl --output ./docs
sed '/LeftDelim:/d; /RightDelim:/d' ./docs/docs.go > ./docs/docs.go.new
mv ./docs/docs.go.new ./docs/docs.go
