# Temp-Alerter

由于业务中有部分用户需要监控机房环境温度，并在机房环境温度过高时发送报警邮件。因此基于该需求，并针对达林科技DL11B-MC-Dx系列温度传感器使用Go语言开发了该程序。该系列传感器采用USB转串口的接口，并使用标准的modbus RTU传输协议。


## 说明

由于手头上的传感器型号为DL11B-MC-D2，该设备有2个温度传感器，因此DL11B-MC-D1/D3用户需要自行测试。用户将传感器插到电脑或者服务器USB接口，然后修改配置启动程序即可。


## 安装

如果是centos 7用户或者rocky 9用户，可以直接从仓库的release中下载rpm文件直接安装即可：

```bash
  rpm -ivh temp-alerter-0.1.0-1.el7.x86_64.rpm
```

其他linux发行版的用户，如果有Go编译环境，可以直接克隆源码，从源码自行编译：

```bash
  git clone https://github.com/sonmihpc/temp-alerter.git
  cd temp-alerter
  make build
  make install
```

windows用户可以自行编译，或者从仓库的release中下载exe二进制可执行文件，然后自行创建一个配置文件。在CMD或者PowerShell中执行如下的命令即可运行：

```cmd
temp-alerter.exe -c config.yaml
```

## 配置/使用

用户安装软件之后，需要进行配置，以下是各个配置选项的说明。linux用户可以编辑 **/etc/temp-alerter/config.yaml**自行修改。

用户需要提供一个开通smtp功能的邮箱用于发送邮件使用，推荐使用163邮箱。

```
serial_port: "/dev/ttyUSB0"       # 温度传感器设备号，win用户改为对应的串口，例如COM1即可
sensor_num: 2                     # 温度传感器路数，DL11B-MC-D2有两路传感器，因此设置为2
smtp_host: smtp.163.com           # 邮箱smtp服务器，示例配置为163邮箱
smtp_port: 25                     # 邮箱smtp服务器端口，一般为25
smtp_email: test@163.com          # 邮箱地址
smtp_username: chia_sender        # 邮箱用户名，即邮箱地址@前面的内容
smtp_password: 123456             # 邮箱的smtp临时密码
mail_receiver:                    # 接收警告邮件的用户列表，示例中有两个用于接受邮件的邮箱
  - user1@163.com
  - user2@163.com
mail_delay: 30                    # 发送报警邮件后睡眠间隔，单位分钟，睡眠是为了防止邮件滥发
max_temp: 25                      # 温度警告最高阈值，单位摄氏度
min_temp: 0                       # 温度警告最低阈值，单位摄氏度
sample_interval: 2                # 温度取样间隔，单位秒
```

配置好之后启动服务即可：

```
systemctl start temp-alerter.service  # 启动服务
systemctl enable temp-alerter.service # 设置自启动
```

用户运行程序后，会在取样间隔时间看到标准输出实时的温度：

```
2024/01/18 20:10:01 environment temperature: [26.1 26.1] °C
2024/01/18 20:10:03 environment temperature: [26.1 26.1] °C
2024/01/18 20:10:05 environment temperature: [26.1 26.1] °C
2024/01/18 20:10:07 environment temperature: [26.1 26.1] °C
2024/01/18 20:10:09 environment temperature: [26.1 26.1] °C
2024/01/18 20:10:11 environment temperature: [26.1 26.1] °C
```

Linux用户也可以通过以下的命令查看到实时的环境温度输出：

```
tail -f /var/log/messages|grep temp-alerter
```

