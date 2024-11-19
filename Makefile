#.DEFAULT_GOAL := clean
all: async-entity-fetcher generate clean

async-entity-fetcher:	
	echo "wellcome aysync fetcher project"

generate:
	@echo "Creating empty text files..."
	touch file-{1..10}.txt

clean:
	@echo "Cleaning up..."
	rm *.txt