build:
	docker build -t historic:latest .
run:
	docker run -p 8080:8080 -e PORT=8080 historic:latest


build-and-run: build run
