#Requirement :

Simple reservation system, restaurant can only accommodate 4 reservations per hour. Functions to add, view and cancel reservation 

1) Add reservation:
Input- customer name, date and time of reservation 
output- confirmation of the reservation if slot is available or not 

2) View -
Input - date
output - all resevations for the customer

3) Cancel
Input- same as add
output - confirmation on cancellation 

Constrains- restaurant operating hour - 10:00 am to 10:00 pm

Deliverable - No need of db. Utilize map or file

#Explaination
1. To Create the simple reservation system I firstly imported "fmt" package for the formatted input/output operations and "time" package for workingh with date and time.
2. Then I've given constants for the program those are openinghour, closehour ad maxresrvaions.
3. Then after the constants I Started defining the types of resrvation as strings.
4. After all the basic code I declaring the variables that holds the reservation data, here []Reservation is the slices holding a list of Reservation struct for each data.
5. From here I declared the functions addReservation, First, it checks if the time slot is available by calling isTimeSlotAvailable. If not, it prints an error message.
   If the slot is available, the function appends the new reservation to the list for that date and confirms it.
6. The next function viewReservations, This function allows viewing all reservations for a specific date. If no reservations exist for the date, it prints an appropriate message.
   Otherwise, it loops through the list of reservations for that date and prints the customer names and their reservation times.
7. And then cancelReservation function, This function allows canceling an existing reservation based on the customer’s name, date, and time. It loops through the reservations for the given date and removes the reservation that matches the provided details.
   If no matching reservation is found, it prints an error message.
8. isTimeSlotAvailable function,This function checks whether a specific time slot is available for a given date. It counts the number of existing reservations for that time slot and compares it to the maximum allowed reservations (4). If the count is less than the maximum, the time slot is available.
9. isValidTime function, This function validates whether the entered reservation time is within the restaurant's operating hours (between 10 AM and 10 PM). It parses the time string and checks whether the hour falls within the allowed range.
10. After declaring all the functions, the main program flow will start.
    -The main function provides a menu system where the user can choose an action:
    - Add Reservation: Prompts the user for the customer name, date, and time, validates the input, and calls addReservation.
    - View Reservations: Prompts the user for a date and calls viewReservations to display all bookings for that date.
    - Cancel Reservation: Prompts the user for the customer name, date, and time, then calls cancelReservation.
    - Exit: Ends the program.
  
