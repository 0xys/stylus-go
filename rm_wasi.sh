#!/bin/bash

cat outputs/mainh_tmp2.wat | grep -v preview1 > tmp.wat

LINE_FD_WRITE=$(cat tmp.wat | grep -n fd_write | cut -d ':' -f1)
START=$(echo "$LINE_FD_WRITE-4" | bc)
END=$(echo "$LINE_FD_WRITE+1" | bc)

echo "$START","$END"d

sed "$START","$END"d tmp.wat > outputs/mainh_tmp3.wat
diff --color outputs/mainh_tmp2.wat outputs/mainh_tmp3.wat
rm tmp.wat
