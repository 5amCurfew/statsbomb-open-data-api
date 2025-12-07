sync:
	@make sync_from_statsbomb_open_data && make create_competition_season_json && make sync_data_to_gcs

sync_from_statsbomb_open_data:
	@echo "Syncing data from StatsBomb open-data repository..."
	@if [ ! -d ".statsbomb-temp" ]; then \
		git clone --depth 1 https://github.com/statsbomb/open-data.git .statsbomb-temp; \
	else \
		cd .statsbomb-temp && git pull && cd ..; \
	fi
	@echo "Copying data files..."
	@mkdir -p competitions events lineups matches three-sixty
	@if [ -f ".statsbomb-temp/data/competitions.json" ]; then \
		echo "Copying competitions.json..."; \
		cp .statsbomb-temp/data/competitions.json ./competitions/competitions.json; \
	fi
	@for dir in events lineups matches three-sixty; do \
		if [ -d ".statsbomb-temp/data/$$dir" ]; then \
			echo "Copying $$dir..."; \
			cp -r .statsbomb-temp/data/$$dir/* ./$$dir/; \
		fi; \
	done
	@echo "Cleaning up temporary directory..."
	@rm -rf .statsbomb-temp
	@echo "Sync complete!"

create_competition_season_json:
	jq -c '.[]' competitions/competitions.json | while read -r line; do \
	  comp_id=$$(echo "$$line" | jq -r '.competition_id'); \
	  season_id=$$(echo "$$line" | jq -r '.season_id'); \
	  echo "$$line" > "competitions/$$comp_id"_$$season_id.json; \
	done

sync_data_to_gcs:
	@for dir in ./*/; do \
		if [ -d $$dir ]; then \
			dir_name=$$(basename $$dir); \
			echo "Syncing $$dir to gs://statsbomb-open-data-api-data/$$dir_name/ ..."; \
			gcloud storage rsync -r --delete-unmatched-destination-objects $$dir gs://statsbomb-open-data-api-data/$$dir_name/; \
			echo "Finished syncing $$dir"; \
			echo "---------------------------"; \
		fi; \
	done
