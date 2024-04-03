#!/bin/bash
# protoc --js_out=import_style=commonjs:. --grpc-web_out=import_style=typescript,mode=grpcwebtext:. protos/*.proto
# protoc --ts_out=./protos --proto_path=./protos protos/*.proto
npx protoc \
  --ts_out protos \
  --ts_opt long_type_string \
  --proto_path protos \
  protos/*/*.proto