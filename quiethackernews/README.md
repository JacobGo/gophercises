# Gophercises: Quiet Hacker News

### Questions
* Do I need to lock my story cache for reads as well?
* Is there a more elegant way to house the cache than in the global space?

### Neat Features
* Added an additional cache for stories so that we remember popular stories that have been on top for a few expirations.

### Ideas
* Refresh the caches in the background + on start up.
* Refactor the overloaded handler and gross globals