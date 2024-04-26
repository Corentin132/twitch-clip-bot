#!/bin/bash          
url=$( cat ./url.txt )
cd ./clips
$(yt-dlp  $url )
cd ../
$(rm url.txt)
