`docker logs [OPTIONS] CONTAINER`    

Options

| Option                                                       | Default | Description                                                  |
| ------------------------------------------------------------ | ------- | ------------------------------------------------------------ |
| `--details`                                                  |         | Show extra details provided to logs                          |
| `-f, --follow`                                               |         | Follow log output                                            |
| `--since`                                                    |         | Show logs since timestamp (e.g. `2013-01-02T13:23:37Z`) or relative (e.g. `42m` for 42 minutes) |
| `-n, --tail`                                                 | `all`   | Number of lines to show from the end of the logs             |
| `-t, --timestamps`                                           |         | Show timestamps                                              |
| `--until`                                                    |         | Show logs before a timestamp (e.g. `2013-01-02T13:23:37Z`) or relative (e.g. `42m` for 42 minutes) |
