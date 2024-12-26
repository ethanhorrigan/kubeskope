KubeSkope

KubeSkope is a graphical CLI tool designed to monitor Kubernetes resource usage. It provides insights into CPU and memory requests, limits, and actual usage for containers running in a specified namespace. This tool helps Kubernetes administrators identify over-dimensioned or under-dimensioned resources and optimize their workloads.

Features

Namespace Monitoring: Analyze all pods and containers in a specified namespace.

Resource Analysis: Display CPU and memory requests, limits, and live usage.

Graphical CLI Output: Outputs data in a table format with clear visualization.

Customizable Kubeconfig: Allows specifying a custom kubeconfig path.

Prerequisites

Kubernetes Cluster: A running Kubernetes cluster (e.g., Kind, Minikube, or a production cluster).

Metrics Server: Ensure the Kubernetes Metrics Server is installed and running.

Go: Installed Go environment (1.20 or later).

Installation

Clone the Repository:

git clone https://github.com/<your-username>/kubeskope.git
cd kubeskope

Build the Application:

go build -o kubeskope

Run the Application:

./kubeskope monitor --namespace <namespace> --kubeconfig <path-to-kubeconfig>

Usage

Monitor a Namespace

Run the following command to monitor resource usage in a specific namespace:

./kubeskope monitor --namespace test-namespace --kubeconfig ~/.kube/config

Example Output

+---------------------------------+-----------+--------------+-----------------+------------+---------------+----------------------------+
|               POD               | CONTAINER | CPU REQUESTS | MEMORY REQUESTS | CPU LIMITS | MEMORY LIMITS |           USAGE            |
+---------------------------------+-----------+--------------+-----------------+------------+---------------+----------------------------+
| nginx-deployment-8b7fcb74-7qsnb | nginx     | 100m         | 128Mi           | 200m       | 256Mi         | CPU: 50m, Memory: 64Mi     |
+---------------------------------+-----------+--------------+-----------------+------------+---------------+----------------------------+
| nginx-deployment-8b7fcb74-vm8cm | nginx     | 100m         | 128Mi           | 200m       | 256Mi         | CPU: 60m, Memory: 70Mi     |
+---------------------------------+-----------+--------------+-----------------+------------+---------------+----------------------------+
| nginx-deployment-8b7fcb74-vsvdp | nginx     | 100m         | 128Mi           | 200m       | 256Mi         | CPU: 40m, Memory: 50Mi     |
+---------------------------------+-----------+--------------+-----------------+------------+---------------+----------------------------+

Flags and Options

--namespace

Description: Specify the namespace to monitor.

Example:

./kubeskope monitor --namespace test-namespace

--kubeconfig

Description: Path to the kubeconfig file.

Example:

./kubeskope monitor --kubeconfig ~/.kube/config

Development

Run the Application Locally

Ensure you have Go installed.

Run the application:

go run main.go monitor --namespace <namespace> --kubeconfig <path-to-kubeconfig>

Run Tests

Use the built-in Go testing framework to run unit tests:

go test ./...

Contributing

Contributions are welcome! To contribute:

Fork the repository.

Create a feature branch.

Commit your changes.

Open a pull request.

License

This project is licensed under the MIT License. See the LICENSE file for details.

Troubleshooting

Common Issues

Metrics Not Available: Ensure the Metrics Server is installed and running in your cluster.

kubectl get apiservices | grep metrics

Connection Errors: Verify the kubeconfig path and cluster access.

kubectl get nodes --kubeconfig <path-to-kubeconfig>

Logs

Enable verbose logging by adding debug statements in the code.

Contact

For support or inquiries, please create an issue in the repository or contact the maintainers.
