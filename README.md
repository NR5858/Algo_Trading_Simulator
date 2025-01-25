# Algo Trading Simulator

## Overview

The Algo Trading Simulator is a project designed to simulate algorithmic trading strategies. It interacts with RabbitMQ for message handling, allowing for asynchronous communication between different components of the system. The project includes a backend service that processes trading orders and sends them to RabbitMQ.

## Features

- **Algorithmic Trading Simulation**: Simulate various trading strategies and evaluate their performance.
- **RabbitMQ Integration**: Use RabbitMQ for message handling and asynchronous communication.
- **REST API**: Expose a REST API for submitting trading orders.
- **FIX Message Conversion**: Convert trading orders to FIX (Financial Information Exchange) format.

## Project Structure

- **cmd/app**: Contains the main application entry point.
- **pkg/handlers**: Contains HTTP handlers for processing API requests.
- **pkg/models**: Contains data models used in the project.
- **pkg/order**: Contains logic for processing trading orders.
- **pkg/rabbitmq**: Contains functions for interacting with RabbitMQ.

## Getting Started

### Prerequisites

- Go 1.16 or later
- RabbitMQ server

### Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/NR5858/Algo_Trading_Simulator.git
   cd Algo_Trading_Simulator
   ```
