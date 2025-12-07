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

// Position ID mappings to pitch coordinates (x: 0-80, y: 0-120 scale, centered at x=40)
const positionMappings = {
    1: { x: 40, y: 10 }, // Goalkeeper
    // Defenders (y=24)
    2: { x: 60, y: 24 },   // Right Back
    3: { x: 52, y: 24 },   // Right Center Back
    4: { x: 40, y: 24 },   // Center Back
    5: { x: 28, y: 24 },   // Left Center Back
    6: { x: 20, y: 24 },   // Left Back
    7: { x: 64, y: 26 },   // Right Wing Back
    8: { x: 16, y: 26 },   // Left Wing Back
    // Defensive Midfielders (y=42)
    9: { x: 52, y: 42 },   // Right Defensive Midfield
    10: { x: 40, y: 42 },  // Center Defensive Midfield
    11: { x: 28, y: 42 },  // Left Defensive Midfield
    // Midfielders (y=60)
    12: { x: 60, y: 60 },  // Right Midfield
    13: { x: 52, y: 60 },  // Right Center Midfield
    14: { x: 40, y: 60 },  // Center Midfield
    15: { x: 28, y: 60 },  // Left Center Midfield
    16: { x: 20, y: 60 },  // Left Midfield
    // Attacking Midfielders (y=78)
    18: { x: 60, y: 78 },  // Right Attacking Midfield
    19: { x: 52, y: 78 },  // Right Midfield/Attacking Midfield
    20: { x: 28, y: 78 },  // Left Attacking Midfield
    // Forwards/Wingers (y=96)
    17: { x: 64, y: 96 },  // Right Wing
    21: { x: 16, y: 96 },  // Left Wing
    22: { x: 48, y: 96 },  // Right Center Forward
    23: { x: 40, y: 106 }, // Center Forward
    24: { x: 32, y: 96 },  // Left Center Forward
    25: { x: 40, y: 84 }   // Secondary Striker
};

// Create lineup players positioned on the pitch
function createLineUpPlayers(svg, data, width, height) {
    const pitchWidth = width * 0.9;
    const pitchHeight = height;
    const pitchOffsetX = (width - pitchWidth) / 2;
    const pitchOffsetY = 0;

    // Create scales matching the pitch (x: 0-80, y: 0-120)
    const xScale = d3.scaleLinear()
        .domain([0, 80])
        .range([pitchOffsetX, pitchOffsetX + pitchWidth]);

    const yScale = d3.scaleLinear()
        .domain([0, 120])
        .range([pitchOffsetY, pitchHeight]);

    // Teams: index 0 is at top (home), index 1 is at bottom (away)
    const teams = data.map((team, idx) => ({
        name: team.team_name,
        teamId: team.team_id,
        isHome: idx === 0,
        players: team.lineup
            .map(p => {
                // Find the starting XI position
                const starting = p.positions.find(pos => 
                    pos.start_reason === "Starting XI"
                );
                if (!starting) return null; // not a starter

                const posId = starting.position_id;
                const mapping = positionMappings[posId] || { x: 40, y: 60 };

                return {
                    id: p.player_id,
                    name: p.player_name,
                    number: p.jersey_number,
                    posId: posId,
                    posName: starting.position,
                    x: mapping.x,
                    y: mapping.y
                };
            })
            .filter(Boolean)
    }));

    teams.forEach((team, ti) => {
        const isHome = ti === 0;
        const color = isHome ? "#1f77b4" : "#ff7f0e";

        // Group players by position to add jitter
        const positionGroups = {};
        team.players.forEach(p => {
            const key = `${p.x},${p.y}`;
            if (!positionGroups[key]) {
                positionGroups[key] = [];
            }
            positionGroups[key].push(p);
        });

        // Apply jitter to players at same position
        team.players.forEach((p, idx) => {
            const key = `${p.x},${p.y}`;
            const playersAtPosition = positionGroups[key];
            
            // Add jitter based on position in group (radius from center)
            const jitterRadius = 3;
            const angle = (idx / playersAtPosition.length) * Math.PI * 2;
            const jitterX = Math.cos(angle) * jitterRadius;
            const jitterY = Math.sin(angle) * jitterRadius;

            const svgX = xScale(p.x + jitterX);
            let svgY;

            if (isHome) {
                // Home team at top half, facing down (0-60 maps to top half)
                svgY = yScale((p.y + jitterY) / 2);
            } else {
                // Away team at bottom half, facing up (flip their y: maps to bottom half)
                svgY = yScale(60 + (60 - (p.y + jitterY)) / 2);
            }

            // Player circle
            svg.append("circle")
                .attr("class", "player")
                .attr("cx", svgX)
                .attr("cy", svgY)
                .attr("r", 16)
                .attr("fill", color)
                .attr("stroke", "white")
                .attr("stroke-width", 2)
                .attr("title", `${p.name} (#${p.number}) - ${p.posName}`);

            // Jersey number
            svg.append("text")
                .attr("x", svgX)
                .attr("y", svgY + 5)
                .attr("text-anchor", "middle")
                .attr("dominant-baseline", "middle")
                .attr("fill", "white")
                .attr("font-size", "11px")
                .attr("font-weight", "bold")
                .text(p.number);
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