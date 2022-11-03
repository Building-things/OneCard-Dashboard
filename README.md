# OneCard-Dashboard
Displays UVic OneCard data nicely with some added data visualization.

## Running
make sure to npm install in the frontend directory first and pip install in the server directory

server (in server directory): go run main.go

frontend (in web directory): npm run dev

navigate to url provided by npm run dev

## Deployment
Make any deployment related changes on heroku-deploy
this command will re publish it to heroku
currently only me (rwn42) has permissions I will change that soon
subtree push --prefix server heroku heroku-deploy:master
:pizza: