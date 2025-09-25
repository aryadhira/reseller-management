# Stage 1: The Builder Stage
# We use a full Node.js environment to install dependencies and build the application.
FROM node:24-alpine AS builder

# Set the working directory inside the container.
WORKDIR /app

# Copy package.json and package-lock.json first to take advantage of Docker layer caching.
# This ensures that npm dependencies are not reinstalled on every code change.
COPY package*.json ./

# Install project dependencies.
RUN npm install

# Copy the rest of the application source code.
COPY . .

# Build the Nuxt application for production.
# This command generates the .output directory which contains the entire server and client bundles.
RUN npm run build

# Stage 2: The Production Stage
# This is a minimal stage that contains only the necessary files to run the application.
FROM node:24-alpine AS production

# Set the working directory.
WORKDIR /app

# Copy the .output directory from the builder stage.
# This is the most crucial step as it copies only the compiled application.
COPY --from=builder /app/.output ./

# Expose the default Nuxt port.
EXPOSE 3000

# Set the command to run the production server as specified in the Nuxt documentation.
CMD ["node", "./server/index.mjs"]