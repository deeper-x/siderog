install:
	@echo "Installing siderog..."
	@go install 

run:
	@echo "Running siderog..."
	@siderog 2>&1 &

stop:
	@echo "Stopping siderog..."
	@pkill siderog 2>&1	