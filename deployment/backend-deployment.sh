#!/bin/bash

# Set the working directory to your git repository
cd /home/orangepi/GreatMailer/ || { echo "Directory not found!"; exit 1; }

# Pull changes from the Git repository
echo "Pulling changes from Git repository..."

git pull origin master

# Check if the git pull was successful
if [ $? -ne 0 ]; then
	echo "Failed to pull from Git repository!"
	exit 1
fi

# Run Go build
echo "Building the project..."
go build -o GreatBackend

# Check if the build was successful
if [ $? -ne 0 ]; then
	echo "Go build failed!"
	exit 1
fi

chmod +x ./greatBackend

# Run the executable
echo "Running the executable..."
./greatBackend

# Check if the executable ran successfully
if [ $? -ne 0 ]; then
	echo "Executable did not run successfully!"
	exit 1
fi

echo "Deployment completed successfully!"
