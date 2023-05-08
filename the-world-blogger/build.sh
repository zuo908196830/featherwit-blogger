npm install
npm audit fix
cd ..
tar -czvf the-world-blogger.tar.gz the-world-blogger/
mv the-world-blogger.tar.gz the-world-blogger/
cd the-world-blogger/
docker build -t world-blogger:0.1 .