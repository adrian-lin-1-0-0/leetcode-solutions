create:
	@read -p "Enter Problem Name:" problem; \
	echo "Creating $$problem"; \
	mkdir -p "$$problem"; \
	touch "$$problem"/README.md; \
	echo "# [$$problem]()" > "$$problem"/README.md; \