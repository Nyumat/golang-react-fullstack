target: startBackend startFrontend

startBackend:
	cd backend && go run main.go

startFrontend:
	cd frontend && npm run dev

