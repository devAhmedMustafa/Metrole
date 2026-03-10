# Software Requirements

## CLI
- Execute commands via terminal
- Pass arguments for configuration
- Display menus and options
- Handle user input and validation
- Provide feedback and error messages
- Support Windows, macOS, and Linux

## Request Handling
- Establish a websocket server to listen for incoming requests
- Parse and validate incoming requests
- Forward requests to the appropriate container
- Handle responses from containers and send them back to the requester

## Container Management
- Assign unique domain name for each container mapped to its port

## Docker Integration
- Create images from docker compose file
- Control resources allocated to containers (CPU and Memory limits)
- Start and stop containers as needed

## Tunneling Server Integration
- Communicate with the tunneling server to manage domain name mappings
- Handle IP changes and update the tunneling server accordingly

## Metrole AI
- Detect project type from file structure
- Generate Docker compose from detection
- Validate generated file before building

## Authentication
- Implement google authentication for secure access to the CLI and tunneling server