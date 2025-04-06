import * as d3 from "https://cdn.jsdelivr.net/npm/d3@7/+esm";
const matchSelect = document.getElementById('matchSelect');
const API_LINEUPS_URL = '/api/lineups';
const lineUpPanel = document.getElementById('lineUpPanel');

// LINEUPS
matchSelect.addEventListener('change', () => {
    const matchId = matchSelect.value;
    if (matchId) {
        loadLineUps(matchId); // use the full match object
    } else {
        document.getElementById('lineUpPanel').innerHTML = 'Select a match to see line ups';
    }
});

async function loadLineUps(matchId) {
    try {
        const response = await fetch(`${API_LINEUPS_URL}?match_id=${matchId}`);
        if (!response.ok) {
            throw new Error(`Server Error: ${response.status}`);
        }

        const lineup = await response.json();
        //showLineUps(lineup);
    } catch (error) {
        console.error('Error fetching line ups:', error);
    }
};

function createLineUpPitch(svg, width, height) {
    const pitchWidth = width * 0.9; // Narrower width for portrait orientation
    const pitchHeight = height; // Taller height for portrait orientation

    var x = d3.scaleLinear()
        .domain([0, 80])
        .range([0, pitchWidth]);

    var y = d3.scaleLinear()
        .domain([0, 120])
        .range([0, pitchHeight]);

    // Draw the football pitch (rectangular, no fill)
    svg.append('rect')
        .attr('x', (width - pitchWidth) / 2)
        .attr('y', (height - pitchHeight) / 2)
        .attr('width', pitchWidth)
        .attr('height', pitchHeight)
        .attr('class', 'pitch')
        .attr('stroke', 'black') // Set stroke color to black
        .attr('stroke-width', 1)
        .attr('fill', 'none'); // No fill for transparent background

    // Draw the horizontal center line (across the middle of the pitch)
    svg.append('line')
        .attr('x1', (width - pitchWidth) / 2)
        .attr('y1', height / 2) // Center line at the vertical middle of the pitch
        .attr('x2', (width + pitchWidth) / 2)
        .attr('y2', height / 2)
        .attr('stroke', 'black') // Set stroke color to black
        .attr('stroke-width', 1);

    // Draw the center circle
    const centerCircleRadius = 50; // Radius of the center circle

    svg.append('circle')
        .attr('cx', width / 2) // Center of the circle (horizontal center of the pitch)
        .attr('cy', height / 2) // Center of the circle (vertical center of the pitch)
        .attr('r', centerCircleRadius) // Radius of the circle
        .attr('stroke', 'black') // Set stroke color to black
        .attr('stroke-width', 1)
        .attr('fill', 'none'); // No fill for transparent background

    // Draw the penalty area on the top side
    svg.append('rect')
        .attr('x', (width - pitchWidth) / 2 + x(18)) // Starting x position
        .attr('y', (height - pitchHeight) / 2) // Top edge of the penalty area
        .attr('width', x(62) - x(18)) // Width of the penalty area
        .attr('height', y(18)) // Height of the penalty area
        .attr('stroke', 'black')
        .attr('stroke-width', 1)
        .attr('fill', 'none')

    // Draw the penalty area on the top side
    svg.append('rect')
        .attr('x', (width - pitchWidth) / 2 + x(18)) // Starting x position
        .attr('y', (height - pitchHeight) / 2 + y(120) - y(18)) // Top edge of the penalty area
        .attr('width', x(62) - x(18)) // Width of the penalty area
        .attr('height', y(18)) // Height of the penalty area
        .attr('stroke', 'black')
        .attr('stroke-width', 1)
        .attr('fill', 'none')

    // Draw the 6-yard box on the top side
    svg.append('rect')
        .attr('x', (width - pitchWidth) / 2 + x(30)) // Starting x position
        .attr('y', (height - pitchHeight) / 2) // Top edge of the 6-yard box
        .attr('width', x(50) - x(30)) // Width of the 6-yard box
        .attr('height', y(6)) // Height of the 6-yard box
        .attr('stroke', 'black')
        .attr('stroke-width', 1)
        .attr('fill', 'none');

    // Draw the 6-yard box on the bottom side
    svg.append('rect')
        .attr('x', (width - pitchWidth) / 2 + x(30)) // Starting x position
        .attr('y', (height - pitchHeight) / 2 + y(120) - y(6)) // Bottom edge of the 6-yard box
        .attr('width', x(50) - x(30)) // Width of the 6-yard box
        .attr('height', y(6)) // Height of the 6-yard box
        .attr('stroke', 'black')
        .attr('stroke-width', 1)
        .attr('fill', 'none');

};

// On page load, initialize the Lineups panel with the pitch
document.addEventListener('DOMContentLoaded', () => {
    // Function to update and redraw the pitch based on the current size of lineUpPanel
    function drawPitch() {
        const width = lineUpPanel.clientWidth; // Get the updated width
        const height = lineUpPanel.clientHeight; // Get the updated height

        // Clear any previous content in lineUpPanel
        lineUpPanel.innerHTML = '';

        // Create an SVG canvas for the pitch
        const svg = d3.select(lineUpPanel)
            .append('svg')
            .attr('width', width)
            .attr('height', height);

        // Create the football pitch
        createLineUpPitch(svg, width, height);
    }

    // Initial draw
    drawPitch();

    // Create a ResizeObserver to monitor changes in the size of lineUpPanel
    const resizeObserver = new ResizeObserver(function() {
        drawPitch(); // Redraw the pitch when the size changes
    });

    // Start observing the lineUpPanel for size changes
    resizeObserver.observe(lineUpPanel);
});