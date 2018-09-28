

—-

## Installation

<details><summary>show</summary>
<p>

```
go get -u github.com/pressly/sup/cmd/sup

```

</p>
</details>

—-

## Usage

sup [OPTIONS] NETWORK COMMAND [...]

<details><summary>show</summary>
<p>

| Option            | Description                      |
|-------------------|----------------------------------|
| `-f Supfile`      | Custom path to Supfile           |
| `-e`, `--env=[]`  | Set environment variables        |
| `--only REGEXP`   | Filter hosts matching regexp     |
| `--except REGEXP` | Filter out hosts matching regexp |
| `--debug`, `-D`   | Enable debug/verbose mode        |
| `--disable-prefix`| Disable hostname prefix          |
| `--help`, `-h`    | Show help/usage                  |
| `--version`, `-v` | Print version                    |

</p>
</details>

—-

## Supfile networks

<details><summary>show</summary>
<p>

```
# Supfile

networks:
    production:
        hosts:
            - api1.example.com
            - api2.example.com
            - api3.example.com
    staging:
        # fetch dynamic list of hosts
        inventory: curl http://example.com/latest/meta-data/hostname

```

sup production COMMAND` will run COMMAND on `api1`, `api2` and `api3` hosts in parallel.

</p>
</details>

—-

## Supfile commands

A shell command(s) to be run remotely.

<details><summary>show</summary>
<p>

```
# Supfile

commands:
    restart:
        desc: Restart example Docker container
        run: sudo docker restart example
    tail-logs:
        desc: Watch tail of Docker logs from all hosts
        run: sudo docker logs --tail=20 -f example

```

`$ sup staging restart` will restart all staging Docker containers in parallel.

`$ sup production tail-logs` will tail Docker logs from all production containers in parallel.

</p>
</details>


—-

## Supfile Serial command (a.k.a. Rolling Update)

<details><summary>show</summary>
<p>

`serial: N` constraints a command to be run on `N` hosts at a time at maximum. Rolling Update for free!

```
# Supfile

commands:
    restart:
        desc: Restart example Docker container
        run: sudo docker restart example
        serial: 2

```

`$ sup production restart` will restart all Docker containers, two at a time at maximum.

</p>
</details>


—-

## Once command (one host only)

<details><summary>show</summary>
<p>

`once: true` constraints a command to be run only on one host. Useful for one-time tasks.

```yaml
# Supfile

commands:
    build:
        desc: Build Docker image and push to registry
        run: sudo docker build -t image:latest . && sudo docker push image:latest
        once: true # one host only
    pull:
        desc: Pull latest Docker image from registry
        run: sudo docker pull image:latest
```

`$ sup production build pull` will build Docker image on one production host only and spread it to all hosts.

### Local command

Runs command always on localhost.

```yaml
# Supfile

commands:
    prepare:
        desc: Prepare to upload
        local: npm run build
```


</p>
</details>


—-

## Upload command

Uploads files/directories to all remote hosts. Uses `tar` under the hood.

<details><summary>show</summary>
<p>

```yaml
# Supfile

commands:
    upload:
        desc: Upload dist files to all hosts
        upload:
          - src: ./dist
            dst: /tmp/
```

</p>
</details>

—-

## Interactive Bash on all hosts

Do you want to interact with multiple hosts at once? Sure!

<details><summary>show</summary>
<p>

```yaml
# Supfile

commands:
    bash:
        desc: Interactive Bash on all hosts
        stdin: true
        run: bash
```

```bash
$ sup production bash
#
# type in commands and see output from all hosts!
# ^C
```

Passing prepared commands to all hosts:
```bash
$ echo 'sudo apt-get update -y' | sup production bash

# or:
$ sup production bash <<< 'sudo apt-get update -y'

# or:
$ cat <<EOF | sup production bash
sudo apt-get update -y
date
uname -a
EOF
```

</p>
</details>


—-

## Interactive Docker Exec on all hosts

<details><summary>show</summary>
<p>

```yaml
# Supfile

commands:
    exec:
        desc: Exec into Docker container on all hosts
        stdin: true
        run: sudo docker exec -i $CONTAINER bash
```

```bash
$ sup production exec
ps aux
strace -p 1 # trace system calls and signals on all your production hosts
```

</p>
</details>

—-

## Target

<details><summary>show</summary>
<p>

Target is an alias for multiple commands. Each command will be run on all hosts in parallel,
`sup` will check return status from all hosts, and run subsequent commands on success only
(thus any error on any host will interrupt the process).

```yaml
# Supfile

targets:
    deploy:
        - build
        - pull
        - migrate-db-up
        - stop-rm-run
        - health
        - slack-notify
        - airbrake-notify
```

</p>
</details>


—-

## Basic structure

<details><summary>show</summary>
<p>

```yaml

version: 0.4

# Global environment variables
env:
  NAME: api
  IMAGE: example/api

networks:
  local:
    hosts:
      - localhost
  staging:
    hosts:
      - stg1.example.com
  production:
    hosts:
      - api1.example.com
      - api2.example.com

commands:
  echo:
    desc: Print some env vars
    run: echo $NAME $IMAGE $SUP_NETWORK
  date:
    desc: Print OS name and current date/time
    run: uname -a; date

targets:
  all:
    - echo
    - date

```



</p>
</details>


—-

## Default environment variables

<details><summary>show</summary>
<p>

- `$SUP_HOST` - Current host.
- `$SUP_NETWORK` - Current network.
- `$SUP_USER` - User who invoked sup command.
- `$SUP_TIME` - Date/time of sup command invocation.
- `$SUP_ENV` - Environment variables provided on sup command invocation. You can pass `$SUP_ENV` to another `sup` or `docker` commands in your Supfile.

</p>
</details>


—-

