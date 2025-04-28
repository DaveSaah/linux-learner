#!/bin/bash

# Directory where your YAML files are located
CHALLENGE_DIR="./cmd/linux-learner/data"
OUTPUT_DIR="./bin"

# Make sure output dir exists
mkdir -p $OUTPUT_DIR

# Loop through each YAML file in the data directory
for challenge_file in $CHALLENGE_DIR/*.yaml; do
  # Get the challenge filename without the path or extension
  challenge_name=$(basename "$challenge_file" .yaml)

  echo "Building Go binary for challenge: $challenge_name"

  # Compile the Go binary with the correct challenge file passed via -ldflags
  go build -ldflags "-X 'main.ChallengeFile=$challenge_name.yaml'" -o "$OUTPUT_DIR/linux-learner-$challenge_name" ./cmd/linux-learner/main.go

  if [[ $? -ne 0 ]]; then
    echo "Go build failed for $challenge_name, skipping..."
    continue
  fi

  echo "Building Docker image for challenge: $challenge_name"

  # Build the Docker image for this challenge and pass the correct binary as a build argument
  docker build -t "linux-learner:$challenge_name" --build-arg CHALLENGE_BIN="$OUTPUT_DIR/linux-learner-$challenge_name" .

  if [[ $? -ne 0 ]]; then
    echo "Docker build failed for $challenge_name, skipping..."
    continue
  fi

  # Clean up the binary after building the image
  rm "$OUTPUT_DIR/linux-learner-$challenge_name"

  echo "Docker image for $challenge_name created successfully!"
done

echo "All challenges have been built and Docker images are created!"
