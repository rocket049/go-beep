# 说明
本程序为 `go` 程序提供一个类似 `Windows API`：`Beep` 的电子音合成对象 `BeepPlayer`。使用方法：

1. 初始化：player,err := NewBeepPlayer()
2. 发音：err = player.Beep(freq, delay) ---- freq表示频率，delay表示以毫秒为单位的时长
3. 卸载：player.Close()

# 安装
1. （以`ubuntu`为例）安装`portaudio19-dev`软件包：`sudo apt install portaudio19-dev`
2. 运行：`go get -v gitee.com/rocket049/go-beep`
