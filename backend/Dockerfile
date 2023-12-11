#### Use the official Golang image as a parent image ####
FROM golang:latest

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY ./ ./

# Manage dependencies
RUN go mod tidy \
    && go mod verify

# Make port 8000 available to the world outside this container
EXPOSE 8000

RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air

ENTRYPOINT ["air"]


