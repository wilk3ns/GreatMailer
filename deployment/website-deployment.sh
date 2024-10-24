#!/bin/bash

# Set the working directory to your git repository
cd /home/orangepi/web-page/ || { echo "Directory not found!"; exit 1; }

# Pull changes from the Git repository
echo "Pulling changes from Git repository..."
git pull origin master

# Check if the git pull was successful
if [ $? -ne 0 ]; then
    echo "Failed to pull from Git repository!"
    exit 1
fi

# Check if node_modules exists
if [ ! -d "node_modules" ]; then
    echo "node_modules not found. Installing npm dependencies..."
    npm install

    # Check if the npm install was successful
    if [ $? -ne 0 ]; then
        echo "npm install failed!"
        exit 1
    fi
else
    echo "node_modules already present. Skipping npm install."
fi

# Check if port 3000 is in use
if lsof -i :3000 > /dev/null; then
    echo "Stopping any process using port 3000..."
    fuser -k 3000/tcp
    sleep 2 # Allow a moment for the process to terminate
else
    echo "No process found using port 3000."
fi

# Start the React app
echo "Starting the React application..."
npm start &

# Check if the application started successfully
if [ $? -ne 0 ]; then
    echo "React application did not run successfully!"
    exit 1
fi

echo "Deployment completed successfully!"