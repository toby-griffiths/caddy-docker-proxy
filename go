#!/usr/bin/env sh

set -e

echo

NETWORK_EXISTS=1
docker network inspect caddy >/dev/null || NETWORK_EXISTS=0

if [ 0 = "${NETWORK_EXISTS}" ]; then
  echo "'caddy' network does not exists.  Creating…"
  docker network create caddy
  echo "   Done!"
else
  echo "'caddy' network already exists."
fi
echo

cat <<EOF | docker-compose -f /dev/stdin -p caddy-proxy up -d
version: "3.7"
services:
  caddy:
    image: lucaslorentz/caddy-docker-proxy:alpine
    ports:
      - 80:80
      - 443:443
      - 2019:2019
    environment:
      - CADDY_INGRESS_NETWORKS=caddy
    networks:
      - caddy
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
      - caddy_data:/data
    restart: unless-stopped
    labels:
      service: caddy-proxy
      function: proxy

  whoami:
    image: jwilder/whoami
    networks:
      - caddy
    labels:
      service: caddy-proxy
      function: test-container
      caddy: whoami.localhost
      caddy.reverse_proxy: '{{upstreams 8000}}'

networks:
  caddy:
    external: true

volumes:
  caddy_data: {}
EOF

echo
echo "Caddy proxy setup"
echo
echo "To proxy an image add the container to the \`caddy\` network, and add the minimal"
echo "caddy labels…"
echo
cat <<TEXT
    services
      nginx:
        networks:
          - caddy
        labels:
          caddy: whoami.localhost
          # caddy.reverse_proxy: '{{upstreams {port}}}'
          # Omitting {port} will expose port 80 by default
          caddy.reverse_proxy: '{{upstreams 8000}}'
TEXT
echo
echo
echo "For more info on defining config labels on the target images see…"
echo "https://github.com/lucaslorentz/caddy-docker-proxy"
echo
echo
echo "To stop the proxy, run…"
echo
echo "    $ docker rm --force \$(docker ps --filter \"label=service=caddy-proxy\" -q)"
