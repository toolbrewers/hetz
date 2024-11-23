default: help

.PHONY: help
help: # Show help for each of the Makefile recipes.
	@grep -E '^[a-zA-Z0-9 -]+:.*#'  Makefile | sort | while read -r l; do printf "\033[1;32m$$(echo $$l | cut -f 1 -d':')\033[00m:$$(echo $$l | cut -f 2- -d'#')\n"; done

.PHONY: generate
generate: # Auto-generates code for .templ files and tailwindcss.
	templ generate
	tailwindcss build -o ./app/assets/stylesheets/tailwind.css --minify

.PHONY: migrate
migrate: # Creates a new migration in migrations with the given name.
	migrate create -ext sql -dir migrations -seq -digits 4 $(NAME)
