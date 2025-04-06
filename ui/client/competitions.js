const competitionSelect = document.getElementById('competitionSelect');
const API_COMPETITIONS_URL = '/api/competitions';

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

// On page load, initialize the Lineups panel with the pitch
document.addEventListener('DOMContentLoaded', function() {
    loadCompetitions();
});