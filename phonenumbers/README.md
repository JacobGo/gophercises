# Gophercises: Phone Number Normalizer

### Questions
* I wasn't too interested in the prompt to really explore this question in depth.
* Jon used a regex for normalizing -- is my loop similarly fast enough?
* Instead of using `WaitGroup` I used a channel so that I could output the message. Is there a way to do this without a 
  WaitGroup? I suppose I could just print my message inside the goroutine without issue.


### Neat Features
* GoLand's Database integration was smooth and quite powerful. Will have to consider this in the future.
* I parallelized the normalization and insertion of each phone number with goroutines and a channel.

### Ideas
* Not that interesting of a project, not sure what else I would do here.