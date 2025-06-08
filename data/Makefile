create_competition_season_json:
	jq -c '.[]' competitions/all.json | while read -r line; do \
	  comp_id=$$(echo "$$line" | jq -r '.competition_id'); \
	  season_id=$$(echo "$$line" | jq -r '.season_id'); \
	  echo "$$line" > "competitions/$$comp_id"_$$season_id.json; \
	done

sync_data:
	@for dir in ./*/; do \
		if [ -d $$dir ]; then \
			dir_name=$$(basename $$dir); \
			echo "Syncing $$dir to gs://statsbomb-open-data-api-data/$$dir_name/ ..."; \
			gcloud storage rsync -r $$dir gs://statsbomb-open-data-api-data/$$dir_name/; \
			echo "Finished syncing $$dir"; \
			echo "---------------------------"; \
		fi; \
	done
