const competitionSelect = document.getElementById('competitionSelect');
const matchSelect = document.getElementById('matchSelect');
const API_COMPETITIONS_URL = '/api/competitions';
const API_MATCHES_URL = '/api/matches';

// COMPETITIONS
async function loadCompetitions() {
  try {
    const response = await fetch(API_COMPETITIONS_URL);
    if (!response.ok) {
      throw new Error(`Server Error: ${response.status}`);
    }

    const competitions = await response.json();
    populateCompetitionDropdown(competitions);
  } catch (error) {
    console.error('Error fetching competitions:', error);
    competitionSelect.innerHTML = '<option value="">Failed to load competitions</option>';
  }
}

function populateCompetitionDropdown(competitions) {
  competitionSelect.innerHTML = '<option value="">Select</option>';
  
  competitions.forEach(competition => {
    const option = document.createElement('option');
    option.value = `${competition.competition_id}-${competition.season_id}`;
    option.textContent = `${competition.competition_name} (${competition.season_name})`;
    competitionSelect.appendChild(option);
  });
}

loadCompetitions();

// MATCHES
competitionSelect.addEventListener('change', () => {
  const value = competitionSelect.value;
  if (!value) {
    matchSelect.innerHTML = '<option value="">Select</option>';
    return;
  }

  const [competitionId, seasonId] = value.split('-');
  loadMatches(competitionId, seasonId);
});

async function loadMatches(competitionId, seasonId) {
  try {
    const response = await fetch(`${API_MATCHES_URL}?competition_id=${competitionId}&season_id=${seasonId}`);
    if (!response.ok) throw new Error(`Server Error: ${response.status}`);
    const matches = await response.json();
    populateMatchDropdown(matches);
  } catch (error) {
    console.error('Error fetching matches:', error);
    matchSelect.innerHTML = '<option value="">Failed to load matches</option>';
  }
}

function populateMatchDropdown(matches) {
  matchSelect.innerHTML = '<option value="">Select</option>';
  matches.forEach(match => {
    const option = document.createElement('option');
    option.value = match.match_id;
    option.textContent = `${match.home_team.home_team_name} vs ${match.away_team.away_team_name} (${match.match_date})`;
    matchSelect.appendChild(option);
  });
}