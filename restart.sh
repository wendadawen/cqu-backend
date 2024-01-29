date=$(date "+%Y-%m-%d_%H:%M.out"); cp nohup.out logs/log_$date; kill -9 $(pidof wechat_linux_linux); sleep 2s; nohup ./wechat_linux_linux &
