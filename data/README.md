## Getting Started

The [data](./data/) is provided as JSON files exported from the StatsBomb Data API, in the following structure:

* Competition and seasons stored in [`competitions.json`](./data/competitions.json).
* Matches for each competition and season, stored in [`matches`](./data/matches/). Each folder within is named for a competition ID, each file is named for a season ID within that competition.
* Events and lineups for each match, stored in [`events`](./data/events/) and [`lineups`](./data/lineups/) respectively. Each file is named for a match ID.
* StatsBomb 360 data for selected matches, stored in [`three-sixty`](./data/three-sixty/). Each file is named for a match ID.

Some documentation about the meaning of different events and the format of the JSON can be found in the [`doc`](./doc) directory.

## Terms & Conditions

If you publish, share or distribute any research, analysis or insights based on this data, please state the data source as StatsBomb and use our logo, available in our [Media Pack](https://statsbomb.com/media-pack/).