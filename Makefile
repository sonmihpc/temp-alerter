VERSION := $(shell git describe --tags --always --match='v*')
version := $(shell echo $(VERSION) |grep -Eo '[0-9]+\.[0-9]+\.[0-9]')

run:
	go run main.go -c config.yaml
build:
	mkdir build && go build -o build/temp-alerter main.go
install:
	mkdir /etc/temp-alerter
	cp -r build/temp-alerter /usr/sbin/
	cp -r config.yaml /etc/temp-alerter/
	cp -r temp-alerter.service /usr/lib/systemd/system/
	systemctl start temp-alerter && systemctl enable temp-alerter
uninstall:
	systemctl stop temp-alerter && systemctl disable temp-alerter
	rm -f /usr/sbin/temp-alerter
	rm -rf /etc/temp-alerter
	rm -f /usr/lib/systemd/system/temp-alerter.service
rpm:
	rm -rf ~/rpmbuild/SOURCES/temp-alerter-$(version)*
	rm -f ~/rpmbuild/SPECS/temp-alerter.spec
	mkdir -p ~/rpmbuild/SOURCES/temp-alerter-$(version)
	go build -o  ~/rpmbuild/SOURCES/temp-alerter-$(version)/temp-alerter main.go
	cp -r temp-alerter.service ~/rpmbuild/SOURCES/temp-alerter-$(version)/
	cp -r config.yaml ~/rpmbuild/SOURCES/temp-alerter-$(version)/
	cd ~/rpmbuild/SOURCES;tar -cvzf temp-alerter-$(version).tar.gz temp-alerter-$(version)/;rm -rf temp-alerter-$(version)/
	cp -r temp-alerter.spec ~/rpmbuild/SPECS/
	sed -i 's/1.0.0/$(version)/g' ~/rpmbuild/SPECS/temp-alerter.spec
	rpmbuild -bb ~/rpmbuild/SPECS/temp-alerter.spec
clean:
	rm -rf ~/rpmbuild/SOURCES/temp-alerter-$(version)*
	rm -f ~/rpmbuild/SPECS/temp-alerter.spec
	rm -f ~/rpmbuild/RPMS/x86_64/temp-alerter-*