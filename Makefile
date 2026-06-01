PROTO_VERSION := v0.1.0
PROTO_PATH    := $(shell go env GOPATH)/pkg/mod/github.com/pav-dev98/pm-proto@$(PROTO_VERSION)
THIRD_PARTY   := $(PROTO_PATH)/third_party/googleapis
OUT_DIR       := ./docs

.PHONY: swagger clean-swagger

swagger:
	@mkdir -p $(OUT_DIR)
	protoc \
		-I $(PROTO_PATH) \
		-I $(THIRD_PARTY) \
		--openapiv2_out=$(OUT_DIR) \
		--openapiv2_opt=allow_merge=true \
		--openapiv2_opt=merge_file_name=api \
		$(PROTO_PATH)/auth/auth.proto

clean-swagger:
	rm -f $(OUT_DIR)/swagger.json
