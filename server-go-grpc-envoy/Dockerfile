FROM envoyproxy/envoy:v1.22.0

COPY envoy-template.yaml /etc/envoy/envoy-template.yaml
RUN apt -qq update
RUN apt install -y gettext
RUN chmod -R 777 /etc/envoy

CMD envsubst < /etc/envoy/envoy-template.yaml > /etc/envoy/envoy.yaml && /usr/local/bin/envoy -c /etc/envoy/envoy.yaml -l trace --log-path /tmp/envoy_info.log
