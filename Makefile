build:
	docker build --build-arg COLOR=yellow -t flower:yellow .
	docker build --build-arg COLOR=blue   -t flower:blue .
	docker build --build-arg COLOR=red   -t flower:red .

	
docx:
	pandoc -o README.docx -f markdown -t docx README.md
