
# Go-Template Plugin

This plugin is designed to work with the `go-template` binary to process templates based on provided values. The plugin is built using Go and is packaged as a Docker container for easy integration.

## How to Build and Run the Docker Image

1. **Build the Docker Image**: Run the following command in your terminal:

    ```bash
    docker build -t go-template-plugin .
    ```

2. **Run the Docker Container**: To run the container with all the necessary parameters, you can use:

    ```bash
    docker run -e PLUGIN_TEMPLATE=path/to/template -e PLUGIN_VALUES=path/to/values -e PLUGIN_OUTPUT=path/to/output go-template-plugin
    ```

## How to Use This Plugin in a Pipeline

### YAML Configuration Example

Here's an example of how you could configure this plugin within a pipeline:

```yaml
type: Plugin
spec:
    connectorRef: <Docker_Hub_Account_or_Other_Registry>
    image: go-template-plugin:latest
    settings:
        template: path/to/template.yaml
        values: path/to/values.yaml
        output: path/to/output/
```

Replace `<Docker_Hub_Account_or_Other_Registry>` with the actual registry where the Docker image for this plugin is hosted.

