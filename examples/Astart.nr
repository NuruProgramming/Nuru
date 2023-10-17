/*############ A*(A-star) Algorithm ##############
             
             By @VictorKariuki

        https://github.com/VictorKariuki

################################################*/


// create a list of numbers
fanya list = unda(first,last,interval){
  fanya list = [first];
  fanya i = first + interval;
  wakati(i < last){
    list.sukuma(i)
    i+=interval;
  }
  rudisha list;
}

// Maths functions
// find the absolute value of a number
fanya abs_namba = unda(namba){
  kama(namba < 0){
    rudisha -1 * namba;
  }

  rudisha namba;
}

// square a number
fanya square = unda(n, i, j){
  fanya kati  = (i+j)/2;
  fanya mul = kati * kati;
  fanya abs_diff = abs_namba(mul-n);

  kama (mul == n || abs_diff < 0.00001){
    rudisha kati;
  }au kama(mul < n){
    rudisha square(n,kati,j)
  }au{
    rudisha square(n,i,kati)
  }
}

// find the square root of a number
fanya sqrt = unda(namba){
  kwa i ktk list(0,namba,1) {
    kama((i*i )== namba){
      rudisha i;
    }au kama ((i*i )> namba){
      rudisha square(namba,i-1,i)
    }
  }
}

// Main function
fanya aStar = unda(start, goal) {
  // Initialize the open and closed lists
  fanya openList = [start];
  fanya closedList = [];

  fanya reconstructPath = unda(node) {
    fanya path = [node];
    wakati (node["parent"]) {
      path = [node["parent"]] + path;
      node = node["parent"];
    }
    rudisha path;
  }

  fanya heuristic = unda(node1, node2) {
    // Calculate the Euclidean distance between the nodes' positions
    fanya dx = node1["x"] - node2["x"];
    fanya dy = node1["y"] - node2["y"];
    rudisha sqrt(dx * dx + dy * dy);
  }

  fanya findMinNode = unda(openList) {
    fanya i = 1;
    fanya minNode = openList[0];

    wakati (i < openList.idadi()) {
      fanya node = openList[i];
      kama (node["f"] < minNode["f"]) {
        minNode = node;
      }
      i++
    }

    rudisha minNode;
  }

  fanya removeNodeFromArray = unda(array, node) {
    fanya newArray = [];
    fanya i = 1;
    wakati (i < array.idadi()) {
      kama (array[i] != node) {
        newArray.sukuma(array[i]);
      }
      i++;
    }
    rudisha newArray;
  }

  fanya urefu = unda(node1, node2) {
    // Assume all edges have a cost of 1
    rudisha 1;
  }

  // Initialize the g and f scores of the starting node
  start["g"] = 0;
  start["f"] = start["g"] + heuristic(start, goal);

  

  // Start the search loop
  wakati (openList.idadi() > 0) {
    // Find the node with the lowest f score in the open list
    fanya current = findMinNode(openList);

    // Check kama the goal node has been reached
    kama (current == goal) {
      rudisha reconstructPath(current);
    }

    // Move the current node from the open to the closed list
    openList = removeNodeFromArray(openList, current);

    closedList.sukuma(current);

    // Explore the neighbors of the current node
    kwa neighbor ktk current["neighbors"] {
      // Skip neighbors that are in the closed list
      kama (neighbor ktk closedList) {
        endelea
      }

      // Calculate the tentative g score of the neighbor
      fanya tentativeG = start["g"] + urefu(current, neighbor);

      // Check kama the neighbor is in the open list
      fanya tentativeIsBetter = sikweli;
      kama (!(neighbor ktk openList)) {
        openList.sukuma(neighbor);
        tentativeIsBetter = kweli;
      } au kama (tentativeG < neighbor["g"]) {
        tentativeIsBetter = kweli;
      }

      // Update the neighbor's g score kama the tentative score is better
      kama (tentativeIsBetter) {
        neighbor["g"] = tentativeG;
        neighbor["f"] = neighbor["g"] + heuristic(neighbor, goal);
        neighbor["parent"] = current;
      }
    }
  }

  // kama the open list is empty, no path was found
  rudisha tupu;
}

// Define the nodes of the graph
fanya nodeA = { "x": 0, "y": 0, "neighbors": [] };
fanya nodeB = { "x": 1, "y": 2, "neighbors": [] };
fanya nodeC = { "x": 3, "y": 1, "neighbors": [] };
fanya nodeD = { "x": 4, "y": 3, "neighbors": [] };

// Define the edges between the nodes
nodeA["neighbors"] = [nodeB];
nodeB["neighbors"] = [nodeA, nodeC];
nodeC["neighbors"] = [nodeB, nodeD];
nodeD["neighbors"] = [nodeC];

// Call the A* function with the start and goal nodes and the heuristic and distance functions
//fanya path = aStar(nodeA, nodeC);

andika(nodeA);