#!/usr/bin/env bash

protoc --plugin=protoc-gen-ts=frontend/web/node_modules/.bin/protoc-gen-grpc-web --js_out=import_style=typescirpt,binary:frontend/web --ts_out=service=grpc-web:frontend/web backend/protos/post.proto
