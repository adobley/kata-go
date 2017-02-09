For this exercise I'll be trying to implement the classic Mars Rover problem.

An initial rover should:
* Take a starting point (x, y) and a direction its facing (N,S,E,W)
* Take command to move forward
* Take command to move backward
* Take command to move right
* Take command to move left

Maybe we'll add some complexity afterwards.
* A world map
* Wrapping of movement around the map (Mars is spherical)
* Cliffs, pits, other dangers - How the simulation should react I don't know. Return an error on a command?