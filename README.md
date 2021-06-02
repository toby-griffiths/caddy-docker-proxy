# Caddy Proxy

Creates a reverse proxy based on Docker container labels.

## Getting started

Download and execute the setup script… 

```
$ curl --silent https://raw.githubusercontent.com/toby-griffiths/caddy-docker-proxy/main/setup | sh
```

See script output for details of how to label your containers for proxying.

Profit!

Personally I use this function to download the latest script, and cache locally,
then run downloaded, or cached version…

```
caddy-proxy-setup () {
  filename=caddy-proxy-setup
  tmpFile="/tmp/$filename"
  dest="${HOME}/bin/$filename"

  _fetchLatestScript() {
    wget -O "$tmpFile" -o /dev/null https://raw.githubusercontent.com/toby-griffiths/caddy-docker-proxy/main/setup && \
      mv -f "$tmpFile" "$dest"
  }

  echo "Fetching latest version of setup script"
  _fetchLatestScript $tmpFile $dest || echo "Unable to fetch latest version.  Using locally stored version."
  ! test -f "$dest" && echo "Setup script not downloaded.  Aborting" && return 1
  chmod +x "$dest"
  "$dest"
}
```
