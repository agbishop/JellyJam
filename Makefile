# Makefile

dev:
	make -j 2 ui api

api:
	go run main.go

ui:
	cd client && npm run dev