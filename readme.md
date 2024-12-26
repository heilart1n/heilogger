# Logger Configuration

This project provides a customizable logger with support for file rotation, multiple output formats, and flexible configuration options. Below are details about the configuration options, example configurations in both JSON and YAML, and usage instructions.

---

## Configuration Options

### Top-Level Config Fields

| Field                | Type      | Description                                                   | Default Value |
|----------------------|-----------|---------------------------------------------------------------|---------------|
| `output_format`      | `string`  | Format of log output. Options: `text`, `json`, `pretty-json`. | `text`        |
| `enable_file_output` | `boolean` | If `true`, logs will be saved to a file.                      | `false`       |
| `output_level`       | `string`  | Minimum log level to output (`info`, `debug`, etc.).          | `info`        |
| `with_source`        | `boolean` | If `true`, includes source file and line number in logs.      | `true`        |
| `output_to_console`  | `boolean` | If `true`, logs will also print to the console.               | `true`        |
| `rotation`           | `object`  | Rotation settings for log files.                              | See below.    |

---

### Rotation Config Fields

| Field              | Type       | Description                                                      | Default Value   |
|--------------------|------------|------------------------------------------------------------------|-----------------|
| `rotate_daily`     | `boolean`  | If `true`, log files will rotate daily.                          | `false`         |
| `max_age`          | `duration` | Maximum time to retain old log files (e.g., `720h` for 30 days). | `168h` (7 days) |
| `rotation_time`    | `duration` | Time interval for rotation (e.g., `24h`).                        | `24h`           |
| `output_directory` | `string`   | Directory where log files are stored.                            | `./logs`        |
| `file_name`        | `string`   | Name of the log file (e.g., `log_YYYY-MM-DD.log`).               | `log.txt`       |

---

## Example Configurations

### Combined JSON and YAML Example

Here are examples of the configuration in both JSON and YAML formats:

#### JSON Configuration

```json
{
  "output_format": "text",
  "enable_file_output": true,
  "output_level": "info",
  "with_source": true,
  "output_to_console": true,
  "rotation": {
    "rotate_daily": true,
    "max_age": "720h",
    "rotation_time": "24h",
    "output_directory": "./logs",
    "file_name": "log_YYYY-MM-DD.log"
  }
}
```
### YAML Confirmation

```yaml
output_format: text
enable_file_output: true
output_level: info
with_source: true
output_to_console: true
rotation:
  rotate_daily: true
  max_age: 720h
  rotation_time: 24h
  output_directory: ./logs
  file_name: log_YYYY-MM-DD.log
```
