.PHONY: gen-frontend-http
gen-frontend-http:
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/home.proto --service frontend -module github.com/emory/gomall/app/frontend -I ../../idl