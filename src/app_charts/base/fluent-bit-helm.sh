#!/bin/bash
# needs at least helm v3.5.0
OUTPUT=./robot/fluent-bit.yaml
TEMPLATE_VERSION=0.20.5
helm repo add fluent https://fluent.github.io/helm-charts
helm repo update fluent
helm template fluent-bit fluent/fluent-bit --version ${TEMPLATE_VERSION} -f fluent-bit-values.yaml --skip-tests > ${OUTPUT}

sed -i '1i\{{ if and (eq .Values.robot_authentication "true") (eq .Values.fluentbit "true") }}' ${OUTPUT}
sed -i '1i\# !!! DO NOT EDIT THIS FILE !!!\n# This file is autogenerated using src/app_charts/base/fluent-bit-helm.sh.\n# See src/app_charts/base/README.md for update instructions.' ${OUTPUT}
sed -i '$a\{{ end }}' ${OUTPUT}
sed -i 's/MY_ROBOT/{{ .Values.robot.name }}/' ${OUTPUT}
# This needs to be an actual Cloud zone so that it can be mapped
# to a Monarch/Stackdriver region. TODO(swolter): We should make
# this zone configurable to avoid confusing users.
sed -i 's/MY_CLUSTER_LOCATION/europe-west1-c/' ${OUTPUT}
