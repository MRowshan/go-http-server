.PHONY = build install uninstall

NAME = go-http-server
INSTALL_LOCATION = /opt/${NAME}
SERVICE_SCRIPT = /etc/systemd/system/${NAME}.service

build:
	@go build app.go

install:
	@cp ${NAME}.service ${SERVICE_SCRIPT}
	@mkdir -p ${INSTALL_LOCATION}
	@cp app ${INSTALL_LOCATION}
	@cp -r static ${INSTALL_LOCATION}
	@systemctl daemon-reload

uninstall:
	@rm -rf ${INSTALL_LOCATION}
	@rm -rf ${SERVICE_SCRIPT}
