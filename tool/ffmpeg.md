# ffmpeg

将当前文件夹所有png文件转为jpg文件，jpg文件更小，`~n`取文件名

`for %a in ("*.png") do ffmpeg -i "%a" "%~na.jpg"`

-t指定截取到某个时间

`ffmpeg -i in.mp4 -t 01:17:27 out.mp4`
