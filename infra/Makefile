sync_data:
	@for dir in ../data/*/; do \
		if [ -d $$dir ]; then \
			dir_name=$$(basename $$dir); \
			echo "Syncing $$dir to gs://statsbomb-open-data-api-data/$$dir_name/ ..."; \
			gcloud storage rsync -r $$dir gs://statsbomb-open-data-api-data/$$dir_name/; \
			echo "Finished syncing $$dir"; \
			echo "---------------------------"; \
		fi; \
	done
