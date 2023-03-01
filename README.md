# GoMongoCRUD
# TODO:
Add a usage  
Add a route to initialize and populate the database with the dataset, and creating in parallel the users' data files on the file system  
Update route: Compare the new data and the file system's one, regenerate the file system's one if it is different  
Add a switch between "no auth" and using godotenv in the usage  
When POSTing a new user, generate a correct timestamp for the "registered" key, then convert it to a string, thus having no problem with mongoDB not accepting the "xDATEx -01:00" format
