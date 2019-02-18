# Parking-Lot

A simple application to operate various operations such as create, park, status or leave

## Run
support two modes of input. Default to interactive mode
1) Interactive command prompt based shell where commands can be typed in

```bash
cd $application_folder
go run main.go
```

2) Another one accepts a filename as a parameter at the command prompt and reads the commands from that file

```bash
cd $application_folder
go run main.go -mode=2 -filePath=examples/file_input.txt
```

## Build and run
Build and run the application using the following commands. 

```bash
cd $application_folder
go build -o parking_lot 
```
The above command builds the go application in the name of parking_lot. now we are supposed to run the same using following commands

To Run in interactive mode (Default mode) Try with the absolute path of the file

```bash
cd $application_folder
 ./parking_lot 
```

To Run in input file mode (Try with the absolute path of the file)
```bash
cd $application_folder
 ./parking_lot -mode=2 -filePath=/examples/file_input.txt
```

## Test
To test coverage run the following commands

```bash
cd $application_folder
go test ./... -cover
```

## Commands

* `create_parking_lot 6` - Creates a parking lot
* `park KA-01-HH-1234 White` - Park the vehicle to the nearest counter
* `leave 4` - leave a parking lot
* `status` - returns the list of vehicles in a parking lot
* `registration_numbers_for_cars_with_colour White` - return the list of vehicle number who's color is white
* `slot_numbers_for_cars_with_colour White` - return the list of slot number who's color is white
* `slot_number_for_registration_number KA-01-HH-3141` - return the slot numbers of a vehicles using registration number of a vehicle

For more examples refer examples folders