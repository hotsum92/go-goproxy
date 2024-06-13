#!/bin/bash

# proxy: http://<username>:<password>@<proxy_host>:<proxy_port>
# export http_proxy=http://<username>:<password>@<proxy_host>:<proxy_port>
# export https_proxy=http://<username>:<password>@<proxy_host>:<proxy_port>
curl -s http://www.google.co.jp -x localhost:8080
