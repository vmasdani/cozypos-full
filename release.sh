mkdir -p release/frontend &&\
go build &&\
cp cozypos-full release &&\
cd frontend &&\
npm run build &&\
cp -r dist/* ../release/frontend &&\
cd ..
