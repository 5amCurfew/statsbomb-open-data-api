import fetch from 'node-fetch';

const API_LOGIN_URL = 'http://localhost:4000/auth/login';

// Simple in-memory cache
let cachedToken = null;

function isTokenExpired(token) {
    try {
        const payload = JSON.parse(Buffer.from(token.split('.')[1], 'base64').toString('utf-8'));
        const expiry = payload.exp * 1000;
        return Date.now() >= expiry;
    } catch (err) {
        console.error('Error decoding token:', err);
        return true;
    }
}

export async function getAuthToken() {
    if (cachedToken && !isTokenExpired(cachedToken)) {
        return cachedToken;
    }

    console.warn('Token expired or not found. Fetching new token...');
    cachedToken = await fetchNewToken();
    return cachedToken;
}

async function fetchNewToken() {
    try {
        const response = await fetch(API_LOGIN_URL, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                email: process.env.API_USER_EMAIL,
                password: process.env.API_USER_PASS
            })
        });

        if (!response.ok) {
            throw new Error(`Login failed: ${response.status}`);
        }

        const data = await response.json();
        cachedToken = data.token;
        console.log('Token updated successfully');
        return cachedToken;
    } catch (err) {
        console.error('Error fetching token:', err);
    }
}

export function clearAuthToken() {
    cachedToken = null;
}