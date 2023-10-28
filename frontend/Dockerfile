#### Use an official Node runtime as a parent image ####
FROM node:latest

# Set the working directory to /app
WORKDIR /app

# Copy package.json file into the container at /app
COPY ./package.json ./

# Copy the current directory contents into the container at /app
COPY ./ ./

# Install any needed packages specified in the package.json
# RUN npm install -g create-vite vite

# Install any needed packages specified in the package.json
RUN npm install

# Make port 3000 available to the world outside this container
EXPOSE 3000

# Define environment variable
ENV BACKEND_URL=http://localhost::8000

# Run the command to start your React app
CMD [ "npm","run","dev" ]
