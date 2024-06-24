# Use the official node image as a base
FROM node:18-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy package.json and package-lock.json
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the app for production
RUN npm run build

# Expose port 3000 to the outside world
EXPOSE 3000

# Serve the built app using a simple HTTP server
CMD ["npm", "run", "preview"]
