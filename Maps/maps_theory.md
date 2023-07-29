## Theory about maps
- collection of unordered set of key value pairs
- maps declaration :
  var <map_name> map[<key_data_type>]<value_data_type>
  This will create a nil map and when we try to add values to it, compiler may throw runtime error
- another map declaration and initialization
  <map_name> := map <key_data_type>]<value_data_type>{<key-value-pairs>}
  example : codes := maptstringlstring{"en": "English", "fr": "French"}
- declaring a map using make function
  <map_name> :=make(map[<key_data_type>] <value_data_type>,<initial_capacity>)
- delete a key can be done using "delete" keyword 
  example : delete(<map_name>,<key>)
- to truncate a map we have two options
  - iterate over it and delete each key
  - assign it with new map
- only maps created/declared through "make" keyword can be assigned new value using map[<key] = value