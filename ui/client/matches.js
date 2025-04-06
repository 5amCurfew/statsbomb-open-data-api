const competitionSelect = document.getElementById('competitionSelect');
const matchSelect = document.getElementById('matchSelect');
const API_MATCHES_URL = '/api/matches';
const matchMap = new Map();

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
    matchMap.clear(); // clear old data

    matchSelect.innerHTML = '<option value="">Select</option>';
    matches.forEach(match => {
        matchMap.set(match.match_id, match); // store full match object

        const option = document.createElement('option');
        option.value = parseInt(match.match_id);
        option.textContent = `${match.home_team.home_team_name} vs ${match.away_team.away_team_name} (${match.match_date})`;
        matchSelect.appendChild(option);
    });
}

// MATCH DETAIL
matchSelect.addEventListener('change', () => {
    const matchId = matchSelect.value;
    const match = matchMap.get(parseInt(matchId));

    if (match) {
        showMatchDetails(match); // use the full match object
    } else {
        document.getElementById('matchDetails').innerHTML = 'Select a match to see details';
    }
});

function showMatchDetails(match) {
    const detailsDiv = document.getElementById('matchDetails');
    detailsDiv.innerHTML = `
      <p><strong>KO:</strong> ${match.match_date} ${match.kick_off}</p>
      <p><strong>Stadium:</strong> ${match.stadium.name}, ${match.stadium.country.name}</p>
      <p><strong>Score:</strong> ${match.home_team.home_team_name} ${match.home_score} - ${match.away_score} ${match.away_team.away_team_name}</p>
      <p><strong>Referee:</strong> ${match.referee.name} (${match.referee.country.name})</p>
      <p><strong>Home Manager:</strong> ${match.home_team.managers[0].name} (${match.home_team.managers[0].country.name})</p>
      <p><strong>Away Manager:</strong> ${match.away_team.managers[0].name} (${match.away_team.managers[0].country.name})</p>
    `;
};