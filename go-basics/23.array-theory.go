/* -Array : array is a collection of similar data elements(arrays or strings) stored at contiguous memory location.

 -Contiguous Memory Locations := The elements are placed side-by-side in a single, unbroken block of RAM with 
  no gaps or scattered pieces in between.

 -Arrays in golng are fixed in length, means once they are declared and size is mentioned, 
  we cannot change the length of the array.
-Array in a golang has a pointer which points at the first elements in the array
int = 4 bits or 1 byte
-since memory is contiguous, we can calculate the address of any element we want to.
-it also has a property called length, which denotes the number of elements in the array,
-property called capacity,which denotes the number of elements it can contain. so in array length and capacity is same
-empty string is a zero value of a string [ - - - ], for int empty array [0 0 0]
- array is a collection of data are also numbered and it is called indices. the indices starts with 0 1 2 3 ...
  the first element has index 0 and the last element has index (l-1), l = length of the array
  we can access our elements using these indexes.