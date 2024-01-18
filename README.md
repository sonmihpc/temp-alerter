# Temp-Alerter

Go实现的环境温度告警系统，该程序适用于达林科技DL11B-MC-Dx系列温度传感器。


## 说明

由于手头上的传感器型号为DL11B-MC-D2，该设备有2个温度传感器，因此DL11B-MC-D1/D3用户需要自行更改源码适配。


## 安装

如果是centos7用户或者rocky 9用户，可以直接从仓库的release中下载rpm文件直接安装即可：

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

用户安装软件之后，需要进行配置，以下是各个配置选项的说明。用户可以编辑 **/etc/temp-alerter/config.yaml**自行修改。

用户需要提供一个开通smtp功能的邮箱用于发送邮件使用，推荐使用163邮箱。

```
serial_port: "/dev/ttyUSB0"       # 温度传感器设备号，win用户改为对应的串口，例如COM1即可
smtp_host: smtp.163.com           # 邮箱smtp服务器，示例配置为163邮箱
smtp_port: 25                     # 邮箱smtp服务器端口，一般为25
smtp_email: test@163.com          # 邮箱地址
smtp_username: chia_sender        # 邮箱用户名，即邮箱地址@前面的内容
smtp_password: 123456             # 邮箱的smtp临时密码
mail_receiver:                    # 接收警告邮件的用户列表，示例中有两个用于接受邮件的邮箱
  - user1@163.com
  - user2@163.com
mail_delay: 30                    # 发送报警邮件后睡眠间隔，单位分钟，睡眠是为了防止邮件滥发
max_temp: 25                      # 温度警告最高阈值
min_temp: 0                       # 温度警告最低阈值
sample_interval: 2                # 温度取样间隔，单位秒
```

配置好之后启动服务即可：

```
systemctl start temp-alerter.service  # 启动服务
systemctl enable temp-alerter.service # 设置自启动
```

