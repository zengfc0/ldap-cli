From harbor.hyperchain.cn/hub/jenkins-slave:centos7-go1.18.9 as builder

WORKDIR /home

COPY . .

RUN go env -w GOPROXY=https://goproxy.cn,direct \
    && go mod tidy \
    && go build

From harbor.hyperchain.cn/hub/jenkins-slave:centos7-java1.8

COPY --from=builder /home/ldap-cli /usr/local/bin/

RUN set -eux \
    && yum -y install jq
