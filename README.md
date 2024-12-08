# Command & Control Server (C2)

## Overview
This is a Command & Control (C2) server project developed as part of the **Luddy Cybersecurity Club** at Indiana University during the second semester of the 2023-2024 academic year. The project is a collaborative effort that provides hands-on experience in designing and implementing a functional C2 server using **Go**.

This server includes features like agent management, command execution, and a custom CLI, making it an excellent platform to learn about cybersecurity operations and programming.

## Features
- **Custom CLI**:
  - Commands to manage agents, execute tasks, and monitor activity.
  - User-friendly interface with auto-completion and history logging.
- **Agent Management**:
  - Handle multiple agents and interact with them in real-time.
- **Payload Generation**:
  - A flexible payload structure to deploy and manage agents effectively.
- **Log Handling**:
  - Logging activities for tracking and debugging purposes.

## Usage
### Build the Server
To build the C2 server and payload, use the included `Makefile`:
```bash
make build
```

### Run the C2 Server
Start the C2 server:
```bash
go run main.go
```

### Interact with the CLI
The custom CLI provides commands to manage agents and execute tasks. Example commands:
```
C2 > help
C2 > agents
C2 > agent <agent_name>
C2 > exec <command>
```

## Contribution
This project was developed as part of a club initiative to learn and explore cybersecurity topics. Contributions and improvements are welcome.

## Acknowledgments
- **Indiana University Luddy Cybersecurity Club** for hosting the development of the project.
- Original repository [C2Dev24](https://github.com/coremedic/C2Dev24)
