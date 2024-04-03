# Use official Node.js image as the base image
FROM node:17-alpine

# Set working directory
WORKDIR /app

# Copy package.json and package-lock.json (if available)
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the entire project directory into the container
COPY . .

# Build React app
RUN npm run build

# Expose port 3000 to the outside world
EXPOSE 5173

# Start the React app
CMD ["npm", "run", "dev"]