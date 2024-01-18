Name:           temp-alerter
Version:        1.0.0
Release:        1%{?dist}
Summary:        Monitor the environment temperature.

License:        GPLv3+
URL:            https://github.com/sonmihpc/temp-alerter
Source0:       https://github.com/sonmihpc/temp-alerter/%{name}-%{version}.tar.gz

%description
Monitor the environment temperature.

%prep
%setup -q

%undefine _missing_build_ids_terminate_build
%global debug_package %{nil}

%build

%install
mkdir -p %{buildroot}/%{_sbindir}
mkdir -p %{buildroot}/%{_sysconfdir}/temp-alerter
mkdir -p %{buildroot}/%{_unitdir}
install -m 0700 temp-alerter %{buildroot}/%{_sbindir}/temp-alerter
install -m 0644 config.yaml %{buildroot}/%{_sysconfdir}/temp-alerter/config.yaml
install -m 0644 temp-alerter.service %{buildroot}/%{_unitdir}/temp-alerter.service

%files
%{_sbindir}/temp-alerter
%{_sysconfdir}/temp-alerter/config.yaml
%{_unitdir}/temp-alerter.service

%changelog
* Tue Jan 18 2024 root
- v1.0.0 release