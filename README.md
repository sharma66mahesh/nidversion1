#### NID enrollment version 1

To run the project 
1. Go to web folder
2. Run export IMAGE_TAG = latest in terminal
3. Run startFabric.sh script
4. Install the node modules by running `npm install`
5. Register admin by running `node enrollAdmin.js`
6. Register user by running `node registerUser.js`
7. Start the webserver by running `node app.js`

Server is started and api can be called at localhost:4000/api/<function>

