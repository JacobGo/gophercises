# Gophercises: Task CLI

### Questions
* I didn't enjoy the `int` <--> `byte[]` conversions for the BoltDB keys.
* Still working on elegant error handling.

### Neat Features
* Added `Done` to Tasks. Represented Task struct as JSON in the db.
* I liked using cobra and BoltDB. Pretty nifty.

### Ideas
* Open a REPl to add interactively describe tasks.
* Save `tasks.db` in user's home directory so it can run from anywhere.
* Add Task removal
* Show completed tasks 