package controller

type AppController struct {
	Auth        *AuthController
	Event       *EventController
	Reservation *ReservationController
}
