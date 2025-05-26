.PHONY: tag
tag:
	@read -p "Enter version (e.g., 1.0.0): " version; \
	if [ -z "$$version" ]; then \
		echo "Version is required. Exiting..."; \
		exit 1; \
	fi; \
	read -p "Do you want to create the tag v$$version? (y/n): " confirm; \
	if [ "$$confirm" != "y" ] && [ "$$confirm" != "Y" ]; then \
		echo "Tag creation aborted."; \
		exit 0; \
	fi; \
	read -p "Enter tag message (press enter for a lightweight tag): " msg; \
	if [ -z "$$msg" ]; then \
		git tag v$$version; \
	else \
		git tag -a v$$version -m "$$msg"; \
	fi; \
	git push origin v$$version