# Go MQTT Server and Client

This project demonstrates how to create a simple MQTT server and client using Go. It includes the following key components:

- **MQTT Server:** Implements a basic MQTT server using the Mochi-MQTT library.
- **MQTT Client:** Connects to the server and publishes/subscribes to messages using the Paho Go MQTT Client library.
- **Docker Support:** Uses Docker Compose to easily build and run both the server and client in containers.

## Running in Docker

1. **Prerequisites:**
   - Docker and Docker Compose installed on your system.

2. **Build and Run:**
   - Navigate to the project directory in your terminal.
   - Run the following command to build and start the containers:

       ```bash
       docker-compose up
       ```

3. **Accessing Server Logs:**
   - To view the server logs, open a new terminal window and run:

       ```bash
       docker logs mqtt-server
       ```

## Manual Execution (Optional)

- **Server:**
   - Build and run:

     ```bash
     go build server/pkg/server.go
     ./server
     ```

- **Client:**
   - Build and run:

     ```bash
     go build server/pkg/client.go
     ./client
     ```

## Additional Notes

- Consider modifying `docker-compose.yml` to expose a different port if 1883 is already in use.
- Explore security configurations for a production-ready MQTT setup.
- Refer to the documentation for Mochi-MQTT and Paho Go MQTT Client for further details and advanced usage.

---

**Key Libraries:**

- Mochi-MQTT: [https://github.com/mochi-mqtt](https://github.com/mochi-mqtt)
- Paho Go MQTT Client: [https://github.com/eclipse/paho.golang](https://github.com/eclipse/paho.golang)

