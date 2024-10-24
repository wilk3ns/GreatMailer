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

# Stop the React app if it's already running
if pgrep -f "npm start" > /dev/null; then
    echo "Stopping the currently running React application..."
    pkill -f "npm start"
    sleep 2 # Allow a moment for the process to terminate
else
    echo "No running React application found."
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