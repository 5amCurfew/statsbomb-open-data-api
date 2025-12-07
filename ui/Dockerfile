# Use official Node 20.10.0 image
FROM node:20.10.0-alpine

# Create app directory
WORKDIR /app

# Copy server package files first (so Docker layer caching works)
COPY package*.json ./

# Install server dependencies
RUN npm install --production

# Copy the rest of the server code
COPY ./ ./

# Expose whatever your server listens on (ex: 8080)
EXPOSE 8080

# Start the server
CMD ["node", "server.js"]
