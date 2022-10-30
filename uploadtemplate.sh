#!/bin/bash

function add_template() {
  namespace=$1
  name=$2
  drone template add --namespace "${namespace}" --name "${name}" --data "@${name}" 2&> /dev/null || true
  drone template update --namespace "${namespace}" --name "${name}" --data "@${name}"
}

add_template kjuulh dagger_go_template.yaml
add_template kjuulh gobin_default_template.yaml


