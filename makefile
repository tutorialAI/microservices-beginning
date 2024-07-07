FOLDERS=kafka movie auth

up:
	$(foreach folder,$(FOLDERS),docker-compose -f $(folder)/docker-compose.yml up -d;)
down:
	$(foreach folder,$(FOLDERS),docker-compose -f $(folder)/docker-compose.yml down;)
