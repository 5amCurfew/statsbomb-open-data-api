create_competition_season_json:
	jq -c '.[]' competitions/all.json | while read -r line; do \
	  comp_id=$$(echo "$$line" | jq -r '.competition_id'); \
	  season_id=$$(echo "$$line" | jq -r '.season_id'); \
	  echo "$$line" > "competitions/$$comp_id"_$$season_id.json; \
	done