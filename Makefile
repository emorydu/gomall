.PHONY: gen-frontend-http
gen-frontend-http:
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/auth_page.proto --service frontend -module github.com/emory/gomall/app/frontend -I ../../idl

.PHONY: gen-user-rpc-client
gen-user-rpc-client:
	@cd rpc_gen && cwgo client --type RPC --service user --module github.com/emory/gomall/rpc_gen -I ../idl --idl ../idl/user.proto

.PHONY: gen-user-rpc-server
gen-user-rpc-server:
	@cd app/user && cwgo server --type RPC --service user --module github.com/emory/gomall/app/user --pass "-use github.com/emory/gomall/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/user.proto

.PHONY: gen-product-rpc-client
gen-product-rpc-client:
	@cd rpc_gen && cwgo client --type RPC --service product --module github.com/emory/gomall/rpc_gen -I ../idl --idl ../idl/product.proto

.PHONY: gen-product-rpc-server
gen-product-rpc-server:
	@cd app/product && cwgo server --type RPC --service product --module github.com/emory/gomall/app/product --pass "-use github.com/emory/gomall/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/product.proto