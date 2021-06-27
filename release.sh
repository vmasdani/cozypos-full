mkdir -p release/frontend &&\
xgo --targets=linux/amd64 . &&\
cp cozypos-full-linux-amd64 release &&\
cd frontend &&\
npm run build &&\
cp -r dist/* ../release/frontend &&\
cd ..
