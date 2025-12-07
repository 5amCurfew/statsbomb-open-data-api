import {
    getAuthToken
} from './auth.js';
import cors from 'cors';
import dotenv from 'dotenv';
import express from 'express';
import fetch from 'node-fetch';
import path from 'path';

dotenv.config();

const app = express();
const PORT = 3000;
// __dirname equivalent in ES modules
const __dirname = path.resolve();

const API_URL = process.env.API_URL

app.use(express.static(path.join(__dirname, "client")));

app.get("/", (req, res) => {
  res.sendFile(path.join(__dirname, "client", "index.html"));
});

app.use(cors({
    origin: `http://localhost:${PORT}`, // only accept from here (deployed domain when deployed)
    methods: ['GET']
}));

app.get('/api/competitions', async (req, res) => {
    try {
        const token = await getAuthToken();
        if (!token) {
            return res.status(500).json({
                error: 'Auth failed'
            });
        }

        const response = await fetch(`${API_URL}/api/competitions`, {
            headers: {
                'Authorization': `Bearer ${token}`,
                'Content-Type': 'application/json'
            }
        });

        if (!response.ok) {
            throw new Error(`API Error: ${response.status}`);
        }

        const data = await response.json();
        res.json(data);
    } catch (err) {
        console.error(err);
        res.status(500).json({
            error: 'Server error'
        });
    }
});

app.get('/api/matches', async (req, res) => {
    const {
        competition_id,
        season_id
    } = req.query;
    if (!competition_id || !season_id) {
        return res.status(400).json({
            error: 'Missing competition_id or season_id'
        });
    }

    try {
        const token = await getAuthToken();
        const apiRes = await fetch(`${API_URL}/api/matches/${competition_id}/${season_id}`, {
            headers: {
                Authorization: `Bearer ${token}`
            }
        });

        if (!apiRes.ok) {
            return res.status(apiRes.status).json({
                error: 'Failed to fetch matches'
            });
        }

        const matches = await apiRes.json();
        res.json(matches);
    } catch (error) {
        console.error('Error fetching matches:', error);
        res.status(500).json({
            error: 'Server error'
        });
    }
});

app.get('/api/lineups', async (req, res) => {
    const {
        match_id
    } = req.query;
    if (!match_id) {
        return res.status(400).json({
            error: 'Missing match id'
        });
    }

    try {
        const token = await getAuthToken();
        const apiRes = await fetch(`${API_URL}/api/lineups/${match_id}`, {
            headers: {
                Authorization: `Bearer ${token}`
            }
        });

        if (!apiRes.ok) {
            return res.status(apiRes.status).json({
                error: 'Failed to fetch matches'
            });
        }

        const matches = await apiRes.json();
        res.json(matches);
    } catch (error) {
        console.error('Error fetching matches:', error);
        res.status(500).json({
            error: 'Server error'
        });
    }
});

app.listen(PORT, () => {
    console.log(`Server running at http://localhost:${PORT}`);
});