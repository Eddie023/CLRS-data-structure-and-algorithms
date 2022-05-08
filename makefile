run: 
	go run main.go

create: 
	mkdir -p $(FILENAME) && touch $(FILENAME)/$(FILENAME).go && echo "package ${FILENAME}" >> $(FILENAME)/$(FILENAME).go
