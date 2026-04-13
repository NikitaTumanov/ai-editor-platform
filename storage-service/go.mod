module github.com/NikitaTumanov/ai-editor-platform/storage-service

go 1.25.4

require (
	github.com/NikitaTumanov/ai-editor-platform/protos v0.0.0-20260411222347-62d91636e633
	github.com/jackc/pgx/v5 v5.8.0
	go.uber.org/zap v1.27.1
	google.golang.org/grpc v1.80.0
)

replace github.com/NikitaTumanov/ai-editor-platform/protos => ../protos

require (
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	golang.org/x/net v0.51.0 // indirect
	golang.org/x/sync v0.19.0 // indirect
	golang.org/x/sys v0.41.0 // indirect
	golang.org/x/text v0.34.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260120221211-b8f7ae30c516 // indirect
	google.golang.org/protobuf v1.36.11 // indirect
)
