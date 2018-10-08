#!/bin/bash

echo -n '{"server": "http://prometheus.openfaas:9090", "query": "sum(increase(gateway_function_invocation_total[1h]))  by (function_name)", "start": "5 hours ago", "end": "now", "step": "1h","format": "table"}' | faas-cli invoke promq --gateway=$OPENFAAS_URL