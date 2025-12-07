import * as d3 from "https://cdn.jsdelivr.net/npm/d3@7/+esm";
const matchSelect = document.getElementById('matchSelect');
const API_LINEUPS_URL = '/api/lineups';
const lineUpPanel = document.getElementById('lineUpPanel');

// LINEUPS
matchSelect.addEventListener('change', () => {
    const matchId = matchSelect.value;
    if (matchId) {
        loadLineUps(matchId);
    } else {
        lineUpPanel.innerHTML = 'Select a match to see line ups';
    }
});

async function loadLineUps(matchId) {
    try {
        const response = await fetch(`${API_LINEUPS_URL}?match_id=${matchId}`);
        if (!response.ok) {
            throw new Error(`Server Error: ${response.status}`);
        }
        const lineup = await response.json();
        drawPitch(lineup);
    } catch (error) {
        console.error('Error fetching line ups:', error);
    }
}

// Create the football pitch
function createLineUpPitch(svg, width, height) {
    const pitchWidth = width * 0.9;
    const pitchHeight = height;

    var x = d3.scaleLinear().domain([0, 80]).range([0, pitchWidth]);
    var y = d3.scaleLinear().domain([0, 120]).range([0, pitchHeight]);

    // Outer rectangle
    svg.append('rect')
        .attr('x', (width - pitchWidth) / 2)
        .attr('y', (height - pitchHeight) / 2)
        .attr('width', pitchWidth)
        .attr('height', pitchHeight)
        .attr('stroke', 'black')
        .attr('stroke-width', 1)
        .attr('fill', 'none');

    // Center line
    svg.append('line')
        .attr('x1', (width - pitchWidth) / 2)
        .attr('y1', height / 2)
        .attr('x2', (width + pitchWidth) / 2)
        .attr('y2', height / 2)
        .attr('stroke', 'black')
        .attr('stroke-width', 1);

    // Center circle
    svg.append('circle')
        .attr('cx', width / 2)
        .attr('cy', height / 2)
        .attr('r', 50)
        .attr('stroke', 'black')
        .attr('stroke-width', 1)
        .attr('fill', 'none');

    // Penalty areas and 6-yard boxes
    const topY = (height - pitchHeight) / 2;
    const bottomY = (height - pitchHeight) / 2 + pitchHeight - y(18);

    // Top penalty
    svg.append('rect').attr('x', (width - pitchWidth) / 2 + x(18)).attr('y', topY)
        .attr('width', x(62) - x(18)).attr('height', y(18))
        .attr('stroke', 'black').attr('fill', 'none');
    // Bottom penalty
    svg.append('rect').attr('x', (width - pitchWidth) / 2 + x(18)).attr('y', bottomY)
        .attr('width', x(62) - x(18)).attr('height', y(18))
        .attr('stroke', 'black').attr('fill', 'none');
    // Top 6-yard box
    svg.append('rect').attr('x', (width - pitchWidth) / 2 + x(30)).attr('y', topY)
        .attr('width', x(50) - x(30)).attr('height', y(6))
        .attr('stroke', 'black').attr('fill', 'none');
    // Bottom 6-yard box
    svg.append('rect').attr('x', (width - pitchWidth) / 2 + x(30)).attr('y', bottomY + y(12))
        .attr('width', x(50) - x(30)).attr('height', y(6))
        .attr('stroke', 'black').attr('fill', 'none');
}

// Create lineup players (row-based)
function createLineUpPlayers(svg, data, width, height) {
    const pitchWidth = width * 0.9;
    const pitchHeight = height;
    const pitchOffsetX = (width - pitchWidth) / 2;

    const rowY = {
        GK: 5,
        DEF: 25,
        MID: 50,
        FWD: 75
    };

    const teams = data.map(team => ({
        name: team.team_name,
        players: team.lineup
            .map(p => {
                // Find the starting XI position - must be at from: "00:00"
                const starting = p.positions.find(pos => 
                    pos.start_reason === "Starting XI"
                );
                if (!starting) return null; // not a starter

                // Determine row for plotting
                let row;
                if ([1].includes(starting.position_id)) row = "GK";
                else if ([2,3,4,5,6].includes(starting.position_id)) row = "DEF";
                else if ([12,13,14,15,16].includes(starting.position_id)) row = "MID";
                else if ([17,21,22,23,24,25].includes(starting.position_id)) row = "FWD";
                else row = "MID";

                return {
                    id: p.player_id,
                    name: p.player_name,
                    number: p.jersey_number,
                    posId: starting.position_id,
                    row
                };
            })
            .filter(Boolean) // remove non-starters
    }));

    teams[0].side = "top";
    teams[1].side = "bottom";

    teams.forEach((team, ti) => {
        const teamPlayers = team.players;
        const yScale = d3.scaleLinear()
            .domain([0, 100])
            .range(ti === 0 ? [0, pitchHeight / 2] : [pitchHeight / 2, pitchHeight]);

        const rowGroups = {};
        ["GK","DEF","MID","FWD"].forEach(r => {
            rowGroups[r] = teamPlayers.filter(p => p.row === r);
        });

        Object.entries(rowGroups).forEach(([rowName, rowPlayers]) => {
            const n = rowPlayers.length;
            rowPlayers.forEach((p, i) => {
                const x = pitchOffsetX + pitchWidth * (i + 1) / (n + 1);
                const y = yScale(rowY[rowName]);

                // Player circle
                svg.append("circle")
                    .attr("class", "player")
                    .attr("cx", x)
                    .attr("cy", y)
                    .attr("r", 16)
                    .attr("fill", ti === 0 ? "#1f77b4" : "#ff7f0e")
                    .attr("stroke", "white")
                    .attr("stroke-width", 2);

                // Jersey number
                svg.append("text")
                    .attr("x", x)
                    .attr("y", y + 5)
                    .attr("text-anchor", "middle")
                    .attr("dominant-baseline", "middle")
                    .attr("fill", "white")
                    .attr("font-size", "11px")
                    .attr("font-weight", "bold")
                    .text(p.number);
            });
        });
    });
}

// Draw pitch and players
function drawPitch(lineupData = null) {
    const width = lineUpPanel.clientWidth;
    const height = lineUpPanel.clientHeight;

    // Clear and create SVG
    lineUpPanel.innerHTML = '';
    const svg = d3.select(lineUpPanel)
        .append('svg')
        .attr('width', '100%')
        .attr('height', '100%')
        .attr('viewBox', `0 0 ${width} ${height}`)
        .attr('preserveAspectRatio', 'xMidYMid meet');

    createLineUpPitch(svg, width, height);

    if (lineupData) {
        createLineUpPlayers(svg, lineupData, width, height);
    }
}

// Resize observer
const resizeObserver = new ResizeObserver(() => {
    drawPitch(lineUpPanel._lineupData || null);
});
resizeObserver.observe(lineUpPanel);

// Initial draw
document.addEventListener('DOMContentLoaded', () => drawPitch());