# Systems Monitor

Systems Monitor is a tool developed in Go for monitoring the status of multiple web pages. It allows you to periodically check the status of configured URLs and receive notifications in case of errors.

## Features

- Periodic monitoring of multiple URLs.
- Flexible configuration of monitoring interval and expected status codes.
- Logging of events and errors with severity levels.
- Easy configuration via YAML files.

## Installation

To install Systems Monitor, you need to have Go installed on your system. Then, you can clone this repository and run the following command:

```bash
go install ./cmd/systems-monitor
```

## Usage

1. Create a urls.yml file with the configuration of the URLs you want to monitor.
2. Run the systems-monitor command in your terminal to start monitoring.

### Example of urls.yml

```yaml
urls:
  - url: http://example.com
    timeout: 5
    interval: 15
    expected_status_codes:
      - 200
      - 201
  - url: http://example.org
    timeout: 10
    interval: 30
    expected_status_codes:
      - 200-209
```

## Contributing

If you want to contribute to Systems Monitor, you are welcome to do so! You can open an issue to discuss new features or problems, or submit a pull request with your contributions.

## License

This project is licensed under the [GNU General Public License v3.0](LICENSE).

